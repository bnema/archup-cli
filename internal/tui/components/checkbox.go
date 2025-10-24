package components

import (
	"github.com/bnema/archup-cli/internal/theme"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// CheckboxItem represents a single checkbox item
type CheckboxItem struct {
	Name        string
	Description string
	Selected    bool
}

// CheckboxList is a reusable multi-select checkbox list component
type CheckboxList struct {
	Items    []CheckboxItem
	Cursor   int
	width    int
	height   int
	Finished bool
}

// NewCheckboxList creates a new checkbox list
func NewCheckboxList(items []CheckboxItem) CheckboxList {
	return CheckboxList{
		Items:  items,
		Cursor: 0,
		width:  80,
		height: 20,
	}
}

// Update handles checkbox list updates
func (m CheckboxList) Update(msg tea.Msg) (CheckboxList, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		return m, nil

	case tea.KeyMsg:
		switch msg.String() {
		case "up", "k":
			if m.Cursor > 0 {
				m.Cursor--
			}

		case "down", "j":
			if m.Cursor < len(m.Items)-1 {
				m.Cursor++
			}

		case " ":
			// Toggle selection
			if m.Cursor < len(m.Items) {
				m.Items[m.Cursor].Selected = !m.Items[m.Cursor].Selected
			}

		case "enter":
			m.Finished = true
			return m, nil
		}
	}

	return m, nil
}

// View renders the checkbox list
func (m CheckboxList) View() string {
	var items string

	for i, item := range m.Items {
		cursor := " "
		if i == m.Cursor {
			cursor = theme.KeybindingStyle.Render(">")
		}

		checkbox := "[ ]"
		if item.Selected {
			checkbox = "[" + IconSuccess + "]"
		}

		nameStyle := theme.BodyStyle
		descStyle := theme.HintStyle

		if i == m.Cursor {
			nameStyle = theme.ListItemFocusedStyle
		}

		line := lipgloss.JoinHorizontal(
			lipgloss.Left,
			cursor+" "+checkbox+"  ",
			nameStyle.Render(item.Name)+"  ",
			descStyle.Render("- "+item.Description),
		)

		items += line + "\n"
	}

	return items
}

// GetSelected returns all selected item names
func (m CheckboxList) GetSelected() []string {
	var selected []string
	for _, item := range m.Items {
		if item.Selected {
			selected = append(selected, item.Name)
		}
	}
	return selected
}
