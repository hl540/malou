package utils

import (
	"github.com/hl540/malou/proto/v1"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
)

func GetCpuPercent() []float64 {
	percent, err := cpu.Percent(0, false)
	if err != nil {
		return make([]float64, 0)
	}
	return percent
}

func GetMemoryPercent() *v1.MemoryInfo {
	memory, err := mem.VirtualMemory()
	if err != nil {
		return &v1.MemoryInfo{}
	}
	return &v1.MemoryInfo{
		Total:       memory.Total,
		Used:        memory.Used,
		Free:        memory.Free,
		UsedPercent: memory.UsedPercent,
	}
}

func GetDiskPercent() *v1.DiskInfo {
	parts, err := disk.Partitions(false)
	if err != nil {
		return &v1.DiskInfo{}
	}
	diskInfo, err := disk.Usage(parts[0].Mountpoint)
	if err != nil {
		return &v1.DiskInfo{}
	}
	return &v1.DiskInfo{
		Total:       diskInfo.Total,
		Used:        diskInfo.Used,
		Free:        diskInfo.Free,
		UsedPercent: diskInfo.UsedPercent,
	}
}
