package main

import (
	"flag"
	"fmt"
	"os"

	todo "github.com/davidandw190/simple-todo-go"
)

const (
	todoFile = ".todo.json"
)

func main() {

	// short form flags
	add := flag.Bool("a", false, "add a new todo item")
	complete := flag.Int("c", 0, "mark a todo item as completed")
	delete := flag.Int("d", 0, "delete a todo item")

	// long form flags
	flag.BoolVar(add, "add", false, "add a new todo item")
	flag.IntVar(complete, "complete", 0, "mark a todo item as completed")
	flag.IntVar(delete, "del", 0, "delete a todo item")

	flag.Parse()

	todos := &todo.Todos{}

	if err := todos.Load(todoFile); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	switch {
	case *add:
		todos.Add("Sample todo")
		err := todos.Store(todoFile)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}

	case *complete > 0:
		err := todos.Complete(*complete)
		if err != nil {
			fmt.Fprint(os.Stderr, err.Error())
			os.Exit(1)
		}
		err = todos.Store(todoFile)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}

	case *delete > 0:
		err := todos.Delete(*delete)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
		err = todos.Store(todoFile)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
	default:
		fmt.Fprint(os.Stdout, "invalid command")
		os.Exit(0)
	}
}
