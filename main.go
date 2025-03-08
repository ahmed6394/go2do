// CLI todo app
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main(){
	var tasks []string
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("CLI-Based To-Do app task")
	fmt.Println("1. Add a Task")
	fmt.Println("2. List of tasks")
	fmt.Println("3. Mark a task as Completed")
	fmt.Println("4. Delete a task")
	fmt.Println("5. Exit")

	for {
		fmt.Print("Enter your choice: ")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		choice, err :=strconv.Atoi(input)
		if err != nil {
			fmt.Println(("Invalid input, please enter a number."))
		}

		switch choice {
		case 1:
			//add task
			fmt.Println("Enter your task.")
			task, _ := reader.ReadString('\n')
			task = strings.TrimSpace(task)

			// add this to task list
			tasks = append(tasks, task)
			fmt.Println("Task added.")
		case 2:
			// task list
			if len(tasks) == 0 {
				fmt.Println("No task available.")
			} else {
				fmt.Println(("Your tasks: "))
				for index, task := range tasks {
					fmt.Printf("%d - %s\n", index+1, task)
				}
			}
		case 3:
			// mark task complete
			if len(tasks) == 0 {
				fmt.Println("No task available.")
			}
			fmt.Print("Enter Task number to complete: ")
			taskIdString, _ := reader.ReadString('\n')
			taskIdString = strings.TrimSpace(taskIdString)
	
			taskId, err :=strconv.Atoi(taskIdString)
			if err == nil && taskId > 0 && taskId <= len(tasks) {
				tasks[taskId - 1] = "[Done] " + tasks[taskId - 1]
				fmt.Print("Task marked as completed.\n")
			} else {
				fmt.Println(("Invalid input, please enter a number."))
			}
		case 4:
			// delete task
		case 5:
			// exit task application
		}
	}
}