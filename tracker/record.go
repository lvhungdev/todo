package tracker

import "time"

type Record struct {
	Name          string
	CreatedDate   time.Time
	CompletedDate time.Time
	DueDate       time.Time
	Priority      Priority
}

type Priority = int

const (
	PriNone Priority = iota
	PriLow
	PriMedium
	PriHeight
)
