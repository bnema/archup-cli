package packages

import "context"

// NOTE: This is a placeholder repository interface for package queries.
// Implementations will be in internal/infra/pacman/

// Repository provides access to package information
type Repository interface {
	// Search finds packages matching a query
	Search(ctx context.Context, query string) ([]Package, error)

	// Get retrieves detailed information about a package
	Get(ctx context.Context, name string) (*Package, error)

	// ListInstalled returns all currently installed packages
	ListInstalled(ctx context.Context) ([]Package, error)

	// ResolveDependencies resolves all dependencies for given packages
	ResolveDependencies(ctx context.Context, packages []string) ([]string, error)

	// CheckConflicts checks if packages have conflicts
	CheckConflicts(ctx context.Context, packages []string) ([]Conflict, error)
}

// Conflict represents a package conflict
type Conflict struct {
	Package1 string
	Package2 string
	Reason   string
}

// RepositoryOption configures the repository
type RepositoryOption func(*repositoryOptions)

type repositoryOptions struct {
	includeAUR bool
	cacheTime  int
}

// WithAUR includes AUR packages in queries
// PLACEHOLDER: Not implemented yet
func WithAUR() RepositoryOption {
	return func(o *repositoryOptions) {
		o.includeAUR = true
	}
}

// WithCacheTime sets the cache duration for package info
// PLACEHOLDER: Not implemented yet
func WithCacheTime(seconds int) RepositoryOption {
	return func(o *repositoryOptions) {
		o.cacheTime = seconds
	}
}
