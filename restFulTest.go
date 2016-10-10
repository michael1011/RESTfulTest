package main

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/yosssi/gohtml"
	"os"
	"runtime"
	"strings"
	"time"
)

var fatal string = "Error: "

var version string = "1.0.0"

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

				parsed, json := parseResponse(response, err, startTime)

				printParsed(parsed, json)

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

				parsed, json := parseResponse(response, err, startTime)

				printParsed(parsed, json)

			} else {
				sendHelp()
			}

		case "gui":
			openGui()

		case "version":
			build, err := Asset("build.txt")

			if err != nil {
				build = []byte("-dev")
			}

			fmt.Println("RESTfulTest " + version + strings.Replace(string(build), "\n", "", -1) +
				" (Go runtime " + runtime.Version() + ")")

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

func printParsed(parsed [5]string, json bool) {
	fmt.Println(parsed[0])
	fmt.Println(parsed[1])
	fmt.Println(parsed[2])
	fmt.Println(parsed[3])

	if json {
		// fixme: remove <nil> at the end
		fmt.Println(prettyJson(parsed[4]))

	} else {
		fmt.Println(gohtml.Format(parsed[4]))
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
