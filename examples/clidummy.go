/*
clidummy demonstrates the basics of using the gocli library
*/
package main

import (
	"github.com/benthor/gocli"
	"strings"
)

func main() {

	cli := gocli.MkCLI("Welcome to this dummy CLI. Type 'help' to get a list of all available commands")

	// register help Option with stub callback calling cli.Help("")
	cli.AddOption("help", "prints this help message", func(args []string) string { return cli.Help(args) })

	// register exit Option with stub callback calling cli.Exit("bye")
	cli.AddOption("exit", "exits the input loop", func(_ []string) string { return cli.Exit("bye") })

	cli.DefaultOption(func(args []string) string {
		return strings.Join(args, " ")
	})

	cli.Loop("dummyprompt? ")
}
