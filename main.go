package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

type myApp struct {
	tasksInUi   []Task
	tasksLists  *widget.List
	currentTask *Task
	app         fyne.App
	win         fyne.Window
}

var guiApp myApp

func main() {
	guiApp.app = app.New()
	guiApp.win = guiApp.app.NewWindow("Gerenciador de tarefas")

	if err := loadTasksFromFile(); err != nil {
		fmt.Printf("Erro critico ao carregar as tarefas %v\n ", err)
	}

	guiApp.tasksInUi = tasks
	content := guiApp.makeUI()
	guiApp.win.SetContent(content)
	guiApp.win.Resize(fyne.NewSize(500, 400))
	guiApp.win.ShowAndRun()
}

func (a *myApp) makeUI() *fyne.Container {
	a.tasksLists = widget.NewList(
		func() int { return len(a.tasksInUi) },
		func() fyne.CanvasObject { return widget.NewLabel("Template Task") },
		func(i widget.ListItemID, o fyne.CanvasObject) {
			task := a.tasksInUi[i]
			prefix := "[ ]"
			if task.Completed {
				prefix = "[X]"
			}
			o.(*widget.Label).SetText(fmt.Sprintf("%s %s", prefix, task.Description))
		},
	)

	a.tasksLists.OnSelected = func(id widget.ListItemID) {
		a.currentTask = &a.tasksInUi[id]
	}

	input := widget.NewEntry()
	input.SetPlaceHolder("Nova tarefa...")
	addButton := widget.NewButton("Adicionar", func() {
		if input.Text != "" {
			a.addTask(input.Text)
			input.SetText("")
		}
	})

	inputBar := container.NewBorder(nil, nil, nil, addButton, input)

	markButton := widget.NewButton("Marcar como Concluída", a.markTask)
	editButton := widget.NewButton("Editar", a.editTask)
	deleteButton := widget.NewButton("Deletar", a.deleteTask)
	actionButtons := container.NewHBox(markButton, editButton, deleteButton)

	return container.NewBorder(inputBar, actionButtons, nil, nil, a.tasksLists)
}

func (a *myApp) addTask(description string) {
	newTask := Task{ID: nextID, Description: description, Completed: false}
	a.tasksInUi = append(a.tasksInUi, newTask)
	nextID++
	a.saveAndRefresh()
}

func (a *myApp) markTask() {
	if a.currentTask == nil {
		return
	}
	a.currentTask.Completed = !a.currentTask.Completed
	a.saveAndRefresh()
}

func (a *myApp) deleteTask() {
	if a.currentTask == nil {
		return
	}
	taskIndex := -1
	for i, task := range a.tasksInUi {
		if task.ID == a.currentTask.ID {
			taskIndex = i
			break
		}
	}

	if taskIndex != -1 {
		a.tasksInUi = append(a.tasksInUi[:taskIndex], a.tasksInUi[taskIndex+1:]...)
		a.currentTask = nil
		a.saveAndRefresh()
	}
}

func (a *myApp) editTask() {
	if a.currentTask == nil {
		return
	}
	entry := widget.NewEntry()
	entry.SetText(a.currentTask.Description)
	formDialog := dialog.NewForm("Editar Tarefa", "Salvar", "Cancelar",
		[]*widget.FormItem{{Text: "Descricao", Widget: entry}},
		func(confirmed bool) {
			if confirmed && entry.Text != "" {
				a.currentTask.Description = entry.Text
				a.saveAndRefresh()
			}
		},
		a.win,
	)
	formDialog.Resize(fyne.NewSize(400, 100))
	formDialog.Show()
}

func (a *myApp) saveAndRefresh() {
	tasks = a.tasksInUi

	if err := saveTasksToFile(); err != nil {
		dialog.ShowError(err, a.win)
	}

	a.tasksLists.Refresh()

}
