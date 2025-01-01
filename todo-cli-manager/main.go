package main

import (
	"fmt"
	"os"
	"todo-cli-manager/commands"
	"todo-cli-manager/repository"
)

func main() {

	repo := &repository.FileTaskRepository{
		FilePath: "../tasks.json",
	}

	if len(os.Args) < 2 {
		fmt.Println("expected 'add', 'list', 'delete', or 'done' subcommands")
		os.Exit(1)
	}

	// CLI commands
	var cmd commands.Command

	switch os.Args[1] {
	case "add":
		checkParams(os.Args, 3, "Info: Add a new TODO List Item\nUsage: todo-cli-manager add <task>")
		cmd = &commands.AddCommand{Repo: repo, Text: os.Args[2]}
	case "list":
		cmd = &commands.ListCommand{Repo: repo}
	case "delete":
		checkParams(os.Args, 3, "Info: Delete a TODO List Item\nUsage: todo-cli-manager delete <task>")
		cmd = &commands.DeleteCommand{Repo: repo, TaskID: os.Args[2]}
	case "done":
		checkParams(os.Args, 3, "Info: Mark a TODO List Item as Done\nUsage: todo-cli-manager done <task>")
		cmd = &commands.DoneCommand{Repo: repo, TaskID: os.Args[2]}
	default:
		fmt.Println("expected 'add', 'list', 'delete', or 'done' subcommands")
		os.Exit(1)
	}

	err := cmd.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}

func checkParams(args []string, expected int, usage string) {
	if len(args) < expected {
		fmt.Println("Usage: ", usage)
		os.Exit(1)
	}
}
