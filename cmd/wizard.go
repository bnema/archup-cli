package cmd

import (
	"fmt"

	"github.com/bnema/archup-cli/internal/logger"
	"github.com/spf13/cobra"
)

// wizardCmd represents the wizard command
var wizardCmd = &cobra.Command{
	Use:   "wizard",
	Short: "Run the interactive desktop setup wizard",
	Long: `Launch the interactive TUI wizard to configure your desktop environment.

The wizard will guide you through:
  1. Compositor selection (Niri, Hyprland, Sway, River)
  2. Hardware detection (GPU, audio, Bluetooth)
  3. Package selection (terminal, browser, file manager, etc.)
  4. Installation and configuration
  5. Service enablement

This wizard installs Tier 2 (Desktop Foundation) and Tier 3 (Compositor + Apps).`,
	Run: func(cmd *cobra.Command, args []string) {
		logger.Info("Starting ArchUp Desktop Wizard")

		// TODO: Launch Bubbletea TUI
		fmt.Println("Wizard implementation coming soon!")
		fmt.Println("")
		fmt.Println("The wizard will:")
		fmt.Println("  1. Detect your hardware")
		fmt.Println("  2. Let you choose a compositor")
		fmt.Println("  3. Select applications to install")
		fmt.Println("  4. Install and configure everything")
	},
}

func init() {
	rootCmd.AddCommand(wizardCmd)
}
