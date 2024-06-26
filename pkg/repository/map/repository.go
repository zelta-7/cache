package repository

import (
	"sync"
)

type MapRepoInter interface {
	// Set sets the value of the key
	Set(key, value string)

	// Get returns the value of the key
	Get(key string) string

	// UpdateValue updates the value of the key
	UpdateValue(key, newValue string) string

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

// Set implements the Set method of the MapRepoInter interface
func (m *MapRepo) Set(key, value string) {
	m.lock.Lock()
	defer m.lock.Unlock()

	m.MapCache[key] = value
}

// Get implements the Get method of the MapRepoInter interface
func (m *MapRepo) Get(key string) string {
	m.lock.Lock()
	defer m.lock.Unlock()

	return m.MapCache[key]
}

// UpdateValue implements the UpdateValue method of the MapRepoInter interface
func (m *MapRepo) UpdateValue(key, newValue string) string {
	m.lock.Lock()
	defer m.lock.Unlock()

	m.MapCache[key] = newValue
	return key
}

// All implements the GetEntryList method of the MapRepoInter interface
func (m *MapRepo) All() map[string]string {
	m.lock.Lock()
	defer m.lock.Unlock()

	return m.MapCache
}
