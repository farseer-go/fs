package flog

// IFormatter 日志格式
type IFormatter interface {
	Formatter(log *logData) string
}
