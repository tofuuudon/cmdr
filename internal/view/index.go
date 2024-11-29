package view

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/tofuuudon/cmdr/internal/loader"
)

var indexStyle = lipgloss.NewStyle().Width(50)

type item struct {
	id, title, description, exec string
}

func (i item) Title() string       { return i.title }
func (i item) Description() string { return i.description }
func (i item) FilterValue() string { return i.title }

type indexModel struct {
	list      list.Model
	commandID string
}

func (m indexModel) Init() tea.Cmd {
	return nil
}

func (m indexModel) Update(msg tea.Msg) (indexModel, tea.Cmd) {
	commands := loader.GetCommands()

	var items []list.Item
	for _, command := range commands {
		items = append(items, item{id: command.ID, title: command.Title, description: command.Description, exec: command.Exec})
	}
	m.list.SetItems(items)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch key := msg.String(); key {
		case "ctrl+c":
			return m, tea.Quit
		case "enter":
			i, ok := m.list.SelectedItem().(item)
			if ok {
				m.commandID = i.id
			}
			return m, nil
		}
	case tea.WindowSizeMsg:
		x, y := commonStyle.GetFrameSize()
		m.list.SetSize(msg.Width-x, msg.Height-y)
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m indexModel) View(focused bool) string {
	if focused {
		return focusedStyle.Inherit(indexStyle).Render(m.list.View())
	}
	return unfocusedStyle.Inherit(indexStyle).Render(m.list.View())
}
