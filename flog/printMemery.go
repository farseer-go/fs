package flog

import (
	"fmt"
	"runtime"
)

// 测试使用
func PrintMemery() {
	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)
	Infof("===== 内存使用统计 =====")
	Infof("当前分配: %.2f MB", float64(memStats.Alloc)/1024/1024)
	Infof("累计分配: %.2f MB", float64(memStats.TotalAlloc)/1024/1024)
	Infof("系统内存: %.2f MB", float64(memStats.Sys)/1024/1024)
	Infof("堆内存  : %.2f MB", float64(memStats.HeapAlloc)/1024/1024)
	Infof("GC 次数 : %d", memStats.NumGC)
	fmt.Println("=======================")
}
