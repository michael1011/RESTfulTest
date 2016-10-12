package main

import (
	"fmt"
	"os"
)

func sendInstructions() {
	fmt.Println("RESTfulTEST is an application to test RESTful services")
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println("	" + os.Args[0] + " <flags> command [arguments]")
	fmt.Println()
	fmt.Println("The commands are:")
	fmt.Println("	get <url>         send a get request to <url>")
	fmt.Println()
	fmt.Println("	post <url> body=\"...\" header=\"key:value;key:value;...\"")
	fmt.Println("			  send a post request to <url>")
	fmt.Println()
	fmt.Println("	version           get the version number")
	fmt.Println()
	fmt.Println("	gui               to start the gui")
	fmt.Println()
	fmt.Println("The flags are:")
	fmt.Println("	-beautify=false   disable beautifying responses")
	fmt.Println("	-beautify=true    enable beautifying responses (default)")
	fmt.Println()
	fmt.Println("	-port=<port>      set the port of the gui (default: 8000)")
	fmt.Println()
}
