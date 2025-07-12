package task

import (
	"errors"
)

type Status string

const (
	Pending Status = "pending"
	Done    Status = "done"
)

type Task struct {
	ID          int
	Description string
	Status      Status
}

// NewTask vaildates input and creates a task

func NewTask(id int, description string) (*Task, error) {
	descriptionLenght := len(description)
	if description == "" || descriptionLenght < 3 {
		return nil, errors.New("task description cannot be empty")
	}
	return &Task{
		ID:          id,
		Description: description,
		Status:      Pending,
	}, nil

}
