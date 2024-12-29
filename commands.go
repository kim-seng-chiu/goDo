package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"
)

type CmdFlags struct {
	Add     string
	Del     int
	Edit    string
	Toggle  int
	Show    bool
	ShowOne int
	Help    bool
	Db      string
	Token   string
}

func NewCmdFlags() *CmdFlags {
	cf := CmdFlags{}

	flag.StringVar(&cf.Add, "add", "", "Add a new todo specifying the task name")
	flag.IntVar(&cf.Del, "del", -1, "Specify a todo by index to delete effectively")
	flag.StringVar(&cf.Edit, "edit", "", "Edit a todo by index & specify a new task name")
	flag.IntVar(&cf.Toggle, "toggle", -1, "Specify a todo by index to complete it or not")
	flag.BoolVar(&cf.Show, "show", false, "List all todos")
	flag.IntVar(&cf.ShowOne, "showone", -1, "Specify a todo by index to show it, including all details")
	flag.BoolVar(&cf.Help, "help", false, "Shows how to use this tool")
	flag.StringVar(&cf.Db, "db", "", "Specify the database url to use")
	flag.StringVar(&cf.Token, "token", "", "Specify the token to attach to the POST request")

	flag.Parse()

	return &cf
}

func (cf *CmdFlags) Execute(todos *Todos) {
	switch {
	case cf.Show:
		if cf.Db != "" {
			todos.show(cf.Db)
		} else {
			fmt.Println("Please provide a db and token")
		}

	case cf.Add != "":
		if cf.Db != "" && cf.Token != "" {
			todos.add(cf.Add, cf.Db, cf.Token)
		} else {
			fmt.Println("Please provide a db and token")
		}
	case cf.Edit != "":
		split := strings.Split(cf.Edit, ":")
		if len(split) < 2 || len(split) > 2 {
			fmt.Println("Edit flag value has an invalid format, please provide it as <index>:<new task name>")
		}
		index, err := strconv.Atoi(split[0])
		if err != nil {
			fmt.Println("Edit flag index value is not a number, please only provide an int without any quotes")
		}
		todos.edit(index, split[1])
	case cf.Toggle != -1:
		todos.toggle(cf.Toggle)
	case cf.Del != -1:
		todos.delete(cf.Del)
	case cf.ShowOne != -1:
		todos.showOne(cf.ShowOne)
	case cf.Help:
		todos.help()
	default:
		fmt.Println("Unknown flag")
	}
}
