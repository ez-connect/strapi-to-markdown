package lib

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
)

type Rest struct {
	baseURL string
	jwt     string
}

func (r *Rest) SetBaseURL(value string) {
	r.baseURL = value
}

func (r *Rest) SetJWT(value string) {
	r.jwt = value
}

func (r *Rest) Get(p string, i interface{}) error {
	return r.Request(http.MethodGet, p, nil, i)
}

func (r *Rest) Post(p string, body, i interface{}) error {
	return r.Request(http.MethodPost, p, body, i)
}

func (r *Rest) Put(p string, body, i interface{}) error {
	return r.Request(http.MethodPut, p, body, i)
}

func (r *Rest) Patch(p string, body, i interface{}) error {
	return r.Request(http.MethodPatch, p, body, i)
}

func (r *Rest) Delete(p string, i interface{}) error {
	return r.Request(http.MethodPatch, p, nil, i)
}

func (r *Rest) Request(method, p string, body, i interface{}) error {
	var req *http.Request
	var res *http.Response
	var err error
	var data []byte

	url := fmt.Sprintf("%s/%s", r.baseURL, p)

	if body != nil {
		data, err = json.Marshal(body)
		if err != nil {
			return err
		}
	}

	if data == nil {
		req, err = http.NewRequest(method, url, nil)
	} else {
		req, err = http.NewRequest(method, url, bytes.NewBuffer(data))
	}

	if r.jwt != "" {
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", r.jwt))
	}

	client := http.Client{}
	if res, err = client.Do(req); err != nil {
		return err
	}

	defer res.Body.Close()
	err = json.NewDecoder(res.Body).Decode(i)
	if err != nil {
		return err
	}

	return nil
}

func (r *Rest) Download(p, outputDir string) error {
	url := fmt.Sprintf("%s%s", r.baseURL, p)
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
