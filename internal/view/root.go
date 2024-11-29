package view

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	commonStyle    = lipgloss.NewStyle().Border(lipgloss.NormalBorder())
	focusedStyle   = lipgloss.NewStyle().Inherit(commonStyle).BorderForeground(lipgloss.Color("63"))
	unfocusedStyle = lipgloss.NewStyle().Inherit(commonStyle)
)

type RootModel struct {
	indexModel   indexModel
	commandModel commandModel
	focused      int8
}

func (m RootModel) Init() tea.Cmd {
	return nil
}

func (m RootModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q":
			return m, tea.Quit
		case "tab":
			m.focused = (m.focused + 1) % 2
		}
	}

	m.indexModel, _ = m.indexModel.Update(msg)
	m.commandModel, _ = m.commandModel.Update(msg)
	return m, nil
}

func (m RootModel) View() string {
	left := m.indexModel.View(m.focused == 0)
	right := m.commandModel.View(m.focused == 1)

	return lipgloss.JoinHorizontal(lipgloss.Top, left, right)
}

func RootView() RootModel {
	indexModel := indexModel{list: list.New([]list.Item{}, list.NewDefaultDelegate(), 0, 50)}
	commandModel := commandModel{}

	return RootModel{
		indexModel:   indexModel,
		commandModel: commandModel,
		focused:      0,
	}
}
