package test

import (
	"github.com/farseer-go/fs/asyncLocal"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAsyncLocal(t *testing.T) {
	al := asyncLocal.New[string]()
	al.Set("A")
	assert.Equal(t, "A", al.Get())
	al.Remove()
	assert.Equal(t, "", al.Get())
}
