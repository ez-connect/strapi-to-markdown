package util

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
)

func Find(baseURL, path string) ([]map[string]interface{}, error) {
	data := []map[string]interface{}{}
	err := request(baseURL, path, &data)
	return data, err
}

func FindOne(baseURL, path string) (map[string]interface{}, error) {
	data := map[string]interface{}{}
	err := request(baseURL, path, &data)
	return data, err
}

func request(baseURL, path string, i interface{}) error {
	res, err := http.Get(fmt.Sprintf("%s/%s", baseURL, path))
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

func download(baseURL, media, outputDir string) error {
	url := fmt.Sprintf("%s%s", baseURL, media)
	res, err := http.Get(url)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	// Create the file
	filename := path.Base(url)
	out, err := os.Create(path.Join(outputDir, filename))
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, res.Body)
	return err

}
