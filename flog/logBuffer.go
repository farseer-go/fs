package flog

// LogBuffer 日志缓冲区
var LogBuffer = make(chan string, 1000)

// LoadLogBuffer 从日志缓冲区读取日志并打印
func LoadLogBuffer() {
	for log := range LogBuffer {
		Println(log)
	}
}

// ClearLogBuffer 清空缓冲区的日志
func ClearLogBuffer() {
	for len(LogBuffer) > 0 {
		Println(<-LogBuffer)
	}
}
