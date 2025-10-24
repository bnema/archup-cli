package linear

import (
	"strings"

	"github.com/bnema/archup-cli/internal/theme"
	"github.com/bnema/archup-cli/internal/tui/components"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type confirmationModel struct {
	compositor string
	apps       []string
	hardware   HardwareInfo
	confirmed  bool
}

func newConfirmationModel() confirmationModel {
	return confirmationModel{}
}

func (m *confirmationModel) setData(compositor string, apps []string, hardware HardwareInfo) {
	m.compositor = compositor
	m.apps = apps
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

	// Hardware section
	hardwareTitle := theme.SubHeaderStyle.Render("Detected Hardware:")
	hardwareContent := lipgloss.JoinVertical(
		lipgloss.Left,
		theme.BodyStyle.Render("  GPU:   "+m.hardware.GPU),
		theme.BodyStyle.Render("  CPU:   "+m.hardware.CPU),
		theme.BodyStyle.Render("  Audio: "+m.hardware.Audio),
	)

	// Compositor section
	compositorTitle := theme.SubHeaderStyle.Render("Selected Compositor:")
	compositorContent := theme.BodyStyle.Render("  " + m.compositor)

	// Applications section
	appsTitle := theme.SubHeaderStyle.Render("Selected Applications:")
	appsLines := make([]string, len(m.apps))
	for i, app := range m.apps {
		appsLines[i] = theme.BodyStyle.Render("  " + components.IconSuccess + " " + app)
	}
	appsContent := strings.Join(appsLines, "\n")

	// Warning
	warning := theme.WarningStyle.Render("⚠ This will install packages and modify system configuration")

	// Buttons
	continueBtn := theme.ButtonPrimaryStyle.Render(" Press Enter to Install ")
	cancelHint := theme.HintStyle.Render("Press Esc to go back")

	buttons := lipgloss.JoinHorizontal(lipgloss.Left, continueBtn, "  ", cancelHint)

	return lipgloss.JoinVertical(
		lipgloss.Left,
		title,
		"",
		hardwareTitle,
		hardwareContent,
		"",
		compositorTitle,
		compositorContent,
		"",
		appsTitle,
		appsContent,
		"",
		"",
		warning,
		"",
		buttons,
	)
}
