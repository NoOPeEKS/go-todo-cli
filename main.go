package main

import (
	"fmt"
	"log"
)

func main() {
	// Connect to the database and get a DB object
	db, err := ConnectSQLite("./todos.db")

	if err != nil {
		log.Fatal("Couldnt connect to database")
	}

	// Parse the program arguments to select the mode and subarguments
	mode, arguments := ParseArguments()

	// Depending on the mode, do one thing or the other
	if mode == "error" {
		log.Fatal("Could not parse arguments correctly")
	}
	if mode == "add" {
		err := AddTodos(db, arguments)
		if err != nil {
			log.Fatal("Could not insert todos correctly")
		}
	} else if mode == "list" {
		todos := GetAllTodos(db)
		for i := 0; i < len(todos); i++ {
			fmt.Println(todos[i])
		}
	} else if mode == "delete" {
		err := DeleteTodo(db, arguments[0])
		if err != nil {
			panic("Could not delete todo correctly")
		}
	}

	// Close the database connection
	db.Close()
}
