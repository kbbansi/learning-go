package commands

import (
	"fmt"
	"todo-cli-manager/models"
	"todo-cli-manager/repository"
)

type AddCommand struct {
	Repo repository.TaskRepository
	Text string
}

func (c *AddCommand) Execute() error {
	tasks, err := c.Repo.LoadTask()
	if err != nil {
		return err
	}

	newTask := models.Task{
		ID:   len(tasks) + 1,
		Text: c.Text,
		Done: false,
	}

	tasks = append(tasks, newTask)
	err = c.Repo.SaveTask(tasks)
	if err != nil {
		return err
	}
	fmt.Println("Task added successfully: ", c.Text)
	return nil
}
