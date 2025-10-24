package hybrid

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
Welcome to the ArchUp Hybrid Wizard Demo!

This demo showcases a hybrid approach combining:
  • Linear step-by-step wizard for focused questions
  • Tree view for hierarchical package selection

The wizard flow:
  1. Hardware detection
  2. System configuration (CPU governor, TLP, Pacman)
  3. Compositor selection
  4. Package selection (tree view)
  5. Theme/config selection (Omarchy vs Bleu vs Stock)
  6. Review & installation

Ready to explore?
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
