package tdlib

import (
	"encoding/json"
	"fmt"
)

// MessageThreadInfo Contains information about a message thread
type MessageThreadInfo struct {
	tdCommon
	ChatId             int64             `json:"chat_id"`              // Identifier of the chat to which the message thread belongs
	MessageThreadId    int64             `json:"message_thread_id"`    // Message thread identifier, unique within the chat
	ReplyInfo          *MessageReplyInfo `json:"reply_info"`           // Information about the message thread
	UnreadMessageCount int32             `json:"unread_message_count"` // Approximate number of unread messages in the message thread
	Messages           []Message         `json:"messages"`             // The messages from which the thread starts. The messages are returned in a reverse chronological order (i.e., in order of decreasing message_id)
	DraftMessage       *DraftMessage     `json:"draft_message"`        // A draft of a message in the message thread; may be null
}

// MessageType return the string telegram-type of MessageThreadInfo
func (messageThreadInfo *MessageThreadInfo) MessageType() string {
	return "messageThreadInfo"
}

// NewMessageThreadInfo creates a new MessageThreadInfo
//
// @param chatId Identifier of the chat to which the message thread belongs
// @param messageThreadId Message thread identifier, unique within the chat
// @param replyInfo Information about the message thread
// @param unreadMessageCount Approximate number of unread messages in the message thread
// @param messages The messages from which the thread starts. The messages are returned in a reverse chronological order (i.e., in order of decreasing message_id)
// @param draftMessage A draft of a message in the message thread; may be null
func NewMessageThreadInfo(chatId int64, messageThreadId int64, replyInfo *MessageReplyInfo, unreadMessageCount int32, messages []Message, draftMessage *DraftMessage) *MessageThreadInfo {
	messageThreadInfoTemp := MessageThreadInfo{
		tdCommon:           tdCommon{Type: "messageThreadInfo"},
		ChatId:             chatId,
		MessageThreadId:    messageThreadId,
		ReplyInfo:          replyInfo,
		UnreadMessageCount: unreadMessageCount,
		Messages:           messages,
		DraftMessage:       draftMessage,
	}

	return &messageThreadInfoTemp
}

// GetMessageThread Returns information about a message thread. Can be used only if message.can_get_message_thread == true
// @param chatId Chat identifier
// @param messageId Identifier of the message
func (client *Client) GetMessageThread(chatId int64, messageId int64) (*MessageThreadInfo, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":      "getMessageThread",
		"chat_id":    chatId,
		"message_id": messageId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var messageThreadInfo MessageThreadInfo
	err = json.Unmarshal(result.Raw, &messageThreadInfo)
	return &messageThreadInfo, err
}
