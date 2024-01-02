package main

import (
	"fmt"
)

func main() {
	// Connect to the database and get a DB object
	db, err := ConnectSQLite("./todos.db")

	if err != nil {
		panic("Couldnt connect to database")
	}

	// Parse the program arguments to select the mode and subarguments
	mode, arguments := ParseArguments()

	// Depending on the mode, do one thing or the other
	if mode == "error" {
		panic("Could not parse arguments correctly")
	}
	if mode == "add" {
		err := AddTodos(db, arguments)
		if err != nil {
			panic("Could not insert todos correctly")
		}
	} else if mode == "list" {
		todos := GetAllTodos(db)
		for i := 0; i < len(todos); i++ {
			fmt.Println(todos[i])
		}
	} else if mode == "delete" {
		panic("Not implemented yet!")
	}

	// Close the database connection
	db.Close()
}
