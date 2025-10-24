package cmd

import (
	"fmt"
	"os"

	"github.com/bnema/archup-cli/internal/logger"
	"github.com/bnema/archup-cli/internal/tui/layouts/linear"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
)

var demoLinearCmd = &cobra.Command{
	Use:   "demo-linear",
	Short: "Run the linear step-by-step wizard demo",
	Long: `Launch the linear wizard demo - a step-by-step guided installation flow.

This is a spike implementation to test the linear wizard UX pattern
using the Bleu design system.

Features:
  - Welcome screen with logo
  - Hardware detection with spinner
  - Compositor selection with list
  - Application selection with checkboxes
  - Confirmation screen with review
  - Installation progress with progress bar`,
	Run: func(cmd *cobra.Command, args []string) {
		logger.Info("Starting Linear Wizard Demo")

		// Create the Bubbletea program
		p := tea.NewProgram(
			linear.New(),
			tea.WithAltScreen(),
			tea.WithMouseCellMotion(),
		)

		// Run the program
		if _, err := p.Run(); err != nil {
			logger.Error("Error running linear wizard demo", "error", err)
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}

		logger.Info("Linear wizard demo completed")
	},
}

func init() {
	rootCmd.AddCommand(demoLinearCmd)
}
