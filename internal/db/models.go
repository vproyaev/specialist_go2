package db

import (
	"time"
)

type Task struct {
	ID    int
	TagId int
	Due   time.Time
}

func (t *Task) IsNull() bool {
	return t.ID == 0
}

type Tag struct {
	ID   int
	Name string
}

func (t *Tag) IsNull() bool {
	return t.ID == 0
}
