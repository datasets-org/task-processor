package managers

import (
	"github.com/datasets-org/task-processor/structs"
	"github.com/datasets-org/task-processor/storage"
	"encoding/json"
)

type Tasks struct {
	Storage storage.Storage
}

func (t *Tasks) GetActiveTasks() (tasks []structs.Task)  {
	for _, i := range t.Storage.Items() {
		t := structs.CreateTask()
		json.Unmarshal([]byte(i), &t)
		tasks = append(tasks, t)
	}
	return
}
