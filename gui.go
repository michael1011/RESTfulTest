package main

import (
	"fmt"
	"github.com/yosssi/gohtml"
	"net/http"
	"strings"
	"time"
)

func startGui(port string) {
	fmt.Println("Open 'localhost:" + port + "' in your browser to see the interface")
	fmt.Println("'Ctrl + C' to stop the server")

	http.HandleFunc("/request", request)
	http.Handle("/", http.FileServer(assetFS()))

	err := http.ListenAndServe(":"+port, nil)

	if err != nil {
		sendError(err)
	}

}

func request(writer http.ResponseWriter, response *http.Request) {
	url := response.URL.Query().Get("url")
	body := response.URL.Query().Get("body")
	rawHeaders := response.URL.Query().Get("headers")

	if url == "" {
		writer.Write([]byte("You have to set an url"))

	} else {
		startTime := time.Now()

		headersLen := len(rawHeaders)

		if len(body) == 0 && headersLen == 0 {
			resp, err := getRequest(url)

			writeResponse(resp, err, startTime, writer)
		} else {
			headers := make(map[string]string)

			if headersLen > 0 {
				for _, value := range strings.Split(rawHeaders, "::") {
					entry := strings.Split(value, ":")

					headers[entry[0]] = entry[1]
				}
			}

			resp, err := postRequest(url, body, headers)

			writeResponse(resp, err, startTime, writer)
		}
	}

}

func writeResponse(response *http.Response, err error, startTime time.Time, writer http.ResponseWriter) {
	if err == nil {
		out, resp := parseResponse(response)

		writer.Write([]byte(outputTemplate[0] + resp.Status + "\n"))

		writer.Write([]byte(outputTemplate[1] + time.Since(startTime).String() + "\n\n"))

		writer.Write([]byte(outputTemplate[2] + "\n"))
		resp.Header.Write(writer)
		writer.Write([]byte("\n\n"))

		if isJson(out) {
			out = prettyJson(out)
		} else {
			out = gohtml.Format(out)
		}

		writer.Write([]byte(out))

	} else {
		writer.Write([]byte(err.Error()))
	}

}
