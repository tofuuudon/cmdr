package view

import (
	"log"
	"os/exec"
	"strings"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var commandStyle = lipgloss.NewStyle()

type commandModel struct {
	spinner spinner.Model
	data    string
	exec    string
}

func (m commandModel) Init() tea.Cmd {
	return nil
}

func (m commandModel) Update(msg tea.Msg) (commandModel, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		x, y := commonStyle.GetFrameSize()
		width := msg.Width/3*2 - x
		height := msg.Height - y
		commandStyle = commandStyle.Width(width).Height(height)
	}

	return m, nil
}

func (m commandModel) View(focused bool) string {
	if focused {
		return focusedStyle.Inherit(commandStyle).Render(m.data)
	}
	return unfocusedStyle.Inherit(commandStyle).Render(m.data)
}

func (m commandModel) ExecuteCommand(execCommand string) (commandModel, tea.Cmd) {
	parts := strings.Fields(execCommand)
	cmd := exec.Command(parts[0], parts[1:]...)
	output, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	m.data = string(output)
	return m, nil
}
