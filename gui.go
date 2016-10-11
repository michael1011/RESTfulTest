package main

import (
	"fmt"
	"net/http"
)

// fixme: add option to change port
func openGui() {
	fmt.Println("Open 'localhost:8000' in your browser to see the interface")
	fmt.Println("'Ctrl + C' to stop the server")

	http.Handle("/", http.FileServer(assetFS()))
	err := http.ListenAndServe(":8000", nil)

	if err != nil {
		sendError(err)
	}

}
