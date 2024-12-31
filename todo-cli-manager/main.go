package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strconv"
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
		listCommand.PrintDefaults()
	}

	deleteCommand := flag.NewFlagSet("delete", flag.ExitOnError)
	deleteCommand.Usage = func() {
		fmt.Println("Info: Delete a TODO List Item\nUsage: todo-cli-manager delete <task>")
		deleteCommand.PrintDefaults()
	}

	doneCommand := flag.NewFlagSet("done", flag.ExitOnError)
	doneCommand.Usage = func() {
		fmt.Println("Info: Mark a TODO List Item as Done\nUsage: todo-cli-manager done <task>")
		doneCommand.PrintDefaults()
	}

	// check the command
	if len(os.Args) < 2 {
		fmt.Println("expected 'add', 'list', 'delete', or 'done' subcommands")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "add":
		addCommand.Parse(os.Args[2:])
		addTask()
	case "list":
		listCommand.Parse(os.Args[2:])
		listTasks()
	case "delete":
		deleteCommand.Parse(os.Args[2:])
		deleteTask()
	case "done":
		doneCommand.Parse(os.Args[2:])
		markTaskAsDone()
	default:
		fmt.Println("expected 'add', 'list', 'delete', or 'done' subcommands")
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
		status := " "
		if task.Done {
			status = "👍🏾"
		}
		fmt.Printf("%d. [%s] %s\n", task.ID, status, task.Text)
	}
}

func deleteTask() {
	// Get the task from CLI arguments
	fmt.Println("Deleting task")

	if len(os.Args) < 3 {
		fmt.Println("No task specified\nUsage: todo-cli-manager delete <task>")
		os.Exit(1)
	}

	taskID := os.Args[2]
	id, err := strconv.Atoi(taskID)
	if err != nil {
		fmt.Println("Invalid task ID: ", taskID)
		os.Exit(1)
	}

	tasks := loadTasks()
	if err != nil {
		fmt.Println("Error loading tasks: ", err)
		os.Exit(1)
	}

	// find the task
	updatedTasks := Tasks{}
	found := false

	for _, task := range tasks {
		if task.ID == id {
			found = true
			continue
		} else {
			updatedTasks = append(updatedTasks, task)
		}
	}

	if !found {
		fmt.Println("Task not found: ", taskID)
		os.Exit(1)
	}

	// save the updated tasks
	saveTasks(updatedTasks)
	if err != nil {
		fmt.Println("Error saving tasks: ", err)
		os.Exit(1)
	}

	fmt.Println("Task deleted successfully: ", taskID)
}

func markTaskAsDone() {
	// Get the task from CLI arguments
	fmt.Println("Marking task as done")

	if len(os.Args) < 3 {
		fmt.Println("No task specified\nUsage: todo-cli-manager done <task>")
		os.Exit(1)
	}

	taskID := os.Args[2]
	id, err := strconv.Atoi(taskID)
	if err != nil {
		fmt.Println("Invalid task ID: ", taskID)
		os.Exit(1)
	}

	tasks := loadTasks()
	if err != nil {
		fmt.Println("Error loading tasks: ", err)
		os.Exit(1)
	}

	// find the task
	updatedTasks := Tasks{}
	found := false

	for _, task := range tasks {
		if task.ID == id {
			found = true
			task.Done = true
		}
		updatedTasks = append(updatedTasks, task)
	}

	if !found {
		fmt.Println("Task not found: ", taskID)
		os.Exit(1)
	}

	// save the updated tasks
	saveTasks(updatedTasks)
	if err != nil {
		fmt.Println("Error saving tasks: ", err)
		os.Exit(1)
	}

	fmt.Println("Task marked as done successfully: ", taskID)
}
