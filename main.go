package main

import (
	"fmt"
	"log"
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
		log.Fatal(err)
	}
	defer s.Close()

	t, err := tracker.NewTracker(s)
	if err != nil {
		log.Fatal(err)
	}

	cmd, err := command.Parse(os.Args[1:])
	if err != nil {
		fmt.Println(err)
		return
	}

	switch cmd := cmd.(type) {
	case command.List:
		printRecords(t)
	case command.Add:
		addNewRecord(t, cmd.Name, cmd.DueDate, cmd.Priority)
	case command.Complete:
		completeRecord(t, cmd.Index)
	}
}

func printRecords(t *tracker.Tracker) {
	records := t.ListActive()
	if len(records) == 0 {
		fmt.Println("empty")
		return
	}

	header := []string{"id", "name", "due", "pri", "urg"}
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

func addNewRecord(t *tracker.Tracker, name string, dueDate time.Time, pri tracker.Priority) {
	index, err := t.Add(name, dueDate, pri)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("added record %v\n", index+1)
}

func completeRecord(t *tracker.Tracker, index int) {
	err := t.Complete(index)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("completed record %v\n", index+1)
}
