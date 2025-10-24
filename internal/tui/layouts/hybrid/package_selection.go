package hybrid

import (
	"github.com/bnema/archup-cli/internal/tui/components"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type packageSelectionModel struct {
	tree      components.TreeModel
	confirmed bool
}

func newPackageSelectionModel() packageSelectionModel {
	// Build the package tree structure
	root := buildPackageTree()

	return packageSelectionModel{
		tree: components.NewTreeModel(root),
	}
}

// buildPackageTree creates a demo package tree
// This demonstrates how to build a tree from any data structure
func buildPackageTree() *components.TreeNode {
	root := components.NewTreeNode("Package Selection", "Select packages to install")

	// Base System
	baseSystem := components.NewTreeNode("Base System", "Core system packages")
	baseSystem.AddChild(components.NewTreeNode("linux", "Linux kernel"))
	baseSystem.AddChild(components.NewTreeNode("linux-firmware", "Hardware firmware files"))
	baseSystem.AddChild(components.NewTreeNode("base-devel", "Development tools"))
	baseSystem.AddChild(components.NewTreeNode("linux-lts", "Long-term support kernel (optional)"))
	baseSystem.State = components.Checked
	baseSystem.CheckAllChildren(true) // Start with base system selected
	root.AddChild(baseSystem)

	// Applications
	applications := components.NewTreeNode("Applications", "Optional applications")

	// Terminal Emulators
	terminals := components.NewTreeNode("Terminal Emulators", "Choose terminal apps")
	terminals.AddChild(components.NewTreeNode("Foot", "Fast Wayland terminal"))
	terminals.AddChild(components.NewTreeNode("Alacritty", "GPU-accelerated terminal"))
	terminals.AddChild(components.NewTreeNode("Kitty", "Feature-rich terminal"))
	terminals.AddChild(components.NewTreeNode("WezTerm", "GPU-accelerated terminal multiplexer"))
	// Select Foot by default
	terminals.Children[0].State = components.Checked
	terminals.UpdateParentState()
	applications.AddChild(terminals)

	// Browsers
	browsers := components.NewTreeNode("Web Browsers", "Internet browsers")
	browsers.AddChild(components.NewTreeNode("Firefox", "Mozilla Firefox browser"))
	browsers.AddChild(components.NewTreeNode("Chromium", "Open-source Chrome"))
	browsers.AddChild(components.NewTreeNode("Brave", "Privacy-focused browser"))
	browsers.AddChild(components.NewTreeNode("Librewolf", "Privacy-hardened Firefox fork"))
	// Select Firefox by default
	browsers.Children[0].State = components.Checked
	browsers.UpdateParentState()
	applications.AddChild(browsers)

	// File Managers
	fileManagers := components.NewTreeNode("File Managers", "GUI file browsers")
	fileManagers.AddChild(components.NewTreeNode("Thunar", "Lightweight file manager"))
	fileManagers.AddChild(components.NewTreeNode("Nautilus", "GNOME file manager"))
	fileManagers.AddChild(components.NewTreeNode("Dolphin", "KDE file manager"))
	fileManagers.AddChild(components.NewTreeNode("nnn", "Terminal file manager"))
	applications.AddChild(fileManagers)

	// Media Players
	media := components.NewTreeNode("Media Players", "Audio and video players")
	media.AddChild(components.NewTreeNode("MPV", "Minimalist video player"))
	media.AddChild(components.NewTreeNode("VLC", "Feature-rich media player"))
	media.AddChild(components.NewTreeNode("imv", "Image viewer"))
	media.AddChild(components.NewTreeNode("Spotify", "Music streaming"))
	applications.AddChild(media)

	// Text Editors
	editors := components.NewTreeNode("Text Editors", "Code and text editors")
	editors.AddChild(components.NewTreeNode("Neovim", "Vim-based text editor"))
	editors.AddChild(components.NewTreeNode("Helix", "Modern modal editor"))
	editors.AddChild(components.NewTreeNode("VS Code", "Microsoft's code editor"))
	editors.AddChild(components.NewTreeNode("Emacs", "Extensible text editor"))
	applications.AddChild(editors)

	// Development Tools
	devTools := components.NewTreeNode("Development Tools", "Programming and development")
	devTools.AddChild(components.NewTreeNode("Git", "Version control system"))
	devTools.AddChild(components.NewTreeNode("Docker", "Container platform"))
	devTools.AddChild(components.NewTreeNode("Node.js", "JavaScript runtime"))
	devTools.AddChild(components.NewTreeNode("Go", "Go programming language"))
	devTools.AddChild(components.NewTreeNode("Rust", "Rust programming language"))
	devTools.AddChild(components.NewTreeNode("Python", "Python programming language"))
	// Select Git by default
	devTools.Children[0].State = components.Checked
	devTools.UpdateParentState()
	applications.AddChild(devTools)

	applications.UpdateParentState()
	root.AddChild(applications)

	// Update root state based on children
	root.UpdateParentState()

	return root
}

func (m packageSelectionModel) Update(msg tea.Msg) (packageSelectionModel, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			if m.tree.Finished {
				m.confirmed = true
				return m, nil
			}
		}
	}

	var cmd tea.Cmd
	m.tree, cmd = m.tree.Update(msg)

	// Check if user pressed enter in the tree
	if m.tree.Finished {
		m.confirmed = true
	}

	return m, cmd
}

func (m packageSelectionModel) View() string {
	return m.ViewWithSize(100, 25) // Default size
}

func (m packageSelectionModel) ViewWithSize(width, contentHeight int) string {
	// Calculate space for header and footer
	headerHeight := 5
	footerHeight := 5
	treeHeight := contentHeight - headerHeight - footerHeight

	header := components.Header(
		"Package Selection",
		"Navigate with ↑↓/jk, expand/collapse with →←/hl or e, select with space",
		80,
	)

	// Render tree
	treeContent := m.tree.Render(width-4, treeHeight)

	// Create container for tree
	treeStyle := lipgloss.NewStyle().
		Width(width).
		Padding(0, 2)

	content := treeStyle.Render(treeContent)

	// Footer with keybindings
	footer := components.Footer([]components.Keybinding{
		{Key: "↑↓/jk", Desc: "navigate"},
		{Key: "→←/hl", Desc: "expand/collapse"},
		{Key: "space", Desc: "toggle"},
		{Key: "enter", Desc: "confirm"},
		{Key: "esc", Desc: "back"},
	}, width)

	return lipgloss.JoinVertical(lipgloss.Left, header, "", content, "", footer)
}
