package tabbed

import (
	"strings"

	"github.com/bnema/archup-cli/internal/theme"
	"github.com/bnema/archup-cli/internal/tui/components"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type reviewModel struct {
	compositor string
	apps       []string
	hardware   HardwareInfo
	confirmed  bool
}

func newReviewModel() reviewModel {
	return reviewModel{}
}

func (m *reviewModel) setData(compositor string, apps []string, hardware HardwareInfo) {
	m.compositor = compositor
	m.apps = apps
	m.hardware = hardware
}

func (m reviewModel) Update(msg tea.Msg) (reviewModel, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			m.confirmed = true
			return m, nil
		}
	}
	return m, nil
}

func (m reviewModel) View() string {
	title := components.Header("Review & Confirm", "Review your selections before installation", 80)

	// Hardware section
	hardwareSection := lipgloss.NewStyle().
		Foreground(lipgloss.Color(theme.PrimaryText)).
		Render(lipgloss.JoinVertical(
			lipgloss.Left,
			theme.BodyStyle.Bold(true).Render("Detected Hardware:"),
			theme.BodyStyle.Render("  GPU:   "+m.hardware.GPU),
			theme.BodyStyle.Render("  CPU:   "+m.hardware.CPU),
			theme.BodyStyle.Render("  Audio: "+m.hardware.Audio),
		))

	// Compositor section
	compositorSection := lipgloss.NewStyle().
		Foreground(lipgloss.Color(theme.PrimaryText)).
		MarginTop(2).
		Render(lipgloss.JoinVertical(
			lipgloss.Left,
			theme.BodyStyle.Bold(true).Render("Selected Compositor:"),
			theme.BodyStyle.Render("  "+components.IconSuccess+" "+m.compositor),
		))

	// Applications section
	appsList := make([]string, len(m.apps))
	for i, app := range m.apps {
		appsList[i] = theme.BodyStyle.Render("  " + components.IconSuccess + " " + app)
	}

	applicationsSection := lipgloss.NewStyle().
		Foreground(lipgloss.Color(theme.PrimaryText)).
		MarginTop(2).
		Render(lipgloss.JoinVertical(
			lipgloss.Left,
			theme.BodyStyle.Bold(true).Render("Selected Applications:"),
			strings.Join(appsList, "\n"),
		))

	// Warning
	warning := lipgloss.NewStyle().
		Foreground(lipgloss.Color(theme.WarmOrange)).
		MarginTop(2).
		Render(components.IconWarning + " This will install packages to your system")

	// Buttons
	installButton := theme.ButtonPrimaryStyle.Render("Press Enter to Install")
	backButton := lipgloss.NewStyle().
		Foreground(lipgloss.Color(theme.DimmedText)).
		Render("Press tab/← to go back")

	buttons := lipgloss.JoinHorizontal(
		lipgloss.Top,
		installButton,
		"  ",
		backButton,
	)

	return lipgloss.JoinVertical(
		lipgloss.Left,
		title,
		"",
		hardwareSection,
		compositorSection,
		applicationsSection,
		"",
		warning,
		"",
		buttons,
	)
}
