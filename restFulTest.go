package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/fatih/color"
	"github.com/yosssi/gohtml"
	"net/http"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"
)

const fatal string = "Error: "

const version string = "1.0.0"

func main() {
	red := color.New(color.FgRed)

	args := os.Args[1:]

	argsLen := len(args)

	if argsLen > 0 {
		for i := 1; strings.HasPrefix(os.Args[i], "-"); i++ {
			args = os.Args[i+1:]
		}

		switch strings.ToLower(args[0]) {
		case "get":
			if argsLen > 1 {
				startTime := time.Now()

				response, err := getRequest(args[1])

				if err == nil {
					output, response := parseResponse(response)

					printParsed(output, response, startTime)

				} else {
					sendError(err)
				}

			} else {
				sendHelp()
			}

		case "post":
			if argsLen > 2 {
				body := ""
				headers := make(map[string]string)

				for _, value := range args[2:] {
					switch {
					case strings.HasPrefix(value, "body="):
						body = body + value[5:]

					case strings.HasPrefix(value, "header="):
						for _, value = range strings.Split(value[5:], "::") {
							entry := strings.Split(value, ":")

							headers[entry[0]] = entry[1]
						}

					default:
						red.Println(fatal + "Unknown argument: " + value)
						os.Exit(1)
					}

				}

				startTime := time.Now()

				response, err := postRequest(args[1], body, headers)

				if err == nil {
					output, response := parseResponse(response)

					flag.Parse()

					printParsed(output, response, startTime)

				} else {
					sendError(err)
				}

			} else {
				sendHelp()
			}

		case "gui":
			rawPort := flag.Int("port", 8000, "change to port of the gui")
			flag.Parse()

			startGui(strconv.Itoa(*rawPort))

		case "version":
			build, _ := Asset("public/build.txt")

			fmt.Println("RESTfulTest " + version + "-" + strings.Replace(string(build[:]), "\n", "", -1) +
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

func printParsed(output string, response *http.Response, startTime time.Time) {
	save := flag.String("save", "", "to save the output to a file")
	beautify := flag.Bool("beautify", true, "disable beautifying")

	flag.Parse()

	timeString := time.Since(startTime).String()

	if *save != "" {
		saveFile(*save, response, output, timeString, *beautify)
	}

	fmt.Println(outputTemplate[0], response.Status)
	fmt.Println(outputTemplate[1], timeString)
	fmt.Println()
	fmt.Println(outputTemplate[2], response.Header)
	fmt.Println()
	fmt.Println()

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

func saveFile(save string, response *http.Response, output string, time string, beautify bool) {
	f, _ := os.Create(save)

	writer := bufio.NewWriter(f)

	writer.WriteString(outputTemplate[0] + response.Status + "\n")
	writer.WriteString(outputTemplate[1] + time + "\n\n")

	writer.WriteString(outputTemplate[2] + "\n")
	response.Header.Write(writer)
	writer.WriteString("\n\n")

	if beautify {
		if isJson(output) {
			writer.WriteString(prettyJson(output))

		} else {
			writer.WriteString(gohtml.Format(output))
		}

	} else {
		writer.WriteString(output)
	}

	writer.Flush()
}

func sendError(error error) {
	red := color.New(color.FgRed)

	red.Print(fatal)
	red.Println(error)
}

func sendHelp() {
	fmt.Println("Run '" + os.Args[0] + " help' to get help")
}
