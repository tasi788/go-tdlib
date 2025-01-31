package tdlib

import (
	"encoding/json"
	"fmt"
)

// ChatActionBar Describes actions which must be possible to do through a chat action bar
type ChatActionBar interface {
	GetChatActionBarEnum() ChatActionBarEnum
}

// ChatActionBarEnum Alias for abstract ChatActionBar 'Sub-Classes', used as constant-enum here
type ChatActionBarEnum string

// ChatActionBar enums
const (
	ChatActionBarReportSpamType              ChatActionBarEnum = "chatActionBarReportSpam"
	ChatActionBarReportUnrelatedLocationType ChatActionBarEnum = "chatActionBarReportUnrelatedLocation"
	ChatActionBarInviteMembersType           ChatActionBarEnum = "chatActionBarInviteMembers"
	ChatActionBarReportAddBlockType          ChatActionBarEnum = "chatActionBarReportAddBlock"
	ChatActionBarAddContactType              ChatActionBarEnum = "chatActionBarAddContact"
	ChatActionBarSharePhoneNumberType        ChatActionBarEnum = "chatActionBarSharePhoneNumber"
	ChatActionBarJoinRequestType             ChatActionBarEnum = "chatActionBarJoinRequest"
)

func unmarshalChatActionBar(rawMsg *json.RawMessage) (ChatActionBar, error) {

	if rawMsg == nil {
		return nil, nil
	}
	var objMap map[string]interface{}
	err := json.Unmarshal(*rawMsg, &objMap)
	if err != nil {
		return nil, err
	}

	switch ChatActionBarEnum(objMap["@type"].(string)) {
	case ChatActionBarReportSpamType:
		var chatActionBarReportSpam ChatActionBarReportSpam
		err := json.Unmarshal(*rawMsg, &chatActionBarReportSpam)
		return &chatActionBarReportSpam, err

	case ChatActionBarReportUnrelatedLocationType:
		var chatActionBarReportUnrelatedLocation ChatActionBarReportUnrelatedLocation
		err := json.Unmarshal(*rawMsg, &chatActionBarReportUnrelatedLocation)
		return &chatActionBarReportUnrelatedLocation, err

	case ChatActionBarInviteMembersType:
		var chatActionBarInviteMembers ChatActionBarInviteMembers
		err := json.Unmarshal(*rawMsg, &chatActionBarInviteMembers)
		return &chatActionBarInviteMembers, err

	case ChatActionBarReportAddBlockType:
		var chatActionBarReportAddBlock ChatActionBarReportAddBlock
		err := json.Unmarshal(*rawMsg, &chatActionBarReportAddBlock)
		return &chatActionBarReportAddBlock, err

	case ChatActionBarAddContactType:
		var chatActionBarAddContact ChatActionBarAddContact
		err := json.Unmarshal(*rawMsg, &chatActionBarAddContact)
		return &chatActionBarAddContact, err

	case ChatActionBarSharePhoneNumberType:
		var chatActionBarSharePhoneNumber ChatActionBarSharePhoneNumber
		err := json.Unmarshal(*rawMsg, &chatActionBarSharePhoneNumber)
		return &chatActionBarSharePhoneNumber, err

	case ChatActionBarJoinRequestType:
		var chatActionBarJoinRequest ChatActionBarJoinRequest
		err := json.Unmarshal(*rawMsg, &chatActionBarJoinRequest)
		return &chatActionBarJoinRequest, err

	default:
		return nil, fmt.Errorf("Error UnMarshaling, unknown type:" + objMap["@type"].(string))
	}
}

// ChatActionBarReportSpam The chat can be reported as spam using the method reportChat with the reason chatReportReasonSpam
type ChatActionBarReportSpam struct {
	tdCommon
	CanUnarchive bool `json:"can_unarchive"` // If true, the chat was automatically archived and can be moved back to the main chat list using addChatToList simultaneously with setting chat notification settings to default using setChatNotificationSettings
}

// MessageType return the string telegram-type of ChatActionBarReportSpam
func (chatActionBarReportSpam *ChatActionBarReportSpam) MessageType() string {
	return "chatActionBarReportSpam"
}

// NewChatActionBarReportSpam creates a new ChatActionBarReportSpam
//
// @param canUnarchive If true, the chat was automatically archived and can be moved back to the main chat list using addChatToList simultaneously with setting chat notification settings to default using setChatNotificationSettings
func NewChatActionBarReportSpam(canUnarchive bool) *ChatActionBarReportSpam {
	chatActionBarReportSpamTemp := ChatActionBarReportSpam{
		tdCommon:     tdCommon{Type: "chatActionBarReportSpam"},
		CanUnarchive: canUnarchive,
	}

	return &chatActionBarReportSpamTemp
}

// GetChatActionBarEnum return the enum type of this object
func (chatActionBarReportSpam *ChatActionBarReportSpam) GetChatActionBarEnum() ChatActionBarEnum {
	return ChatActionBarReportSpamType
}

// ChatActionBarReportUnrelatedLocation The chat is a location-based supergroup, which can be reported as having unrelated location using the method reportChat with the reason chatReportReasonUnrelatedLocation
type ChatActionBarReportUnrelatedLocation struct {
	tdCommon
}

// MessageType return the string telegram-type of ChatActionBarReportUnrelatedLocation
func (chatActionBarReportUnrelatedLocation *ChatActionBarReportUnrelatedLocation) MessageType() string {
	return "chatActionBarReportUnrelatedLocation"
}

// NewChatActionBarReportUnrelatedLocation creates a new ChatActionBarReportUnrelatedLocation
//
func NewChatActionBarReportUnrelatedLocation() *ChatActionBarReportUnrelatedLocation {
	chatActionBarReportUnrelatedLocationTemp := ChatActionBarReportUnrelatedLocation{
		tdCommon: tdCommon{Type: "chatActionBarReportUnrelatedLocation"},
	}

	return &chatActionBarReportUnrelatedLocationTemp
}

// GetChatActionBarEnum return the enum type of this object
func (chatActionBarReportUnrelatedLocation *ChatActionBarReportUnrelatedLocation) GetChatActionBarEnum() ChatActionBarEnum {
	return ChatActionBarReportUnrelatedLocationType
}

// ChatActionBarInviteMembers The chat is a recently created group chat to which new members can be invited
type ChatActionBarInviteMembers struct {
	tdCommon
}

// MessageType return the string telegram-type of ChatActionBarInviteMembers
func (chatActionBarInviteMembers *ChatActionBarInviteMembers) MessageType() string {
	return "chatActionBarInviteMembers"
}

// NewChatActionBarInviteMembers creates a new ChatActionBarInviteMembers
//
func NewChatActionBarInviteMembers() *ChatActionBarInviteMembers {
	chatActionBarInviteMembersTemp := ChatActionBarInviteMembers{
		tdCommon: tdCommon{Type: "chatActionBarInviteMembers"},
	}

	return &chatActionBarInviteMembersTemp
}

// GetChatActionBarEnum return the enum type of this object
func (chatActionBarInviteMembers *ChatActionBarInviteMembers) GetChatActionBarEnum() ChatActionBarEnum {
	return ChatActionBarInviteMembersType
}

// ChatActionBarReportAddBlock The chat is a private or secret chat, which can be reported using the method reportChat, or the other user can be blocked using the method toggleMessageSenderIsBlocked, or the other user can be added to the contact list using the method addContact
type ChatActionBarReportAddBlock struct {
	tdCommon
	CanUnarchive bool  `json:"can_unarchive"` // If true, the chat was automatically archived and can be moved back to the main chat list using addChatToList simultaneously with setting chat notification settings to default using setChatNotificationSettings
	Distance     int32 `json:"distance"`      // If non-negative, the current user was found by the peer through searchChatsNearby and this is the distance between the users
}

// MessageType return the string telegram-type of ChatActionBarReportAddBlock
func (chatActionBarReportAddBlock *ChatActionBarReportAddBlock) MessageType() string {
	return "chatActionBarReportAddBlock"
}

// NewChatActionBarReportAddBlock creates a new ChatActionBarReportAddBlock
//
// @param canUnarchive If true, the chat was automatically archived and can be moved back to the main chat list using addChatToList simultaneously with setting chat notification settings to default using setChatNotificationSettings
// @param distance If non-negative, the current user was found by the peer through searchChatsNearby and this is the distance between the users
func NewChatActionBarReportAddBlock(canUnarchive bool, distance int32) *ChatActionBarReportAddBlock {
	chatActionBarReportAddBlockTemp := ChatActionBarReportAddBlock{
		tdCommon:     tdCommon{Type: "chatActionBarReportAddBlock"},
		CanUnarchive: canUnarchive,
		Distance:     distance,
	}

	return &chatActionBarReportAddBlockTemp
}

// GetChatActionBarEnum return the enum type of this object
func (chatActionBarReportAddBlock *ChatActionBarReportAddBlock) GetChatActionBarEnum() ChatActionBarEnum {
	return ChatActionBarReportAddBlockType
}

// ChatActionBarAddContact The chat is a private or secret chat and the other user can be added to the contact list using the method addContact
type ChatActionBarAddContact struct {
	tdCommon
}

// MessageType return the string telegram-type of ChatActionBarAddContact
func (chatActionBarAddContact *ChatActionBarAddContact) MessageType() string {
	return "chatActionBarAddContact"
}

// NewChatActionBarAddContact creates a new ChatActionBarAddContact
//
func NewChatActionBarAddContact() *ChatActionBarAddContact {
	chatActionBarAddContactTemp := ChatActionBarAddContact{
		tdCommon: tdCommon{Type: "chatActionBarAddContact"},
	}

	return &chatActionBarAddContactTemp
}

// GetChatActionBarEnum return the enum type of this object
func (chatActionBarAddContact *ChatActionBarAddContact) GetChatActionBarEnum() ChatActionBarEnum {
	return ChatActionBarAddContactType
}

// ChatActionBarSharePhoneNumber The chat is a private or secret chat with a mutual contact and the user's phone number can be shared with the other user using the method sharePhoneNumber
type ChatActionBarSharePhoneNumber struct {
	tdCommon
}

// MessageType return the string telegram-type of ChatActionBarSharePhoneNumber
func (chatActionBarSharePhoneNumber *ChatActionBarSharePhoneNumber) MessageType() string {
	return "chatActionBarSharePhoneNumber"
}

// NewChatActionBarSharePhoneNumber creates a new ChatActionBarSharePhoneNumber
//
func NewChatActionBarSharePhoneNumber() *ChatActionBarSharePhoneNumber {
	chatActionBarSharePhoneNumberTemp := ChatActionBarSharePhoneNumber{
		tdCommon: tdCommon{Type: "chatActionBarSharePhoneNumber"},
	}

	return &chatActionBarSharePhoneNumberTemp
}

// GetChatActionBarEnum return the enum type of this object
func (chatActionBarSharePhoneNumber *ChatActionBarSharePhoneNumber) GetChatActionBarEnum() ChatActionBarEnum {
	return ChatActionBarSharePhoneNumberType
}

// ChatActionBarJoinRequest The chat is a private chat with an administrator of a chat to which the user sent join request
type ChatActionBarJoinRequest struct {
	tdCommon
	Title       string `json:"title"`        // Title of the chat to which the join request was sent
	IsChannel   bool   `json:"is_channel"`   // True, if the join request was sent to a channel chat
	RequestDate int32  `json:"request_date"` // Point in time (Unix timestamp) when the join request was sent
}

// MessageType return the string telegram-type of ChatActionBarJoinRequest
func (chatActionBarJoinRequest *ChatActionBarJoinRequest) MessageType() string {
	return "chatActionBarJoinRequest"
}

// NewChatActionBarJoinRequest creates a new ChatActionBarJoinRequest
//
// @param title Title of the chat to which the join request was sent
// @param isChannel True, if the join request was sent to a channel chat
// @param requestDate Point in time (Unix timestamp) when the join request was sent
func NewChatActionBarJoinRequest(title string, isChannel bool, requestDate int32) *ChatActionBarJoinRequest {
	chatActionBarJoinRequestTemp := ChatActionBarJoinRequest{
		tdCommon:    tdCommon{Type: "chatActionBarJoinRequest"},
		Title:       title,
		IsChannel:   isChannel,
		RequestDate: requestDate,
	}

	return &chatActionBarJoinRequestTemp
}

// GetChatActionBarEnum return the enum type of this object
func (chatActionBarJoinRequest *ChatActionBarJoinRequest) GetChatActionBarEnum() ChatActionBarEnum {
	return ChatActionBarJoinRequestType
}
