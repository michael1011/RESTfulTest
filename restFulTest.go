package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"flag"
)

var fatal string = "Error: "

// todo: add gui

func main() {
	args := os.Args[1:]

	if(strings.HasPrefix(os.Args[1], "-")) {
		args = os.Args[2:]
	}

	argsLen := len(args)

	if argsLen > 0 {
		switch strings.ToLower(args[0]) {
		// fixme: add post request

		case "get":
			if argsLen > 1 {
				response, err := getRequest(args[1])

				if err == nil {
					defer response.Body.Close()

					rawResp, readErr := ioutil.ReadAll(response.Body)

					if readErr != nil {
						sendError(readErr)
					}

					fmt.Println("Status: " + response.Status)
					fmt.Println()

					readResp := string(rawResp)


					beautify := flag.Bool("beautify", true, "disable beautifying json")

					flag.Parse()

					// fixme: add xml beautifier

					if *beautify {
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
						fmt.Println(readResp)
					}

				} else {
					sendError(err)
				}

			} else {
				sendHelp()
			}

		case "help":
			sendInstructions()

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
