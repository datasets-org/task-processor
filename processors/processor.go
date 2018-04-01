package processors

import (
	"github.com/datasets-org/task-processor/structs"
	"fmt"
	"github.com/datasets-org/task-processor/managers"
	"time"
	"errors"
)

func Process(task *structs.Task, tm *managers.Tasks) {
	var err error
	message := ""
	switch task.Task.Operation {
	case "dummy":
		err, message = dummyProcessor(*task)
	default:
		err = errors.New(fmt.Sprintf("Unknown operation %s", task.Task.Operation))
	}
	if err == nil {
		task.Complete(true, message)
	} else {
		task.Complete(false, err.Error())
	}
	tm.Store(*task)
}

func Reactor(tm *managers.Tasks) {
	for {
		for _, t := range tm.GetActiveTasks() {
			if !t.Locked {
				t.Locked = true
				tm.Store(t)
				go Process(&t, tm)
			}
		}
		time.Sleep(time.Second)
	}
}
