package timingWheel

import (
	"github.com/farseer-go/fs/flog"
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
	}
}

// Start 开始运行
func (receiver *timingWheel) Start() {
	receiver.onceStart.Do(func() {
		receiver.ticker = time.NewTicker(receiver.duration)
		//receiver.clock[0] = 0
		// 启动时间轮盘
		go receiver.turning()
		// 启动获取到达时间的任务
		go receiver.popTimer()
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
	var curLevelTimeHandDuration time.Duration
	if level == 0 {
		curLevelTimeHandDuration = receiver.duration
	} else {
		curLevelTimeHandDuration = time.Duration(math.Pow(float64(receiver.totalDuration), float64(level)))
		// 减去上一层的总时长，得出当前层的相对时长
		d -= curLevelTimeHandDuration
	}

	// 得到level层所在的指针
	curLevelTimeHand := timeHand(d / curLevelTimeHandDuration)

	// 如果是第0层，则还要算上当前已经走的指针
	if level == 0 {
		// 要算上当前已经走的指针
		curLevelTimeHand += receiver.clock[0]
		// 当前走的指针大于每一层的格子数量，则要跳到下一层
		if curLevelTimeHand >= receiver.bucketsNum {
			level++
			// 跳到下一层时，指针要按下一层的时长重新计算
			d -= time.Duration(receiver.bucketsNum-receiver.clock[0]) * receiver.duration
			curLevelTimeHandDuration = time.Duration(math.Pow(float64(receiver.totalDuration), float64(level)))
			curLevelTimeHand = timeHand(d / curLevelTimeHandDuration)
			//curLevelTimeHand -= receiver.bucketsNum
		}
	}

	flog.Debugf("添加时间:+%s %s level：%d 指针：%d", planDuration, planAt.Format("15:04:05.000"), level, curLevelTimeHand)
	t := &Timer{
		C:        make(chan time.Time),
		duration: d,
		planAt:   planAt,
	}

	_, exists := receiver.timeHandTimer[level]
	// 如果所在的层没有初始化过，则先初始化
	if !exists {
		receiver.timeHandTimer[level] = make(map[timeHand][]*Timer)
		for i := 0; i < receiver.bucketsNum; i++ {
			receiver.timeHandTimer[level][i] = []*Timer{}
		}
	}

	// 找到对应层的指针
	timers := receiver.timeHandTimer[level][curLevelTimeHand]
	timers = append(timers, t)
	receiver.timeHandTimer[level][curLevelTimeHand] = timers
	return t
}

func (receiver *timingWheel) AddPrecision(d time.Duration) *Timer {
	timer := receiver.Add(d)
	timer.isPrecision = true
	return timer
}

// 时间轮开始转动
func (receiver *timingWheel) turning() {
	for {
		// 每timingWheel.duration 转动一次
		<-receiver.ticker.C
		//fmt.Printf("ticker:%s 指针：%d 任务：%d个\n", time.Now().Format("15:04:05.000"), receiver.clock[0], len(receiver.timeHandTimer[0][receiver.clock[0]]))

		// 取出当前格子的任务
		timers := receiver.timeHandTimer[0][receiver.clock[0]]
		for i := 0; i < len(timers); i++ {
			if !timers[i].isStop {
				receiver.timerQueue <- timers[i]
			}
		}
		delete(receiver.timeHandTimer[0], receiver.clock[0])

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
	timesHand, exists := receiver.timeHandTimer[level]
	if exists {
		clockHand := receiver.clock[level]
		timers := timesHand[clockHand]
		// 得到level层一格的时长（也相当于level-1一圈的时长）
		curLevelTimeHandDuration := time.Duration(math.Pow(float64(receiver.totalDuration), float64(level)))
		for i := 0; i < len(timers); i++ {
			// 得到level-1层所在的指针
			nextLevelTimeHand := timeHand(timers[i].duration / curLevelTimeHandDuration)

			// 将上层任务降级到下层
			receiver.timeHandTimer[level-1][nextLevelTimeHand] = append(receiver.timeHandTimer[level-1][nextLevelTimeHand], timers[i])
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
func (receiver *timingWheel) popTimer() {
	for timer := range receiver.timerQueue {
		// 使用精确的时间
		if timer.isPrecision {
			go timer.precision()
			continue
		}

		// 留出5ms做最后精确控制
		milliseconds := timer.planAt.Sub(time.Now()).Milliseconds()
		time.Sleep(time.Duration(milliseconds) * time.Millisecond)
		timer.C <- timer.planAt
	}
}
