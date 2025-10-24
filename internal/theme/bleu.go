package theme

import "github.com/charmbracelet/lipgloss"

// Bleu Design System Color Palette
// Based on /home/brice/projects/bleu-theme/docs/bleu-design-system.md

// Background Colors
const (
	DeepNavy      = "#050a14" // Primary background
	DarkBlue      = "#070c16" // Borders, dividers
	DarkerBlue    = "#0a1018" // Panels, status bar
	ActiveTabBlue = "#0f1520" // Active tabs, highlighted sections
	CardBlue      = "#1a2332" // Hover states, card backgrounds
	OceanBlue     = "#2d4a6b" // Selections, hover states
)

// Text Colors
const (
	PrimaryText = "#e8f4f8" // Main body text (WCAG AAA - 14:1)
	PureWhite   = "#fefefe" // Emphasis, selected items (WCAG AAA - 16:1)
	DimmedText  = "#708090" // Comments, inactive elements (WCAG AA - 5:1)
)

// Accent Colors
const (
	BrightCyan    = "#00d4ff" // Keywords, links, primary actions
	PureBlue      = "#5588cc" // Types, focus states, secondary actions
	LightSkyBlue  = "#87ceeb" // Strings, info states
	SkyBlue       = "#4a7ba7" // Numbers, constants
)

// Status Colors
const (
	SuccessGreen = "#99FFE4" // Success states, confirmations
	SoftRed      = "#ff6b8a" // Errors, destructive actions
	WarmOrange   = "#ffb347" // Warnings, attention needed
)

// Headers
var (
	HeaderStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color(PureWhite)).
			Bold(true).
			Padding(0, 1)

	SubHeaderStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color(PrimaryText)).
			Padding(0, 1)
)

// Interactive Elements - Buttons
var (
	ButtonPrimaryStyle = lipgloss.NewStyle().
				Background(lipgloss.Color(PureBlue)).
				Foreground(lipgloss.Color(PureWhite)).
				Padding(0, 2).
				Bold(true)

	ButtonSecondaryStyle = lipgloss.NewStyle().
				Background(lipgloss.Color(OceanBlue)).
				Foreground(lipgloss.Color(PrimaryText)).
				Padding(0, 2)

	ButtonDisabledStyle = lipgloss.NewStyle().
				Background(lipgloss.Color(DarkBlue)).
				Foreground(lipgloss.Color(DimmedText)).
				Padding(0, 2)
)

// Interactive Elements - Lists
var (
	ListItemStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color(PrimaryText)).
			Padding(0, 2)

	ListItemSelectedStyle = lipgloss.NewStyle().
				Background(lipgloss.Color(PureBlue)).
				Foreground(lipgloss.Color(PureWhite)).
				Padding(0, 2).
				Bold(true)

	ListItemFocusedStyle = lipgloss.NewStyle().
				Background(lipgloss.Color(OceanBlue)).
				Foreground(lipgloss.Color(PrimaryText)).
				Padding(0, 2)
)

// Containers
var (
	PanelStyle = lipgloss.NewStyle().
			Background(lipgloss.Color(DarkerBlue)).
			Foreground(lipgloss.Color(PrimaryText)).
			Padding(1, 2).
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color(DarkBlue))

	CardStyle = lipgloss.NewStyle().
			Background(lipgloss.Color(CardBlue)).
			Foreground(lipgloss.Color(PrimaryText)).
			Padding(1, 2).
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color(DarkBlue))

	BorderStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color(DarkBlue))

	BorderActiveStyle = lipgloss.NewStyle().
				Border(lipgloss.RoundedBorder()).
				BorderForeground(lipgloss.Color(PureBlue))
)

// Status Indicators
var (
	SpinnerStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color(BrightCyan))

	ProgressBarStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color(PureBlue))

	SuccessStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color(SuccessGreen)).
			Bold(true)

	ErrorStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color(SoftRed)).
			Bold(true)

	WarningStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color(WarmOrange)).
			Bold(true)
)

// Text Hierarchy
var (
	TitleStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color(PureWhite)).
			Bold(true).
			Padding(1, 0)

	BodyStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color(PrimaryText))

	HintStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color(DimmedText)).
			Italic(true)

	KeybindingStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color(BrightCyan))
)

// Focused returns a focused version of a style with bright cyan accents
func Focused(s lipgloss.Style) lipgloss.Style {
	return s.BorderForeground(lipgloss.Color(BrightCyan))
}

// Dimmed returns a dimmed version of any style
func Dimmed(s lipgloss.Style) lipgloss.Style {
	return s.Foreground(lipgloss.Color(DimmedText))
}
