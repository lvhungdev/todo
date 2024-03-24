package core

import "time"

type Todo struct {
	Id            int       `db:"id"`
	Name          string    `db:"name"`
	CreatedDate   time.Time `db:"created_date"`
	CompletedDate time.Time `db:"completed_date"`
	IsCompleted   bool      `db:"is_completed"`
	DueDate       time.Time `db:"due_date"`
	Priority      Priority  `db:"priority"`
}

type Priority int

const (
	PriNone Priority = iota
	PriLow
	PriMedium
	PriHigh
)
