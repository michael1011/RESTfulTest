package main

import (
	"flag"
	"fmt"
	"github.com/fatih/color"
	"github.com/yosssi/gohtml"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

var fatal string = "Error: "

// todo: add gui
// todo: option to save response in file

func main() {
	red := color.New(color.FgRed)

	args := os.Args[1:]

	argsLen := len(args)

	if argsLen > 0 {
		if strings.HasPrefix(os.Args[1], "-") {
			args = os.Args[2:]
		}

		switch strings.ToLower(args[0]) {
		case "get":
			if argsLen > 1 {
				startTime := time.Now()

				response, err := getRequest(args[1])

				sendResponse(response, err, startTime)

			} else {
				sendHelp()
			}

		case "post":
			if argsLen > 2 {
				startTime := time.Now()

				body := ""
				headers := make(map[string]string)

				for _, value := range args[2:] {
					switch {
					case strings.HasPrefix(value, "body="):
						body = body + value[5:]

					case strings.HasPrefix(value, "header="):
						for _, value = range strings.Split(value[5:], ";") {
							entry := strings.Split(value, ":")

							headers[entry[0]] = entry[1]
						}

					default:
						red.Println(fatal + "Unknown argument: " + value)
						os.Exit(1)
					}

				}

				response, err := postRequest(args[1], body, headers)

				sendResponse(response, err, startTime)

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

func sendResponse(response *http.Response, err error, startTime time.Time) {
	if err == nil {
		defer response.Body.Close()

		rawResp, readErr := ioutil.ReadAll(response.Body)

		if readErr != nil {
			sendError(readErr)
		}

		fmt.Println("Status: " + response.Status)
		fmt.Println("Time: ", time.Since(startTime))
		fmt.Println("Headers: ", response.Header)
		fmt.Println()

		readResp := string(rawResp)

		beautify := flag.Bool("beautify", true, "disable beautifying")

		flag.Parse()

		if *beautify {
			if isJson(readResp) {
				jsonRead, jsonErr := prettyJson(readResp)

				if jsonErr == nil {
					fmt.Println(jsonRead)
				} else {
					sendError(jsonErr)
				}

			} else {
				fmt.Println(gohtml.Format(readResp))
			}

		} else {
			fmt.Println(readResp)
		}

	} else {
		sendError(err)
	}
}

func sendError(error error) {
	red := color.New(color.FgRed)

	red.Print(fatal)
	red.Println(error)
}

func sendHelp() {
	fmt.Println("Run '" + os.Args[0] + " help' to get help")
}
