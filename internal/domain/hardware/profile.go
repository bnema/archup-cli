package hardware

// NOTE: This is a placeholder domain model for hardware detection.
// Implementation will be added when we migrate from TUI-embedded logic.

// Profile represents detected hardware characteristics
type Profile struct {
	GPU      GPUInfo
	CPU      CPUInfo
	Audio    AudioSystem
	IsLaptop bool
}

// GPUInfo contains graphics hardware information
type GPUInfo struct {
	Vendor       string // "nvidia", "amd", "intel"
	Model        string
	Driver       string // Recommended driver: "nvidia", "mesa", etc.
	IsIntegrated bool
}

// CPUInfo contains processor information
type CPUInfo struct {
	Vendor         string // "AMD", "Intel"
	Model          string
	Cores          int
	Threads        int
	FrequencyRange FrequencyRange
	AvailableGovs  []string // Available CPU governors
}

// FrequencyRange represents CPU frequency capabilities
type FrequencyRange struct {
	MinMHz int
	MaxMHz int
}

// AudioSystem represents audio hardware
type AudioSystem struct {
	Server string // "pipewire", "pulseaudio", etc.
	Cards  []string
}

// Validate checks if the profile has required information
// PLACEHOLDER: Not implemented yet
func (p *Profile) Validate() error {
	return nil
}

// SupportsGPUDriver checks if a specific GPU driver is supported
// PLACEHOLDER: Not implemented yet
func (g *GPUInfo) SupportsGPUDriver(driver string) bool {
	return false
}

// RecommendedPackages returns GPU-specific packages needed
// PLACEHOLDER: Not implemented yet
func (g *GPUInfo) RecommendedPackages() []string {
	return []string{}
}
