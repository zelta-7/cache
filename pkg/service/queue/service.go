package service

import (
	"time"

	"github.com/zelta-7/cache/common"
	repository "github.com/zelta-7/cache/pkg/repository/queue"
)

type QueueServiceInterface interface {
	// Set adds a value to the queue
	Set(key, value string) string

	// Get retrives the first value from the queue
	Get() (key, value string)

	// All returns all the values in the queue
	All() []repository.CacheEntry

	// GetEntryList returns the first n entries in the queue
	GetEntryList(n int) []repository.CacheEntry

	// GetSortedEntries returns the first n entries in the queue sorted by key or value
	GetSortedEntries(selector, n int) []repository.CacheEntry

	// UpdateValue updates the value of a given key
	UpdateValue(key, newValue string) string

	// SetCacheTimetoLive adds a value to the queue with a time to live
	SetCacheTimetoLive(key, value string, ttl ...int) string
}

type queueService struct {
	queueInterface repository.QueueRepoInterface
}

func NewQueueService(queueRepoInter repository.QueueRepoInterface) QueueServiceInterface {
	return &queueService{
		queueInterface: queueRepoInter,
	}
}

// Set implements the Set method of the QueueServiceInterface
func (q *queueService) Set(key, value string) string {
	hashedKey := common.HashKey(key)
	q.queueInterface.Set(hashedKey, value)
	return key
}

// Get implements the Get method of the QueueServiceInterface
func (q *queueService) Get() (string, string) {
	key, value := q.queueInterface.Get()
	return key, value
}

// All implements the All method of the QueueServiceInterface
func (q *queueService) All() []repository.CacheEntry {
	return q.queueInterface.All()
}

// GetEntryList implements the GetEntryList method of the QueueServiceInterface
func (q *queueService) GetEntryList(n int) []repository.CacheEntry {
	result := []repository.CacheEntry{}
	for _, entry := range q.queueInterface.All() {
		result = append(result, entry)
		n--
		if n == 0 {
			break
		}
	}
	return result
}

// GetSortedEntries implements the GetSortedEntries method of the QueueServiceInterface
func (q *queueService) GetSortedEntries(selector, n int) []repository.CacheEntry {
	if selector != 1 && selector != 0 {
		return nil
	}
	if selector == 1 {
		return repository.SortQueueByKey(q.GetEntryList(n))
	} else {
		return repository.SortQueueByValue(q.GetEntryList(n))
	}
}

// UpdateValue implements the UpdateValue method of the QueueServiceInterface
func (q *queueService) UpdateValue(key, newValue string) string {
	hashedKey := common.HashKey(key)
	q.queueInterface.Update(hashedKey, newValue)
	return key
}

// SetCacheTimetoLive implements the SetCacheTimetoLive method of the QueueServiceInterface
func (q *queueService) SetCacheTimetoLive(key, value string, ttl ...int) string {
	q.queueInterface.Set(key, value, ttl[0])
	time.NewTimer(time.Duration(ttl[0]) * time.Second)
	return key
}
