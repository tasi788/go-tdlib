package tdlib

import (
	"encoding/json"
	"fmt"
)

// StorageStatistics Contains the exact storage usage statistics split by chats and file type
type StorageStatistics struct {
	tdCommon
	Size   int64                     `json:"size"`    // Total size of files, in bytes
	Count  int32                     `json:"count"`   // Total number of files
	ByChat []StorageStatisticsByChat `json:"by_chat"` // Statistics split by chats
}

// MessageType return the string telegram-type of StorageStatistics
func (storageStatistics *StorageStatistics) MessageType() string {
	return "storageStatistics"
}

// NewStorageStatistics creates a new StorageStatistics
//
// @param size Total size of files, in bytes
// @param count Total number of files
// @param byChat Statistics split by chats
func NewStorageStatistics(size int64, count int32, byChat []StorageStatisticsByChat) *StorageStatistics {
	storageStatisticsTemp := StorageStatistics{
		tdCommon: tdCommon{Type: "storageStatistics"},
		Size:     size,
		Count:    count,
		ByChat:   byChat,
	}

	return &storageStatisticsTemp
}

// GetStorageStatistics Returns storage usage statistics. Can be called before authorization
// @param chatLimit The maximum number of chats with the largest storage usage for which separate statistics need to be returned. All other chats will be grouped in entries with chat_id == 0. If the chat info database is not used, the chat_limit is ignored and is always set to 0
func (client *Client) GetStorageStatistics(chatLimit int32) (*StorageStatistics, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":      "getStorageStatistics",
		"chat_limit": chatLimit,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var storageStatistics StorageStatistics
	err = json.Unmarshal(result.Raw, &storageStatistics)
	return &storageStatistics, err
}

// OptimizeStorage Optimizes storage usage, i.e. deletes some files and returns new storage usage statistics. Secret thumbnails can't be deleted
// @param size Limit on the total size of files after deletion, in bytes. Pass -1 to use the default limit
// @param ttl Limit on the time that has passed since the last time a file was accessed (or creation time for some filesystems). Pass -1 to use the default limit
// @param count Limit on the total count of files after deletion. Pass -1 to use the default limit
// @param immunityDelay The amount of time after the creation of a file during which it can't be deleted, in seconds. Pass -1 to use the default value
// @param fileTypes If non-empty, only files with the given types are considered. By default, all types except thumbnails, profile photos, stickers and wallpapers are deleted
// @param chatIds If non-empty, only files from the given chats are considered. Use 0 as chat identifier to delete files not belonging to any chat (e.g., profile photos)
// @param excludeChatIds If non-empty, files from the given chats are excluded. Use 0 as chat identifier to exclude all files not belonging to any chat (e.g., profile photos)
// @param returnDeletedFileStatistics Pass true if statistics about the files that were deleted must be returned instead of the whole storage usage statistics. Affects only returned statistics
// @param chatLimit Same as in getStorageStatistics. Affects only returned statistics
func (client *Client) OptimizeStorage(size int64, ttl int32, count int32, immunityDelay int32, fileTypes []FileType, chatIds []int64, excludeChatIds []int64, returnDeletedFileStatistics bool, chatLimit int32) (*StorageStatistics, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":                          "optimizeStorage",
		"size":                           size,
		"ttl":                            ttl,
		"count":                          count,
		"immunity_delay":                 immunityDelay,
		"file_types":                     fileTypes,
		"chat_ids":                       chatIds,
		"exclude_chat_ids":               excludeChatIds,
		"return_deleted_file_statistics": returnDeletedFileStatistics,
		"chat_limit":                     chatLimit,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var storageStatistics StorageStatistics
	err = json.Unmarshal(result.Raw, &storageStatistics)
	return &storageStatistics, err
}
