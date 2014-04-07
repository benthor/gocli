/*
Package gocli implements a simple command line interface (cli) library

It is a higher-level wrapper around the line editor Liner (https://github.com/peterh/liner),
mainly adding convenience functions for

 * registering commands
 * automatic history
 * tab completion
 * running a REPL-ish loop

A few (currently) built-in defaults:

 * whitespaces in commands are not allowed
 * the tab completer will complete non-ambiguous but incomplete commands while preserving all arguments
 * options registered with empty help message won't show up in help list


(May or may not be idiomatic Go, the author is just a pedestrian hacker messing around with an interesting new programming language)
*/
package gocli

import (
	"errors"
	"fmt"
	"github.com/peterh/liner"
	"strings"
)

// type Option may or may not contain accessible attributes in the future
type Option struct {
	Cmd      string
	Help     string
	Function func(args []string) string
}

// type CLI may or may not contain accessible attributes in the future
type CLI struct {
	Liner    liner.State
	Options  map[string]Option
	Default  Option
	Greeting string
	looping  bool
	longest  int
}

// AddOption registers a command (cmd), appropriate documentation string (help) and callback function with the CLI
// Also registers cmd with the in-built Tab completer
// Returns an error if cmd string contains white spaces
func (cli *CLI) AddOption(cmd string, help string, function func(args []string) string) error {
	if strings.Count(cmd, " ") > 0 {
		return errors.New("cmd string can not contain white spaces")
	}
	cli.Options[cmd] = Option{cmd, help, function}
	// need this for pretty printing the help message
	if cli.longest < len(cmd) {
		cli.longest = len(cmd)
	}
	return nil
}

// Register callback to process the (white-space split) cmdline that could not be parsed
func (cli *CLI) DefaultOption(function func(args []string) string) {
	cli.Default = Option{"", "", function}
}

// Loop is a REPL-inspired loop, prompting for input and running the registered callbacks
func (cli *CLI) Loop(prompt string) {
	cli.looping = true
	fmt.Println(cli.Greeting)
	for cli.looping {
		cmd, err := cli.Liner.Prompt(prompt)
		if err != nil {
			// l.Println(err)
			cli.Exit([]string{fmt.Sprintf("error: %q", err.Error())})
		} else {
			tmp := strings.Split(cmd, " ")
			if option, ok := cli.Options[tmp[0]]; ok {
				cli.Liner.AppendHistory(cmd)
				fmt.Println(option.Function(tmp[1:]))
			} else {
				fmt.Println(cli.Default.Function(tmp))
			}
		}
	}
}

// Help returns a documentation string about all registered Options.
// It is meant to be used as the callback of a registered "help" Option
func (cli *CLI) Help(args []string) string {
	var result string
	for cmd, option := range cli.Options {
		if len(option.Help) > 0 {
			fmt.Printf("%"+fmt.Sprintf("%d", cli.longest)+"s  -  %s\n", cmd, option.Help)
		}
	}
	return result
}

// Exit terminates the loop, returning a concatenation of all string arguments
// (The signature purposefully matches that of an Option callback to be easily called as such)
func (cli *CLI) Exit(args []string) string {
	cli.looping = false
	cli.Liner.Close()
	return strings.Join(args, " ")
}

// MkCLI returns new CLI
func MkCLI(greeting string) CLI {
	tmp := CLI{*liner.NewLiner(), make(map[string]Option), Option{}, greeting, true, 0}
	tmp.Liner.SetCompleter(func(line string) []string {
		tokens := strings.Split(line, " ")
		// first word is already a valid command
		if _, ok := tmp.Options[tokens[0]]; ok {
			// add whitespace to indicate validity of command by jumping cursor
			return []string{line + " "}
		}
		candidates := []string{}
		filtered := []string{}
		for candidate, _ := range tmp.Options {
			// only do prefix here
			if strings.HasPrefix(candidate, tokens[0]) {
				// make sure that any arguments are carried through the tab completion
				candidates = append(candidates, candidate+" "+strings.Join(tokens[1:], " "))
			} else {
				filtered = append(filtered, candidate)
			}
		}
		// test for substring in the rest
		// TODO: could also disallow matching of single (or just a few) letters by doing a length check on tokens[0] here
		for _, candidate := range filtered {
			if strings.Contains(candidate, tokens[0]) {
				candidates = append(candidates, candidate+" "+strings.Join(tokens[1:], " "))
			}
		}
		return candidates
	})
	return tmp
}
