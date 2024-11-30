package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/tofuuudon/cmdr/internal/view"
)

func main() {
	p := tea.NewProgram(view.RootView(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
}
