package preset

// NOTE: This is a placeholder domain aggregate for presets/themes.
// Presets define collections of configurations, dotfiles, and scripts.

// Preset represents a complete system configuration preset
type Preset struct {
	ID          string
	Name        string
	Description string
	Author      string
	Version     string

	// Configuration files to deploy
	Dotfiles []Dotfile

	// Scripts to run during installation
	Scripts []Script

	// Packages specific to this preset
	Packages []string

	// Theme configuration
	Theme Theme
}

// Dotfile represents a configuration file to deploy
type Dotfile struct {
	Source      string // Path in preset repo
	Destination string // Target path on system
	Template    bool   // If true, render as template
	Backup      bool   // If true, backup existing file
}

// Script represents a script to execute
type Script struct {
	Name        string
	Description string
	Path        string
	Phase       Phase // When to run: pre-install, post-install, etc.
	Critical    bool  // If true, abort installation on failure
}

// Phase represents when a script should run
type Phase string

const (
	PhasePreInstall  Phase = "pre-install"
	PhasePostInstall Phase = "post-install"
	PhasePreConfig   Phase = "pre-config"
	PhasePostConfig  Phase = "post-config"
)

// Theme contains visual theme configuration
type Theme struct {
	Name        string
	ColorScheme string
	GTKTheme    string
	IconTheme   string
	CursorTheme string
	FontFamily  string
}

// NewPreset creates a new preset
// PLACEHOLDER: Not implemented yet
func NewPreset(id, name string) *Preset {
	return &Preset{
		ID:       id,
		Name:     name,
		Dotfiles: []Dotfile{},
		Scripts:  []Script{},
		Packages: []string{},
	}
}

// AddDotfile adds a dotfile to the preset
// PLACEHOLDER: Not implemented yet
func (p *Preset) AddDotfile(df Dotfile) {
	p.Dotfiles = append(p.Dotfiles, df)
}

// AddScript adds a script to the preset
// PLACEHOLDER: Not implemented yet
func (p *Preset) AddScript(s Script) {
	p.Scripts = append(p.Scripts, s)
}

// Validate checks if the preset is valid
// PLACEHOLDER: Not implemented yet
func (p *Preset) Validate() error {
	return nil
}

// GetScriptsByPhase returns scripts for a specific phase
// PLACEHOLDER: Not implemented yet
func (p *Preset) GetScriptsByPhase(phase Phase) []Script {
	var scripts []Script
	for _, s := range p.Scripts {
		if s.Phase == phase {
			scripts = append(scripts, s)
		}
	}
	return scripts
}
