package test

import (
	"fmt"
	"github.com/farseer-go/fs"
	"github.com/farseer-go/fs/modules"
	"github.com/farseer-go/fs/timingWheel"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestTimingWheel(t *testing.T) {
	fs.Initialize[modules.FarseerKernelModule]("unit test")
	tw := timingWheel.New(500*time.Millisecond, 60)
	tw.Start()
	tw.AddPrecision(-305 * time.Millisecond)
	//time.Sleep(300 * time.Millisecond)
	// 添加一个505ms以后的任务
	timer1 := tw.AddPrecision(305 * time.Millisecond)
	// 添加一个1843ms以后的任务
	timer2 := tw.Add(102 * time.Millisecond)
	timer2.Stop()

	timer3 := tw.Add(599 * time.Millisecond)

	timer4 := tw.AddPrecision(1102 * time.Millisecond)
	timer5 := tw.AddPrecision(1303 * time.Millisecond)
	timer6 := tw.AddPrecision(1304 * time.Millisecond)
	assert.Equal(t, (<-timer1.C).Format("15:04:05.000"), time.Now().Format("15:04:05.000"))

	fmt.Printf("timer3计划时间：%s\n", (<-timer3.C).Format("15:04:05.000"))
	fmt.Printf("现在时间：%s\n\n", time.Now().Format("15:04:05.000"))

	assert.Equal(t, (<-timer4.C).Format("15:04:05.000"), time.Now().Format("15:04:05.000"))
	assert.Equal(t, (<-timer5.C).Format("15:04:05.000"), time.Now().Format("15:04:05.000"))
	assert.Equal(t, (<-timer6.C).Format("15:04:05.000"), time.Now().Format("15:04:05.000"))
}
