// This file contains logic for managing TO-DO tasks

package main

import (
	"encoding/json"
	"errors"
	"os"
)

// Single to-do item
type Task struct {
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

const datafile = "data/tasks.json"

// saves the tasks to JSON file
func saveTasks(tasks []Task) error {
	file, err := json.MarshalIndent(tasks, "", " ")
	if err != nil {
		return err
	}
	return os.WriteFile(datafile, file, 0644)
}

// loads tasks from a JSON file
func loadTasks() ([]Task, error) {
	file, err := os.ReadFile(datafile)
	if err != nil {
		if os.IsNotExist(err) {
			return []Task{}, nil
		}
		return nil, err
	}
	var tasks []Task
	err = json.Unmarshal(file, &tasks)
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

// adding task
func addTask(description string) error {
	tasks, err := loadTasks()
	if err != nil {
		return err
	}
	tasks = append(tasks, Task{Description: description, Completed: false})
	return saveTasks(tasks)
}

// prints all tasks
func listTasks() error {
	tasks, err := loadTasks()
	if err != nil {
		return err
	}
	if len(tasks) == 0 {
		println("No tasks to show!")
		return nil
	}

	println("Your TO+DO List: ")
	for i, task := range tasks {
		status := " "
		if task.Completed {
			status = "X"
		}
		println(i+1, "[", status, "]", task.Description)
	}
	return nil
}

// marking as completed
func completeTask(index int) error {
	tasks, err := loadTasks()
	if err != nil {
		return err
	}
	if index < 1 || index > len(tasks) {
		return errors.New("invalid task index")
	}
	tasks[index-1].Completed = true
	return saveTasks(tasks)
}

func deleteTask(index int) error {
	tasks, err := loadTasks()
	if err != nil {
		return err
	}
	if index < 1 || index > len(tasks) {
		return errors.New("inavlid task index")
	}
	tasks = append(tasks[:index-1], tasks[index:]...)
	return saveTasks(tasks)
}
