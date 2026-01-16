package test

import (
	"github.com/farseer-go/fs/parse"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHashCode(t *testing.T) {
	for i := 0; i < 100; i++ {
		assert.Less(t, uint32(0), parse.HashCode(parse.RandString(10)))
	}
	assert.Less(t, uint32(0), parse.HashCode("abcdefghijklmnopqrstuvwxyz"))
	assert.Equal(t, uint32(0x9cb642c2), parse.HashCodes([]string{"1", "2", "3"}))
}
