package components

import (
	"strings"

	"github.com/bnema/archup-cli/internal/theme"
	"github.com/charmbracelet/lipgloss"
)

// Keybinding represents a single keybinding with its description
type Keybinding struct {
	Key  string
	Desc string
}

// Footer renders a footer with keybindings
func Footer(keybindings []Keybinding, width int) string {
	var parts []string

	for _, kb := range keybindings {
		key := theme.KeybindingStyle.Render(kb.Key)
		desc := theme.HintStyle.Render(kb.Desc)
		parts = append(parts, key+" "+desc)
	}

	content := strings.Join(parts, " • ")

	footerStyle := lipgloss.NewStyle().
		Width(width).
		Padding(1, 2).
		Background(lipgloss.Color(theme.DarkerBlue)).
		Foreground(lipgloss.Color(theme.PrimaryText))

	return footerStyle.Render(content)
}

// StandardFooter returns common keybindings for navigation
func StandardFooter(width int, includeBack bool) string {
	keybindings := []Keybinding{
		{"↑/k", "up"},
		{"↓/j", "down"},
		{"enter", "select"},
	}

	if includeBack {
		keybindings = append(keybindings, Keybinding{"esc", "back"})
	}

	keybindings = append(keybindings, Keybinding{"q", "quit"})

	return Footer(keybindings, width)
}
