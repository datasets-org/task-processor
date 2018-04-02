package structs

import (
	"errors"
	"fmt"
	"github.com/simplereach/timeutils"
	"time"
	"encoding/json"
	"github.com/golang/glog"
	"reflect"
)

type Dataset struct {
	Id         string   `yaml:"id" json:"id"`
	Data       []string `yaml:"data" json:"data"`
	Name       string   `yaml:"name" json:"name"`
	Internal   bool     `yaml:"internal" json:"internal"`
	From       string   `yaml:"from_ds" json:"from_ds"`
	Url        string   `yaml:"url" json:"url"`
	Maintainer string   `yaml:"maintainer" json:"maintainer"`
	Tags       []string `yaml:"tags" json:"tags"`
	// todo usage type struct
	Usages []map[string]string `yaml:"usages" json:"tags"`
	// todo changelog type struct
	Changelog [][][]string `yaml:"changelog" json:"changelog"`
	Markdowns []string     `yaml:"markdowns" json:"markdowns"`
	// todo characteristics type struct
	Characteristics map[string]string `yaml:"characteristics" json:"characteristics"`
	Links           []string          `yaml:"links" json:"links"`
	Path            []string          `yaml:"path" json:"path"`
	Type            string            `yaml:"type" json:"type"`
	Servers         []string          `yaml:"servers" json:"servers"`
}

type Change struct {
	Field string
	From  string
	To    string
}

type Changes struct {
	Created timeutils.Time
	Changes []Change
}

func (ds *Dataset) Diff(other Dataset) error {
	if ds.Id != other.Id {
		return errors.New(fmt.Sprintf("Generating diff for non-matching datasets %s != %s", ds.Id, other.Id))
	}
	changes := Changes{Created: timeutils.NewTime(time.Now(), timeutils.RFC3339)}
	if !reflect.DeepEqual(ds.Data, other.Data) {
		fromData, err1 := json.Marshal(ds.Data)
		if err1 != nil {
			glog.Errorf("Data not serializable in %s (%v)", ds.Id, ds.Data)
		}
		toData, err2 := json.Marshal(other.Data)
		if err2 != nil {
			glog.Errorf("Data not serializable in %s (%v)", ds.Id, ds.Data)
		}
		if err1 == nil && err2 == nil {
			changes.Changes = append(changes.Changes, Change{Field: "data", From: string(fromData), To: string(toData)})
		}
	}
	// todo compare other fields
	return nil
}

func (ds *Dataset) Use(usage map[string]string) {
	// todo add date
	ds.Usages = append(ds.Usages, usage)
}
