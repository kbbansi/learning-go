package commands

import (
	"fmt"
	"todo-cli-manager/repository"
)

type ListCommand struct {
	Repo repository.TaskRepository
}

func (c *ListCommand) Execute() error {
	tasks, err := c.Repo.LoadTask()
	if err != nil {
		return err
	}

	if len(tasks) == 0 {
		fmt.Println("No tasks found")
		return nil
	}

	for _, task := range tasks {
		status := " "
		if task.Done {
			status = "👍🏾"
		}
		fmt.Printf("%d. [%s] %s\n", task.ID, status, task.Text)
	}
	return nil
}
