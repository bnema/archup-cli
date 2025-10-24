package hybrid

import (
	"github.com/bnema/archup-cli/internal/theme"
	"github.com/bnema/archup-cli/internal/tui/components"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type themeConfigModel struct {
	list      components.ListModel
	selected  string
	confirmed bool
}

func newThemeConfigModel() themeConfigModel {
	items := []components.ListItem{
		{
			Title: "Omarchy Defaults",
			Description: "Complete Omarchy theme with curated scripts, configs, and dotfiles.\n" +
				"Includes custom Niri config, themed terminal, and productivity scripts.",
		},
		{
			Title: "ArchUp Bleu Theme",
			Description: "Modern Bleu-themed configuration with clean aesthetics.\n" +
				"Minimalist approach with carefully selected color palette.",
		},
		{
			Title: "Stock Configuration",
			Description: "Vanilla Arch Linux configuration with minimal customization.\n" +
				"Start with a clean slate and configure everything yourself.",
		},
	}

	return themeConfigModel{
		list: components.NewListModel(items),
	}
}

func (m themeConfigModel) Update(msg tea.Msg) (themeConfigModel, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter", " ":
			m.selected = m.list.SelectedItem().Title
			m.confirmed = true
			return m, nil
		}
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m themeConfigModel) View() string {
	title := components.Header(
		"Theme & Configuration",
		"Choose your preferred theme and configuration approach",
		80,
	)

	description := theme.HintStyle.Render(
		"This choice determines:\n" +
			"  • Desktop theme and color scheme\n" +
			"  • Default dotfiles and configurations\n" +
			"  • Pre-installed scripts and utilities\n",
	)

	listView := m.list.View()

	hint := theme.HintStyle.Render("Press Enter to select")

	content := lipgloss.JoinVertical(
		lipgloss.Left,
		title,
		"",
		description,
		"",
		listView,
		"",
		hint,
	)

	return content
}
