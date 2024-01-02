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
	todos := GetAllTodos()
	for i := 0; i < len(todos); i++ {
		fmt.Println(todos[i])
	}
}
