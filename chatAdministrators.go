package tdlib

import (
	"encoding/json"
	"fmt"
)

// ChatAdministrators Represents a list of chat administrators
type ChatAdministrators struct {
	tdCommon
	Administrators []ChatAdministrator `json:"administrators"` // A list of chat administrators
}

// MessageType return the string telegram-type of ChatAdministrators
func (chatAdministrators *ChatAdministrators) MessageType() string {
	return "chatAdministrators"
}

// NewChatAdministrators creates a new ChatAdministrators
//
// @param administrators A list of chat administrators
func NewChatAdministrators(administrators []ChatAdministrator) *ChatAdministrators {
	chatAdministratorsTemp := ChatAdministrators{
		tdCommon:       tdCommon{Type: "chatAdministrators"},
		Administrators: administrators,
	}

	return &chatAdministratorsTemp
}

// GetChatAdministrators Returns a list of administrators of the chat with their custom titles
// @param chatId Chat identifier
func (client *Client) GetChatAdministrators(chatId int64) (*ChatAdministrators, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "getChatAdministrators",
		"chat_id": chatId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var chatAdministrators ChatAdministrators
	err = json.Unmarshal(result.Raw, &chatAdministrators)
	return &chatAdministrators, err
}
