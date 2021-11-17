package tdlib

import (
	"encoding/json"
	"fmt"
)

// MessageLinkInfo Contains information about a link to a message in a chat
type MessageLinkInfo struct {
	tdCommon
	IsPublic       bool     `json:"is_public"`       // True, if the link is a public link for a message in a chat
	ChatId         int64    `json:"chat_id"`         // If found, identifier of the chat to which the message belongs, 0 otherwise
	Message        *Message `json:"message"`         // If found, the linked message; may be null
	MediaTimestamp int32    `json:"media_timestamp"` // Timestamp from which the video/audio/video note/voice note playing must start, in seconds; 0 if not specified. The media can be in the message content or in its web page preview
	ForAlbum       bool     `json:"for_album"`       // True, if the whole media album to which the message belongs is linked
	ForComment     bool     `json:"for_comment"`     // True, if the message is linked as a channel post comment or from a message thread
}

// MessageType return the string telegram-type of MessageLinkInfo
func (messageLinkInfo *MessageLinkInfo) MessageType() string {
	return "messageLinkInfo"
}

// NewMessageLinkInfo creates a new MessageLinkInfo
//
// @param isPublic True, if the link is a public link for a message in a chat
// @param chatId If found, identifier of the chat to which the message belongs, 0 otherwise
// @param message If found, the linked message; may be null
// @param mediaTimestamp Timestamp from which the video/audio/video note/voice note playing must start, in seconds; 0 if not specified. The media can be in the message content or in its web page preview
// @param forAlbum True, if the whole media album to which the message belongs is linked
// @param forComment True, if the message is linked as a channel post comment or from a message thread
func NewMessageLinkInfo(isPublic bool, chatId int64, message *Message, mediaTimestamp int32, forAlbum bool, forComment bool) *MessageLinkInfo {
	messageLinkInfoTemp := MessageLinkInfo{
		tdCommon:       tdCommon{Type: "messageLinkInfo"},
		IsPublic:       isPublic,
		ChatId:         chatId,
		Message:        message,
		MediaTimestamp: mediaTimestamp,
		ForAlbum:       forAlbum,
		ForComment:     forComment,
	}

	return &messageLinkInfoTemp
}

// GetMessageLinkInfo Returns information about a public or private message link. Can be called for any internal link of the type internalLinkTypeMessage
// @param url The message link
func (client *Client) GetMessageLinkInfo(url string) (*MessageLinkInfo, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "getMessageLinkInfo",
		"url":   url,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var messageLinkInfo MessageLinkInfo
	err = json.Unmarshal(result.Raw, &messageLinkInfo)
	return &messageLinkInfo, err
}
