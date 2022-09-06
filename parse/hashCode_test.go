package parse

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHashCode(t *testing.T) {
	for i := 0; i < 100; i++ {
		assert.Less(t, 0, HashCode(RandString(10)))
	}
}
