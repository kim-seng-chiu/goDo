package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
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

type StoredTodos struct {
	Items []struct {
		CollectionID   string `json:"collectionId"`
		CollectionName string `json:"collectionName"`
		Completed      bool   `json:"completed"`
		CompletedAt    string `json:"completed_at"`
		CreatedOn      string `json:"created_on"`
		ID             string `json:"id"`
		ModifiedOn     string `json:"modified_on"`
		TaskName       string `json:"task_name"`
	} `json:"items"`
	Page       int `json:"page"`
	PerPage    int `json:"perPage"`
	TotalItems int `json:"totalItems"`
	TotalPages int `json:"totalPages"`
}

func (todos *Todos) add(taskName string, db string, token string) {
	todo := Todo{
		TaskName:    taskName,
		Completed:   false,
		CreatedOn:   time.Now(),
		ModifiedOn:  nil,
		CompletedAt: nil,
	}

	jsonTodo, err := json.Marshal(todo)
	if err != nil {
		fmt.Println("Error marshalling todo:", err)
		os.Exit(1)
	}

	req, err := http.NewRequest("POST", db, bytes.NewBuffer(jsonTodo))
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", token)

	client := &http.Client{}

	response, err := client.Do(req)

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("Unable to read response body: %s", err.Error())
	}

	if responseBody != nil {
		fmt.Printf("Todo created successfully: %s", responseBody)
	} else if response.StatusCode != 200 {
		fmt.Printf("Error creating todo: %d", response.StatusCode)
	}
	fmt.Printf("Todo created successfully: %s", response)

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

func (todos *Todos) show(db string) {
	req, err := http.Get(db)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	bodyResponse, err := io.ReadAll(req.Body)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	stringOfTodos := string(bodyResponse)
	var tasks StoredTodos
	err = json.Unmarshal([]byte(stringOfTodos), &tasks)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	table := table.New(os.Stdout)
	table.SetRowLines(false)
	table.SetHeaders("#", "Task Name", "Completed?", "Created On", "Completed At")

	for i := 0; i < len(tasks.Items); i++ {
		completed := "❌"
		completedAt := ""
		createdOnTime := ""

		if tasks.Items[i].Completed {
			completed = "✅"
			if tasks.Items[i].CompletedAt != "" {
				completedTime, err := time.Parse(time.RFC3339, tasks.Items[i].CompletedAt)
				if err == nil {
					completedAt = completedTime.Format(time.RFC3339)
				}
			}
		}

		createdOn, err := time.Parse(time.RFC3339, tasks.Items[i].CreatedOn)
		if err == nil {
			createdOnTime = createdOn.Format(time.RFC3339)
		}

		table.AddRow(strconv.Itoa(i), tasks.Items[i].TaskName, completed, createdOnTime, completedAt)
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
