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

	cmd, err := parser.Parse(os.Args[1:])
	if err != nil {
		fmt.Println(err)
		return
	}

	switch cmd.Command {
	case parser.CmdList:
		todos, err := handler.List()
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Printf("%v\n", todos)
		break

	case parser.CmdAdd:
		todo, err := handler.Add(cmd.Add.Name)
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Printf("%v\n", todo)
		break

	default:
		break
	}
}
