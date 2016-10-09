package main

import (
	"io/ioutil"
	"os"
	"strings"
	"fmt"
)

var fatal string = "Error: "

func main() {
	args := os.Args[1:]

	argsLen := len(args)

	if argsLen > 0 {
		switch strings.ToLower(args[0]) {
		case "get":
			if argsLen == 2 {
				response, err := getRequest(args[1])

				if err == nil {
					defer response.Body.Close()

					rawResp, readErr := ioutil.ReadAll(response.Body)

					if readErr != nil {
						sendError(readErr)
					}

					fmt.Println("Status: "+response.Status)
					fmt.Println()

					readResp := string(rawResp)

					if isJson(readResp) {
						jsonRead, jsonErr := prettyJson(readResp)

						if jsonErr == nil {
							fmt.Println(jsonRead)
						} else {
							sendError(jsonErr)
						}

					} else {
						fmt.Println(readResp)
					}

				} else {
					sendError(err)
				}

			} else {
				sendHelp()
			}

		default:
			sendHelp()
		}

	} else {
		sendHelp()
	}

}

func sendError(error error) {
	fmt.Print(fatal)
	fmt.Println(error)
}

func sendHelp() {
	fmt.Println("Run '" + os.Args[0] + " help' to get help")
}