package tree

import (
	"github.com/bnema/archup-cli/internal/theme"
	"github.com/bnema/archup-cli/internal/tui/components"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// Model represents the tree/checklist demo state
type Model struct {
	tree   *TreeView
	width  int
	height int

	confirmed         bool
	installing        bool
	installationModel installationModel

	quitting bool
}

// installationModel wraps the installation component
type installationModel struct {
	components.InstallationModel
}

// New creates a new tree demo model
func New() Model {
	// Build the package tree structure
	root := buildPackageTree()

	// Create installation tasks
	tasks := []string{
		"Updating package database",
		"Installing base system packages",
		"Installing desktop environment",
		"Installing applications",
		"Configuring system",
		"Setting up user environment",
		"Cleaning up package cache",
	}

	return Model{
		tree:              NewTreeView(root),
		width:             120,
		height:            30,
		installationModel: installationModel{components.NewInstallationModel(tasks)},
	}
}

// Init initializes the installation model
func (m installationModel) Init() tea.Cmd {
	return m.InstallationModel.Init()
}

// Update wraps the installation model update
func (m installationModel) Update(msg tea.Msg) (installationModel, tea.Cmd) {
	var cmd tea.Cmd
	m.InstallationModel, cmd = m.InstallationModel.Update(msg)
	return m, cmd
}

// View wraps the installation model view
func (m installationModel) View() string {
	return m.InstallationModel.View()
}

// buildPackageTree creates the demo package tree
func buildPackageTree() *Node {
	root := NewNode("Package Selection", "Select packages to install")

	// Base System
	baseSystem := NewNode("Base System", "Core system packages")
	baseSystem.AddChild(NewNode("linux", "Linux kernel"))
	baseSystem.AddChild(NewNode("linux-firmware", "Hardware firmware files"))
	baseSystem.AddChild(NewNode("base-devel", "Development tools"))
	baseSystem.AddChild(NewNode("linux-lts", "Long-term support kernel"))
	baseSystem.State = Checked
	baseSystem.CheckAllChildren(true) // Start with base system selected
	root.AddChild(baseSystem)

	// Desktop Environment (Radio group)
	desktopEnv := NewNode("Desktop Environment", "Choose your compositor")
	desktopEnv.IsRadio = true
	desktopEnv.AddChild(NewNode("Niri", "Scrollable-tiling Wayland compositor"))
	desktopEnv.AddChild(NewNode("Hyprland", "Dynamic tiling Wayland compositor"))
	desktopEnv.AddChild(NewNode("Sway", "i3-compatible Wayland compositor"))
	desktopEnv.AddChild(NewNode("River", "Dynamic tiling Wayland compositor"))
	// Select first one by default
	desktopEnv.Children[0].State = Checked
	desktopEnv.UpdateParentState()
	root.AddChild(desktopEnv)

	// Applications
	applications := NewNode("Applications", "Optional applications")

	// Terminal Emulators
	terminals := NewNode("Terminal Emulators", "Choose terminal apps")
	terminals.AddChild(NewNode("Foot", "Fast Wayland terminal"))
	terminals.AddChild(NewNode("Alacritty", "GPU-accelerated terminal"))
	terminals.AddChild(NewNode("Kitty", "Feature-rich terminal"))
	// Select first one by default
	terminals.Children[0].State = Checked
	terminals.UpdateParentState()
	applications.AddChild(terminals)

	// Browsers
	browsers := NewNode("Web Browsers", "Internet browsers")
	browsers.AddChild(NewNode("Firefox", "Mozilla Firefox browser"))
	browsers.AddChild(NewNode("Chromium", "Open-source Chrome"))
	browsers.AddChild(NewNode("Brave", "Privacy-focused browser"))
	// Select Firefox by default
	browsers.Children[0].State = Checked
	browsers.UpdateParentState()
	applications.AddChild(browsers)

	// File Managers
	fileManagers := NewNode("File Managers", "GUI file browsers")
	fileManagers.AddChild(NewNode("Thunar", "Lightweight file manager"))
	fileManagers.AddChild(NewNode("Nautilus", "GNOME file manager"))
	fileManagers.AddChild(NewNode("Dolphin", "KDE file manager"))
	applications.AddChild(fileManagers)

	// Media Players
	media := NewNode("Media Players", "Audio and video players")
	media.AddChild(NewNode("MPV", "Minimalist video player"))
	media.AddChild(NewNode("VLC", "Feature-rich media player"))
	media.AddChild(NewNode("imv", "Image viewer"))
	applications.AddChild(media)

	// Text Editors
	editors := NewNode("Text Editors", "Code and text editors")
	editors.AddChild(NewNode("Neovim", "Vim-based text editor"))
	editors.AddChild(NewNode("Helix", "Modern modal editor"))
	editors.AddChild(NewNode("VS Code", "Microsoft's code editor"))
	applications.AddChild(editors)

	applications.UpdateParentState()
	root.AddChild(applications)

	// Update root state based on children
	root.UpdateParentState()

	return root
}

// Init initializes the model
func (m Model) Init() tea.Cmd {
	return nil
}

// Update handles messages and updates the model
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// Handle installation updates first
	if m.installing {
		var cmd tea.Cmd
		m.installationModel, cmd = m.installationModel.Update(msg)

		// Check if installation is complete and user wants to close
		if m.installationModel.ShouldClose || m.installationModel.ShouldReboot {
			m.quitting = true
			return m, tea.Quit
		}

		return m, cmd
	}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			m.quitting = true
			return m, tea.Quit

		case "esc":
			// Go back from confirmation view
			if m.confirmed {
				m.confirmed = false
				return m, nil
			}
			return m, nil

		case "up", "k":
			m.tree.MoveCursor(-1)
			return m, nil

		case "down", "j":
			m.tree.MoveCursor(1)
			return m, nil

		case "left", "h":
			// Collapse current node
			if node := m.tree.GetCurrentNode(); node != nil {
				if node.Expanded && len(node.Children) > 0 {
					node.Expanded = false
				} else if node.Parent != nil {
					// Move to parent
					nodes := m.tree.GetVisibleNodes()
					for i, n := range nodes {
						if n == node.Parent {
							m.tree.cursor = i
							break
						}
					}
				}
			}
			return m, nil

		case "right", "l":
			// Expand current node
			if node := m.tree.GetCurrentNode(); node != nil && len(node.Children) > 0 {
				node.Expanded = true
			}
			return m, nil

		case " ":
			// Toggle checkbox
			m.tree.ToggleCurrentCheck()
			return m, nil

		case "enter":
			// On tree view: go to confirmation
			if !m.confirmed {
				m.confirmed = true
				return m, nil
			}
			// On confirmation view: start installation
			if m.confirmed && !m.installing {
				m.installing = true
				return m, m.installationModel.Init()
			}
			return m, nil

		case "e":
			// Toggle expand/collapse
			m.tree.ToggleCurrentExpanded()
			return m, nil
		}

	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		return m, nil
	}

	return m, nil
}

// View renders the tree demo
func (m Model) View() string {
	if m.quitting {
		return ""
	}

	// Show installation view if installing
	if m.installing {
		return m.installationModel.View()
	}

	if m.confirmed {
		return m.viewConfirmation()
	}

	// Calculate content height (total - header - footer - padding)
	headerHeight := 5
	footerHeight := 3
	contentHeight := m.height - headerHeight - footerHeight - 1

	// Header
	header := components.Header(
		"Package Selection",
		"Navigate with ↑↓/jk, expand/collapse with →←/hl or e, select with space",
		m.width,
	)

	// Tree content
	treeContent := m.tree.Render(m.width-4, contentHeight)

	// Create fixed-height container for content
	contentStyle := lipgloss.NewStyle().
		Height(contentHeight).
		Width(m.width).
		Padding(0, 2)

	content := contentStyle.Render(treeContent)

	// Footer with keybindings
	footer := components.Footer([]components.Keybinding{
		{Key: "↑↓/jk", Desc: "move"},
		{Key: "→←/hl", Desc: "expand"},
		{Key: "space", Desc: "toggle"},
		{Key: "enter", Desc: "next"},
		{Key: "q", Desc: "quit"},
	}, m.width)

	return lipgloss.JoinVertical(lipgloss.Left, header, "", content, footer)
}

// viewConfirmation renders the confirmation screen
func (m Model) viewConfirmation() string {
	// Calculate content height (total - header - footer - padding)
	headerHeight := 5
	footerHeight := 3
	contentHeight := m.height - headerHeight - footerHeight - 2

	header := components.Header(
		"Review Selection",
		"Review your selections and press Enter to install",
		m.width,
	)

	selected := m.tree.GetRoot().GetSelectedItems()

	var itemsSection string
	if len(selected) == 0 {
		itemsSection = components.RenderStatus(components.StatusMessage{
			Type:    components.StatusWarning,
			Message: "No packages selected",
		})
	} else {
		var itemsList []string
		itemsList = append(itemsList, lipgloss.NewStyle().
			Foreground(lipgloss.Color(theme.BrightCyan)).
			Bold(true).
			Render("Selected Packages:"))
		itemsList = append(itemsList, "")

		for _, item := range selected {
			itemsList = append(itemsList, "  "+components.IconSuccess+" "+item)
		}

		itemsSection = lipgloss.JoinVertical(lipgloss.Left, itemsList...)
	}

	// Install button
	installButton := theme.ButtonPrimaryStyle.
		Background(lipgloss.Color(theme.BrightCyan)).
		Foreground(lipgloss.Color(theme.PureWhite)).
		Render("[ Install ]")

	buttonHint := theme.HintStyle.Render("Press Enter to install")

	buttonSection := lipgloss.JoinVertical(
		lipgloss.Left,
		"",
		"",
		installButton,
		"",
		buttonHint,
	)

	// Combine content
	content := lipgloss.JoinVertical(
		lipgloss.Left,
		itemsSection,
		buttonSection,
	)

	// Create fixed-height container for content
	contentStyle := lipgloss.NewStyle().
		Height(contentHeight).
		Width(m.width).
		Padding(0, 2)

	content = contentStyle.Render(content)

	footer := components.Footer([]components.Keybinding{
		{Key: "enter", Desc: "install"},
		{Key: "esc", Desc: "back"},
		{Key: "q", Desc: "quit"},
	}, m.width)

	return lipgloss.JoinVertical(
		lipgloss.Left,
		header,
		"",
		content,
		footer,
	)
}
