package main

import (
	"os"
)

func ParseArguments() (string, []string) {
	// Get every argument except program name
	args := os.Args[1:]

	if args[0] == "add" { // If first argument is add, return add and todos to add
		if len(args) >= 2 {
			return "add", args[1:]
		} else {
			panic("Not enough arguments. Add subcommand requires at least one name")
		}
	} else if args[0] == "list" { // If first and only argument is list, return every todo in db
		if len(args) != 1 {
			panic("Too much arguments. List does not require subcommands.")
		}
		return "list", nil
	} else if args[0] == "delete" { // TODO
		panic("Not implemented yet")
	} else {
		return "error", nil
	}
}
