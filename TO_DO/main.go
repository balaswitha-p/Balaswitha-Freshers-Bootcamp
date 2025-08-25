package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: go run main.go [command] [arguments]")
		fmt.Println("Commands: add,list,complete,delete")
		return
	}

	command := os.Args[1]

	switch command {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("usage: go run main.go add \"task description\"")
			return
		}
		description := os.Args[2]
		err := addTask(description)
		if err != nil {
			fmt.Println("Error adding task: ", err)
		} else {
			fmt.Println("Task added successfully")
		}

	case "list":
		err := listTasks()
		if err != nil {
			fmt.Println("Error adding task: ", err)
		}

	case "complete":
		if len(os.Args) < 3 {
			fmt.Println("Usage: go run main.go complete [task number]")
			return
		}
		index, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Error: Invalid task number.")
			return
		}
		err = completeTask(index)
		if err != nil {
			fmt.Println("Error completing task:", err)
		} else {
			fmt.Println("Task marked as complete.")
		}

	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("Usage: go run main.go delete [task number]")
			return
		}
		index, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Error: Invalid task number")
			return
		}
		err = deleteTask(index)
		if err != nil {
			fmt.Println("Error deletig the task: ", err)
		} else {
			fmt.Println("Task Deleted")
		}

	default:
		fmt.Println("Unknown Command: ", command)
	}
}
