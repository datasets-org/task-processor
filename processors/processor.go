package processors

import (
	"github.com/datasets-org/task-processor/structs"
	"fmt"
	"github.com/datasets-org/task-processor/managers"
	"time"
	"errors"
	"github.com/golang/glog"
)

const datasetFilename = "dataset.yaml"

func Process(task *structs.Task, tm *managers.Tasks) {
	var err error
	message := ""
	switch task.Task.Operation {
	case "dummy":
		err, message = dummyProcessor(*task)
	case "scan":
		err, message = scanProcessor(*task)
	default:
		glog.Infof(`Processing task %s - unknown operation (%s) params "%s"`, task.Id, task.Task.Operation,
			task.Task.Params)
		err = errors.New(fmt.Sprintf("Unknown operation %s", task.Task.Operation))
	}
	if err == nil {
		task.Complete(true, message)
	} else {
		task.Complete(false, err.Error())
	}
	task.Locked = false
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
