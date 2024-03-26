package parser

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type Command int

const (
	CmdList Command = iota
	CmdAdd
	CmdComplete
)

type resultListCmd struct {
}

type resultAddCmd struct {
	Name string
	Due  time.Time
}

type resultCompleteCmd struct {
	Id int
}

type ResultCmd struct {
	Command  Command
	List     resultListCmd
	Add      resultAddCmd
	Complete resultCompleteCmd
}

func ParseCmd(args []string) (ResultCmd, error) {
	if len(args) == 0 {
		return ResultCmd{
			Command: CmdList,
			List:    resultListCmd{},
		}, nil
	}

	switch args[0] {
	case "list":
		return parseCmdList()
	case "add":
		return parseCmdAdd(args[1:])
	case "cmp":
		return parseCmdComplete(args[1:])
	default:
		return ResultCmd{}, errors.New(fmt.Sprintf("[ERR] unknown command %v", args[0]))
	}
}

func parseCmdList() (ResultCmd, error) {
	return ResultCmd{
		Command: CmdList,
		List:    resultListCmd{},
	}, nil
}

func parseCmdAdd(args []string) (ResultCmd, error) {
	if len(args) == 0 {
		return ResultCmd{}, errors.New("[ERR] name is required")
	}

	nameArgs := []string{}
	dueDate := time.Time{}
	// priority := core.PriNone

	for _, arg := range args {
		option := strings.Split(arg, "=")
		if len(option) != 2 {
			nameArgs = append(nameArgs, arg)
			continue
		}

		switch option[0] {
		case "due":
			due, err := parseTime(option[1])
			dueDate = due
			if err != nil {
				return ResultCmd{}, err
			}
		case "pri":
			break
		default:
			return ResultCmd{}, errors.New(fmt.Sprintf("[ERR] unknown option %v", option[0]))
		}
	}

	return ResultCmd{
		Command: CmdAdd,
		Add: resultAddCmd{
			Name: strings.Join(nameArgs, " "),
			Due:  dueDate,
		},
	}, nil
}

func parseCmdComplete(args []string) (ResultCmd, error) {
	if len(args) == 0 {
		return ResultCmd{}, errors.New("id is required")
	}

	id, err := strconv.Atoi(args[0])
	if err != nil {
		return ResultCmd{}, errors.New(fmt.Sprintf("[ERR] invalid id %v", args[0]))
	}

	return ResultCmd{
		Command: CmdComplete,
		Complete: resultCompleteCmd{
			Id: id,
		},
	}, nil
}
