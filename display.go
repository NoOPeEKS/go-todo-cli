package main

import (
	"github.com/fatih/color"
)

func Deleted(name string) {
	color.Green("Todo item '%s': deleted correctly.", name)
}

func Show(names []string) {
	for i := 0; i < len(names); i++ {
		color.HiYellow(names[i])
	}
}

func Added(names []string) {
	for i := 0; i < len(names); i++ {
		color.Green("Todo item '%s' added.", names[i])
	}
}
