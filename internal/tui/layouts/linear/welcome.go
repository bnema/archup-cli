package linear

import (
	"github.com/bnema/archup-cli/internal/theme"
	"github.com/bnema/archup-cli/internal/tui/components"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type welcomeModel struct {
	confirmed bool
}

func newWelcomeModel() welcomeModel {
	return welcomeModel{}
}

func (m welcomeModel) Update(msg tea.Msg) (welcomeModel, tea.Cmd) {
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

func (m welcomeModel) View() string {
	logo := components.HeaderWithLogo(100)

	introText := `
Welcome to the ArchUp Desktop Environment Setup Wizard!

This wizard will guide you through:

  1. Hardware detection (GPU, CPU, Audio)
  2. Compositor selection (Wayland compositor)
  3. Application selection (terminal, browser, etc.)
  4. Configuration and installation

The wizard will automatically install and configure your selected
desktop environment with sensible defaults.

Ready to begin?
`

	introStyle := theme.BodyStyle.
		Width(80).
		Align(lipgloss.Left).
		Padding(1, 2)

	continuePrompt := theme.ButtonPrimaryStyle.Render(" Press Enter to Continue ")

	return lipgloss.JoinVertical(
		lipgloss.Center,
		logo,
		"",
		"",
		introStyle.Render(introText),
		"",
		continuePrompt,
	)
}
