package main

import (
	"fmt"
	"time"
)

// Task represents a task in the time management application.
type Task struct {
	ID          int
	Description string
	DueDate     time.Time
}

// NewTask creates a new Task instance with the given ID, description, and due date.
func NewTask(id int, description string, dueDate time.Time) *Task {
	return &Task{
		ID:          id,
		Description: description,
		DueDate:     dueDate,
	}
}

// TimeManagementApp represents a time management application.
type TimeManagementApp struct {
	Tasks []*Task
}

// NewTimeManagementApp creates a new TimeManagementApp instance.
func NewTimeManagementApp() *TimeManagementApp {
	return &TimeManagementApp{
		Tasks: []*Task{},
	}
}

// AddTask adds a new task to the time management application.
func (tma *TimeManagementApp) AddTask(task *Task) {
	tma.Tasks = append(tma.Tasks, task)
}

// GetTasks returns all the tasks in the time management application.
func (tma *TimeManagementApp) GetTasks() []*Task {
	return tma.Tasks
}

// GetTasksDueToday returns the tasks that are due today.
func (tma *TimeManagementApp) GetTasksDueToday() []*Task {
	tasksDueToday := []*Task{}
	today := time.Now().Truncate(24 * time.Hour)

	for _, task := range tma.Tasks {
		if task.DueDate.Truncate(24*time.Hour).Equal(today) {
			tasksDueToday = append(tasksDueToday, task)
		}
	}

	return tasksDueToday
}

// Usage Example for TimeManagementApp

func main() {
	// Create a new time management application
	app := NewTimeManagementApp()

	// Add tasks to the application
	task1 := NewTask(1, "Complete project", time.Date(2022, time.January, 15, 0, 0, 0, 0, time.UTC))
	task2 := NewTask(2, "Submit report", time.Date(2022, time.January, 20, 0, 0, 0, 0, time.UTC))
	task3 := NewTask(3, "Attend meeting", time.Date(2022, time.January, 15, 15, 0, 0, 0, time.UTC))

	app.AddTask(task1)
	app.AddTask(task2)
	app.AddTask(task3)

	// Get all tasks in the application
	allTasks := app.GetTasks()
	fmt.Println("All tasks:")
	for _, task := range allTasks {
		fmt.Printf("Task ID: %d, Description: %s, Due Date: %s\n", task.ID, task.Description, task.DueDate.Format("2006-01-02"))
	}

	// Get tasks due today
	tasksDueToday := app.GetTasksDueToday()
	fmt.Println("\nTasks due today:")
	for _, task := range tasksDueToday {
		fmt.Printf("Task ID: %d, Description: %s, Due Date: %s\n", task.ID, task.Description, task.DueDate.Format("2006-01-02"))
	}
}
