package repository

import "sort"

func SortMapByKey(m map[string]string) map[string]string {
	var keys []string
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	sortedMap := make(map[string]string, len(keys))
	for _, k := range keys {
		sortedMap[k] = m[k]
	}
	return sortedMap
}

func SortMapByValue(m map[string]string) map[string]string {
	var keys []string
	for k := range m {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool {
		return m[keys[i]] < m[keys[j]]
	})

	sortedMap := make(map[string]string, len(keys))
	for _, k := range keys {
		sortedMap[k] = m[k]
	}
	return sortedMap
}
