package processors

import (
	"github.com/datasets-org/task-processor/structs"
	"github.com/golang/glog"
	"fmt"
	"errors"
	"encoding/json"
	"path/filepath"
	"path"
	"os"
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
	switch p.Fs {
	case "local":
		scanFs(p.Path)
	case "http":
		scanHttp(p.Path)
	default:
		e := fmt.Sprintf("Unknown FS %s", p.Fs)
		glog.Error(e)
		return errors.New(e), ""
	}

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

func scanFs(scanPath string) {
	err := filepath.Walk(scanPath, func(pth string, f os.FileInfo, err error) error {
		if !f.IsDir() && f.Name() == datasetFilename {
			dir, _ := path.Split(pth)
			// todo add this info to task message
			glog.Infof("Found dataset %s", pth)
			go DatasetProcessor(dir)
		}
		return nil
	})

	if err != nil {
		glog.Errorf("Directory scan failed %s", err)
	}
}

func scanHttp(path string) {

}