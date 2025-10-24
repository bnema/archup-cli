package linear

import (
	"github.com/bnema/archup-cli/internal/tui/components"
	tea "github.com/charmbracelet/bubbletea"
)

type installationModel struct {
	components.InstallationModel
}

func newInstallationModel() installationModel {
	tasks := []string{
		"Installing base packages",
		"Installing compositor",
		"Installing applications",
		"Configuring services",
		"Deploying configurations",
		"Finalizing setup",
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
	return m, cmd
}

func (m installationModel) View() string {
	return m.InstallationModel.View()
}
