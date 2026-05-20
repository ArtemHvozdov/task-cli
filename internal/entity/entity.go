package entity

import "time"

type Task struct {
	ID uint16
	Description string
	Status Status // todo | in progress | done
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Status string

const (
	Todo         Status = "todo"
	InProgress   Status = "in-progress"
	Done         Status = "done"
)