package tdlib

import (
	"encoding/json"
	"fmt"
)

// ChatInviteLink Contains a chat invite link
type ChatInviteLink struct {
	tdCommon
	InviteLink              string `json:"invite_link"`                // Chat invite link
	Name                    string `json:"name"`                       // Name of the link
	CreatorUserId           int64  `json:"creator_user_id"`            // User identifier of an administrator created the link
	Date                    int32  `json:"date"`                       // Point in time (Unix timestamp) when the link was created
	EditDate                int32  `json:"edit_date"`                  // Point in time (Unix timestamp) when the link was last edited; 0 if never or unknown
	ExpireDate              int32  `json:"expire_date"`                // Point in time (Unix timestamp) when the link will expire; 0 if never
	MemberLimit             int32  `json:"member_limit"`               // The maximum number of members, which can join the chat using the link simultaneously; 0 if not limited. Always 0 if the link requires approval
	MemberCount             int32  `json:"member_count"`               // Number of chat members, which joined the chat using the link
	PendingJoinRequestCount int32  `json:"pending_join_request_count"` // Number of pending join requests created using this link
	CreatesJoinRequest      bool   `json:"creates_join_request"`       // True, if the link only creates join request. If true, total number of joining members will be unlimited
	IsPrimary               bool   `json:"is_primary"`                 // True, if the link is primary. Primary invite link can't have name, expire date or usage limit. There is exactly one primary invite link for each administrator with can_invite_users right at a given time
	IsRevoked               bool   `json:"is_revoked"`                 // True, if the link was revoked
}

// MessageType return the string telegram-type of ChatInviteLink
func (chatInviteLink *ChatInviteLink) MessageType() string {
	return "chatInviteLink"
}

// NewChatInviteLink creates a new ChatInviteLink
//
// @param inviteLink Chat invite link
// @param name Name of the link
// @param creatorUserId User identifier of an administrator created the link
// @param date Point in time (Unix timestamp) when the link was created
// @param editDate Point in time (Unix timestamp) when the link was last edited; 0 if never or unknown
// @param expireDate Point in time (Unix timestamp) when the link will expire; 0 if never
// @param memberLimit The maximum number of members, which can join the chat using the link simultaneously; 0 if not limited. Always 0 if the link requires approval
// @param memberCount Number of chat members, which joined the chat using the link
// @param pendingJoinRequestCount Number of pending join requests created using this link
// @param createsJoinRequest True, if the link only creates join request. If true, total number of joining members will be unlimited
// @param isPrimary True, if the link is primary. Primary invite link can't have name, expire date or usage limit. There is exactly one primary invite link for each administrator with can_invite_users right at a given time
// @param isRevoked True, if the link was revoked
func NewChatInviteLink(inviteLink string, name string, creatorUserId int64, date int32, editDate int32, expireDate int32, memberLimit int32, memberCount int32, pendingJoinRequestCount int32, createsJoinRequest bool, isPrimary bool, isRevoked bool) *ChatInviteLink {
	chatInviteLinkTemp := ChatInviteLink{
		tdCommon:                tdCommon{Type: "chatInviteLink"},
		InviteLink:              inviteLink,
		Name:                    name,
		CreatorUserId:           creatorUserId,
		Date:                    date,
		EditDate:                editDate,
		ExpireDate:              expireDate,
		MemberLimit:             memberLimit,
		MemberCount:             memberCount,
		PendingJoinRequestCount: pendingJoinRequestCount,
		CreatesJoinRequest:      createsJoinRequest,
		IsPrimary:               isPrimary,
		IsRevoked:               isRevoked,
	}

	return &chatInviteLinkTemp
}

// ReplacePrimaryChatInviteLink Replaces current primary invite link for a chat with a new primary invite link. Available for basic groups, supergroups, and channels. Requires administrator privileges and can_invite_users right
// @param chatId Chat identifier
func (client *Client) ReplacePrimaryChatInviteLink(chatId int64) (*ChatInviteLink, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "replacePrimaryChatInviteLink",
		"chat_id": chatId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var chatInviteLink ChatInviteLink
	err = json.Unmarshal(result.Raw, &chatInviteLink)
	return &chatInviteLink, err
}

// CreateChatInviteLink Creates a new invite link for a chat. Available for basic groups, supergroups, and channels. Requires administrator privileges and can_invite_users right in the chat
// @param chatId Chat identifier
// @param name Invite link name; 0-32 characters
// @param expireDate Point in time (Unix timestamp) when the link will expire; pass 0 if never
// @param memberLimit The maximum number of chat members that can join the chat by the link simultaneously; 0-99999; pass 0 if not limited
// @param createsJoinRequest True, if the link only creates join request. If true, member_limit must not be specified
func (client *Client) CreateChatInviteLink(chatId int64, name string, expireDate int32, memberLimit int32, createsJoinRequest bool) (*ChatInviteLink, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":                "createChatInviteLink",
		"chat_id":              chatId,
		"name":                 name,
		"expire_date":          expireDate,
		"member_limit":         memberLimit,
		"creates_join_request": createsJoinRequest,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var chatInviteLink ChatInviteLink
	err = json.Unmarshal(result.Raw, &chatInviteLink)
	return &chatInviteLink, err
}

// EditChatInviteLink Edits a non-primary invite link for a chat. Available for basic groups, supergroups, and channels. Requires administrator privileges and can_invite_users right in the chat for own links and owner privileges for other links
// @param chatId Chat identifier
// @param inviteLink Invite link to be edited
// @param name Invite link name; 0-32 characters
// @param expireDate Point in time (Unix timestamp) when the link will expire; pass 0 if never
// @param memberLimit The maximum number of chat members that can join the chat by the link simultaneously; 0-99999; pass 0 if not limited
// @param createsJoinRequest True, if the link only creates join request. If true, member_limit must not be specified
func (client *Client) EditChatInviteLink(chatId int64, inviteLink string, name string, expireDate int32, memberLimit int32, createsJoinRequest bool) (*ChatInviteLink, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":                "editChatInviteLink",
		"chat_id":              chatId,
		"invite_link":          inviteLink,
		"name":                 name,
		"expire_date":          expireDate,
		"member_limit":         memberLimit,
		"creates_join_request": createsJoinRequest,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var chatInviteLink ChatInviteLink
	err = json.Unmarshal(result.Raw, &chatInviteLink)
	return &chatInviteLink, err
}

// GetChatInviteLink Returns information about an invite link. Requires administrator privileges and can_invite_users right in the chat to get own links and owner privileges to get other links
// @param chatId Chat identifier
// @param inviteLink Invite link to get
func (client *Client) GetChatInviteLink(chatId int64, inviteLink string) (*ChatInviteLink, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":       "getChatInviteLink",
		"chat_id":     chatId,
		"invite_link": inviteLink,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var chatInviteLink ChatInviteLink
	err = json.Unmarshal(result.Raw, &chatInviteLink)
	return &chatInviteLink, err
}
