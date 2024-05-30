package main

import (
	"fmt"
	"os"
	"time"

	"github.com/lvhungdev/todo/command"
	"github.com/lvhungdev/todo/storage"
	"github.com/lvhungdev/todo/tracker"
	"github.com/lvhungdev/todo/ui"
)

func main() {
	s, err := storage.NewStore("./todo.db")
	if err != nil {
		panic(err)
	}
	defer s.Close()

	t, err := tracker.NewTracker(s)
	if err != nil {
		panic(err)
	}

	cmd, err := command.Parse(os.Args[1:])
	if err != nil {
		panic(err)
	}

	switch cmd := cmd.(type) {
	case command.List:
		printRecords(t)
	case command.Add:
		addNewRecord(t, cmd.Name, cmd.DueDate, cmd.Priority)
	}
}

func addNewRecord(t *tracker.Tracker, name string, dueDate time.Time, pri tracker.Priority) {
	if err := t.Add(name, dueDate, pri); err != nil {
		panic(err)
	}
	fmt.Println("Added")
}

func printRecords(t *tracker.Tracker) {
	records := t.ListActive()

	header := []string{"Id", "Name", "Due", "Pri", "Urg"}
	content := [][]string{}

	for i, r := range records {
		c := []string{
			fmt.Sprint(i + 1),
			r.Name,
			ui.RelativeTime(r.DueDate),
			ui.Priority(r.Priority),
			ui.Urgency(r.Urgency()),
		}
		content = append(content, c)
	}

	fmt.Print(ui.Table(header, content))
}
