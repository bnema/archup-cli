package hybrid

import (
	"fmt"
	"time"

	"github.com/bnema/archup-cli/internal/theme"
	"github.com/bnema/archup-cli/internal/tui/components"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type hardwareModel struct {
	spinner  spinner.Model
	hardware HardwareInfo
	complete bool
	detected bool
}

type hardwareDetectedMsg HardwareInfo

func newHardwareModel() hardwareModel {
	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = theme.SpinnerStyle

	return hardwareModel{
		spinner: s,
	}
}

func (m hardwareModel) Init() tea.Cmd {
	return tea.Batch(
		m.spinner.Tick,
		detectHardware(),
	)
}

func (m hardwareModel) Update(msg tea.Msg) (hardwareModel, tea.Cmd) {
	switch msg := msg.(type) {
	case hardwareDetectedMsg:
		m.hardware = HardwareInfo(msg)
		m.detected = true
		return m, nil

	case tea.KeyMsg:
		if m.detected {
			switch msg.String() {
			case "enter", " ":
				m.complete = true
				return m, nil
			}
		}

	case spinner.TickMsg:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd
	}

	return m, nil
}

func (m hardwareModel) View() string {
	title := components.Header("Hardware Detection", "Detecting your system hardware...", 80)

	if !m.detected {
		return lipgloss.JoinVertical(
			lipgloss.Left,
			title,
			"",
			m.spinner.View()+" Scanning hardware...",
		)
	}

	// Show detected hardware
	gpuLine := theme.BodyStyle.Render(components.IconSuccess + " GPU: " + m.hardware.GPU)
	cpuLine := theme.BodyStyle.Render(components.IconSuccess + " CPU: " + m.hardware.CPU)
	audioLine := theme.BodyStyle.Render(components.IconSuccess + " Audio: " + m.hardware.Audio)

	// Show laptop detection
	laptopStatus := "Desktop"
	if m.hardware.IsLaptop {
		laptopStatus = "Laptop (TLP will be configured)"
	}
	laptopLine := theme.BodyStyle.Render(components.IconSuccess + " Type: " + laptopStatus)

	// Show CPU frequencies
	var freqLine string
	if m.hardware.MinCPUFreq > 0 && m.hardware.MaxCPUFreq > 0 {
		freqLine = theme.BodyStyle.Render(
			fmt.Sprintf(components.IconSuccess+" CPU Frequency: %d MHz - %d MHz",
				m.hardware.MinCPUFreq, m.hardware.MaxCPUFreq),
		)
	}

	content := lipgloss.JoinVertical(
		lipgloss.Left,
		title,
		"",
		"Detected hardware:",
		"",
		gpuLine,
		cpuLine,
		audioLine,
		laptopLine,
		freqLine,
		"",
		"",
		theme.ButtonPrimaryStyle.Render(" Press Enter to Continue "),
	)

	return content
}

// detectHardware simulates hardware detection
func detectHardware() tea.Cmd {
	return func() tea.Msg {
		// Simulate detection delay
		time.Sleep(2 * time.Second)

		// TODO: Replace with actual hardware detection
		return hardwareDetectedMsg{
			GPU:        "AMD Radeon RX 6700 XT",
			CPU:        "AMD Ryzen 7 5800X",
			CPUModel:   "AMD Ryzen 7 5800X 8-Core Processor",
			Audio:      "PipeWire",
			IsLaptop:   false,
			MinCPUFreq: 2200,
			MaxCPUFreq: 4700,
		}
	}
}
