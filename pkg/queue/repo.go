package repository

import "sync"

type QueueRepoInterface interface {
	// Set adds a value to the queue
	Set(value int) []int
	// Get retrives the first value from the queue
	Get() int
	// All returns all the values in the queue
	All() []int
}

type QueueRepo struct {
	queueCache []int
	lock       sync.RWMutex
}

func NewQueueRepo() QueueRepoInterface {
	return &QueueRepo{
		queueCache: make([]int, 0),
		lock:       sync.RWMutex{},
	}
}

// Set implements the Set method of the QueueRepoInterface
func (q *QueueRepo) Set(value int) []int {
	q.lock.Lock()
	defer q.lock.Unlock()

	q.queueCache = append(q.queueCache, value)
	return q.queueCache
}

// Get implements the Get method of the QueueRepoInterface
func (q *QueueRepo) Get() int {
	q.lock.Lock()
	defer q.lock.Unlock()

	if len(q.queueCache) == 0 {
		return -1
	}

	value := q.queueCache[0]
	q.queueCache = q.queueCache[1:]
	return value
}

// All implements the All method of the QueueRepoInterface
func (q *QueueRepo) All() []int {
	q.lock.Lock()
	defer q.lock.Unlock()

	return q.queueCache
}
