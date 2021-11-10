package tdlib

import (
	"encoding/json"
	"fmt"
)

// ChatInviteLinks Contains a list of chat invite links
type ChatInviteLinks struct {
	tdCommon
	TotalCount  int32            `json:"total_count"`  // Approximate total count of chat invite links found
	InviteLinks []ChatInviteLink `json:"invite_links"` // List of invite links
}

// MessageType return the string telegram-type of ChatInviteLinks
func (chatInviteLinks *ChatInviteLinks) MessageType() string {
	return "chatInviteLinks"
}

// NewChatInviteLinks creates a new ChatInviteLinks
//
// @param totalCount Approximate total count of chat invite links found
// @param inviteLinks List of invite links
func NewChatInviteLinks(totalCount int32, inviteLinks []ChatInviteLink) *ChatInviteLinks {
	chatInviteLinksTemp := ChatInviteLinks{
		tdCommon:    tdCommon{Type: "chatInviteLinks"},
		TotalCount:  totalCount,
		InviteLinks: inviteLinks,
	}

	return &chatInviteLinksTemp
}

// GetChatInviteLinks Returns invite links for a chat created by specified administrator. Requires administrator privileges and can_invite_users right in the chat to get own links and owner privileges to get other links
// @param chatId Chat identifier
// @param creatorUserId User identifier of a chat administrator. Must be an identifier of the current user for non-owner
// @param isRevoked Pass true if revoked links needs to be returned instead of active or expired
// @param offsetDate Creation date of an invite link starting after which to return invite links; use 0 to get results from the beginning
// @param offsetInviteLink Invite link starting after which to return invite links; use empty string to get results from the beginning
// @param limit The maximum number of invite links to return; up to 100
func (client *Client) GetChatInviteLinks(chatId int64, creatorUserId int64, isRevoked bool, offsetDate int32, offsetInviteLink string, limit int32) (*ChatInviteLinks, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":              "getChatInviteLinks",
		"chat_id":            chatId,
		"creator_user_id":    creatorUserId,
		"is_revoked":         isRevoked,
		"offset_date":        offsetDate,
		"offset_invite_link": offsetInviteLink,
		"limit":              limit,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var chatInviteLinks ChatInviteLinks
	err = json.Unmarshal(result.Raw, &chatInviteLinks)
	return &chatInviteLinks, err
}

// RevokeChatInviteLink Revokes invite link for a chat. Available for basic groups, supergroups, and channels. Requires administrator privileges and can_invite_users right in the chat for own links and owner privileges for other links. If a primary link is revoked, then additionally to the revoked link returns new primary link
// @param chatId Chat identifier
// @param inviteLink Invite link to be revoked
func (client *Client) RevokeChatInviteLink(chatId int64, inviteLink string) (*ChatInviteLinks, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":       "revokeChatInviteLink",
		"chat_id":     chatId,
		"invite_link": inviteLink,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var chatInviteLinks ChatInviteLinks
	err = json.Unmarshal(result.Raw, &chatInviteLinks)
	return &chatInviteLinks, err
}
