package components

import (
	"github.com/bnema/archup-cli/internal/theme"
	"github.com/charmbracelet/lipgloss"
)

const archupLogo = `  ░█████╗░██████╗░░█████╗░██╗░░██╗██╗░░░██╗██████╗░
  ██╔══██╗██╔══██╗██╔══██╗██║░░██║██║░░░██║██╔══██╗
  ███████║██████╔╝██║░░╚═╝███████║██║░░░██║██████╔╝
  ██╔══██║██╔══██╗██║░░██╗██╔══██║██║░░░██║██╔═══╝░
  ██║░░██║██║░░██║╚█████╔╝██║░░██║╚██████╔╝██║░░░░░
  ╚═╝░░╚═╝╚═╝░░╚═╝░╚════╝░╚═╝░░╚═╝░╚═════╝░╚═╝░░░░░

  v0.3.0`

// Header renders a standard header with title and optional subtitle
func Header(title, subtitle string, width int) string {
	titleStyle := theme.TitleStyle.Width(width)
	title = titleStyle.Render(title)

	if subtitle == "" {
		return title
	}

	subtitleStyle := theme.HintStyle.Width(width)
	subtitle = subtitleStyle.Render(subtitle)

	return lipgloss.JoinVertical(lipgloss.Left, title, subtitle)
}

// HeaderWithLogo renders a header with the ArchUp ASCII logo
func HeaderWithLogo(width int) string {
	logoStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color(theme.BrightCyan)).
		Width(width).
		Align(lipgloss.Center)

	subtitle := "Desktop Environment Setup Wizard"
	subtitleStyle := theme.HintStyle.Width(width).Align(lipgloss.Center)

	return lipgloss.JoinVertical(
		lipgloss.Center,
		logoStyle.Render(archupLogo),
		subtitleStyle.Render(subtitle),
	)
}
