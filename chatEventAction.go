package tdlib

import (
	"encoding/json"
	"fmt"
)

// ChatEventAction Represents a chat event
type ChatEventAction interface {
	GetChatEventActionEnum() ChatEventActionEnum
}

// ChatEventActionEnum Alias for abstract ChatEventAction 'Sub-Classes', used as constant-enum here
type ChatEventActionEnum string

// ChatEventAction enums
const (
	ChatEventMessageEditedType                          ChatEventActionEnum = "chatEventMessageEdited"
	ChatEventMessageDeletedType                         ChatEventActionEnum = "chatEventMessageDeleted"
	ChatEventPollStoppedType                            ChatEventActionEnum = "chatEventPollStopped"
	ChatEventMessagePinnedType                          ChatEventActionEnum = "chatEventMessagePinned"
	ChatEventMessageUnpinnedType                        ChatEventActionEnum = "chatEventMessageUnpinned"
	ChatEventMemberJoinedType                           ChatEventActionEnum = "chatEventMemberJoined"
	ChatEventMemberJoinedByInviteLinkType               ChatEventActionEnum = "chatEventMemberJoinedByInviteLink"
	ChatEventMemberJoinedByRequestType                  ChatEventActionEnum = "chatEventMemberJoinedByRequest"
	ChatEventMemberLeftType                             ChatEventActionEnum = "chatEventMemberLeft"
	ChatEventMemberInvitedType                          ChatEventActionEnum = "chatEventMemberInvited"
	ChatEventMemberPromotedType                         ChatEventActionEnum = "chatEventMemberPromoted"
	ChatEventMemberRestrictedType                       ChatEventActionEnum = "chatEventMemberRestricted"
	ChatEventTitleChangedType                           ChatEventActionEnum = "chatEventTitleChanged"
	ChatEventPermissionsChangedType                     ChatEventActionEnum = "chatEventPermissionsChanged"
	ChatEventDescriptionChangedType                     ChatEventActionEnum = "chatEventDescriptionChanged"
	ChatEventUsernameChangedType                        ChatEventActionEnum = "chatEventUsernameChanged"
	ChatEventPhotoChangedType                           ChatEventActionEnum = "chatEventPhotoChanged"
	ChatEventInvitesToggledType                         ChatEventActionEnum = "chatEventInvitesToggled"
	ChatEventLinkedChatChangedType                      ChatEventActionEnum = "chatEventLinkedChatChanged"
	ChatEventSlowModeDelayChangedType                   ChatEventActionEnum = "chatEventSlowModeDelayChanged"
	ChatEventMessageTtlChangedType                      ChatEventActionEnum = "chatEventMessageTtlChanged"
	ChatEventSignMessagesToggledType                    ChatEventActionEnum = "chatEventSignMessagesToggled"
	ChatEventHasProtectedContentToggledType             ChatEventActionEnum = "chatEventHasProtectedContentToggled"
	ChatEventStickerSetChangedType                      ChatEventActionEnum = "chatEventStickerSetChanged"
	ChatEventLocationChangedType                        ChatEventActionEnum = "chatEventLocationChanged"
	ChatEventIsAllHistoryAvailableToggledType           ChatEventActionEnum = "chatEventIsAllHistoryAvailableToggled"
	ChatEventInviteLinkEditedType                       ChatEventActionEnum = "chatEventInviteLinkEdited"
	ChatEventInviteLinkRevokedType                      ChatEventActionEnum = "chatEventInviteLinkRevoked"
	ChatEventInviteLinkDeletedType                      ChatEventActionEnum = "chatEventInviteLinkDeleted"
	ChatEventVideoChatCreatedType                       ChatEventActionEnum = "chatEventVideoChatCreated"
	ChatEventVideoChatEndedType                         ChatEventActionEnum = "chatEventVideoChatEnded"
	ChatEventVideoChatParticipantIsMutedToggledType     ChatEventActionEnum = "chatEventVideoChatParticipantIsMutedToggled"
	ChatEventVideoChatParticipantVolumeLevelChangedType ChatEventActionEnum = "chatEventVideoChatParticipantVolumeLevelChanged"
	ChatEventVideoChatMuteNewParticipantsToggledType    ChatEventActionEnum = "chatEventVideoChatMuteNewParticipantsToggled"
)

func unmarshalChatEventAction(rawMsg *json.RawMessage) (ChatEventAction, error) {

	if rawMsg == nil {
		return nil, nil
	}
	var objMap map[string]interface{}
	err := json.Unmarshal(*rawMsg, &objMap)
	if err != nil {
		return nil, err
	}

	switch ChatEventActionEnum(objMap["@type"].(string)) {
	case ChatEventMessageEditedType:
		var chatEventMessageEdited ChatEventMessageEdited
		err := json.Unmarshal(*rawMsg, &chatEventMessageEdited)
		return &chatEventMessageEdited, err

	case ChatEventMessageDeletedType:
		var chatEventMessageDeleted ChatEventMessageDeleted
		err := json.Unmarshal(*rawMsg, &chatEventMessageDeleted)
		return &chatEventMessageDeleted, err

	case ChatEventPollStoppedType:
		var chatEventPollStopped ChatEventPollStopped
		err := json.Unmarshal(*rawMsg, &chatEventPollStopped)
		return &chatEventPollStopped, err

	case ChatEventMessagePinnedType:
		var chatEventMessagePinned ChatEventMessagePinned
		err := json.Unmarshal(*rawMsg, &chatEventMessagePinned)
		return &chatEventMessagePinned, err

	case ChatEventMessageUnpinnedType:
		var chatEventMessageUnpinned ChatEventMessageUnpinned
		err := json.Unmarshal(*rawMsg, &chatEventMessageUnpinned)
		return &chatEventMessageUnpinned, err

	case ChatEventMemberJoinedType:
		var chatEventMemberJoined ChatEventMemberJoined
		err := json.Unmarshal(*rawMsg, &chatEventMemberJoined)
		return &chatEventMemberJoined, err

	case ChatEventMemberJoinedByInviteLinkType:
		var chatEventMemberJoinedByInviteLink ChatEventMemberJoinedByInviteLink
		err := json.Unmarshal(*rawMsg, &chatEventMemberJoinedByInviteLink)
		return &chatEventMemberJoinedByInviteLink, err

	case ChatEventMemberJoinedByRequestType:
		var chatEventMemberJoinedByRequest ChatEventMemberJoinedByRequest
		err := json.Unmarshal(*rawMsg, &chatEventMemberJoinedByRequest)
		return &chatEventMemberJoinedByRequest, err

	case ChatEventMemberLeftType:
		var chatEventMemberLeft ChatEventMemberLeft
		err := json.Unmarshal(*rawMsg, &chatEventMemberLeft)
		return &chatEventMemberLeft, err

	case ChatEventMemberInvitedType:
		var chatEventMemberInvited ChatEventMemberInvited
		err := json.Unmarshal(*rawMsg, &chatEventMemberInvited)
		return &chatEventMemberInvited, err

	case ChatEventMemberPromotedType:
		var chatEventMemberPromoted ChatEventMemberPromoted
		err := json.Unmarshal(*rawMsg, &chatEventMemberPromoted)
		return &chatEventMemberPromoted, err

	case ChatEventMemberRestrictedType:
		var chatEventMemberRestricted ChatEventMemberRestricted
		err := json.Unmarshal(*rawMsg, &chatEventMemberRestricted)
		return &chatEventMemberRestricted, err

	case ChatEventTitleChangedType:
		var chatEventTitleChanged ChatEventTitleChanged
		err := json.Unmarshal(*rawMsg, &chatEventTitleChanged)
		return &chatEventTitleChanged, err

	case ChatEventPermissionsChangedType:
		var chatEventPermissionsChanged ChatEventPermissionsChanged
		err := json.Unmarshal(*rawMsg, &chatEventPermissionsChanged)
		return &chatEventPermissionsChanged, err

	case ChatEventDescriptionChangedType:
		var chatEventDescriptionChanged ChatEventDescriptionChanged
		err := json.Unmarshal(*rawMsg, &chatEventDescriptionChanged)
		return &chatEventDescriptionChanged, err

	case ChatEventUsernameChangedType:
		var chatEventUsernameChanged ChatEventUsernameChanged
		err := json.Unmarshal(*rawMsg, &chatEventUsernameChanged)
		return &chatEventUsernameChanged, err

	case ChatEventPhotoChangedType:
		var chatEventPhotoChanged ChatEventPhotoChanged
		err := json.Unmarshal(*rawMsg, &chatEventPhotoChanged)
		return &chatEventPhotoChanged, err

	case ChatEventInvitesToggledType:
		var chatEventInvitesToggled ChatEventInvitesToggled
		err := json.Unmarshal(*rawMsg, &chatEventInvitesToggled)
		return &chatEventInvitesToggled, err

	case ChatEventLinkedChatChangedType:
		var chatEventLinkedChatChanged ChatEventLinkedChatChanged
		err := json.Unmarshal(*rawMsg, &chatEventLinkedChatChanged)
		return &chatEventLinkedChatChanged, err

	case ChatEventSlowModeDelayChangedType:
		var chatEventSlowModeDelayChanged ChatEventSlowModeDelayChanged
		err := json.Unmarshal(*rawMsg, &chatEventSlowModeDelayChanged)
		return &chatEventSlowModeDelayChanged, err

	case ChatEventMessageTtlChangedType:
		var chatEventMessageTtlChanged ChatEventMessageTtlChanged
		err := json.Unmarshal(*rawMsg, &chatEventMessageTtlChanged)
		return &chatEventMessageTtlChanged, err

	case ChatEventSignMessagesToggledType:
		var chatEventSignMessagesToggled ChatEventSignMessagesToggled
		err := json.Unmarshal(*rawMsg, &chatEventSignMessagesToggled)
		return &chatEventSignMessagesToggled, err

	case ChatEventHasProtectedContentToggledType:
		var chatEventHasProtectedContentToggled ChatEventHasProtectedContentToggled
		err := json.Unmarshal(*rawMsg, &chatEventHasProtectedContentToggled)
		return &chatEventHasProtectedContentToggled, err

	case ChatEventStickerSetChangedType:
		var chatEventStickerSetChanged ChatEventStickerSetChanged
		err := json.Unmarshal(*rawMsg, &chatEventStickerSetChanged)
		return &chatEventStickerSetChanged, err

	case ChatEventLocationChangedType:
		var chatEventLocationChanged ChatEventLocationChanged
		err := json.Unmarshal(*rawMsg, &chatEventLocationChanged)
		return &chatEventLocationChanged, err

	case ChatEventIsAllHistoryAvailableToggledType:
		var chatEventIsAllHistoryAvailableToggled ChatEventIsAllHistoryAvailableToggled
		err := json.Unmarshal(*rawMsg, &chatEventIsAllHistoryAvailableToggled)
		return &chatEventIsAllHistoryAvailableToggled, err

	case ChatEventInviteLinkEditedType:
		var chatEventInviteLinkEdited ChatEventInviteLinkEdited
		err := json.Unmarshal(*rawMsg, &chatEventInviteLinkEdited)
		return &chatEventInviteLinkEdited, err

	case ChatEventInviteLinkRevokedType:
		var chatEventInviteLinkRevoked ChatEventInviteLinkRevoked
		err := json.Unmarshal(*rawMsg, &chatEventInviteLinkRevoked)
		return &chatEventInviteLinkRevoked, err

	case ChatEventInviteLinkDeletedType:
		var chatEventInviteLinkDeleted ChatEventInviteLinkDeleted
		err := json.Unmarshal(*rawMsg, &chatEventInviteLinkDeleted)
		return &chatEventInviteLinkDeleted, err

	case ChatEventVideoChatCreatedType:
		var chatEventVideoChatCreated ChatEventVideoChatCreated
		err := json.Unmarshal(*rawMsg, &chatEventVideoChatCreated)
		return &chatEventVideoChatCreated, err

	case ChatEventVideoChatEndedType:
		var chatEventVideoChatEnded ChatEventVideoChatEnded
		err := json.Unmarshal(*rawMsg, &chatEventVideoChatEnded)
		return &chatEventVideoChatEnded, err

	case ChatEventVideoChatParticipantIsMutedToggledType:
		var chatEventVideoChatParticipantIsMutedToggled ChatEventVideoChatParticipantIsMutedToggled
		err := json.Unmarshal(*rawMsg, &chatEventVideoChatParticipantIsMutedToggled)
		return &chatEventVideoChatParticipantIsMutedToggled, err

	case ChatEventVideoChatParticipantVolumeLevelChangedType:
		var chatEventVideoChatParticipantVolumeLevelChanged ChatEventVideoChatParticipantVolumeLevelChanged
		err := json.Unmarshal(*rawMsg, &chatEventVideoChatParticipantVolumeLevelChanged)
		return &chatEventVideoChatParticipantVolumeLevelChanged, err

	case ChatEventVideoChatMuteNewParticipantsToggledType:
		var chatEventVideoChatMuteNewParticipantsToggled ChatEventVideoChatMuteNewParticipantsToggled
		err := json.Unmarshal(*rawMsg, &chatEventVideoChatMuteNewParticipantsToggled)
		return &chatEventVideoChatMuteNewParticipantsToggled, err

	default:
		return nil, fmt.Errorf("Error UnMarshaling, unknown type:" + objMap["@type"].(string))
	}
}

// ChatEventMessageEdited A message was edited
type ChatEventMessageEdited struct {
	tdCommon
	OldMessage *Message `json:"old_message"` // The original message before the edit
	NewMessage *Message `json:"new_message"` // The message after it was edited
}

// MessageType return the string telegram-type of ChatEventMessageEdited
func (chatEventMessageEdited *ChatEventMessageEdited) MessageType() string {
	return "chatEventMessageEdited"
}

// NewChatEventMessageEdited creates a new ChatEventMessageEdited
//
// @param oldMessage The original message before the edit
// @param newMessage The message after it was edited
func NewChatEventMessageEdited(oldMessage *Message, newMessage *Message) *ChatEventMessageEdited {
	chatEventMessageEditedTemp := ChatEventMessageEdited{
		tdCommon:   tdCommon{Type: "chatEventMessageEdited"},
		OldMessage: oldMessage,
		NewMessage: newMessage,
	}

	return &chatEventMessageEditedTemp
}

// GetChatEventActionEnum return the enum type of this object
func (chatEventMessageEdited *ChatEventMessageEdited) GetChatEventActionEnum() ChatEventActionEnum {
	return ChatEventMessageEditedType
}

// ChatEventMessageDeleted A message was deleted
type ChatEventMessageDeleted struct {
	tdCommon
	Message *Message `json:"message"` // Deleted message
}

// MessageType return the string telegram-type of ChatEventMessageDeleted
func (chatEventMessageDeleted *ChatEventMessageDeleted) MessageType() string {
	return "chatEventMessageDeleted"
}

// NewChatEventMessageDeleted creates a new ChatEventMessageDeleted
//
// @param message Deleted message
func NewChatEventMessageDeleted(message *Message) *ChatEventMessageDeleted {
	chatEventMessageDeletedTemp := ChatEventMessageDeleted{
		tdCommon: tdCommon{Type: "chatEventMessageDeleted"},
		Message:  message,
	}

	return &chatEventMessageDeletedTemp
}

// GetChatEventActionEnum return the enum type of this object
func (chatEventMessageDeleted *ChatEventMessageDeleted) GetChatEventActionEnum() ChatEventActionEnum {
	return ChatEventMessageDeletedType
}

// ChatEventPollStopped A poll in a message was stopped
type ChatEventPollStopped struct {
	tdCommon
	Message *Message `json:"message"` // The message with the poll
}

// MessageType return the string telegram-type of ChatEventPollStopped
func (chatEventPollStopped *ChatEventPollStopped) MessageType() string {
	return "chatEventPollStopped"
}

// NewChatEventPollStopped creates a new ChatEventPollStopped
//
// @param message The message with the poll
func NewChatEventPollStopped(message *Message) *ChatEventPollStopped {
	chatEventPollStoppedTemp := ChatEventPollStopped{
		tdCommon: tdCommon{Type: "chatEventPollStopped"},
		Message:  message,
	}

	return &chatEventPollStoppedTemp
}

// GetChatEventActionEnum return the enum type of this object
func (chatEventPollStopped *ChatEventPollStopped) GetChatEventActionEnum() ChatEventActionEnum {
	return ChatEventPollStoppedType
}

// ChatEventMessagePinned A message was pinned
type ChatEventMessagePinned struct {
	tdCommon
	Message *Message `json:"message"` // Pinned message
}

// MessageType return the string telegram-type of ChatEventMessagePinned
func (chatEventMessagePinned *ChatEventMessagePinned) MessageType() string {
	return "chatEventMessagePinned"
}

// NewChatEventMessagePinned creates a new ChatEventMessagePinned
//
// @param message Pinned message
func NewChatEventMessagePinned(message *Message) *ChatEventMessagePinned {
	chatEventMessagePinnedTemp := ChatEventMessagePinned{
		tdCommon: tdCommon{Type: "chatEventMessagePinned"},
		Message:  message,
	}

	return &chatEventMessagePinnedTemp
}

// GetChatEventActionEnum return the enum type of this object
func (chatEventMessagePinned *ChatEventMessagePinned) GetChatEventActionEnum() ChatEventActionEnum {
	return ChatEventMessagePinnedType
}

// ChatEventMessageUnpinned A message was unpinned
type ChatEventMessageUnpinned struct {
	tdCommon
	Message *Message `json:"message"` // Unpinned message
}

// MessageType return the string telegram-type of ChatEventMessageUnpinned
func (chatEventMessageUnpinned *ChatEventMessageUnpinned) MessageType() string {
	return "chatEventMessageUnpinned"
}

// NewChatEventMessageUnpinned creates a new ChatEventMessageUnpinned
//
// @param message Unpinned message
func NewChatEventMessageUnpinned(message *Message) *ChatEventMessageUnpinned {
	chatEventMessageUnpinnedTemp := ChatEventMessageUnpinned{
		tdCommon: tdCommon{Type: "chatEventMessageUnpinned"},
		Message:  message,
	}

	return &chatEventMessageUnpinnedTemp
}

// GetChatEventActionEnum return the enum type of this object
func (chatEventMessageUnpinned *ChatEventMessageUnpinned) GetChatEventActionEnum() ChatEventActionEnum {
	return ChatEventMessageUnpinnedType
}

// ChatEventMemberJoined A new member joined the chat
type ChatEventMemberJoined struct {
	tdCommon
}

// MessageType return the string telegram-type of ChatEventMemberJoined
func (chatEventMemberJoined *ChatEventMemberJoined) MessageType() string {
	return "chatEventMemberJoined"
}

// NewChatEventMemberJoined creates a new ChatEventMemberJoined
//
func NewChatEventMemberJoined() *ChatEventMemberJoined {
	chatEventMemberJoinedTemp := ChatEventMemberJoined{
		tdCommon: tdCommon{Type: "chatEventMemberJoined"},
	}

	return &chatEventMemberJoinedTemp
}

// GetChatEventActionEnum return the enum type of this object
func (chatEventMemberJoined *ChatEventMemberJoined) GetChatEventActionEnum() ChatEventActionEnum {
	return ChatEventMemberJoinedType
}

// ChatEventMemberJoinedByInviteLink A new member joined the chat via an invite link
type ChatEventMemberJoinedByInviteLink struct {
	tdCommon
	InviteLink *ChatInviteLink `json:"invite_link"` // Invite link used to join the chat
}

// MessageType return the string telegram-type of ChatEventMemberJoinedByInviteLink
func (chatEventMemberJoinedByInviteLink *ChatEventMemberJoinedByInviteLink) MessageType() string {
	return "chatEventMemberJoinedByInviteLink"
}

// NewChatEventMemberJoinedByInviteLink creates a new ChatEventMemberJoinedByInviteLink
//
// @param inviteLink Invite link used to join the chat
func NewChatEventMemberJoinedByInviteLink(inviteLink *ChatInviteLink) *ChatEventMemberJoinedByInviteLink {
	chatEventMemberJoinedByInviteLinkTemp := ChatEventMemberJoinedByInviteLink{
		tdCommon:   tdCommon{Type: "chatEventMemberJoinedByInviteLink"},
		InviteLink: inviteLink,
	}

	return &chatEventMemberJoinedByInviteLinkTemp
}

// GetChatEventActionEnum return the enum type of this object
func (chatEventMemberJoinedByInviteLink *ChatEventMemberJoinedByInviteLink) GetChatEventActionEnum() ChatEventActionEnum {
	return ChatEventMemberJoinedByInviteLinkType
}

// ChatEventMemberJoinedByRequest A new member was accepted to the chat by an administrator
type ChatEventMemberJoinedByRequest struct {
	tdCommon
	ApproverUserId int64           `json:"approver_user_id"` // User identifier of the chat administrator, approved user join request
	InviteLink     *ChatInviteLink `json:"invite_link"`      // Invite link used to join the chat; may be null
}

// MessageType return the string telegram-type of ChatEventMemberJoinedByRequest
func (chatEventMemberJoinedByRequest *ChatEventMemberJoinedByRequest) MessageType() string {
	return "chatEventMemberJoinedByRequest"
}

// NewChatEventMemberJoinedByRequest creates a new ChatEventMemberJoinedByRequest
//
// @param approverUserId User identifier of the chat administrator, approved user join request
// @param inviteLink Invite link used to join the chat; may be null
func NewChatEventMemberJoinedByRequest(approverUserId int64, inviteLink *ChatInviteLink) *ChatEventMemberJoinedByRequest {
	chatEventMemberJoinedByRequestTemp := ChatEventMemberJoinedByRequest{
		tdCommon:       tdCommon{Type: "chatEventMemberJoinedByRequest"},
		ApproverUserId: approverUserId,
		InviteLink:     inviteLink,
	}

	return &chatEventMemberJoinedByRequestTemp
}

// GetChatEventActionEnum return the enum type of this object
func (chatEventMemberJoinedByRequest *ChatEventMemberJoinedByRequest) GetChatEventActionEnum() ChatEventActionEnum {
	return ChatEventMemberJoinedByRequestType
}

// ChatEventMemberLeft A member left the chat
type ChatEventMemberLeft struct {
	tdCommon
}

// MessageType return the string telegram-type of ChatEventMemberLeft
func (chatEventMemberLeft *ChatEventMemberLeft) MessageType() string {
	return "chatEventMemberLeft"
}

// NewChatEventMemberLeft creates a new ChatEventMemberLeft
//
func NewChatEventMemberLeft() *ChatEventMemberLeft {
	chatEventMemberLeftTemp := ChatEventMemberLeft{
		tdCommon: tdCommon{Type: "chatEventMemberLeft"},
	}

	return &chatEventMemberLeftTemp
}

// GetChatEventActionEnum return the enum type of this object
func (chatEventMemberLeft *ChatEventMemberLeft) GetChatEventActionEnum() ChatEventActionEnum {
	return ChatEventMemberLeftType
}

// ChatEventMemberInvited A new chat member was invited
type ChatEventMemberInvited struct {
	tdCommon
	UserId int64            `json:"user_id"` // New member user identifier
	Status ChatMemberStatus `json:"status"`  // New member status
}

// MessageType return the string telegram-type of ChatEventMemberInvited
func (chatEventMemberInvited *ChatEventMemberInvited) MessageType() string {
	return "chatEventMemberInvited"
}

// NewChatEventMemberInvited creates a new ChatEventMemberInvited
//
// @param userId New member user identifier
// @param status New member status
func NewChatEventMemberInvited(userId int64, status ChatMemberStatus) *ChatEventMemberInvited {
	chatEventMemberInvitedTemp := ChatEventMemberInvited{
		tdCommon: tdCommon{Type: "chatEventMemberInvited"},
		UserId:   userId,
		Status:   status,
	}

	return &chatEventMemberInvitedTemp
}

// UnmarshalJSON unmarshal to json
func (chatEventMemberInvited *ChatEventMemberInvited) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	tempObj := struct {
		tdCommon
		UserId int64 `json:"user_id"` // New member user identifier

	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	chatEventMemberInvited.tdCommon = tempObj.tdCommon
	chatEventMemberInvited.UserId = tempObj.UserId

	fieldStatus, _ := unmarshalChatMemberStatus(objMap["status"])
	chatEventMemberInvited.Status = fieldStatus

	return nil
}

// GetChatEventActionEnum return the enum type of this object
func (chatEventMemberInvited *ChatEventMemberInvited) GetChatEventActionEnum() ChatEventActionEnum {
	return ChatEventMemberInvitedType
}

// ChatEventMemberPromoted A chat member has gained/lost administrator status, or the list of their administrator privileges has changed
type ChatEventMemberPromoted struct {
	tdCommon
	UserId    int64            `json:"user_id"`    // Affected chat member user identifier
	OldStatus ChatMemberStatus `json:"old_status"` // Previous status of the chat member
	NewStatus ChatMemberStatus `json:"new_status"` // New status of the chat member
}

// MessageType return the string telegram-type of ChatEventMemberPromoted
func (chatEventMemberPromoted *ChatEventMemberPromoted) MessageType() string {
	return "chatEventMemberPromoted"
}

// NewChatEventMemberPromoted creates a new ChatEventMemberPromoted
//
// @param userId Affected chat member user identifier
// @param oldStatus Previous status of the chat member
// @param newStatus New status of the chat member
func NewChatEventMemberPromoted(userId int64, oldStatus ChatMemberStatus, newStatus ChatMemberStatus) *ChatEventMemberPromoted {
	chatEventMemberPromotedTemp := ChatEventMemberPromoted{
		tdCommon:  tdCommon{Type: "chatEventMemberPromoted"},
		UserId:    userId,
		OldStatus: oldStatus,
		NewStatus: newStatus,
	}

	return &chatEventMemberPromotedTemp
}

// UnmarshalJSON unmarshal to json
func (chatEventMemberPromoted *ChatEventMemberPromoted) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	tempObj := struct {
		tdCommon
		UserId int64 `json:"user_id"` // Affected chat member user identifier

	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	chatEventMemberPromoted.tdCommon = tempObj.tdCommon
	chatEventMemberPromoted.UserId = tempObj.UserId

	fieldOldStatus, _ := unmarshalChatMemberStatus(objMap["old_status"])
	chatEventMemberPromoted.OldStatus = fieldOldStatus

	fieldNewStatus, _ := unmarshalChatMemberStatus(objMap["new_status"])
	chatEventMemberPromoted.NewStatus = fieldNewStatus

	return nil
}

// GetChatEventActionEnum return the enum type of this object
func (chatEventMemberPromoted *ChatEventMemberPromoted) GetChatEventActionEnum() ChatEventActionEnum {
	return ChatEventMemberPromotedType
}

// ChatEventMemberRestricted A chat member was restricted/unrestricted or banned/unbanned, or the list of their restrictions has changed
type ChatEventMemberRestricted struct {
	tdCommon
	MemberId  MessageSender    `json:"member_id"`  // Affected chat member identifier
	OldStatus ChatMemberStatus `json:"old_status"` // Previous status of the chat member
	NewStatus ChatMemberStatus `json:"new_status"` // New status of the chat member
}

// MessageType return the string telegram-type of ChatEventMemberRestricted
func (chatEventMemberRestricted *ChatEventMemberRestricted) MessageType() string {
	return "chatEventMemberRestricted"
}

// NewChatEventMemberRestricted creates a new ChatEventMemberRestricted
//
// @param memberId Affected chat member identifier
// @param oldStatus Previous status of the chat member
// @param newStatus New status of the chat member
func NewChatEventMemberRestricted(memberId MessageSender, oldStatus ChatMemberStatus, newStatus ChatMemberStatus) *ChatEventMemberRestricted {
	chatEventMemberRestrictedTemp := ChatEventMemberRestricted{
		tdCommon:  tdCommon{Type: "chatEventMemberRestricted"},
		MemberId:  memberId,
		OldStatus: oldStatus,
		NewStatus: newStatus,
	}

	return &chatEventMemberRestrictedTemp
}

// UnmarshalJSON unmarshal to json
func (chatEventMemberRestricted *ChatEventMemberRestricted) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	tempObj := struct {
		tdCommon
	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	chatEventMemberRestricted.tdCommon = tempObj.tdCommon

	fieldMemberId, _ := unmarshalMessageSender(objMap["member_id"])
	chatEventMemberRestricted.MemberId = fieldMemberId

	fieldOldStatus, _ := unmarshalChatMemberStatus(objMap["old_status"])
	chatEventMemberRestricted.OldStatus = fieldOldStatus

	fieldNewStatus, _ := unmarshalChatMemberStatus(objMap["new_status"])
	chatEventMemberRestricted.NewStatus = fieldNewStatus

	return nil
}

// GetChatEventActionEnum return the enum type of this object
func (chatEventMemberRestricted *ChatEventMemberRestricted) GetChatEventActionEnum() ChatEventActionEnum {
	return ChatEventMemberRestrictedType
}

// ChatEventTitleChanged The chat title was changed
type ChatEventTitleChanged struct {
	tdCommon
	OldTitle string `json:"old_title"` // Previous chat title
	NewTitle string `json:"new_title"` // New chat title
}

// MessageType return the string telegram-type of ChatEventTitleChanged
func (chatEventTitleChanged *ChatEventTitleChanged) MessageType() string {
	return "chatEventTitleChanged"
}

// NewChatEventTitleChanged creates a new ChatEventTitleChanged
//
// @param oldTitle Previous chat title
// @param newTitle New chat title
func NewChatEventTitleChanged(oldTitle string, newTitle string) *ChatEventTitleChanged {
	chatEventTitleChangedTemp := ChatEventTitleChanged{
		tdCommon: tdCommon{Type: "chatEventTitleChanged"},
		OldTitle: oldTitle,
		NewTitle: newTitle,
	}

	return &chatEventTitleChangedTemp
}

// GetChatEventActionEnum return the enum type of this object
func (chatEventTitleChanged *ChatEventTitleChanged) GetChatEventActionEnum() ChatEventActionEnum {
	return ChatEventTitleChangedType
}

// ChatEventPermissionsChanged The chat permissions was changed
type ChatEventPermissionsChanged struct {
	tdCommon
	OldPermissions *ChatPermissions `json:"old_permissions"` // Previous chat permissions
	NewPermissions *ChatPermissions `json:"new_permissions"` // New chat permissions
}

// MessageType return the string telegram-type of ChatEventPermissionsChanged
func (chatEventPermissionsChanged *ChatEventPermissionsChanged) MessageType() string {
	return "chatEventPermissionsChanged"
}

// NewChatEventPermissionsChanged creates a new ChatEventPermissionsChanged
//
// @param oldPermissions Previous chat permissions
// @param newPermissions New chat permissions
func NewChatEventPermissionsChanged(oldPermissions *ChatPermissions, newPermissions *ChatPermissions) *ChatEventPermissionsChanged {
	chatEventPermissionsChangedTemp := ChatEventPermissionsChanged{
		tdCommon:       tdCommon{Type: "chatEventPermissionsChanged"},
		OldPermissions: oldPermissions,
		NewPermissions: newPermissions,
	}

	return &chatEventPermissionsChangedTemp
}

// GetChatEventActionEnum return the enum type of this object
func (chatEventPermissionsChanged *ChatEventPermissionsChanged) GetChatEventActionEnum() ChatEventActionEnum {
	return ChatEventPermissionsChangedType
}

// ChatEventDescriptionChanged The chat description was changed
type ChatEventDescriptionChanged struct {
	tdCommon
	OldDescription string `json:"old_description"` // Previous chat description
	NewDescription string `json:"new_description"` // New chat description
}

// MessageType return the string telegram-type of ChatEventDescriptionChanged
func (chatEventDescriptionChanged *ChatEventDescriptionChanged) MessageType() string {
	return "chatEventDescriptionChanged"
}

// NewChatEventDescriptionChanged creates a new ChatEventDescriptionChanged
//
// @param oldDescription Previous chat description
// @param newDescription New chat description
func NewChatEventDescriptionChanged(oldDescription string, newDescription string) *ChatEventDescriptionChanged {
	chatEventDescriptionChangedTemp := ChatEventDescriptionChanged{
		tdCommon:       tdCommon{Type: "chatEventDescriptionChanged"},
		OldDescription: oldDescription,
		NewDescription: newDescription,
	}

	return &chatEventDescriptionChangedTemp
}

// GetChatEventActionEnum return the enum type of this object
func (chatEventDescriptionChanged *ChatEventDescriptionChanged) GetChatEventActionEnum() ChatEventActionEnum {
	return ChatEventDescriptionChangedType
}

// ChatEventUsernameChanged The chat username was changed
type ChatEventUsernameChanged struct {
	tdCommon
	OldUsername string `json:"old_username"` // Previous chat username
	NewUsername string `json:"new_username"` // New chat username
}

// MessageType return the string telegram-type of ChatEventUsernameChanged
func (chatEventUsernameChanged *ChatEventUsernameChanged) MessageType() string {
	return "chatEventUsernameChanged"
}

// NewChatEventUsernameChanged creates a new ChatEventUsernameChanged
//
// @param oldUsername Previous chat username
// @param newUsername New chat username
func NewChatEventUsernameChanged(oldUsername string, newUsername string) *ChatEventUsernameChanged {
	chatEventUsernameChangedTemp := ChatEventUsernameChanged{
		tdCommon:    tdCommon{Type: "chatEventUsernameChanged"},
		OldUsername: oldUsername,
		NewUsername: newUsername,
	}

	return &chatEventUsernameChangedTemp
}

// GetChatEventActionEnum return the enum type of this object
func (chatEventUsernameChanged *ChatEventUsernameChanged) GetChatEventActionEnum() ChatEventActionEnum {
	return ChatEventUsernameChangedType
}

// ChatEventPhotoChanged The chat photo was changed
type ChatEventPhotoChanged struct {
	tdCommon
	OldPhoto *ChatPhoto `json:"old_photo"` // Previous chat photo value; may be null
	NewPhoto *ChatPhoto `json:"new_photo"` // New chat photo value; may be null
}

// MessageType return the string telegram-type of ChatEventPhotoChanged
func (chatEventPhotoChanged *ChatEventPhotoChanged) MessageType() string {
	return "chatEventPhotoChanged"
}

// NewChatEventPhotoChanged creates a new ChatEventPhotoChanged
//
// @param oldPhoto Previous chat photo value; may be null
// @param newPhoto New chat photo value; may be null
func NewChatEventPhotoChanged(oldPhoto *ChatPhoto, newPhoto *ChatPhoto) *ChatEventPhotoChanged {
	chatEventPhotoChangedTemp := ChatEventPhotoChanged{
		tdCommon: tdCommon{Type: "chatEventPhotoChanged"},
		OldPhoto: oldPhoto,
		NewPhoto: newPhoto,
	}

	return &chatEventPhotoChangedTemp
}

// GetChatEventActionEnum return the enum type of this object
func (chatEventPhotoChanged *ChatEventPhotoChanged) GetChatEventActionEnum() ChatEventActionEnum {
	return ChatEventPhotoChangedType
}

// ChatEventInvitesToggled The can_invite_users permission of a supergroup chat was toggled
type ChatEventInvitesToggled struct {
	tdCommon
	CanInviteUsers bool `json:"can_invite_users"` // New value of can_invite_users permission
}

// MessageType return the string telegram-type of ChatEventInvitesToggled
func (chatEventInvitesToggled *ChatEventInvitesToggled) MessageType() string {
	return "chatEventInvitesToggled"
}

// NewChatEventInvitesToggled creates a new ChatEventInvitesToggled
//
// @param canInviteUsers New value of can_invite_users permission
func NewChatEventInvitesToggled(canInviteUsers bool) *ChatEventInvitesToggled {
	chatEventInvitesToggledTemp := ChatEventInvitesToggled{
		tdCommon:       tdCommon{Type: "chatEventInvitesToggled"},
		CanInviteUsers: canInviteUsers,
	}

	return &chatEventInvitesToggledTemp
}

// GetChatEventActionEnum return the enum type of this object
func (chatEventInvitesToggled *ChatEventInvitesToggled) GetChatEventActionEnum() ChatEventActionEnum {
	return ChatEventInvitesToggledType
}

// ChatEventLinkedChatChanged The linked chat of a supergroup was changed
type ChatEventLinkedChatChanged struct {
	tdCommon
	OldLinkedChatId int64 `json:"old_linked_chat_id"` // Previous supergroup linked chat identifier
	NewLinkedChatId int64 `json:"new_linked_chat_id"` // New supergroup linked chat identifier
}

// MessageType return the string telegram-type of ChatEventLinkedChatChanged
func (chatEventLinkedChatChanged *ChatEventLinkedChatChanged) MessageType() string {
	return "chatEventLinkedChatChanged"
}

// NewChatEventLinkedChatChanged creates a new ChatEventLinkedChatChanged
//
// @param oldLinkedChatId Previous supergroup linked chat identifier
// @param newLinkedChatId New supergroup linked chat identifier
func NewChatEventLinkedChatChanged(oldLinkedChatId int64, newLinkedChatId int64) *ChatEventLinkedChatChanged {
	chatEventLinkedChatChangedTemp := ChatEventLinkedChatChanged{
		tdCommon:        tdCommon{Type: "chatEventLinkedChatChanged"},
		OldLinkedChatId: oldLinkedChatId,
		NewLinkedChatId: newLinkedChatId,
	}

	return &chatEventLinkedChatChangedTemp
}

// GetChatEventActionEnum return the enum type of this object
func (chatEventLinkedChatChanged *ChatEventLinkedChatChanged) GetChatEventActionEnum() ChatEventActionEnum {
	return ChatEventLinkedChatChangedType
}

// ChatEventSlowModeDelayChanged The slow_mode_delay setting of a supergroup was changed
type ChatEventSlowModeDelayChanged struct {
	tdCommon
	OldSlowModeDelay int32 `json:"old_slow_mode_delay"` // Previous value of slow_mode_delay, in seconds
	NewSlowModeDelay int32 `json:"new_slow_mode_delay"` // New value of slow_mode_delay, in seconds
}

// MessageType return the string telegram-type of ChatEventSlowModeDelayChanged
func (chatEventSlowModeDelayChanged *ChatEventSlowModeDelayChanged) MessageType() string {
	return "chatEventSlowModeDelayChanged"
}

// NewChatEventSlowModeDelayChanged creates a new ChatEventSlowModeDelayChanged
//
// @param oldSlowModeDelay Previous value of slow_mode_delay, in seconds
// @param newSlowModeDelay New value of slow_mode_delay, in seconds
func NewChatEventSlowModeDelayChanged(oldSlowModeDelay int32, newSlowModeDelay int32) *ChatEventSlowModeDelayChanged {
	chatEventSlowModeDelayChangedTemp := ChatEventSlowModeDelayChanged{
		tdCommon:         tdCommon{Type: "chatEventSlowModeDelayChanged"},
		OldSlowModeDelay: oldSlowModeDelay,
		NewSlowModeDelay: newSlowModeDelay,
	}

	return &chatEventSlowModeDelayChangedTemp
}

// GetChatEventActionEnum return the enum type of this object
func (chatEventSlowModeDelayChanged *ChatEventSlowModeDelayChanged) GetChatEventActionEnum() ChatEventActionEnum {
	return ChatEventSlowModeDelayChangedType
}

// ChatEventMessageTtlChanged The message TTL was changed
type ChatEventMessageTtlChanged struct {
	tdCommon
	OldMessageTtl int32 `json:"old_message_ttl"` // Previous value of message_ttl
	NewMessageTtl int32 `json:"new_message_ttl"` // New value of message_ttl
}

// MessageType return the string telegram-type of ChatEventMessageTtlChanged
func (chatEventMessageTtlChanged *ChatEventMessageTtlChanged) MessageType() string {
	return "chatEventMessageTtlChanged"
}

// NewChatEventMessageTtlChanged creates a new ChatEventMessageTtlChanged
//
// @param oldMessageTtl Previous value of message_ttl
// @param newMessageTtl New value of message_ttl
func NewChatEventMessageTtlChanged(oldMessageTtl int32, newMessageTtl int32) *ChatEventMessageTtlChanged {
	chatEventMessageTtlChangedTemp := ChatEventMessageTtlChanged{
		tdCommon:      tdCommon{Type: "chatEventMessageTtlChanged"},
		OldMessageTtl: oldMessageTtl,
		NewMessageTtl: newMessageTtl,
	}

	return &chatEventMessageTtlChangedTemp
}

// GetChatEventActionEnum return the enum type of this object
func (chatEventMessageTtlChanged *ChatEventMessageTtlChanged) GetChatEventActionEnum() ChatEventActionEnum {
	return ChatEventMessageTtlChangedType
}

// ChatEventSignMessagesToggled The sign_messages setting of a channel was toggled
type ChatEventSignMessagesToggled struct {
	tdCommon
	SignMessages bool `json:"sign_messages"` // New value of sign_messages
}

// MessageType return the string telegram-type of ChatEventSignMessagesToggled
func (chatEventSignMessagesToggled *ChatEventSignMessagesToggled) MessageType() string {
	return "chatEventSignMessagesToggled"
}

// NewChatEventSignMessagesToggled creates a new ChatEventSignMessagesToggled
//
// @param signMessages New value of sign_messages
func NewChatEventSignMessagesToggled(signMessages bool) *ChatEventSignMessagesToggled {
	chatEventSignMessagesToggledTemp := ChatEventSignMessagesToggled{
		tdCommon:     tdCommon{Type: "chatEventSignMessagesToggled"},
		SignMessages: signMessages,
	}

	return &chatEventSignMessagesToggledTemp
}

// GetChatEventActionEnum return the enum type of this object
func (chatEventSignMessagesToggled *ChatEventSignMessagesToggled) GetChatEventActionEnum() ChatEventActionEnum {
	return ChatEventSignMessagesToggledType
}

// ChatEventHasProtectedContentToggled The has_protected_content setting of a channel was toggled
type ChatEventHasProtectedContentToggled struct {
	tdCommon
	HasProtectedContent bool `json:"has_protected_content"` // New value of has_protected_content
}

// MessageType return the string telegram-type of ChatEventHasProtectedContentToggled
func (chatEventHasProtectedContentToggled *ChatEventHasProtectedContentToggled) MessageType() string {
	return "chatEventHasProtectedContentToggled"
}

// NewChatEventHasProtectedContentToggled creates a new ChatEventHasProtectedContentToggled
//
// @param hasProtectedContent New value of has_protected_content
func NewChatEventHasProtectedContentToggled(hasProtectedContent bool) *ChatEventHasProtectedContentToggled {
	chatEventHasProtectedContentToggledTemp := ChatEventHasProtectedContentToggled{
		tdCommon:            tdCommon{Type: "chatEventHasProtectedContentToggled"},
		HasProtectedContent: hasProtectedContent,
	}

	return &chatEventHasProtectedContentToggledTemp
}

// GetChatEventActionEnum return the enum type of this object
func (chatEventHasProtectedContentToggled *ChatEventHasProtectedContentToggled) GetChatEventActionEnum() ChatEventActionEnum {
	return ChatEventHasProtectedContentToggledType
}

// ChatEventStickerSetChanged The supergroup sticker set was changed
type ChatEventStickerSetChanged struct {
	tdCommon
	OldStickerSetId JSONInt64 `json:"old_sticker_set_id"` // Previous identifier of the chat sticker set; 0 if none
	NewStickerSetId JSONInt64 `json:"new_sticker_set_id"` // New identifier of the chat sticker set; 0 if none
}

// MessageType return the string telegram-type of ChatEventStickerSetChanged
func (chatEventStickerSetChanged *ChatEventStickerSetChanged) MessageType() string {
	return "chatEventStickerSetChanged"
}

// NewChatEventStickerSetChanged creates a new ChatEventStickerSetChanged
//
// @param oldStickerSetId Previous identifier of the chat sticker set; 0 if none
// @param newStickerSetId New identifier of the chat sticker set; 0 if none
func NewChatEventStickerSetChanged(oldStickerSetId JSONInt64, newStickerSetId JSONInt64) *ChatEventStickerSetChanged {
	chatEventStickerSetChangedTemp := ChatEventStickerSetChanged{
		tdCommon:        tdCommon{Type: "chatEventStickerSetChanged"},
		OldStickerSetId: oldStickerSetId,
		NewStickerSetId: newStickerSetId,
	}

	return &chatEventStickerSetChangedTemp
}

// GetChatEventActionEnum return the enum type of this object
func (chatEventStickerSetChanged *ChatEventStickerSetChanged) GetChatEventActionEnum() ChatEventActionEnum {
	return ChatEventStickerSetChangedType
}

// ChatEventLocationChanged The supergroup location was changed
type ChatEventLocationChanged struct {
	tdCommon
	OldLocation *ChatLocation `json:"old_location"` // Previous location; may be null
	NewLocation *ChatLocation `json:"new_location"` // New location; may be null
}

// MessageType return the string telegram-type of ChatEventLocationChanged
func (chatEventLocationChanged *ChatEventLocationChanged) MessageType() string {
	return "chatEventLocationChanged"
}

// NewChatEventLocationChanged creates a new ChatEventLocationChanged
//
// @param oldLocation Previous location; may be null
// @param newLocation New location; may be null
func NewChatEventLocationChanged(oldLocation *ChatLocation, newLocation *ChatLocation) *ChatEventLocationChanged {
	chatEventLocationChangedTemp := ChatEventLocationChanged{
		tdCommon:    tdCommon{Type: "chatEventLocationChanged"},
		OldLocation: oldLocation,
		NewLocation: newLocation,
	}

	return &chatEventLocationChangedTemp
}

// GetChatEventActionEnum return the enum type of this object
func (chatEventLocationChanged *ChatEventLocationChanged) GetChatEventActionEnum() ChatEventActionEnum {
	return ChatEventLocationChangedType
}

// ChatEventIsAllHistoryAvailableToggled The is_all_history_available setting of a supergroup was toggled
type ChatEventIsAllHistoryAvailableToggled struct {
	tdCommon
	IsAllHistoryAvailable bool `json:"is_all_history_available"` // New value of is_all_history_available
}

// MessageType return the string telegram-type of ChatEventIsAllHistoryAvailableToggled
func (chatEventIsAllHistoryAvailableToggled *ChatEventIsAllHistoryAvailableToggled) MessageType() string {
	return "chatEventIsAllHistoryAvailableToggled"
}

// NewChatEventIsAllHistoryAvailableToggled creates a new ChatEventIsAllHistoryAvailableToggled
//
// @param isAllHistoryAvailable New value of is_all_history_available
func NewChatEventIsAllHistoryAvailableToggled(isAllHistoryAvailable bool) *ChatEventIsAllHistoryAvailableToggled {
	chatEventIsAllHistoryAvailableToggledTemp := ChatEventIsAllHistoryAvailableToggled{
		tdCommon:              tdCommon{Type: "chatEventIsAllHistoryAvailableToggled"},
		IsAllHistoryAvailable: isAllHistoryAvailable,
	}

	return &chatEventIsAllHistoryAvailableToggledTemp
}

// GetChatEventActionEnum return the enum type of this object
func (chatEventIsAllHistoryAvailableToggled *ChatEventIsAllHistoryAvailableToggled) GetChatEventActionEnum() ChatEventActionEnum {
	return ChatEventIsAllHistoryAvailableToggledType
}

// ChatEventInviteLinkEdited A chat invite link was edited
type ChatEventInviteLinkEdited struct {
	tdCommon
	OldInviteLink *ChatInviteLink `json:"old_invite_link"` // Previous information about the invite link
	NewInviteLink *ChatInviteLink `json:"new_invite_link"` // New information about the invite link
}

// MessageType return the string telegram-type of ChatEventInviteLinkEdited
func (chatEventInviteLinkEdited *ChatEventInviteLinkEdited) MessageType() string {
	return "chatEventInviteLinkEdited"
}

// NewChatEventInviteLinkEdited creates a new ChatEventInviteLinkEdited
//
// @param oldInviteLink Previous information about the invite link
// @param newInviteLink New information about the invite link
func NewChatEventInviteLinkEdited(oldInviteLink *ChatInviteLink, newInviteLink *ChatInviteLink) *ChatEventInviteLinkEdited {
	chatEventInviteLinkEditedTemp := ChatEventInviteLinkEdited{
		tdCommon:      tdCommon{Type: "chatEventInviteLinkEdited"},
		OldInviteLink: oldInviteLink,
		NewInviteLink: newInviteLink,
	}

	return &chatEventInviteLinkEditedTemp
}

// GetChatEventActionEnum return the enum type of this object
func (chatEventInviteLinkEdited *ChatEventInviteLinkEdited) GetChatEventActionEnum() ChatEventActionEnum {
	return ChatEventInviteLinkEditedType
}

// ChatEventInviteLinkRevoked A chat invite link was revoked
type ChatEventInviteLinkRevoked struct {
	tdCommon
	InviteLink *ChatInviteLink `json:"invite_link"` // The invite link
}

// MessageType return the string telegram-type of ChatEventInviteLinkRevoked
func (chatEventInviteLinkRevoked *ChatEventInviteLinkRevoked) MessageType() string {
	return "chatEventInviteLinkRevoked"
}

// NewChatEventInviteLinkRevoked creates a new ChatEventInviteLinkRevoked
//
// @param inviteLink The invite link
func NewChatEventInviteLinkRevoked(inviteLink *ChatInviteLink) *ChatEventInviteLinkRevoked {
	chatEventInviteLinkRevokedTemp := ChatEventInviteLinkRevoked{
		tdCommon:   tdCommon{Type: "chatEventInviteLinkRevoked"},
		InviteLink: inviteLink,
	}

	return &chatEventInviteLinkRevokedTemp
}

// GetChatEventActionEnum return the enum type of this object
func (chatEventInviteLinkRevoked *ChatEventInviteLinkRevoked) GetChatEventActionEnum() ChatEventActionEnum {
	return ChatEventInviteLinkRevokedType
}

// ChatEventInviteLinkDeleted A revoked chat invite link was deleted
type ChatEventInviteLinkDeleted struct {
	tdCommon
	InviteLink *ChatInviteLink `json:"invite_link"` // The invite link
}

// MessageType return the string telegram-type of ChatEventInviteLinkDeleted
func (chatEventInviteLinkDeleted *ChatEventInviteLinkDeleted) MessageType() string {
	return "chatEventInviteLinkDeleted"
}

// NewChatEventInviteLinkDeleted creates a new ChatEventInviteLinkDeleted
//
// @param inviteLink The invite link
func NewChatEventInviteLinkDeleted(inviteLink *ChatInviteLink) *ChatEventInviteLinkDeleted {
	chatEventInviteLinkDeletedTemp := ChatEventInviteLinkDeleted{
		tdCommon:   tdCommon{Type: "chatEventInviteLinkDeleted"},
		InviteLink: inviteLink,
	}

	return &chatEventInviteLinkDeletedTemp
}

// GetChatEventActionEnum return the enum type of this object
func (chatEventInviteLinkDeleted *ChatEventInviteLinkDeleted) GetChatEventActionEnum() ChatEventActionEnum {
	return ChatEventInviteLinkDeletedType
}

// ChatEventVideoChatCreated A video chat was created
type ChatEventVideoChatCreated struct {
	tdCommon
	GroupCallId int32 `json:"group_call_id"` // Identifier of the video chat. The video chat can be received through the method getGroupCall
}

// MessageType return the string telegram-type of ChatEventVideoChatCreated
func (chatEventVideoChatCreated *ChatEventVideoChatCreated) MessageType() string {
	return "chatEventVideoChatCreated"
}

// NewChatEventVideoChatCreated creates a new ChatEventVideoChatCreated
//
// @param groupCallId Identifier of the video chat. The video chat can be received through the method getGroupCall
func NewChatEventVideoChatCreated(groupCallId int32) *ChatEventVideoChatCreated {
	chatEventVideoChatCreatedTemp := ChatEventVideoChatCreated{
		tdCommon:    tdCommon{Type: "chatEventVideoChatCreated"},
		GroupCallId: groupCallId,
	}

	return &chatEventVideoChatCreatedTemp
}

// GetChatEventActionEnum return the enum type of this object
func (chatEventVideoChatCreated *ChatEventVideoChatCreated) GetChatEventActionEnum() ChatEventActionEnum {
	return ChatEventVideoChatCreatedType
}

// ChatEventVideoChatEnded A video chat was ended
type ChatEventVideoChatEnded struct {
	tdCommon
	GroupCallId int32 `json:"group_call_id"` // Identifier of the video chat. The video chat can be received through the method getGroupCall
}

// MessageType return the string telegram-type of ChatEventVideoChatEnded
func (chatEventVideoChatEnded *ChatEventVideoChatEnded) MessageType() string {
	return "chatEventVideoChatEnded"
}

// NewChatEventVideoChatEnded creates a new ChatEventVideoChatEnded
//
// @param groupCallId Identifier of the video chat. The video chat can be received through the method getGroupCall
func NewChatEventVideoChatEnded(groupCallId int32) *ChatEventVideoChatEnded {
	chatEventVideoChatEndedTemp := ChatEventVideoChatEnded{
		tdCommon:    tdCommon{Type: "chatEventVideoChatEnded"},
		GroupCallId: groupCallId,
	}

	return &chatEventVideoChatEndedTemp
}

// GetChatEventActionEnum return the enum type of this object
func (chatEventVideoChatEnded *ChatEventVideoChatEnded) GetChatEventActionEnum() ChatEventActionEnum {
	return ChatEventVideoChatEndedType
}

// ChatEventVideoChatParticipantIsMutedToggled A video chat participant was muted or unmuted
type ChatEventVideoChatParticipantIsMutedToggled struct {
	tdCommon
	ParticipantId MessageSender `json:"participant_id"` // Identifier of the affected group call participant
	IsMuted       bool          `json:"is_muted"`       // New value of is_muted
}

// MessageType return the string telegram-type of ChatEventVideoChatParticipantIsMutedToggled
func (chatEventVideoChatParticipantIsMutedToggled *ChatEventVideoChatParticipantIsMutedToggled) MessageType() string {
	return "chatEventVideoChatParticipantIsMutedToggled"
}

// NewChatEventVideoChatParticipantIsMutedToggled creates a new ChatEventVideoChatParticipantIsMutedToggled
//
// @param participantId Identifier of the affected group call participant
// @param isMuted New value of is_muted
func NewChatEventVideoChatParticipantIsMutedToggled(participantId MessageSender, isMuted bool) *ChatEventVideoChatParticipantIsMutedToggled {
	chatEventVideoChatParticipantIsMutedToggledTemp := ChatEventVideoChatParticipantIsMutedToggled{
		tdCommon:      tdCommon{Type: "chatEventVideoChatParticipantIsMutedToggled"},
		ParticipantId: participantId,
		IsMuted:       isMuted,
	}

	return &chatEventVideoChatParticipantIsMutedToggledTemp
}

// UnmarshalJSON unmarshal to json
func (chatEventVideoChatParticipantIsMutedToggled *ChatEventVideoChatParticipantIsMutedToggled) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	tempObj := struct {
		tdCommon
		IsMuted bool `json:"is_muted"` // New value of is_muted
	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	chatEventVideoChatParticipantIsMutedToggled.tdCommon = tempObj.tdCommon
	chatEventVideoChatParticipantIsMutedToggled.IsMuted = tempObj.IsMuted

	fieldParticipantId, _ := unmarshalMessageSender(objMap["participant_id"])
	chatEventVideoChatParticipantIsMutedToggled.ParticipantId = fieldParticipantId

	return nil
}

// GetChatEventActionEnum return the enum type of this object
func (chatEventVideoChatParticipantIsMutedToggled *ChatEventVideoChatParticipantIsMutedToggled) GetChatEventActionEnum() ChatEventActionEnum {
	return ChatEventVideoChatParticipantIsMutedToggledType
}

// ChatEventVideoChatParticipantVolumeLevelChanged A video chat participant volume level was changed
type ChatEventVideoChatParticipantVolumeLevelChanged struct {
	tdCommon
	ParticipantId MessageSender `json:"participant_id"` // Identifier of the affected group call participant
	VolumeLevel   int32         `json:"volume_level"`   // New value of volume_level; 1-20000 in hundreds of percents
}

// MessageType return the string telegram-type of ChatEventVideoChatParticipantVolumeLevelChanged
func (chatEventVideoChatParticipantVolumeLevelChanged *ChatEventVideoChatParticipantVolumeLevelChanged) MessageType() string {
	return "chatEventVideoChatParticipantVolumeLevelChanged"
}

// NewChatEventVideoChatParticipantVolumeLevelChanged creates a new ChatEventVideoChatParticipantVolumeLevelChanged
//
// @param participantId Identifier of the affected group call participant
// @param volumeLevel New value of volume_level; 1-20000 in hundreds of percents
func NewChatEventVideoChatParticipantVolumeLevelChanged(participantId MessageSender, volumeLevel int32) *ChatEventVideoChatParticipantVolumeLevelChanged {
	chatEventVideoChatParticipantVolumeLevelChangedTemp := ChatEventVideoChatParticipantVolumeLevelChanged{
		tdCommon:      tdCommon{Type: "chatEventVideoChatParticipantVolumeLevelChanged"},
		ParticipantId: participantId,
		VolumeLevel:   volumeLevel,
	}

	return &chatEventVideoChatParticipantVolumeLevelChangedTemp
}

// UnmarshalJSON unmarshal to json
func (chatEventVideoChatParticipantVolumeLevelChanged *ChatEventVideoChatParticipantVolumeLevelChanged) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	tempObj := struct {
		tdCommon
		VolumeLevel int32 `json:"volume_level"` // New value of volume_level; 1-20000 in hundreds of percents
	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	chatEventVideoChatParticipantVolumeLevelChanged.tdCommon = tempObj.tdCommon
	chatEventVideoChatParticipantVolumeLevelChanged.VolumeLevel = tempObj.VolumeLevel

	fieldParticipantId, _ := unmarshalMessageSender(objMap["participant_id"])
	chatEventVideoChatParticipantVolumeLevelChanged.ParticipantId = fieldParticipantId

	return nil
}

// GetChatEventActionEnum return the enum type of this object
func (chatEventVideoChatParticipantVolumeLevelChanged *ChatEventVideoChatParticipantVolumeLevelChanged) GetChatEventActionEnum() ChatEventActionEnum {
	return ChatEventVideoChatParticipantVolumeLevelChangedType
}

// ChatEventVideoChatMuteNewParticipantsToggled The mute_new_participants setting of a video chat was toggled
type ChatEventVideoChatMuteNewParticipantsToggled struct {
	tdCommon
	MuteNewParticipants bool `json:"mute_new_participants"` // New value of the mute_new_participants setting
}

// MessageType return the string telegram-type of ChatEventVideoChatMuteNewParticipantsToggled
func (chatEventVideoChatMuteNewParticipantsToggled *ChatEventVideoChatMuteNewParticipantsToggled) MessageType() string {
	return "chatEventVideoChatMuteNewParticipantsToggled"
}

// NewChatEventVideoChatMuteNewParticipantsToggled creates a new ChatEventVideoChatMuteNewParticipantsToggled
//
// @param muteNewParticipants New value of the mute_new_participants setting
func NewChatEventVideoChatMuteNewParticipantsToggled(muteNewParticipants bool) *ChatEventVideoChatMuteNewParticipantsToggled {
	chatEventVideoChatMuteNewParticipantsToggledTemp := ChatEventVideoChatMuteNewParticipantsToggled{
		tdCommon:            tdCommon{Type: "chatEventVideoChatMuteNewParticipantsToggled"},
		MuteNewParticipants: muteNewParticipants,
	}

	return &chatEventVideoChatMuteNewParticipantsToggledTemp
}

// GetChatEventActionEnum return the enum type of this object
func (chatEventVideoChatMuteNewParticipantsToggled *ChatEventVideoChatMuteNewParticipantsToggled) GetChatEventActionEnum() ChatEventActionEnum {
	return ChatEventVideoChatMuteNewParticipantsToggledType
}
