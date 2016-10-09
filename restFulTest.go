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

					readResp, readErr := ioutil.ReadAll(response.Body)

					if readErr != nil {
						fmt.Print(fatal)
						fmt.Println(readErr)
					}

					fmt.Println("Status: "+response.Status)
					fmt.Println()
					fmt.Println(string(readResp))

				} else {
					fmt.Print(fatal)
					fmt.Println(err)
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

func sendHelp() {
	fmt.Println("Run '" + os.Args[0] + " help' to get help")
}
