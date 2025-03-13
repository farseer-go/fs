package snc

import (
	// "github.com/bytedance/sonic"
	// "github.com/bytedance/sonic/option"
	"bytes"

	jsoniter "github.com/json-iterator/go"
)

//var snc sonic.API

func init() {
	// 设置较小的缓冲区，以减少内存使用
	// option.LimitBufferSize = 2 * 1024
	// option.DefaultDecoderBufferSize = 2 * 1024
	// option.DefaultEncoderBufferSize = 2 * 1024
	// option.DefaultAstBufferSize = 2 * 1024

	// snc = sonic.Config{
	// 	CompactMarshaler: true,
	// 	UseNumber:        true,
	// 	CopyString:       true,
	// }.Froze()
}

// 将json转换成对象（反序列化）
func Unmarshal(data []byte, v any) error {
	// d := json.NewDecoder(bytes.NewReader(data))
	// d.UseNumber()
	// return d.Decode(v)
	//return snc.Unmarshal(data, v)
	d := jsoniter.NewDecoder(bytes.NewReader(data))
	d.UseNumber()
	return d.Decode(v)
}

// 将对象转换成json（序列化）
func Marshal(val any) ([]byte, error) {
	return jsoniter.Marshal(val)
}

func MarshalIndent(v any, prefix, indent string) ([]byte, error) {
	return jsoniter.MarshalIndent(v, prefix, indent)
}
