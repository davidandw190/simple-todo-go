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

const layoutOld string = " 2 Aug - 15:04"
const layoutToday string = "Today - 15:04"

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
		Task:      task,
		Done:      false,
		CreatedAt: time.Now(),
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

// DeleteAll removes all the existing  tasks.
func (t *Todos) DeleteAll() error {
	if len(*t) == 0 {
		return errors.New("todo list already empty")
	}

	*t = []Item{}

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

	table := createTable(t)

	table.Println()

	fmt.Println()

}

func createTable(todos Todos) *simpletable.Table {
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

	for index, item := range todos {
		index++
		task := formatTask(item)
		status := formatStatus(item.Done)
		createdAt := formatTimestamp(item.CreatedAt)
		completedAt := formatTimestamp(item.CompletedAt)

		row := []*simpletable.Cell{
			{Text: fmt.Sprintf("%d", index)},
			{Text: task},
			{Align: simpletable.AlignCenter, Text: status},
			{Align: simpletable.AlignLeft, Text: createdAt},
			{Align: simpletable.AlignCenter, Text: completedAt},
		}

		table.Body.Cells = append(table.Body.Cells, row)
	}

	table.SetStyle(simpletable.StyleUnicode)
	return table
}

func formatTask(item Item) string {
	if item.Done {
		return Green(fmt.Sprintf("\u2713 %s", item.Task))
	}
	return Blue(fmt.Sprintf("\u2501 %s", item.Task))
}

func formatStatus(done bool) string {
	if done {
		return Green("COMPLETED")
	}
	return Red("...")
}

func (t Todos) countPending() int {
	total := 0
	for _, item := range t {
		if !item.Done {
			total++
		}
	}

	return total
}

func (t Todos) countCompleted() int {
	total := 0
	for _, item := range t {
		if item.Done {
			total++
		}
	}

	return total
}

func formatTimestamp(timestamp time.Time) string {
	currentTime := time.Now()
	ut := time.Time{}

	if timestamp == ut {
		return Red("...")

	} else if timestamp.Year() == currentTime.Year() && timestamp.YearDay() == currentTime.YearDay() {
		return timestamp.Format(layoutToday)
	}
	return timestamp.Format(layoutOld)
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
