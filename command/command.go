package command

import (
	"fmt"
	"strings"
	"time"

	"github.com/lvhungdev/todo/tracker"
)

type Command interface {
	time() time.Time
}

type baseCommand struct {
	at time.Time
}

func (b baseCommand) time() time.Time {
	return b.at
}

type List struct {
	baseCommand
}

type Add struct {
	baseCommand
	Name     string
	DueDate  time.Time
	Priority tracker.Priority
}

func Parse(args []string) (Command, error) {
	if len(args) == 0 {
		return parseListCommand()
	}

	switch args[0] {
	case "list":
		return parseListCommand()
	case "add":
		return parseAddCommand(args[1:])
	default:
		return nil, fmt.Errorf("[ERR] unknown command: %s", args[0])
	}
}

func parseListCommand() (Command, error) {
	return List{baseCommand{time.Now()}}, nil
}

func parseAddCommand(args []string) (Command, error) {
	if len(args) == 0 {
		return nil, fmt.Errorf("[ERR] name is required")
	}

	nameArgs := []string{}
	dueDate := time.Time{}
	priority := tracker.PriNone

	for _, arg := range args {
		opt := strings.Split(arg, "=")
		if len(opt) != 2 {
			nameArgs = append(nameArgs, arg)
			continue
		}

		switch opt[0] {
		case "due":
			due, err := parseTime(opt[1])
			dueDate = due
			if err != nil {
				return nil, err
			}

		case "pri":
			switch opt[1] {
			case "n":
				priority = tracker.PriNone
			case "l":
				priority = tracker.PriLow
			case "m":
				priority = tracker.PriMedium
			case "h":
				priority = tracker.PriHigh
			}

		default:
			return nil, fmt.Errorf("[ERR] unknown option %v", opt[0])
		}
	}

	return Add{
		baseCommand: baseCommand{time.Now()},
		Name:        strings.Join(nameArgs, " "),
		DueDate:     dueDate,
		Priority:    priority,
	}, nil
}