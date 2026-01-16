package test

import (
	"strconv"
	"testing"

	"github.com/farseer-go/fs/sonyflake"
	"github.com/stretchr/testify/assert"
)

func TestSonyflake(t *testing.T) {
	assert.Equal(t, 18, len(strconv.FormatInt(sonyflake.GenerateId(), 10)))
}
