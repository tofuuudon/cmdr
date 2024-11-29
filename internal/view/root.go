package view

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/tofuuudon/cmdr/internal/loader"
)

var (
	commonStyle    = lipgloss.NewStyle().Border(lipgloss.NormalBorder())
	focusedStyle   = lipgloss.NewStyle().Inherit(commonStyle).BorderForeground(lipgloss.Color("63"))
	unfocusedStyle = lipgloss.NewStyle().Inherit(commonStyle)
)

type RootModel struct {
	indexModel   indexModel
	commandModel commandModel
	command      list.Item
}

func (m RootModel) Init() tea.Cmd {
	return nil
}

func (m RootModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	commands := loader.GetCommands()

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q":
			return m, tea.Quit
		case "enter":
			item, ok := m.indexModel.list.SelectedItem().(loader.Command)
			if ok {
				m.command = item
			}
		case "backspace":
			m.command = nil
		}
	}

	m.indexModel, _ = m.indexModel.Update(msg, &commands, &m.command)
	m.commandModel, _ = m.commandModel.Update(msg, &m.command)
	return m, nil
}

func (m RootModel) View() string {
	left := m.indexModel.View(m.command == nil)
	right := m.commandModel.View(m.command != nil)

	return lipgloss.JoinHorizontal(lipgloss.Top, left, right)
}

func RootView() RootModel {
	indexModel := indexModel{list: list.New([]list.Item{}, list.NewDefaultDelegate(), 0, 50)}
	commandModel := commandModel{}

	return RootModel{
		indexModel:   indexModel,
		commandModel: commandModel,
		command:      nil,
	}
}
