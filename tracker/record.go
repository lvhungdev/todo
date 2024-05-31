package tracker

import "time"

type Record struct {
	Id            string
	Name          string
	CreatedDate   time.Time
	CompletedDate time.Time
	DueDate       time.Time
	Priority      Priority
}

func (r Record) Completed() bool {
	return !r.CompletedDate.IsZero()
}

func (r Record) Urgency() float64 {
	urgency := 0.0

	if !r.DueDate.IsZero() {
		urgency += 0.2

		d := r.DueDate.Sub(time.Now())
		dInSecs := max(0, (60*60*24*7)-d.Seconds())

		urgencyPerDay := 1.0
		urgencyPerSec := urgencyPerDay / 84600
		urgency += urgencyPerSec * dInSecs
	}

	switch r.Priority {
	case PriLow:
		urgency += 1.0
	case PriMedium:
		urgency += 2.0
	case PriHigh:
		urgency += 4.0
	default:
		break
	}

	return urgency
}

type Priority = int

const (
	PriNone Priority = iota
	PriLow
	PriMedium
	PriHigh
)
