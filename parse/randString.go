package parse

import (
	"encoding/hex"
	"math/rand"
	"time"
)

// RandString 随机字符串
func RandString(length int) string {
	b := make([]byte, length)
	rand.New(rand.NewSource(time.Now().UnixNano())).Read(b)
	return hex.EncodeToString(b)[0:length]
}
