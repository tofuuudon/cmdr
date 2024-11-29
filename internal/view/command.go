package view

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var commandStyle = lipgloss.NewStyle().Width(60)

type commandModel struct {
	data string
}

func (m commandModel) Init() tea.Cmd {
	return nil
}

func (m commandModel) Update(msg tea.Msg) (commandModel, tea.Cmd) {
	m.data = "No data yet..."

	return m, nil
}

func (m commandModel) View(focused bool) string {
	if focused {
		return focusedStyle.Inherit(commandStyle).Render(m.data)
	}

	return unfocusedStyle.Inherit(commandStyle).Render(m.data)
}
