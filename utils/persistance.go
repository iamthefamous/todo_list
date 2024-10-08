package utils

import (
	"encoding/json"
	"io"
	"os"
	"todo_list/models"
)

const filePath = "task.json"

func LoadTasks() ([]models.Task, error) {
	file, err := os.Open(filePath)
	if os.IsNotExist(err) {
		return []models.Task{}, nil
	}

	if err != nil {
		return nil, err
	}

	defer file.Close()

	var tasks []models.Task

	byteValue, _ := io.ReadAll(file)
	err = json.Unmarshal(byteValue, &tasks)
	return tasks, err
}

func SaveTasks(tasks []models.Task) error {
	file, err := json.MarshalIndent(tasks, "", " ")
	if err != nil {
		return err
	}
	return os.WriteFile(filePath, file, 0644)

}
