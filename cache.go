package cache

import "time"

type Cache struct {
	items             map[string]*CacheItem
	defaultExpiration time.Duration
}

type CacheItem struct {
	KeyValue   interface{}
	Expiration *time.Time
}

func New(defaultExpiration, cleanInterval time.Duration) *Cache {
	c := &Cache{
		items:             map[string]*CacheItem{},
		defaultExpiration: defaultExpiration,
	}
	if cleanInterval > 0 {
		go func() {
			for {
				time.Sleep(cleanInterval)
				c.DeleteExpired()
			}
		}()
	}
	return c
}

func (item *CacheItem) Expired() bool {
	if item.Expiration == nil {
		return false
	}
	return item.Expiration.Before(time.Now())
}

func (c *Cache) Set(key string, value interface{}, duration time.Duration) {
	var timeCache *time.Time
	if duration == 0 {
		duration = c.defaultExpiration
	}
	c.items[key] = &CacheItem{
		KeyValue:   value,
		Expiration: timeCache,
	}
}

func (c *Cache) Get(key string) interface{} {
	return c.items[key]
}

func (c *Cache) Delete(key string) {
	delete(c.items, key)
}

func (c *Cache) DeleteExpired() {
	for k, v := range c.items {
		if v.Expired() {
			delete(c.items, k)
		}
	}
}
