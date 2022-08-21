package parse

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsEqual(t *testing.T) {
	assert.True(t, IsEqual(true, true))
	assert.False(t, IsEqual(true, false))

	assert.True(t, IsEqual("1", "1"))
	assert.False(t, IsEqual("steden1", "steden2"))

	assert.True(t, IsEqual(1, 1))
	assert.False(t, IsEqual(1, 2))

	assert.True(t, IsEqual(int8(1), int8(1)))
	assert.False(t, IsEqual(int8(1), int8(2)))

	assert.True(t, IsEqual(int16(1), int16(1)))
	assert.False(t, IsEqual(int16(1), int16(2)))

	assert.True(t, IsEqual(int32(1), int32(1)))
	assert.False(t, IsEqual(int32(1), int32(2)))

	assert.True(t, IsEqual(int64(1), int64(1)))
	assert.False(t, IsEqual(int64(1), int64(2)))

	assert.True(t, IsEqual(uint(1), uint(1)))
	assert.False(t, IsEqual(uint(1), uint(2)))

	assert.True(t, IsEqual(uint8(1), uint8(1)))
	assert.False(t, IsEqual(uint8(1), uint8(2)))

	assert.True(t, IsEqual(uint16(1), uint16(1)))
	assert.False(t, IsEqual(uint16(1), uint16(2)))

	assert.True(t, IsEqual(uint32(1), uint32(1)))
	assert.False(t, IsEqual(uint32(1), uint32(2)))

	assert.True(t, IsEqual(uint64(1), uint64(1)))
	assert.False(t, IsEqual(uint64(1), uint64(2)))

	assert.True(t, IsEqual(float32(1), float32(1)))
	assert.False(t, IsEqual(float32(1), float32(2)))

	assert.True(t, IsEqual(float64(1), float64(1)))
	assert.False(t, IsEqual(float64(1), float64(2)))
}
