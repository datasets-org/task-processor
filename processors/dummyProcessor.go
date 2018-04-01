package processors

import (
	"github.com/datasets-org/task-processor/structs"
	"github.com/golang/glog"
	"fmt"
)

func dummyProcessor(task structs.Task) (error, string) {
	glog.Infof(`Processing task %s with dummy processor and params "%s"`, task.Id, task.Task.Params)
	return nil, fmt.Sprintf(`Processed with dummy processor and params "%s"`, task.Task.Params)
}