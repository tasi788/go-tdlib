package tdlib

import (
	"encoding/json"
	"fmt"
)

// ChatList Describes a list of chats
type ChatList interface {
	GetChatListEnum() ChatListEnum
}

// ChatListEnum Alias for abstract ChatList 'Sub-Classes', used as constant-enum here
type ChatListEnum string

// ChatList enums
const (
	ChatListMainType    ChatListEnum = "chatListMain"
	ChatListArchiveType ChatListEnum = "chatListArchive"
	ChatListFilterType  ChatListEnum = "chatListFilter"
)

func unmarshalChatList(rawMsg *json.RawMessage) (ChatList, error) {

	if rawMsg == nil {
		return nil, nil
	}
	var objMap map[string]interface{}
	err := json.Unmarshal(*rawMsg, &objMap)
	if err != nil {
		return nil, err
	}

	switch ChatListEnum(objMap["@type"].(string)) {
	case ChatListMainType:
		var chatListMain ChatListMain
		err := json.Unmarshal(*rawMsg, &chatListMain)
		return &chatListMain, err

	case ChatListArchiveType:
		var chatListArchive ChatListArchive
		err := json.Unmarshal(*rawMsg, &chatListArchive)
		return &chatListArchive, err

	case ChatListFilterType:
		var chatListFilter ChatListFilter
		err := json.Unmarshal(*rawMsg, &chatListFilter)
		return &chatListFilter, err

	default:
		return nil, fmt.Errorf("Error UnMarshaling, unknown type:" + objMap["@type"].(string))
	}
}

// ChatListMain A main list of chats
type ChatListMain struct {
	tdCommon
}

// MessageType return the string telegram-type of ChatListMain
func (chatListMain *ChatListMain) MessageType() string {
	return "chatListMain"
}

// NewChatListMain creates a new ChatListMain
//
func NewChatListMain() *ChatListMain {
	chatListMainTemp := ChatListMain{
		tdCommon: tdCommon{Type: "chatListMain"},
	}

	return &chatListMainTemp
}

// GetChatListEnum return the enum type of this object
func (chatListMain *ChatListMain) GetChatListEnum() ChatListEnum {
	return ChatListMainType
}

// ChatListArchive A list of chats usually located at the top of the main chat list. Unmuted chats are automatically moved from the Archive to the Main chat list when a new message arrives
type ChatListArchive struct {
	tdCommon
}

// MessageType return the string telegram-type of ChatListArchive
func (chatListArchive *ChatListArchive) MessageType() string {
	return "chatListArchive"
}

// NewChatListArchive creates a new ChatListArchive
//
func NewChatListArchive() *ChatListArchive {
	chatListArchiveTemp := ChatListArchive{
		tdCommon: tdCommon{Type: "chatListArchive"},
	}

	return &chatListArchiveTemp
}

// GetChatListEnum return the enum type of this object
func (chatListArchive *ChatListArchive) GetChatListEnum() ChatListEnum {
	return ChatListArchiveType
}

// ChatListFilter A list of chats belonging to a chat filter
type ChatListFilter struct {
	tdCommon
	ChatFilterId int32 `json:"chat_filter_id"` // Chat filter identifier
}

// MessageType return the string telegram-type of ChatListFilter
func (chatListFilter *ChatListFilter) MessageType() string {
	return "chatListFilter"
}

// NewChatListFilter creates a new ChatListFilter
//
// @param chatFilterId Chat filter identifier
func NewChatListFilter(chatFilterId int32) *ChatListFilter {
	chatListFilterTemp := ChatListFilter{
		tdCommon:     tdCommon{Type: "chatListFilter"},
		ChatFilterId: chatFilterId,
	}

	return &chatListFilterTemp
}

// GetChatListEnum return the enum type of this object
func (chatListFilter *ChatListFilter) GetChatListEnum() ChatListEnum {
	return ChatListFilterType
}
