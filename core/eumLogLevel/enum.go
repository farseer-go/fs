package eumLogLevel

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

// GetName 获取标签名称
func GetName(eum Enum) string {
	switch eum {
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
