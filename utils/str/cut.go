package str

// CutRight 如果str末尾包含lastTag，则裁剪掉lastTag部份
func CutRight(str string, lastTag string) string {
	if str[len(str)-len(lastTag):] == lastTag {
		return str[0 : len(str)-len(lastTag)]
	}
	return str
}
