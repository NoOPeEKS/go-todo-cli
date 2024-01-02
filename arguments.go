package main

import (
	"log"
	"os"
)

func ParseArguments() (string, []string) {
	// Get every argument except program name
	args := os.Args[1:]

	if args[0] == "add" { // If first argument is add, return add and todos to add
		if len(args) >= 2 {
			return "add", args[1:]
		} else {
			log.Fatal("Not enough arguments. Add subcommand requires at least one name")
			return "error", nil
		}
	} else if args[0] == "list" { // If first and only argument is list, return every todo in db
		if len(args) != 1 {
			log.Fatal("Too much arguments. List does not require subcommands.")
		}
		return "list", nil
	} else if args[0] == "delete" { // If first argument is delete, delete next argument in db
		if len(args) == 2 {
			return "delete", args[1:]
		} else {
			log.Fatal("Incorrect arguments. Need to have only 1 argument to delete")
			return "error", nil
		}
	} else {
		return "error", nil
	}
}
