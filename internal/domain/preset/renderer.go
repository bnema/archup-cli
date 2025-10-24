package preset

import "context"

// NOTE: This is a placeholder for template rendering logic.
// Will handle rendering dotfile templates with user-specific variables.

// Renderer renders preset templates with context
type Renderer interface {
	// Render renders a template with given context
	Render(ctx context.Context, template string, data RenderContext) (string, error)

	// RenderFile renders a file template
	RenderFile(ctx context.Context, path string, data RenderContext) (string, error)
}

// RenderContext contains variables for template rendering
type RenderContext struct {
	Username  string
	Hostname  string
	HomeDir   string
	Compositor string
	Theme      Theme
	Custom     map[string]interface{}
}

// NewRenderContext creates a render context
// PLACEHOLDER: Not implemented yet
func NewRenderContext(username, hostname string) *RenderContext {
	return &RenderContext{
		Username: username,
		Hostname: hostname,
		Custom:   make(map[string]interface{}),
	}
}

// Set sets a custom variable
// PLACEHOLDER: Not implemented yet
func (rc *RenderContext) Set(key string, value interface{}) {
	rc.Custom[key] = value
}

// Get gets a custom variable
// PLACEHOLDER: Not implemented yet
func (rc *RenderContext) Get(key string) (interface{}, bool) {
	val, ok := rc.Custom[key]
	return val, ok
}
