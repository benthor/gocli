/*
Package gocli implements a simple command line interface (cli) library

It is a higher-level wrapper around the line editor Liner (https://github.com/peterh/liner),
mainly adding convenience functions for

 - registering commands
 - automatic history
 - tab completion
 - running a REPL-ish loop

(May or may not be idiomatic Go, the author is just a pedestrian hacker messing around with an interesting new programming language)
*/

package gocli

import (
	"errors"
	"fmt"
	"github.com/peterh/liner"
	"strings"
)

type Option struct {
	Cmd      string
	Help     string
	Function func(args []string) string
}

type CLI struct {
	Liner     liner.State
	Options   map[string]Option
	Default   Option
	exit_chan chan int
}

// AddOption registers a command (cmd), appropriate documentation string (help) and callback function with the CLI
// Returns an error if cmd string contains white spaces
func (cli *CLI) AddOption(cmd string, help string, function func(args []string) string) error {
	if strings.Count(cmd, " ") > 0 {
		return errors.New("cmd string can not contain white spaces")
	}
	cli.Options[cmd] = Option{cmd, help, function}
	return nil
}

// Register callback to process the (white-space split) cmdline that could not be parsed
func (cli *CLI) DefaultOption(function func(args []string) string) {
	cli.Default = Option{"", "", function}
}

// Loop is a REPL-inspired loop, prompting for input and running the registered callbacks
func (cli *CLI) Loop(prompt string) {
	inner := func() {
		for {
			cmd, err := cli.Liner.Prompt(prompt)
			if err != nil {
				// l.Println(err)
				cli.Exit(fmt.Sprintf("error: %q", err.Error()))
			} else {
				tmp := strings.Split(cmd, " ")
				if option, ok := cli.Options[tmp[0]]; ok {
					fmt.Println(option.Function(tmp[1:]))
				} else {
					fmt.Println(cli.Default.Function(tmp))
				}

			}
		}
	}
	go inner()
	<-cli.exit_chan
}

// Help prints documentation about all registered Options
func (cli *CLI) Help(args []string) string {
	return fmt.Sprintf("help message for '%s' not yet implemented", strings.Join(args, " "))
}

// Exit terminates the loop, returning the specified message
func (cli *CLI) Exit(message string) string {
	cli.Liner.Close()
	cli.exit_chan <- 1
	return message
}

func MkCLI() CLI {
	return CLI{*liner.NewLiner(), make(map[string]Option), Option{}, make(chan int)}
}
