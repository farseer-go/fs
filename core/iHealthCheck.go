package core

// IHealthCheck 健康检查
type IHealthCheck interface {
	// Check 检查
	Check() (string, error)
}
