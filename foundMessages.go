package tdlib

import (
	"encoding/json"
	"fmt"
)

// FoundMessages Contains a list of messages found by a search
type FoundMessages struct {
	tdCommon
	TotalCount int32     `json:"total_count"` // Approximate total count of messages found; -1 if unknown
	Messages   []Message `json:"messages"`    // List of messages
	NextOffset string    `json:"next_offset"` // The offset for the next request. If empty, there are no more results
}

// MessageType return the string telegram-type of FoundMessages
func (foundMessages *FoundMessages) MessageType() string {
	return "foundMessages"
}

// NewFoundMessages creates a new FoundMessages
//
// @param totalCount Approximate total count of messages found; -1 if unknown
// @param messages List of messages
// @param nextOffset The offset for the next request. If empty, there are no more results
func NewFoundMessages(totalCount int32, messages []Message, nextOffset string) *FoundMessages {
	foundMessagesTemp := FoundMessages{
		tdCommon:   tdCommon{Type: "foundMessages"},
		TotalCount: totalCount,
		Messages:   messages,
		NextOffset: nextOffset,
	}

	return &foundMessagesTemp
}

// SearchSecretMessages Searches for messages in secret chats. Returns the results in reverse chronological order. For optimal performance, the number of returned messages is chosen by TDLib
// @param chatId Identifier of the chat in which to search. Specify 0 to search in all secret chats
// @param query Query to search for. If empty, searchChatMessages must be used instead
// @param offset Offset of the first entry to return as received from the previous request; use empty string to get first chunk of results
// @param limit The maximum number of messages to be returned; up to 100. For optimal performance, the number of returned messages is chosen by TDLib and can be smaller than the specified limit
// @param filter Additional filter for messages to search; pass null to search for all messages
func (client *Client) SearchSecretMessages(chatId int64, query string, offset string, limit int32, filter SearchMessagesFilter) (*FoundMessages, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "searchSecretMessages",
		"chat_id": chatId,
		"query":   query,
		"offset":  offset,
		"limit":   limit,
		"filter":  filter,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var foundMessages FoundMessages
	err = json.Unmarshal(result.Raw, &foundMessages)
	return &foundMessages, err
}

// GetMessagePublicForwards Returns forwarded copies of a channel message to different public channels. For optimal performance, the number of returned messages is chosen by TDLib
// @param chatId Chat identifier of the message
// @param messageId Message identifier
// @param offset Offset of the first entry to return as received from the previous request; use empty string to get first chunk of results
// @param limit The maximum number of messages to be returned; must be positive and can't be greater than 100. For optimal performance, the number of returned messages is chosen by TDLib and can be smaller than the specified limit
func (client *Client) GetMessagePublicForwards(chatId int64, messageId int64, offset string, limit int32) (*FoundMessages, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":      "getMessagePublicForwards",
		"chat_id":    chatId,
		"message_id": messageId,
		"offset":     offset,
		"limit":      limit,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var foundMessages FoundMessages
	err = json.Unmarshal(result.Raw, &foundMessages)
	return &foundMessages, err
}
