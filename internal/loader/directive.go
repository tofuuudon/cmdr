package loader

import (
	"encoding/json"
	"io"
	"os"

	"github.com/charmbracelet/bubbles/list"
)

type Command struct {
	ID   string `json:"id"`
	Name string `json:"title"`
	Desc string `json:"description"`
	Exec string `json:"exec"`
}

func (c Command) Title() string       { return c.Name }
func (c Command) Description() string { return c.Desc }
func (c Command) FilterValue() string { return c.Name }

type Directive struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Commands    []Command `json:"commands"`
}

func getDirectives() []Directive {
	config := GetConfig()

	var directives []Directive
	for _, path := range config.Directives {
		file, _ := os.Open(path)
		bytes, _ := io.ReadAll(file)

		var directive Directive
		json.Unmarshal(bytes, &directive)

		directives = append(directives, directive)
	}
	return directives
}

func GetCommands() []list.Item {
	directives := getDirectives()

	var commands []list.Item
	for _, directive := range directives {
		for _, command := range directive.Commands {
			commands = append(commands, command)
		}
	}
	return commands
}
