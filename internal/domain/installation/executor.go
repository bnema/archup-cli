package installation

import "context"

// NOTE: This is a placeholder executor interface for running installation plans.
// Implementations will be in internal/infra/ and will handle actual system changes.

// Executor executes installation plans
type Executor interface {
	// Execute runs the entire plan
	Execute(ctx context.Context, plan *Plan) error

	// ExecuteStep runs a single step
	ExecuteStep(ctx context.Context, step Step) error

	// Rollback reverses a step's changes
	Rollback(ctx context.Context, step Step) error

	// Subscribe allows monitoring progress
	Subscribe(handler ProgressHandler)
}

// ProgressHandler handles progress updates
type ProgressHandler func(event ProgressEvent)

// ProgressEvent represents a progress update
type ProgressEvent struct {
	Type    EventType
	Step    *Step
	Message string
	Error   error
	Progress int // 0-100
}

// EventType categorizes progress events
type EventType string

const (
	EventStepStarted   EventType = "step-started"
	EventStepProgress  EventType = "step-progress"
	EventStepCompleted EventType = "step-completed"
	EventStepFailed    EventType = "step-failed"
	EventPlanCompleted EventType = "plan-completed"
	EventPlanFailed    EventType = "plan-failed"
)

// ExecutorOption configures the executor
type ExecutorOption func(*executorOptions)

type executorOptions struct {
	dryRun      bool
	continueOnError bool
	maxRetries  int
}

// WithDryRun enables dry-run mode (no actual changes)
// PLACEHOLDER: Not implemented yet
func WithDryRun() ExecutorOption {
	return func(o *executorOptions) {
		o.dryRun = true
	}
}

// WithContinueOnError continues execution even if non-critical steps fail
// PLACEHOLDER: Not implemented yet
func WithContinueOnError() ExecutorOption {
	return func(o *executorOptions) {
		o.continueOnError = true
	}
}

// WithMaxRetries sets the maximum number of retries for failed steps
// PLACEHOLDER: Not implemented yet
func WithMaxRetries(n int) ExecutorOption {
	return func(o *executorOptions) {
		o.maxRetries = n
	}
}
