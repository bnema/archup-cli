package tabbed

import (
	"time"

	"github.com/bnema/archup-cli/internal/theme"
	"github.com/bnema/archup-cli/internal/tui/components"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type systemModel struct {
	spinner  spinner.Model
	hardware HardwareInfo
	complete bool
	detected bool
}

type hardwareDetectedMsg HardwareInfo

func newSystemModel() systemModel {
	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color(theme.BrightCyan))
	return systemModel{
		spinner: s,
	}
}

func (m systemModel) Init() tea.Cmd {
	return tea.Batch(
		m.spinner.Tick,
		detectHardware(),
	)
}

func detectHardware() tea.Cmd {
	return tea.Tick(2*time.Second, func(t time.Time) tea.Msg {
		return hardwareDetectedMsg{
			GPU:   "AMD Radeon RX 7900 XTX",
			CPU:   "AMD Ryzen 9 7950X",
			Audio: "Focusrite Scarlett 2i2",
		}
	})
}

func (m systemModel) Update(msg tea.Msg) (systemModel, tea.Cmd) {
	switch msg := msg.(type) {
	case hardwareDetectedMsg:
		m.hardware = HardwareInfo(msg)
		m.detected = true
		m.complete = true
		return m, nil

	case spinner.TickMsg:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd
	}

	return m, nil
}

func (m systemModel) View() string {
	title := components.Header("System Configuration", "Base system packages and hardware detection", 80)

	var status string
	if !m.detected {
		status = lipgloss.NewStyle().
			Foreground(lipgloss.Color(theme.BrightCyan)).
			Render(m.spinner.View() + " Detecting hardware...")
	} else {
		// Hardware detected - show results
		status = lipgloss.NewStyle().
			Foreground(lipgloss.Color(theme.SuccessGreen)).
			Bold(true).
			Render(components.IconSuccess + " Hardware detected")

		hardwareInfo := lipgloss.NewStyle().
			Foreground(lipgloss.Color(theme.PrimaryText)).
			MarginTop(2).
			Render(lipgloss.JoinVertical(
				lipgloss.Left,
				theme.BodyStyle.Render("GPU:   "+m.hardware.GPU),
				theme.BodyStyle.Render("CPU:   "+m.hardware.CPU),
				theme.BodyStyle.Render("Audio: "+m.hardware.Audio),
			))

		basePackages := lipgloss.NewStyle().
			Foreground(lipgloss.Color(theme.PrimaryText)).
			MarginTop(2).
			Render(lipgloss.JoinVertical(
				lipgloss.Left,
				theme.BodyStyle.Bold(true).Render("Base System Packages:"),
				theme.BodyStyle.Render("  • linux, linux-firmware"),
				theme.BodyStyle.Render("  • base-devel, git"),
				theme.BodyStyle.Render("  • networkmanager, pipewire"),
				theme.BodyStyle.Render("  • wayland, xdg-desktop-portal"),
			))

		status = lipgloss.JoinVertical(
			lipgloss.Left,
			status,
			hardwareInfo,
			basePackages,
		)
	}

	hint := theme.HintStyle.Render("Use tab/← → to switch between tabs")

	return lipgloss.JoinVertical(
		lipgloss.Left,
		title,
		"",
		status,
		"",
		hint,
	)
}
