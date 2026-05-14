package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"time"
)

const tasksFile string = "tasks.json"

type Task struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Done      bool      `json:"done"`
	CreatedAt time.Time `json:"created_at"`
}

func main() {
	if len(os.Args) < 2 {
		usage()
		os.Exit(1)
	}
	switch os.Args[1] {
	case "add":

		if len(os.Args) < 3 {
			log.Println("Not enough arguments, can't add a task")
			usage()
			os.Exit(1)
		}
		newTaskTitle := os.Args[2]
		if newTaskTitle == "" {
			log.Println("No title's been passed in")
			usage()
			os.Exit(1)
		}
		tasks, err := loadTasks(tasksFile)
		if err != nil {
			log.Fatal("Failed to load tasks: ", err)
		}
		updatedTasks, newTask := cmdAdd(tasks, newTaskTitle)
		err = saveTasks(tasksFile, updatedTasks)
		if err != nil {
			log.Fatal("Failed to save tasks:", err)
		}
		fmt.Println("Added task", newTask.ID)
	case "list":
		fmt.Println("we'll be listing tasks here")
	case "complete":
		fmt.Println("we'll be completing a task here")
	case "delete":
		fmt.Println("We'll be deleting a task here")
	default:
		usage()
		os.Exit(1)
	}
}

func usage() {
	fmt.Println(`Usage: todo <command> [args]

Commands:
  add <title>    Add a new task
  list           List all tasks
  complete <id>  Mark a task complete
  delete <id>    Delete a task`)
}

func cmdAdd(tasks []Task, title string) ([]Task, Task) {
	newTask := Task{
		ID:        nextID(tasks),
		Title:     title,
		Done:      false,
		CreatedAt: time.Now(),
	}
	return append(tasks, newTask), newTask
}

func nextID(tasks []Task) int {
	var max int
	for _, task := range tasks {
		if task.ID > max {
			max = task.ID
		}
	}

	return max + 1
}

func loadTasks(path string) ([]Task, error) {
	var tasks []Task
	data, err := os.ReadFile(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {

			return tasks, nil
		}
		return nil, err
	}
	err = json.Unmarshal(data, &tasks)
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func saveTasks(path string, tasks []Task) error {
	jsonData, err := json.MarshalIndent(tasks, "", " ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, jsonData, 0644)
}
