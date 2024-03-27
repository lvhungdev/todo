package parser

import (
	"errors"
	"fmt"
	"strconv"
	"time"
)

func parseTime(value string) (time.Time, error) {
	if parsedTime, err := parseAbsoluteTime(value); err == nil {
		return parsedTime, nil
	} else if parsedTime, err := parseRelativeTime(value); err == nil {
		return parsedTime, nil
	} else if parsedTime, err := parseEndOfTime(value); err == nil {
		return parsedTime, nil
	}

	return time.Time{}, errors.New(fmt.Sprintf("[ERR] invalid time format %v", value))
}

func parseAbsoluteTime(value string) (time.Time, error) {
	if parsedTime, err := time.Parse("2006-01-02T15:04:05", value); err == nil {
		return parsedTime, nil
	} else if parsedTime, err := time.Parse("2006-01-02", value); err == nil {
		newParsedTime := time.Date(parsedTime.Year(), parsedTime.Month(), parsedTime.Day(), 23, 59, 59, 0, parsedTime.Location())
		return newParsedTime, nil
	} else if parsedTime, err := time.Parse("15:04:05", value); err == nil {
		now := time.Now()
		newParsedTime := time.Date(now.Year(), now.Month(), now.Day(), parsedTime.Hour(), parsedTime.Minute(), parsedTime.Second(), parsedTime.Nanosecond(), parsedTime.Location())
		return newParsedTime, nil
	} else if parsedTime, err := time.Parse("15:04", value); err == nil {
		now := time.Now()
		newParsedTime := time.Date(now.Year(), now.Month(), now.Day(), parsedTime.Hour(), parsedTime.Minute(), parsedTime.Second(), parsedTime.Nanosecond(), parsedTime.Location())
		return newParsedTime, nil
	}

	return time.Time{}, errors.New("")
}

func parseRelativeTime(value string) (time.Time, error) {
	if len(value) < 2 {
		return time.Time{}, errors.New("")
	}

	now := time.Now()

	amount := value[:len(value)-1]
	unit := value[len(value)-1:]

	parsedAmount, err := strconv.Atoi(amount)
	if err != nil {
		return time.Time{}, errors.New("")
	}

	switch unit {
	case "s":
		return now.Add(time.Second * time.Duration(parsedAmount)), nil
	case "m":
		return now.Add(time.Minute * time.Duration(parsedAmount)), nil
	case "h":
		return now.Add(time.Hour * time.Duration(parsedAmount)), nil
	case "d":
		return now.Add(time.Hour * 24 * time.Duration(parsedAmount)), nil
	case "w":
		return now.Add(time.Hour * 24 * 7 * time.Duration(parsedAmount)), nil
	default:
		return time.Time{}, nil
	}
}

func parseEndOfTime(value string) (time.Time, error) {
	return time.Time{}, nil
}
