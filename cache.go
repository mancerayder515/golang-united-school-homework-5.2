package cache

import "time"

type Cache struct {
	data     map[string]string
	tempKeys map[string]time.Time
}

func NewCache() Cache {
	return Cache{data: map[string]string{}, tempKeys: map[string]time.Time{}}
}

func (c *Cache) Get(key string) (string, bool) {
	c.clean()
	value, ok := c.data[key]
	return value, ok
}

func (c *Cache) Put(key, value string) {
	if _, ok := c.tempKeys[key]; ok {
		delete(c.tempKeys, key)
	}
	c.data[key] = value
}

func (c *Cache) Keys() []string {
	c.clean()
	keys := make([]string, len(c.data))
	for k, _ := range c.data {
		keys = append(keys, k)
	}
	return keys
}

func (c *Cache) PutTill(key, value string, deadline time.Time) {
	c.tempKeys[key] = deadline
	c.data[key] = value
}

func (c *Cache) clean() {
	for k, deadline := range c.tempKeys {
		if deadline.Before(time.Now()) {
			delete(c.data, k)
			delete(c.tempKeys, k)
		}
	}
}
