package test

import (
	"github.com/farseer-go/fs/system"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestResourceResource(t *testing.T) {
	resource := system.GetResource()
	assert.Greater(t, resource.CpuCores, 0)
	resource.ToString()
}
