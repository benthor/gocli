[Archived] gocli
=====

Simple, pure Go command line interface (cli) library

It wraps the (pure Go) line editor **Liner** (peterh/liner) and adds:

 * command registration with callbacks and short description
 * automatic history
 * automatic tab completion
 * running a REPL-ish loop

A few (currently) built-in defaults:

 * whitespaces in commands are not allowed
 * the tab completer will complete non-ambiguous but incomplete commands while preserving all arguments
 * options registered with empty help message won't show up in help list

Roadmapped for after the initial release:

 * nested contexts


_(May or may not be idiomatic Go, the author is just a pedestrian hacker messing around with an interesting new programming language)_
