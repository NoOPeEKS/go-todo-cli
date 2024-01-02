package main

import (
	"os"
)

func ParseArguments() (string, []string) {
	args := os.Args[1:]
	if args[0] == "add" {
		if len(args) >= 2 {
			return "add", args[1:]
		} else {
			panic("Not enough arguments. Add subcommand requires at least one name")
		}
	} else if args[0] == "list" {
		if len(args) >= 2 {
			panic("Too much arguments. List does not require subcommands.")
		}
		return "list", nil
	} else if args[0] == "delete" {
		panic("Not implemented yet")
	} else {
		return "error", nil
	}
}
