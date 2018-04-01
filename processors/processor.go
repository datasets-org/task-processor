package processors

import (
	"github.com/datasets-org/task-processor/structs"
	"fmt"
	"github.com/datasets-org/task-processor/managers"
	"time"
)

func Process(task *structs.Task) {
	var err error
	switch task.Task.Operation {
	case "dummy":
		err = dummyProcessor(*task)
	default:
		task.Complete(false, fmt.Sprintf("Uknown operation %s", task.Task.Operation))
	}
	if err == nil {
		task.Complete(true, "")
	} else {
		task.Complete(false, err.Error())
	}
}

func Reactor(tm managers.Tasks) {
	for {
		for _, t := range tm.GetActiveTasks() {
			// todo lock processed tasks
			go Process(&t)
			// todo store result (send via channel)
		}
		time.Sleep(time.Second)
	}
}