package models

import "fmt"

type Task struct {
	ID   int    `json:"id"`
	Text string `json:"text"`
	Done bool   `json:"done"`
}

type Tasks []Task

func (t *Task) print() {
	status := " "
	// check if the task is done
	if t.Done {
		status = "✅"
	}
	fmt.Printf("%d. [%s] %s\n", t.ID, status, t.Text)
}
