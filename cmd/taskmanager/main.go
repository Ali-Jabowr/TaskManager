package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	//Define subcommands

	addCmd := flag.NewFlagSet("add", flag.ExitOnError)
	listCmd := flag.NewFlagSet("list", flag.ExitOnError)

	// Add command flags
	addDescription := addCmd.String("description", "", "Task description")

	if len(os.Args) < 2 {
		fmt.Println("Expected 'add' or 'list' subcommands")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "add":
		addCmd.Parse(os.Args[2:])
		fmt.Printf("Adding task: %s\n", *addDescription)
	case "list":
		listCmd.Parse(os.Args[2:])
		fmt.Println("Listing all tasks")
	default:
		fmt.Println("Unksnown command")
		os.Exit(1)
	}
}
