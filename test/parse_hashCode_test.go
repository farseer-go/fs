package test

import (
	"github.com/farseer-go/fs/parse"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHashCode(t *testing.T) {
	for i := 0; i < 100; i++ {
		assert.Less(t, 0, parse.HashCode(parse.RandString(10)))
	}
}
