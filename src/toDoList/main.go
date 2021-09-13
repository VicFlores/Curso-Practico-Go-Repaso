package main

import "fmt"

// Struct
type Task struct {
	nombre      string
	description string
	completado  bool
}

type TaskList struct {
	task []*Task
}

// Actualizar tareas
func (t *Task) TaskCompleted() {
	t.completado = true
}

func (t *Task) UpdateDescription(description string) {
	t.description = description
}

func (t *Task) updateName(name string) {
	t.nombre = name
}

func (tl *TaskList) AddList(t *Task) {
	tl.task = append(tl.task, t)
}

// Eliminar task de lista
func (tl *TaskList) deleteList(index int) {
	tl.task = append(tl.task[:index], tl.task[index+1:]...)
}

// Imprimir task en lista
func (t *TaskList) PrintList() {

	for _, task := range t.task {
		fmt.Println("Nombre:", task.nombre)
		fmt.Println("Description:", task.description)
	}
}

// Imprimir tasks completados
func (t *TaskList) PrintCompleted() {

	for _, task := range t.task {
		if task.completado {
			fmt.Println("Nombre:", task.nombre)
			fmt.Println("Description:", task.description)
		}
	}
}

func main() {

	// Task
	task1 := &Task{
		nombre:      "Completar curso practico de Go",
		description: "Terminarlo en esta semana",
	}

	task2 := &Task{
		nombre:      "Completar curso avanzado de Go",
		description: "Terminarlo en 3 dias",
	}

	task3 := &Task{
		nombre:      "Completar curso docker",
		description: "Terminarlo en 5 dias",
	}

	// Task list
	list := &TaskList{
		task: []*Task{
			task1, task2,
		},
	}

	// eliminar de lista
	// list.deleteList(1)

	// agregar a lista
	list.AddList(task3)

	// 1 forma de iterar
	fmt.Println("Primera forma de iterar")
	for i := 0; i < len(list.task); i++ {
		fmt.Println("Index:", i, "Nombre:", list.task[i].nombre)
	}

	// 2 forma de iterar
	fmt.Println("Segunda forma de iterar")
	for index, task := range list.task {
		fmt.Println("Index:", index, "Nombre:", task.nombre)
	}

	// Function que itera tasks
	list.PrintList()

	// Function para tareas completadas
	list.task[0].TaskCompleted()

	fmt.Println("Tareas completadas")

	list.PrintCompleted()

	// Maps
	mapTask := make(map[string]*TaskList)

	mapTask["Vic"] = list

	task4 := &Task{
		nombre:      "Completar curso avanzado de JS",
		description: "Terminarlo en 3 dias",
	}

	task5 := &Task{
		nombre:      "Completar curso Js",
		description: "Terminarlo en 5 dias",
	}

	list2 := &TaskList{
		task: []*Task{
			task4, task5,
		},
	}

	mapTask["Alejandra"] = list2

	fmt.Println("Tareas Vic")
	mapTask["Vic"].PrintList()

	fmt.Println("Tareas Alejandra")
	mapTask["Alejandra"].PrintList()

}
