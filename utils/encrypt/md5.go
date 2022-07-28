package encrypt

import (
	"crypto/md5"
	"fmt"
)

// Md5 对字符串做MD5加密
func Md5(str string) string {
	data := []byte(str)
	sum := md5.Sum(data)
	return fmt.Sprintf("%x", sum) //将[]byte转成16进制
}
