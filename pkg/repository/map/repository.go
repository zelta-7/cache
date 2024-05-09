package repository

import (
	"sync"
)

type MapRepoInter interface {
	// Get returns the value of the key
	Get(key string) string
	// Set sets the value of the key
	Set(key, value string)
	// All return all the entries in the map
	All() map[string]string
}

type MapRepo struct {
	MapCache map[string]string
	lock     sync.RWMutex
}

func NewMapRepo() MapRepoInter {
	return &MapRepo{
		MapCache: make(map[string]string),
		lock:     sync.RWMutex{},
	}
}

// Get implements the Get method of the MapRepoInter interface
func (m *MapRepo) Get(key string) string {
	m.lock.Lock()
	defer m.lock.Lock()

	return m.MapCache[key]
}

// Set implements the Set method of the MapRepoInter interface
func (m *MapRepo) Set(key, value string) {
	m.lock.Lock()
	defer m.lock.Lock()

	m.MapCache[key] = value
}

// All implements the GetEntryList method of the MapRepoInter interface
func (m *MapRepo) All() map[string]string {
	m.lock.Lock()
	defer m.lock.Lock()

	return m.MapCache
}

// In Go:
// There are two fundamental types
// 1. Concrete types
// 	- struct
// 	- int
// 	- string
// 	- float64
// 	- Anything that has a memory representation

// 2. Abstract types
// 	- interface
// 	- Interfaces are implemented by concrete types