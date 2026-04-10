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

const (
	filename       = "tasks.json"
	markInProgress = "mark-in-progress"
	markDone       = "mark-done"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: task-cli <operation> [arguments]")
		os.Exit(1)
	}

	operation := os.Args[1]

	// Check if file exists
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		if err := os.WriteFile(filename, []byte("[]"), 0644); err != nil {
			panic(err)
		}
	}

	if strings.HasPrefix(operation, markInProgress) || strings.HasPrefix(operation, markDone) {
		operation = "markTask"
	}

	switch operation {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Usage: task-cli add <description>")
			os.Exit(1)
		}
		addTask(strings.Join(os.Args[2:], " "))
	case "update":
		updateTask(os.Args)
	case "delete":
		deleteTask(os.Args)
	case "markTask":
		markTask(os.Args)
	case "list":
		listTasks(os.Args)
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

	maxId := 0
	for _, task := range tasks {
		if task.ID > maxId {
			maxId = task.ID
		}
	}

	newTask := Task{
		ID:          maxId + 1,
		Description: description,
		Status:      "todo",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	tasks = append(tasks, newTask)
	saveTasks(tasks)
	fmt.Printf("Task added successfully (ID: %d)\n", newTask.ID)
}

func updateTask(args []string) {
	if len(args) < 4 {
		fmt.Println("Usage: task-cli update <id> <description>")
		return
	}

	id := parseID(args[2])
	description := strings.Join(args[3:], " ")

	if strings.TrimSpace(description) == "" {
		fmt.Println("Error: task description cannot be empty")
		return
	}

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

func deleteTask(args []string) {
	if len(args) < 3 {
		fmt.Println("Usage: task-cli delete <id>")
		return
	}
	id := parseID(args[2])
	tasks := loadTasks()
	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			saveTasks(tasks)
			fmt.Printf("Task with ID %d deleted successfully\n", id)
			return
		}
	}
	fmt.Printf("Task with ID %d not found\n", id)
}

func markTask(args []string) {
	if len(args) < 3 {
		fmt.Println("Usage: task-cli mark-in-progress|mark-done <id>")
		return
	}
	id := parseID(args[2])
	tasks := loadTasks()
	for i, task := range tasks {
		if task.ID == id {
			if strings.HasPrefix(args[1], markInProgress) {
				tasks[i].Status = "in-progress"
			} else if strings.HasPrefix(args[1], markDone) {
				tasks[i].Status = "done"
			}
			tasks[i].UpdatedAt = time.Now()
			saveTasks(tasks)
			fmt.Printf("Task with ID %d marked as %s\n", id, tasks[i].Status)
			return
		}
	}
	fmt.Printf("Task with ID %d not found\n", id)
}

func listTasks(args []string) {
	tasks := loadTasks()
	if len(tasks) == 0 {
		fmt.Println("No tasks found.")
		return
	}

	validStatuses := map[string]bool{
		"todo":        true,
		"in-progress": true,
		"done":        true,
	}

	// Filter by status if provided
	if len(args) >= 3 {
		filter := strings.Join(args[2:], " ")
		if !validStatuses[filter] {
			fmt.Printf("Error: unknown status '%s'. Use: todo, in-progress, done\n", filter)
			return
		}
		printHeader()
		for _, task := range tasks {
			if task.Status == filter {
				printTask(task)
			}
		}
		return
	}

	printHeader()
	for _, task := range tasks {
		printTask(task)
	}
}

func printHeader() {
	fmt.Println("ID | Description | Status | Created At | Updated At")
	fmt.Println("------------------------------------------------------")
}

func printTask(task Task) {
	fmt.Printf("%d | %s | %s | %s | %s\n",
		task.ID,
		task.Description,
		task.Status,
		task.CreatedAt.Format(time.RFC3339),
		task.UpdatedAt.Format(time.RFC3339),
	)
}
