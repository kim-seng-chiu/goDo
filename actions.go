package main

import (
	"os"
	"strconv"
	"time"

	"github.com/aquasecurity/table"
)

type Todo struct {
	TaskName    string
	Completed   bool
	CreatedOn   time.Time
	ModifiedOn  *time.Time
	CompletedAt *time.Time
}

type Todos []Todo

func (todos *Todos) add(taskName string) {
	todo := Todo{
		TaskName:    taskName,
		Completed:   false,
		CreatedOn:   time.Now(),
		ModifiedOn:  nil,
		CompletedAt: nil,
	}

	*todos = append(*todos, todo)
}

func (todos *Todos) delete(index int) error {
	list := *todos

	*todos = append(list[:index], list[index+1:]...)

	return nil
}

func (todos *Todos) toggle(index int) error {
	list := *todos
	todo := &list[index]

	if !todo.Completed {
		completedTime := time.Now()
		todo.CompletedAt = &completedTime
	} else {
		todo.CompletedAt = nil
	}

	todo.Completed = !todo.Completed

	return nil
}

func (todos *Todos) show() {
	table := table.New(os.Stdout)
	table.SetRowLines(false)
	table.SetHeaders("#", "Task Name", "Completed?", "Created On", "Completed At")

	for index, t := range *todos {
		completed := "❌"
		completedAt := ""

		if t.Completed {
			completed = "✅"
			if t.CompletedAt != nil {
				completedAt = t.CompletedAt.Format(time.RFC3339)
			}
		}

		table.AddRow(strconv.Itoa(index), t.TaskName, completed, t.CreatedOn.Format(time.RFC1123), completedAt)
	}

	table.Render()
}
