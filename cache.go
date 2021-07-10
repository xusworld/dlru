package dcache

import "time"

// Cache cache is the interface of in-memory cache
type Cache interface {
	// Get returns single item from the backend if the requested item is not
	// found, returns NotFound err
	Get(key string) (interface{}, error)

	// Set set or update a key/value pair in in-memory cache
	Set(key string, value interface{})

	// SetWithExpire Set set or update a key/value pair in in-memory cache  with an expiration time
	SetWithExpire(key string, value interface{}, expiration time.Duration)

	// Delete deletes single item from backend
	Delete(key string) error

	// Len returns the number of items in cache
	Len() int

	// ForEach
	ForEach(func(key string, val interface{}))
}
