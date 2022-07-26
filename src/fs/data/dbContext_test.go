package data

import (
	"testing"
	"time"
)

type TestMysqlContext struct {
	Admin TableSet[TestPO] `data:"name=admin"`
}

type TestPO struct {
	Id int `gorm:"primaryKey"`
	// 管理员名称
	UserName string
	// 管理员密码
	UserPwd string
	// 是否启用
	IsEnable bool
	// 上次登陆时间
	LastLoginAt time.Time
	// 上次登陆IP
	LastLoginIp string
	// 创建时间
	CreateAt time.Time
	// 创建人
	CreateUser string
	// 创建人ID
	CreateId string
}

func TestInit(t *testing.T) {
	context := Init[TestMysqlContext]("fops")
	if context.Admin.DbContext == nil {
		t.Errorf("context.DbContext = nil")
	}
}
