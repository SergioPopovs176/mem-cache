package cache

type Cache interface {
	Set(key string, value interface{})
	Get(key string)
	Delete(key string)
}
