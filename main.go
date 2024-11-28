package main

import "fmt"


func main() {
	tasks := make(map[string]interface{})
	tasks["Task 1"] = true
	tasks["Task 2"] = false

	for task := range tasks {
		status := "Icompleted"
		fmt.Println(task, status)
	}
}
