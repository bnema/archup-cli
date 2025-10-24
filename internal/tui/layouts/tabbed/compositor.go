package tabbed

import (
	"github.com/bnema/archup-cli/internal/theme"
	"github.com/bnema/archup-cli/internal/tui/components"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type compositorItem struct {
	name        string
	description string
}

func (i compositorItem) Title() string       { return i.name }
func (i compositorItem) Description() string { return i.description }
func (i compositorItem) FilterValue() string { return i.name }

type compositorModel struct {
	list     list.Model
	selected string
}

func newCompositorModel() compositorModel {
	items := []list.Item{
		compositorItem{
			name:        "Niri",
			description: "Scrollable-tiling Wayland compositor with smooth animations",
		},
		compositorItem{
			name:        "Hyprland",
			description: "Dynamic tiling Wayland compositor with eye candy",
		},
		compositorItem{
			name:        "Sway",
			description: "i3-compatible Wayland compositor (stable and minimal)",
		},
		compositorItem{
			name:        "River",
			description: "Dynamic tiling Wayland compositor written in Zig",
		},
	}

	delegate := components.NewBleuDelegate()
	l := list.New(items, delegate, 80, 20)
	l.Title = "Select Compositor"
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(false)
	l.Styles.Title = theme.HeaderStyle

	return compositorModel{
		list: l,
	}
}

func (m compositorModel) Update(msg tea.Msg) (compositorModel, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.list.SetWidth(msg.Width - 4)
		m.list.SetHeight(15)
		return m, nil

	case tea.KeyMsg:
		switch msg.String() {
		case "enter", " ":
			if selected, ok := m.list.SelectedItem().(compositorItem); ok {
				m.selected = selected.name
				return m, nil
			}
		}
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m compositorModel) View() string {
	title := components.Header("Compositor Selection", "Choose your Wayland compositor", 80)

	hint := theme.HintStyle.Render("Use arrow keys to navigate, Enter to select, tab/← → to switch tabs")

	selectedInfo := ""
	if m.selected != "" {
		selectedInfo = lipgloss.NewStyle().
			Foreground(lipgloss.Color(theme.SuccessGreen)).
			MarginTop(1).
			Render(components.IconSuccess + " Selected: " + m.selected)
	}

	return lipgloss.JoinVertical(
		lipgloss.Left,
		title,
		"",
		m.list.View(),
		selectedInfo,
		"",
		hint,
	)
}
