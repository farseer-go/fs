package snc

import (
	"bytes"
	"encoding/json"
)

// 将json转换成对象（反序列化）
func Unmarshal(data []byte, v any) error {
	d := json.NewDecoder(bytes.NewReader(data))
	d.UseNumber()
	return d.Decode(v)
}

// 将对象转换成json（序列化）
func Marshal(val any) ([]byte, error) {
	return json.Marshal(val)
}

func MarshalIndent(v any, prefix, indent string) ([]byte, error) {
	return json.MarshalIndent(v, prefix, indent)
}
