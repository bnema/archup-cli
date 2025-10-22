package cmd

import (
	"context"
	"fmt"

	"github.com/bnema/archup-cli/internal/command"
	"github.com/bnema/archup-cli/internal/logger"
	"github.com/spf13/cobra"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update archup-cli to the latest version",
	Long: `Update archup-cli to the latest version from GitHub.

This command runs: go install github.com/bnema/archup-cli@latest

The binary will be installed to $GOBIN (or $GOPATH/bin or ~/go/bin).`,
	Run: func(cmd *cobra.Command, args []string) {
		logger.Info("Updating archup-cli to latest version...")

		ctx := context.Background()
		output, err := command.Run(ctx, "go", "install", "github.com/bnema/archup-cli@latest")

		if err != nil {
			logger.Error("Update failed", "error", err)
			fmt.Println("Update failed:", err)
			if output != "" {
				fmt.Println(output)
			}
			return
		}

		logger.Info("Update successful")
		fmt.Println("archup-cli updated to latest version")
		fmt.Println("")
		fmt.Println("Run 'archup version' to verify the new version")
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
}
