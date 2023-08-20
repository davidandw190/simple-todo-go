// Package todo provides functionality for a simple CLI todo application.
package todo

import (
	"encoding/json"
	"errors"
	"os"
	"time"
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
