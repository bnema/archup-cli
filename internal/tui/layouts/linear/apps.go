package linear

import (
	"github.com/bnema/archup-cli/internal/theme"
	"github.com/bnema/archup-cli/internal/tui/components"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type applicationsModel struct {
	checkboxList components.CheckboxList
	selectedApps []string
}

func newApplicationsModel() applicationsModel {
	items := []components.CheckboxItem{
		{Name: "Foot", Description: "Fast, lightweight Wayland terminal", Selected: true},
		{Name: "Firefox", Description: "Open source web browser", Selected: true},
		{Name: "Thunar", Description: "Modern file manager", Selected: false},
		{Name: "MPV", Description: "Minimalist video player", Selected: false},
		{Name: "imv", Description: "Command line image viewer", Selected: false},
		{Name: "grim + slurp", Description: "Screenshot utilities", Selected: true},
	}

	return applicationsModel{
		checkboxList: components.NewCheckboxList(items),
	}
}

func (m applicationsModel) Update(msg tea.Msg) (applicationsModel, tea.Cmd) {
	var cmd tea.Cmd
	m.checkboxList, cmd = m.checkboxList.Update(msg)

	if m.checkboxList.Finished {
		m.selectedApps = m.checkboxList.GetSelected()
	}

	return m, cmd
}

func (m applicationsModel) View() string {
	title := components.Header("Application Selection", "Choose applications to install", 80)
	items := m.checkboxList.View()
	hint := theme.HintStyle.Render("Space to toggle, Enter to continue")

	return lipgloss.JoinVertical(
		lipgloss.Left,
		title,
		"",
		items,
		"",
		hint,
	)
}
