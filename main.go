// CLI todo app
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func displayMessage() {
	fmt.Println("Welcome to the Task Manager")
	fmt.Println("1. Add Task")
	fmt.Println("2. List Task")
	fmt.Println("3. Mark Task as Complete")
	fmt.Println("4. Delete Task")
	fmt.Println("5. Exit")
}

func getChoice(reader *bufio.Reader) (int, error) {
	
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	return strconv.Atoi(input)
}

func getTaskName(reader bufio.Reader) string {
	task, _ := reader.ReadString('\n')
	return strings.TrimSpace(task)
}

func markTaskComplete(tasks []string, taskId int) ([]string, bool) {
	if taskId > 0 && taskId <= len(tasks) {
		tasks[taskId - 1] = "[Done] " + tasks[taskId - 1]
		fmt.Print("Task marked as completed.\n")
		return tasks, true
	}
	fmt.Println("Invalid task number.")
	return tasks, false
}
func deleteTask(tasks []string, taskId int) ([]string, bool) {
	if taskId > 0 && taskId <= len(tasks) {
		tasks = append(tasks[:taskId - 1], tasks[taskId:]...)
		fmt.Print("Task deleted.\n")
		return tasks, true
	}
	fmt.Println("Invalid task number.")
	return tasks, false
}

func main(){
	var tasks []string
	reader := bufio.NewReader(os.Stdin)
	

	for {
		displayMessage()
		fmt.Print("Enter your choice: ")
		choice, err :=getChoice(reader)
		if err != nil {
			fmt.Println(("Invalid input, please enter a number."))
		}

		switch choice {
		case 1:
			//add task
			fmt.Println("Enter your task.")
			// add this to task list
			tasks = append(tasks, getTaskName(*reader))
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
	
			taskId, err :=getChoice(reader)
			if err == nil  {
				markTaskComplete(tasks, taskId)
			} else {
				fmt.Println(("Invalid input, please enter a number."))
			}
		case 4:
			// delete task
			if len(tasks) == 0 {
				fmt.Println("No task available.")
			}
			fmt.Print("Enter Task number to delete: ")

			taskId, err :=getChoice(reader)
			if err == nil {
				tasks, _ = deleteTask(tasks, taskId)
			} else {
				fmt.Println(("Invalid input, please enter a number."))
			}
		case 5:
			// exit task application
			fmt.Println("Exiting the application.")
			os.Exit(0)
		}
	}
}