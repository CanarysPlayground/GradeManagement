package cache

// Cache defines minimal cache interface
type Cache interface {
	Get(key string) (string, error)
	Set(key string, value string) error
}

// NewRedisCache returns a simple stub cache that ignores redisAddr for now
func NewRedisCache(redisAddr string) Cache {
	return &noopCache{}
}

type noopCache struct{}

func (n *noopCache) Get(key string) (string, error) { return "", nil }
func (n *noopCache) Set(key string, value string) error { return nil }
