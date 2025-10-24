package hardware

import "context"

// NOTE: This is a placeholder interface for hardware detection.
// Implementations will be in internal/infra/hwdetect/

// Detector is responsible for detecting system hardware
type Detector interface {
	// Detect scans the system and returns a hardware profile
	Detect(ctx context.Context) (*Profile, error)

	// DetectGPU detects only GPU information
	DetectGPU(ctx context.Context) (*GPUInfo, error)

	// DetectCPU detects only CPU information
	DetectCPU(ctx context.Context) (*CPUInfo, error)

	// IsLaptop determines if the system is a laptop
	IsLaptop(ctx context.Context) (bool, error)
}

// DetectorOption configures the detector
type DetectorOption func(*detectorOptions)

type detectorOptions struct {
	skipGPU   bool
	skipAudio bool
}

// WithSkipGPU skips GPU detection
// PLACEHOLDER: Not implemented yet
func WithSkipGPU() DetectorOption {
	return func(o *detectorOptions) {
		o.skipGPU = true
	}
}

// WithSkipAudio skips audio detection
// PLACEHOLDER: Not implemented yet
func WithSkipAudio() DetectorOption {
	return func(o *detectorOptions) {
		o.skipAudio = true
	}
}
