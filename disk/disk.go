package disk

import (
	"bufio"
	"github.com/TomasTomecek/tmux-top/conf"
	"os"
	"strings"
	"syscall"
	"text/template"
)

type DiskStat struct {
	MountPoint  string  `json:"mountpoint"`
	Device      string  `json:"device"`
	Total       uint64  `json:"total"`
	Used        uint64  `json:"used"`
	Available   uint64  `json:"available"`
	UsedPercent float64 `json:"used_percent"`
}

type DiskStats struct {
	Mounts []DiskStat `json:"mounts"`
}

func getDiskStatForMount(mountpoint string) (DiskStat, error) {
	var stat syscall.Statfs_t
	err := syscall.Statfs(mountpoint, &stat)
	if err != nil {
		return DiskStat{}, err
	}

	total := stat.Blocks * uint64(stat.Bsize)
	available := stat.Bavail * uint64(stat.Bsize)
	used := total - (stat.Bfree * uint64(stat.Bsize))
	usedPercent := float64(used) / float64(total) * 100

	return DiskStat{
		MountPoint:  mountpoint,
		Total:       total,
		Used:        used,
		Available:   available,
		UsedPercent: usedPercent,
	}, nil
}

func getMountPoints() (map[string]string, error) {
	file, err := os.Open("/proc/mounts")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	mounts := make(map[string]string)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		if len(fields) < 3 {
			continue
		}

		device := fields[0]
		mountpoint := fields[1]

		mounts[mountpoint] = device
	}

	return mounts, scanner.Err()
}

func GetDiskStats(c *conf.ConfigurationManager) DiskStats {
	diskStats := make([]DiskStat, 0)
	stats := DiskStats{
		Mounts: diskStats,
	}
	mounts, err := getMountPoints()
	if err != nil {
		return stats
	}

	confMounts, err := c.GetDiskMounts()
	if err != nil {
		return stats
	}

	for _, mountpoint := range confMounts {
		stat, err := getDiskStatForMount(mountpoint)
		stat.Device = mounts[mountpoint]
		if err != nil {
			continue
		}
		diskStats = append(diskStats, stat)
	}

	stats.Mounts = diskStats
	return stats
}

func PrintDiskStats(t template.Template, c *conf.ConfigurationManager) {
	d := GetDiskStats(c)
	err := t.Execute(os.Stdout, d)
	if err != nil {
		panic(err)
	}
}
