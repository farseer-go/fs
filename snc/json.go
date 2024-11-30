package snc

import (
	"github.com/bytedance/sonic"
	"github.com/bytedance/sonic/option"
)

var snc sonic.API

func init() {
	// 设置较小的缓冲区，以减少内存使用
	option.LimitBufferSize = 2 * 1024
	option.DefaultDecoderBufferSize = 2 * 1024
	option.DefaultEncoderBufferSize = 2 * 1024
	option.DefaultAstBufferSize = 2 * 1024

	snc = sonic.Config{
		CompactMarshaler: true,
		UseNumber:        true,
		CopyString:       true,
	}.Froze()
}

func Unmarshal(data []byte, v any) error {
	return snc.Unmarshal(data, v)
}

func Marshal(val any) ([]byte, error) {
	return snc.Marshal(val)
}

func MarshalIndent(v any, prefix, indent string) ([]byte, error) {
	return snc.MarshalIndent(v, prefix, indent)
}
