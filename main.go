package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

// Command constants
const (
	cmdHelp          = "help"
	cmdAdd           = "add"
	cmdUpdate        = "update"
	cmdDelete        = "delete"
	cmdMarkDone      = "mark-done"
	cmdMarkProgress  = "mark-in-progress"
	cmdList          = "list"
	defaultTasksFile = "tasks.json"
)

// Status constants
const (
	statusTodo       = "todo"
	statusDone       = "done"
	statusInProgress = "in-progress"
)

// Task represents a task with its properties
type Task struct {
	ID          int       `json:"id"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func run() error {
	if len(os.Args) <= 1 {
		printUsage()
		return errors.New("no command provided")
	}

	command := os.Args[1]
	args := os.Args[2:]

	switch command {
	case cmdHelp, "--help", "-h":
		printUsage()
		return nil

	case cmdAdd:
		if len(args) == 0 || strings.TrimSpace(args[0]) == "" {
			return fmt.Errorf("usage: ./task-cli %s <description>", cmdAdd)
		}
		return addTask(args[0])

	case cmdUpdate:
		if len(args) < 2 {
			return fmt.Errorf("usage: ./task-cli %s <id> <description>", cmdUpdate)
		}
		id, err := parseID(args[0])
		if err != nil {
			return err
		}
		return updateTask(id, args[1])

	case cmdDelete:
		if len(args) == 0 {
			return fmt.Errorf("usage: ./task-cli %s <id>", cmdDelete)
		}
		id, err := parseID(args[0])
		if err != nil {
			return err
		}
		return deleteTask(id)

	case cmdMarkDone, cmdMarkProgress:
		if len(args) == 0 {
			return fmt.Errorf("usage: ./task-cli %s <id>", command)
		}
		id, err := parseID(args[0])
		if err != nil {
			return err
		}
		status := strings.TrimPrefix(command, "mark-")
		status = strings.ReplaceAll(status, "-", " ")
		return markTask(id, status)

	case cmdList:
		filter := ""
		if len(args) > 0 {
			filter = args[0]
		}
		return listTasks(filter)

	default:
		printUsage()
		return fmt.Errorf("invalid command: %s", command)
	}
}

func parseID(idStr string) (int, error) {
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return 0, errors.New("invalid task ID: must be a number")
	}
	if id <= 0 {
		return 0, errors.New("invalid task ID: must be greater than zero")
	}
	return id, nil
}

func printUsage() {
	fmt.Println("            Task Tracker CLI            ")
	fmt.Println("        a simple CLI task tracker       ")
	fmt.Println("========================================")
	fmt.Println("Usage: ./task-cli <command> [arguments]")
	fmt.Println("========================================")
	fmt.Println("Commands:")
	fmt.Println("  add <description>            Add a new task")
	fmt.Println("  update <id> <description>    Update a task's description")
	fmt.Println("  delete <id>                  Delete a task")
	fmt.Println("  mark-done <id>               Mark a task as done")
	fmt.Println("  mark-in-progress <id>        Mark a task as in-progress")
	fmt.Println("  list                         List all tasks")
	fmt.Println("  list [filter]                List tasks (filter: todo|in-progress|done)")
	fmt.Println("========================================")
}

func getTasksFilePath() string {
	return defaultTasksFile
}

func readTasks() ([]Task, error) {
	filePath := getTasksFilePath()

	// Check if file exists
	if _, err := os.Stat(filePath); errors.Is(err, os.ErrNotExist) {
		// Return empty tasks list if file doesn't exist yet
		return []Task{}, nil
	}

	f, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open tasks file: %w", err)
	}
	defer f.Close()

	var tasks []Task
	data, err := io.ReadAll(f)
	if err != nil {
		return nil, fmt.Errorf("failed to read tasks file: %w", err)
	}

	// Empty file case
	if len(data) == 0 {
		return []Task{}, nil
	}

	if err := json.Unmarshal(data, &tasks); err != nil {
		return nil, fmt.Errorf("failed to parse tasks data: %w", err)
	}

	return tasks, nil
}

func writeTasks(tasks []Task) error {
	filePath := getTasksFilePath()

	// Create directories if they don't exist
	dir := filepath.Dir(filePath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}

	// Create or truncate the file
	f, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("failed to create tasks file: %w", err)
	}
	defer f.Close()

	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to serialize tasks: %w", err)
	}

	if _, err := f.Write(data); err != nil {
		return fmt.Errorf("failed to write tasks to file: %w", err)
	}

	return nil
}

func generateNextID(tasks []Task) int {
	maxID := 0
	for _, task := range tasks {
		if task.ID > maxID {
			maxID = task.ID
		}
	}
	return maxID + 1
}

func addTask(description string) error {
	description = strings.TrimSpace(description)
	if description == "" {
		return errors.New("task description cannot be empty")
	}

	tasks, err := readTasks()
	if err != nil {
		return err
	}

	newTask := Task{
		ID:          generateNextID(tasks),
		Description: description,
		Status:      statusTodo,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	tasks = append(tasks, newTask)

	if err := writeTasks(tasks); err != nil {
		return err
	}

	fmt.Printf("Task added successfully (ID: %d)\n", newTask.ID)
	return nil
}

func updateTask(id int, newDescription string) error {
	newDescription = strings.TrimSpace(newDescription)
	if newDescription == "" {
		return errors.New("task description cannot be empty")
	}

	tasks, err := readTasks()
	if err != nil {
		return err
	}

	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Description = newDescription
			tasks[i].UpdatedAt = time.Now()

			if err := writeTasks(tasks); err != nil {
				return err
			}

			fmt.Printf("Task updated successfully (ID: %d, Description: %s)\n", id, newDescription)
			return nil
		}
	}

	return fmt.Errorf("task with ID %d not found", id)
}

func deleteTask(id int) error {
	tasks, err := readTasks()
	if err != nil {
		return err
	}

	newTasks := make([]Task, 0, len(tasks))
	found := false

	for _, task := range tasks {
		if task.ID != id {
			newTasks = append(newTasks, task)
		} else {
			found = true
		}
	}

	if !found {
		return fmt.Errorf("task with ID %d not found", id)
	}

	if err := writeTasks(newTasks); err != nil {
		return err
	}

	fmt.Printf("Task deleted successfully (ID: %d)\n", id)
	return nil
}

func validateStatus(status string) error {
	validStatuses := map[string]bool{
		statusTodo:       true,
		statusDone:       true,
		statusInProgress: true,
	}

	if !validStatuses[status] {
		return fmt.Errorf("invalid status: %s (must be todo, done, or in-progress)", status)
	}
	return nil
}

func listTasks(filter string) error {
	tasks, err := readTasks()
	if err != nil {
		return err
	}

	if len(tasks) == 0 {
		fmt.Println("No tasks found.")
		return nil
	}

	printTask := func(task Task) {
		fmt.Printf("ID: %d | Description: %s | Status: %s\n",
			task.ID, task.Description, task.Status)
	}

	switch filter {
	case statusTodo, statusDone, statusInProgress:
		found := false
		for _, task := range tasks {
			if task.Status == filter {
				printTask(task)
				found = true
			}
		}
		if !found {
			fmt.Printf("No tasks found with status: %s\n", filter)
		}
	case "":
		for _, task := range tasks {
			printTask(task)
		}
	default:
		return fmt.Errorf("invalid filter: %s (use todo, done, or in-progress)", filter)
	}

	return nil
}

func markTask(id int, status string) error {
	if err := validateStatus(status); err != nil {
		return err
	}

	tasks, err := readTasks()
	if err != nil {
		return err
	}

	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Status = status
			tasks[i].UpdatedAt = time.Now()

			if err := writeTasks(tasks); err != nil {
				return err
			}

			fmt.Printf("Task marked as %s successfully!\n", status)
			return nil
		}
	}

	return fmt.Errorf("task with ID %d not found", id)
}
