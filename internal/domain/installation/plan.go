package installation

// NOTE: This is a placeholder for installation plan domain logic.
// The Plan is an ordered list of steps that will be executed.

// Plan represents an ordered installation plan
type Plan struct {
	Steps []Step
}

// Step represents a single installation step
type Step struct {
	ID          string
	Name        string
	Description string
	Type        StepType
	Action      Action
	Required    bool // If true, failure aborts installation
	Rollback    Action // Optional rollback action
}

// StepType categorizes installation steps
type StepType string

const (
	StepTypePackageInstall StepType = "package-install"
	StepTypeConfigDeploy   StepType = "config-deploy"
	StepTypeScriptRun      StepType = "script-run"
	StepTypeServiceEnable  StepType = "service-enable"
	StepTypeFileWrite      StepType = "file-write"
)

// Action represents an executable action
type Action interface {
	// Execute runs the action
	Execute() error

	// Estimate returns estimated time in seconds
	Estimate() int
}

// NewPlan creates a new empty plan
// PLACEHOLDER: Not implemented yet
func NewPlan() *Plan {
	return &Plan{
		Steps: []Step{},
	}
}

// AddStep adds a step to the plan
// PLACEHOLDER: Not implemented yet
func (p *Plan) AddStep(step Step) {
	p.Steps = append(p.Steps, step)
}

// Validate checks if the plan is valid
// PLACEHOLDER: Not implemented yet
func (p *Plan) Validate() error {
	return nil
}

// EstimatedDuration returns total estimated time in seconds
// PLACEHOLDER: Not implemented yet
func (p *Plan) EstimatedDuration() int {
	total := 0
	for _, step := range p.Steps {
		if step.Action != nil {
			total += step.Action.Estimate()
		}
	}
	return total
}

// TotalSteps returns the number of steps
// PLACEHOLDER: Not implemented yet
func (p *Plan) TotalSteps() int {
	return len(p.Steps)
}
