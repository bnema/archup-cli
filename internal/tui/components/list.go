package components

import (
	"fmt"
	"io"

	"github.com/bnema/archup-cli/internal/theme"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// ListItem represents an item in a simple list
type ListItem struct {
	Title       string
	Description string
}

// ListModel is a simple list component for selection
type ListModel struct {
	items         []ListItem
	selectedIndex int
}

// NewListModel creates a new list model
func NewListModel(items []ListItem) ListModel {
	return ListModel{
		items:         items,
		selectedIndex: 0,
	}
}

// Update handles list navigation
func (m ListModel) Update(msg tea.Msg) (ListModel, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "up", "k":
			if m.selectedIndex > 0 {
				m.selectedIndex--
			}
		case "down", "j":
			if m.selectedIndex < len(m.items)-1 {
				m.selectedIndex++
			}
		}
	}
	return m, nil
}

// View renders the list
func (m ListModel) View() string {
	var lines []string
	for i, item := range m.items {
		if i == m.selectedIndex {
			// Selected item
			title := lipgloss.NewStyle().
				Foreground(lipgloss.Color(theme.BrightCyan)).
				Bold(true).
				Render("▶ " + item.Title)
			desc := lipgloss.NewStyle().
				Foreground(lipgloss.Color(theme.PrimaryText)).
				Render("  " + item.Description)
			lines = append(lines, title)
			lines = append(lines, desc)
		} else {
			// Unselected item
			title := lipgloss.NewStyle().
				Foreground(lipgloss.Color(theme.PrimaryText)).
				Render("  " + item.Title)
			desc := lipgloss.NewStyle().
				Foreground(lipgloss.Color(theme.DimmedText)).
				Render("  " + item.Description)
			lines = append(lines, title)
			lines = append(lines, desc)
		}
		// Add spacing between items
		if i < len(m.items)-1 {
			lines = append(lines, "")
		}
	}
	return lipgloss.JoinVertical(lipgloss.Left, lines...)
}

// SelectedItem returns the currently selected item
func (m ListModel) SelectedItem() ListItem {
	if m.selectedIndex >= 0 && m.selectedIndex < len(m.items) {
		return m.items[m.selectedIndex]
	}
	return ListItem{}
}

// SelectIndex sets the selected index
func (m *ListModel) SelectIndex(index int) {
	if index >= 0 && index < len(m.items) {
		m.selectedIndex = index
	}
}

// BleuDelegate is a custom list delegate styled with Bleu theme
type BleuDelegate struct {
	height  int
	spacing int
}

// NewBleuDelegate creates a new Bleu-themed list delegate
func NewBleuDelegate() BleuDelegate {
	return BleuDelegate{
		height:  2,
		spacing: 1,
	}
}

// Height returns the height of list items
func (d BleuDelegate) Height() int { return d.height }

// Spacing returns the spacing between items
func (d BleuDelegate) Spacing() int { return d.spacing }

// Update handles delegate updates
func (d BleuDelegate) Update(_ tea.Msg, _ *list.Model) tea.Cmd { return nil }

// Render renders a list item with Bleu styling
func (d BleuDelegate) Render(w io.Writer, m list.Model, index int, item list.Item) {
	// Type assert to get Title and Description
	type listItem interface {
		Title() string
		Description() string
	}

	i, ok := item.(listItem)
	if !ok {
		return
	}

	var str string
	if index == m.Index() {
		// Selected item - bright cyan with indicator
		str = lipgloss.NewStyle().
			Foreground(lipgloss.Color(theme.BrightCyan)).
			Bold(true).
			Render("▶ " + i.Title())
		str += "\n  " + lipgloss.NewStyle().
			Foreground(lipgloss.Color(theme.PrimaryText)).
			Render(i.Description())
	} else {
		// Unselected item - normal text with dimmed description
		str = lipgloss.NewStyle().
			Foreground(lipgloss.Color(theme.PrimaryText)).
			Render("  " + i.Title())
		str += "\n  " + lipgloss.NewStyle().
			Foreground(lipgloss.Color(theme.DimmedText)).
			Render(i.Description())
	}

	_, _ = fmt.Fprint(w, str)
}
