package processors

import (
	"github.com/datasets-org/task-processor/structs"
	"github.com/golang/glog"
	"fmt"
	"errors"
	"encoding/json"
)

type params struct {
	Fs   string `json:"Fs"`
	Path string `json:"Path"`
}

func scanProcessor(task structs.Task) (error, string) {
	glog.Infof(`Processing task %s with scan processor and params "%s"`, task.Id, task.Task.Params)
	p, err := parseParams(task.Task.Params)
	if err != nil {
		e := fmt.Sprintf("Params cannot be parsed val: %s", task.Task.Params)
		glog.Error(e)
		return errors.New(e), ""
	}
	glog.Infof("%s %s", p.Fs, p.Path)
	return nil, fmt.Sprintf(`Processed with scan processor and params "%s"`, task.Task.Params)
}

func parseParams(param string) (params, error) {
	p := params{}
	err := json.Unmarshal([]byte(param), &p)
	if err != nil {
		glog.Errorf("Params cannot be parsed\n%s", param)
		return p, err
	} else {
		return p, nil
	}
}
