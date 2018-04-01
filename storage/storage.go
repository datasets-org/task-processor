package storage

type Storage interface {
	Put(key string, data string)
	Update(key string, data string)
	Delete(key string)
	Get(key string) (string, error)
	Items() []string
}
