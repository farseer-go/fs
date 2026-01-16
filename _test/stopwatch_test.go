package test

import (
	"fmt"
	"testing"
	"time"

	"github.com/farseer-go/fs/stopwatch"
	"github.com/stretchr/testify/assert"
)

func TestStopwatch(t *testing.T) {
	sw := stopwatch.New()
	assert.Equal(t, time.Duration(0), sw.ElapsedDuration())
	sw.Start()
	time.Sleep(time.Millisecond)
	sw.Stop()
	assert.Equal(t, int64(1), sw.ElapsedDuration().Milliseconds())
	sw.Start()
	time.Sleep(time.Millisecond * 2)
	fmt.Println(sw.GetMillisecondsText())
	fmt.Println(sw.GetMicrosecondsText())
	fmt.Println(sw.GetNanosecondsText())
	sw.Restart()
	assert.LessOrEqual(t, int64(0), sw.ElapsedDuration().Milliseconds())
	time.Sleep(time.Millisecond * 300)
}
