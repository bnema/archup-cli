package packages

// NOTE: This is a placeholder domain model for package selection and management.
// Will contain business logic for package dependencies, conflicts, and resolution.

// Selection represents a user's package selection
type Selection struct {
	BaseSystem   []string
	Compositor   string
	Applications []Application
	DevTools     []string
}

// Application represents an application with metadata
type Application struct {
	Name         string
	Category     string
	Dependencies []string
	Conflicts    []string
	Optional     []string
}

// Category represents a package category
type Category struct {
	Name        string
	Description string
	Packages    []Package
	Required    bool // If true, user must select at least one
}

// Package represents a single package
type Package struct {
	Name         string
	Description  string
	Repository   string // "core", "extra", "aur"
	Dependencies []string
	Conflicts    []string
	Optional     []string
	Size         int64 // In bytes
}

// Add adds a package to the selection
// PLACEHOLDER: Not implemented yet
func (s *Selection) Add(pkg string) error {
	return nil
}

// Remove removes a package from the selection
// PLACEHOLDER: Not implemented yet
func (s *Selection) Remove(pkg string) error {
	return nil
}

// Resolve resolves all dependencies for the selection
// PLACEHOLDER: Not implemented yet
func (s *Selection) Resolve() ([]string, error) {
	return []string{}, nil
}

// Validate checks for conflicts and missing dependencies
// PLACEHOLDER: Not implemented yet
func (s *Selection) Validate() error {
	return nil
}

// TotalSize calculates the total download size
// PLACEHOLDER: Not implemented yet
func (s *Selection) TotalSize() int64 {
	return 0
}
