package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	todo "github.com/davidandw190/simple-todo-go"
)

const (
	todoFile = ".todo.json"
)

func main() {

	// short form flags
	list := flag.Bool("l", false, "List all todo items")
	add := flag.Bool("a", false, "Add a new todo item")
	complete := flag.Int("c", 0, "Mark a todo item as completed")
	delete := flag.Int("d", 0, "Delete a todo item")

	// long form flags
	flag.BoolVar(list, "list", false, "List all todo items")
	flag.BoolVar(add, "add", false, "Add a new todo item")
	flag.IntVar(complete, "complete", 0, "Mark a todo item as completed")
	flag.IntVar(delete, "del", 0, "Delete a todo item")

	flag.Parse()

	todos := &todo.Todos{}

	if err := todos.Load(todoFile); err != nil {
		fmt.Fprintln(os.Stderr, "Error loading todo items:", err)
		os.Exit(1)
	}

	switch {
	case *list:
		todos.Print()

	case *add:
		task, err := getInput(os.Stdin, flag.Args()...)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error getting input:", err)
			os.Exit(1)
		}
		todos.Add(task)

	case *complete > 0:
		err := todos.Complete(*complete)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error completing todo item:", err)
			os.Exit(1)
		}

	case *delete > 0:
		err := todos.Delete(*delete)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error deleting todo item:", err)
			os.Exit(1)
		}

	default:
		fmt.Fprintln(os.Stdout, "Invalid command")
		os.Exit(0)
	}

	err := todos.Store(todoFile)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error storing todo items:", err)
		os.Exit(1)
	}
}

func getInput(r io.Reader, args ...string) (string, error) {
	if len(args) > 0 {
		return strings.Join(args, " "), nil
	}

	scanner := bufio.NewScanner(r)
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		return "", err
	}

	text := scanner.Text()

	if len(text) == 0 {
		return "", errors.New("empty todo is not allowed")
	}

	return text, nil
}
