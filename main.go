package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var tasks []Task
var nextID int = 1

func main() {

	if err := loadTaskFromFile(); err != nil {
		fmt.Printf("Erro ao carregar tarefas: %v\n", err)
	}

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\n--- Gerenciador de Tarefas ---")
		fmt.Println("1. Adicionar Tarefa")
		fmt.Println("2. Listar Tarefas")
		fmt.Println("3. Marcar Tarefa como Concluída")
		fmt.Println("4. Editar uma tarefa")
		fmt.Println("5. Deletar uma tarefa")
		fmt.Println("6. Sair")
		fmt.Print("Escolha uma opção: ")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch input {
		case "1":
			addTask(reader)
		case "2":
			listTasks()
		case "3":
			markAsCompleted(reader)
		case "4":
			editTask(reader)
		case "5":
			deleteTask(reader)
		case "6":
			fmt.Println("Saindo do gerenciador de tarefas. Até mais!")
			return
		default:
			fmt.Println("Opção inválida. Por favor, escolha novamente.")
		}
	}
}

func editTask(reader *bufio.Reader) {
	if len(tasks) == 0 {
		fmt.Println("Nenhuma tarefa para editar!")
		return
	}

	listTasks()
	fmt.Print("Digite o ID da tarefa que deseja editar: ")
	inputId, _ := reader.ReadString('\n')
	inputId = strings.TrimSpace(inputId)

	id, err := strconv.Atoi(inputId)
	if err != nil {
		fmt.Println("ID inválido. Por favor, digite um número.")
		return
	}

	taskIndex := -1
	for i := range tasks {
		if tasks[i].ID == id {
			taskIndex = i
			break
		}
	}

	if taskIndex == -1 {
		fmt.Println("Tarefa com ID especificado não encontrada.")
		return
	}

	fmt.Printf("Digite a nova descrição da tarefa %d: ", id)
	newDescription, _ := reader.ReadString('\n')
	newDescription = strings.TrimSpace(newDescription)

	if newDescription == "" {
		fmt.Println("A descrição não pode ser vazia.")
		return
	}

	tasks[taskIndex].Description = newDescription

	if err := saveTaskToFile(); err != nil {
		fmt.Printf("Erro ao salvar a alteração: %v\n", err)

	} else {
		fmt.Println("Tarefa editada e salva com sucesso!")
	}
}

func deleteTask(reader *bufio.Reader) {
	if len(tasks) == 0 {
		fmt.Println("Nenhuma tarefa para deletar.")
		return
	}

	listTasks()
	fmt.Println("Digite o ID da tarefa que deseja deletar: ")
	inputId, _ := reader.ReadString('\n')
	inputId = strings.TrimSpace(inputId)

	id, err := strconv.Atoi(inputId)
	if err != nil {
		fmt.Println("ID inválido. Por favor, digite um número.")
		return
	}

	taskIndex := -1
	for i := range tasks {
		if tasks[i].ID == id {
			taskIndex = i
			break
		}
	}

	tasks = append(tasks[:taskIndex], tasks[taskIndex+1:]...)

	if err := saveTaskToFile(); err != nil {
		fmt.Printf("Erro ao salvar após deletar: %v\n", err)
	} else {
		fmt.Println("Tarefa deletada com sucesso!")
	}

}

func addTask(reader *bufio.Reader) {
	fmt.Print("Digite a descrição da tarefa: ")
	description, _ := reader.ReadString('\n')
	description = strings.TrimSpace(description)

	if description == "" {
		fmt.Println("A descrição não pode ser vazia.")
		return
	}

	newTask := Task{ID: nextID, Description: description, Completed: false}
	tasks = append(tasks, newTask)
	nextID++

	if err := saveTaskToFile(); err != nil {
		fmt.Printf("Erro ao salvar a tarefa: %v\n", err)
	} else {
		fmt.Println("Tarefa adicionada com sucesso")
	}

}

func listTasks() {
	if len(tasks) == 0 {
		fmt.Println("Nenhuma tarefa adicionada.")
		return
	}

	fmt.Println("\n-----suas tarefas-----")
	for _, task := range tasks {
		status := "PENDENTE"
		if task.Completed == true {
			status = "COMPLETA"
		}

		fmt.Printf("%d. [%s] %s\n", task.ID, status, task.Description)
	}
}

func markAsCompleted(reader *bufio.Reader) {
	if len(tasks) == 0 {
		fmt.Println("Lista de tarefas vazia...")
		return
	}

	listTasks()
	fmt.Print("Digite o ID da tarefa que deseja marcar como concluída: ")
	inputID, _ := reader.ReadString('\n')
	inputID = strings.TrimSpace(inputID)
	id, err := strconv.Atoi(inputID)
	if err != nil {
		fmt.Println("Id inválido. Por favor digite um número. ")
		return
	}

	found := false
	for i := range tasks {
		if tasks[i].ID == id {
			if tasks[i].Completed == true {
				fmt.Println("Tarefa já está marcada como concluída.")
			} else {
				tasks[i].Completed = true
				fmt.Println("Tarefa marcada como concluída com sucesso! ")
			}

		}
		found = true
		break
	}

	if !found {
		fmt.Println("Tarefa não encontrada!")

	}

}
