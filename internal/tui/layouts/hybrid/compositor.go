package hybrid

import (
	"github.com/bnema/archup-cli/internal/theme"
	"github.com/bnema/archup-cli/internal/tui/components"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type compositorModel struct {
	list      components.ListModel
	selected  string
	confirmed bool
}

func newCompositorModel() compositorModel {
	items := []components.ListItem{
		{
			Title:       "Niri",
			Description: "Scrollable-tiling Wayland compositor with smooth animations",
		},
		{
			Title:       "Hyprland",
			Description: "Dynamic tiling Wayland compositor with eye candy",
		},
		{
			Title:       "Sway",
			Description: "i3-compatible Wayland compositor (stable and minimal)",
		},
		{
			Title:       "River",
			Description: "Dynamic tiling Wayland compositor written in Zig",
		},
	}

	return compositorModel{
		list: components.NewListModel(items),
	}
}

func (m compositorModel) Update(msg tea.Msg) (compositorModel, tea.Cmd) {
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

func (m compositorModel) View() string {
	title := components.Header("Compositor Selection", "Choose your Wayland compositor", 80)

	listView := m.list.View()

	hint := theme.HintStyle.Render("Use arrow keys to navigate, Enter to select")

	return lipgloss.JoinVertical(
		lipgloss.Left,
		title,
		"",
		listView,
		"",
		hint,
	)
}
