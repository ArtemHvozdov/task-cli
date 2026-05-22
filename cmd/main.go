package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/ArtemHvozdov/task-cli/internal/entity"
	"github.com/ArtemHvozdov/task-cli/internal/service"
	"github.com/ArtemHvozdov/task-cli/internal/storage"
)

func printTasks(tasks []entity.Task) {
	if len(tasks) == 0 {
		fmt.Println("No tasks found.")
		return
	}
	for _, task := range tasks {
		fmt.Printf("ID: %d, Description: %s, Status: %s, CreatedAt: %s, UpdatedAt: %s\n",
			task.ID, task.Description, task.Status, task.CreatedAt.Format("2006-01-02 15:04:05"), task.UpdatedAt.Format("2006-01-02 15:04:05"))
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("command required")
		return
	}

	mainStorage, err := storage.NewJSONStorage("tasks.json")
	if err != nil {
		fmt.Printf("failed to initialize storage: %v\n", err)
		return
	}

	mainService := service.NewService(mainStorage)

	command := os.Args[1]

	switch command {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("description required for add command")
			return
		}
		description := os.Args[2]
		id, err := mainService.Add(description)
		if err != nil {
			fmt.Printf("failed to add task: %v\n", err)
			return
		}
		fmt.Printf("Task added successfully (ID: %d)\n", id)
	case "update":
		if len(os.Args) < 4 {
			fmt.Println("id and description required for update command")
			return
		}
		id := os.Args[2]
		iD, err := strconv.ParseUint(id, 10, 16)
		if err != nil {
			fmt.Printf("invalid id: %v\n", err)
			return
		}
		description := os.Args[3]
		err = mainService.Update(uint16(iD), description)
		if err != nil {
			fmt.Printf("failed to update task: %v\n", err)
			return
		}
		fmt.Printf("Task %d updated successfully\n", uint16(iD))
	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("id required for delete command")
			return
		}
		id := os.Args[2]
		iD, err := strconv.ParseUint(id, 10, 16)
		if err != nil {
			fmt.Printf("invalid id: %v\n", err)
			return
		}
		err = mainService.Delete(uint16(iD))
		if err != nil {
			fmt.Printf("failed to delete task: %v\n", err)
			return
		}
		fmt.Printf("Task %d deleted successfully\n", uint16(iD))
	case "mark-in-progress":
		if len(os.Args) < 3 {
			fmt.Println("id required for mark-in-progress command")
			return
		}
		id := os.Args[2]
		iD, err := strconv.ParseUint(id, 10, 16)
		if err != nil {
			fmt.Printf("invalid id: %v\n", err)
			return
		}
		err = mainService.UpdateStatus(uint16(iD), entity.InProgress)
		if err != nil {
			fmt.Printf("failed to update task status: %v\n", err)
			return
		}
		fmt.Printf("Task %d marked as in progress\n", uint16(iD))
	case "mark-done":
		if len(os.Args) < 3 {
			fmt.Println("id required for mark-done command")
			return
		}
		id := os.Args[2]
		iD, err := strconv.ParseUint(id, 10, 16)
		if err != nil {
			fmt.Printf("invalid id: %v\n", err)
			return
		}
		err = mainService.UpdateStatus(uint16(iD), entity.Done)
		if err != nil {
			fmt.Printf("failed to update task status: %v\n", err)
			return
		}
		fmt.Printf("Task %d marked as done\n", uint16(iD))
	case "list":
		if len(os.Args) < 3 {
			tasks, err := mainService.GetAll()
			if err != nil {
				fmt.Printf("failed to get tasks: %v\n", err)
				return
			}
			printTasks(tasks)
			return
		}

		parameter := entity.Status(os.Args[2])
		if parameter != entity.Todo && parameter != entity.InProgress && parameter != entity.Done {
			fmt.Printf("invalid status: %s\n", parameter)
			return
		}

		tasks, err := mainService.GetByStatus(parameter)
		if err != nil {
			fmt.Printf("failed to get tasks: %v\n", err)
			return
		}
		printTasks(tasks)
	default:
		fmt.Printf("unknown command: %s\n", command)
	}

}