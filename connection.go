package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func ConnectSQLite(path string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", path)

	if err != nil {
		return nil, err
	}

	return db, nil
}

func GetAllTodos(database *sql.DB) []string {
	rows, err := database.Query("SELECT * FROM todos")
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
	return todos
}

func AddTodos(database *sql.DB, todos []string) error {
	for _, todo := range todos {
		_, err := database.Exec("INSERT INTO todos VALUES(?);", todo)
		if err != nil {
			return err
		}
	}
	return nil
}
