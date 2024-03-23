package parser

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Command int

const (
	CmdList Command = iota
	CmdAdd
	CmdCmp
)

type resultList struct {
}

type resultAdd struct {
	Name string
}

type resultCmp struct {
	Id int
}

type Result struct {
	Command Command
	List    resultList
	Add     resultAdd
	Cmp     resultCmp
}

func Parse(args []string) (Result, error) {
	if len(args) == 0 {
		return Result{
			Command: CmdList,
			Add:     resultAdd{},
		}, nil
	}

	switch args[0] {
	case "list":
		return parseList()
	case "add":
		return parseAdd(args[1:])
	case "cmp":
		return parseCmp(args[1:])
	default:
		return Result{}, errors.New(fmt.Sprintf("[ERR] unknown command %v", args[0]))
	}
}

func parseList() (Result, error) {
	return Result{
		Command: CmdList,
		List:    resultList{},
	}, nil
}

func parseAdd(args []string) (Result, error) {
	if len(args) == 0 {
		return Result{}, errors.New("[ERR] name is required")
	}

	return Result{
		Command: CmdAdd,
		Add: resultAdd{
			Name: strings.Join(args, " "),
		},
	}, nil
}

func parseCmp(args []string) (Result, error) {
	if len(args) == 0 {
		return Result{}, errors.New("id is required")
	}

	id, err := strconv.Atoi(args[0])
	if err != nil {
		return Result{}, errors.New(fmt.Sprintf("[ERR] invalid id %v", args[0]))
	}

	return Result{
		Command: CmdCmp,
		Cmp: resultCmp{
			Id: id,
		},
	}, nil
}
