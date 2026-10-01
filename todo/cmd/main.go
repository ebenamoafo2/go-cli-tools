package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/eben/todo"
)

const todoFileName = ".todo.json"

func main() {

	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "%s tool. Developed for Task Management\n", os.Args[0])
		fmt.Fprintf(flag.CommandLine.Output(), "Copyright 2026\n")
		fmt.Fprintln(flag.CommandLine.Output(), "Usage information:")
		flag.PrintDefaults()
	}

	//Parsing the command line flags
	task := flag.String("task", "", "Task to be included in the ToDo list")
	list := flag.Bool("list", false, "List all tasks")
	complete := flag.Int("complete", 0, "Item to be completed")
	flag.Parse()

	l := &todo.List{}
	if err := l.Get(todoFileName); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// Decide what to do based on the number of arguments provided
	switch {
	case *list:
		pending := 0
		//List current todo items
		for _, item := range *l {
			if !item.Done {
				fmt.Println(item.Task)
				pending++
			}

		}
		switch {
		case len(*l) == 0:
			fmt.Println("No tasks yet. Add one with -task")
		case pending == 0:
			fmt.Println("All tasks completed")
		}
	case *complete > 0:
		//complete the given item
		if err := l.Complete(*complete); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		if err := l.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	case *task != "":
		//Add the task
		l.Add(*task)
		fmt.Println("Task added successfully:", *task)

		//Save the new list
		if err := l.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	default:
		//Invalid flag provided
		fmt.Fprintln(os.Stderr, "Invalid option")
		os.Exit(1)
	}
}
