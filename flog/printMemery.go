package flog

import (
	"runtime"
)

// 测试使用
func PrintMemery(title string) {
	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)
	Infof("=== 内存统计（%s）GC 次数 : %d ===", title, memStats.NumGC)
	Infof("当前分配: %.2f MB，累计分配: %.2f MB，申请: %.2f MB", float64(memStats.Alloc)/1024/1024, float64(memStats.TotalAlloc)/1024/1024, float64(memStats.Sys)/1024/1024)
}
