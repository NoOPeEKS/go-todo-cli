package main

import (
	"fmt"
)

func main() {
	db, err := ConnectSQLite("./todos.db")

	if err != nil {
		panic("Couldnt connect to database")
	}

	mode, _ := ParseArguments()
	if mode == "error" {
		panic("Could not parse arguments correctly")
	}

	if mode == "add" {
		panic("Not implemented yet!")
	} else if mode == "list" {
		todos := GetAllTodos(db)
		for i := 0; i < len(todos); i++ {
			fmt.Println(todos[i])
		}
	} else if mode == "delete" {
		panic("Not implemented yet!")
	}

	db.Close()
}
