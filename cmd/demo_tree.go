package cmd

import (
	"fmt"
	"os"

	"github.com/bnema/archup-cli/internal/tui/layouts/tree"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
)

var demoTreeCmd = &cobra.Command{
	Use:   "demo-tree",
	Short: "Run the tree/checklist navigation demo",
	Long: `Launches the tree/checklist demo for the ArchUp TUI.

This demo showcases:
- Hierarchical package selection with tree structure
- Expand/collapse navigation
- Tri-state checkboxes (checked, partial, unchecked)
- Radio button groups for exclusive selection
- Vim-style navigation (hjkl)
- Real-time state propagation`,
	Run: func(cmd *cobra.Command, args []string) {
		// Create the tree demo model
		m := tree.New()

		// Create the Bubbletea program
		p := tea.NewProgram(
			m,
			tea.WithAltScreen(),
			tea.WithMouseCellMotion(),
		)

		// Run the program
		if _, err := p.Run(); err != nil {
			fmt.Printf("Error running tree demo: %v\n", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(demoTreeCmd)
}
