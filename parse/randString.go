package parse

import (
	"encoding/hex"
	"math/rand"
)

// RandString 随机字符串
func RandString(length int) string {
	b := make([]byte, length)
	rand.Read(b)
	return hex.EncodeToString(b)[0:length]
}
