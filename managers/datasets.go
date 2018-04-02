package managers

import (
	"github.com/datasets-org/task-processor/storage"
	"github.com/datasets-org/task-processor/structs"
	"github.com/golang/glog"
	"encoding/json"
	"fmt"
	"errors"
)

type Datasets struct {
	Storage storage.Storage
}

func (t *Datasets) list() (ds []structs.Dataset) {
	for _, i := range t.Storage.Items() {
		dataset := structs.Dataset{}
		err := json.Unmarshal([]byte(i), &dataset)
		if err != nil {
			glog.Errorf("Dataset cannot be parsed\n%s", i)
		} else {
			ds = append(ds, dataset)
		}
	}
	return
}

func (t *Datasets) Get(key string) (structs.Dataset, error) {
	v, err := t.Storage.Get(key)
	dataset := structs.Dataset{}
	if err != nil {
		glog.Errorf("Dataset %s not found", key)
		return dataset, errors.New(fmt.Sprintf("Dataset %s not found", key))
	}
	err = json.Unmarshal([]byte(v), &dataset)
	if err != nil {
		glog.Errorf("Dataset cannot be parsed\n%s", v)
	}
	return dataset, nil
}

func (t *Datasets) Store(dataset structs.Dataset) {
	dsJson, err := json.Marshal(dataset)
	if err != nil {
		glog.Errorf("Dataset %s cannot be serialized", dataset.Id)
	} else {
		glog.Infof("%s", dsJson)
		t.Storage.Update(dataset.Id, string(dsJson))
	}
}
