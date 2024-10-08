package sevices

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"todo_list/models"
	"todo_list/utils"
)

func AddTask() error {
	tasks, err := utils.LoadTasks()
	if err != nil {
		return err
	}

	id := len(tasks) + 1
	status := "progress"
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Add task description below: ")
	scanner.Scan()
	description := scanner.Text()

	task := models.Task{Id: id, Description: description, Status: status}

	tasks = append(tasks, task)
	err = utils.SaveTasks(tasks)
	if err != nil {
		return err
	}

	fmt.Println("Task successfully added")
	return nil
}

func DisplayTasks() error {
	tasks, err := utils.LoadTasks()
	if err != nil {
		return err
	}
	content := ""
	for _, task := range tasks {
		content += fmt.Sprintf("%d) %s \t %s\n", task.Id, task.Description, task.Status) +
	}

	if len(tasks) == 0 {
		content = "No tasks found"
	}

	var choice int
	for {
		fmt.Print("To exit Tasks menu type 0: ")
		fmt.Scan(&choice)
		fmt.Print(content)
		if choice == 0 {
			break
		}
	}


	return nil
}

func DeleteTask(id int) error {
	tasks, err := utils.LoadTasks()
	if err != nil {
		return err
	}

	if id <= 0 || id > len(tasks) {
		return errors.New("invalid task id")
	}

	tasks = append(tasks[:id-1], tasks[id:]...)

	err = utils.SaveTasks(tasks)
	if err != nil {
		return err
	}
	fmt.Println("Task successfully deleted")

	return nil
}

func MarkAsDone(id int) error {
	tasks, err := utils.LoadTasks()
	if err != nil {
		return err
	}

	if id <= 0 || id > len(tasks) {
		return errors.New("invalid task id")
	}

	tasks[id-1].Status = "done"

	err = utils.SaveTasks(tasks)
	if err != nil {
		return err
	}
	fmt.Println("Task successfully marked as done")

	return nil
}
