package hwdetect

import (
	"context"

	"github.com/bnema/archup-cli/internal/domain/hardware"
)

// NOTE: This is a placeholder infrastructure implementation for hardware detection.
// Will use tools like lscpu, lspci, dmidecode, etc.

// SystemDetector implements hardware.Detector using system tools
type SystemDetector struct {
	// Could have config for which tools to use
}

// NewSystemDetector creates a new system detector
// PLACEHOLDER: Not implemented yet
func NewSystemDetector() *SystemDetector {
	return &SystemDetector{}
}

// Detect scans the system and returns a hardware profile
// PLACEHOLDER: Not implemented yet
func (d *SystemDetector) Detect(ctx context.Context) (*hardware.Profile, error) {
	// TODO: Implement using:
	// - lspci for GPU detection
	// - lscpu for CPU info
	// - cat /proc/cpuinfo for detailed CPU info
	// - dmidecode for laptop detection
	// - pactl/pw-cli for audio detection

	return &hardware.Profile{}, nil
}

// DetectGPU detects only GPU information
// PLACEHOLDER: Not implemented yet
func (d *SystemDetector) DetectGPU(ctx context.Context) (*hardware.GPUInfo, error) {
	// TODO: Use lspci | grep VGA
	return &hardware.GPUInfo{}, nil
}

// DetectCPU detects only CPU information
// PLACEHOLDER: Not implemented yet
func (d *SystemDetector) DetectCPU(ctx context.Context) (*hardware.CPUInfo, error) {
	// TODO: Use lscpu and /proc/cpuinfo
	return &hardware.CPUInfo{}, nil
}

// IsLaptop determines if the system is a laptop
// PLACEHOLDER: Not implemented yet
func (d *SystemDetector) IsLaptop(ctx context.Context) (bool, error) {
	// TODO: Check for battery in /sys/class/power_supply/
	// or use dmidecode
	return false, nil
}
