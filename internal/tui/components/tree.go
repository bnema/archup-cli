package components

import (
	"strings"

	"github.com/bnema/archup-cli/internal/theme"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// CheckState represents the state of a checkbox in the tree
type CheckState int

const (
	Unchecked CheckState = iota // None of the children are checked
	Partial                     // Some children are checked
	Checked                     // All children are checked
)

// TreeNode represents a single node in the tree
type TreeNode struct {
	Name        string
	Description string
	Children    []*TreeNode
	Parent      *TreeNode
	Expanded    bool
	State       CheckState
	IsRadio     bool // If true, only one child can be selected (like compositor choice)
}

// NewTreeNode creates a new tree node
func NewTreeNode(name, description string) *TreeNode {
	return &TreeNode{
		Name:        name,
		Description: description,
		Children:    []*TreeNode{},
		Expanded:    true, // Start expanded by default
		State:       Unchecked,
	}
}

// AddChild adds a child node
func (n *TreeNode) AddChild(child *TreeNode) *TreeNode {
	child.Parent = n
	n.Children = append(n.Children, child)
	return n
}

// ToggleExpanded toggles the expanded state
func (n *TreeNode) ToggleExpanded() {
	if len(n.Children) > 0 {
		n.Expanded = !n.Expanded
	}
}

// ToggleCheck toggles the check state
func (n *TreeNode) ToggleCheck() {
	// If this is a radio group, handle special logic
	if n.Parent != nil && n.Parent.IsRadio {
		// Uncheck all siblings first
		for _, sibling := range n.Parent.Children {
			sibling.State = Unchecked
		}
		// Check this node
		n.State = Checked
		n.Parent.UpdateParentState()
		return
	}

	// Regular checkbox logic
	switch n.State {
	case Unchecked, Partial:
		n.State = Checked
		n.CheckAllChildren(true)
	case Checked:
		n.State = Unchecked
		n.CheckAllChildren(false)
	}

	// Update parent states up the tree
	if n.Parent != nil {
		n.Parent.UpdateParentState()
	}
}

// CheckAllChildren recursively checks/unchecks all children
func (n *TreeNode) CheckAllChildren(checked bool) {
	state := Unchecked
	if checked {
		state = Checked
	}

	for _, child := range n.Children {
		child.State = state
		child.CheckAllChildren(checked)
	}
}

// UpdateParentState updates this node's state based on children
func (n *TreeNode) UpdateParentState() {
	if len(n.Children) == 0 {
		return
	}

	// If this is a radio group, check if any child is checked
	if n.IsRadio {
		for _, child := range n.Children {
			switch child.State {
			case Checked:
				n.State = Checked
				if n.Parent != nil {
					n.Parent.UpdateParentState()
				}
				return
			}
		}
		n.State = Unchecked
		if n.Parent != nil {
			n.Parent.UpdateParentState()
		}
		return
	}

	// Regular checkbox logic
	checkedCount := 0
	for _, child := range n.Children {
		switch child.State {
		case Checked:
			checkedCount++
		case Partial:
			n.State = Partial
			if n.Parent != nil {
				n.Parent.UpdateParentState()
			}
			return
		}
	}

	if checkedCount == 0 {
		n.State = Unchecked
	} else if checkedCount == len(n.Children) {
		n.State = Checked
	} else {
		n.State = Partial
	}

	if n.Parent != nil {
		n.Parent.UpdateParentState()
	}
}

// GetSelectedItems returns a flat list of all checked leaf nodes
func (n *TreeNode) GetSelectedItems() []string {
	var items []string

	// Only add leaf nodes that are checked
	if len(n.Children) == 0 && n.State == Checked {
		items = append(items, n.Name)
		return items
	}

	// Recurse through children
	for _, child := range n.Children {
		items = append(items, child.GetSelectedItems()...)
	}

	return items
}

// FlattenVisible returns a flat list of visible nodes for cursor navigation
func (n *TreeNode) FlattenVisible() []*TreeNode {
	var nodes []*TreeNode
	nodes = append(nodes, n)

	if n.Expanded {
		for _, child := range n.Children {
			nodes = append(nodes, child.FlattenVisible()...)
		}
	}

	return nodes
}

// CountDepth returns the depth of this node in the tree
func (n *TreeNode) CountDepth() int {
	depth := 0
	current := n
	for current.Parent != nil {
		depth++
		current = current.Parent
	}
	return depth
}

// TreeModel represents a reusable tree component
type TreeModel struct {
	Root   *TreeNode
	cursor int // Index in the flattened visible nodes

	// Callback when selection is confirmed
	Finished bool
}

// NewTreeModel creates a new reusable tree model
func NewTreeModel(root *TreeNode) TreeModel {
	return TreeModel{
		Root:   root,
		cursor: 0,
	}
}

// Init initializes the tree model
func (m TreeModel) Init() tea.Cmd {
	return nil
}

// Update handles tree navigation and selection
func (m TreeModel) Update(msg tea.Msg) (TreeModel, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "up", "k":
			m.MoveCursor(-1)
		case "down", "j":
			m.MoveCursor(1)
		case "left", "h":
			m.handleLeft()
		case "right", "l":
			m.handleRight()
		case " ":
			m.ToggleCurrentCheck()
		case "e":
			m.ToggleCurrentExpanded()
		case "enter":
			m.Finished = true
			return m, nil
		}
	}
	return m, nil
}

// View renders the tree
func (m TreeModel) View() string {
	return m.Render(100, 20) // Default size, will be overridden by parent
}

// GetVisibleNodes returns all currently visible nodes
func (m *TreeModel) GetVisibleNodes() []*TreeNode {
	return m.Root.FlattenVisible()
}

// GetCurrentNode returns the node at the cursor position
func (m *TreeModel) GetCurrentNode() *TreeNode {
	nodes := m.GetVisibleNodes()
	if m.cursor >= 0 && m.cursor < len(nodes) {
		return nodes[m.cursor]
	}
	return nil
}

// MoveCursor moves the cursor up or down
func (m *TreeModel) MoveCursor(delta int) {
	nodes := m.GetVisibleNodes()
	m.cursor += delta

	if m.cursor < 0 {
		m.cursor = 0
	}
	if m.cursor >= len(nodes) {
		m.cursor = len(nodes) - 1
	}
}

// ToggleCurrentExpanded toggles the expanded state of current node
func (m *TreeModel) ToggleCurrentExpanded() {
	if node := m.GetCurrentNode(); node != nil {
		node.ToggleExpanded()
	}
}

// ToggleCurrentCheck toggles the check state of current node
func (m *TreeModel) ToggleCurrentCheck() {
	if node := m.GetCurrentNode(); node != nil {
		node.ToggleCheck()
	}
}

// handleLeft handles left arrow key
func (m *TreeModel) handleLeft() {
	if node := m.GetCurrentNode(); node != nil {
		if node.Expanded && len(node.Children) > 0 {
			node.Expanded = false
		} else if node.Parent != nil {
			// Move to parent
			nodes := m.GetVisibleNodes()
			for i, n := range nodes {
				if n == node.Parent {
					m.cursor = i
					break
				}
			}
		}
	}
}

// handleRight handles right arrow key
func (m *TreeModel) handleRight() {
	if node := m.GetCurrentNode(); node != nil && len(node.Children) > 0 {
		node.Expanded = true
	}
}

// Render renders the tree view
func (m *TreeModel) Render(width, height int) string {
	nodes := m.GetVisibleNodes()
	var lines []string

	for i, node := range nodes {
		line := m.renderNode(node, i == m.cursor)
		lines = append(lines, line)
	}

	// Join lines and ensure we don't exceed height
	content := strings.Join(lines, "\n")

	// Add scrolling if needed (simple truncation for now)
	contentLines := strings.Split(content, "\n")
	if len(contentLines) > height {
		// Keep cursor in view
		start := m.cursor - height/2
		if start < 0 {
			start = 0
		}
		end := start + height
		if end > len(contentLines) {
			end = len(contentLines)
			start = end - height
			if start < 0 {
				start = 0
			}
		}
		contentLines = contentLines[start:end]
		content = strings.Join(contentLines, "\n")
	}

	return content
}

// renderNode renders a single node with proper indentation and styling
func (m *TreeModel) renderNode(node *TreeNode, isCursor bool) string {
	depth := node.CountDepth()

	// Build the line parts
	var parts []string

	// Indentation (tree structure lines)
	indent := strings.Repeat("  ", depth)
	if depth > 0 {
		// Replace last indent with tree connector
		if len(node.Parent.Children) > 0 && node.Parent.Children[len(node.Parent.Children)-1] == node {
			indent = strings.Repeat("  ", depth-1) + "└─"
		} else {
			indent = strings.Repeat("  ", depth-1) + "├─"
		}
	}

	// Cursor indicator
	cursor := " "
	if isCursor {
		cursor = theme.KeybindingStyle.Render("▶")
	}

	// Expand/collapse indicator
	expandIndicator := " "
	if len(node.Children) > 0 {
		if node.Expanded {
			expandIndicator = "▼"
		} else {
			expandIndicator = "▶"
		}
	}

	// Checkbox/radio indicator
	var checkBox string
	if node.Parent != nil && node.Parent.IsRadio {
		// Radio button
		switch node.State {
		case Checked:
			checkBox = "(●)"
		default:
			checkBox = "(○)"
		}
	} else {
		// Checkbox
		switch node.State {
		case Checked:
			checkBox = "[✓]"
		case Partial:
			checkBox = "[~]"
		default:
			checkBox = "[ ]"
		}
	}

	// Node name and description
	nameStyle := theme.BodyStyle
	if isCursor {
		nameStyle = theme.ListItemFocusedStyle.Bold(true)
	}

	name := nameStyle.Render(node.Name)
	desc := ""
	if node.Description != "" {
		descStyle := theme.HintStyle
		if isCursor {
			descStyle = descStyle.Foreground(lipgloss.Color(theme.PrimaryText))
		}
		desc = " " + descStyle.Render("- "+node.Description)
	}

	// Color the checkbox based on state
	checkBoxStyle := lipgloss.NewStyle()
	switch node.State {
	case Checked:
		checkBoxStyle = checkBoxStyle.Foreground(lipgloss.Color(theme.SuccessGreen))
	case Partial:
		checkBoxStyle = checkBoxStyle.Foreground(lipgloss.Color(theme.WarmOrange))
	default:
		checkBoxStyle = checkBoxStyle.Foreground(lipgloss.Color(theme.DimmedText))
	}

	// Assemble the line
	parts = append(parts, indent)
	parts = append(parts, cursor)
	parts = append(parts, " ")
	parts = append(parts, expandIndicator)
	parts = append(parts, " ")
	parts = append(parts, checkBoxStyle.Render(checkBox))
	parts = append(parts, "  ")
	parts = append(parts, name)
	parts = append(parts, desc)

	return strings.Join(parts, "")
}
