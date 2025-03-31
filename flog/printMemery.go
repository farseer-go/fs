package flog

import (
	"runtime"
)

// 测试使用
func PrintMemery(title string) {
	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)
	Infof("===== 内存统计（%s） =====", title)
	Infof("当前分配，还未释放: %.2f MB", float64(memStats.Alloc)/1024/1024)
	Infof("累计分配（含释放）: %.2f MB", float64(memStats.TotalAlloc)/1024/1024)
	Infof("系统内存: %.2f MB", float64(memStats.Sys)/1024/1024)
	Infof("堆内存，还未释放: %.2f MB", float64(memStats.HeapAlloc)/1024/1024)
	Infof("GC 次数 : %d", memStats.NumGC)
}
