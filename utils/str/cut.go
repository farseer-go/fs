package str

// CutRight 裁剪末尾标签
func CutRight(str string, lastTag string) string {
	if str[len(str)-len(lastTag):] == lastTag {
		return str[0 : len(str)-len(lastTag)]
	}
	return str
}
