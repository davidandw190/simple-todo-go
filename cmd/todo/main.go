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

	add := flag.Bool("a", false, "add a new todo item")

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
	default:
		fmt.Fprint(os.Stdout, "invalid command")
		os.Exit(0)
	}
}
