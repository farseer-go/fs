package test

import (
	"github.com/farseer-go/fs/parse"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsEqual(t *testing.T) {
	assert.False(t, parse.IsEqual(struct{}{}, struct{}{}))
	assert.True(t, parse.IsEqual(true, true))
	assert.False(t, parse.IsEqual(true, false))

	assert.True(t, parse.IsEqual("1", "1"))
	assert.False(t, parse.IsEqual("steden1", "steden2"))

	assert.True(t, parse.IsEqual(1, 1))
	assert.False(t, parse.IsEqual(1, 2))

	assert.True(t, parse.IsEqual(int8(1), int8(1)))
	assert.False(t, parse.IsEqual(int8(1), int8(2)))

	assert.True(t, parse.IsEqual(int16(1), int16(1)))
	assert.False(t, parse.IsEqual(int16(1), int16(2)))

	assert.True(t, parse.IsEqual(int32(1), int32(1)))
	assert.False(t, parse.IsEqual(int32(1), int32(2)))

	assert.True(t, parse.IsEqual(int64(1), int64(1)))
	assert.False(t, parse.IsEqual(int64(1), int64(2)))

	assert.True(t, parse.IsEqual(uint(1), uint(1)))
	assert.False(t, parse.IsEqual(uint(1), uint(2)))

	assert.True(t, parse.IsEqual(uint8(1), uint8(1)))
	assert.False(t, parse.IsEqual(uint8(1), uint8(2)))

	assert.True(t, parse.IsEqual(uint16(1), uint16(1)))
	assert.False(t, parse.IsEqual(uint16(1), uint16(2)))

	assert.True(t, parse.IsEqual(uint32(1), uint32(1)))
	assert.False(t, parse.IsEqual(uint32(1), uint32(2)))

	assert.True(t, parse.IsEqual(uint64(1), uint64(1)))
	assert.False(t, parse.IsEqual(uint64(1), uint64(2)))

	assert.True(t, parse.IsEqual(float32(1), float32(1)))
	assert.False(t, parse.IsEqual(float32(1), float32(2)))

	assert.True(t, parse.IsEqual(float64(1), float64(1)))
	assert.False(t, parse.IsEqual(float64(1), float64(2)))
}
