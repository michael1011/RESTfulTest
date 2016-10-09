package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strings"
)

func getRequest(url string) (*http.Response, error) {
	request, err := http.Get(completeUrl(url))

	return request, err
}

func postRequest(url string, body string, headers map[string]string) (*http.Response, error) {
	request, _ := http.NewRequest("POST", completeUrl(url), bytes.NewBuffer([]byte(body)))

	for key, value := range headers {
		request.Header.Add(key, value)
	}

	response, err := http.DefaultClient.Do(request)

	return response, err
}

func completeUrl(url string) string {
	if !strings.HasPrefix(url, "http") {
		url = "http://" + url
	}

	return url
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
