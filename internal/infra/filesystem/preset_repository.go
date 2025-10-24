package filesystem

import (
	"context"

	"github.com/bnema/archup-cli/internal/domain/preset"
)

// NOTE: This is a placeholder infrastructure for preset storage on filesystem.
// Will handle loading presets from disk, git repos, or embedded files.

// PresetRepository implements preset.Repository using filesystem
type PresetRepository struct {
	basePath string // Base path for presets
}

// NewPresetRepository creates a new filesystem preset repository
// PLACEHOLDER: Not implemented yet
func NewPresetRepository(basePath string) *PresetRepository {
	return &PresetRepository{
		basePath: basePath,
	}
}

// Get retrieves a preset by ID
// PLACEHOLDER: Not implemented yet
func (r *PresetRepository) Get(ctx context.Context, id string) (*preset.Preset, error) {
	// TODO: Load preset from <basePath>/<id>/preset.yaml
	// Parse YAML into Preset struct
	return &preset.Preset{}, nil
}

// List returns all available presets
// PLACEHOLDER: Not implemented yet
func (r *PresetRepository) List(ctx context.Context) ([]preset.Preset, error) {
	// TODO: Scan <basePath> directory for preset.yaml files
	return []preset.Preset{}, nil
}

// Save persists a preset
// PLACEHOLDER: Not implemented yet
func (r *PresetRepository) Save(ctx context.Context, p *preset.Preset) error {
	// TODO: Write preset to <basePath>/<id>/preset.yaml
	return nil
}

// Delete removes a preset
// PLACEHOLDER: Not implemented yet
func (r *PresetRepository) Delete(ctx context.Context, id string) error {
	// TODO: Remove <basePath>/<id> directory
	return nil
}

// CloneFromGit clones a preset from a git repository
// PLACEHOLDER: Not implemented yet
func (r *PresetRepository) CloneFromGit(ctx context.Context, url string, id string) error {
	// TODO: git clone <url> <basePath>/<id>
	return nil
}
