package view

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var indexStyle = lipgloss.NewStyle()

type item struct {
	id, title, description, exec string
}

func (i item) Title() string       { return i.title }
func (i item) Description() string { return i.description }
func (i item) FilterValue() string { return i.title }

type indexModel struct {
	list list.Model
}

func (m indexModel) Init() tea.Cmd {
	return nil
}

func (m indexModel) Update(msg tea.Msg, commands *[]list.Item, command *list.Item) (indexModel, tea.Cmd) {
	if *command != nil {
		return m, nil
	}

	m.list.SetItems(*commands)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch key := msg.String(); key {
		case "ctrl+c":
			return m, tea.Quit
		}
	case tea.WindowSizeMsg:
		x, y := commonStyle.GetFrameSize()
		width := msg.Width/3 - x
		height := msg.Height - y

		m.list.SetSize(width, height)
		indexStyle = indexStyle.Width(width)
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
