package preset

import "context"

// NOTE: This is a placeholder repository interface for preset management.
// Implementations will be in internal/infra/filesystem/ or could fetch from git repos.

// Repository provides access to presets
type Repository interface {
	// Get retrieves a preset by ID
	Get(ctx context.Context, id string) (*Preset, error)

	// List returns all available presets
	List(ctx context.Context) ([]Preset, error)

	// Save persists a preset
	Save(ctx context.Context, preset *Preset) error

	// Delete removes a preset
	Delete(ctx context.Context, id string) error
}

// BuiltinPresets returns the built-in presets (Omarchy, Bleu, Stock)
// PLACEHOLDER: Not implemented yet
func BuiltinPresets() []Preset {
	return []Preset{
		{
			ID:          "omarchy",
			Name:        "Omarchy Defaults",
			Description: "Complete Omarchy theme with curated scripts, configs, and dotfiles",
			Author:      "bnema",
			Version:     "1.0.0",
		},
		{
			ID:          "bleu",
			Name:        "ArchUp Bleu Theme",
			Description: "Modern Bleu-themed configuration with clean aesthetics",
			Author:      "bnema",
			Version:     "1.0.0",
		},
		{
			ID:          "stock",
			Name:        "Stock Configuration",
			Description: "Vanilla Arch Linux configuration with minimal customization",
			Author:      "archlinux",
			Version:     "1.0.0",
		},
	}
}
