package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Task struct {
	ID          int       `json:"id"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

const filename = "tasks.json"

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: task-cli <operation> [arguments]")
		os.Exit(1)
	}

	operation := os.Args[1]
	taskString := strings.Join(os.Args[2:], " ")

	//check if file exists
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		if err := os.WriteFile(filename, []byte("[]"), 0644); err != nil {
			panic(err)
		}
	}

	switch operation {
	case "add":
		addTask(taskString)
	case "update":
		updateTask(os.Args)

	default:
		fmt.Printf("Unknown operation: %s\n", operation)
		os.Exit(1)
	}

}

func loadTasks() []Task {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	var tasks []Task
	if err := json.Unmarshal(data, &tasks); err != nil {
		panic(err)
	}
	return tasks
}

func saveTasks(tasks []Task) {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		panic(err)
	}
	if err := os.WriteFile(filename, data, 0644); err != nil {
		panic(err)
	}
}

func addTask(description string) {
	if strings.TrimSpace(description) == "" {
		fmt.Println("Error: task description cannot be empty")
		return
	}

	tasks := loadTasks()

	newId := len(tasks) + 1

	newTask := Task{
		ID:          newId,
		Description: description,
		Status:      "todo",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	tasks = append(tasks, newTask)
	saveTasks(tasks)

	fmt.Printf("Task added successfully (ID: %d)\n", newId)
}

func updateTask(args []string) {
	if len(args) < 4 {
		fmt.Println("Usage: task-cli update <id> <description>")
		return
	}

	id := parseID(args[2])
	description := strings.Join(args[3:], " ")

	tasks := loadTasks()

	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Description = description
			tasks[i].UpdatedAt = time.Now()
			saveTasks(tasks)
			fmt.Printf("Task with ID %d updated successfully\n", id)
			return
		}
	}
	fmt.Printf("Task with ID %d not found\n", id)
}

func parseID(idStr string) int {
	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Printf("Error: invalid task ID '%s'\n", idStr)
		os.Exit(1)
	}
	return id
}
