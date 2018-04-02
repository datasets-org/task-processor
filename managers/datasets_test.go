package managers

import (
	"testing"
	"github.com/datasets-org/task-processor/storage"
	"github.com/datasets-org/task-processor/structs"
)

func TestDatasets(t *testing.T) {
	ds := Datasets{Storage: storage.CreateMapStorage()}
	if len(ds.list()) > 0 {
		t.Error("Datasets not empty")
	}
	ds.Store(structs.Dataset{Id: "0"})
	if len(ds.list()) != 1 {
		t.Error("Dataset not stored")
	}
	_, err := ds.Get("-1")
	if err == nil {
		t.Error("Get for non existent key passes")
	}
	task, err := ds.Get("0")
	if err != nil {
		t.Error("Dataset not stored")
	}
	if task.Id != "0" {
		t.Error("Dataset malformed")
	}
	ds.Store(structs.Dataset{Id: "1"})
	ds.Store(structs.Dataset{Id: "2"})
	if len(ds.list()) != 3 {
		t.Error("Datasets not stored")
	}
}
