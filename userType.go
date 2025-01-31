package tdlib

import (
	"encoding/json"
	"fmt"
)

// UserType Represents the type of a user. The following types are possible: regular users, deleted users and bots
type UserType interface {
	GetUserTypeEnum() UserTypeEnum
}

// UserTypeEnum Alias for abstract UserType 'Sub-Classes', used as constant-enum here
type UserTypeEnum string

// UserType enums
const (
	UserTypeRegularType UserTypeEnum = "userTypeRegular"
	UserTypeDeletedType UserTypeEnum = "userTypeDeleted"
	UserTypeBotType     UserTypeEnum = "userTypeBot"
	UserTypeUnknownType UserTypeEnum = "userTypeUnknown"
)

func unmarshalUserType(rawMsg *json.RawMessage) (UserType, error) {

	if rawMsg == nil {
		return nil, nil
	}
	var objMap map[string]interface{}
	err := json.Unmarshal(*rawMsg, &objMap)
	if err != nil {
		return nil, err
	}

	switch UserTypeEnum(objMap["@type"].(string)) {
	case UserTypeRegularType:
		var userTypeRegular UserTypeRegular
		err := json.Unmarshal(*rawMsg, &userTypeRegular)
		return &userTypeRegular, err

	case UserTypeDeletedType:
		var userTypeDeleted UserTypeDeleted
		err := json.Unmarshal(*rawMsg, &userTypeDeleted)
		return &userTypeDeleted, err

	case UserTypeBotType:
		var userTypeBot UserTypeBot
		err := json.Unmarshal(*rawMsg, &userTypeBot)
		return &userTypeBot, err

	case UserTypeUnknownType:
		var userTypeUnknown UserTypeUnknown
		err := json.Unmarshal(*rawMsg, &userTypeUnknown)
		return &userTypeUnknown, err

	default:
		return nil, fmt.Errorf("Error UnMarshaling, unknown type:" + objMap["@type"].(string))
	}
}

// UserTypeRegular A regular user
type UserTypeRegular struct {
	tdCommon
}

// MessageType return the string telegram-type of UserTypeRegular
func (userTypeRegular *UserTypeRegular) MessageType() string {
	return "userTypeRegular"
}

// NewUserTypeRegular creates a new UserTypeRegular
//
func NewUserTypeRegular() *UserTypeRegular {
	userTypeRegularTemp := UserTypeRegular{
		tdCommon: tdCommon{Type: "userTypeRegular"},
	}

	return &userTypeRegularTemp
}

// GetUserTypeEnum return the enum type of this object
func (userTypeRegular *UserTypeRegular) GetUserTypeEnum() UserTypeEnum {
	return UserTypeRegularType
}

// UserTypeDeleted A deleted user or deleted bot. No information on the user besides the user identifier is available. It is not possible to perform any active actions on this type of user
type UserTypeDeleted struct {
	tdCommon
}

// MessageType return the string telegram-type of UserTypeDeleted
func (userTypeDeleted *UserTypeDeleted) MessageType() string {
	return "userTypeDeleted"
}

// NewUserTypeDeleted creates a new UserTypeDeleted
//
func NewUserTypeDeleted() *UserTypeDeleted {
	userTypeDeletedTemp := UserTypeDeleted{
		tdCommon: tdCommon{Type: "userTypeDeleted"},
	}

	return &userTypeDeletedTemp
}

// GetUserTypeEnum return the enum type of this object
func (userTypeDeleted *UserTypeDeleted) GetUserTypeEnum() UserTypeEnum {
	return UserTypeDeletedType
}

// UserTypeBot A bot (see https://core.telegram.org/bots)
type UserTypeBot struct {
	tdCommon
	CanJoinGroups           bool   `json:"can_join_groups"`             // True, if the bot can be invited to basic group and supergroup chats
	CanReadAllGroupMessages bool   `json:"can_read_all_group_messages"` // True, if the bot can read all messages in basic group or supergroup chats and not just those addressed to the bot. In private and channel chats a bot can always read all messages
	IsInline                bool   `json:"is_inline"`                   // True, if the bot supports inline queries
	InlineQueryPlaceholder  string `json:"inline_query_placeholder"`    // Placeholder for inline queries (displayed on the application input field)
	NeedLocation            bool   `json:"need_location"`               // True, if the location of the user is expected to be sent with every inline query to this bot
}

// MessageType return the string telegram-type of UserTypeBot
func (userTypeBot *UserTypeBot) MessageType() string {
	return "userTypeBot"
}

// NewUserTypeBot creates a new UserTypeBot
//
// @param canJoinGroups True, if the bot can be invited to basic group and supergroup chats
// @param canReadAllGroupMessages True, if the bot can read all messages in basic group or supergroup chats and not just those addressed to the bot. In private and channel chats a bot can always read all messages
// @param isInline True, if the bot supports inline queries
// @param inlineQueryPlaceholder Placeholder for inline queries (displayed on the application input field)
// @param needLocation True, if the location of the user is expected to be sent with every inline query to this bot
func NewUserTypeBot(canJoinGroups bool, canReadAllGroupMessages bool, isInline bool, inlineQueryPlaceholder string, needLocation bool) *UserTypeBot {
	userTypeBotTemp := UserTypeBot{
		tdCommon:                tdCommon{Type: "userTypeBot"},
		CanJoinGroups:           canJoinGroups,
		CanReadAllGroupMessages: canReadAllGroupMessages,
		IsInline:                isInline,
		InlineQueryPlaceholder:  inlineQueryPlaceholder,
		NeedLocation:            needLocation,
	}

	return &userTypeBotTemp
}

// GetUserTypeEnum return the enum type of this object
func (userTypeBot *UserTypeBot) GetUserTypeEnum() UserTypeEnum {
	return UserTypeBotType
}

// UserTypeUnknown No information on the user besides the user identifier is available, yet this user has not been deleted. This object is extremely rare and must be handled like a deleted user. It is not possible to perform any actions on users of this type
type UserTypeUnknown struct {
	tdCommon
}

// MessageType return the string telegram-type of UserTypeUnknown
func (userTypeUnknown *UserTypeUnknown) MessageType() string {
	return "userTypeUnknown"
}

// NewUserTypeUnknown creates a new UserTypeUnknown
//
func NewUserTypeUnknown() *UserTypeUnknown {
	userTypeUnknownTemp := UserTypeUnknown{
		tdCommon: tdCommon{Type: "userTypeUnknown"},
	}

	return &userTypeUnknownTemp
}

// GetUserTypeEnum return the enum type of this object
func (userTypeUnknown *UserTypeUnknown) GetUserTypeEnum() UserTypeEnum {
	return UserTypeUnknownType
}
