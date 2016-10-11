package main

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/yosssi/gohtml"
	"net/http"
	"os"
	"runtime"
	"strings"
	"time"
)

var fatal string = "Error: "

var version string = "1.0.0"

// fixme: ci build bindata-assetfs
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

				if err == nil {
					output, response, beautify := parseResponse(response)

					printParsed(output, response, beautify, startTime)

				} else {
					sendError(err)
				}

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

				if err == nil {
					output, response, beautify := parseResponse(response)

					printParsed(output, response, beautify, startTime)

				} else {
					sendError(err)
				}

			} else {
				sendHelp()
			}

		case "gui":
			openGui()

		case "version":
			// todo: add build number
			fmt.Println("RESTfulTest " + version+"(Go runtime " + runtime.Version() + ")")
			fmt.Println("Copyright (c) 2016, michael1011")

		case "help":
			sendInstructions()

		default:
			sendHelp()
		}

	} else {
		sendHelp()
	}

}

func printParsed(output string, response *http.Response, beautify *bool, startTime time.Time) {
	fmt.Println(outputTemplate[0], response.Status)
	fmt.Println(outputTemplate[1], time.Since(startTime))
	fmt.Println(outputTemplate[2], response.Header)
	fmt.Println(outputTemplate[3])

	if *beautify {
		if isJson(output) {
			fmt.Println(prettyJson(output))

		} else {
			fmt.Println(gohtml.Format(output))
		}

	} else {
		fmt.Println(output)
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
