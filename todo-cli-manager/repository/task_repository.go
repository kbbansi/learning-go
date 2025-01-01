package repository

import "todo-cli-manager/models"

type TaskRepository interface {
	SaveTask(task models.Tasks) error
	LoadTask() (models.Tasks, error)
}
