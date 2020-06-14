package remote

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

func do(method string, url string, content interface{}) (io.ReadCloser, error) {

	reqBody, err := json.Marshal(content)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp.Body, err
}

func Get(url string) (io.ReadCloser, error) {
	return do(http.MethodGet, url, nil)
}
