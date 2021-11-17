package tdlib

import (
	"encoding/json"
	"fmt"
)

// Count Contains a counter
type Count struct {
	tdCommon
	Count int32 `json:"count"` // Count
}

// MessageType return the string telegram-type of Count
func (count *Count) MessageType() string {
	return "count"
}

// NewCount creates a new Count
//
// @param count Count
func NewCount(count int32) *Count {
	countTemp := Count{
		tdCommon: tdCommon{Type: "count"},
		Count:    count,
	}

	return &countTemp
}

// GetChatMessageCount Returns approximate number of messages of the specified type in the chat
// @param chatId Identifier of the chat in which to count messages
// @param filter Filter for message content; searchMessagesFilterEmpty is unsupported in this function
// @param returnLocal If true, returns count that is available locally without sending network requests, returning -1 if the number of messages is unknown
func (client *Client) GetChatMessageCount(chatId int64, filter SearchMessagesFilter, returnLocal bool) (*Count, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":        "getChatMessageCount",
		"chat_id":      chatId,
		"filter":       filter,
		"return_local": returnLocal,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var count Count
	err = json.Unmarshal(result.Raw, &count)
	return &count, err
}

// GetFileDownloadedPrefixSize Returns file downloaded prefix size from a given offset, in bytes
// @param fileId Identifier of the file
// @param offset Offset from which downloaded prefix size needs to be calculated
func (client *Client) GetFileDownloadedPrefixSize(fileId int32, offset int32) (*Count, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "getFileDownloadedPrefixSize",
		"file_id": fileId,
		"offset":  offset,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var count Count
	err = json.Unmarshal(result.Raw, &count)
	return &count, err
}

// GetImportedContactCount Returns the total number of imported contacts
func (client *Client) GetImportedContactCount() (*Count, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "getImportedContactCount",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var count Count
	err = json.Unmarshal(result.Raw, &count)
	return &count, err
}
