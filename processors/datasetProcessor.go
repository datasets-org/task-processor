package processors

import (
	"github.com/golang/glog"
	"gopkg.in/yaml.v2"
	"path"
	"github.com/datasets-org/task-processor/structs"
	"io/ioutil"
	"github.com/datasets-org/task-processor/managers"
	"github.com/datasets-org/task-processor/storage"
)

func DatasetProcessor(dsPath string) error {
	glog.Infof("Processing dataset %s", dsPath)
	dsFilePath := path.Join(dsPath, datasetFilename)

	ds := structs.Dataset{}
	dat, err := ioutil.ReadFile(dsFilePath)
	if err != nil {
		glog.Errorf("Dataset file (%s) cannot be read %s", dsFilePath, err)
		return err
	}
	err = yaml.Unmarshal(dat, &ds)
	glog.Info(ds)
	// todo param
	datasets := managers.Datasets{Storage: storage.SyncMapStorage{}}
	datasets.Store(ds)
	if err != nil {
		glog.Errorf("Dataset yaml malformed %s", err)
		return err
	}
	return nil
}
