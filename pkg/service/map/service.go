package service

import (
	"reflect"

	repository "github.com/zelta-7/cache/pkg/repository/map"
)

type MapServiceInterface interface {
	// Set sets the value of the key
	Set(key, value string)

	// Get returns the value of the key
	Get(key string) string

	// All returns all the entries in the map
	All() map[string]string

	// GetEntryList returns the first n entries in the map
	GetEntryList(n int) map[string]string

	// GetSortedEntryList returns the first n entries in the map sorted by value
	GetSortedEntryList(n int) map[string]string

	// GetCacheMetadata returns the type of the value stored in the key
	GetCacheMetadata(key string) string

	// GetCacheMetadataAll returns the type of all the values stored in the map
	GetCacheMetadataAll() []string

	// UpdateCacheEntry updates the value of the key
	UpdateCacheEntry(key, value string)

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
func (m *mapService) Set(key, value string) {
	m.mapInterface.Set(key, value)
}

// Get implements the Get method of the MapServiceInterface
func (m *mapService) Get(key string) string {
	return m.mapInterface.Get(key)
}

// All implements the All method of the MapServiceInterface
func (m *mapService) All() map[string]string {
	return m.mapInterface.All()
}

// GetEntryList implements the GetEntryList method of the MapServiceInterface
func (m *mapService) GetEntryList(n int) map[string]string {
	entryList := make(map[string]string)
	for key, value := range m.mapInterface.All() {
		entryList[key] = value
		n--
		if n == 0 {
			break
		}
	}
	return entryList
}

// TODO: Update GetSortedEntryList for sorting by keys, add functionality as required

// GetSortedEntryList implements the GetSortedEntryList method of the MapServiceInterface
func (m *mapService) GetSortedEntryList(n int) map[string]string {
	entryList := make(map[string]string)
	for key, value := range m.mapInterface.All() {
		entryList[key] = value
		n--
		if n == 0 {
			break
		}
	}
	// entryList = util.x(entryList)
	return entryList
}

// GetCacheMetadata implemets the GetCacheMetadata method of the MapServiceInterface
func (m *mapService) GetCacheMetadata(key string) string {
	value := m.mapInterface.Get(key)
	return reflect.TypeOf(value).String()
}

// GetCacheMetadataAll implements the GetCacheMetadataAll method of the MapServiceInterface
func (m *mapService) GetCacheMetadataAll() []string {
	metadata := make([]string, 0)
	for key, value := range m.mapInterface.All() {
		metadata = append(metadata, key+" : "+reflect.TypeOf(value).String())
	}
	return metadata
}

// UpdateCacheEntry implements the UpdateCacheEntry method of the MapServiceInterface
func (m *mapService) UpdateCacheEntry(key, value string) {
	m.mapInterface.Set(key, value)
}

// GetListofValues implements the GetListofValues method of the MapServiceInterface
func (m *mapService) GetListofValues(keys []string) []string {
	values := make([]string, 0)
	for _, key := range keys {
		values = append(values, m.mapInterface.Get(key))
	}
	return values
}

// TODO: Implement SetCacheTimetoLive function
// func (m *mapService) SetCacheTimetoLive(key string, ttl int)
