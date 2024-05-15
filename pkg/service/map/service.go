package service

import (
	"github.com/zelta-7/cache/common"
	repository "github.com/zelta-7/cache/pkg/repository/map"
	"k8s.io/klog/v2"
)

type MapServiceInterface interface {
	// Set sets the value of the key
	Set(key, value string) string

	// Get returns the value of the key
	Get(key string) string

	// All returns all the entries in the map
	All() map[string]string

	// GetEntryList returns the first n entries in the map
	GetEntryList(n int) map[string]string

	// GetSortedEntryList returns the first n entries in the map sorted by value
	GetSortedEntryList(selector, n int) map[string]string

	// UpdateCacheEntry updates the value of the key
	UpdateCacheEntry(key, value string) string

	// GetListofValues returns the array of values for the given keys
	GetListofValues(keys []string) []string
}

type mapService struct {
	mapInterface repository.MapRepoInter
}

func NewMapService(mapRepoInterface repository.MapRepoInter) MapServiceInterface {
	return &mapService{
		mapInterface: mapRepoInterface,
	}
}

// Set implements the Set method of the MapServiceInterface
func (m *mapService) Set(key, value string) string {
	hashedKey := common.HashKey(key)
	m.mapInterface.Set(hashedKey, value)
	return key
}

// Get implements the Get method of the MapServiceInterface
func (m *mapService) Get(key string) string {
	hashedKey := common.HashKey(key)
	return m.mapInterface.Get(hashedKey)
}

// All implements the All method of the MapServiceInterface
func (m *mapService) All() map[string]string {
	return m.mapInterface.All()
}

// GetEntryList implements the GetEntryList method of the MapServiceInterface
func (m *mapService) GetEntryList(n int) map[string]string {
	entryList := make(map[string]string)
	for hashedKey, value := range m.mapInterface.All() {
		key, err := common.DecodeHashedKey(hashedKey)
		if err != nil {
			klog.ErrorS(err, "Error decoding hashed key", "key", hashedKey)
			continue
		}
		entryList[key] = value
		n--
		if n == 0 {
			break
		}
	}
	return entryList
}

// GetSortedEntryList implements the GetSortedEntryList method of the MapServiceInterface
func (m *mapService) GetSortedEntryList(selector, n int) map[string]string {
	if selector != 0 && selector != 1 {
		klog.Warning("Invalid selector value")
		return nil
	}
	if selector == 0 {
		return repository.SortMapByValue(m.GetEntryList(n))
	}
	if selector == 1 {
		return repository.SortMapByKey(m.GetEntryList(n))
	}
	return nil
}

// UpdateCacheEntry implements the UpdateCacheEntry method of the MapServiceInterface
func (m *mapService) UpdateCacheEntry(key, value string) string {
	hashedKey := common.HashKey(key)
	m.mapInterface.UpdateValue(hashedKey, value)
	return key
}

// GetListofValues implements the GetListofValues method of the MapServiceInterface
func (m *mapService) GetListofValues(keys []string) []string {
	values := make([]string, 0)
	for _, hasedKey := range keys {
		key, err := common.DecodeHashedKey(hasedKey)
		if err != nil {
			klog.ErrorS(err, "Error decoding hashed key", "key", hasedKey)
			continue
		}
		values = append(values, m.mapInterface.Get(key))
	}
	return values
}

// TODO: Implement SetCacheTimetoLive function
// func (m *mapService) SetCacheTimetoLive(key string, ttl int)
