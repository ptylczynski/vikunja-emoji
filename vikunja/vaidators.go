package vikunja

import "fmt"

func Validate(task *Task) error {
	if task == nil {
		return fmt.Errorf("Task cannot be nil")
	}
	if task.Title == "" {
		return fmt.Errorf("Title cannot be empty")
	}
	return nil
}
