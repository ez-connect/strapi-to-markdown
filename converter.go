package main

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

func markdown(data map[string]interface{}, body string, output string) error {
	f, err := os.Create(output)
	if err != nil {
		return err
	}

	defer f.Close()

	content := ""
	if body != "" {
		content = fmt.Sprintf("%s", data[body])
		delete(data, body)
	}

	f.WriteString("---\n")
	buf, err := yaml.Marshal(data)
	if err != nil {
		return err
	}

	f.WriteString(string(buf))
	f.WriteString("---\n\n")
	f.WriteString(content)
	f.WriteString("\n")

	return nil
}
