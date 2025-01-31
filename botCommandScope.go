package tdlib

import (
	"encoding/json"
	"fmt"
)

// BotCommandScope Represents the scope to which bot commands are relevant
type BotCommandScope interface {
	GetBotCommandScopeEnum() BotCommandScopeEnum
}

// BotCommandScopeEnum Alias for abstract BotCommandScope 'Sub-Classes', used as constant-enum here
type BotCommandScopeEnum string

// BotCommandScope enums
const (
	BotCommandScopeDefaultType               BotCommandScopeEnum = "botCommandScopeDefault"
	BotCommandScopeAllPrivateChatsType       BotCommandScopeEnum = "botCommandScopeAllPrivateChats"
	BotCommandScopeAllGroupChatsType         BotCommandScopeEnum = "botCommandScopeAllGroupChats"
	BotCommandScopeAllChatAdministratorsType BotCommandScopeEnum = "botCommandScopeAllChatAdministrators"
	BotCommandScopeChatType                  BotCommandScopeEnum = "botCommandScopeChat"
	BotCommandScopeChatAdministratorsType    BotCommandScopeEnum = "botCommandScopeChatAdministrators"
	BotCommandScopeChatMemberType            BotCommandScopeEnum = "botCommandScopeChatMember"
)

func unmarshalBotCommandScope(rawMsg *json.RawMessage) (BotCommandScope, error) {

	if rawMsg == nil {
		return nil, nil
	}
	var objMap map[string]interface{}
	err := json.Unmarshal(*rawMsg, &objMap)
	if err != nil {
		return nil, err
	}

	switch BotCommandScopeEnum(objMap["@type"].(string)) {
	case BotCommandScopeDefaultType:
		var botCommandScopeDefault BotCommandScopeDefault
		err := json.Unmarshal(*rawMsg, &botCommandScopeDefault)
		return &botCommandScopeDefault, err

	case BotCommandScopeAllPrivateChatsType:
		var botCommandScopeAllPrivateChats BotCommandScopeAllPrivateChats
		err := json.Unmarshal(*rawMsg, &botCommandScopeAllPrivateChats)
		return &botCommandScopeAllPrivateChats, err

	case BotCommandScopeAllGroupChatsType:
		var botCommandScopeAllGroupChats BotCommandScopeAllGroupChats
		err := json.Unmarshal(*rawMsg, &botCommandScopeAllGroupChats)
		return &botCommandScopeAllGroupChats, err

	case BotCommandScopeAllChatAdministratorsType:
		var botCommandScopeAllChatAdministrators BotCommandScopeAllChatAdministrators
		err := json.Unmarshal(*rawMsg, &botCommandScopeAllChatAdministrators)
		return &botCommandScopeAllChatAdministrators, err

	case BotCommandScopeChatType:
		var botCommandScopeChat BotCommandScopeChat
		err := json.Unmarshal(*rawMsg, &botCommandScopeChat)
		return &botCommandScopeChat, err

	case BotCommandScopeChatAdministratorsType:
		var botCommandScopeChatAdministrators BotCommandScopeChatAdministrators
		err := json.Unmarshal(*rawMsg, &botCommandScopeChatAdministrators)
		return &botCommandScopeChatAdministrators, err

	case BotCommandScopeChatMemberType:
		var botCommandScopeChatMember BotCommandScopeChatMember
		err := json.Unmarshal(*rawMsg, &botCommandScopeChatMember)
		return &botCommandScopeChatMember, err

	default:
		return nil, fmt.Errorf("Error UnMarshaling, unknown type:" + objMap["@type"].(string))
	}
}

// BotCommandScopeDefault A scope covering all users
type BotCommandScopeDefault struct {
	tdCommon
}

// MessageType return the string telegram-type of BotCommandScopeDefault
func (botCommandScopeDefault *BotCommandScopeDefault) MessageType() string {
	return "botCommandScopeDefault"
}

// NewBotCommandScopeDefault creates a new BotCommandScopeDefault
//
func NewBotCommandScopeDefault() *BotCommandScopeDefault {
	botCommandScopeDefaultTemp := BotCommandScopeDefault{
		tdCommon: tdCommon{Type: "botCommandScopeDefault"},
	}

	return &botCommandScopeDefaultTemp
}

// GetBotCommandScopeEnum return the enum type of this object
func (botCommandScopeDefault *BotCommandScopeDefault) GetBotCommandScopeEnum() BotCommandScopeEnum {
	return BotCommandScopeDefaultType
}

// BotCommandScopeAllPrivateChats A scope covering all private chats
type BotCommandScopeAllPrivateChats struct {
	tdCommon
}

// MessageType return the string telegram-type of BotCommandScopeAllPrivateChats
func (botCommandScopeAllPrivateChats *BotCommandScopeAllPrivateChats) MessageType() string {
	return "botCommandScopeAllPrivateChats"
}

// NewBotCommandScopeAllPrivateChats creates a new BotCommandScopeAllPrivateChats
//
func NewBotCommandScopeAllPrivateChats() *BotCommandScopeAllPrivateChats {
	botCommandScopeAllPrivateChatsTemp := BotCommandScopeAllPrivateChats{
		tdCommon: tdCommon{Type: "botCommandScopeAllPrivateChats"},
	}

	return &botCommandScopeAllPrivateChatsTemp
}

// GetBotCommandScopeEnum return the enum type of this object
func (botCommandScopeAllPrivateChats *BotCommandScopeAllPrivateChats) GetBotCommandScopeEnum() BotCommandScopeEnum {
	return BotCommandScopeAllPrivateChatsType
}

// BotCommandScopeAllGroupChats A scope covering all group and supergroup chats
type BotCommandScopeAllGroupChats struct {
	tdCommon
}

// MessageType return the string telegram-type of BotCommandScopeAllGroupChats
func (botCommandScopeAllGroupChats *BotCommandScopeAllGroupChats) MessageType() string {
	return "botCommandScopeAllGroupChats"
}

// NewBotCommandScopeAllGroupChats creates a new BotCommandScopeAllGroupChats
//
func NewBotCommandScopeAllGroupChats() *BotCommandScopeAllGroupChats {
	botCommandScopeAllGroupChatsTemp := BotCommandScopeAllGroupChats{
		tdCommon: tdCommon{Type: "botCommandScopeAllGroupChats"},
	}

	return &botCommandScopeAllGroupChatsTemp
}

// GetBotCommandScopeEnum return the enum type of this object
func (botCommandScopeAllGroupChats *BotCommandScopeAllGroupChats) GetBotCommandScopeEnum() BotCommandScopeEnum {
	return BotCommandScopeAllGroupChatsType
}

// BotCommandScopeAllChatAdministrators A scope covering all group and supergroup chat administrators
type BotCommandScopeAllChatAdministrators struct {
	tdCommon
}

// MessageType return the string telegram-type of BotCommandScopeAllChatAdministrators
func (botCommandScopeAllChatAdministrators *BotCommandScopeAllChatAdministrators) MessageType() string {
	return "botCommandScopeAllChatAdministrators"
}

// NewBotCommandScopeAllChatAdministrators creates a new BotCommandScopeAllChatAdministrators
//
func NewBotCommandScopeAllChatAdministrators() *BotCommandScopeAllChatAdministrators {
	botCommandScopeAllChatAdministratorsTemp := BotCommandScopeAllChatAdministrators{
		tdCommon: tdCommon{Type: "botCommandScopeAllChatAdministrators"},
	}

	return &botCommandScopeAllChatAdministratorsTemp
}

// GetBotCommandScopeEnum return the enum type of this object
func (botCommandScopeAllChatAdministrators *BotCommandScopeAllChatAdministrators) GetBotCommandScopeEnum() BotCommandScopeEnum {
	return BotCommandScopeAllChatAdministratorsType
}

// BotCommandScopeChat A scope covering all members of a chat
type BotCommandScopeChat struct {
	tdCommon
	ChatId int64 `json:"chat_id"` // Chat identifier
}

// MessageType return the string telegram-type of BotCommandScopeChat
func (botCommandScopeChat *BotCommandScopeChat) MessageType() string {
	return "botCommandScopeChat"
}

// NewBotCommandScopeChat creates a new BotCommandScopeChat
//
// @param chatId Chat identifier
func NewBotCommandScopeChat(chatId int64) *BotCommandScopeChat {
	botCommandScopeChatTemp := BotCommandScopeChat{
		tdCommon: tdCommon{Type: "botCommandScopeChat"},
		ChatId:   chatId,
	}

	return &botCommandScopeChatTemp
}

// GetBotCommandScopeEnum return the enum type of this object
func (botCommandScopeChat *BotCommandScopeChat) GetBotCommandScopeEnum() BotCommandScopeEnum {
	return BotCommandScopeChatType
}

// BotCommandScopeChatAdministrators A scope covering all administrators of a chat
type BotCommandScopeChatAdministrators struct {
	tdCommon
	ChatId int64 `json:"chat_id"` // Chat identifier
}

// MessageType return the string telegram-type of BotCommandScopeChatAdministrators
func (botCommandScopeChatAdministrators *BotCommandScopeChatAdministrators) MessageType() string {
	return "botCommandScopeChatAdministrators"
}

// NewBotCommandScopeChatAdministrators creates a new BotCommandScopeChatAdministrators
//
// @param chatId Chat identifier
func NewBotCommandScopeChatAdministrators(chatId int64) *BotCommandScopeChatAdministrators {
	botCommandScopeChatAdministratorsTemp := BotCommandScopeChatAdministrators{
		tdCommon: tdCommon{Type: "botCommandScopeChatAdministrators"},
		ChatId:   chatId,
	}

	return &botCommandScopeChatAdministratorsTemp
}

// GetBotCommandScopeEnum return the enum type of this object
func (botCommandScopeChatAdministrators *BotCommandScopeChatAdministrators) GetBotCommandScopeEnum() BotCommandScopeEnum {
	return BotCommandScopeChatAdministratorsType
}

// BotCommandScopeChatMember A scope covering a member of a chat
type BotCommandScopeChatMember struct {
	tdCommon
	ChatId int64 `json:"chat_id"` // Chat identifier
	UserId int64 `json:"user_id"` // User identifier
}

// MessageType return the string telegram-type of BotCommandScopeChatMember
func (botCommandScopeChatMember *BotCommandScopeChatMember) MessageType() string {
	return "botCommandScopeChatMember"
}

// NewBotCommandScopeChatMember creates a new BotCommandScopeChatMember
//
// @param chatId Chat identifier
// @param userId User identifier
func NewBotCommandScopeChatMember(chatId int64, userId int64) *BotCommandScopeChatMember {
	botCommandScopeChatMemberTemp := BotCommandScopeChatMember{
		tdCommon: tdCommon{Type: "botCommandScopeChatMember"},
		ChatId:   chatId,
		UserId:   userId,
	}

	return &botCommandScopeChatMemberTemp
}

// GetBotCommandScopeEnum return the enum type of this object
func (botCommandScopeChatMember *BotCommandScopeChatMember) GetBotCommandScopeEnum() BotCommandScopeEnum {
	return BotCommandScopeChatMemberType
}
