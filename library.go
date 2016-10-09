package main

import (
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
