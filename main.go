package main

import (
	"github.com/datasets-org/task-processor/storage"
	"github.com/datasets-org/task-processor/managers"
	"fmt"
	"encoding/json"
)

func main() {
	ms := storage.CreateMapStorage()
	ms.Put("0", `{"id": "ID-0", "created": "2018-04-01T13:07:59.1234", "completed": false}`)
	t := managers.Tasks{Storage: ms}
	for _, i := range t.GetActiveTasks() {
		fmt.Println(i.Id)
		fmt.Println(i.Completed)
		v, _ := json.Marshal(i)
		fmt.Println(string(v))
		i.Complete(true, "ok")
		v, _ = json.Marshal(i)
		fmt.Println(string(v))
	}
}
