package main

import (
	"flag"
	"fmt"
	"net/http"
	"strconv"
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
	url := response.URL.Query().Get("url")
	body := response.URL.Query().Get("body")
	header := response.URL.Query().Get("header")

	if url == "" {
		writer.Write([]byte("You have to set an url"))

	} else {
		if body == "" && header == "" {
			resp, err := getRequest(url)

			if err == nil {
				// todo: send time, status code, ...
				out, _ := parseResponse(resp)

				writer.Write([]byte(out))

			} else {
				writer.Write([]byte(err.Error()))
			}
		} else {
			// fixme: add post requests
		}
	}

}
