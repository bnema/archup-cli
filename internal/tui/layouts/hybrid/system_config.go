package hybrid

import (
	"fmt"

	"github.com/bnema/archup-cli/internal/theme"
	"github.com/bnema/archup-cli/internal/tui/components"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type systemConfigModel struct {
	hardware HardwareInfo

	// Current focus
	focusIndex int
	fields     []string

	// Selections
	selectedCPUGov     string
	selectedTLPProfile string
	pacmanConfig       PacmanConfig

	// Lists for each field
	cpuGovList      components.ListModel
	tlpProfileList  components.ListModel
	parallelDlList  components.ListModel
	colorToggle     bool
	verboseToggle   bool
	candyToggle     bool

	confirmed bool
}

const (
	fieldCPUGov = iota
	fieldTLPProfile
	fieldParallelDownloads
	fieldColorOutput
	fieldVerbosePkgLists
	fieldILoveCandy
	fieldConfirm
	fieldCount
)

func newSystemConfigModel() systemConfigModel {
	// CPU Governor options
	cpuGovItems := []components.ListItem{
		{
			Title:       "performance",
			Description: "Maximum CPU performance (higher power consumption)",
		},
		{
			Title:       "powersave",
			Description: "Better battery life (lower performance)",
		},
		{
			Title:       "schedutil",
			Description: "Balanced - CPU scheduler-driven (recommended)",
		},
		{
			Title:       "ondemand",
			Description: "Dynamic scaling based on load",
		},
		{
			Title:       "conservative",
			Description: "Gradual scaling for smooth performance",
		},
	}

	// TLP profile options
	tlpProfileItems := []components.ListItem{
		{
			Title:       "Balanced",
			Description: "Balanced performance and battery life",
		},
		{
			Title:       "Performance",
			Description: "Prioritize performance over battery",
		},
		{
			Title:       "Battery",
			Description: "Maximize battery life",
		},
		{
			Title:       "Skip TLP",
			Description: "Don't configure TLP",
		},
	}

	// Pacman parallel downloads
	parallelItems := []components.ListItem{
		{Title: "3 downloads", Description: "Conservative"},
		{Title: "5 downloads", Description: "Recommended"},
		{Title: "10 downloads", Description: "Fast"},
	}

	m := systemConfigModel{
		fields:        make([]string, fieldCount),
		cpuGovList:    components.NewListModel(cpuGovItems),
		tlpProfileList: components.NewListModel(tlpProfileItems),
		parallelDlList: components.NewListModel(parallelItems),
		colorToggle:   true,
		verboseToggle: false,
		candyToggle:   false,
	}

	// Set default selections
	m.cpuGovList.SelectIndex(2)       // schedutil
	m.tlpProfileList.SelectIndex(0)   // Balanced
	m.parallelDlList.SelectIndex(1)   // 5 downloads

	return m
}

func (m systemConfigModel) Update(msg tea.Msg) (systemConfigModel, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "up", "k":
			if m.focusIndex > 0 {
				m.focusIndex--
			}
			return m, nil

		case "down", "j":
			if m.focusIndex < fieldCount-1 {
				m.focusIndex++
			}
			return m, nil

		case "left", "h", "right", "l", " ":
			return m.handleFieldInteraction()

		case "enter":
			if m.focusIndex == fieldConfirm {
				m.collectData()
				m.confirmed = true
				return m, nil
			}
			return m.handleFieldInteraction()
		}
	}

	// Update active list
	var cmd tea.Cmd
	switch m.focusIndex {
	case fieldCPUGov:
		m.cpuGovList, cmd = m.cpuGovList.Update(msg)
	case fieldTLPProfile:
		m.tlpProfileList, cmd = m.tlpProfileList.Update(msg)
	case fieldParallelDownloads:
		m.parallelDlList, cmd = m.parallelDlList.Update(msg)
	}

	return m, cmd
}

func (m systemConfigModel) handleFieldInteraction() (systemConfigModel, tea.Cmd) {
	switch m.focusIndex {
	case fieldColorOutput:
		m.colorToggle = !m.colorToggle
	case fieldVerbosePkgLists:
		m.verboseToggle = !m.verboseToggle
	case fieldILoveCandy:
		m.candyToggle = !m.candyToggle
	}
	return m, nil
}

func (m *systemConfigModel) collectData() {
	m.selectedCPUGov = m.cpuGovList.SelectedItem().Title
	m.selectedTLPProfile = m.tlpProfileList.SelectedItem().Title

	// Extract parallel downloads number
	dlText := m.parallelDlList.SelectedItem().Title
	switch dlText {
	case "3 downloads":
		m.pacmanConfig.ParallelDownloads = 3
	case "5 downloads":
		m.pacmanConfig.ParallelDownloads = 5
	case "10 downloads":
		m.pacmanConfig.ParallelDownloads = 10
	}

	m.pacmanConfig.ColorEnabled = m.colorToggle
	m.pacmanConfig.VerbosePkgLists = m.verboseToggle
	m.pacmanConfig.ILoveCandy = m.candyToggle
}

func (m systemConfigModel) View() string {
	title := components.Header(
		"System Configuration",
		"Configure CPU governor, TLP, and Pacman settings",
		80,
	)

	sections := []string{
		m.renderCPUSection(),
		m.renderTLPSection(),
		m.renderPacmanSection(),
		m.renderConfirmButton(),
	}

	content := lipgloss.JoinVertical(lipgloss.Left, sections...)

	return lipgloss.JoinVertical(
		lipgloss.Left,
		title,
		"",
		content,
	)
}

func (m systemConfigModel) renderCPUSection() string {
	sectionTitle := theme.BodyStyle.Bold(true).Render("CPU Governor")
	focused := m.focusIndex == fieldCPUGov

	var listView string
	if focused {
		listView = m.cpuGovList.View()
	} else {
		selected := m.cpuGovList.SelectedItem()
		listView = theme.HintStyle.Render("  → " + selected.Title)
	}

	indicator := " "
	if focused {
		indicator = theme.KeybindingStyle.Render("▶")
	}

	return lipgloss.JoinVertical(
		lipgloss.Left,
		indicator+" "+sectionTitle,
		listView,
		"",
	)
}

func (m systemConfigModel) renderTLPSection() string {
	if !m.hardware.IsLaptop {
		return theme.HintStyle.Render("  TLP: Skipped (not a laptop)")
	}

	sectionTitle := theme.BodyStyle.Bold(true).Render("TLP Profile (Laptop Power Management)")
	focused := m.focusIndex == fieldTLPProfile

	var listView string
	if focused {
		listView = m.tlpProfileList.View()
	} else {
		selected := m.tlpProfileList.SelectedItem()
		listView = theme.HintStyle.Render("  → " + selected.Title)
	}

	indicator := " "
	if focused {
		indicator = theme.KeybindingStyle.Render("▶")
	}

	return lipgloss.JoinVertical(
		lipgloss.Left,
		indicator+" "+sectionTitle,
		listView,
		"",
	)
}

func (m systemConfigModel) renderPacmanSection() string {
	sectionTitle := theme.BodyStyle.Bold(true).Render("Pacman Configuration")

	// Parallel downloads
	dlFocused := m.focusIndex == fieldParallelDownloads
	var dlView string
	if dlFocused {
		dlView = m.parallelDlList.View()
	} else {
		selected := m.parallelDlList.SelectedItem()
		dlView = theme.HintStyle.Render("  → " + selected.Title)
	}

	dlIndicator := " "
	if dlFocused {
		dlIndicator = theme.KeybindingStyle.Render("▶")
	}

	parallelSection := lipgloss.JoinVertical(
		lipgloss.Left,
		dlIndicator+" Parallel Downloads",
		dlView,
	)

	// Toggles
	colorLine := m.renderToggle("Color Output", fieldColorOutput, m.colorToggle)
	verboseLine := m.renderToggle("Verbose Package Lists", fieldVerbosePkgLists, m.verboseToggle)
	candyLine := m.renderToggle("ILoveCandy (Pac-Man animation)", fieldILoveCandy, m.candyToggle)

	return lipgloss.JoinVertical(
		lipgloss.Left,
		sectionTitle,
		parallelSection,
		"",
		colorLine,
		verboseLine,
		candyLine,
		"",
	)
}

func (m systemConfigModel) renderToggle(label string, fieldIndex int, enabled bool) string {
	focused := m.focusIndex == fieldIndex

	indicator := " "
	if focused {
		indicator = theme.KeybindingStyle.Render("▶")
	}

	checkbox := "[ ]"
	if enabled {
		checkbox = "[✓]"
	}

	checkStyle := theme.HintStyle
	if enabled {
		checkStyle = lipgloss.NewStyle().Foreground(lipgloss.Color(theme.SuccessGreen))
	}

	labelStyle := theme.BodyStyle
	if focused {
		labelStyle = theme.ListItemFocusedStyle
	}

	return fmt.Sprintf("%s %s %s",
		indicator,
		checkStyle.Render(checkbox),
		labelStyle.Render(label),
	)
}

func (m systemConfigModel) renderConfirmButton() string {
	focused := m.focusIndex == fieldConfirm

	button := "[ Continue ]"
	if focused {
		button = theme.ButtonPrimaryStyle.Render("[ Continue ]")
	} else {
		button = theme.HintStyle.Render(button)
	}

	indicator := " "
	if focused {
		indicator = theme.KeybindingStyle.Render("▶")
	}

	return lipgloss.JoinVertical(
		lipgloss.Left,
		"",
		indicator+" "+button,
	)
}
