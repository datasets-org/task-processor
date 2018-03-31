package managers

import (
	"github.com/datasets-org/task-processor/structs"
	"github.com/datasets-org/task-processor/storage"
)

type Tasks struct {
	Storage storage.Storage
}

func (t *Tasks) GetActiveTasks() (tasks []structs.Task)  {
	for _, i := range t.Storage.Items() {
		// todo load from json
		tasks = append(tasks, structs.Task{Id: i["id"]})
	}
	return
}
