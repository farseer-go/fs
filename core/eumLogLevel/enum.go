package eumLogLevel

import (
	"strings"

	"github.com/farseer-go/fs/snc"
)

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

func (receiver Enum) ToString() string {
	switch receiver {
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
	default:
		return "None"
	}
}

// MarshalJSON to output non base64 encoded []byte
// 此处不能用指针，否则json序列化时不执行
func (receiver Enum) MarshalJSON() ([]byte, error) {
	return snc.Marshal(receiver.ToString())
}

// UnmarshalJSON to deserialize []byte
func (receiver *Enum) UnmarshalJSON(b []byte) error {
	var numStr string
	err := snc.Unmarshal(b, &numStr)
	*receiver = GetEnum(numStr)
	return err
}
