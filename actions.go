package main

import (
	"fmt"
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

func (todos *Todos) edit(index int, taskName string) {
	list := *todos
	todo := &list[index]

	todo.TaskName = taskName
	modifiedTime := time.Now()
	todo.ModifiedOn = &modifiedTime
}

func (todos *Todos) help() {
	fmt.Print(`
Usage of todo:
	-add <string>    		Adds a new todo specifying the task name
	-delete <index>  		Deletes a todo at the specified index
	-edit <index>:<string>   	Edits the task name of the todo at the specified index
	-toggle <index>  		Toggles the completion status of the todo at the specified index
	-show            		Shows all todos
	-showOne <index> 		Shows details of the todo at the specified index
	-help            		Shows this help message
	`)
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

func (todos *Todos) showOne(index int) {
	if index < 0 || index >= len(*todos) {
		fmt.Println("No task found with given index")
		return
	}

	todo := (*todos)[index]

	table := table.New(os.Stdout)
	table.SetRowLines(false)
	table.SetHeaders("Task Name", "Completed?", "Created On", "Completed At", "Modified On")

	completed := "❌"
	completedAt := ""

	modifiedOn := ""

	if todo.Completed {
		completed = "✅"
		if todo.CompletedAt != nil {
			completedAt = todo.CompletedAt.Format(time.RFC3339)
		}
	}

	if todo.ModifiedOn != nil {
		modifiedOn = todo.ModifiedOn.Format(time.RFC3339)
	}

	table.AddRow(todo.TaskName, completed, todo.CreatedOn.Format(time.RFC1123), completedAt, modifiedOn)
	table.Render()
}
