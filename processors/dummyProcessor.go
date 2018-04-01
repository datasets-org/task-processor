package processors

import (
	"github.com/datasets-org/task-processor/structs"
	"github.com/golang/glog"
)

func dummyProcessor(task structs.Task) error {
	glog.Infof("Processing task %s with dummy processor and params %s", task.Id, task.Task.Params)
	return nil
}