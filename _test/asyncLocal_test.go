package test

import (
	"testing"

	"github.com/farseer-go/fs/asyncLocal"
	"github.com/stretchr/testify/assert"
)

func TestAsyncLocal(t *testing.T) {
	al := asyncLocal.New[string]()
	al.Set("A")
	assert.Equal(t, "A", al.Get())
	asyncLocal.Release()
	assert.Equal(t, "", al.Get())
}
