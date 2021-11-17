package tdlib

import (
	"encoding/json"
	"fmt"
)

// MessagePositions Contains a list of message positions
type MessagePositions struct {
	tdCommon
	TotalCount int32             `json:"total_count"` // Total count of messages found
	Positions  []MessagePosition `json:"positions"`   // List of message positions
}

// MessageType return the string telegram-type of MessagePositions
func (messagePositions *MessagePositions) MessageType() string {
	return "messagePositions"
}

// NewMessagePositions creates a new MessagePositions
//
// @param totalCount Total count of messages found
// @param positions List of message positions
func NewMessagePositions(totalCount int32, positions []MessagePosition) *MessagePositions {
	messagePositionsTemp := MessagePositions{
		tdCommon:   tdCommon{Type: "messagePositions"},
		TotalCount: totalCount,
		Positions:  positions,
	}

	return &messagePositionsTemp
}

// GetChatSparseMessagePositions Returns sparse positions of messages of the specified type in the chat to be used for shared media scroll implementation. Returns the results in reverse chronological order (i.e., in order of decreasing message_id). Cannot be used in secret chats or with searchMessagesFilterFailedToSend filter without an enabled message database
// @param chatId Identifier of the chat in which to return information about message positions
// @param filter Filter for message content. Filters searchMessagesFilterEmpty, searchMessagesFilterCall, searchMessagesFilterMissedCall, searchMessagesFilterMention and searchMessagesFilterUnreadMention are unsupported in this function
// @param fromMessageId The message identifier from which to return information about message positions
// @param limit The expected number of message positions to be returned; 50-2000. A smaller number of positions can be returned, if there are not enough appropriate messages
func (client *Client) GetChatSparseMessagePositions(chatId int64, filter SearchMessagesFilter, fromMessageId int64, limit int32) (*MessagePositions, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":           "getChatSparseMessagePositions",
		"chat_id":         chatId,
		"filter":          filter,
		"from_message_id": fromMessageId,
		"limit":           limit,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var messagePositions MessagePositions
	err = json.Unmarshal(result.Raw, &messagePositions)
	return &messagePositions, err
}
