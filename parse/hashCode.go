package parse

import (
	"bytes"
	"fmt"
	"hash/crc32"
)

// HashCode 获取哈希值
func HashCode(s string) int {
	v := int(crc32.ChecksumIEEE([]byte(s)))
	if v >= 0 {
		return v
	}
	if -v >= 0 {
		return -v
	}
	return 0
}

// HashCodes 获取哈希值
func HashCodes(strings []string) string {
	var buf bytes.Buffer
	for _, s := range strings {
		buf.WriteString(fmt.Sprintf("%s-", s))
	}
	return fmt.Sprintf("%d", HashCode(buf.String()))
}
