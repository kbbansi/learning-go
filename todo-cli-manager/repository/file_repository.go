package repository

import (
	"encoding/json"
	"os"
	"todo-cli-manager/models"
)

type FileTaskRepository struct {
	FilePath string `json:"filePath"`
}

func (f *FileTaskRepository) setFilePath(path string) (string, error) {
	if path != "" {
		f.FilePath = path
		return path, nil
	}
	return "task.json", nil
}

func (f *FileTaskRepository) SaveTask(task models.Tasks) error {
	file, err := json.MarshalIndent(task, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(f.FilePath, file, 0644)
}

func (f *FileTaskRepository) LoadTask() (models.Tasks, error) {
	if _, err := os.Stat(f.FilePath); os.IsNotExist(err) {
		return models.Tasks{}, nil
	}

	file, err := os.ReadFile(f.FilePath)
	if err != nil {
		return nil, err
	}

	var tasks models.Tasks
	err = json.Unmarshal(file, &tasks)
	return tasks, err
}
