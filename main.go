package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
)

func main() {
	/// Flag args
	baseURL := flag.String("baseURL", "http://localhost:1337", "Strapi base url")
	singleTypes := flag.String("single", "nav", "Single types, seperated by ,")
	contentTypes := flag.String("content", "post", "Single types, seperated by ,")

	flag.Parse()

	fmt.Println(*baseURL)
	fmt.Println(*singleTypes)
	fmt.Println(*contentTypes)

	res, err := http.Get(fmt.Sprintf("%s/%s", *baseURL, *singleTypes))
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()
	data := map[string]interface{}{}
	err = json.NewDecoder(res.Body).Decode(&data)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(data)
}
