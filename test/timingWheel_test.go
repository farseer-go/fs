package test

import (
	"github.com/farseer-go/fs/timingWheel"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestTimingWheel(t *testing.T) {
	//	fs.Initialize[modules.FarseerKernelModule]("unit test")
	/*
			0： 100 * 60 = 6000 ms						0.12s
			1： 100 * 60 * 60 = 1440 ms					1.44s
			2： 100 * 60 * 60 * 60 * 17280 ms			17.28
			3： 100 * 60 * 60 * 60 * 60 * 207360 ms		207.36

		第0层： 10 20 30 40 50 60 70 80 90 100 110 120
		第1层： 120 240 360 480 600 720 840 960 1080 1200 1320 1440
				  0秒 123：第1层，第2格，3
				  78秒 123：第1层，第2格，3
					201秒
	*/
	timingWheel.NewDefault(10*time.Millisecond, 12)
	timingWheel.Start()
	<-timingWheel.AddPrecision(-1 * time.Millisecond).C
	<-timingWheel.AddPrecision(1 * time.Millisecond).C
	time.Sleep(131 * time.Millisecond)
	timingWheel.AddPrecision(123 * time.Millisecond)
	timer4 := timingWheel.AddPrecision(122 * time.Millisecond)
	timer5 := timingWheel.AddPrecision(1003 * time.Millisecond)
	timer6 := timingWheel.AddPrecision(1443 * time.Millisecond)
	timingWheel.AddPrecision(-35 * time.Millisecond).Stop()
	timingWheel.Add(102 * time.Millisecond)
	timingWheel.AddTime(time.Now().Add(1304 * time.Millisecond))
	timer1 := timingWheel.AddPrecision(12 * time.Millisecond)
	assert.WithinDuration(t, timer1.PlanAt, <-timer1.C, time.Millisecond)

	timer2 := timingWheel.AddPrecision(1102 * time.Millisecond)
	timer3 := timingWheel.AddTimePrecision(time.Now().Add(1203 * time.Millisecond))
	assert.WithinDuration(t, timer2.PlanAt, <-timer2.C, time.Millisecond)
	assert.WithinDuration(t, timer3.PlanAt, <-timer3.C, time.Millisecond)
	assert.WithinDuration(t, timer4.PlanAt, <-timer4.C, time.Millisecond)
	assert.WithinDuration(t, timer5.PlanAt, <-timer5.C, time.Millisecond)
	assert.WithinDuration(t, timer6.PlanAt, <-timer6.C, time.Millisecond)
}
