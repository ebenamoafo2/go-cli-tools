package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/eben/todo"
)

const todoFileName = ".todo.json"

func main() {
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "%s tool. Developed for Task Management\n", filepath.Base(os.Args[0]))
		fmt.Fprintln(flag.CommandLine.Output(), "Copyright 2026")
		fmt.Fprintln(flag.CommandLine.Output(), "Usage information:")
		flag.PrintDefaults()
	}

	task := flag.String("task", "", "Task to be included in the ToDo list")
	list := flag.Bool("list", false, "List all tasks")
	complete := flag.Int("complete", 0, "Item to be completed")
	flag.Parse()

	l := &todo.List{}
	if err := l.Get(todoFileName); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	switch {
	case *list:
		if len(*l) == 0 {
			fmt.Println("No tasks yet. Add one with -task")
			break
		}
		fmt.Print(l)

	case *complete > 0:
		if err := l.Complete(*complete); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		if err := l.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		fmt.Println("Task completed:", *complete)

	case *task != "":
		l.Add(*task)
		if err := l.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		fmt.Println("Task added successfully:", *task)

	default:
		fmt.Fprintln(os.Stderr, "Invalid option")
		flag.Usage()
		os.Exit(1)
	}
}
