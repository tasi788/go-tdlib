package tdlib

import (
	"encoding/json"
	"fmt"
)

// MessageLink Contains an HTTPS link to a message in a supergroup or channel
type MessageLink struct {
	tdCommon
	Link     string `json:"link"`      // Message link
	IsPublic bool   `json:"is_public"` // True, if the link will work for non-members of the chat
}

// MessageType return the string telegram-type of MessageLink
func (messageLink *MessageLink) MessageType() string {
	return "messageLink"
}

// NewMessageLink creates a new MessageLink
//
// @param link Message link
// @param isPublic True, if the link will work for non-members of the chat
func NewMessageLink(link string, isPublic bool) *MessageLink {
	messageLinkTemp := MessageLink{
		tdCommon: tdCommon{Type: "messageLink"},
		Link:     link,
		IsPublic: isPublic,
	}

	return &messageLinkTemp
}

// GetMessageLink Returns an HTTPS link to a message in a chat. Available only for already sent messages in supergroups and channels, or if message.can_get_media_timestamp_links and a media timestamp link is generated. This is an offline request
// @param chatId Identifier of the chat to which the message belongs
// @param messageId Identifier of the message
// @param mediaTimestamp If not 0, timestamp from which the video/audio/video note/voice note playing must start, in seconds. The media can be in the message content or in its web page preview
// @param forAlbum Pass true to create a link for the whole media album
// @param forComment Pass true to create a link to the message as a channel post comment, or from a message thread
func (client *Client) GetMessageLink(chatId int64, messageId int64, mediaTimestamp int32, forAlbum bool, forComment bool) (*MessageLink, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":           "getMessageLink",
		"chat_id":         chatId,
		"message_id":      messageId,
		"media_timestamp": mediaTimestamp,
		"for_album":       forAlbum,
		"for_comment":     forComment,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var messageLink MessageLink
	err = json.Unmarshal(result.Raw, &messageLink)
	return &messageLink, err
}
