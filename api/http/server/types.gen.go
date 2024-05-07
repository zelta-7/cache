// Package apiSpec provides primitives to interact with the openapi HTTP API.
//
// Code generated by unknown module path version unknown version DO NOT EDIT.
package apiSpec

// Defines values for GetCacheListSortNParamsSort.
const (
	Asc  GetCacheListSortNParamsSort = "asc"
	Desc GetCacheListSortNParamsSort = "desc"
)

// CacheEntry defines model for CacheEntry.
type CacheEntry struct {
	Key        string `json:"key"`
	TimeToLive *int   `json:"time-to-live,omitempty"`
	Value      string `json:"value"`
}

// CacheEntryResponse defines model for CacheEntryResponse.
type CacheEntryResponse struct {
	Hkey *int `json:"hkey,omitempty"`
}

// GetEntry defines model for GetEntry.
type GetEntry struct {
	Value string `json:"value"`
}

// GetEntryList defines model for GetEntryList.
type GetEntryList struct {
	Entries []CacheEntry `json:"entries"`
}

// GetCacheEntriesParams defines parameters for GetCacheEntries.
type GetCacheEntriesParams struct {
	// Key List of keys
	Key []string `form:"key" json:"key"`
}

// GetCacheListSortNParamsSort defines parameters for GetCacheListSortN.
type GetCacheListSortNParamsSort string

// PutCacheKeyJSONBody defines parameters for PutCacheKey.
type PutCacheKeyJSONBody struct {
	NewVal     *string `json:"new-val,omitempty"`
	TimeToLive *int    `json:"time-to-live,omitempty"`
}

// PostCacheJSONRequestBody defines body for PostCache for application/json ContentType.
type PostCacheJSONRequestBody = CacheEntry

// PutCacheKeyJSONRequestBody defines body for PutCacheKey for application/json ContentType.
type PutCacheKeyJSONRequestBody PutCacheKeyJSONBody

// PostCacheTimeToLiveJSONRequestBody defines body for PostCacheTimeToLive for application/json ContentType.
type PostCacheTimeToLiveJSONRequestBody = CacheEntry
