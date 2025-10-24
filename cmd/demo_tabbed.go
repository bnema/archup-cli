package cmd

import (
	"fmt"
	"os"

	"github.com/bnema/archup-cli/internal/logger"
	"github.com/bnema/archup-cli/internal/tui/layouts/tabbed"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
)

var demoTabbedCmd = &cobra.Command{
	Use:   "demo-tabbed",
	Short: "Run the tabbed configuration panel demo",
	Long:  `Demonstrates a tabbed interface for configuring the system, compositor, and applications.`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := runDemoTabbed(); err != nil {
			logger.Error("Failed to run tabbed demo", "error", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(demoTabbedCmd)
}

func runDemoTabbed() error {
	p := tea.NewProgram(
		tabbed.New(),
		tea.WithAltScreen(),
		tea.WithMouseCellMotion(),
	)

	if _, err := p.Run(); err != nil {
		return fmt.Errorf("error running tabbed demo: %w", err)
	}

	return nil
}
