package parse

import (
	"bytes"
	"fmt"
	"hash/crc32"
)

// HashCode 获取哈希值
func HashCode(s string) uint32 {
	return crc32.ChecksumIEEE([]byte(s))
}

// HashCode64 获取哈希值
func HashCode64(s string) int64 {
	return int64(HashCode(s))
}

// HashCodes 获取哈希值
func HashCodes(strings []string) uint32 {
	var buf bytes.Buffer
	for _, s := range strings {
		buf.WriteString(fmt.Sprintf("%s-", s))
	}
	return HashCode(buf.String())
}
