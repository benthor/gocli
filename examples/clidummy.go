/*
clidummy demonstrates the basics of using the gocli library
*/
package main

import (
	"fmt"
	"github.com/benthor/gocli"
)

func main() {

	cli := gocli.MkCLI("Welcome to this dummy CLI.")

	// register help Option with cli.Help as callback
	cli.AddOption("help", "prints this help message", cli.Help)

	// register exit Option with cli.Exit as callback
	cli.AddOption("exit", "exits the input loop", cli.Exit)

	// register rather long command as Option with custom callback function
	cli.AddOption("kapitänsmützenabzeichen", "just an example of a long cmd name not breaking the help formatting", func(_ []string) string { return cli.Exit([]string{"fnord"}) })

	// register hidden quit Option with cli.Exit as callback. Should not appear in "help" list
	cli.AddOption("quit", "", cli.Exit)

	cli.DefaultOption(func(args []string) string {
		return fmt.Sprintf("%s: command not found, type 'help' for help", args[0])
	})

	// run the main loop
	cli.Loop("dummyprompt? ")

	fmt.Println("(this part of the code is only reached when the cli loop returns)")

}
