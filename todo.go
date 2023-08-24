// Package todo provides functionality for a simple CLI todo application.
package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"time"

	"github.com/alexeyco/simpletable"
)

// Item represents a single todo item.
type Item struct {
	Task        string
	Done        bool
	CreatedAt   time.Time
	CompletedAt time.Time
}

// Todos is a collection of todo items.
type Todos []Item

// Add adds a new task to the todo list.
func (t *Todos) Add(task string) {
	todo := Item{
		Task:        task,
		Done:        false,
		CreatedAt:   time.Now(),
		CompletedAt: time.Time{},
	}

	*t = append(*t, todo)
}

// Complete marks a task as completed by index.
func (t *Todos) Complete(index int) error {
	if index < 1 || index > len(*t) {
		return errors.New("invalid index")
	}

	item := &(*t)[index-1]
	item.CompletedAt = time.Now()
	item.Done = true

	return nil
}

// Delete removes a task by index.
func (t *Todos) Delete(index int) error {
	if index < 1 || index > len(*t) {
		return errors.New("invalid index")
	}

	*t = append((*t)[:index-1], (*t)[index:]...)

	return nil
}

// Load reads and deserializes todo items from a file.
func (t *Todos) Load(filename string) error {
	file, err := os.ReadFile(filename)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return err
	}

	if len(file) == 0 {
		return nil
	}

	err = json.Unmarshal(file, t)
	if err != nil {
		return err
	}

	return nil
}

// Store writes the todo items to a file in JSON format.
func (t *Todos) Store(filename string) error {
	data, err := json.Marshal(t)
	if err != nil {
		return err
	}

	err = os.WriteFile(filename, data, 0644)
	if err != nil {
		return err
	}

	return nil
}

// Print prints all the todo items to the console.
func (t Todos) Print() {

	clearScreen()
	fmt.Printf("\n\n")
	printCurrentDateTime()

	if len(t) < 1 {
		fmt.Println("(empty)")
	}

	table := simpletable.New()

	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "#"},
			{Align: simpletable.AlignCenter, Text: "Task"},
			{Align: simpletable.AlignCenter, Text: "Status"},
			{Align: simpletable.AlignCenter, Text: "Created At"},
			{Align: simpletable.AlignCenter, Text: "Completed At"},
		},
	}

	var cells [][]*simpletable.Cell

	for index, item := range t {
		index++
		var task string
		var completed string
		if item.Done {
			task = green(fmt.Sprintf("\u2713 %s", item.Task))
			completed = green(fmt.Sprintf("COMPLETED"))
		} else {
			task = blue(fmt.Sprintf("\u2501 %s", item.Task))
			completed = red(fmt.Sprintf("PENDING"))
		}

		cells = append(cells, []*simpletable.Cell{
			{Text: fmt.Sprintf("%d", index)},
			{Text: task},
			{Text: fmt.Sprintf("%s", completed)},
			{Text: item.CreatedAt.Format(time.RFC822)},
			{Text: item.CompletedAt.Format(time.RFC822)},
		})
	}

	table.Body = &simpletable.Body{Cells: cells}

	table.SetStyle(simpletable.StyleUnicode)

	table.Println()

	fmt.Println()

	fmt.Println("\t\t\t" + red(fmt.Sprintf("pending: %d", t.CountPending())) + "\t\t\t" + green(fmt.Sprintf("completed: %d", t.CountCompleted())))

	fmt.Printf("\n\n\n")

}

func (t *Todos) CountPending() int {
	total := 0
	for _, item := range *t {
		if !item.Done {
			total += 1
		}
	}

	return total
}

func (t *Todos) CountCompleted() int {
	total := 0
	for _, item := range *t {
		if item.Done {
			total += 1
		}
	}

	return total
}

func printCurrentDateTime() {
	currentTime := time.Now()
	fmt.Printf("%s\n\n", currentTime.Format(time.RFC1123))
}

func clearScreen() {
	var clearCmd string

	switch runtime.GOOS {
	case "windows":
		clearCmd = "cls"
	default:
		clearCmd = "clear"
	}

	cmd := exec.Command(clearCmd)
	cmd.Stdout = os.Stdout
	cmd.Run()
}
