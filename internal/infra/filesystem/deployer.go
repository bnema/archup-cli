package filesystem

import (
	"context"

	"github.com/bnema/archup-cli/internal/domain/preset"
)

// NOTE: This is a placeholder infrastructure for deploying dotfiles and configs.
// Will handle copying files, creating symlinks, backing up existing files.

// Deployer handles file deployment operations
type Deployer struct {
	backupDir string // Directory for backups
}

// NewDeployer creates a new file deployer
// PLACEHOLDER: Not implemented yet
func NewDeployer(backupDir string) *Deployer {
	return &Deployer{
		backupDir: backupDir,
	}
}

// DeployDotfile deploys a single dotfile
// PLACEHOLDER: Not implemented yet
func (d *Deployer) DeployDotfile(ctx context.Context, df preset.Dotfile) error {
	// TODO:
	// 1. If df.Backup, backup existing file to d.backupDir
	// 2. If df.Template, render template
	// 3. Copy/symlink file to df.Destination
	return nil
}

// DeployAll deploys all dotfiles from a preset
// PLACEHOLDER: Not implemented yet
func (d *Deployer) DeployAll(ctx context.Context, p *preset.Preset) error {
	// TODO: Deploy all dotfiles in preset
	for _, df := range p.Dotfiles {
		if err := d.DeployDotfile(ctx, df); err != nil {
			return err
		}
	}
	return nil
}

// Backup backs up a file
// PLACEHOLDER: Not implemented yet
func (d *Deployer) Backup(ctx context.Context, path string) error {
	// TODO: Copy file to d.backupDir with timestamp
	return nil
}

// Rollback restores files from backup
// PLACEHOLDER: Not implemented yet
func (d *Deployer) Rollback(ctx context.Context, files []string) error {
	// TODO: Restore files from d.backupDir
	return nil
}

// CreateSymlink creates a symbolic link
// PLACEHOLDER: Not implemented yet
func (d *Deployer) CreateSymlink(ctx context.Context, source, dest string) error {
	// TODO: os.Symlink(source, dest)
	return nil
}
