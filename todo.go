package todo

import (
	"errors"
	"time"
)

type item struct {
	Task        string
	Done        bool
	CreatedAt   time.Time
	CompletedAt time.Time
}

type Todos []item

func (t *Todos) Add(task string) {
	todo := item{
		Task:        task,
		Done:        false,
		CreatedAt:   time.Now(),
		CompletedAt: time.Time{},
	}

	*t = append(*t, todo)
}

func (t *Todos) Complete(index int) error {
	lst := *t
	if index <= 0 || index > len(lst) {
		return errors.New("Invalid index!")
	}

	lst[index-1].CompletedAt = time.Now()
	lst[index-1].Done = true

	return nil
}
