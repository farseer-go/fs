package parse

import (
	"encoding/hex"
	"math/rand"
)

// RandString 随机字符串
func RandString(length int) string {
	//rand.Seed(time.Now().UnixNano())
	uLen := 6
	b := make([]byte, uLen)
	rand.Read(b)
	return hex.EncodeToString(b)[0:uLen]
}
