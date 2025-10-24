package hybrid

import (
	"github.com/bnema/archup-cli/internal/tui/components"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// WizardStep represents each step in the hybrid wizard
type WizardStep int

const (
	StepWelcome WizardStep = iota
	StepHardware
	StepSystemConfig
	StepCompositor
	StepPackageSelection
	StepThemeConfig
	StepConfirmation
	StepInstallation
	StepComplete
)

// Model represents the hybrid wizard state
type Model struct {
	currentStep WizardStep
	totalSteps  int
	width       int
	height      int

	// Screen models
	welcomeModel       welcomeModel
	hardwareModel      hardwareModel
	themeConfigModel   themeConfigModel
	systemConfigModel  systemConfigModel
	compositorModel    compositorModel
	packageModel       packageSelectionModel
	confirmationModel  confirmationModel
	installationModel  installationModel

	// Shared state - collected from user
	selectedTheme      string // "omarchy", "bleu", "stock"
	selectedCPUGov     string
	selectedTLPProfile string
	pacmanConfig       PacmanConfig
	selectedCompositor string
	selectedPackages   []string
	detectedHardware   HardwareInfo

	quitting bool
}

// HardwareInfo stores detected hardware information
type HardwareInfo struct {
	GPU       string
	CPU       string
	CPUModel  string
	Audio     string
	IsLaptop  bool
	MinCPUFreq int
	MaxCPUFreq int
}

// PacmanConfig stores pacman configuration choices
type PacmanConfig struct {
	ParallelDownloads int
	ColorEnabled      bool
	VerbosePkgLists   bool
	ILoveCandy        bool // Pac-Man easter egg
}

// New creates a new hybrid wizard model
func New() Model {
	return Model{
		currentStep: StepWelcome,
		totalSteps:  7,
		width:       120,
		height:      30,
		welcomeModel:      newWelcomeModel(),
		hardwareModel:     newHardwareModel(),
		themeConfigModel:  newThemeConfigModel(),
		systemConfigModel: newSystemConfigModel(),
		compositorModel:   newCompositorModel(),
		packageModel:      newPackageSelectionModel(),
		confirmationModel: newConfirmationModel(),
		installationModel: newInstallationModel(),
		pacmanConfig: PacmanConfig{
			ParallelDownloads: 5,
			ColorEnabled:      true,
			VerbosePkgLists:   false,
			ILoveCandy:        false,
		},
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
			// Don't quit during installation
			if m.currentStep != StepInstallation {
				m.quitting = true
				return m, tea.Quit
			}

		case "esc":
			// Go back one step (except on welcome screen and during installation)
			if m.currentStep > StepWelcome && m.currentStep != StepInstallation {
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
			m.currentStep = StepSystemConfig
			// Pass hardware info to system config
			m.systemConfigModel.hardware = m.detectedHardware
		}

	case StepSystemConfig:
		m.systemConfigModel, cmd = m.systemConfigModel.Update(msg)
		if m.systemConfigModel.confirmed {
			m.selectedCPUGov = m.systemConfigModel.selectedCPUGov
			m.selectedTLPProfile = m.systemConfigModel.selectedTLPProfile
			m.pacmanConfig = m.systemConfigModel.pacmanConfig
			m.currentStep = StepCompositor
		}

	case StepCompositor:
		m.compositorModel, cmd = m.compositorModel.Update(msg)
		if m.compositorModel.confirmed {
			m.selectedCompositor = m.compositorModel.selected
			m.currentStep = StepPackageSelection
		}

	case StepPackageSelection:
		m.packageModel, cmd = m.packageModel.Update(msg)
		if m.packageModel.confirmed {
			m.selectedPackages = m.packageModel.tree.Root.GetSelectedItems()
			m.currentStep = StepThemeConfig
		}

	case StepThemeConfig:
		m.themeConfigModel, cmd = m.themeConfigModel.Update(msg)
		if m.themeConfigModel.confirmed {
			m.selectedTheme = m.themeConfigModel.selected
			m.currentStep = StepConfirmation
			m.confirmationModel.setData(
				m.selectedTheme,
				m.selectedCPUGov,
				m.selectedTLPProfile,
				m.pacmanConfig,
				m.selectedCompositor,
				m.selectedPackages,
				m.detectedHardware,
			)
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
	case StepSystemConfig:
		content = m.systemConfigModel.View()
	case StepCompositor:
		content = m.compositorModel.View()
	case StepPackageSelection:
		content = m.packageModel.ViewWithSize(m.width, contentHeight)
	case StepThemeConfig:
		content = m.themeConfigModel.View()
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
	footer := components.StandardFooter(m.width, m.currentStep > StepWelcome && m.currentStep != StepInstallation)

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
