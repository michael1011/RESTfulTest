package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strings"
	"time"
	"flag"
	"io/ioutil"
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

func parseResponse(response *http.Response, err error, startTime time.Time) (output [5]string, json bool) {
	if err == nil {
		defer response.Body.Close()

		rawResp, readErr := ioutil.ReadAll(response.Body)

		if readErr != nil {
			sendError(readErr)
		}

		output[0] = "Status: " + response.Status

		// fixme: fix headers and time
		output[1] = "Time: "//string(time.Since(startTime))
		output[2] = "Headers: "//+string(response.Header)
		output[3] = ""

		readResp := string(rawResp)

		beautify := flag.Bool("beautify", true, "disable beautifying")

		flag.Parse()

		if *beautify {
			if isJson(readResp) {
				json = true
			}

			output[4] = readResp

		} else {
			output[4] = readResp
		}

	} else {
		sendError(err)
	}

	return output, json
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
