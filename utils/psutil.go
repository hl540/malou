package utils

import (
	"github.com/hl540/malou/proto/v1"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
	"math"
	"time"
)

func GetCpuPercent() float64 {
	percent, err := cpu.Percent(time.Second, false)
	if err != nil {
		return 0
	}
	return math.Round(percent[0]*100) / 100
}

func GetMemoryPercent() *v1.MemoryInfo {
	memory, err := mem.VirtualMemory()
	if err != nil {
		return &v1.MemoryInfo{}
	}
	return &v1.MemoryInfo{
		Total:       bytesToGB(memory.Total),
		Used:        bytesToGB(memory.Used),
		Free:        bytesToGB(memory.Free),
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
		Total:       bytesToGB(diskInfo.Total),
		Used:        bytesToGB(diskInfo.Used),
		Free:        bytesToGB(diskInfo.Free),
		UsedPercent: diskInfo.UsedPercent,
	}
}

func bytesToGB(bytes uint64) float64 {
	value := float64(bytes) / (1024 * 1024 * 1024)
	return math.Round(value*100) / 100
}
