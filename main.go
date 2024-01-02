package main

import (
	"fmt"
	"log"
)

func main() {

	if !DatabaseFileExists() {
		CreateDatabase()
		db, err := ConnectSQLite("./todos.db")
		if err != nil {
			log.Fatal("Could not connect to database after creating it.")
		}
		_, error := db.Exec("CREATE TABLE todos (name primary key);")
		if error != nil {
			log.Fatal("Could not create todos table when creating database")
		}
	}
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
		Added(arguments)
	} else if mode == "list" {
		todos := GetAllTodos(db)
		Show(todos)
	} else if mode == "delete" {
		err := DeleteTodo(db, arguments[0])
		if err != nil {
			log.Fatal("Could not delete todo correctly")
		}
		Deleted(arguments[0])
	} else if mode == "help" {
		help_text := "Usage: \n add <task_name_1> <task_name_2> ...\n list (Display all pending tasks)\n delete <task_name> (Remove a task from the list)"
		fmt.Println(help_text)
	}

	// Close the database connection
	db.Close()
}
