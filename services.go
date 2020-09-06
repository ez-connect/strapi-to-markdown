package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func request(baseURL, endpoint string, i interface{}) error {
	res, err := http.Get(fmt.Sprintf("%s/%s", baseURL, endpoint))
	if err != nil {
		return err
	}

	defer res.Body.Close()
	err = json.NewDecoder(res.Body).Decode(i)
	if err != nil {
		return err
	}

	return nil
}

func find(baseURL, endpoint string) ([]map[string]interface{}, error) {
	data := []map[string]interface{}{}
	err := request(baseURL, endpoint, &data)
	return data, err
}

func findOne(baseURL, endpoint string) (map[string]interface{}, error) {
	data := map[string]interface{}{}
	err := request(baseURL, endpoint, &data)
	return data, err
}

func download(baseURL, path, output string) error {
	res, err := http.Get(fmt.Sprintf("%s/%s", baseURL, path))
	if err != nil {
		return err
	}

	defer res.Body.Close()

	// Create the file
	out, err := os.Create(fmt.Sprintf("%s%s", output, path))
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, res.Body)
	return err

}
