package storage

type Storage interface {
	Put(key string, data map[string]string)
	Update(key string, data map[string]string)
	Delete(key string)
	Get(key string) (map[string]string, error)
	Items() []map[string]string
}
