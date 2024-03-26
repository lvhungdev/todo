package parser

import (
	"errors"
	"fmt"
	"time"
)

func parseTime(value string) (time.Time, error) {
	if t, err := parseAbsoluteTime(value); err == nil {
		return t, nil
	} else if t, err := parseRelativeTime(value); err == nil {
		return t, nil
	} else if t, err := parseEndOfTime(value); err == nil {
		return t, nil
	}

	return time.Time{}, errors.New(fmt.Sprintf("[ERR] invalid time format %v", value))
}

func parseAbsoluteTime(value string) (time.Time, error) {
	if t, err := time.Parse("2006-01-02T15:04:05", value); err == nil {
		return t, nil
	} else if t, err := time.Parse("2006-01-02", value); err == nil {
		parsedT := time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, 0, t.Location())
		return parsedT, nil
	} else if t, err := time.Parse("15:04:05", value); err == nil {
		now := time.Now()
		parsedT := time.Date(now.Year(), now.Month(), now.Day(), t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), t.Location())
		return parsedT, nil
	} else if t, err := time.Parse("15:04", value); err == nil {
		now := time.Now()
		parsedT := time.Date(now.Year(), now.Month(), now.Day(), t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), t.Location())
		return parsedT, nil
	}

	return time.Time{}, errors.New("")
}

func parseRelativeTime(value string) (time.Time, error) {
	return time.Time{}, nil
}

func parseEndOfTime(value string) (time.Time, error) {
	return time.Time{}, nil
}
