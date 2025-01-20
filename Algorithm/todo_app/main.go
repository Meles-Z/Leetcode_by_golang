package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

type Task struct {
	ID        int
	Name      string
	Status    string
	CreatedAt time.Time
}

type TaskManager struct {
	tasks []Task
}

// add task
func (tm *TaskManager) AddTask(name, status string) {
	id := len(tm.tasks) + 1
	task := Task{
		ID:        id,
		Name:      name,
		Status:    status,
		CreatedAt: time.Now(),
	}
	tm.tasks = append(tm.tasks, task)
}

// display task
func (tm *TaskManager) display() {
	if len(tm.tasks) == 0 {
		fmt.Println("Taks is empty!")
		return
	}
	fmt.Println("Todo list is:")
	for _, task := range tm.tasks {
		fmt.Printf("[%d] %s-%s (Created %s)\n", task.ID, task.Name, task.Status, task.CreatedAt.Format("2006-01-02 15:04:05"))
	}
}

// mark as done
func (tm *TaskManager) MarkAsDone(id int) {
	for i, task := range tm.tasks {
		if task.ID == id {
			if task.Status == "Done" {
				fmt.Println("Task already marked as done!")
				return
			}
			tm.tasks[i].Status = "Done"
			fmt.Printf("Task marked as done is:%s", task.Name)
			return
		}
	}
	fmt.Println("Task not found")
}

// delete task
func (tm *TaskManager) DeleteTask(id int) {
	for i, task := range tm.tasks {
		if task.ID == id {
			tm.tasks = append(tm.tasks[:i], tm.tasks[i+1:]...)
			fmt.Printf("Deleted task is:%s", task.Name)
			return
		}
	}
	fmt.Println("Task not found")
}

func main() {
	taskManager := &TaskManager{}
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Welcome to my todo Appication!")

	for {
		fmt.Println("\n Choose the option you want:")
		fmt.Println("1. Add Task")
		fmt.Println("2. Display Task")
		fmt.Println("3. Mark as done")
		fmt.Println("4. Delete Task")
		fmt.Println("5. Exit")
		fmt.Print("Enter your choice:")
		scanner.Scan()
		choice := scanner.Text()
		switch choice {
		case "1":
			fmt.Print("Enter task name:")
			scanner.Scan()
			taskName := scanner.Text()
			fmt.Println("Enter task status:")
			scanner.Scan()
			taskStatus := scanner.Text()
			taskManager.AddTask(taskName, taskStatus)

		case "2":
			taskManager.display()
		case "3":
			fmt.Print("Enter task id to mark as done:")
			scanner.Scan()
			id, err := strconv.Atoi(scanner.Text())
			if err != nil {
				fmt.Println("Invalid id")
				continue
			}
			taskManager.MarkAsDone(id)

		case "4":
			fmt.Print("Enter task id to mark as done:")
			scanner.Scan()
			id, err := strconv.Atoi(scanner.Text())
			if err != nil {
				fmt.Println("Invalid id")
				continue
			}
			taskManager.DeleteTask(id)
		case "5":
			fmt.Println("Goodbye!")
			return

		default:
			fmt.Println("your option is not correct please try again")
		}
	}
}
