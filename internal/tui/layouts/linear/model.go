package linear

import (
	"github.com/bnema/archup-cli/internal/tui/components"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// WizardStep represents each step in the linear wizard
type WizardStep int

const (
	StepWelcome WizardStep = iota
	StepHardware
	StepCompositor
	StepApplications
	StepConfirmation
	StepInstallation
	StepComplete
)

// Model represents the linear wizard state
type Model struct {
	currentStep WizardStep
	totalSteps  int
	width       int
	height      int

	// Screen models
	welcomeModel       welcomeModel
	hardwareModel      hardwareModel
	compositorModel    compositorModel
	applicationsModel  applicationsModel
	confirmationModel  confirmationModel
	installationModel  installationModel

	// Shared state
	selectedCompositor string
	selectedApps       []string
	detectedHardware   HardwareInfo

	quitting bool
}

// HardwareInfo stores detected hardware information
type HardwareInfo struct {
	GPU   string
	CPU   string
	Audio string
}

// New creates a new linear wizard model
func New() Model {
	return Model{
		currentStep: StepWelcome,
		totalSteps:  6,
		width:       120,
		height:      30,
		welcomeModel: newWelcomeModel(),
		hardwareModel: newHardwareModel(),
		compositorModel: newCompositorModel(),
		applicationsModel: newApplicationsModel(),
		confirmationModel: newConfirmationModel(),
		installationModel: newInstallationModel(),
	}
}

// Init initializes the model
func (m Model) Init() tea.Cmd {
	return nil
}

// Update handles messages and updates the model
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			m.quitting = true
			return m, tea.Quit

		case "esc":
			// Go back one step (except on welcome screen)
			if m.currentStep > StepWelcome {
				m.currentStep--
			}
			return m, nil
		}

	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		return m, nil
	}

	// Delegate to current step's update
	return m.updateCurrentStep(msg)
}

// updateCurrentStep delegates update to the current step
func (m Model) updateCurrentStep(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch m.currentStep {
	case StepWelcome:
		m.welcomeModel, cmd = m.welcomeModel.Update(msg)
		if m.welcomeModel.confirmed {
			m.currentStep = StepHardware
			return m, m.hardwareModel.Init()
		}

	case StepHardware:
		m.hardwareModel, cmd = m.hardwareModel.Update(msg)
		if m.hardwareModel.complete {
			m.detectedHardware = m.hardwareModel.hardware
			m.currentStep = StepCompositor
		}

	case StepCompositor:
		// Pass window size to compositor list
		if _, ok := msg.(tea.WindowSizeMsg); ok {
			msg = tea.WindowSizeMsg{Width: m.width, Height: m.height}
		}
		m.compositorModel, cmd = m.compositorModel.Update(msg)
		if m.compositorModel.confirmed {
			m.selectedCompositor = m.compositorModel.selected
			m.currentStep = StepApplications
		}

	case StepApplications:
		m.applicationsModel, cmd = m.applicationsModel.Update(msg)
		if m.applicationsModel.checkboxList.Finished {
			m.selectedApps = m.applicationsModel.selectedApps
			m.currentStep = StepConfirmation
			m.confirmationModel.setData(m.selectedCompositor, m.selectedApps, m.detectedHardware)
		}

	case StepConfirmation:
		m.confirmationModel, cmd = m.confirmationModel.Update(msg)
		if m.confirmationModel.confirmed {
			m.currentStep = StepInstallation
			return m, m.installationModel.Init()
		}

	case StepInstallation:
		m.installationModel, cmd = m.installationModel.Update(msg)
		if m.installationModel.Complete {
			m.currentStep = StepComplete
		}
	}

	return m, cmd
}

// View renders the current step
func (m Model) View() string {
	if m.quitting {
		return ""
	}

	// Calculate content height (total height - footer - progress - padding)
	footerHeight := 3
	progressHeight := 2
	contentHeight := m.height - footerHeight - progressHeight - 2

	var content string

	// Render current step
	switch m.currentStep {
	case StepWelcome:
		content = m.welcomeModel.View()
	case StepHardware:
		content = m.hardwareModel.View()
	case StepCompositor:
		content = m.compositorModel.View()
	case StepApplications:
		content = m.applicationsModel.View()
	case StepConfirmation:
		content = m.confirmationModel.View()
	case StepInstallation:
		content = m.installationModel.View()
	case StepComplete:
		content = m.viewComplete()
	}

	// Create fixed-height container for content
	contentStyle := lipgloss.NewStyle().
		Height(contentHeight).
		Width(m.width)

	content = contentStyle.Render(content)

	// Add progress indicator at top (except welcome and complete)
	if m.currentStep > StepWelcome && m.currentStep < StepComplete {
		progress := components.ProgressIndicator(int(m.currentStep), m.totalSteps, m.width)
		content = lipgloss.JoinVertical(lipgloss.Left, progress, "", content)
	}

	// Add footer with keybindings
	footer := components.StandardFooter(m.width, m.currentStep > StepWelcome)

	return lipgloss.JoinVertical(lipgloss.Left, content, footer)
}

// viewComplete renders the completion screen
func (m Model) viewComplete() string {
	msg := components.StatusMessage{
		Type:    components.StatusSuccess,
		Message: "Installation complete! Your desktop environment is ready.",
	}
	return components.RenderStatus(msg)
}
