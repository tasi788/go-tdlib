package tdlib

import (
	"encoding/json"
	"fmt"
)

// ChatInviteLinkInfo Contains information about a chat invite link
type ChatInviteLinkInfo struct {
	tdCommon
	ChatId             int64          `json:"chat_id"`              // Chat identifier of the invite link; 0 if the user has no access to the chat before joining
	AccessibleFor      int32          `json:"accessible_for"`       // If non-zero, the amount of time for which read access to the chat will remain available, in seconds
	Type               ChatType       `json:"type"`                 // Type of the chat
	Title              string         `json:"title"`                // Title of the chat
	Photo              *ChatPhotoInfo `json:"photo"`                // Chat photo; may be null
	Description        string         `json:"description"`          // Chat description
	MemberCount        int32          `json:"member_count"`         // Number of members in the chat
	MemberUserIds      []int64        `json:"member_user_ids"`      // User identifiers of some chat members that may be known to the current user
	CreatesJoinRequest bool           `json:"creates_join_request"` // True, if the link only creates join request
	IsPublic           bool           `json:"is_public"`            // True, if the chat is a public supergroup or channel, i.e. it has a username or it is a location-based supergroup
}

// MessageType return the string telegram-type of ChatInviteLinkInfo
func (chatInviteLinkInfo *ChatInviteLinkInfo) MessageType() string {
	return "chatInviteLinkInfo"
}

// NewChatInviteLinkInfo creates a new ChatInviteLinkInfo
//
// @param chatId Chat identifier of the invite link; 0 if the user has no access to the chat before joining
// @param accessibleFor If non-zero, the amount of time for which read access to the chat will remain available, in seconds
// @param typeParam Type of the chat
// @param title Title of the chat
// @param photo Chat photo; may be null
// @param description Chat description
// @param memberCount Number of members in the chat
// @param memberUserIds User identifiers of some chat members that may be known to the current user
// @param createsJoinRequest True, if the link only creates join request
// @param isPublic True, if the chat is a public supergroup or channel, i.e. it has a username or it is a location-based supergroup
func NewChatInviteLinkInfo(chatId int64, accessibleFor int32, typeParam ChatType, title string, photo *ChatPhotoInfo, description string, memberCount int32, memberUserIds []int64, createsJoinRequest bool, isPublic bool) *ChatInviteLinkInfo {
	chatInviteLinkInfoTemp := ChatInviteLinkInfo{
		tdCommon:           tdCommon{Type: "chatInviteLinkInfo"},
		ChatId:             chatId,
		AccessibleFor:      accessibleFor,
		Type:               typeParam,
		Title:              title,
		Photo:              photo,
		Description:        description,
		MemberCount:        memberCount,
		MemberUserIds:      memberUserIds,
		CreatesJoinRequest: createsJoinRequest,
		IsPublic:           isPublic,
	}

	return &chatInviteLinkInfoTemp
}

// UnmarshalJSON unmarshal to json
func (chatInviteLinkInfo *ChatInviteLinkInfo) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	tempObj := struct {
		tdCommon
		ChatId             int64          `json:"chat_id"`              // Chat identifier of the invite link; 0 if the user has no access to the chat before joining
		AccessibleFor      int32          `json:"accessible_for"`       // If non-zero, the amount of time for which read access to the chat will remain available, in seconds
		Title              string         `json:"title"`                // Title of the chat
		Photo              *ChatPhotoInfo `json:"photo"`                // Chat photo; may be null
		Description        string         `json:"description"`          // Chat description
		MemberCount        int32          `json:"member_count"`         // Number of members in the chat
		MemberUserIds      []int64        `json:"member_user_ids"`      // User identifiers of some chat members that may be known to the current user
		CreatesJoinRequest bool           `json:"creates_join_request"` // True, if the link only creates join request
		IsPublic           bool           `json:"is_public"`            // True, if the chat is a public supergroup or channel, i.e. it has a username or it is a location-based supergroup
	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	chatInviteLinkInfo.tdCommon = tempObj.tdCommon
	chatInviteLinkInfo.ChatId = tempObj.ChatId
	chatInviteLinkInfo.AccessibleFor = tempObj.AccessibleFor
	chatInviteLinkInfo.Title = tempObj.Title
	chatInviteLinkInfo.Photo = tempObj.Photo
	chatInviteLinkInfo.Description = tempObj.Description
	chatInviteLinkInfo.MemberCount = tempObj.MemberCount
	chatInviteLinkInfo.MemberUserIds = tempObj.MemberUserIds
	chatInviteLinkInfo.CreatesJoinRequest = tempObj.CreatesJoinRequest
	chatInviteLinkInfo.IsPublic = tempObj.IsPublic

	fieldType, _ := unmarshalChatType(objMap["type"])
	chatInviteLinkInfo.Type = fieldType

	return nil
}

// CheckChatInviteLink Checks the validity of an invite link for a chat and returns information about the corresponding chat
// @param inviteLink Invite link to be checked
func (client *Client) CheckChatInviteLink(inviteLink string) (*ChatInviteLinkInfo, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":       "checkChatInviteLink",
		"invite_link": inviteLink,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var chatInviteLinkInfo ChatInviteLinkInfo
	err = json.Unmarshal(result.Raw, &chatInviteLinkInfo)
	return &chatInviteLinkInfo, err
}
