package main

import (
	"fmt"
	"net/http"
)

// fixme: add option to change port
func openGui() {
	fmt.Println("Open 'localhost:8000' in your browser to see the interface")
	fmt.Println("'Ctrl + C' to stop the server")

	http.HandleFunc("/request", request)
	http.Handle("/", http.FileServer(assetFS()))

	err := http.ListenAndServe(":8000", nil)

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
		if body == "" && header == ""{
			resp, err := getRequest(url)

			if err == nil {
				// todo: send time, status code, ...
				out, _ := parseResponse(resp)

				writer.Write([]byte(out))

			} else {
				writer.Write([]byte(err.Error()))
			}
		} else {
			// fixme: add post request
		}
	}

}