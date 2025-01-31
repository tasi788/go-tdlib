package tdlib

// StorageStatisticsByChat Contains the storage usage statistics for a specific chat
type StorageStatisticsByChat struct {
	tdCommon
	ChatId     int64                         `json:"chat_id"`      // Chat identifier; 0 if none
	Size       int64                         `json:"size"`         // Total size of the files in the chat, in bytes
	Count      int32                         `json:"count"`        // Total number of files in the chat
	ByFileType []StorageStatisticsByFileType `json:"by_file_type"` // Statistics split by file types
}

// MessageType return the string telegram-type of StorageStatisticsByChat
func (storageStatisticsByChat *StorageStatisticsByChat) MessageType() string {
	return "storageStatisticsByChat"
}

// NewStorageStatisticsByChat creates a new StorageStatisticsByChat
//
// @param chatId Chat identifier; 0 if none
// @param size Total size of the files in the chat, in bytes
// @param count Total number of files in the chat
// @param byFileType Statistics split by file types
func NewStorageStatisticsByChat(chatId int64, size int64, count int32, byFileType []StorageStatisticsByFileType) *StorageStatisticsByChat {
	storageStatisticsByChatTemp := StorageStatisticsByChat{
		tdCommon:   tdCommon{Type: "storageStatisticsByChat"},
		ChatId:     chatId,
		Size:       size,
		Count:      count,
		ByFileType: byFileType,
	}

	return &storageStatisticsByChatTemp
}
