package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/jrvldam/command-line-applications/todo"
)

const todoFilename = ".todo.json"

func main() {
	l := &todo.List{}

	if err := l.Get(todoFilename); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	switch {
	case len(os.Args) == 1:
		for _, todo := range *l {
			fmt.Println(todo.Task)
		}
	default:
		todo := strings.Join(os.Args[1:], " ")
		l.Add(todo)
		if err := l.Save(todoFilename); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}
}
