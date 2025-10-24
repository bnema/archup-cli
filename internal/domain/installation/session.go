package installation

import (
	"time"

	"github.com/bnema/archup-cli/internal/domain/config"
	"github.com/bnema/archup-cli/internal/domain/hardware"
	"github.com/bnema/archup-cli/internal/domain/packages"
	"github.com/bnema/archup-cli/internal/domain/preset"
)

// NOTE: This is a placeholder domain aggregate for installation orchestration.
// Session is the root aggregate that coordinates the entire installation process.

// Session represents an installation session
type Session struct {
	ID        string
	CreatedAt time.Time
	UpdatedAt time.Time
	State     State

	// User selections
	Hardware   *hardware.Profile
	Config     *config.System
	Packages   *packages.Selection
	Preset     *preset.Preset
	Compositor string

	// Installation tracking
	Plan     *Plan
	Progress Progress
}

// State represents the current state of the installation
type State string

const (
	StateNew        State = "new"
	StatePlanning   State = "planning"
	StateReady      State = "ready"
	StateInstalling State = "installing"
	StateCompleted  State = "completed"
	StateFailed     State = "failed"
	StateRolledBack State = "rolled-back"
)

// Progress tracks installation progress
type Progress struct {
	CurrentStep int
	TotalSteps  int
	CurrentTask string
	Percentage  int
	Errors      []error
}

// NewSession creates a new installation session
// PLACEHOLDER: Not implemented yet
func NewSession() *Session {
	return &Session{
		ID:        generateID(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		State:     StateNew,
		Progress: Progress{
			CurrentStep: 0,
			TotalSteps:  0,
			Percentage:  0,
		},
	}
}

// SetHardware sets the detected hardware profile
// PLACEHOLDER: Not implemented yet
func (s *Session) SetHardware(hw *hardware.Profile) {
	s.Hardware = hw
	s.UpdatedAt = time.Now()
}

// SetConfig sets the system configuration
// PLACEHOLDER: Not implemented yet
func (s *Session) SetConfig(cfg *config.System) {
	s.Config = cfg
	s.UpdatedAt = time.Now()
}

// SetPackages sets the package selection
// PLACEHOLDER: Not implemented yet
func (s *Session) SetPackages(pkgs *packages.Selection) {
	s.Packages = pkgs
	s.UpdatedAt = time.Now()
}

// SetPreset sets the selected preset
// PLACEHOLDER: Not implemented yet
func (s *Session) SetPreset(p *preset.Preset) {
	s.Preset = p
	s.UpdatedAt = time.Now()
}

// GeneratePlan creates an installation plan from the session
// PLACEHOLDER: Not implemented yet
func (s *Session) GeneratePlan() (*Plan, error) {
	s.State = StatePlanning
	// Will generate ordered steps based on selections
	plan := NewPlan()
	s.Plan = plan
	s.State = StateReady
	return plan, nil
}

// Validate checks if the session is ready for installation
// PLACEHOLDER: Not implemented yet
func (s *Session) Validate() error {
	return nil
}

// MarkCompleted marks the session as completed
// PLACEHOLDER: Not implemented yet
func (s *Session) MarkCompleted() {
	s.State = StateCompleted
	s.UpdatedAt = time.Now()
}

// MarkFailed marks the session as failed
// PLACEHOLDER: Not implemented yet
func (s *Session) MarkFailed(err error) {
	s.State = StateFailed
	s.Progress.Errors = append(s.Progress.Errors, err)
	s.UpdatedAt = time.Now()
}

// generateID generates a unique session ID
// PLACEHOLDER: Not implemented yet
func generateID() string {
	return "session-" + time.Now().Format("20060102-150405")
}
