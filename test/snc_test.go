package test

import (
	"context"
	"testing"

	"github.com/farseer-go/fs/dateTime"
	"github.com/farseer-go/fs/snc"
	"github.com/stretchr/testify/assert"
)

type UnmarshalVO struct {
	Id         string            // 客户端ID
	Name       string            // 客户端名称
	Ip         string            // 客户端IP
	Port       int               // 客户端端口
	ActivateAt dateTime.DateTime // 活动时间
	ScheduleAt dateTime.DateTime // 任务调度时间
	QueueCount int               // 排队中的任务数量
	WorkCount  int               // 正在处理的任务数量
	ErrorCount int               // 错误次数
	Ctx        context.Context   `json:"-"` // 用于通知应用端是否断开连接
	IsMaster   bool              // 是否为主客户端
}

func TestUnmarshal(t *testing.T) {
	json := "{\"Id\":\"10.100.0.110:56426\",\"Name\":\"aaa\",\"Ip\":\"10.100.0.110\",\"Port\":56426,\"ActivateAt\":\"2025-06-08 12:21:19\",\"ScheduleAt\":\"0001-01-01 00:00:00\",\"Status\":0,\"QueueCount\":0,\"WorkCount\":0,\"ErrorCount\":0,\"Job\":{\"Name\":\"aaa.1\",\"Ver\":1},\"Ctx\":{\"Context\":{}},\"IsMaster\":true}"
	var item UnmarshalVO
	snc.Unmarshal([]byte(json), &item)

	assert.True(t, item.IsMaster)
}
