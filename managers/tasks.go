package managers

import (
	"github.com/datasets-org/task-processor/structs"
	"github.com/datasets-org/task-processor/storage"
	"encoding/json"
	"github.com/golang/glog"
	"errors"
	"fmt"
	"sort"
)

type Tasks struct {
	Storage storage.Storage
}

func (t *Tasks) list() (tasks []structs.Task) {
	for _, i := range t.Storage.Items() {
		task := structs.CreateTask()
		err := json.Unmarshal([]byte(i), &task)
		if err != nil {
			glog.Errorf("Task cannot be parsed\n%s", i)
		} else {
			tasks = append(tasks, task)
		}
	}
	return
}

func (t *Tasks) GetActiveTasks() (tasks []structs.Task) {
	for _, i := range t.list() {
		if i.Completed == false {
			tasks = append(tasks, i)
		}
	}
	sort.Slice(tasks, func(i, j int) bool {
		return tasks[i].Created.Time.After(tasks[j].Created.Time)
	})
	return
}

func (t *Tasks) GetCompleteTasks() (tasks []structs.Task) {
	for _, i := range t.list() {
		if i.Completed == true {
			tasks = append(tasks, i)
		}
	}
	sort.Slice(tasks, func(i, j int) bool {
		return tasks[i].Created.Time.After(tasks[j].Created.Time)
	})
	return
}

func (t *Tasks) Get(key string) (structs.Task, error) {
	v, err := t.Storage.Get(key)
	task := structs.CreateTask()
	if err != nil {
		glog.Errorf("Task %s not found", key)
		return task, errors.New(fmt.Sprintf("Task %s not found", key))
	}
	err = json.Unmarshal([]byte(v), &task)
	if err != nil {
		glog.Errorf("Task cannot be parsed\n%s", v)
	}
	return task, nil
}

func (t *Tasks) Store(task structs.Task) {
	taskJson, err := json.Marshal(task)
	if err != nil {
		glog.Errorf("Task %s cannot be serialized", task.Id)
	} else {
		t.Storage.Update(task.Id, string(taskJson))
	}
}
