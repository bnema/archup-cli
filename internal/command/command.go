package command

import (
	"context"
	"fmt"
	"os/exec"
	"strings"

	"github.com/bnema/archup-cli/internal/logger"
)

// Run executes a command with context and returns stdout, stderr, and error
func Run(ctx context.Context, name string, args ...string) (string, error) {
	cmdStr := fmt.Sprintf("%s %s", name, strings.Join(args, " "))
	logger.Debug("Running command", "cmd", cmdStr)

	cmd := exec.CommandContext(ctx, name, args...)
	output, err := cmd.CombinedOutput()
	outputStr := strings.TrimSpace(string(output))

	if err != nil {
		logger.Error("Command failed", "cmd", cmdStr, "error", err)
		if outputStr != "" {
			logger.Debug("Command output", "output", outputStr)
		}
		return outputStr, err
	}

	logger.Debug("Command succeeded", "cmd", cmdStr)
	if outputStr != "" {
		logger.Debug("Command output", "output", outputStr)
	}
	return outputStr, nil
}

// MustRun executes a command and panics on error
func MustRun(ctx context.Context, name string, args ...string) string {
	output, err := Run(ctx, name, args...)
	if err != nil {
		panic(fmt.Sprintf("command failed: %s %v: %v", name, args, err))
	}
	return output
}
