package tree

// CheckState represents the state of a checkbox in the tree
type CheckState int

const (
	Unchecked CheckState = iota // None of the children are checked
	Partial                     // Some children are checked
	Checked                     // All children are checked
)

// Node represents a single node in the tree
type Node struct {
	Name        string
	Description string
	Children    []*Node
	Parent      *Node
	Expanded    bool
	State       CheckState
	IsRadio     bool // If true, only one child can be selected (like compositor choice)
}

// NewNode creates a new tree node
func NewNode(name, description string) *Node {
	return &Node{
		Name:        name,
		Description: description,
		Children:    []*Node{},
		Expanded:    true, // Start expanded by default
		State:       Unchecked,
	}
}

// AddChild adds a child node
func (n *Node) AddChild(child *Node) *Node {
	child.Parent = n
	n.Children = append(n.Children, child)
	return n
}

// ToggleExpanded toggles the expanded state
func (n *Node) ToggleExpanded() {
	if len(n.Children) > 0 {
		n.Expanded = !n.Expanded
	}
}

// ToggleCheck toggles the check state
func (n *Node) ToggleCheck() {
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
func (n *Node) CheckAllChildren(checked bool) {
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
func (n *Node) UpdateParentState() {
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
func (n *Node) GetSelectedItems() []string {
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
func (n *Node) FlattenVisible() []*Node {
	var nodes []*Node
	nodes = append(nodes, n)

	if n.Expanded {
		for _, child := range n.Children {
			nodes = append(nodes, child.FlattenVisible()...)
		}
	}

	return nodes
}

// CountDepth returns the depth of this node in the tree
func (n *Node) CountDepth() int {
	depth := 0
	current := n
	for current.Parent != nil {
		depth++
		current = current.Parent
	}
	return depth
}
