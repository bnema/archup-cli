package config

import "github.com/bnema/archup-cli/internal/domain/hardware"

// NOTE: This is a placeholder for CPU configuration domain logic.
// Contains business rules for CPU governor selection and frequency scaling.

// CPUConfig represents CPU-related configuration
type CPUConfig struct {
	Governor      hardware.Governor
	MinFrequency  int // MHz
	MaxFrequency  int // MHz
	BoostEnabled  bool
	TurboEnabled  bool
}

// NewCPUConfig creates CPU configuration based on hardware profile
// PLACEHOLDER: Not implemented yet
func NewCPUConfig(profile *hardware.Profile) CPUConfig {
	return CPUConfig{
		Governor:     hardware.RecommendGovernor(profile),
		MinFrequency: profile.CPU.FrequencyRange.MinMHz,
		MaxFrequency: profile.CPU.FrequencyRange.MaxMHz,
		BoostEnabled: true,
		TurboEnabled: true,
	}
}

// SetGovernor sets the CPU governor
// PLACEHOLDER: Not implemented yet
func (c *CPUConfig) SetGovernor(gov hardware.Governor) {
	c.Governor = gov
}

// SetFrequencyRange sets the frequency limits
// PLACEHOLDER: Not implemented yet
func (c *CPUConfig) SetFrequencyRange(min, max int) error {
	if min > max {
		return &ValidationError{Field: "frequency", Message: "min cannot be greater than max"}
	}
	c.MinFrequency = min
	c.MaxFrequency = max
	return nil
}

// Validate checks if the configuration is valid
// PLACEHOLDER: Not implemented yet
func (c *CPUConfig) Validate() error {
	return nil
}

// ValidationError represents a configuration validation error
type ValidationError struct {
	Field   string
	Message string
}

func (e *ValidationError) Error() string {
	return e.Field + ": " + e.Message
}
