package main

import (
	"fmt"
	"net/http"
)

func openGui() {
	fmt.Println("Open 'localhost:8080' in your browser to see the interface")
	fmt.Println("'Ctrl + C' to stop the server")

	http.Handle("/", http.FileServer(assetFS()))
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		sendError(err)
	}

}
