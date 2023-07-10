package test

import (
	"github.com/farseer-go/fs/flog2"
	"github.com/farseer-go/fs/stopwatch"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestStopwatch(t *testing.T) {
	sw := stopwatch.New()
	assert.Equal(t, int64(0), sw.ElapsedMilliseconds())
	assert.Equal(t, int64(0), sw.ElapsedMicroseconds())
	assert.Equal(t, int64(0), sw.ElapsedNanoseconds())
	sw.Start()
	time.Sleep(time.Millisecond)
	sw.Stop()
	assert.Equal(t, int64(1), sw.ElapsedMilliseconds())
	sw.Start()
	time.Sleep(time.Millisecond * 2)
	flog2.Println(sw.GetMillisecondsText())
	flog2.Println(sw.GetMicrosecondsText())
	flog2.Println(sw.GetNanosecondsText())
	sw.Restart()
	assert.LessOrEqual(t, int64(0), sw.ElapsedMilliseconds())
	time.Sleep(time.Millisecond * 300)
}
