package test

import (
	"github.com/farseer-go/fs/system"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestResourceResource(t *testing.T) {
	//ticker := time.NewTicker(time.Second)
	//for range ticker.C {
	//	fmt.Println(system.GetResource().ToString())
	//}

	//fmt.Println(system.GetResource().ToString())
	resource := system.GetResource()
	assert.Greater(t, resource.CpuCores, 0)
}
