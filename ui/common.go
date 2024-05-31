package ui

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/lvhungdev/todo/tracker"
)

func Record(record tracker.Record) string {
	return strings.Join([]string{
		"name: " + record.Name,
		"due:  " + RelativeTime(record.DueDate),
		"pri:  " + Priority(record.Priority),
		"urg:  " + Urgency(record.Urgency()),
	}, "\n")
}

func Records(records []tracker.Record) string {
	if len(records) == 0 {
		return "empty"
	}

	header := []string{"id", "name", "due", "pri", "urg"}
	content := [][]string{}

	for i, r := range records {
		c := []string{
			fmt.Sprint(i + 1),
			r.Name,
			RelativeTime(r.DueDate),
			Priority(r.Priority),
			Urgency(r.Urgency()),
		}
		content = append(content, c)
	}

	return Table(header, content)
}

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
