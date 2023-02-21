package test

import (
	"github.com/farseer-go/fs"
	"github.com/farseer-go/fs/flog"
	"github.com/farseer-go/fs/modules"
	"github.com/farseer-go/fs/timingWheel"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestTimingWheel(t *testing.T) {
	fs.Initialize[modules.FarseerKernelModule]("unit test")
	tw := timingWheel.New(50*time.Millisecond, 10)
	tw.Start()
	timer4 := tw.AddPrecision(503 * time.Millisecond)
	timer5 := tw.AddPrecision(1003 * time.Millisecond)
	tw.AddPrecision(-35 * time.Millisecond).Stop()
	tw.Add(102 * time.Millisecond)
	tw.AddTime(time.Now().Add(1304 * time.Millisecond))
	timer1 := tw.AddPrecision(10 * time.Millisecond)
	flog.Info("timer1------------------------------------------")
	assert.Equal(t, timer1.PlanAt.Format("15:04:05.000"), (<-timer1.C).Format("15:04:05.000"))

	timer2 := tw.AddPrecision(1102 * time.Millisecond)
	timer3 := tw.AddTimePrecision(time.Now().Add(1203 * time.Millisecond))
	assert.WithinDuration(t, timer2.PlanAt, <-timer2.C, time.Millisecond)
	assert.WithinDuration(t, timer3.PlanAt, <-timer3.C, time.Millisecond)
	assert.WithinDuration(t, timer4.PlanAt, <-timer4.C, time.Millisecond)
	assert.WithinDuration(t, timer5.PlanAt, <-timer5.C, time.Millisecond)
}
