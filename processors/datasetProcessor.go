package processors

import (
	"github.com/golang/glog"
	"gopkg.in/yaml.v2"
	"path"
	"github.com/datasets-org/task-processor/structs"
	"io/ioutil"
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
	if err != nil {
		glog.Errorf("Dataset yaml malformed %s", err)
		return err
	}
	return nil
}
