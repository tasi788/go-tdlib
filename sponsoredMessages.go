package tdlib

import (
	"encoding/json"
	"fmt"
)

// SponsoredMessages Contains a list of sponsored messages
type SponsoredMessages struct {
	tdCommon
	Messages []SponsoredMessage `json:"messages"` // List of sponsored messages
}

// MessageType return the string telegram-type of SponsoredMessages
func (sponsoredMessages *SponsoredMessages) MessageType() string {
	return "sponsoredMessages"
}

// NewSponsoredMessages creates a new SponsoredMessages
//
// @param messages List of sponsored messages
func NewSponsoredMessages(messages []SponsoredMessage) *SponsoredMessages {
	sponsoredMessagesTemp := SponsoredMessages{
		tdCommon: tdCommon{Type: "sponsoredMessages"},
		Messages: messages,
	}

	return &sponsoredMessagesTemp
}

// GetChatSponsoredMessages Returns sponsored messages to be shown in a chat; for channel chats only
// @param chatId Identifier of the chat
func (client *Client) GetChatSponsoredMessages(chatId int64) (*SponsoredMessages, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "getChatSponsoredMessages",
		"chat_id": chatId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var sponsoredMessages SponsoredMessages
	err = json.Unmarshal(result.Raw, &sponsoredMessages)
	return &sponsoredMessages, err
}
