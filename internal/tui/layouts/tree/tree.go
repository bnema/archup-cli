package tree

import (
	"strings"

	"github.com/bnema/archup-cli/internal/theme"
	"github.com/charmbracelet/lipgloss"
)

// TreeView handles rendering of the tree structure
type TreeView struct {
	root   *Node
	cursor int // Index in the flattened visible nodes
}

// NewTreeView creates a new tree view
func NewTreeView(root *Node) *TreeView {
	return &TreeView{
		root:   root,
		cursor: 0,
	}
}

// GetVisibleNodes returns all currently visible nodes
func (tv *TreeView) GetVisibleNodes() []*Node {
	return tv.root.FlattenVisible()
}

// GetCurrentNode returns the node at the cursor position
func (tv *TreeView) GetCurrentNode() *Node {
	nodes := tv.GetVisibleNodes()
	if tv.cursor >= 0 && tv.cursor < len(nodes) {
		return nodes[tv.cursor]
	}
	return nil
}

// MoveCursor moves the cursor up or down
func (tv *TreeView) MoveCursor(delta int) {
	nodes := tv.GetVisibleNodes()
	tv.cursor += delta

	if tv.cursor < 0 {
		tv.cursor = 0
	}
	if tv.cursor >= len(nodes) {
		tv.cursor = len(nodes) - 1
	}
}

// ToggleCurrentExpanded toggles the expanded state of current node
func (tv *TreeView) ToggleCurrentExpanded() {
	if node := tv.GetCurrentNode(); node != nil {
		node.ToggleExpanded()
	}
}

// ToggleCurrentCheck toggles the check state of current node
func (tv *TreeView) ToggleCurrentCheck() {
	if node := tv.GetCurrentNode(); node != nil {
		node.ToggleCheck()
	}
}

// Render renders the tree view
func (tv *TreeView) Render(width, height int) string {
	nodes := tv.GetVisibleNodes()
	var lines []string

	for i, node := range nodes {
		line := tv.renderNode(node, i == tv.cursor)
		lines = append(lines, line)
	}

	// Join lines and ensure we don't exceed height
	content := strings.Join(lines, "\n")

	// Add scrolling if needed (simple truncation for now)
	contentLines := strings.Split(content, "\n")
	if len(contentLines) > height {
		// Keep cursor in view
		start := tv.cursor - height/2
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
func (tv *TreeView) renderNode(node *Node, isCursor bool) string {
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

// GetRoot returns the root node
func (tv *TreeView) GetRoot() *Node {
	return tv.root
}
