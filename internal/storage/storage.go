package storage

type Storage interface {
	Save(key string, value interface{}) error
	Load(key string, result interface{}) error
	Delete(key string) error
	ListKeys() ([]string, error)
}
