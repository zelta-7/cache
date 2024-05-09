package repository

import (
	"sync"
)

type QueueRepoInterface interface {
	// Set adds a value to the queue
	Set(key, value string)
	// Get retrives the first value from the queue
	Get() (key, value string)
	// All returns all the values in the queue
	All() []CacheEntry
}

type CacheEntry struct {
	Value string
	Key   string
}

type QueueRepo struct {
	queueCache []CacheEntry
	lock       sync.RWMutex
}

func NewQueueRepo() QueueRepoInterface {
	return &QueueRepo{
		queueCache: make([]CacheEntry, 0),
		lock:       sync.RWMutex{},
	}
}

// Set implements the Set method of the QueueRepoInterface
func (q *QueueRepo) Set(value, key string) {
	q.lock.Lock()
	defer q.lock.Unlock()

	q.queueCache = append(q.queueCache, CacheEntry{Value: value, Key: key})
}

// Get implements the Get method of the QueueRepoInterface
func (q *QueueRepo) Get() (string, string) {
	q.lock.Lock()
	defer q.lock.Unlock()

	if len(q.queueCache) == 0 {
		return "", ""
	}

	result := q.queueCache[0]
	q.queueCache = q.queueCache[1:]
	return result.Key, result.Value
}

// All implements the All method of the QueueRepoInterface
func (q *QueueRepo) All() []CacheEntry {
	q.lock.Lock()
	defer q.lock.Unlock()

	return q.queueCache
}
