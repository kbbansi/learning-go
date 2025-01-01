package commands

import (
	"fmt"
	"strconv"
	"todo-cli-manager/repository"
)

type DoneCommand struct {
	Repo   repository.TaskRepository
	TaskID string
}

func (c *DoneCommand) Execute() error {
	id, err := strconv.Atoi(c.TaskID)
	if err != nil {
		return fmt.Errorf("invalid task id: %s", c.TaskID)
	}

	tasks, err := c.Repo.LoadTask()
	if err != nil {
		return err
	}

	found := false
	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Done = true
			found = true
			break
		}
	}

	if !found {
		return fmt.Errorf("task not found: %s", c.TaskID)
	}

	err = c.Repo.SaveTask(tasks)
	if err != nil {
		return err
	}

	fmt.Println("Task marked as done: ", c.TaskID)
	return nil
}
