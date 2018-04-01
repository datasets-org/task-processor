package storage

import (
	"errors"
	"sync"
)

type SyncMapStorage struct {
	data sync.Map
}

func CreateSyncMapStorage() SyncMapStorage {
	ms := SyncMapStorage{}
	ms.data.Store("0", "0")
	ms.data.Delete("0")
	return ms
}

func (ms SyncMapStorage) Put(key string, data string) {
	ms.data.Store(key, data)
}

func (ms SyncMapStorage) Update(key string, data string) {
	ms.data.Store(key, data)
}

func (ms SyncMapStorage) Delete(key string) {
	ms.data.Delete(key)
}

func (ms SyncMapStorage) Get(key string) (string, error) {
	data, ok := ms.data.Load(key)
	if !ok {
		return "", errors.New("key does not exist")
	}
	return data.(string), nil
}

func (ms SyncMapStorage) Items() (data []string) {
	ms.data.Range(func(_, value interface{}) bool {
		data = append(data, value.(string))
		return true
	})
	return
}
