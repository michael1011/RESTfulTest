package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strings"
)

func getRequest(url string) (*http.Response, error) {
	if !strings.HasPrefix(url, "http") {
		url = "http://" + url
	}

	response, err := http.Get(url)

	return response, err
}

func isJson(input string) bool {
	var js json.RawMessage
	return json.Unmarshal([]byte(input), &js) == nil
}

func prettyJson(input string) (string, error) {
	var out bytes.Buffer
	err := json.Indent(&out, []byte(input), "", "\t")

	if err != nil {
		return "", err
	}

	return out.String(), err
}
