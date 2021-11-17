package tdlib

import (
	"encoding/json"
	"fmt"
)

// ChatInviteLinkMembers Contains a list of chat members joined a chat by an invite link
type ChatInviteLinkMembers struct {
	tdCommon
	TotalCount int32                  `json:"total_count"` // Approximate total count of chat members found
	Members    []ChatInviteLinkMember `json:"members"`     // List of chat members, joined a chat by an invite link
}

// MessageType return the string telegram-type of ChatInviteLinkMembers
func (chatInviteLinkMembers *ChatInviteLinkMembers) MessageType() string {
	return "chatInviteLinkMembers"
}

// NewChatInviteLinkMembers creates a new ChatInviteLinkMembers
//
// @param totalCount Approximate total count of chat members found
// @param members List of chat members, joined a chat by an invite link
func NewChatInviteLinkMembers(totalCount int32, members []ChatInviteLinkMember) *ChatInviteLinkMembers {
	chatInviteLinkMembersTemp := ChatInviteLinkMembers{
		tdCommon:   tdCommon{Type: "chatInviteLinkMembers"},
		TotalCount: totalCount,
		Members:    members,
	}

	return &chatInviteLinkMembersTemp
}

// GetChatInviteLinkMembers Returns chat members joined a chat by an invite link. Requires administrator privileges and can_invite_users right in the chat for own links and owner privileges for other links
// @param chatId Chat identifier
// @param inviteLink Invite link for which to return chat members
// @param offsetMember A chat member from which to return next chat members; pass null to get results from the beginning
// @param limit The maximum number of chat members to return; up to 100
func (client *Client) GetChatInviteLinkMembers(chatId int64, inviteLink string, offsetMember *ChatInviteLinkMember, limit int32) (*ChatInviteLinkMembers, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":         "getChatInviteLinkMembers",
		"chat_id":       chatId,
		"invite_link":   inviteLink,
		"offset_member": offsetMember,
		"limit":         limit,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var chatInviteLinkMembers ChatInviteLinkMembers
	err = json.Unmarshal(result.Raw, &chatInviteLinkMembers)
	return &chatInviteLinkMembers, err
}
