package eumLogLevel

import (
	"encoding/json"
	"strings"
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

// MarshalJSON to output non base64 encoded []byte
// 此处不能用指针，否则json序列化时不执行
func (d Enum) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.ToString())
}

// UnmarshalJSON to deserialize []byte
func (d *Enum) UnmarshalJSON(b []byte) error {
	var numStr string
	err := json.Unmarshal(b, &numStr)
	*d = GetEnum(numStr)
	return err
}
