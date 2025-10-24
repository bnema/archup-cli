package tabbed

import (
	"github.com/bnema/archup-cli/internal/tui/components"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// TabIndex represents which tab is currently active
type TabIndex int

const (
	TabSystem TabIndex = iota
	TabCompositor
	TabApplications
	TabReview
)

// Model represents the tabbed wizard state
type Model struct {
	activeTab TabIndex
	width     int
	height    int

	// Tab models
	systemModel       systemModel
	compositorModel   compositorModel
	applicationsModel applicationsModel
	reviewModel       reviewModel
	installationModel installationModel

	// Shared state
	selectedCompositor string
	selectedApps       []string
	detectedHardware   HardwareInfo

	// Installation state
	installing bool
	complete   bool

	quitting bool
}

// HardwareInfo stores detected hardware information
type HardwareInfo struct {
	GPU   string
	CPU   string
	Audio string
}

// New creates a new tabbed wizard model
func New() Model {
	return Model{
		activeTab:         TabSystem,
		width:             120,
		height:            30,
		systemModel:       newSystemModel(),
		compositorModel:   newCompositorModel(),
		applicationsModel: newApplicationsModel(),
		reviewModel:       newReviewModel(),
		installationModel: newInstallationModel(),
	}
}

// Init initializes the model
func (m Model) Init() tea.Cmd {
	return m.systemModel.Init()
}

// Update handles messages and updates the model
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// Handle installation updates first
	if m.installing {
		var cmd tea.Cmd
		m.installationModel, cmd = m.installationModel.Update(msg)
		if m.installationModel.Complete {
			m.complete = true
			m.installing = false
		}
		return m, cmd
	}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			m.quitting = true
			return m, tea.Quit

		case "tab", "right", "l":
			// Move to next tab
			if m.activeTab < TabReview {
				// Update shared state before switching
				m.updateSharedState()
				m.activeTab++
			}
			return m, nil

		case "shift+tab", "left", "h":
			// Move to previous tab
			if m.activeTab > TabSystem {
				// Update shared state before switching
				m.updateSharedState()
				m.activeTab--
			}
			return m, nil
		}

	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		return m, nil
	}

	// Delegate to current tab's update
	return m.updateCurrentTab(msg)
}

// updateSharedState updates the shared state from all tab models
func (m *Model) updateSharedState() {
	// Update hardware info from system tab
	if m.systemModel.complete {
		m.detectedHardware = m.systemModel.hardware
	}

	// Update compositor selection
	if m.compositorModel.selected != "" {
		m.selectedCompositor = m.compositorModel.selected
	}

	// Update app selections
	m.selectedApps = m.applicationsModel.checkboxList.GetSelected()
}

// updateCurrentTab delegates update to the current tab
func (m Model) updateCurrentTab(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch m.activeTab {
	case TabSystem:
		m.systemModel, cmd = m.systemModel.Update(msg)
		if m.systemModel.complete {
			m.detectedHardware = m.systemModel.hardware
		}

	case TabCompositor:
		// Pass window size to compositor list
		if _, ok := msg.(tea.WindowSizeMsg); ok {
			msg = tea.WindowSizeMsg{Width: m.width, Height: m.height}
		}
		m.compositorModel, cmd = m.compositorModel.Update(msg)
		if m.compositorModel.selected != "" {
			m.selectedCompositor = m.compositorModel.selected
		}

	case TabApplications:
		m.applicationsModel, cmd = m.applicationsModel.Update(msg)
		m.selectedApps = m.applicationsModel.checkboxList.GetSelected()

	case TabReview:
		m.reviewModel.setData(m.selectedCompositor, m.selectedApps, m.detectedHardware)
		m.reviewModel, cmd = m.reviewModel.Update(msg)
		if m.reviewModel.confirmed {
			m.installing = true
			return m, m.installationModel.Init()
		}
	}

	return m, cmd
}

// View renders the tabbed interface
func (m Model) View() string {
	if m.quitting {
		return ""
	}

	// If installing or complete, show installation view instead of tabs
	if m.installing || m.complete {
		return m.installationModel.View()
	}

	// Calculate content height (total - tabs - footer - padding)
	tabsHeight := 3
	footerHeight := 3
	contentHeight := m.height - tabsHeight - footerHeight - 2

	// Render tab bar
	tabs := []components.Tab{
		{Name: "System", Active: m.activeTab == TabSystem},
		{Name: "Compositor", Active: m.activeTab == TabCompositor},
		{Name: "Applications", Active: m.activeTab == TabApplications},
		{Name: "Review", Active: m.activeTab == TabReview},
	}
	tabBar := components.RenderTabs(tabs, m.width)

	// Render current tab content
	var content string
	switch m.activeTab {
	case TabSystem:
		content = m.systemModel.View()
	case TabCompositor:
		content = m.compositorModel.View()
	case TabApplications:
		content = m.applicationsModel.View()
	case TabReview:
		// Update review model with latest data before rendering
		m.reviewModel.setData(m.selectedCompositor, m.selectedApps, m.detectedHardware)
		content = m.reviewModel.View()
	}

	// Create fixed-height container for content
	contentStyle := lipgloss.NewStyle().
		Height(contentHeight).
		Width(m.width)

	content = contentStyle.Render(content)

	// Footer with keybindings
	footer := components.Footer([]components.Keybinding{
		{Key: "tab/←→", Desc: "switch tabs"},
		{Key: "q", Desc: "quit"},
	}, m.width)

	return lipgloss.JoinVertical(lipgloss.Left, tabBar, "", content, footer)
}
