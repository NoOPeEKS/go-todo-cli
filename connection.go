package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func GetAllTodos() []string {
	db, err := sql.Open("sqlite3", "./todos.db")
	if err != nil {
		fmt.Println("Could not open sqlite3 db")
	}
	rows, err := db.Query("SELECT * FROM todos")
	if err != nil {
		fmt.Println("Could not query correctly")
	}
	var todos []string
	for rows.Next() {
		var todo string
		err = rows.Scan(&todo)
		if err != nil {
			fmt.Println("could not scan row")
		}
		todos = append(todos, todo)
	}
	rows.Close()
	db.Close()
	return todos
}
