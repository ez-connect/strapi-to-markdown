package main

import (
	"flag"
	"fmt"
	"log"
)

func main() {
	/// Flag args
	baseURL := *flag.String("baseURL", "http://localhost:1337", "Strapi base url")
	singleTypes := *flag.String("single", "nav", "Single types, seperated by ,")
	contentTypes := *flag.String("content", "post", "Single types, seperated by ,")
	body := *flag.String("body", "body", "Body field")
	output := *flag.String("output", "content", "Output location")

	flag.Parse()

	fmt.Println(baseURL)
	fmt.Println(singleTypes)
	fmt.Println(contentTypes)
	fmt.Println(body)
	fmt.Println(output)

	data, err := find(baseURL, "posts")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(data)

	markdown(data[0], body, "test.md")
}
