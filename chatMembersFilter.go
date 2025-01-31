package tdlib

import (
	"encoding/json"
	"fmt"
)

// ChatMembersFilter Specifies the kind of chat members to return in searchChatMembers
type ChatMembersFilter interface {
	GetChatMembersFilterEnum() ChatMembersFilterEnum
}

// ChatMembersFilterEnum Alias for abstract ChatMembersFilter 'Sub-Classes', used as constant-enum here
type ChatMembersFilterEnum string

// ChatMembersFilter enums
const (
	ChatMembersFilterContactsType       ChatMembersFilterEnum = "chatMembersFilterContacts"
	ChatMembersFilterAdministratorsType ChatMembersFilterEnum = "chatMembersFilterAdministrators"
	ChatMembersFilterMembersType        ChatMembersFilterEnum = "chatMembersFilterMembers"
	ChatMembersFilterMentionType        ChatMembersFilterEnum = "chatMembersFilterMention"
	ChatMembersFilterRestrictedType     ChatMembersFilterEnum = "chatMembersFilterRestricted"
	ChatMembersFilterBannedType         ChatMembersFilterEnum = "chatMembersFilterBanned"
	ChatMembersFilterBotsType           ChatMembersFilterEnum = "chatMembersFilterBots"
)

func unmarshalChatMembersFilter(rawMsg *json.RawMessage) (ChatMembersFilter, error) {

	if rawMsg == nil {
		return nil, nil
	}
	var objMap map[string]interface{}
	err := json.Unmarshal(*rawMsg, &objMap)
	if err != nil {
		return nil, err
	}

	switch ChatMembersFilterEnum(objMap["@type"].(string)) {
	case ChatMembersFilterContactsType:
		var chatMembersFilterContacts ChatMembersFilterContacts
		err := json.Unmarshal(*rawMsg, &chatMembersFilterContacts)
		return &chatMembersFilterContacts, err

	case ChatMembersFilterAdministratorsType:
		var chatMembersFilterAdministrators ChatMembersFilterAdministrators
		err := json.Unmarshal(*rawMsg, &chatMembersFilterAdministrators)
		return &chatMembersFilterAdministrators, err

	case ChatMembersFilterMembersType:
		var chatMembersFilterMembers ChatMembersFilterMembers
		err := json.Unmarshal(*rawMsg, &chatMembersFilterMembers)
		return &chatMembersFilterMembers, err

	case ChatMembersFilterMentionType:
		var chatMembersFilterMention ChatMembersFilterMention
		err := json.Unmarshal(*rawMsg, &chatMembersFilterMention)
		return &chatMembersFilterMention, err

	case ChatMembersFilterRestrictedType:
		var chatMembersFilterRestricted ChatMembersFilterRestricted
		err := json.Unmarshal(*rawMsg, &chatMembersFilterRestricted)
		return &chatMembersFilterRestricted, err

	case ChatMembersFilterBannedType:
		var chatMembersFilterBanned ChatMembersFilterBanned
		err := json.Unmarshal(*rawMsg, &chatMembersFilterBanned)
		return &chatMembersFilterBanned, err

	case ChatMembersFilterBotsType:
		var chatMembersFilterBots ChatMembersFilterBots
		err := json.Unmarshal(*rawMsg, &chatMembersFilterBots)
		return &chatMembersFilterBots, err

	default:
		return nil, fmt.Errorf("Error UnMarshaling, unknown type:" + objMap["@type"].(string))
	}
}

// ChatMembersFilterContacts Returns contacts of the user
type ChatMembersFilterContacts struct {
	tdCommon
}

// MessageType return the string telegram-type of ChatMembersFilterContacts
func (chatMembersFilterContacts *ChatMembersFilterContacts) MessageType() string {
	return "chatMembersFilterContacts"
}

// NewChatMembersFilterContacts creates a new ChatMembersFilterContacts
//
func NewChatMembersFilterContacts() *ChatMembersFilterContacts {
	chatMembersFilterContactsTemp := ChatMembersFilterContacts{
		tdCommon: tdCommon{Type: "chatMembersFilterContacts"},
	}

	return &chatMembersFilterContactsTemp
}

// GetChatMembersFilterEnum return the enum type of this object
func (chatMembersFilterContacts *ChatMembersFilterContacts) GetChatMembersFilterEnum() ChatMembersFilterEnum {
	return ChatMembersFilterContactsType
}

// ChatMembersFilterAdministrators Returns the owner and administrators
type ChatMembersFilterAdministrators struct {
	tdCommon
}

// MessageType return the string telegram-type of ChatMembersFilterAdministrators
func (chatMembersFilterAdministrators *ChatMembersFilterAdministrators) MessageType() string {
	return "chatMembersFilterAdministrators"
}

// NewChatMembersFilterAdministrators creates a new ChatMembersFilterAdministrators
//
func NewChatMembersFilterAdministrators() *ChatMembersFilterAdministrators {
	chatMembersFilterAdministratorsTemp := ChatMembersFilterAdministrators{
		tdCommon: tdCommon{Type: "chatMembersFilterAdministrators"},
	}

	return &chatMembersFilterAdministratorsTemp
}

// GetChatMembersFilterEnum return the enum type of this object
func (chatMembersFilterAdministrators *ChatMembersFilterAdministrators) GetChatMembersFilterEnum() ChatMembersFilterEnum {
	return ChatMembersFilterAdministratorsType
}

// ChatMembersFilterMembers Returns all chat members, including restricted chat members
type ChatMembersFilterMembers struct {
	tdCommon
}

// MessageType return the string telegram-type of ChatMembersFilterMembers
func (chatMembersFilterMembers *ChatMembersFilterMembers) MessageType() string {
	return "chatMembersFilterMembers"
}

// NewChatMembersFilterMembers creates a new ChatMembersFilterMembers
//
func NewChatMembersFilterMembers() *ChatMembersFilterMembers {
	chatMembersFilterMembersTemp := ChatMembersFilterMembers{
		tdCommon: tdCommon{Type: "chatMembersFilterMembers"},
	}

	return &chatMembersFilterMembersTemp
}

// GetChatMembersFilterEnum return the enum type of this object
func (chatMembersFilterMembers *ChatMembersFilterMembers) GetChatMembersFilterEnum() ChatMembersFilterEnum {
	return ChatMembersFilterMembersType
}

// ChatMembersFilterMention Returns users which can be mentioned in the chat
type ChatMembersFilterMention struct {
	tdCommon
	MessageThreadId int64 `json:"message_thread_id"` // If non-zero, the identifier of the current message thread
}

// MessageType return the string telegram-type of ChatMembersFilterMention
func (chatMembersFilterMention *ChatMembersFilterMention) MessageType() string {
	return "chatMembersFilterMention"
}

// NewChatMembersFilterMention creates a new ChatMembersFilterMention
//
// @param messageThreadId If non-zero, the identifier of the current message thread
func NewChatMembersFilterMention(messageThreadId int64) *ChatMembersFilterMention {
	chatMembersFilterMentionTemp := ChatMembersFilterMention{
		tdCommon:        tdCommon{Type: "chatMembersFilterMention"},
		MessageThreadId: messageThreadId,
	}

	return &chatMembersFilterMentionTemp
}

// GetChatMembersFilterEnum return the enum type of this object
func (chatMembersFilterMention *ChatMembersFilterMention) GetChatMembersFilterEnum() ChatMembersFilterEnum {
	return ChatMembersFilterMentionType
}

// ChatMembersFilterRestricted Returns users under certain restrictions in the chat; can be used only by administrators in a supergroup
type ChatMembersFilterRestricted struct {
	tdCommon
}

// MessageType return the string telegram-type of ChatMembersFilterRestricted
func (chatMembersFilterRestricted *ChatMembersFilterRestricted) MessageType() string {
	return "chatMembersFilterRestricted"
}

// NewChatMembersFilterRestricted creates a new ChatMembersFilterRestricted
//
func NewChatMembersFilterRestricted() *ChatMembersFilterRestricted {
	chatMembersFilterRestrictedTemp := ChatMembersFilterRestricted{
		tdCommon: tdCommon{Type: "chatMembersFilterRestricted"},
	}

	return &chatMembersFilterRestrictedTemp
}

// GetChatMembersFilterEnum return the enum type of this object
func (chatMembersFilterRestricted *ChatMembersFilterRestricted) GetChatMembersFilterEnum() ChatMembersFilterEnum {
	return ChatMembersFilterRestrictedType
}

// ChatMembersFilterBanned Returns users banned from the chat; can be used only by administrators in a supergroup or in a channel
type ChatMembersFilterBanned struct {
	tdCommon
}

// MessageType return the string telegram-type of ChatMembersFilterBanned
func (chatMembersFilterBanned *ChatMembersFilterBanned) MessageType() string {
	return "chatMembersFilterBanned"
}

// NewChatMembersFilterBanned creates a new ChatMembersFilterBanned
//
func NewChatMembersFilterBanned() *ChatMembersFilterBanned {
	chatMembersFilterBannedTemp := ChatMembersFilterBanned{
		tdCommon: tdCommon{Type: "chatMembersFilterBanned"},
	}

	return &chatMembersFilterBannedTemp
}

// GetChatMembersFilterEnum return the enum type of this object
func (chatMembersFilterBanned *ChatMembersFilterBanned) GetChatMembersFilterEnum() ChatMembersFilterEnum {
	return ChatMembersFilterBannedType
}

// ChatMembersFilterBots Returns bot members of the chat
type ChatMembersFilterBots struct {
	tdCommon
}

// MessageType return the string telegram-type of ChatMembersFilterBots
func (chatMembersFilterBots *ChatMembersFilterBots) MessageType() string {
	return "chatMembersFilterBots"
}

// NewChatMembersFilterBots creates a new ChatMembersFilterBots
//
func NewChatMembersFilterBots() *ChatMembersFilterBots {
	chatMembersFilterBotsTemp := ChatMembersFilterBots{
		tdCommon: tdCommon{Type: "chatMembersFilterBots"},
	}

	return &chatMembersFilterBotsTemp
}

// GetChatMembersFilterEnum return the enum type of this object
func (chatMembersFilterBots *ChatMembersFilterBots) GetChatMembersFilterEnum() ChatMembersFilterEnum {
	return ChatMembersFilterBotsType
}
