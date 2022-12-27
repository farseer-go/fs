package parse

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

// 为了保证百分百覆盖率
func TestParse(t *testing.T) {
	assert.Equal(t, "123", numberToNumber(8, "123", reflect.String))
	assert.Equal(t, "123", numberToString(8, "123", reflect.String))
	assert.Equal(t, 8, stringToNumber("123f", 8, reflect.String))
}
