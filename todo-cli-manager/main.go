package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
)

// Task struct representing a single todo item
type Task struct {
	ID   int    `json:"id"`
	Text string `json:"text"`
	Done bool   `json:"done"`
}

// Tasks is a slice (or list) of Task
type Tasks []Task

func main() {

	// CLI commands
	addCommand := flag.NewFlagSet("add", flag.ExitOnError)
	addCommand.Usage = func() {
		fmt.Println("Info: Add a new TODO List Item\nUsage: todo-cli-manager add <task>")
		addCommand.PrintDefaults()
	}

	listCommand := flag.NewFlagSet("list", flag.ExitOnError)
	listCommand.Usage = func() {
		fmt.Println("Usage: todo-cli-manager list")
		addCommand.PrintDefaults()
	}

	// check the command
	if len(os.Args) < 2 {
		fmt.Println("expected 'add' or 'list' subcommands")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "add":
		addCommand.Parse(os.Args[2:])
		addTask()
	case "list":
		listCommand.Parse(os.Args[2:])
		listTasks()
	default:
		fmt.Println("expected 'add' or 'list' subcommands")
		os.Exit(1)
	}

}

func addTask() {
	// Get the task from CLI arguments
	fmt.Println("Adding task")

	if len(os.Args) < 3 {
		fmt.Println("No task specified\nUsage: todo-cli-manager add <task>")
		os.Exit(1)
	}

	taskText := os.Args[2]

	// load existing tasks
	tasks := loadTasks()

	// create a new task
	newTask := Task{
		ID:   len(tasks) + 1,
		Text: taskText,
		Done: false,
	}
	tasks = append(tasks, newTask)

	// save tasks
	saveTasks(tasks)
	fmt.Println("Task added successfully: ", taskText)
}

func loadTasks() Tasks {
	// check if file exists
	if _, err := os.Stat("tasks.json"); os.IsNotExist(err) {
		// file does not exist
		return Tasks{}
	}

	// read the file
	file, err := os.ReadFile("tasks.json")
	if err != nil {
		fmt.Println("Error reading tasks.json file: ", err)
		os.Exit(1)
	}

	// Unmarshal JSON
	var tasks Tasks
	err = json.Unmarshal(file, &tasks)
	if err != nil {
		fmt.Println("Error reading tasks file: ", err)
		os.Exit(1)
	}
	return tasks
}

func saveTasks(tasks Tasks) {
	// Marshal tasks to JSON

	file, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		fmt.Println("Error marshalling tasks: ", err)
		os.Exit(1)
	}

	// write to file
	err = os.WriteFile("tasks.json", file, 0644)
	if err != nil {
		fmt.Println("Error writing tasks to file: ", err)
		os.Exit(1)
	}
}

func listTasks() {
	// load tasks
	tasks := loadTasks()

	// display tasks
	if len(tasks) == 0 {
		fmt.Println("No tasks found")
		return
	}

	for _, task := range tasks {
		status := "Pending"
		if task.Done {
			status = "Done"
		}
		fmt.Printf("%d. %s [%s]\n", task.ID, task.Text, status)
	}
}
