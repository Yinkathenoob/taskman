package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Task represents one item in the task list
type Task struct {
	ID    int
	Title string
	Done  bool
}

func addTask(tasks []Task, title string, id int) ([]Task, int) {
	task := Task{ID: id, Title: title, Done: false}
	tasks = append(tasks, task)
	id++
	return tasks, id
}

// saveTasks writes the current list of tasks to a JSON file. It returns an error if the operation fails.
func saveTasks(tasks []Task) error {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile("tasks.json", data, 0644)
}

// loadTasks reads the list of tasks from a JSON file. If the file doesn't exist, it returns an empty list. It returns an error if the operation fails.
func loadTasks() ([]Task, error) {
	data, err := os.ReadFile("tasks.json")
	if err != nil {
		if os.IsNotExist(err) {
			return []Task{}, nil // no file yet — start with an empty list, no error
		}
		return nil, err // some other real error
	}

	// Unmarshal the JSON data into the slice of tasks
	var tasks []Task
	err = json.Unmarshal(data, &tasks)
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

// markDone sets the Done field of the task with the given ID to true. It returns true if the task was found and marked, or false if no task with that ID exists.
func markDone(tasks []Task, id int) bool {
	for i := range tasks {
		if tasks[i].ID == id {
			tasks[i].Done = true
			return true // found and marked it
		}
	}
	return false // no task with that ID
}

// main is the entry point of the application. It provides a terminal interface for managing tasks.
func main() {
	reader := bufio.NewReader(os.Stdin)

	// Load existing tasks from the JSON file
	tasks, err := loadTasks()
	if err != nil {
		fmt.Println("Error loading tasks:", err)
		return
	}
	nextID := 1
	for _, t := range tasks {
		if t.ID >= nextID {
			nextID = t.ID + 1
		}
	}
	// Main loop for the task manager interface
	for {
		fmt.Println("\n--- Task Manager ---")
		fmt.Println("1. Add a task")
		fmt.Println("2. List tasks")
		fmt.Println("3. Mark a task done")
		fmt.Println("4. Quit")
		fmt.Print("Choose an option: ")

		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		// Handle the user's menu choice
		switch choice {
		case "1":
			fmt.Print("Enter task title: ")
			title, _ := reader.ReadString('\n')
			title = strings.TrimSpace(title)

			task := Task{ID: nextID, Title: title, Done: false}
			tasks, nextID = addTask(tasks, title, nextID)

			fmt.Printf("Added task #%d: %s\n", task.ID, task.Title)
		// Save the tasks after adding a new one
		case "2":
			if len(tasks) == 0 {
				fmt.Println("No tasks yet.")
				break
			}
			fmt.Println("\nYour tasks:")
			for _, t := range tasks {
				status := " "
				if t.Done {
					status = "x"
				}
				fmt.Printf("  [%s] #%d %s\n", status, t.ID, t.Title)
			}
			// Save the tasks after listing them
		case "3":
			fmt.Print("Enter task ID to mark done: ")
			idInput, _ := reader.ReadString('\n')
			idInput = strings.TrimSpace(idInput)

			id, err := strconv.Atoi(idInput)
			if err != nil {
				fmt.Println("That's not a valid ID:", idInput)
				break
			}

			if markDone(tasks, id) {
				fmt.Printf("Marked task #%d as done.\n", id)
			} else {
				fmt.Println("No task found with ID:", id)
			}
			// Save the tasks after marking one as done
		case "4":
			if err := saveTasks(tasks); err != nil {
				fmt.Println("failed to save tasks:", err)
			} else {
				fmt.Println("Tasks saved. Goodbye!")
			}
			return

		default:
			fmt.Println("Invalid option, please choose 1, 2, 3, or 4.")
		}
	}
}
