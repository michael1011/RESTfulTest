package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"io/ioutil"
	"net/http"
	"strings"
)

var outputTemplate []string = []string{"Status: ", "Time: ", "Headers: ", ""}

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

func parseResponse(response *http.Response) (string, *http.Response, *bool) {
	defer response.Body.Close()

	rawResp, readErr := ioutil.ReadAll(response.Body)

	if readErr != nil {
		sendError(readErr)
	}

	readResp := string(rawResp)

	beautify := flag.Bool("beautify", true, "disable beautifying")

	flag.Parse()

	return readResp, response, beautify
}

func isJson(input string) bool {
	var js json.RawMessage
	return json.Unmarshal([]byte(input), &js) == nil
}

func prettyJson(input string) string {
	var out bytes.Buffer
	err := json.Indent(&out, []byte(input), "", "\t")

	if err != nil {
		sendError(err)

		return ""
	}

	return out.String()
}
