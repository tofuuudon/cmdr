package loader

import (
	"encoding/json"
	"io"
	"os"
)

type Directive struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

func GetDirectives() []Directive {
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
