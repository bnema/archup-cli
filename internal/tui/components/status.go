package components

import (
	"fmt"
	"strings"

	"github.com/bnema/archup-cli/internal/theme"
	"github.com/charmbracelet/lipgloss"
)

// Status icons
const (
	IconInfo    = "ℹ"
	IconSuccess = "✓"
	IconError   = "✗"
	IconWarning = "!"
)

// StatusMessage represents different types of status messages
type StatusMessage struct {
	Type    StatusType
	Message string
}

type StatusType int

const (
	StatusInfo StatusType = iota
	StatusSuccess
	StatusError
	StatusWarning
)

// RenderStatus renders a status message with appropriate styling
func RenderStatus(msg StatusMessage) string {
	var icon string
	var style lipgloss.Style

	switch msg.Type {
	case StatusSuccess:
		icon = IconSuccess
		style = theme.SuccessStyle
	case StatusError:
		icon = IconError
		style = theme.ErrorStyle
	case StatusWarning:
		icon = IconWarning
		style = theme.WarningStyle
	default: // StatusInfo
		icon = IconInfo
		style = theme.BodyStyle
	}

	return style.Render(icon + " " + msg.Message)
}

// ProgressIndicator renders a simple progress indicator
func ProgressIndicator(current, total int, width int) string {
	if total == 0 {
		return ""
	}

	percentage := float64(current) / float64(total)
	barWidth := width - 20 // Leave space for text
	filled := int(percentage * float64(barWidth))

	bar := strings.Repeat("━", filled) + strings.Repeat("─", barWidth-filled)
	text := fmt.Sprintf("Step %d/%d", current, total)

	barStyle := theme.ProgressBarStyle
	textStyle := theme.BodyStyle

	return lipgloss.JoinHorizontal(
		lipgloss.Left,
		textStyle.Render(text+"  "),
		barStyle.Render(bar),
		textStyle.Render(fmt.Sprintf("  %.0f%%", percentage*100)),
	)
}
