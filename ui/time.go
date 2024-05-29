package ui

import (
	"fmt"
	"time"
)

func GetRelativeTime(t time.Time) string {
	if t.IsZero() {
		return ""
	}

	now := time.Now()
	duration := now.Sub(t)
	seconds := int(duration.Seconds())

	format := func(value int, prefix string) string {
		if value < 60 {
			return fmt.Sprintf("%v%vs", prefix, value)
		} else if value < 3600 {
			return fmt.Sprintf("%v%vm %vs", prefix, value/60, value%60)
		} else if value < 86400 {
			return fmt.Sprintf("%v%vh %vm", prefix, value/3600, (value%3600)/60)
		} else {
			return fmt.Sprintf("%v%vd %vh", prefix, value/86400, (value%86400)/3600)
		}
	}

	if seconds < 0 {
		return format(seconds*-1, "")
	} else {
		return format(seconds, "-")
	}
}
