package ui

import "github.com/lvhungdev/todo/tracker"

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
