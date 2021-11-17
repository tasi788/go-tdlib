package tdlib

import (
	"encoding/json"
	"fmt"
)

// ChatInviteLinkCounts Contains a list of chat invite link counts
type ChatInviteLinkCounts struct {
	tdCommon
	InviteLinkCounts []ChatInviteLinkCount `json:"invite_link_counts"` // List of invite linkcounts
}

// MessageType return the string telegram-type of ChatInviteLinkCounts
func (chatInviteLinkCounts *ChatInviteLinkCounts) MessageType() string {
	return "chatInviteLinkCounts"
}

// NewChatInviteLinkCounts creates a new ChatInviteLinkCounts
//
// @param inviteLinkCounts List of invite linkcounts
func NewChatInviteLinkCounts(inviteLinkCounts []ChatInviteLinkCount) *ChatInviteLinkCounts {
	chatInviteLinkCountsTemp := ChatInviteLinkCounts{
		tdCommon:         tdCommon{Type: "chatInviteLinkCounts"},
		InviteLinkCounts: inviteLinkCounts,
	}

	return &chatInviteLinkCountsTemp
}

// GetChatInviteLinkCounts Returns list of chat administrators with number of their invite links. Requires owner privileges in the chat
// @param chatId Chat identifier
func (client *Client) GetChatInviteLinkCounts(chatId int64) (*ChatInviteLinkCounts, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "getChatInviteLinkCounts",
		"chat_id": chatId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var chatInviteLinkCounts ChatInviteLinkCounts
	err = json.Unmarshal(result.Raw, &chatInviteLinkCounts)
	return &chatInviteLinkCounts, err
}
