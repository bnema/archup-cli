package cmd

import (
	"fmt"
	"os"

	"github.com/bnema/archup-cli/internal/logger"
	"github.com/bnema/archup-cli/internal/tui/layouts/hybrid"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
)

var demoHybridCmd = &cobra.Command{
	Use:   "demo-hybrid",
	Short: "Run the hybrid wizard demo (linear + tree)",
	Long: `Launch the hybrid wizard demo - combining linear wizard with tree view.

This demo showcases a hybrid approach that combines:
  • Linear step-by-step wizard for focused configuration questions
  • Tree view for hierarchical package selection

The wizard flow includes:
  1. Welcome screen
  2. Hardware detection with detailed information
  3. System configuration (CPU governor, TLP, Pacman settings)
  4. Compositor selection
  5. Package selection using tree view
  6. Theme/Config selection (Omarchy vs Bleu vs Stock)
  7. Review & confirmation
  8. Installation progress

This demonstrates the recommended UX pattern for ArchUp CLI.`,
	Run: func(cmd *cobra.Command, args []string) {
		logger.Info("Starting Hybrid Wizard Demo")

		// Create the Bubbletea program
		p := tea.NewProgram(
			hybrid.New(),
			tea.WithAltScreen(),
			tea.WithMouseCellMotion(),
		)

		// Run the program
		if _, err := p.Run(); err != nil {
			logger.Error("Error running hybrid wizard demo", "error", err)
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}

		logger.Info("Hybrid wizard demo completed")
	},
}

func init() {
	rootCmd.AddCommand(demoHybridCmd)
}
