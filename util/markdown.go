package util

import (
	"fmt"
	"os"
	"path"
	"regexp"
	"strings"

	"github.com/gosimple/slug"
	"gopkg.in/yaml.v3"
)

// func mapping(m map[string]string, data map[string]interface{}) map[string]interface{} {
// 	for k, v := range m {
// 		if data[v] != nil {
// 			data[k] = data[v]
// 			delete(data, v)
// 		}
// 	}

// 	return data
// }

func WriteMarkdown(data map[string]interface{}, exclude, contentDir, mediaDir string) error {
	slug := slug.Make(fmt.Sprintf("%v", data["title"]))
	name := fmt.Sprintf("%s-%v", slug, data["id"])
	category := ""
	if data["category"] != nil {
		category = fmt.Sprintf("%v", data["category"])
	}

	var filename string
	if category == "" {
		filename = path.Join(contentDir, name)
	} else {
		filename = path.Join(contentDir, category, name)
	}

	f, err := os.Create(fmt.Sprintf("%s.md", filename))
	if err != nil {
		return err
	}

	defer f.Close()

	if exclude != "" {
		items := strings.Split(exclude, ",")
		for _, v := range items {
			delete(data, v)
		}
	}

	f.WriteString("---\n")
	frontMatter := map[string]interface{}{}
	for k, v := range data {
		if k == "content" {
			continue
		}

		frontMatter[k] = v
	}

	buf, err := yaml.Marshal(frontMatter)
	if err != nil {
		return err
	}

	f.WriteString(string(buf))
	f.WriteString("---\n\n")

	content := ""
	if data["content"] != nil {
		content = fmt.Sprintf("%s", data["content"])
	}
	if content != "" {
		f.WriteString(content)
		f.WriteString("\n")
	}

	return nil
}

func WriteMedia(baseURL string, path string, data map[string]interface{}, mediaDir string) error {
	content := ""
	if data["content"] != nil {
		content = fmt.Sprintf("%s", data["content"])
	}

	// https://regex101.com/codegen?language=golang
	re := regexp.MustCompile(`\[.*\]\((\/uploads\/.*)\)`)
	match := re.FindAllStringSubmatch(content, -1)
	for _, v := range match {
		if len(v) > 0 {
			path := v[1]
			if err := download(baseURL, path, mediaDir); err != nil {
				return err
			}
		}
	}

	return nil
}
