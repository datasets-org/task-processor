package main

import (
	"github.com/datasets-org/task-processor/storage"
	"github.com/datasets-org/task-processor/managers"
	"github.com/datasets-org/task-processor/processors"
	"flag"
	"github.com/datasets-org/task-processor/structs"
	"time"
	"math/rand"
	"fmt"
	"os"
	"path"
)

func main() {
	flag.Set("stderrthreshold", "INFO")
	flag.Parse()

	ms := storage.CreateSyncMapStorage()
	//ms.Put("ID-0", `{"id": "ID-0", "task": {"operation": "dummy"}, "created": "2018-04-01T13:07:59.1234", "completed": false}`)
	t := managers.Tasks{Storage: ms}
	//t.Store(structs.Task{Id:"ID-0", Task: structs.TaskDefinition{Operation: "dummy"}})
	dir, _ := os.Getwd()
	path := path.Join(dir, "data")
	t.Store(structs.Task{Id:"ID-0", Task: structs.TaskDefinition{Operation: "scan", Params: fmt.Sprintf(`{"fs": "local", "path": "%s"}`, path)}})

	go taskCreator(&t)
	processors.Reactor(&t)
}

func taskCreator(tasks *managers.Tasks) {
	for {
		var op string
		if rand.Intn(10) > 3 {
			op = "dummy"
		} else {
			op = "random"
		}
		tasks.Store(structs.Task{
			Id: fmt.Sprintf("ID-%d", rand.Intn(100)),
			Task: structs.TaskDefinition{Operation: op}})
		time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
	}
}
