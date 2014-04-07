/*
clidummy demonstrates the basics of using the gocli library
*/
package main

import (
	"fmt"
	"github.com/benthor/gocli"
	"strings"
)

func main() {

	cli := gocli.MkCLI("Welcome to this dummy CLI. Type 'help' to get a list of all available commands")

	// register help Option with stub callback calling cli.Help("")
	cli.AddOption("help", "prints this help message", func(args []string) string { return cli.Help(args) })

	// register exit Option with stub callback calling cli.Exit with any further provided tokens as argument
	cli.AddOption("exit", "exits the input loop", func(args []string) string { return cli.Exit(strings.Join(args, " ")) })

	// register exit Option with stub callback calling cli.Exit with any further provided tokens as argument
	cli.AddOption("kapitänsmützenabzeichen", "just an example of a long cmd name not breaking the help formatting", func(args []string) string { return cli.Exit(strings.Join(args, " ")) })

	cli.DefaultOption(func(args []string) string {
		return strings.Join(args, " ")
	})

	// run the main loop
	cli.Loop("dummyprompt? ")

	fmt.Println("this part of the code is only reached when the cli loop returns")
}
