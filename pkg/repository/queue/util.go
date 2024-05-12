package repository

import "sort"

func SortQueueByValue(q []CacheEntry) []CacheEntry {
	sort.Slice(q, func(i, j int) bool {
		return q[i].Value < q[j].Value
	})
	return q
}

func SortQueueByKey(q []CacheEntry) []CacheEntry {
	sort.Slice(q, func(i, j int) bool {
		return q[i].Key < q[j].Key
	})
	return q
}
