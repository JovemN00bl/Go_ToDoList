package main

import (
	"encoding/json"
	"fmt"
	"os"
)

var tasks []Task
var nextID int = 1

const fileName = "tasks.json"

func saveTasksToFile() error {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return fmt.Errorf("erro ao converter tarefas para JSON: %w", err)
	}
	err = os.WriteFile(fileName, data, 0644)
	if err != nil {
		return fmt.Errorf("erro ao salvar tarefas no arquivo: %w", err)
	}
	return nil
}

func loadTasksFromFile() error {
	data, err := os.ReadFile(fileName)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return fmt.Errorf("erro ao ler o arquivo de tarefas: %w", err)
	}
	if len(data) == 0 {
		return nil // Arquivo vazio.
	}
	err = json.Unmarshal(data, &tasks)
	if err != nil {
		return fmt.Errorf("erro ao converter JSON para tarefas: %w", err)
	}

	if len(tasks) > 0 {
		maxID := 0
		for _, task := range tasks {
			if task.ID > maxID {
				maxID = task.ID
			}
		}
		nextID = maxID + 1
	}
	return nil
}
