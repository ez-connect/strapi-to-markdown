package main

import (
	"flag"
	"fmt"
	"log"
)

func main() {
	/// Flag args
	baseURL := flag.String("baseURL", "http://localhost:1337", "Strapi base url")
	singleType := flag.String("single", "", "A single type name")
	collectionType := flag.String("collection", "post", "A colletion type name")
	body := flag.String("body", "body", "Field name to write to markdown content")
	output := flag.String("output", "content", "Output directory")
	name := flag.String("name", "slug", "Output file name for single type or a field name of collection type")

	flag.Parse()

	if *singleType == "" && *collectionType == "" {
		log.Fatal("Missing type name")
	}

	if *output == "" {
		log.Fatal("Missing output directory")
	}

	if *name == "" {
		log.Fatal("Missing output file name")
	}

	/// Single type
	if *singleType != "" {
		data, err := findOne(*baseURL, *singleType)
		if err != nil {
			log.Fatal(err)
		}

		err = markdown(data, "", fmt.Sprintf("%s/%s.md", *output, *name))
		if err != nil {
			log.Fatal(err)
		}
	}

	/// Collection type
	if *collectionType != "" {
		data, err := find(*baseURL, fmt.Sprintf("%ss", *collectionType))
		if err != nil {
			log.Fatal(err)
		}

		for _, v := range data {
			date := fmt.Sprintf("%s", v["created_at"])[0:10]
			filename := fmt.Sprintf("%s-%s", date, v[*name])
			err = markdown(v, *body, fmt.Sprintf("%s/%s.md", *output, filename))
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}
