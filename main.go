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
)

func main() {
	flag.Set("stderrthreshold", "INFO")
	flag.Parse()

	ms := storage.CreateMapStorage()
	//ms.Put("ID-0", `{"id": "ID-0", "task": {"operation": "dummy"}, "created": "2018-04-01T13:07:59.1234", "completed": false}`)
	t := managers.Tasks{Storage: ms}
	//t.Store(structs.Task{Id:"ID-0", Task: structs.TaskDefinition{Operation: "dummy"}})

	//for _, i := range t.GetActiveTasks() {
	//	fmt.Println(i.Id)
	//	fmt.Println(i.Completed)
	//	v, _ := json.Marshal(i)
	//	fmt.Println(string(v))
	//	processors.Process(&i, &t)
	//	v, _ = json.Marshal(i)
	//	fmt.Println(string(v))
	//}

	go taskCreator(&t)
	processors.Reactor(&t)
}

func taskCreator(tasks *managers.Tasks) {
	for {
		tasks.Store(structs.Task{Id: fmt.Sprintf("ID-%d", rand.Intn(100)), Task: structs.TaskDefinition{Operation: "dummy"}})
		time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
	}
}