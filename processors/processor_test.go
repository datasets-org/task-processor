package processors

import (
	"testing"
	"github.com/datasets-org/task-processor/managers"
	"github.com/datasets-org/task-processor/storage"
	"github.com/datasets-org/task-processor/structs"
)

func TestProcessor(t *testing.T) {
	ms := storage.CreateMapStorage()
	tm := managers.Tasks{Storage: ms}
	task1 := structs.Task{Id: "ID-0", Task: structs.TaskDefinition{Operation: "dummy"}}
	Process(&task1, &tm)
	t1, _ := tm.Get("ID-0")
	if t1.Completed == false {
		t.Error("Task not processed")
	}
	if t1.Success == false {
		t.Error("Dummy processor failed")
	}
	task2 := structs.Task{Id: "ID-1", Task: structs.TaskDefinition{Operation: "unknown"}}
	Process(&task2, &tm)
	t2, _ := tm.Get("ID-1")
	if t2.Completed == false {
		t.Error("Task not processed")
	}
	if t2.Success == true {
		t.Error("Uknown processor succeeded")
	}
	if t2.Message != "Unknown operation unknown" {
		t.Error("message for unknown processor does not match")
	}
}
