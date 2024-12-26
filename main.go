package main

import (
	"fmt"
)

func main() {
	todos := Todos{}

	storage := NewStorage[Todos]("/todos.json")
	err := storage.Load(&todos)
	if err != nil {
		fmt.Println("Could not find any tasks. Starting afresh.")
	}
	// Parse & Execute
	cmdFlags := NewCmdFlags()
	cmdFlags.Execute(&todos)
	// Save to storage
	err = storage.Save(todos)
	if err != nil {
		fmt.Printf("Error saving todos in storage: %v\n", err)
	}
}
