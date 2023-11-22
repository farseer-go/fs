package system

import (
	"fmt"
	"github.com/farseer-go/fs/net"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/mem"
)

type resource struct {
	OS                 string  // 操作系统名称
	Processes          uint64  // 进程数
	IP                 string  // IP
	CpuName            string  // CPU名称
	CpuMhz             float64 // CPU总赫兹
	CpuCores           int     // CPU核心数
	CpuUsagePercent    float64 // CPU使用百分比
	MemoryTotal        uint64  // 总内存
	MemoryAvailable    uint64  // 内存可用量
	MemoryUsage        uint64  // 内存已使用
	MemoryUsagePercent float64 // 内存使用百分比
	DiskTotal          uint64  // 硬盘总容量
	DiskAvailable      uint64  // 硬盘可用空间
	DiskUsage          uint64  // 硬盘已用空间
	DiskUsagePercent   float64 // 硬盘使用百分比
}

func (receiver *resource) ToString() string {
	return fmt.Sprintf("%+v", receiver)
}

// GetResource 获取当前环境信息
func GetResource() *resource {
	info, _ := cpu.Percent(0, false)
	infoStats, _ := cpu.Info()
	memory, _ := mem.VirtualMemory()
	hostInfo, _ := host.Info()
	diskUsage, _ := disk.Usage("/")
	return &resource{
		OS:                 hostInfo.OS,
		Processes:          hostInfo.Procs,
		IP:                 net.GetIp(),
		CpuName:            infoStats[0].ModelName,
		CpuMhz:             infoStats[0].Mhz,
		CpuUsagePercent:    info[0],
		CpuCores:           int(infoStats[0].Cores),
		MemoryUsage:        memory.Used,
		MemoryUsagePercent: memory.UsedPercent,
		MemoryTotal:        memory.Total,
		MemoryAvailable:    memory.Available,
		DiskTotal:          diskUsage.Total,
		DiskAvailable:      diskUsage.Free,
		DiskUsage:          diskUsage.Used,
		DiskUsagePercent:   diskUsage.UsedPercent,
	}
}
