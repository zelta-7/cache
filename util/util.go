package util

import (
	"crypto/sha256"
	"encoding/hex"
	"sort"

	repository "github.com/zelta-7/cache/pkg/repository/queue"
)

func SortQueueByValue(q []repository.CacheEntry) []repository.CacheEntry {
	sort.Slice(q, func(i, j int) bool {
		return q[i].Value < q[j].Value
	})
	return q
}

func SortQueueByKey(q []repository.CacheEntry) []repository.CacheEntry {
	sort.Slice(q, func(i, j int) bool {
		return q[i].Key < q[j].Key
	})
	return q
}

func HashKey(key string) string {
	hasher := sha256.New()
	hasher.Write([]byte(key))
	return hex.EncodeToString(hasher.Sum(nil))
}
