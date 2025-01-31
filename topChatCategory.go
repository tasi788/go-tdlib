package tdlib

import (
	"encoding/json"
	"fmt"
)

// TopChatCategory Represents the categories of chats for which a list of frequently used chats can be retrieved
type TopChatCategory interface {
	GetTopChatCategoryEnum() TopChatCategoryEnum
}

// TopChatCategoryEnum Alias for abstract TopChatCategory 'Sub-Classes', used as constant-enum here
type TopChatCategoryEnum string

// TopChatCategory enums
const (
	TopChatCategoryUsersType        TopChatCategoryEnum = "topChatCategoryUsers"
	TopChatCategoryBotsType         TopChatCategoryEnum = "topChatCategoryBots"
	TopChatCategoryGroupsType       TopChatCategoryEnum = "topChatCategoryGroups"
	TopChatCategoryChannelsType     TopChatCategoryEnum = "topChatCategoryChannels"
	TopChatCategoryInlineBotsType   TopChatCategoryEnum = "topChatCategoryInlineBots"
	TopChatCategoryCallsType        TopChatCategoryEnum = "topChatCategoryCalls"
	TopChatCategoryForwardChatsType TopChatCategoryEnum = "topChatCategoryForwardChats"
)

func unmarshalTopChatCategory(rawMsg *json.RawMessage) (TopChatCategory, error) {

	if rawMsg == nil {
		return nil, nil
	}
	var objMap map[string]interface{}
	err := json.Unmarshal(*rawMsg, &objMap)
	if err != nil {
		return nil, err
	}

	switch TopChatCategoryEnum(objMap["@type"].(string)) {
	case TopChatCategoryUsersType:
		var topChatCategoryUsers TopChatCategoryUsers
		err := json.Unmarshal(*rawMsg, &topChatCategoryUsers)
		return &topChatCategoryUsers, err

	case TopChatCategoryBotsType:
		var topChatCategoryBots TopChatCategoryBots
		err := json.Unmarshal(*rawMsg, &topChatCategoryBots)
		return &topChatCategoryBots, err

	case TopChatCategoryGroupsType:
		var topChatCategoryGroups TopChatCategoryGroups
		err := json.Unmarshal(*rawMsg, &topChatCategoryGroups)
		return &topChatCategoryGroups, err

	case TopChatCategoryChannelsType:
		var topChatCategoryChannels TopChatCategoryChannels
		err := json.Unmarshal(*rawMsg, &topChatCategoryChannels)
		return &topChatCategoryChannels, err

	case TopChatCategoryInlineBotsType:
		var topChatCategoryInlineBots TopChatCategoryInlineBots
		err := json.Unmarshal(*rawMsg, &topChatCategoryInlineBots)
		return &topChatCategoryInlineBots, err

	case TopChatCategoryCallsType:
		var topChatCategoryCalls TopChatCategoryCalls
		err := json.Unmarshal(*rawMsg, &topChatCategoryCalls)
		return &topChatCategoryCalls, err

	case TopChatCategoryForwardChatsType:
		var topChatCategoryForwardChats TopChatCategoryForwardChats
		err := json.Unmarshal(*rawMsg, &topChatCategoryForwardChats)
		return &topChatCategoryForwardChats, err

	default:
		return nil, fmt.Errorf("Error UnMarshaling, unknown type:" + objMap["@type"].(string))
	}
}

// TopChatCategoryUsers A category containing frequently used private chats with non-bot users
type TopChatCategoryUsers struct {
	tdCommon
}

// MessageType return the string telegram-type of TopChatCategoryUsers
func (topChatCategoryUsers *TopChatCategoryUsers) MessageType() string {
	return "topChatCategoryUsers"
}

// NewTopChatCategoryUsers creates a new TopChatCategoryUsers
//
func NewTopChatCategoryUsers() *TopChatCategoryUsers {
	topChatCategoryUsersTemp := TopChatCategoryUsers{
		tdCommon: tdCommon{Type: "topChatCategoryUsers"},
	}

	return &topChatCategoryUsersTemp
}

// GetTopChatCategoryEnum return the enum type of this object
func (topChatCategoryUsers *TopChatCategoryUsers) GetTopChatCategoryEnum() TopChatCategoryEnum {
	return TopChatCategoryUsersType
}

// TopChatCategoryBots A category containing frequently used private chats with bot users
type TopChatCategoryBots struct {
	tdCommon
}

// MessageType return the string telegram-type of TopChatCategoryBots
func (topChatCategoryBots *TopChatCategoryBots) MessageType() string {
	return "topChatCategoryBots"
}

// NewTopChatCategoryBots creates a new TopChatCategoryBots
//
func NewTopChatCategoryBots() *TopChatCategoryBots {
	topChatCategoryBotsTemp := TopChatCategoryBots{
		tdCommon: tdCommon{Type: "topChatCategoryBots"},
	}

	return &topChatCategoryBotsTemp
}

// GetTopChatCategoryEnum return the enum type of this object
func (topChatCategoryBots *TopChatCategoryBots) GetTopChatCategoryEnum() TopChatCategoryEnum {
	return TopChatCategoryBotsType
}

// TopChatCategoryGroups A category containing frequently used basic groups and supergroups
type TopChatCategoryGroups struct {
	tdCommon
}

// MessageType return the string telegram-type of TopChatCategoryGroups
func (topChatCategoryGroups *TopChatCategoryGroups) MessageType() string {
	return "topChatCategoryGroups"
}

// NewTopChatCategoryGroups creates a new TopChatCategoryGroups
//
func NewTopChatCategoryGroups() *TopChatCategoryGroups {
	topChatCategoryGroupsTemp := TopChatCategoryGroups{
		tdCommon: tdCommon{Type: "topChatCategoryGroups"},
	}

	return &topChatCategoryGroupsTemp
}

// GetTopChatCategoryEnum return the enum type of this object
func (topChatCategoryGroups *TopChatCategoryGroups) GetTopChatCategoryEnum() TopChatCategoryEnum {
	return TopChatCategoryGroupsType
}

// TopChatCategoryChannels A category containing frequently used channels
type TopChatCategoryChannels struct {
	tdCommon
}

// MessageType return the string telegram-type of TopChatCategoryChannels
func (topChatCategoryChannels *TopChatCategoryChannels) MessageType() string {
	return "topChatCategoryChannels"
}

// NewTopChatCategoryChannels creates a new TopChatCategoryChannels
//
func NewTopChatCategoryChannels() *TopChatCategoryChannels {
	topChatCategoryChannelsTemp := TopChatCategoryChannels{
		tdCommon: tdCommon{Type: "topChatCategoryChannels"},
	}

	return &topChatCategoryChannelsTemp
}

// GetTopChatCategoryEnum return the enum type of this object
func (topChatCategoryChannels *TopChatCategoryChannels) GetTopChatCategoryEnum() TopChatCategoryEnum {
	return TopChatCategoryChannelsType
}

// TopChatCategoryInlineBots A category containing frequently used chats with inline bots sorted by their usage in inline mode
type TopChatCategoryInlineBots struct {
	tdCommon
}

// MessageType return the string telegram-type of TopChatCategoryInlineBots
func (topChatCategoryInlineBots *TopChatCategoryInlineBots) MessageType() string {
	return "topChatCategoryInlineBots"
}

// NewTopChatCategoryInlineBots creates a new TopChatCategoryInlineBots
//
func NewTopChatCategoryInlineBots() *TopChatCategoryInlineBots {
	topChatCategoryInlineBotsTemp := TopChatCategoryInlineBots{
		tdCommon: tdCommon{Type: "topChatCategoryInlineBots"},
	}

	return &topChatCategoryInlineBotsTemp
}

// GetTopChatCategoryEnum return the enum type of this object
func (topChatCategoryInlineBots *TopChatCategoryInlineBots) GetTopChatCategoryEnum() TopChatCategoryEnum {
	return TopChatCategoryInlineBotsType
}

// TopChatCategoryCalls A category containing frequently used chats used for calls
type TopChatCategoryCalls struct {
	tdCommon
}

// MessageType return the string telegram-type of TopChatCategoryCalls
func (topChatCategoryCalls *TopChatCategoryCalls) MessageType() string {
	return "topChatCategoryCalls"
}

// NewTopChatCategoryCalls creates a new TopChatCategoryCalls
//
func NewTopChatCategoryCalls() *TopChatCategoryCalls {
	topChatCategoryCallsTemp := TopChatCategoryCalls{
		tdCommon: tdCommon{Type: "topChatCategoryCalls"},
	}

	return &topChatCategoryCallsTemp
}

// GetTopChatCategoryEnum return the enum type of this object
func (topChatCategoryCalls *TopChatCategoryCalls) GetTopChatCategoryEnum() TopChatCategoryEnum {
	return TopChatCategoryCallsType
}

// TopChatCategoryForwardChats A category containing frequently used chats used to forward messages
type TopChatCategoryForwardChats struct {
	tdCommon
}

// MessageType return the string telegram-type of TopChatCategoryForwardChats
func (topChatCategoryForwardChats *TopChatCategoryForwardChats) MessageType() string {
	return "topChatCategoryForwardChats"
}

// NewTopChatCategoryForwardChats creates a new TopChatCategoryForwardChats
//
func NewTopChatCategoryForwardChats() *TopChatCategoryForwardChats {
	topChatCategoryForwardChatsTemp := TopChatCategoryForwardChats{
		tdCommon: tdCommon{Type: "topChatCategoryForwardChats"},
	}

	return &topChatCategoryForwardChatsTemp
}

// GetTopChatCategoryEnum return the enum type of this object
func (topChatCategoryForwardChats *TopChatCategoryForwardChats) GetTopChatCategoryEnum() TopChatCategoryEnum {
	return TopChatCategoryForwardChatsType
}
