package cache

type cacheValue struct {
	value any
}

type MemoryCahe struct {
	storage map[string]cacheValue
}

func (c *MemoryCahe) Set(key string, value interface{}) {
	c.storage[key] = cacheValue{value: value}
}

func (c *MemoryCahe) Get(key string) interface{} {
	return c.storage[key].value
}

func (c *MemoryCahe) Delete(key string) {
	delete(c.storage, key)
}

func New() *MemoryCahe {
	c := MemoryCahe{storage: make(map[string]cacheValue)}

	return &(c)
}
