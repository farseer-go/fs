package timingWheel

import (
	"github.com/farseer-go/fs/flog"
	"github.com/farseer-go/fs/snowflake"
	"math"
	"sync"
	"time"
)

type wheelLevel = int // 时间轮的层数
type timeHand = int   // 第几个格子（桶）

// 多层时间轮，与时钟表的秒分时钟一个原理
// 秒钟走一圈，分钟走一格，以此类推
// 秒钟一圈时长，等于分钟的一格时长，以此类推
// 相比传统的多层，每层时长一样的算法，更科学。
type timingWheel struct {
	ticker        *time.Ticker            // 定时器
	duration      time.Duration           // 第一层时间每格的时长
	bucketsNum    int                     // 一圈的格子数量
	totalDuration time.Duration           // 第一层的时长 duration * bucketsNum
	onceStart     sync.Once               // 保证只执行一次
	timerQueue    chan *Timer             // 到达时间的任务，会立即放到此队列中
	clock         map[wheelLevel]timeHand // 时钟模型
	// 两组，第一组以时间轮的层数为KEY
	// 第二组为当前时间轮的层数的时间格做为KEY
	// 根据当前时间轮层数 + 时间格子 ，就能找出对应的Timer
	timeHandTimer map[wheelLevel]map[timeHand][]*Timer
	lock          *sync.RWMutex
}

// New 初始化
// duration 每个时间格子的时长
func New(interval time.Duration, bucketsNum int) *timingWheel {
	timeHandTimer := make(map[wheelLevel]map[timeHand][]*Timer)
	timeHandTimer[0] = make(map[timeHand][]*Timer)
	return &timingWheel{
		duration:      interval,
		bucketsNum:    bucketsNum,
		totalDuration: interval * time.Duration(bucketsNum),
		timerQueue:    make(chan *Timer, 1000),
		clock:         make(map[wheelLevel]timeHand),
		timeHandTimer: timeHandTimer,
		lock:          &sync.RWMutex{},
	}
}

// Start 开始运行
func (receiver *timingWheel) Start() {
	receiver.onceStart.Do(func() {
		receiver.ticker = time.NewTicker(receiver.duration)
		// 启动时间轮盘
		go receiver.turning()
	})
}

// Add 添加一个定时任务
func (receiver *timingWheel) Add(d time.Duration) *Timer {
	planAt := time.Now().Add(d)
	planDuration := d.String()

	// 公式原理：receiver.totalDuration = 第一层的时长
	// 每一层的一格时长：receiver.totalDuration^level	math.Pow(receiver.totalDuration, level)
	// 每一层的一圈时长：receiver.totalDuration^level+1		math.Pow(receiver.totalDuration, level+1)

	// 计算出第几层 根据d、receiver.bucketDuration开根号，会得出第几层
	level := receiver.getLevel(d)
	// 得到level层一格的时长（也相当于level-1一圈的时长）
	curLevelTimeHandDuration := receiver.getLevelTimeHandDuration(level)
	if level > 0 {
		// 减去上一层的总时长，得出当前层的相对时长
		d -= curLevelTimeHandDuration
	}

	// 得到level层所在的指针
	curLevelTimeHand := timeHand(d / curLevelTimeHandDuration)

	// 要算上level已经走的指针
	curLevelTimeHand += receiver.clock[level]
	// 当前走的指针大于每一层的格子数量，则要跳到下一层
	if curLevelTimeHand >= receiver.bucketsNum {
		level++
		// 跳到下一层时，指针要按下一层的时长重新计算
		d -= time.Duration(receiver.bucketsNum-receiver.clock[level]) * receiver.duration
		curLevelTimeHandDuration = receiver.getLevelTimeHandDuration(level)
		curLevelTimeHand = timeHand(d / curLevelTimeHandDuration)
	}

	// 用于更精确控制时间
	level, curLevelTimeHand = receiver.rewind(d, curLevelTimeHandDuration, level, curLevelTimeHand)

	t := &Timer{
		Id:       snowflake.GenerateId(),
		C:        make(chan time.Time, 1),
		duration: d,
		PlanAt:   planAt,
	}
	// 在同一格，说明需要立即执行
	if (level == 0 && curLevelTimeHand == receiver.clock[0]) || d < 0 {
		go receiver.popTimer(t)
		return t
	}
	flog.Debugf("添加时间(%d):+%s %s level：%d 指针：%d，当前指针：%d", t.Id, planDuration, planAt.Format("15:04:05.000"), level, curLevelTimeHand, receiver.clock[0])

	receiver.lock.Lock()
	defer receiver.lock.Unlock()
	_, exists := receiver.timeHandTimer[level]
	// 如果所在的层没有初始化过，则先初始化
	if !exists {
		receiver.timeHandTimer[level] = make(map[timeHand][]*Timer)
	}

	// 找到对应层的指针
	timers := receiver.timeHandTimer[level][curLevelTimeHand]
	timers = append(timers, t)
	receiver.timeHandTimer[level][curLevelTimeHand] = timers
	return t
}

// AddTime 添加一个定时任务
func (receiver *timingWheel) AddTime(t time.Time) *Timer {
	return receiver.Add(t.Sub(time.Now()))
}

// AddPrecision 添加一个定时任务（高精度）
func (receiver *timingWheel) AddPrecision(d time.Duration) *Timer {
	timer := receiver.Add(d)
	timer.isPrecision = true
	return timer
}

// AddTimePrecision 添加一个定时任务（高精度）
func (receiver *timingWheel) AddTimePrecision(t time.Time) *Timer {
	timer := receiver.Add(t.Sub(time.Now()))
	timer.isPrecision = true
	return timer
}

// 时间轮开始转动
func (receiver *timingWheel) turning() {
	for {
		go func(tHand timeHand) {
			flog.Debugf("当前指针，%d.%d.%d", receiver.clock[2], receiver.clock[1], tHand)
			// 取出当前格子的任务
			timers := receiver.timeHandTimer[0][tHand]
			if len(timers) > 0 {
				// 需要提前删除，否则会与降级任务冲突
				receiver.lock.Lock()
				delete(receiver.timeHandTimer[0], tHand)
				receiver.lock.Unlock()
			}
			if len(timers) > 1 {
				receiver.order(timers)
			}

			for i := 0; i < len(timers); i++ {
				if !timers[i].isStop {
					flog.Debugf("休眠时间(%d):+%s %d us", timers[i].Id, timers[i].PlanAt.Format("15:04:05.000"), timers[i].PlanAt.Sub(time.Now()).Microseconds())
					receiver.popTimer(timers[i])
				}
			}
		}(receiver.clock[0])

		// 每timingWheel.duration 转动一次
		<-receiver.ticker.C
		// 时间指针向前一格
		receiver.turningNextLevel(0)
	}
}

// 下一层指针+1
func (receiver *timingWheel) turningNextLevel(level wheelLevel) {
	// 时间指针向前一格
	receiver.clock[level]++
	// 超出最大格子时
	if receiver.clock[level] >= receiver.bucketsNum {
		// 指针重新回到0
		receiver.clock[level] = 0
		// 先降级任务
		receiver.timerDowngrade(level + 1)
		// 下一层指针+1
		receiver.turningNextLevel(level + 1)
	}
}

// 任务降级，把level层的任务降到level-1层
func (receiver *timingWheel) timerDowngrade(level wheelLevel) {
	flog.Debugf("第%d层降级任务", level)
	timesHand, exists := receiver.timeHandTimer[level]
	if exists {
		clockHand := receiver.clock[level]
		timers := timesHand[clockHand]
		// 得到level-1层一格的时长
		preLevelTimeHandDuration := receiver.getLevelTimeHandDuration(level - 1)
		// 得到level层一格的时长
		curLevelTimeHandDuration := receiver.getLevelTimeHandDuration(level)
		for i := 0; i < len(timers); i++ {
			// 如果不是第0格，则要回退一格
			if clockHand > 0 {
				timers[i].duration -= curLevelTimeHandDuration
			}
			// 得到level-1层所在的指针
			nextLevelTimeHand := timeHand(timers[i].duration / preLevelTimeHandDuration)
			if nextLevelTimeHand == receiver.bucketsNum {
				nextLevelTimeHand--
			}

			// 用于更精确控制时间
			level, nextLevelTimeHand = receiver.rewind(timers[i].duration, preLevelTimeHandDuration, level, nextLevelTimeHand)
			// 将上层任务降级到下层
			receiver.timeHandTimer[level-1][nextLevelTimeHand] = append(receiver.timeHandTimer[level-1][nextLevelTimeHand], timers[i])
		}
		receiver.lock.Lock()
		defer receiver.lock.Unlock()
		// 把当前这一层这一格的任务移除
		delete(receiver.timeHandTimer[level], clockHand)
		if len(receiver.timeHandTimer[level]) == 0 {
			delete(receiver.timeHandTimer, level)
		}
	}
}

// 计算出第几层 将第一层的时长开根号，会得出第几层
func (receiver *timingWheel) getLevel(d time.Duration) int {
	count := 0
	for d/receiver.totalDuration >= 1 {
		d /= receiver.totalDuration
		count++
	}
	return count
}

// 将到达时间的任务推送给C
func (receiver *timingWheel) popTimer(timer *Timer) {
	microseconds := timer.PlanAt.Sub(time.Now()).Microseconds()
	flog.Debugf("休眠时间(%d):+%s %d us", timer.Id, timer.PlanAt.Format("15:04:05.000"), timer.PlanAt.Sub(time.Now()).Microseconds())
	if microseconds > 0 {
		// 使用精确的时间
		if timer.isPrecision {
			timer.precision()
		}
		time.Sleep(timer.PlanAt.Sub(time.Now()))
	}
	timer.C <- time.Now()
	flog.Debugf("推送时间(%d):+%s 当前指针：%d 精确度：%v", timer.Id, timer.PlanAt.Format("15:04:05.000"), receiver.clock[0], timer.isPrecision)
}

// 排序任务，按时间从小到大
func (receiver *timingWheel) order(lst []*Timer) {
	// 首先拿数组第0个出来做为左边值
	for leftIndex := 0; leftIndex < len(lst); leftIndex++ {
		// 拿这个值与后面的值作比较
		leftValue := lst[leftIndex].duration

		// 再拿出左边值索引后面的值一一对比
		for rightIndex := leftIndex + 1; rightIndex < len(lst); rightIndex++ {
			rightValue := lst[rightIndex].duration // 这个就是后面的值，会陆续跟数组后面的值做比较
			rightItem := lst[rightIndex]

			// 后面的值比前面的值小，说明要交换数据
			if leftValue >= rightValue {
				// 开始交换数据，先从后面交换到前面
				for swapIndex := rightIndex; swapIndex > leftIndex; swapIndex-- {
					lst[swapIndex] = lst[swapIndex-1]
				}
				lst[leftIndex] = rightItem
				leftValue = lst[leftIndex].duration
			}
		}
	}
}

// 获取对应层的一格时长
func (receiver *timingWheel) getLevelTimeHandDuration(level wheelLevel) time.Duration {
	if level == 0 {
		return receiver.duration
	}
	return time.Duration(math.Pow(float64(receiver.totalDuration), float64(level)))
}

// 当剩余时间小于3ms时，回退一格（用于更精确控制时间）
func (receiver *timingWheel) rewind(d time.Duration, curLevelTimeHandDuration time.Duration, level wheelLevel, nextLevelTimeHand timeHand) (wheelLevel, timeHand) {
	// 在当前时间格中，超出的时间
	remainingDuration := d % curLevelTimeHandDuration
	// 如果小于3ms，则退一格（用于更精确控制时间），并且不能是第一层，第一格，否则无法再退了
	if remainingDuration <= 3*time.Millisecond && (level > 0 || nextLevelTimeHand > 0) {
		nextLevelTimeHand--
		// 当前时间格无法再往后退，只能下降一层
		if nextLevelTimeHand == -1 {
			nextLevelTimeHand = receiver.bucketsNum - 1
			level--
		}
	}
	return level, nextLevelTimeHand
}
