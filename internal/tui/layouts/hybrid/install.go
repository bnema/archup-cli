package hybrid

import (
	"github.com/bnema/archup-cli/internal/tui/components"
	tea "github.com/charmbracelet/bubbletea"
)

type installationModel struct {
	components.InstallationModel
	Complete bool
}

func newInstallationModel() installationModel {
	tasks := []string{
		"Updating package database",
		"Installing base system packages",
		"Configuring CPU governor",
		"Configuring TLP power management",
		"Configuring Pacman",
		"Installing compositor",
		"Installing desktop environment",
		"Installing selected applications",
		"Deploying theme and configurations",
		"Configuring system services",
		"Setting up user environment",
		"Cleaning up package cache",
	}

	return installationModel{
		InstallationModel: components.NewInstallationModel(tasks),
	}
}

func (m installationModel) Init() tea.Cmd {
	return m.InstallationModel.Init()
}

func (m installationModel) Update(msg tea.Msg) (installationModel, tea.Cmd) {
	var cmd tea.Cmd
	m.InstallationModel, cmd = m.InstallationModel.Update(msg)

	// Check if installation is complete
	if m.InstallationModel.ShouldClose || m.InstallationModel.ShouldReboot {
		m.Complete = true
	}

	return m, cmd
}

func (m installationModel) View() string {
	return m.InstallationModel.View()
}
