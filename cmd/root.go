package cmd

import (
	"os"

	"github.com/bnema/archup-cli/internal/logger"
	"github.com/spf13/cobra"
)

var (
	// Version is set via build flags
	Version = "dev"
	// Debug enables debug logging
	Debug bool
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "archup",
	Short: "ArchUp Desktop Wizard - Install and configure Wayland compositors",
	Long: `ArchUp is a desktop environment setup wizard for Arch Linux.

After installing the barebone system with the ArchUp installer,
run 'archup wizard' to install desktop infrastructure and your
chosen compositor (Niri, Hyprland, Sway, or River).

Features:
  • Auto-detects hardware (GPU, CPU)
  • Installs desktop foundation (graphics, audio, Wayland, Bluetooth, printing)
  • User selects compositor and applications
  • Deploys configs and themes
  • Enables required services

Example:
  archup wizard          # Run interactive setup wizard
  archup version         # Show version information
  archup --debug wizard  # Run with debug logging`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		// Initialize logger
		if err := logger.Init(Debug); err != nil {
			// Non-fatal, just print to stderr
			os.Stderr.WriteString("Warning: could not initialize log file: " + err.Error() + "\n")
		}
	},
	PersistentPostRun: func(cmd *cobra.Command, args []string) {
		// Close log file
		logger.Close()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Global flags
	rootCmd.PersistentFlags().BoolVar(&Debug, "debug", false, "Enable debug logging")
}
