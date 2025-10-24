package components

import (
	"github.com/bnema/archup-cli/internal/theme"
	"github.com/charmbracelet/lipgloss"
)

// Tab represents a single tab with a name
type Tab struct {
	Name   string
	Active bool
}

// RenderTabs renders a horizontal tab bar
func RenderTabs(tabs []Tab, width int) string {
	var renderedTabs []string

	for _, tab := range tabs {
		var style lipgloss.Style
		if tab.Active {
			// Active tab - bright cyan with bottom border
			style = lipgloss.NewStyle().
				Foreground(lipgloss.Color(theme.BrightCyan)).
				Bold(true).
				Padding(0, 2).
				BorderStyle(lipgloss.Border{
					Top:    "",
					Bottom: "▀",
					Left:   "",
					Right:  "",
				}).
				BorderForeground(lipgloss.Color(theme.BrightCyan)).
				BorderBottom(true)
		} else {
			// Inactive tab - dimmed text
			style = lipgloss.NewStyle().
				Foreground(lipgloss.Color(theme.DimmedText)).
				Padding(0, 2)
		}
		renderedTabs = append(renderedTabs, style.Render(tab.Name))
	}

	// Join tabs horizontally
	tabBar := lipgloss.JoinHorizontal(lipgloss.Top, renderedTabs...)

	// Add bottom border line
	borderLine := lipgloss.NewStyle().
		Foreground(lipgloss.Color(theme.DarkBlue)).
		Render(lipgloss.NewStyle().Width(width).Render("━"))

	return lipgloss.JoinVertical(lipgloss.Left, tabBar, borderLine)
}
