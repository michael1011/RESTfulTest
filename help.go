package main

import (
	"fmt"
	"os"
)

func sendInstructions() {
	fmt.Println("RESTfulTEST is an application to test RESTful services")
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println("	"+os.Args[0]+" <flags> command [arguments]")
	fmt.Println()
	fmt.Println("The commands are:")
	fmt.Println("	get <url>         send a get request to <url>")
	fmt.Println()
	fmt.Println("Available flags are:")
	fmt.Println("	-beautify=false   disable beautifying json")
	fmt.Println("	-beautify=true    enable beautifying json (default)")
	fmt.Println()
}