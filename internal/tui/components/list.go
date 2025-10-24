package components

import (
	"fmt"
	"io"

	"github.com/bnema/archup-cli/internal/theme"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

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
