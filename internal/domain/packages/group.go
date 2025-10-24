package packages

// NOTE: This is a placeholder for package group domain logic.
// Groups organize packages into logical categories (terminals, browsers, etc.)

// Group represents a logical grouping of packages
type Group struct {
	ID          string
	Name        string
	Description string
	Packages    []Package
	IsRadio     bool // If true, only one package can be selected from this group
	Required    bool // If true, at least one package must be selected
}

// NewGroup creates a new package group
// PLACEHOLDER: Not implemented yet
func NewGroup(id, name, description string) *Group {
	return &Group{
		ID:          id,
		Name:        name,
		Description: description,
		Packages:    []Package{},
	}
}

// AddPackage adds a package to the group
// PLACEHOLDER: Not implemented yet
func (g *Group) AddPackage(pkg Package) {
	g.Packages = append(g.Packages, pkg)
}

// Validate checks group constraints (radio, required, etc.)
// PLACEHOLDER: Not implemented yet
func (g *Group) Validate(selectedPackages []string) error {
	return nil
}

// PredefinedGroups returns common package groups
// PLACEHOLDER: Not implemented yet - will load from config/database
func PredefinedGroups() []Group {
	return []Group{
		{
			ID:          "terminals",
			Name:        "Terminal Emulators",
			Description: "Choose your terminal application",
			Required:    true,
		},
		{
			ID:          "browsers",
			Name:        "Web Browsers",
			Description: "Internet browsers",
			Required:    false,
		},
		{
			ID:          "file-managers",
			Name:        "File Managers",
			Description: "GUI file browsers",
			Required:    false,
		},
	}
}
