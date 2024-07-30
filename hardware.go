package common

import (
	"github.com/google/uuid"
	"github.com/klauspost/cpuid/v2"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

func GetCpuArch() string {
	cpuArch := "unknown"
	switch runtime.GOARCH {
	case "amd64":
		cpuArch = "x86_64"
	case "386":
		cpuArch = "x86"
	case "arm":
		cpuArch = "arm"
	case "arm64":
		cpuArch = "arm64"
	case "aarch64":
		cpuArch = "arm64"
	}
	return cpuArch
}

func GetCpuBrandName() string {
	return cpuid.CPU.BrandName
}

func GetCpuCores() int {
	return runtime.NumCPU()
}

func GetMemorySize() uint64 {
	mem := new(runtime.MemStats)
	runtime.ReadMemStats(mem)
	return mem.TotalAlloc
}

func GetOsVersion() string {
	return runtime.GOOS
}

// GetCurrTimestamp return current timestamp
func GetCurrTimestamp() int64 {
	return time.Now().Unix()
}

// GetUUID return a new UUID
func GetUUID() string {
	return uuid.New().String()
}

func GetDisks() []string {
	switch runtime.GOOS {
	case "darwin":
		return GetMacDiskList()
	case "linux":
		return GetLinuxDiskList()
	case "windows":
		return GetWindowsDiskList()
	default:
		return nil
	}
}

func GetLinuxDiskList() []string {
	cmd := "lsblk --list|grep disk|awk '{print $1}'"
	result := exec.Command("sh", "-c", cmd)
	output, err := result.CombinedOutput()
	if err != nil {
		return nil
	}
	var disks []string
	for _, line := range strings.Split(string(output), "\n") {
		if line != "" {
			disks = append(disks, line)
		}
	}
	return disks
}

func GetWindowsDiskList() []string {
	return nil
}

func GetMacDiskList() []string {
	cmd := "diskutil list|grep '^/dev/disk'|awk '{print $1}'|awk -F/ '{print $3}'"
	result := exec.Command("sh", "-c", cmd)
	output, err := result.CombinedOutput()
	if err != nil {
		return nil
	}
	var disks []string
	for _, line := range strings.Split(string(output), "\n") {
		if line != "" {
			disks = append(disks, line)
		}
	}
	return disks
}
