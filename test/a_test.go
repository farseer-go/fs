package test

import (
	"github.com/farseer-go/fs/container"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFirst(t *testing.T) {
	test := &mysql{}
	assert.Panics(t, func() {
		container.Register(func() IDatabase {
			return &mysql{}
		})
	})
	assert.Panics(t, func() {
		container.RegisterInstance[IDatabase](test)
	})

	assert.Panics(t, func() {
		container.RegisterTransient(func() IDatabase {
			return &mysql{}
		}, "test")
	})

}
