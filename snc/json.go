package snc

import (
	"bytes"
	"encoding/json"

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
		CompactMarshaler: true, // 输出紧凑json
		NoNullSliceOrMap: true, // 空对象编码为：[] {}
		UseInt64:         true, // 整数对象转换为int64，否则为float64
		UseNumber:        true, // 不要转换成float64而是json.number
		CopyString:       true, // 不要引用字符串，而是复制一份出来
	}.Froze()
}

// 将json转换成对象（反序列化）
func Unmarshal(data []byte, v any) error {
	d := json.NewDecoder(bytes.NewReader(data))
	d.UseNumber()
	return d.Decode(v)
	//return snc.Unmarshal(data, v)
}

// 将对象转换成json（序列化）
func Marshal(val any) ([]byte, error) {
	return snc.Marshal(val)
}

func MarshalIndent(v any, prefix, indent string) ([]byte, error) {
	return snc.MarshalIndent(v, prefix, indent)
}
