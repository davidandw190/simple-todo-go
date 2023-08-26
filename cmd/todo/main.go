package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"

	todo "github.com/davidandw190/simple-todo-go"
)

const (
	todoFile = ".todo.json"
)

func main() {
	list := flag.Bool("l", false, "List all todo items")
	add := flag.Bool("a", false, "Add a new todo item")
	edit := flag.Int("e", 0, "Edit existing todo item")
	complete := flag.Int("c", 0, "Mark a todo item as completed")
	delete := flag.Int("d", 0, "Delete a todo item")
	deleteAll := flag.Bool("da", false, "Delete all the existing todo items")

	flag.BoolVar(list, "list", false, "List all todo items")
	flag.BoolVar(add, "add", false, "Add a new todo item")
	flag.IntVar(edit, "edit", 0, "Edit existing todo item")
	flag.IntVar(complete, "complete", 0, "Mark a todo item as completed")
	flag.IntVar(delete, "del", 0, "Delete a todo item")
	flag.BoolVar(deleteAll, "delall", false, "Delete all the existing todo items")

	flag.Parse()

	todos := &todo.Todos{}

	if err := todos.Load(todoFile); err != nil {
		todo.PrintRedStderr("[!] Error loading todo items: " + err.Error())
		os.Exit(1)
	}

	switch {
	case *list:
		todos.Print()

	case *add:
		task := getInput(false, flag.Args()...)
		todos.Add(task)

	case *edit > 0:
		newTask := getInput(true, flag.Args()...)
		err := todos.Edit(*edit, newTask)
		if err != nil {
			todo.PrintRedStderr("[!] Error editing todo item: " + err.Error())
			os.Exit(1)
		}

	case *complete > 0:
		err := todos.Complete(*complete)
		if err != nil {
			todo.PrintRedStderr("[!] Error completing todo item: " + err.Error())
			os.Exit(1)
		}

	case *delete > 0:
		err := todos.Delete(*delete)
		if err != nil {
			todo.PrintRedStderr("[!] Error deleting todo item: " + err.Error())
			os.Exit(1)
		}

	case *deleteAll:
		err := todos.DeleteAll()
		if err != nil {
			todo.PrintRedStderr("[!] Error deleting todo item: " + err.Error())
			os.Exit(1)
		}

	default:
		todo.PrintBlue(os.Stdout, "[?] Invalid command\n")
		os.Exit(0)
	}

	err := todos.Store(todoFile)
	if err != nil {
		todo.PrintRedStderr("[!] Error storing todo items: " + err.Error())
		os.Exit(1)
	}
}

func getInput(editMode bool, args ...string) string {
	if len(args) > 0 {
		return strings.Join(args, " ")
	}

	reader := bufio.NewReader(os.Stdin)

	if editMode {
		fmt.Print("> Enter the modified task: ")
	} else {
		fmt.Print("> Enter new task:")
	}

	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(strings.Replace(input, "\r", "", -1))

	if len(input) == 0 {
		todo.PrintRedStderr("[!] Empty todo is not allowed\n")
		os.Exit(1)
	}

	return input
}
