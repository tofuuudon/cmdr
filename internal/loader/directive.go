package loader

import (
	"encoding/json"
	"io"
	"os"
)

type Command struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Exec        string `json:"exec"`
}

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

func GetCommands() []Command {
	directives := getDirectives()

	var commands []Command
	for _, directive := range directives {
		for _, command := range directive.Commands {
			commands = append(commands, command)
		}
	}
	return commands
}
