package transport

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	mapservice "github.com/zelta-7/cache/pkg/service/map"
	queueservice "github.com/zelta-7/cache/pkg/service/queue"
	"k8s.io/klog/v2"
)

type CacheHandlerInterface interface {
	// SetQueueValue sets a value in the queue
	SetQueueValue(c *gin.Context)

	// GetQueueValue gets a value from the queue
	GetQueueValue(c *gin.Context)

	// GetAllQueueValues gets all the values from the queue
	GetAllQueueValues(c *gin.Context)

	// GetQueueEntryList gets the first n entries from the queue
	GetQueueEntryList(c *gin.Context)

	// GetSortedQueueEntries gets the first n entries from the queue sorted by key or value
	GetSortedQueueEntries(c *gin.Context)

	// UpdateQueueValue updates the value of a given key in the queue
	UpdateQueueValue(c *gin.Context)

	// SetMapValue sets a value in the map
	SetMapValue(c *gin.Context)

	// GetMapValue gets a value from the map
	GetMapValue(c *gin.Context)

	// GetAllMapValues gets all the values from the map
	GetAllMapValues(c *gin.Context)

	// GetMapEntryList gets the first n entries from the map
	GetMapEntryList(c *gin.Context)

	// GetSortedMapEntries gets the first n entries from the map sorted by key or value
	GetSortedMapEntries(c *gin.Context)

	// UpdateMapEntry updates the value of a given key in the map
	UpdateMapEntry(c *gin.Context)

	// GetListofMapValues gets the values from a given list of key of the map
	GetListofMapValues(c *gin.Context)

	// TODO: Implement SetCacheTimetoLive for queue and map
}

type SetRequest struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type cacheHandler struct {
	cacheMapService   mapservice.MapServiceInterface
	cacheQueueService queueservice.QueueServiceInterface
}

func NewCacheHandler(mapservice mapservice.MapServiceInterface, queueservice queueservice.QueueServiceInterface) CacheHandlerInterface {
	return &cacheHandler{
		cacheMapService:   mapservice,
		cacheQueueService: queueservice,
	}
}

// SetQueueValue implements the SetQueueValue method of the CacheHandlerInterface
func (handler *cacheHandler) SetQueueValue(c *gin.Context) {
	var setValue SetRequest
	err := c.BindJSON(&setValue)
	if err != nil {
		klog.ErrorS(err, "Error binding json")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	key := handler.cacheQueueService.Set(setValue.Key, setValue.Value)

	c.JSON(http.StatusOK, gin.H{"status": "value set sucessfully", "key": key})
}

// GetQueueValue implements the GetQueueValue method of the CacheHandlerInterface
func (handler *cacheHandler) GetQueueValue(c *gin.Context) {
	key, value := handler.cacheQueueService.Get()
	klog.Info("Key: ", key, " Value: ", value)

	c.JSON(http.StatusOK, gin.H{"key": key, "value": value})
}

// GetAllQueueValues implements the GetAllQueueValues method of the CacheHandlerInterface
func (handler *cacheHandler) GetAllQueueValues(c *gin.Context) {
	values := handler.cacheQueueService.All()
	klog.Info("Values: ", values)
	c.JSON(http.StatusOK, gin.H{"values": values})
}

// GetQueueEntryList implements the GetQueueEntryList method of the CacheHandlerInterface
func (handler *cacheHandler) GetQueueEntryList(c *gin.Context) {
	n, err := strconv.Atoi(c.Param("n"))
	if err != nil {
		klog.ErrorS(err, "Error converting string to int")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	list := handler.cacheQueueService.GetEntryList(n)
	klog.Info("List: ", list)
	c.JSON(http.StatusOK, gin.H{"list": list})
}

// GetSortedQueueEntries implements the GetSortedQueueEntries method of the CacheHandlerInterface
func (handler *cacheHandler) GetSortedQueueEntries(c *gin.Context) {
	selector, err := strconv.Atoi(c.Param("selector"))
	if err != nil {
		klog.ErrorS(err, "Error converting string to int")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	n, err := strconv.Atoi(c.Param("n"))
	if err != nil {
		klog.ErrorS(err, "Error converting string to int")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	list := handler.cacheQueueService.GetSortedEntries(selector, n)
	klog.Info("List: ", list)
	c.JSON(http.StatusOK, gin.H{"list": list})
}

// UpdateQueueValue implements the UpdateQueueValue method of the CacheHandlerInterface
func (handler *cacheHandler) UpdateQueueValue(c *gin.Context) {
	key := c.Param("key")
	newValue := c.Param("newValue")

	responseKey := handler.cacheQueueService.UpdateValue(key, newValue)
	klog.Info("Key:", responseKey)
	c.JSON(http.StatusOK, gin.H{"key": responseKey})
}

// SetMapValue implements the SetMapValue method of the CacheHandlerInterface
func (handler *cacheHandler) SetMapValue(c *gin.Context) {
	var setValue SetRequest
	err := c.BindJSON(&setValue)
	if err != nil {
		klog.ErrorS(err, "Error binding json")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	key := handler.cacheMapService.Set(setValue.Key, setValue.Value)
	klog.Info("Key: ", key)
	c.JSON(http.StatusOK, gin.H{"status": "value set sucessfully", "key": key})
}

// GetMapValue implements the GetMapValue method of the CacheHandlerInterface
func (handler *cacheHandler) GetMapValue(c *gin.Context) {
	key := c.Param("key")
	value := handler.cacheMapService.Get(key)

	klog.Info("Key: ", key, " Value: ", value)
	c.JSON(http.StatusOK, gin.H{"key": key, "value": value})
}

// GetAllMapValues implements the GetAllMapValues method of the CacheHandlerInterface
func (handler *cacheHandler) GetAllMapValues(c *gin.Context) {
	values := handler.cacheMapService.All()

	klog.Info("Values: ", values)
	c.JSON(http.StatusOK, gin.H{"values": values})
}

// GetMapEntryList implements the GetMapEntryList method of the CacheHandlerInterface
func (handler *cacheHandler) GetMapEntryList(c *gin.Context) {
	n, err := strconv.Atoi(c.Param("n"))
	if err != nil {
		klog.ErrorS(err, "Error converting string to int")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	values := handler.cacheMapService.GetEntryList(n)
	klog.Info("Values: ", values)
	c.JSON(http.StatusOK, gin.H{"values": values})
}

// GetSortedMapEntries implements the GetSortedMapEntries method of the CacheHandlerInterface
func (handler *cacheHandler) GetSortedMapEntries(c *gin.Context) {
	selector, err := strconv.Atoi(c.Param("selector"))
	if err != nil {
		klog.ErrorS(err, "Error converting string to int")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	n, err := strconv.Atoi(c.Param("n"))
	if err != nil {
		klog.ErrorS(err, "Error converting string to int")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	values := handler.cacheMapService.GetSortedEntryList(selector, n)

	klog.Info("Values: ", values)
	c.JSON(http.StatusOK, gin.H{"values": values})
}

// UpdateMapEntry implements the UpdateMapEntry method of the CacheHandlerInterface
func (handler *cacheHandler) UpdateMapEntry(c *gin.Context) {
	key := c.Param("key")
	newValue := c.Param("newValue")

	responseKey := handler.cacheMapService.UpdateCacheEntry(key, newValue)

	klog.Info("Key:", responseKey)
	c.JSON(http.StatusOK, gin.H{"key": responseKey})
}

// GetListofMapValues implements the GetListofMapValues method of the CacheHandlerInterface
func (handler *cacheHandler) GetListofMapValues(c *gin.Context) {
	keys := c.QueryArray("keys")
	values := handler.cacheMapService.GetListofValues(keys)

	klog.Info("Values: ", values)
	c.JSON(http.StatusOK, gin.H{"values": values})
}
