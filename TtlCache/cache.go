package ttlcache

type TtlCache interface {
	Set(key string, value interface{}, ttl int)
	Get(key string)
	Delete(key string)
}
