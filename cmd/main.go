package main

import (
	"bufio"
	"fmt"
	"os"
	"todo_list/sevices"
)

func main() {
	var choice int
	scanner := bufio.NewScanner(os.Stdin)

	running := true
	for running {

		fmt.Println("Welcome To the TODO list app")
		fmt.Println("Options: ")
		fmt.Println("1. Display tasks")
		fmt.Println("2. Add task")
		fmt.Println("3. Remove task")
		fmt.Println("4. Mark Task as done")
		fmt.Println("0. Exit")
		fmt.Print("Choose an option: ")

		fmt.Scan(&choice)
		scanner.Scan()

		switch choice {
		case 1:
			err := sevices.DisplayTasks()
			if err != nil {
				fmt.Println(err)
			}
			break
		case 2:
			err := sevices.AddTask()
			if err != nil {
				fmt.Println(err)
			}
			break
		case 3:
			var id int

			fmt.Print("Choose a task to delete: ")
			fmt.Scan(&id)
			err := sevices.DeleteTask(id)
			if err != nil {
				fmt.Println(err)
			}
			break
		case 4:
			var id int
			fmt.Print("Choose a task to mark as done: ")
			fmt.Scan(&id)
			err := sevices.MarkAsDone(id)
			if err != nil {
				fmt.Println(err)
			}
			break
		case 0:
			fmt.Println("See you next time")
			running = false
			break
		default:
			fmt.Println("Invalid option")
			break
		}
	}
}
