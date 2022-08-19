package parse

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConvert(t *testing.T) {
	t.Run("int转bool", func(t *testing.T) {
		result := Convert(0, false)
		assert.False(t, result)
		result = Convert(1, false)
		assert.True(t, result)
	})

	t.Run("int转字符串", func(t *testing.T) {
		result := Convert(1, "")
		assert.Equal(t, result, "1")
	})

	t.Run("int转int64", func(t *testing.T) {
		result := Convert(1, int64(0))
		assert.Equal(t, result, int64(1))
	})

	t.Run("字符串转bool", func(t *testing.T) {
		result := Convert("true", false)
		assert.True(t, result)
	})

	t.Run("字符串转int", func(t *testing.T) {
		result := Convert("123", 0)
		assert.Equal(t, result, 123)
	})

	t.Run("字符串转int64", func(t *testing.T) {
		result := Convert("123", int64(0))
		assert.Equal(t, result, int64(123))
	})
}
