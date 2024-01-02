package main

import (
	"fmt"
)

func main() {
	/*
		args := os.Args[1:]
		switch args[0] {
		case "add":
			fmt.Println("Added input")

		default:
			fmt.Println("Default case")
		}*/
	db, err := ConnectSQLite("./todos.db")

	if err != nil {
		panic("Couldnt connect to database")
	}

	todos := GetAllTodos(db)

	for i := 0; i < len(todos); i++ {
		fmt.Println(todos[i])
	}
	db.Close()
}
