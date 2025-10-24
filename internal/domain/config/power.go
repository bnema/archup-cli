package config

import "github.com/bnema/archup-cli/internal/domain/hardware"

// NOTE: This is a placeholder for power management configuration (TLP).
// Contains business rules for laptop power management profiles.

// PowerConfig represents power management configuration (TLP)
type PowerConfig struct {
	Enabled bool
	Profile TLPProfile
}

// TLPProfile represents a TLP power profile
type TLPProfile string

const (
	TLPBalanced    TLPProfile = "balanced"
	TLPPerformance TLPProfile = "performance"
	TLPBattery     TLPProfile = "battery"
	TLPDisabled    TLPProfile = "disabled"
)

// TLPProfileInfo contains metadata about a TLP profile
type TLPProfileInfo struct {
	Profile     TLPProfile
	Description string
	BatteryLife string // "high", "medium", "low"
	Performance string // "high", "medium", "low"
}

// NewPowerConfig creates power configuration based on hardware
// PLACEHOLDER: Not implemented yet
func NewPowerConfig(profile *hardware.Profile) PowerConfig {
	if !profile.IsLaptop {
		return PowerConfig{
			Enabled: false,
			Profile: TLPDisabled,
		}
	}
	return PowerConfig{
		Enabled: true,
		Profile: TLPBalanced,
	}
}

// AvailableProfiles returns all TLP profiles with metadata
// PLACEHOLDER: Not implemented yet
func AvailableProfiles() []TLPProfileInfo {
	return []TLPProfileInfo{
		{
			Profile:     TLPBalanced,
			Description: "Balanced performance and battery life",
			BatteryLife: "medium",
			Performance: "medium",
		},
		{
			Profile:     TLPPerformance,
			Description: "Prioritize performance over battery",
			BatteryLife: "low",
			Performance: "high",
		},
		{
			Profile:     TLPBattery,
			Description: "Maximize battery life",
			BatteryLife: "high",
			Performance: "low",
		},
		{
			Profile:     TLPDisabled,
			Description: "Don't configure TLP",
			BatteryLife: "n/a",
			Performance: "n/a",
		},
	}
}

// SetProfile sets the TLP profile
// PLACEHOLDER: Not implemented yet
func (p *PowerConfig) SetProfile(profile TLPProfile) {
	p.Profile = profile
	p.Enabled = profile != TLPDisabled
}

// Validate checks if the configuration is valid
// PLACEHOLDER: Not implemented yet
func (p *PowerConfig) Validate() error {
	return nil
}
