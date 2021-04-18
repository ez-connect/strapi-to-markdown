package main

import (
	"app/util"
	"flag"
	"fmt"
)

const (
	_version    = "v0.1.0"
	_baseURL    = "http://localhost:1337"
	_contentDir = "content"
	_mediaDir   = "static/images"
)

func main() {
	fmt.Println("Strapi to markdown", _version)

	/// Flag args
	baseURL := flag.String("b", _baseURL, "Base URL")
	path := flag.String("p", "articles", "Strapi base url")
	isSingle := flag.Bool("s", false, "Fetch single record")
	exclude := flag.String("e", "", "Exclude field name")
	contentDir := flag.String("c", _contentDir, "Output directory")
	mediaDir := flag.String("m", _mediaDir, "Static directory")
	flag.Parse()

	if *contentDir == "" {
		panic("Missing content output directory")
	}

	if *mediaDir == "" {
		panic("Missing media output directory")
	}

	if !*isSingle {
		data, err := util.Find(*baseURL, *path)
		if err != nil {
			panic(err)
		}

		for _, v := range data {
			err := util.WriteMarkdown(v, *exclude, *contentDir, *mediaDir)
			if err != nil {
				panic(err)
			}

			// Download media
			if err := util.WriteMedia(*baseURL, *path, v, *mediaDir); err != nil {
				panic(err)
			}
		}
	} else {
		data, err := util.FindOne(*baseURL, *path)
		if err != nil {
			panic(err)
		}

		err = util.WriteMarkdown(data, *exclude, *contentDir, *mediaDir)
		if err != nil {
			panic(err)
		}

		// Download media
		if err := util.WriteMedia(*baseURL, *path, data, *mediaDir); err != nil {
			panic(err)
		}
	}
}
