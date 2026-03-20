package eumLogLevel

import (
	"strings"

	"github.com/farseer-go/fs/snc"
	"github.com/vmihailenco/msgpack/v5"
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

// MarshalMsgpack 将 Enum 转为字符串二进制
// 同样不建议用指针接收者，确保值传递时也能触发
func (receiver Enum) MarshalMsgpack() ([]byte, error) {
	// 1. 获取枚举的字符串表示
	s := receiver.ToString()

	// 2. 序列化为 Msgpack 字符串 (会自动处理长度标识，不带引号)
	return msgpack.Marshal(s)
}

// UnmarshalMsgpack 从二进制中恢复 Enum
func (receiver *Enum) UnmarshalMsgpack(b []byte) error {
	var numStr string

	// 1. 解出字符串
	err := msgpack.Unmarshal(b, &numStr)
	if err != nil {
		return err
	}

	// 2. 通过你现有的工厂函数还原枚举值
	*receiver = GetEnum(numStr)

	return nil
}
