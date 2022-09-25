package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/jrvldam/command-line-applications/todo"
)

const todoFilename = ".todo.json"

func main() {
	task := flag.String("task", "", "Taks to be included in the ToDo list")
	list := flag.Bool("list", false, "List all tasks")
	complete := flag.Int("complete", 0, "Task to be completed")
	flag.Parse()

	l := &todo.List{}

	if err := l.Get(todoFilename); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	switch {
	case *list:
		for _, todo := range *l {
			if !todo.Done {
				fmt.Println(todo.Task)
			}
		}
	case *complete > 0:
		if err := l.Complete(*complete); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		if err := l.Save(todoFilename); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		fmt.Println("list: v%", l)
	case *task != "":
		l.Add(*task)

		if err := l.Save(todoFilename); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	default:
		fmt.Fprintln(os.Stderr, "Invalid option")
		os.Exit(1)
	}
}
