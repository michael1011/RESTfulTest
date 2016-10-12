package main

import (
	"flag"
	"fmt"
	"github.com/yosssi/gohtml"
	"net/http"
	"strconv"
	"time"
)

func startGui() {
	rawPort := flag.Int("port", 8000, "change to port of the gui")
	flag.Parse()

	port := strconv.Itoa(*rawPort)

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
	startTime := time.Now()

	url := response.URL.Query().Get("url")
	body := response.URL.Query().Get("body")
	header := response.URL.Query().Get("header")

	if url == "" {
		writer.Write([]byte("You have to set an url"))

	} else {
		if body == "" && header == "" {
			resp, err := getRequest(url)

			writeResponse(resp, err, startTime, writer)

		} else {
			// fixme: add post requests
		}
	}

}

func writeResponse(resp *http.Response, err error, startTime time.Time, writer http.ResponseWriter) {
	if err == nil {
		out, resp := parseResponse(resp)

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
