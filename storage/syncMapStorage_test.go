package storage

import (
	"testing"
)

func TestSyncMapStorage(t *testing.T) {
	ms := CreateSyncMapStorage()
	items := ms.Items()
	if len(items) > 0 {
		t.Error("Storage not empty")
	}
	_, err := ms.Get("0")
	if err == nil {
		t.Error("Empty key retrieval")
	}
	ms.Put("1", "task")
	val, err := ms.Get("1")
	if err != nil {
		t.Error("key not found")
	}
	if len(val) < 1 {
		t.Error("corrupted data")
	}
	if val != "task" {
		t.Error("corrupted data content")
	}
	items = ms.Items()
	if len(items) != 1 {
		t.Error("Storage size missmatch")
	}
	ms.Put("2", "task2")
	items = ms.Items()
	if len(items) != 2 {
		t.Error("Storage size missmatch")
	}
	ms.Update("1", "task0")
	if len(items) != 2 {
		t.Error("Update appends")
	}
	val, _ = ms.Get("1")
	if val != "task0" {
		t.Error("Update corrupts data")
	}
	ms.Delete("1")
	items = ms.Items()
	if len(items) != 1 {
		t.Error("Delete not deleting")
	}
	_, err = ms.Get("0")
	if err == nil {
		t.Error("Key not delete")
	}
}
