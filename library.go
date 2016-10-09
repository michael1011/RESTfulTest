package main

import (
	"net/http"
	"strings"
	"encoding/json"
	"bytes"
)

func getRequest(url string) (*http.Response, error) {
	if !strings.HasPrefix(url, "http") {
		url = "http://" + url
	}

	response, err := http.Get(url)

	return response, err
}


func isJson(str string) (bool) {
	var js json.RawMessage
	return json.Unmarshal([]byte(str), &js) == nil
}

func prettyJson(in string) (string, error) {
	var out bytes.Buffer
	err := json.Indent(&out, []byte(in), "", "\t")

	if err != nil {
		return "", err
	}

	return out.String(), err
}
