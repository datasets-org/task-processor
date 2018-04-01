package storage

import (
	"errors"
)

type MapStorage struct {
	data map[string]string
}

func CreateMapStorage() MapStorage {
	ms := MapStorage{}
	ms.data = make(map[string]string)
	return ms
}

func (ms MapStorage) Put(key string, data string) {
	ms.data[key] = data
}

func (ms MapStorage) Update(key string, data string) {
	ms.data[key] = data
}

func (ms MapStorage) Delete(key string) {
	delete(ms.data, key)
}

func (ms MapStorage) Get(key string) (string, error) {
	data, ok := ms.data[key]
	if !ok {
		return "", errors.New("key does not exist")
	}
	return data, nil
}

func (ms MapStorage) Items() (data []string) {
	for _, v := range ms.data {
		data = append(data, v)
	}
	return
}
