package commands

import (
	"fmt"
	"strconv"
	"todo-cli-manager/models"
	"todo-cli-manager/repository"
)

type DeleteCommand struct {
	Repo   repository.TaskRepository
	TaskID string
}

func (c *DeleteCommand) Execute() error {

	id, err := strconv.Atoi(c.TaskID)
	if err != nil {
		return fmt.Errorf("invalid task id: %s", c.TaskID)
	}

	tasks, err := c.Repo.LoadTask()
	if err != nil {
		return err
	}

	updatedTasks := models.Tasks{}
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
		return fmt.Errorf("task not found: %s", c.TaskID)
	}

	err = c.Repo.SaveTask(updatedTasks)
	if err != nil {
		return err
	}

	fmt.Printf("Task %s deleted successfully\n", c.TaskID)
	return nil
}
