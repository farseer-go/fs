package timingWheel

import (
	"context"
	"math"
	"sort"
	"sync"
	"time"

	"github.com/farseer-go/fs/sonyflake"
)

// OpOption 选项
type OpOption func(*Op)

// Op 选项
type Op struct {
	IsPrecision bool // 是否使用高精度时间格
}
type wheelLevel = int // 时间轮的层数
type timeHand = int   // 第几个格子（桶）

// 多层时间轮，与时钟表的秒分时钟一个原理
// 秒钟走一圈，分钟走一格，以此类推
// 秒钟一圈时长，等于分钟的一格时长，以此类推
// 相比传统的多层，每层时长一样的算法，更科学。
type timingWheel struct {
	ticker        *time.Ticker    // 定时器
	duration      []time.Duration // 每一层时间每格的时长
	bucketsNum    int             // 一圈的格子数量
	totalDuration time.Duration   // 第一层的时长 duration * bucketsNum
	onceStart     sync.Once       // 保证只执行一次
	clock         []timeHand      // 时钟模型（数组索引 = wheelLevel）
	clockLock     *sync.RWMutex   // 锁
	// 数组索引 = wheelLevel
	// 当前时间轮的层数的时间格做为MAP KEY
	// 根据当前时间轮层数 + 时间格子 ，就能找出对应的Timer
	timeHandTimer []map[timeHand][]*Timer
	timerLock     *sync.RWMutex // 锁
	ctx           context.Context
	cancel        context.CancelFunc
	stopped       chan struct{} // 停止完成的信号
}

// New 初始化
// interval 每个时间格子的时长，建议设置为100ms
func New(interval time.Duration, bucketsNum int) *timingWheel {
	return &timingWheel{
		duration:      []time.Duration{interval},
		bucketsNum:    bucketsNum,
		totalDuration: interval * time.Duration(bucketsNum),
		timerLock:     &sync.RWMutex{},
		clockLock:     &sync.RWMutex{},
		timeHandTimer: []map[timeHand][]*Timer{make(map[timeHand][]*Timer)},
		clock:         []timeHand{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	}
}

// Start 开始运行
func (receiver *timingWheel) Start() {
	receiver.onceStart.Do(func() {
		receiver.initLevelTimeHandDuration(10)
		receiver.ticker = time.NewTicker(receiver.duration[0])
		receiver.ctx, receiver.cancel = context.WithCancel(context.Background())
		receiver.stopped = make(chan struct{})

		// 启动时间轮盘
		go receiver.turning()
	})
}

// Stop 停止时间轮
func (receiver *timingWheel) Stop() {
	if receiver.cancel != nil {
		receiver.cancel()
		<-receiver.stopped
	}
}

// Add 添加一个定时任务
func (receiver *timingWheel) Add(d time.Duration, opts ...OpOption) *Timer {
	op := &Op{}
	for _, opt := range opts {
		opt(op)
	}

	t := &Timer{
		Id:                sonyflake.GenerateId(),
		C:                 make(chan time.Time, 1),
		duration:          d,
		remainingDuration: 0,
		PlanAt:            time.Now().Add(d),
		isPrecision:       op.IsPrecision,
	}

	// 计算第几层第几格、剩余时长（需要持有 clockLock，因为 rewind 会读写 clock）
	receiver.clockLock.Lock()
	level, curLevelTimeHand, remainingDuration := receiver.rewind(d)
	curLevelTimeHand += receiver.clock[level]

	// 超出了桶，则跳到下一级
	for curLevelTimeHand >= receiver.bucketsNum {
		remainingDuration += time.Duration(curLevelTimeHand-receiver.bucketsNum) * receiver.duration[level]
		level++
		curLevelTimeHand = receiver.clock[level] + 1
	}

	t.remainingDuration = remainingDuration
	immediateRun := t.remainingDuration < 0 || t.PlanAt.Before(time.Now())
	sameSlot := level == 0 && curLevelTimeHand == receiver.clock[0]
	receiver.clockLock.Unlock()

	// 晚于当前时间，需要立即推送
	if immediateRun {
		receiver.popTimer(t)
		return t
	}

	// 在同一格，说明需要立即执行
	if sameSlot {
		go receiver.popTimer(t)
		return t
	}

	// 初始化所在层的任务队列
	receiver.initTimeHandTimer(level)

	receiver.timerLock.Lock()
	defer receiver.timerLock.Unlock()
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
	return receiver.Add(d, func(op *Op) {
		op.IsPrecision = true
	})
}

// AddTimePrecision 添加一个定时任务（高精度）
func (receiver *timingWheel) AddTimePrecision(t time.Time) *Timer {
	return receiver.Add(t.Sub(time.Now()), func(op *Op) {
		op.IsPrecision = true
	})
}

// 时间轮开始转动
func (receiver *timingWheel) turning() {
	defer close(receiver.stopped)
	for {
		select {
		case <-receiver.ctx.Done():
			receiver.ticker.Stop()
			return
		default:
		}

		tHand := receiver.clock[0]
		receiver.timerLock.Lock()
		// 取出当前格子的任务
		timers := receiver.timeHandTimer[0][tHand]
		if len(timers) > 0 {
			// 需要提前删除，否则会与降级任务冲突
			delete(receiver.timeHandTimer[0], tHand)
		}
		receiver.timerLock.Unlock()

		if len(timers) > 1 {
			receiver.order(timers)
		}

		go func(timers []*Timer) {
			for i := 0; i < len(timers); i++ {
				if !timers[i].isStopped() {
					// 每个 timer 独立 goroutine，避免阻塞连锁
					go receiver.popTimer(timers[i])
				}
			}
		}(timers)

		// 每timingWheel.duration 转动一次
		<-receiver.ticker.C

		receiver.clockLock.Lock()
		// 时间指针向前一格
		receiver.turningNextLevel(0)
		receiver.clockLock.Unlock()
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

		for level+1 >= len(receiver.clock) {
			receiver.clock = append(receiver.clock, 0)
		}

		// 下一层指针+1
		receiver.turningNextLevel(level + 1)

		// 先降级任务
		receiver.timerDowngrade(level + 1)
		receiver.initTimeHandTimer(level + 1)
	}
}

// 任务降级，把level层的任务降到level-1层
func (receiver *timingWheel) timerDowngrade(curLevel wheelLevel) {
	clockHand := receiver.clock[curLevel]
	//flog.Debugf("任务第%d层第%d格 降级", curLevel, clockHand)
	if curLevel < len(receiver.timeHandTimer) {
		receiver.timerLock.Lock()
		defer receiver.timerLock.Unlock()

		timers := receiver.timeHandTimer[curLevel][clockHand]
		for i := 0; i < len(timers); i++ {
			// 用于更精确控制时间
			level, curLevelTimeHand, remainingDuration := receiver.rewind(timers[i].remainingDuration)
			timers[i].remainingDuration = remainingDuration

			// 将上层任务降级到下层
			receiver.timeHandTimer[level][curLevelTimeHand] = append(receiver.timeHandTimer[level][curLevelTimeHand], timers[i])
			//flog.Debugf("任务(%d)，放到第%d层第%d格 %s", timers[i].Id, level, curLevelTimeHand, timers[i].PlanAt.Format("15:04:05.000"))
		}

		// 把当前这一层这一格的任务移除
		delete(receiver.timeHandTimer[curLevel], clockHand)
	}
}

// 计算出第几层 将第一层的时长开根号，会得出第几层
func (receiver *timingWheel) getLevel(d time.Duration) int {
	//	0： 10 * 12 = 120 ms						0.12s
	//	1： 10 * 12 * 12 = 1440 ms					1.44s
	//	2： 10 * 12 * 12 * 12 * 17280 ms			17.28
	// 503
	count := 0
	for d > receiver.totalDuration {
		d /= time.Duration(receiver.bucketsNum)
		count++
	}
	return count
}

// 将到达时间的任务推送给C
func (receiver *timingWheel) popTimer(timer *Timer) {
	remaining := timer.PlanAt.Sub(time.Now())
	if remaining > 0 {
		if timer.isPrecision {
			// 高精度模式：先 Sleep 大部分时间，最后 busy-wait
			timer.precision()
			// precision 已经 spin 到目标时间，不再调用 time.Sleep
		} else {
			// 普通模式：直接 Sleep
			time.Sleep(remaining)
		}
	}
	timer.C <- time.Now()
}

// 排序任务，按时间从小到大
func (receiver *timingWheel) order(lst []*Timer) {
	sort.Slice(lst, func(i, j int) bool {
		return lst[i].duration < lst[j].duration
	})
}

// 初始化这一层的一格时长
func (receiver *timingWheel) initLevelTimeHandDuration(level wheelLevel) {
	// 为了优化，将下一层的时长提前计算好
	for level >= len(receiver.duration) {
		curLevelTimeHandDuration := receiver.duration[0] * time.Duration(math.Pow(float64(receiver.bucketsNum), float64(len(receiver.duration))))
		receiver.duration = append(receiver.duration, curLevelTimeHandDuration)
	}
}

// 计算第几层第几格、剩余时长
func (receiver *timingWheel) rewind(duration time.Duration) (wheelLevel, timeHand, time.Duration) {
	// 计算出第几层
	level := receiver.getLevel(duration)

	// 初始化这一层的一格时长
	receiver.initLevelTimeHandDuration(level)

	// 时间 / 当前一格的时间，就能算出第几格（向上取整）
	curLevelTimeHand := timeHand(math.Ceil(float64(duration)/float64(receiver.duration[level]))) - 1
	if curLevelTimeHand == -1 {
		curLevelTimeHand = 0
	}

	// 得到在level层的剩余时长 duration - (receiver.duration[level] * time.Duration(curLevelTimeHand))
	remainingDuration := duration % receiver.duration[level]

	// 如果小于3ms，则退一格（用于更精确控制时间），并且不能是第一层，第一格，否则无法再退了
	if remainingDuration <= 3*time.Millisecond && (level > 0 || curLevelTimeHand > 1) {
		curLevelTimeHand--

		// 第1层+的第0格，是不会被运行到的，因为默认就是第0格
		if curLevelTimeHand == 0 {
			curLevelTimeHand = receiver.bucketsNum - 1
			level--
		}
		remainingDuration += receiver.duration[level]
	}

	// 这里+1，是因为在后面执行的时候，有可能会跳到下一层
	// 注意：调用方需保证已持有 clockLock
	for level+1 >= len(receiver.clock) {
		receiver.clock = append(receiver.clock, 0)
	}

	return level, curLevelTimeHand, remainingDuration
}

// 初始化所在层的任务队列
func (receiver *timingWheel) initTimeHandTimer(level wheelLevel) {
	receiver.timerLock.Lock()
	defer receiver.timerLock.Unlock()
	// 如果所在的层没有初始化过，则先初始化
	for level >= len(receiver.timeHandTimer) {
		receiver.timeHandTimer = append(receiver.timeHandTimer, make(map[timeHand][]*Timer))
	}
}
