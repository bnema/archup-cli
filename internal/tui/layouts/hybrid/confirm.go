package hybrid

import (
	"fmt"
	"strings"

	"github.com/bnema/archup-cli/internal/theme"
	"github.com/bnema/archup-cli/internal/tui/components"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type confirmationModel struct {
	themeChoice    string
	cpuGov         string
	tlpProfile     string
	pacmanCfg      PacmanConfig
	compositor     string
	packages       []string
	hardware       HardwareInfo
	confirmed      bool
}

func newConfirmationModel() confirmationModel {
	return confirmationModel{}
}

func (m *confirmationModel) setData(
	theme string,
	cpuGov string,
	tlpProfile string,
	pacmanCfg PacmanConfig,
	compositor string,
	packages []string,
	hardware HardwareInfo,
) {
	m.themeChoice = theme
	m.cpuGov = cpuGov
	m.tlpProfile = tlpProfile
	m.pacmanCfg = pacmanCfg
	m.compositor = compositor
	m.packages = packages
	m.hardware = hardware
}

func (m confirmationModel) Update(msg tea.Msg) (confirmationModel, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter", " ":
			m.confirmed = true
			return m, nil
		}
	}
	return m, nil
}

func (m confirmationModel) View() string {
	title := components.Header("Confirmation", "Review your selections before installation", 80)

	sections := []string{
		m.renderHardwareSection(),
		m.renderThemeSection(),
		m.renderSystemConfigSection(),
		m.renderCompositorSection(),
		m.renderPackagesSection(),
		m.renderWarningAndButtons(),
	}

	content := lipgloss.JoinVertical(lipgloss.Left, sections...)

	return lipgloss.JoinVertical(
		lipgloss.Left,
		title,
		"",
		content,
	)
}

func (m confirmationModel) renderHardwareSection() string {
	title := theme.SubHeaderStyle.Render("Detected Hardware:")

	lines := []string{
		fmt.Sprintf("  GPU:   %s", m.hardware.GPU),
		fmt.Sprintf("  CPU:   %s", m.hardware.CPU),
		fmt.Sprintf("  Audio: %s", m.hardware.Audio),
	}

	if m.hardware.IsLaptop {
		lines = append(lines, "  Type:  Laptop")
	}

	content := theme.BodyStyle.Render(strings.Join(lines, "\n"))

	return lipgloss.JoinVertical(lipgloss.Left, title, content, "")
}

func (m confirmationModel) renderThemeSection() string {
	title := theme.SubHeaderStyle.Render("Theme & Configuration:")
	content := theme.BodyStyle.Render("  " + m.themeChoice)
	return lipgloss.JoinVertical(lipgloss.Left, title, content, "")
}

func (m confirmationModel) renderSystemConfigSection() string {
	title := theme.SubHeaderStyle.Render("System Configuration:")

	lines := []string{
		fmt.Sprintf("  CPU Governor: %s", m.cpuGov),
	}

	if m.hardware.IsLaptop && m.tlpProfile != "Skip TLP" {
		lines = append(lines, fmt.Sprintf("  TLP Profile:  %s", m.tlpProfile))
	}

	lines = append(lines, "  Pacman:")
	lines = append(lines, fmt.Sprintf("    • Parallel downloads: %d", m.pacmanCfg.ParallelDownloads))
	lines = append(lines, fmt.Sprintf("    • Color output: %v", m.pacmanCfg.ColorEnabled))
	lines = append(lines, fmt.Sprintf("    • Verbose package lists: %v", m.pacmanCfg.VerbosePkgLists))
	if m.pacmanCfg.ILoveCandy {
		lines = append(lines, "    • ILoveCandy: true (Pac-Man animation enabled!)")
	}

	content := theme.BodyStyle.Render(strings.Join(lines, "\n"))

	return lipgloss.JoinVertical(lipgloss.Left, title, content, "")
}

func (m confirmationModel) renderCompositorSection() string {
	title := theme.SubHeaderStyle.Render("Selected Compositor:")
	content := theme.BodyStyle.Render("  " + m.compositor)
	return lipgloss.JoinVertical(lipgloss.Left, title, content, "")
}

func (m confirmationModel) renderPackagesSection() string {
	title := theme.SubHeaderStyle.Render(fmt.Sprintf("Selected Packages (%d):", len(m.packages)))

	// Show first 10 packages, then "and N more..."
	displayLimit := 10
	var lines []string
	for i, pkg := range m.packages {
		if i >= displayLimit {
			remaining := len(m.packages) - displayLimit
			lines = append(lines, fmt.Sprintf("  ... and %d more packages", remaining))
			break
		}
		lines = append(lines, "  "+components.IconSuccess+" "+pkg)
	}

	content := theme.BodyStyle.Render(strings.Join(lines, "\n"))

	return lipgloss.JoinVertical(lipgloss.Left, title, content, "")
}

func (m confirmationModel) renderWarningAndButtons() string {
	warning := theme.WarningStyle.Render("⚠ This will install packages and modify system configuration")

	continueBtn := theme.ButtonPrimaryStyle.Render(" Press Enter to Install ")
	cancelHint := theme.HintStyle.Render("Press Esc to go back")

	buttons := lipgloss.JoinHorizontal(lipgloss.Left, continueBtn, "  ", cancelHint)

	return lipgloss.JoinVertical(
		lipgloss.Left,
		"",
		warning,
		"",
		buttons,
	)
}
