package task

import "fmt"

type Task struct {
	Name     string
	Finished bool
}

type TaskList struct {
	tasks []Task
}

func NewTaskList() *TaskList {
	return &TaskList{tasks: []Task{}}
}

func (t *TaskList) Add(task Task) {
	t.tasks = append(t.tasks, task)
}

func (t *TaskList) List() {
	for i, task := range t.tasks {
		status := "pendiente"
		if task.Finished {
			status = "completada"
		}
		fmt.Printf("[%d] %s - %s\n", i, task.Name, status)
	}
}

func (t *TaskList) Delete(index int) {
	if index < 0 || index >= len(t.tasks) {
		fmt.Println("Índice fuera de rango")
		return
	}
	t.tasks = append(t.tasks[:index], t.tasks[index+1:]...)
	fmt.Println("Tarea eliminada.")
}
