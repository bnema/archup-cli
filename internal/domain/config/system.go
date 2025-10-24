package config

import "github.com/bnema/archup-cli/internal/domain/hardware"

// NOTE: This is a placeholder domain aggregate for system configuration.
// This is the main aggregate that combines all system configuration concerns.

// System represents the complete system configuration
type System struct {
	CPU    CPUConfig
	Power  PowerConfig
	Pacman PacmanConfig
}

// NewSystem creates a new system configuration with defaults
// PLACEHOLDER: Not implemented yet
func NewSystem(profile *hardware.Profile) *System {
	return &System{
		CPU:    NewCPUConfig(profile),
		Power:  NewPowerConfig(profile),
		Pacman: NewPacmanConfig(),
	}
}

// Validate checks if the configuration is valid
// PLACEHOLDER: Not implemented yet
func (s *System) Validate() error {
	return nil
}

// Apply applies the configuration to the system
// PLACEHOLDER: Not implemented yet - will delegate to infrastructure layer
func (s *System) Apply() error {
	return nil
}

// ToFiles generates configuration files that need to be written
// PLACEHOLDER: Not implemented yet
func (s *System) ToFiles() map[string]string {
	// Returns map of filepath -> content
	return map[string]string{}
}
