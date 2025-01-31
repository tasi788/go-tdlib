package tdlib

import (
	"encoding/json"
	"fmt"
)

// ChatMemberStatus Provides information about the status of a member in a chat
type ChatMemberStatus interface {
	GetChatMemberStatusEnum() ChatMemberStatusEnum
}

// ChatMemberStatusEnum Alias for abstract ChatMemberStatus 'Sub-Classes', used as constant-enum here
type ChatMemberStatusEnum string

// ChatMemberStatus enums
const (
	ChatMemberStatusCreatorType       ChatMemberStatusEnum = "chatMemberStatusCreator"
	ChatMemberStatusAdministratorType ChatMemberStatusEnum = "chatMemberStatusAdministrator"
	ChatMemberStatusMemberType        ChatMemberStatusEnum = "chatMemberStatusMember"
	ChatMemberStatusRestrictedType    ChatMemberStatusEnum = "chatMemberStatusRestricted"
	ChatMemberStatusLeftType          ChatMemberStatusEnum = "chatMemberStatusLeft"
	ChatMemberStatusBannedType        ChatMemberStatusEnum = "chatMemberStatusBanned"
)

func unmarshalChatMemberStatus(rawMsg *json.RawMessage) (ChatMemberStatus, error) {

	if rawMsg == nil {
		return nil, nil
	}
	var objMap map[string]interface{}
	err := json.Unmarshal(*rawMsg, &objMap)
	if err != nil {
		return nil, err
	}

	switch ChatMemberStatusEnum(objMap["@type"].(string)) {
	case ChatMemberStatusCreatorType:
		var chatMemberStatusCreator ChatMemberStatusCreator
		err := json.Unmarshal(*rawMsg, &chatMemberStatusCreator)
		return &chatMemberStatusCreator, err

	case ChatMemberStatusAdministratorType:
		var chatMemberStatusAdministrator ChatMemberStatusAdministrator
		err := json.Unmarshal(*rawMsg, &chatMemberStatusAdministrator)
		return &chatMemberStatusAdministrator, err

	case ChatMemberStatusMemberType:
		var chatMemberStatusMember ChatMemberStatusMember
		err := json.Unmarshal(*rawMsg, &chatMemberStatusMember)
		return &chatMemberStatusMember, err

	case ChatMemberStatusRestrictedType:
		var chatMemberStatusRestricted ChatMemberStatusRestricted
		err := json.Unmarshal(*rawMsg, &chatMemberStatusRestricted)
		return &chatMemberStatusRestricted, err

	case ChatMemberStatusLeftType:
		var chatMemberStatusLeft ChatMemberStatusLeft
		err := json.Unmarshal(*rawMsg, &chatMemberStatusLeft)
		return &chatMemberStatusLeft, err

	case ChatMemberStatusBannedType:
		var chatMemberStatusBanned ChatMemberStatusBanned
		err := json.Unmarshal(*rawMsg, &chatMemberStatusBanned)
		return &chatMemberStatusBanned, err

	default:
		return nil, fmt.Errorf("Error UnMarshaling, unknown type:" + objMap["@type"].(string))
	}
}

// ChatMemberStatusCreator The user is the owner of the chat and has all the administrator privileges
type ChatMemberStatusCreator struct {
	tdCommon
	CustomTitle string `json:"custom_title"` // A custom title of the owner; 0-16 characters without emojis; applicable to supergroups only
	IsAnonymous bool   `json:"is_anonymous"` // True, if the creator isn't shown in the chat member list and sends messages anonymously; applicable to supergroups only
	IsMember    bool   `json:"is_member"`    // True, if the user is a member of the chat
}

// MessageType return the string telegram-type of ChatMemberStatusCreator
func (chatMemberStatusCreator *ChatMemberStatusCreator) MessageType() string {
	return "chatMemberStatusCreator"
}

// NewChatMemberStatusCreator creates a new ChatMemberStatusCreator
//
// @param customTitle A custom title of the owner; 0-16 characters without emojis; applicable to supergroups only
// @param isAnonymous True, if the creator isn't shown in the chat member list and sends messages anonymously; applicable to supergroups only
// @param isMember True, if the user is a member of the chat
func NewChatMemberStatusCreator(customTitle string, isAnonymous bool, isMember bool) *ChatMemberStatusCreator {
	chatMemberStatusCreatorTemp := ChatMemberStatusCreator{
		tdCommon:    tdCommon{Type: "chatMemberStatusCreator"},
		CustomTitle: customTitle,
		IsAnonymous: isAnonymous,
		IsMember:    isMember,
	}

	return &chatMemberStatusCreatorTemp
}

// GetChatMemberStatusEnum return the enum type of this object
func (chatMemberStatusCreator *ChatMemberStatusCreator) GetChatMemberStatusEnum() ChatMemberStatusEnum {
	return ChatMemberStatusCreatorType
}

// ChatMemberStatusAdministrator The user is a member of the chat and has some additional privileges. In basic groups, administrators can edit and delete messages sent by others, add new members, ban unprivileged members, and manage video chats. In supergroups and channels, there are more detailed options for administrator privileges
type ChatMemberStatusAdministrator struct {
	tdCommon
	CustomTitle         string `json:"custom_title"`           // A custom title of the administrator; 0-16 characters without emojis; applicable to supergroups only
	CanBeEdited         bool   `json:"can_be_edited"`          // True, if the current user can edit the administrator privileges for the called user
	CanManageChat       bool   `json:"can_manage_chat"`        // True, if the administrator can get chat event log, get chat statistics, get message statistics in channels, get channel members, see anonymous administrators in supergroups and ignore slow mode. Implied by any other privilege; applicable to supergroups and channels only
	CanChangeInfo       bool   `json:"can_change_info"`        // True, if the administrator can change the chat title, photo, and other settings
	CanPostMessages     bool   `json:"can_post_messages"`      // True, if the administrator can create channel posts; applicable to channels only
	CanEditMessages     bool   `json:"can_edit_messages"`      // True, if the administrator can edit messages of other users and pin messages; applicable to channels only
	CanDeleteMessages   bool   `json:"can_delete_messages"`    // True, if the administrator can delete messages of other users
	CanInviteUsers      bool   `json:"can_invite_users"`       // True, if the administrator can invite new users to the chat
	CanRestrictMembers  bool   `json:"can_restrict_members"`   // True, if the administrator can restrict, ban, or unban chat members; always true for channels
	CanPinMessages      bool   `json:"can_pin_messages"`       // True, if the administrator can pin messages; applicable to basic groups and supergroups only
	CanPromoteMembers   bool   `json:"can_promote_members"`    // True, if the administrator can add new administrators with a subset of their own privileges or demote administrators that were directly or indirectly promoted by them
	CanManageVideoChats bool   `json:"can_manage_video_chats"` // True, if the administrator can manage video chats
	IsAnonymous         bool   `json:"is_anonymous"`           // True, if the administrator isn't shown in the chat member list and sends messages anonymously; applicable to supergroups only
}

// MessageType return the string telegram-type of ChatMemberStatusAdministrator
func (chatMemberStatusAdministrator *ChatMemberStatusAdministrator) MessageType() string {
	return "chatMemberStatusAdministrator"
}

// NewChatMemberStatusAdministrator creates a new ChatMemberStatusAdministrator
//
// @param customTitle A custom title of the administrator; 0-16 characters without emojis; applicable to supergroups only
// @param canBeEdited True, if the current user can edit the administrator privileges for the called user
// @param canManageChat True, if the administrator can get chat event log, get chat statistics, get message statistics in channels, get channel members, see anonymous administrators in supergroups and ignore slow mode. Implied by any other privilege; applicable to supergroups and channels only
// @param canChangeInfo True, if the administrator can change the chat title, photo, and other settings
// @param canPostMessages True, if the administrator can create channel posts; applicable to channels only
// @param canEditMessages True, if the administrator can edit messages of other users and pin messages; applicable to channels only
// @param canDeleteMessages True, if the administrator can delete messages of other users
// @param canInviteUsers True, if the administrator can invite new users to the chat
// @param canRestrictMembers True, if the administrator can restrict, ban, or unban chat members; always true for channels
// @param canPinMessages True, if the administrator can pin messages; applicable to basic groups and supergroups only
// @param canPromoteMembers True, if the administrator can add new administrators with a subset of their own privileges or demote administrators that were directly or indirectly promoted by them
// @param canManageVideoChats True, if the administrator can manage video chats
// @param isAnonymous True, if the administrator isn't shown in the chat member list and sends messages anonymously; applicable to supergroups only
func NewChatMemberStatusAdministrator(customTitle string, canBeEdited bool, canManageChat bool, canChangeInfo bool, canPostMessages bool, canEditMessages bool, canDeleteMessages bool, canInviteUsers bool, canRestrictMembers bool, canPinMessages bool, canPromoteMembers bool, canManageVideoChats bool, isAnonymous bool) *ChatMemberStatusAdministrator {
	chatMemberStatusAdministratorTemp := ChatMemberStatusAdministrator{
		tdCommon:            tdCommon{Type: "chatMemberStatusAdministrator"},
		CustomTitle:         customTitle,
		CanBeEdited:         canBeEdited,
		CanManageChat:       canManageChat,
		CanChangeInfo:       canChangeInfo,
		CanPostMessages:     canPostMessages,
		CanEditMessages:     canEditMessages,
		CanDeleteMessages:   canDeleteMessages,
		CanInviteUsers:      canInviteUsers,
		CanRestrictMembers:  canRestrictMembers,
		CanPinMessages:      canPinMessages,
		CanPromoteMembers:   canPromoteMembers,
		CanManageVideoChats: canManageVideoChats,
		IsAnonymous:         isAnonymous,
	}

	return &chatMemberStatusAdministratorTemp
}

// GetChatMemberStatusEnum return the enum type of this object
func (chatMemberStatusAdministrator *ChatMemberStatusAdministrator) GetChatMemberStatusEnum() ChatMemberStatusEnum {
	return ChatMemberStatusAdministratorType
}

// ChatMemberStatusMember The user is a member of the chat, without any additional privileges or restrictions
type ChatMemberStatusMember struct {
	tdCommon
}

// MessageType return the string telegram-type of ChatMemberStatusMember
func (chatMemberStatusMember *ChatMemberStatusMember) MessageType() string {
	return "chatMemberStatusMember"
}

// NewChatMemberStatusMember creates a new ChatMemberStatusMember
//
func NewChatMemberStatusMember() *ChatMemberStatusMember {
	chatMemberStatusMemberTemp := ChatMemberStatusMember{
		tdCommon: tdCommon{Type: "chatMemberStatusMember"},
	}

	return &chatMemberStatusMemberTemp
}

// GetChatMemberStatusEnum return the enum type of this object
func (chatMemberStatusMember *ChatMemberStatusMember) GetChatMemberStatusEnum() ChatMemberStatusEnum {
	return ChatMemberStatusMemberType
}

// ChatMemberStatusRestricted The user is under certain restrictions in the chat. Not supported in basic groups and channels
type ChatMemberStatusRestricted struct {
	tdCommon
	IsMember            bool             `json:"is_member"`             // True, if the user is a member of the chat
	RestrictedUntilDate int32            `json:"restricted_until_date"` // Point in time (Unix timestamp) when restrictions will be lifted from the user; 0 if never. If the user is restricted for more than 366 days or for less than 30 seconds from the current time, the user is considered to be restricted forever
	Permissions         *ChatPermissions `json:"permissions"`           // User permissions in the chat
}

// MessageType return the string telegram-type of ChatMemberStatusRestricted
func (chatMemberStatusRestricted *ChatMemberStatusRestricted) MessageType() string {
	return "chatMemberStatusRestricted"
}

// NewChatMemberStatusRestricted creates a new ChatMemberStatusRestricted
//
// @param isMember True, if the user is a member of the chat
// @param restrictedUntilDate Point in time (Unix timestamp) when restrictions will be lifted from the user; 0 if never. If the user is restricted for more than 366 days or for less than 30 seconds from the current time, the user is considered to be restricted forever
// @param permissions User permissions in the chat
func NewChatMemberStatusRestricted(isMember bool, restrictedUntilDate int32, permissions *ChatPermissions) *ChatMemberStatusRestricted {
	chatMemberStatusRestrictedTemp := ChatMemberStatusRestricted{
		tdCommon:            tdCommon{Type: "chatMemberStatusRestricted"},
		IsMember:            isMember,
		RestrictedUntilDate: restrictedUntilDate,
		Permissions:         permissions,
	}

	return &chatMemberStatusRestrictedTemp
}

// GetChatMemberStatusEnum return the enum type of this object
func (chatMemberStatusRestricted *ChatMemberStatusRestricted) GetChatMemberStatusEnum() ChatMemberStatusEnum {
	return ChatMemberStatusRestrictedType
}

// ChatMemberStatusLeft The user or the chat is not a chat member
type ChatMemberStatusLeft struct {
	tdCommon
}

// MessageType return the string telegram-type of ChatMemberStatusLeft
func (chatMemberStatusLeft *ChatMemberStatusLeft) MessageType() string {
	return "chatMemberStatusLeft"
}

// NewChatMemberStatusLeft creates a new ChatMemberStatusLeft
//
func NewChatMemberStatusLeft() *ChatMemberStatusLeft {
	chatMemberStatusLeftTemp := ChatMemberStatusLeft{
		tdCommon: tdCommon{Type: "chatMemberStatusLeft"},
	}

	return &chatMemberStatusLeftTemp
}

// GetChatMemberStatusEnum return the enum type of this object
func (chatMemberStatusLeft *ChatMemberStatusLeft) GetChatMemberStatusEnum() ChatMemberStatusEnum {
	return ChatMemberStatusLeftType
}

// ChatMemberStatusBanned The user or the chat was banned (and hence is not a member of the chat). Implies the user can't return to the chat, view messages, or be used as a participant identifier to join a video chat of the chat
type ChatMemberStatusBanned struct {
	tdCommon
	BannedUntilDate int32 `json:"banned_until_date"` // Point in time (Unix timestamp) when the user will be unbanned; 0 if never. If the user is banned for more than 366 days or for less than 30 seconds from the current time, the user is considered to be banned forever. Always 0 in basic groups
}

// MessageType return the string telegram-type of ChatMemberStatusBanned
func (chatMemberStatusBanned *ChatMemberStatusBanned) MessageType() string {
	return "chatMemberStatusBanned"
}

// NewChatMemberStatusBanned creates a new ChatMemberStatusBanned
//
// @param bannedUntilDate Point in time (Unix timestamp) when the user will be unbanned; 0 if never. If the user is banned for more than 366 days or for less than 30 seconds from the current time, the user is considered to be banned forever. Always 0 in basic groups
func NewChatMemberStatusBanned(bannedUntilDate int32) *ChatMemberStatusBanned {
	chatMemberStatusBannedTemp := ChatMemberStatusBanned{
		tdCommon:        tdCommon{Type: "chatMemberStatusBanned"},
		BannedUntilDate: bannedUntilDate,
	}

	return &chatMemberStatusBannedTemp
}

// GetChatMemberStatusEnum return the enum type of this object
func (chatMemberStatusBanned *ChatMemberStatusBanned) GetChatMemberStatusEnum() ChatMemberStatusEnum {
	return ChatMemberStatusBannedType
}
