package tdlib

import (
	"encoding/json"
	"fmt"
)

// MessageSenders Represents a list of message senders
type MessageSenders struct {
	tdCommon
	TotalCount int32           `json:"total_count"` // Approximate total count of messages senders found
	Senders    []MessageSender `json:"senders"`     // List of message senders
}

// MessageType return the string telegram-type of MessageSenders
func (messageSenders *MessageSenders) MessageType() string {
	return "messageSenders"
}

// NewMessageSenders creates a new MessageSenders
//
// @param totalCount Approximate total count of messages senders found
// @param senders List of message senders
func NewMessageSenders(totalCount int32, senders []MessageSender) *MessageSenders {
	messageSendersTemp := MessageSenders{
		tdCommon:   tdCommon{Type: "messageSenders"},
		TotalCount: totalCount,
		Senders:    senders,
	}

	return &messageSendersTemp
}

// GetChatAvailableMessageSenders Returns list of message sender identifiers, which can be used to send messages in a chat
// @param chatId Chat identifier
func (client *Client) GetChatAvailableMessageSenders(chatId int64) (*MessageSenders, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "getChatAvailableMessageSenders",
		"chat_id": chatId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var messageSenders MessageSenders
	err = json.Unmarshal(result.Raw, &messageSenders)
	return &messageSenders, err
}

// GetVideoChatAvailableParticipants Returns list of participant identifiers, on whose behalf a video chat in the chat can be joined
// @param chatId Chat identifier
func (client *Client) GetVideoChatAvailableParticipants(chatId int64) (*MessageSenders, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "getVideoChatAvailableParticipants",
		"chat_id": chatId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var messageSenders MessageSenders
	err = json.Unmarshal(result.Raw, &messageSenders)
	return &messageSenders, err
}

// GetBlockedMessageSenders Returns users and chats that were blocked by the current user
// @param offset Number of users and chats to skip in the result; must be non-negative
// @param limit The maximum number of users and chats to return; up to 100
func (client *Client) GetBlockedMessageSenders(offset int32, limit int32) (*MessageSenders, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":  "getBlockedMessageSenders",
		"offset": offset,
		"limit":  limit,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var messageSenders MessageSenders
	err = json.Unmarshal(result.Raw, &messageSenders)
	return &messageSenders, err
}
