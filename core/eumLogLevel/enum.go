package eumLogLevel

import "strings"

// Enum 日志等级
type Enum int

const (
	Trace Enum = iota
	Debug
	Information
	Warning
	Error
	Critical
	NoneLevel
)

// GetEnum 名称转枚举
func GetEnum(name string) Enum {
	switch strings.ToLower(name) {
	case "trace":
		return Trace
	case "debug":
		return Debug
	case "information", "info":
		return Information
	case "warning", "warn":
		return Warning
	case "error":
		return Error
	case "critical":
		return Critical
	}
	return NoneLevel
}

func (r Enum) ToString() string {
	switch r {
	case Trace:
		return "Trace"
	case Debug:
		return "Debug"
	case Information:
		return "Info"
	case Warning:
		return "Warn"
	case Error:
		return "Error"
	case Critical:
		return "Critical"
	}
	return "Info"
}
