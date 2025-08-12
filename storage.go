package main

import (
	"encoding/json"
	"fmt"
	"os"
)

const fileName = "tasks.json"

func loadTaskFromFile() error {
	data, err := os.ReadFile(fileName)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return fmt.Errorf("Erro ao ler o arquivo de tarefas %w", err)
	}

	if len(data) == 0 {
		return nil
	}

	err = json.Unmarshal(data, &tasks)
	if err != nil {
		return fmt.Errorf("erro ao converter JSON para tarefas %w", err)
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

func saveTaskToFile() error {
	data, err := json.MarshalIndent(tasks, "", " ")
	if err != nil {
		return fmt.Errorf("erro ao converter tarefas em JSON %w", err)
	}
	err = os.WriteFile(fileName, data, 0664)
	if err != nil {
		return fmt.Errorf("erro ao salvar as tarefas no arquivo %w", err)
	}
	return nil
}
