package main

import (
	"fmt"
	"log"
	"os"
	"path"
	"time"

	"github.com/lvhungdev/todo/command"
	"github.com/lvhungdev/todo/storage"
	"github.com/lvhungdev/todo/tracker"
	"github.com/lvhungdev/todo/ui"
)

func main() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("[WRN] Failed to get home dir, using current dir to store data instead")
		homeDir = "."
	}

	s, err := storage.NewStore(path.Join(homeDir, "todo.db"))
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
		fmt.Println(ui.Records(t.GetActive()))
	case command.Next:
		listNextRecord(t)
	case command.Add:
		addNewRecord(t, cmd.Name, cmd.DueDate, cmd.Priority)
	case command.Complete:
		completeRecord(t, cmd.Index)
	}
}

func listNextRecord(t *tracker.Tracker) {
	records := t.GetActive()
	if len(records) == 0 {
		fmt.Println("empty")
		return
	}

	fmt.Println(ui.Record(records[0]))
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
	if err := t.Complete(index); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("completed record %v\n", index+1)
}
