package structs

import (
	"github.com/simplereach/timeutils"
)

type Task struct {
	Id        string         `json:"id"`
	Completed bool           `json:"completed"`
	Created   timeutils.Time `json:"created"`
	Task      string         `json:"task"`
	Message   string         `json:"message"`
	Success   bool           `json:"success"`
	Finished  timeutils.Time `json:"finished"`
}

func CreateTask() Task {
	t := Task{}
	t.Created = t.Created.FormatMode(timeutils.RFC3339)
	t.Finished = t.Finished.FormatMode(timeutils.RFC3339)
	return t
}
