package main

import (
	"github.com/datasets/task-processor/storage"
	"github.com/datasets/task-processor/managers"
	"fmt"
)

func main() {
	ms := storage.MapStorage{}
	ms.Create()
	ms.Put("0", map[string]string{"id": "ID 0"})
	t := managers.Tasks{Storage: ms}
	for _, i := range t.GetActiveTasks() {
		fmt.Println(i.Id)
	}
}
