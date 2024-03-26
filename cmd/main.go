package main

import (
	"fmt"
	"os"

	"github.com/lvhungdev/todo/core"
	"github.com/lvhungdev/todo/parser"
	"github.com/lvhungdev/todo/repo"
)

func main() {
	repo := repo.NewRepo()
	if err := repo.Init("sqlite3", "./todo.db"); err != nil {
		fmt.Printf("[ERR] %v\n", err)
		return
	}

	handler := core.NewHandler(repo)

	cmd, err := parser.ParseCmd(os.Args[1:])
	if err != nil {
		fmt.Println(err)
		return
	}

	switch cmd.Command {
	case parser.CmdList:
		todos, err := handler.List()
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("%+v\n", todos)
		return

	case parser.CmdAdd:
		todo, err := handler.Add(cmd.Add.Name, cmd.Add.Due)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("%+v\n", todo)
		return

	default:
		return
	}
}
