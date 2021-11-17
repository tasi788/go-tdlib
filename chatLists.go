package tdlib

import (
	"encoding/json"
	"fmt"
)

// ChatLists Contains a list of chat lists
type ChatLists struct {
	tdCommon
	ChatLists []ChatList `json:"chat_lists"` // List of chat lists
}

// MessageType return the string telegram-type of ChatLists
func (chatLists *ChatLists) MessageType() string {
	return "chatLists"
}

// NewChatLists creates a new ChatLists
//
// @param chatLists List of chat lists
func NewChatLists(chatLists []ChatList) *ChatLists {
	chatListsTemp := ChatLists{
		tdCommon:  tdCommon{Type: "chatLists"},
		ChatLists: chatLists,
	}

	return &chatListsTemp
}

// GetChatListsToAddChat Returns chat lists to which the chat can be added. This is an offline request
// @param chatId Chat identifier
func (client *Client) GetChatListsToAddChat(chatId int64) (*ChatLists, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "getChatListsToAddChat",
		"chat_id": chatId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var chatLists ChatLists
	err = json.Unmarshal(result.Raw, &chatLists)
	return &chatLists, err
}
