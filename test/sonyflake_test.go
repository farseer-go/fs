package test

import (
	"github.com/farseer-go/fs/parse"
	"github.com/farseer-go/fs/sonyflake"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSonyflake(t *testing.T) {
	assert.Equal(t, 18, len(parse.ToString(sonyflake.GenerateId())))
}
