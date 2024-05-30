package ui

import (
	"strconv"

	"github.com/lvhungdev/todo/tracker"
)

func Priority(priority tracker.Priority) string {
	switch priority {
	case tracker.PriNone:
		return ""
	case tracker.PriLow:
		return "L"
	case tracker.PriMedium:
		return "M"
	case tracker.PriHigh:
		return "H"
	default:
		return ""
	}
}

func Urgency(urgency float64) string {
	if urgency == 0.0 {
		return ""
	}

	return strconv.FormatFloat(urgency, 'f', 2, 64)
}
