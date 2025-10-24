package services

import (
	"context"

	"github.com/bnema/archup-cli/internal/domain/config"
	"github.com/bnema/archup-cli/internal/domain/hardware"
	"github.com/bnema/archup-cli/internal/domain/installation"
	"github.com/bnema/archup-cli/internal/domain/packages"
	"github.com/bnema/archup-cli/internal/domain/preset"
)

// NOTE: This is a placeholder application service for installation orchestration.
// Services coordinate between domain objects and infrastructure.

// InstallationService orchestrates the installation process
type InstallationService struct {
	hardwareDetector hardware.Detector
	packageRepo      packages.Repository
	presetRepo       preset.Repository
	executor         installation.Executor
}

// NewInstallationService creates a new installation service
// PLACEHOLDER: Not implemented yet
func NewInstallationService(
	hwDetector hardware.Detector,
	pkgRepo packages.Repository,
	presetRepo preset.Repository,
	executor installation.Executor,
) *InstallationService {
	return &InstallationService{
		hardwareDetector: hwDetector,
		packageRepo:      pkgRepo,
		presetRepo:       presetRepo,
		executor:         executor,
	}
}

// CreateSession creates a new installation session with hardware detection
// PLACEHOLDER: Not implemented yet
func (s *InstallationService) CreateSession(ctx context.Context) (*installation.Session, error) {
	session := installation.NewSession()

	// Detect hardware
	hw, err := s.hardwareDetector.Detect(ctx)
	if err != nil {
		return nil, err
	}
	session.SetHardware(hw)

	// Initialize default config based on hardware
	cfg := config.NewSystem(hw)
	session.SetConfig(cfg)

	return session, nil
}

// ConfigureSystem updates the system configuration in the session
// PLACEHOLDER: Not implemented yet
func (s *InstallationService) ConfigureSystem(
	session *installation.Session,
	cpuGov hardware.Governor,
	tlpProfile config.TLPProfile,
	pacmanCfg config.PacmanConfig,
) error {
	session.Config.CPU.SetGovernor(cpuGov)
	session.Config.Power.SetProfile(tlpProfile)
	session.Config.Pacman = pacmanCfg
	return nil
}

// SelectPackages updates the package selection in the session
// PLACEHOLDER: Not implemented yet
func (s *InstallationService) SelectPackages(
	ctx context.Context,
	session *installation.Session,
	selectedPackages []string,
) error {
	selection := &packages.Selection{}
	for _, pkg := range selectedPackages {
		_ = selection.Add(pkg)
	}

	// Resolve dependencies
	_, err := selection.Resolve()
	if err != nil {
		return err
	}

	session.SetPackages(selection)
	return nil
}

// ApplyPreset applies a preset to the session
// PLACEHOLDER: Not implemented yet
func (s *InstallationService) ApplyPreset(
	ctx context.Context,
	session *installation.Session,
	presetID string,
) error {
	p, err := s.presetRepo.Get(ctx, presetID)
	if err != nil {
		return err
	}

	session.SetPreset(p)
	return nil
}

// Install executes the installation
// PLACEHOLDER: Not implemented yet
func (s *InstallationService) Install(
	ctx context.Context,
	session *installation.Session,
) error {
	// Validate session
	if err := session.Validate(); err != nil {
		return err
	}

	// Generate plan
	plan, err := session.GeneratePlan()
	if err != nil {
		return err
	}

	// Execute plan
	if err := s.executor.Execute(ctx, plan); err != nil {
		session.MarkFailed(err)
		return err
	}

	session.MarkCompleted()
	return nil
}
