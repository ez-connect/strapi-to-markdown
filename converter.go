package main

import (
	"fmt"
	"os"
	"regexp"

	"gopkg.in/yaml.v3"
)

func markdown(data map[string]interface{}, body, exclude, output, baseURL, staticDir string) error {
	f, err := os.Create(output)
	if err != nil {
		return err
	}

	defer f.Close()

	content := ""
	if body != "" {
		if data[body] != nil {
			content = fmt.Sprintf("%s", data[body])
		}
		delete(data, body)
	}

	if exclude != "" {
		delete(data, exclude)
	}

	f.WriteString("---\n")
	buf, err := yaml.Marshal(data)
	if err != nil {
		return err
	}

	f.WriteString(string(buf))
	f.WriteString("---\n\n")

	if content != "" {
		f.WriteString(content)
		f.WriteString("\n")

		// Find media
		// https://regex101.com/codegen?language=golang
		re := regexp.MustCompile(`\[.*\]\((\/uploads\/.*)\)`)
		match := re.FindAllStringSubmatch(content, -1)
		for _, v := range match {
			if len(v) > 0 {
				path := v[1]
				download(
					baseURL,
					path,
					staticDir,
				)
			}
		}
	}

	return nil
}
