package main

import (
	"flag"
	"fmt"
)

type CmdFlags struct {
	Add    string
	Del    int
	Edit   string
	Toggle int
	Show   bool
}

func NewCmdFlags() *CmdFlags {
	cf := CmdFlags{}

	flag.StringVar(&cf.Add, "add", "", "Add a new todo specifying the task name")
	flag.IntVar(&cf.Del, "del", -1, "Specify a todo by index to delete effectively")
	flag.StringVar(&cf.Edit, "edit", "", "Edit a todo by index & specify a new task name")
	flag.IntVar(&cf.Toggle, "toggle", -1, "Specify a todo by index to complete it or not")
	flag.BoolVar(&cf.Show, "show", false, "List all todos")

	flag.Parse()

	return &cf
}

func (cf *CmdFlags) Execute(todos *Todos) {
	switch {
	case cf.Show:
		todos.show()
	case cf.Add != "":
		todos.add(cf.Add)
	case cf.Edit != "":
		fmt.Println("Not implemented yet")
	case cf.Toggle != -1:
		todos.toggle(cf.Toggle)
	case cf.Del != -1:
		todos.delete(cf.Del)
	default:
		fmt.Println("Unknown flag")
	}
}
