package repository

import (
	"sync"
	"time"

	"github.com/zelta-7/cache/common"
)

type QueueRepoInterface interface {
	// Set adds a value to the queue
	Set(key, value string, ttl ...int)

	// Get retrives the first value from the queue
	Get() (key, value string)

	// Update the value of a given key
	Update(key, value string) string

	// All returns all the values in the queue
	All() []CacheEntry
}

type CacheEntry struct {
	Value string
	Key   string
	TTL   time.Duration
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
func (q *QueueRepo) Set(value, key string, ttl ...int) {
	q.lock.Lock()
	defer q.lock.Unlock()

	if len(ttl) > 0 {
		q.queueCache = append(q.queueCache, CacheEntry{Value: value, Key: key, TTL: time.Duration(ttl[0])})
		return
	} else {
		q.queueCache = append(q.queueCache, CacheEntry{Value: value, Key: key})
		return
	}
}

// Get implements the Get method of the QueueRepoInterface
func (q *QueueRepo) Get() (string, string) {
	q.lock.Lock()
	defer q.lock.Unlock()

	if len(q.queueCache) == 0 {
		return "", ""
	}

	result := q.queueCache[0]
	// q.queueCache = q.queueCache[1:]
	return result.Key, result.Value
}

func (q *QueueRepo) Update(key, value string) string {
	q.lock.Lock()
	defer q.lock.Unlock()

	hashedKey := common.HashKey(key)
	for i, entry := range q.queueCache {
		if entry.Key == hashedKey {
			q.queueCache[i].Value = value
			break
		}
	}
	return key
}

// All implements the All method of the QueueRepoInterface
func (q *QueueRepo) All() []CacheEntry {
	q.lock.Lock()
	defer q.lock.Unlock()

	return q.queueCache
}
