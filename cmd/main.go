package main

import (
	"fmt"
	"os"

	"github.com/lvhungdev/todo/core"
	"github.com/lvhungdev/todo/parser"
	"github.com/lvhungdev/todo/repo"
	"github.com/lvhungdev/todo/ui"
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

		printTodoList(todos)
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

func printTodoList(todos []core.Todo) {
	header := []string{"Id", "Name", "Due"}
	content := [][]string{}

	for _, t := range todos {
		c := []string{fmt.Sprint(t.Id), t.Name, ui.GetRelativeTime(t.DueDate)}
		content = append(content, c)
	}

	table := ui.NewTable(header, content)
	fmt.Print(table.Build())
}
