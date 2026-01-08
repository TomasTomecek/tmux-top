package disk

import (
	"github.com/TomasTomecek/tmux-top/conf"
	"testing"
)

func TestGetDiskStats(t *testing.T) {
	c := &conf.ConfigurationManager{
		Default: conf.LoadConfFromBytes([]byte(conf.GetDefaultConf())),
	}

	stats := GetDiskStats(c)

	if len(stats.Mounts) == 0 {
		t.Error("Expected at least one mount point, got none")
	}

	for _, mount := range stats.Mounts {
		if mount.MountPoint == "" {
			t.Error("Mount point should not be empty")
		}
		if mount.Total == 0 {
			t.Error("Total disk space should not be zero")
		}
		if mount.UsedPercent < 0 || mount.UsedPercent > 100 {
			t.Errorf("Used percentage should be between 0 and 100, got %.2f", mount.UsedPercent)
		}
	}
}

func TestGetDiskStatsDefaultTemplate(t *testing.T) {
	c := &conf.ConfigurationManager{
		Default: conf.LoadConfFromBytes([]byte(conf.GetDefaultConf())),
	}

	template := c.GetDiskTemplate("")
	PrintDiskStats(template, c)
}

func TestGetMountPoints(t *testing.T) {
	mounts, err := getMountPoints()
	if err != nil {
		t.Fatalf("Failed to get mount points: %v", err)
	}

	if len(mounts) == 0 {
		t.Error("Expected at least one mount point")
	}

	// Check that root filesystem is present
	if _, exists := mounts["/"]; !exists {
		t.Error("Root filesystem (/) should be present in mount points")
	}
}

func TestGetDiskStatForMount(t *testing.T) {
	stat, err := getDiskStatForMount("/")
	if err != nil {
		t.Fatalf("Failed to get disk stats for /: %v", err)
	}

	if stat.MountPoint != "/" {
		t.Errorf("Expected mount point to be /, got %s", stat.MountPoint)
	}

	if stat.Total == 0 {
		t.Error("Total should not be zero")
	}

	if stat.Used > stat.Total {
		t.Error("Used space cannot be greater than total space")
	}

	if stat.Available > stat.Total {
		t.Error("Available space cannot be greater than total space")
	}
}