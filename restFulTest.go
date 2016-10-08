package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	args := os.Args[1:]

	argsLen := len(args)

	if argsLen > 0 {
		switch strings.ToLower(args[0]) {
		case "get":
			if argsLen == 2 {
				fmt.Println("test")

			} else {
				sendHelp()
			}

		}

	} else {
		sendHelp()
	}

}

func sendHelp() {
	fmt.Println("Run '" + os.Args[0] + " help' to get help")
}
