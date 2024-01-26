package core

import "database/sql"

// ITransaction 事务
type ITransaction interface {
	// Begin 开始
	// isolationLevels：事务等级
	Begin(isolationLevels ...sql.IsolationLevel) error
	// Commit 提交
	Commit()
	// Rollback 回滚
	Rollback()
	// Transaction 使用事务
	Transaction(executeFn func(), isolationLevels ...sql.IsolationLevel)
}
