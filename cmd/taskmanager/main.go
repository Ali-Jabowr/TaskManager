package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/Ali-Jabowr/TaskManager/internal/task" // Fixed import path
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Expected 'add' or 'list' subcommands")
		os.Exit(1)
	}

	addCmd := flag.NewFlagSet("add", flag.ExitOnError)
	addDescription := addCmd.String("description", "", "Task description (required)")

	listCmd := flag.NewFlagSet("list", flag.ExitOnError)

	switch os.Args[1] {
	case "add":
		addCmd.Parse(os.Args[2:])
		if *addDescription == "" {
			fmt.Println("Error: description is required")
			addCmd.Usage()
			os.Exit(1)
		}
		err := task.Add(*addDescription)
		if err != nil {
			fmt.Printf("Error adding task: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("Task added successfully")

	case "list":
		listCmd.Parse(os.Args[2:])
		tasks, err := task.List()
		if err != nil {
			fmt.Printf("Error listing tasks: %v\n", err)
			os.Exit(1)
		}
		for i, t := range tasks {
			fmt.Printf("%d. %s\n", i+1, t.Description)
		}

	default:
		fmt.Println("Unknown command")
		os.Exit(1)
	}
}
