package managers

import (
	"testing"
	"github.com/datasets-org/task-processor/storage"
	"github.com/datasets-org/task-processor/structs"
	"time"
)

func TestTasks(t *testing.T) {
	tasks := Tasks{Storage: storage.CreateMapStorage()}
	if len(tasks.list()) > 0 {
		t.Error("Tasks not empty")
	}
	if len(tasks.GetActiveTasks()) > 0 {
		t.Error("Tasks not empty")
	}
	if len(tasks.GetCompleteTasks()) > 0 {
		t.Error("Tasks not empty")
	}
	tasks.Store(structs.Task{Id: "0"})
	if len(tasks.list()) != 1 {
		t.Error("Task not stored")
	}
	if len(tasks.GetActiveTasks()) != 1 {
		t.Error("Added task is not active")
	}
	if len(tasks.GetCompleteTasks()) > 0 {
		t.Error("Complete Tasks not empty")
	}
	_, err := tasks.Get("-1")
	if err == nil {
		t.Error("Get for non existent ket passes")
	}
	task, err := tasks.Get("0")
	if err != nil {
		t.Error("Task not stored")
	}
	if task.Id != "0" {
		t.Error("Task malformed")
	}
	task.Complete(true, "")
	tasks.Store(task)
	if len(tasks.list()) != 1 {
		t.Error("Task not stored")
	}
	if len(tasks.GetActiveTasks()) > 0 {
		t.Error("Task already complete")
	}
	if len(tasks.GetCompleteTasks()) != 1 {
		t.Error("Complete task not found")
	}
	task = structs.Task{Id: "1"}
	task.Created.Time = time.Now()
	tasks.Store(task)
	task = structs.Task{Id: "2"}
	task.Created.Time = time.Now().Add(time.Minute)
	tasks.Store(task)
	if len(tasks.list()) != 3 {
		t.Error("Tasks not stored")
	}
	if len(tasks.GetActiveTasks()) != 2 {
		t.Error("2 active tasks not present")
	}
	if len(tasks.GetCompleteTasks()) != 1 {
		t.Error("Complete task not found")
	}
	t1, _ := tasks.Get("1")
	t2, _ := tasks.Get("2")
	if t2.Id != tasks.GetActiveTasks()[0].Id {
		t.Error("Tasks not ordered")
	}
	if t1.Id != tasks.GetActiveTasks()[1].Id {
		t.Error("Tasks not ordered")
	}
	t2.Complete(false, "failed")
	tasks.Store(t2)
	if t2.Id != tasks.GetCompleteTasks()[0].Id {
		t.Error("Completed tasks not ordered")
	}
	t2, _ = tasks.Get("2")
	if t2.Message != "failed" {
		t.Error("Task fail not stored")
	}
}
