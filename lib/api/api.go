package api

import (
	"bytes"
	"net/http"
)

const (
	Get = iota
	Post
)

func Request(url string, payload []byte, method int, headers map[string]string) (*http.Response, error) {
	req, err := http.NewRequest(
		methodToString(method),
		url,
		bytes.NewBuffer(payload),
	)
	if err != nil {
		return nil, err
	}

	setHeader(req, headers)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func methodToString(method int) string {
	switch method {
	case Get:
		return "GET"
	case Post:
		return "POST"
	}
	return ""
}

func setHeader(req *http.Request, headers map[string]string) {
	_, ok := headers["Content-Type"]
	if !ok {
		headers["Content-Type"] = "application/x-www-form-urlencoded"
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}
}
