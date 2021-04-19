package main

import (
	"flag"
	"fmt"

	"app/lib"
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
	username := flag.String("u", "", "User name")
	password := flag.String("p", "", "Password")
	baseURL := flag.String("b", _baseURL, "Base URL")
	path := flag.String("P", "articles", "Strapi base url")
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

	strapi := lib.Strapi{}
	strapi.SetBaseURL(*baseURL)
	markdown := lib.Markdown{}
	markdown.SetBaseURL(*baseURL)

	if *username != "" {
		if jwt, err := strapi.SignIn(*username, *password); err == nil {
			markdown.SetJWT(jwt)
		} else {
			panic(err)
		}
	}

	if !*isSingle {
		data, err := strapi.Find(*path)
		if err != nil {
			panic(err)
		}

		for _, v := range data {
			err := markdown.Write(v, *exclude, *contentDir, *mediaDir)
			if err != nil {
				panic(err)
			}

			// Download media
			if err := markdown.WriteMedia(*path, v, *mediaDir); err != nil {
				panic(err)
			}
		}
	} else {
		data, err := strapi.FindOne(*path)
		if err != nil {
			panic(err)
		}

		err = markdown.Write(data, *exclude, *contentDir, *mediaDir)
		if err != nil {
			panic(err)
		}

		// Download media
		if err := markdown.WriteMedia(*path, data, *mediaDir); err != nil {
			panic(err)
		}
	}
}
