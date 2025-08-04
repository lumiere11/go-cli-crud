package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/lumiere11/go-example/internal/task"
)

func main() {
	list := task.NewTaskList()

	args := os.Args
	if len(args) < 2 {
		fmt.Println("Uso: cli add|list|delete <nombre>")
		return
	}

	command := strings.ToLower(args[1])

	switch command {
	case "add":
		if len(args) < 3 {
			fmt.Println("Falta el nombre de la tarea")
			return
		}
		name := args[2]
		list.Add(task.Task{Name: name, Finished: false})
		fmt.Println("Tarea agregada:", name)

	case "list":
		list.List()

	case "delete":
		if len(args) < 3 {
			fmt.Println("Falta el índice de la tarea a eliminar")
			return
		}
		var index int
		fmt.Sscanf(args[2], "%d", &index)
		list.Delete(index)

	default:
		fmt.Println("Comando no reconocido:", command)
	}
}
