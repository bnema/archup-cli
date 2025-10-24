package components

import (
	"fmt"
	"time"

	"github.com/bnema/archup-cli/internal/theme"
	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// InstallationModel handles installation progress display
type InstallationModel struct {
	progress       progress.Model
	percent        float64
	currentTask    string
	Complete       bool
	Tasks          []string
	currentIndex   int
	selectedButton int // 0 = Reboot, 1 = Close
	ShouldReboot   bool
	ShouldClose    bool
	logLines       []string // Live log output
	maxLogLines    int      // Maximum number of log lines to show
}

// ProgressMsg represents installation progress
type ProgressMsg struct {
	Percent float64
	Task    string
}

// LogMsg represents a new log line
type LogMsg struct {
	Line string
}

// NewInstallationModel creates a new installation model with custom tasks
func NewInstallationModel(tasks []string) InstallationModel {
	return InstallationModel{
		progress: progress.New(
			progress.WithDefaultGradient(),
			progress.WithWidth(60),
		),
		Tasks:       tasks,
		logLines:    []string{},
		maxLogLines: 8, // Show last 8 lines of logs
	}
}

// generateDemoLog simulates pacman-style installation output
func generateDemoLog(taskIndex int) tea.Cmd {
	logs := [][]string{
		// Task 0: Updating package database
		{
			":: Synchronizing package databases...",
			"   core is up to date",
			"   extra is up to date",
			"   multilib is up to date",
		},
		// Task 1: Installing base system
		{
			":: Retrieving packages...",
			"   downloading linux-6.11.1...",
			"   downloading linux-firmware-20241015...",
			"   downloading base-devel-1-1...",
			":: Installing packages...",
			"   [1/3] installing linux...",
			"   [2/3] installing linux-firmware...",
			"   [3/3] installing base-devel...",
		},
		// Task 2: Installing desktop environment
		{
			":: Installing compositor...",
			"   downloading niri-0.1.7...",
			"   downloading wayland-1.23.0...",
			"   [1/2] installing wayland...",
			"   [2/2] installing niri...",
		},
		// Task 3: Installing applications
		{
			":: Installing applications...",
			"   downloading firefox-130.0...",
			"   downloading foot-1.17.2...",
			"   [1/5] installing foot...",
			"   [2/5] installing firefox...",
			"   [3/5] installing thunar...",
			"   [4/5] installing mpv...",
			"   [5/5] installing neovim...",
		},
		// Task 4: Configuring system
		{
			":: Configuring system...",
			"   generating locales...",
			"   setting up systemd services...",
			"   configuring network...",
		},
		// Task 5: Setting up user environment
		{
			":: Setting up user environment...",
			"   creating user directories...",
			"   copying dotfiles...",
			"   setting shell to zsh...",
		},
		// Task 6: Cleaning up
		{
			":: Cleaning up...",
			"   removing package cache...",
			"   cleaning build files...",
			"   done!",
		},
	}

	if taskIndex >= len(logs) {
		taskIndex = len(logs) - 1
	}

	// Return a batch of log messages
	var cmds []tea.Cmd
	for _, line := range logs[taskIndex] {
		logLine := line
		cmds = append(cmds, func() tea.Msg {
			return LogMsg{Line: logLine}
		})
	}

	return tea.Batch(cmds...)
}

// Init starts the installation process
func (m InstallationModel) Init() tea.Cmd {
	if len(m.Tasks) == 0 {
		m.Complete = true
		return nil
	}

	return tea.Tick(500*time.Millisecond, func(t time.Time) tea.Msg {
		return ProgressMsg{Percent: 0.0, Task: m.Tasks[0]}
	})
}

// Update handles installation progress updates
func (m InstallationModel) Update(msg tea.Msg) (InstallationModel, tea.Cmd) {
	// Handle button selection when complete
	if m.Complete {
		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch msg.String() {
			case "left", "h":
				m.selectedButton = 0 // Reboot
				return m, nil
			case "right", "l":
				m.selectedButton = 1 // Close
				return m, nil
			case "enter":
				if m.selectedButton == 0 {
					m.ShouldReboot = true
				} else {
					m.ShouldClose = true
				}
				return m, tea.Quit
			}
		}
		return m, nil
	}

	switch msg := msg.(type) {
	case ProgressMsg:
		m.percent = msg.Percent
		m.currentTask = msg.Task
		m.currentIndex = int(msg.Percent * float64(len(m.Tasks)))

		if m.percent >= 1.0 {
			m.Complete = true
			return m, nil
		}

		newPercent := m.percent + (1.0 / float64(len(m.Tasks)))
		taskIndex := int(newPercent * float64(len(m.Tasks)))
		if taskIndex >= len(m.Tasks) {
			taskIndex = len(m.Tasks) - 1
		}

		// Generate demo logs for current task
		logCmd := generateDemoLog(m.currentIndex)

		return m, tea.Batch(
			tea.Tick(800*time.Millisecond, func(t time.Time) tea.Msg {
				return ProgressMsg{
					Percent: newPercent,
					Task:    m.Tasks[taskIndex],
				}
			}),
			logCmd,
		)

	case LogMsg:
		// Add log line and keep only the last N lines
		m.logLines = append(m.logLines, msg.Line)
		if len(m.logLines) > m.maxLogLines {
			m.logLines = m.logLines[len(m.logLines)-m.maxLogLines:]
		}
		return m, nil

	case progress.FrameMsg:
		var cmd tea.Cmd
		progressModel, cmd := m.progress.Update(msg)
		m.progress = progressModel.(progress.Model)
		return m, cmd
	}

	return m, nil
}

// View renders the installation progress
func (m InstallationModel) View() string {
	if m.Complete {
		title := Header("Installation Complete", "Your desktop environment is ready!", 80)

		successMsg := lipgloss.NewStyle().
			Foreground(lipgloss.Color(theme.SuccessGreen)).
			Bold(true).
			MarginTop(2).
			Render(IconSuccess + " Installation completed successfully!")

		// Buttons
		rebootButton := theme.ButtonPrimaryStyle
		closeButton := theme.ButtonPrimaryStyle

		if m.selectedButton == 0 {
			// Reboot selected
			rebootButton = rebootButton.
				Background(lipgloss.Color(theme.BrightCyan)).
				Foreground(lipgloss.Color(theme.PureWhite))
			closeButton = lipgloss.NewStyle().
				Foreground(lipgloss.Color(theme.DimmedText)).
				Padding(0, 2)
		} else {
			// Close selected
			closeButton = closeButton.
				Background(lipgloss.Color(theme.BrightCyan)).
				Foreground(lipgloss.Color(theme.PureWhite))
			rebootButton = lipgloss.NewStyle().
				Foreground(lipgloss.Color(theme.DimmedText)).
				Padding(0, 2)
		}

		buttons := lipgloss.JoinHorizontal(
			lipgloss.Top,
			rebootButton.Render("Reboot"),
			"  ",
			closeButton.Render("Close"),
		)

		hint := theme.HintStyle.Render("Use ←/→ to select, Enter to confirm")

		return lipgloss.JoinVertical(
			lipgloss.Left,
			title,
			"",
			successMsg,
			"",
			"",
			buttons,
			"",
			hint,
		)
	}

	title := Header("Installing", "Please wait while packages are installed", 80)

	// Progress bar
	progressView := m.progress.ViewAs(m.percent)

	// Percentage and current task
	percentage := lipgloss.NewStyle().
		Foreground(lipgloss.Color(theme.BrightCyan)).
		Bold(true).
		Render(fmt.Sprintf("%.0f%%", m.percent*100))

	taskText := lipgloss.NewStyle().
		Foreground(lipgloss.Color(theme.PrimaryText)).
		Render(m.currentTask)

	progressInfo := lipgloss.JoinVertical(
		lipgloss.Left,
		progressView,
		"",
		percentage+" - "+taskText,
	)

	// Task list with checkmarks
	var taskList []string
	for i, task := range m.Tasks {
		if i < m.currentIndex {
			taskList = append(taskList, theme.BodyStyle.Foreground(lipgloss.Color(theme.SuccessGreen)).Render(IconSuccess+" "+task))
		} else if i == m.currentIndex {
			taskList = append(taskList, theme.BodyStyle.Foreground(lipgloss.Color(theme.BrightCyan)).Render("▶ "+task))
		} else {
			taskList = append(taskList, theme.BodyStyle.Foreground(lipgloss.Color(theme.DimmedText)).Render("  "+task))
		}
	}

	tasks := lipgloss.NewStyle().
		MarginTop(2).
		Render(lipgloss.JoinVertical(lipgloss.Left, taskList...))

	// Live log output with faded colors
	var logOutput string
	if len(m.logLines) > 0 {
		logStyle := lipgloss.NewStyle().
			Foreground(lipgloss.Color(theme.DimmedText)).
			Faint(true).
			MarginTop(2)

		logTitle := lipgloss.NewStyle().
			Foreground(lipgloss.Color(theme.DimmedText)).
			Render("─── Live Output ───")

		var fadedLogs []string
		fadedLogs = append(fadedLogs, logTitle)

		for _, line := range m.logLines {
			fadedLogs = append(fadedLogs, logStyle.Render(line))
		}

		logOutput = lipgloss.JoinVertical(lipgloss.Left, fadedLogs...)
	}

	return lipgloss.JoinVertical(
		lipgloss.Left,
		title,
		"",
		progressInfo,
		tasks,
		logOutput,
	)
}
