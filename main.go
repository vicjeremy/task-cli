package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type Task struct {
	ID          int       `json:"id" `
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func main() {
	if len(os.Args) <= 1 {
		printUsage()
		os.Exit(1)
	}

	command := os.Args[1]
	args := os.Args[2:]
	id := parse(command, args)

	switch command {
	case "help", "--help", "-h":
		printUsage()
	case "add":
		description := args[0]
		addTask(description)
	case "update":
		newDescription := args[1]
		updateTask(id, newDescription)
	case "delete":
		deleteTask(id)
	case "mark-done", "mark-in-progress":
		status := strings.TrimPrefix(command, "mark-")
		markTask(id, status)
	case "list":
		filter := ""
		if len(args) > 0 {
			filter = args[0]
		}
		listTasks(filter)
	default:
		fmt.Println("Invalid command")
		printUsage()
		os.Exit(1)
	}

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

func parse(command string, args []string) int {
	if command == "update" && len(args) < 2 {
		fmt.Printf("Usage: ./task-cli %s <id> <description>\n", command)
		os.Exit(1)
	}
	if command == "add" && (len(args) == 0 || strings.TrimSpace(args[0]) == "") {
		fmt.Printf("Usage: ./task-cli %s <description>\n", command)
		os.Exit(1)
	}
	if command == "update" || command == "delete" || command == "mark-in-progress" || command == "mark-done" {
		if len(args) == 0 {
			fmt.Printf("Usage: ./task-cli %s <id>\n", command)
			os.Exit(1)
		}
		id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Invalid id")
			os.Exit(1)
		}
		return id
	} else {
		return 0
	}

}

func readTasks() []Task {
	f, err := os.OpenFile("tasks.json", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var tasks []Task
	decoder := json.NewDecoder(f)
	err = decoder.Decode(&tasks)
	if err != nil && err.Error() != "EOF" {
		fmt.Println("Error decoding JSON:", err)
	}
	return tasks
}

func writeTasks(tasks []Task) {
	f, err := os.OpenFile("tasks.json", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	encoder := json.NewEncoder(f)
	if err := encoder.Encode(tasks); err != nil {
		log.Fatalf("Error encoding JSON: %v", err)
	}
}

func addTask(description string) {
	if strings.TrimSpace(description) == "" {
		fmt.Println("Error: Description cannot be empty")
		os.Exit(1)
		return
	}
	tasks := readTasks()

	newTask := Task{
		ID:          len(tasks) + 1,
		Description: description,
		Status:      "todo",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	tasks = append(tasks, newTask)

	writeTasks(tasks)
	fmt.Println("Task added successfully (ID: ", newTask.ID, ")")
}

func updateTask(id int, newDescription string) {
	tasks := readTasks()

	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Description = newDescription
			tasks[i].UpdatedAt = time.Now()
			writeTasks(tasks)
			fmt.Println("Task updated successfully (ID: ", id, ", Description: ", newDescription, ")")
			return
		}
	}

	fmt.Println("Task not found!")
}

func deleteTask(id int) {
	tasks := readTasks()
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
		fmt.Println("Task not found!")
		return
	}

	writeTasks(newTasks)
	fmt.Println("Task deleted successfully (ID:", id, ")")
}

func listTasks(filter string) {
	tasks := readTasks()
	if len(tasks) == 0 {
		fmt.Println("No tasks found.")
		return
	}
	printTask := func(task Task) {
		fmt.Printf("ID: %d | Description: %s | Status: %s\n", task.ID, task.Description, task.Status)
	}
	switch filter {
	case "done", "todo", "in-progress":
		filter = strings.ReplaceAll(filter, "-", " ")
		found := false
		for _, task := range tasks {
			if task.Status == filter {
				printTask(task)
				found = true
			}
		}
		if !found {
			fmt.Println("No tasks found with status:", filter)
		}
	case "":
		for _, task := range tasks {
			printTask(task)
		}
	default:
		fmt.Println("Invalid command, please use one of the following: done, todo, in-progress")
	}
}
func markTask(id int, status string) {
	tasks := readTasks()

	for i, task := range tasks {
		if task.ID == id {
			status = strings.ReplaceAll(status, "-", " ")
			tasks[i].Status = status
			tasks[i].UpdatedAt = time.Now()
			writeTasks(tasks)
			fmt.Println("Task marked as", status, "successfully!")
			return
		}
	}

	fmt.Println("Task not found!")
}
