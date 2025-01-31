package tdlib

import (
	"encoding/json"
	"fmt"
)

// NotificationGroupType Describes the type of notifications in a notification group
type NotificationGroupType interface {
	GetNotificationGroupTypeEnum() NotificationGroupTypeEnum
}

// NotificationGroupTypeEnum Alias for abstract NotificationGroupType 'Sub-Classes', used as constant-enum here
type NotificationGroupTypeEnum string

// NotificationGroupType enums
const (
	NotificationGroupTypeMessagesType   NotificationGroupTypeEnum = "notificationGroupTypeMessages"
	NotificationGroupTypeMentionsType   NotificationGroupTypeEnum = "notificationGroupTypeMentions"
	NotificationGroupTypeSecretChatType NotificationGroupTypeEnum = "notificationGroupTypeSecretChat"
	NotificationGroupTypeCallsType      NotificationGroupTypeEnum = "notificationGroupTypeCalls"
)

func unmarshalNotificationGroupType(rawMsg *json.RawMessage) (NotificationGroupType, error) {

	if rawMsg == nil {
		return nil, nil
	}
	var objMap map[string]interface{}
	err := json.Unmarshal(*rawMsg, &objMap)
	if err != nil {
		return nil, err
	}

	switch NotificationGroupTypeEnum(objMap["@type"].(string)) {
	case NotificationGroupTypeMessagesType:
		var notificationGroupTypeMessages NotificationGroupTypeMessages
		err := json.Unmarshal(*rawMsg, &notificationGroupTypeMessages)
		return &notificationGroupTypeMessages, err

	case NotificationGroupTypeMentionsType:
		var notificationGroupTypeMentions NotificationGroupTypeMentions
		err := json.Unmarshal(*rawMsg, &notificationGroupTypeMentions)
		return &notificationGroupTypeMentions, err

	case NotificationGroupTypeSecretChatType:
		var notificationGroupTypeSecretChat NotificationGroupTypeSecretChat
		err := json.Unmarshal(*rawMsg, &notificationGroupTypeSecretChat)
		return &notificationGroupTypeSecretChat, err

	case NotificationGroupTypeCallsType:
		var notificationGroupTypeCalls NotificationGroupTypeCalls
		err := json.Unmarshal(*rawMsg, &notificationGroupTypeCalls)
		return &notificationGroupTypeCalls, err

	default:
		return nil, fmt.Errorf("Error UnMarshaling, unknown type:" + objMap["@type"].(string))
	}
}

// NotificationGroupTypeMessages A group containing notifications of type notificationTypeNewMessage and notificationTypeNewPushMessage with ordinary unread messages
type NotificationGroupTypeMessages struct {
	tdCommon
}

// MessageType return the string telegram-type of NotificationGroupTypeMessages
func (notificationGroupTypeMessages *NotificationGroupTypeMessages) MessageType() string {
	return "notificationGroupTypeMessages"
}

// NewNotificationGroupTypeMessages creates a new NotificationGroupTypeMessages
//
func NewNotificationGroupTypeMessages() *NotificationGroupTypeMessages {
	notificationGroupTypeMessagesTemp := NotificationGroupTypeMessages{
		tdCommon: tdCommon{Type: "notificationGroupTypeMessages"},
	}

	return &notificationGroupTypeMessagesTemp
}

// GetNotificationGroupTypeEnum return the enum type of this object
func (notificationGroupTypeMessages *NotificationGroupTypeMessages) GetNotificationGroupTypeEnum() NotificationGroupTypeEnum {
	return NotificationGroupTypeMessagesType
}

// NotificationGroupTypeMentions A group containing notifications of type notificationTypeNewMessage and notificationTypeNewPushMessage with unread mentions of the current user, replies to their messages, or a pinned message
type NotificationGroupTypeMentions struct {
	tdCommon
}

// MessageType return the string telegram-type of NotificationGroupTypeMentions
func (notificationGroupTypeMentions *NotificationGroupTypeMentions) MessageType() string {
	return "notificationGroupTypeMentions"
}

// NewNotificationGroupTypeMentions creates a new NotificationGroupTypeMentions
//
func NewNotificationGroupTypeMentions() *NotificationGroupTypeMentions {
	notificationGroupTypeMentionsTemp := NotificationGroupTypeMentions{
		tdCommon: tdCommon{Type: "notificationGroupTypeMentions"},
	}

	return &notificationGroupTypeMentionsTemp
}

// GetNotificationGroupTypeEnum return the enum type of this object
func (notificationGroupTypeMentions *NotificationGroupTypeMentions) GetNotificationGroupTypeEnum() NotificationGroupTypeEnum {
	return NotificationGroupTypeMentionsType
}

// NotificationGroupTypeSecretChat A group containing a notification of type notificationTypeNewSecretChat
type NotificationGroupTypeSecretChat struct {
	tdCommon
}

// MessageType return the string telegram-type of NotificationGroupTypeSecretChat
func (notificationGroupTypeSecretChat *NotificationGroupTypeSecretChat) MessageType() string {
	return "notificationGroupTypeSecretChat"
}

// NewNotificationGroupTypeSecretChat creates a new NotificationGroupTypeSecretChat
//
func NewNotificationGroupTypeSecretChat() *NotificationGroupTypeSecretChat {
	notificationGroupTypeSecretChatTemp := NotificationGroupTypeSecretChat{
		tdCommon: tdCommon{Type: "notificationGroupTypeSecretChat"},
	}

	return &notificationGroupTypeSecretChatTemp
}

// GetNotificationGroupTypeEnum return the enum type of this object
func (notificationGroupTypeSecretChat *NotificationGroupTypeSecretChat) GetNotificationGroupTypeEnum() NotificationGroupTypeEnum {
	return NotificationGroupTypeSecretChatType
}

// NotificationGroupTypeCalls A group containing notifications of type notificationTypeNewCall
type NotificationGroupTypeCalls struct {
	tdCommon
}

// MessageType return the string telegram-type of NotificationGroupTypeCalls
func (notificationGroupTypeCalls *NotificationGroupTypeCalls) MessageType() string {
	return "notificationGroupTypeCalls"
}

// NewNotificationGroupTypeCalls creates a new NotificationGroupTypeCalls
//
func NewNotificationGroupTypeCalls() *NotificationGroupTypeCalls {
	notificationGroupTypeCallsTemp := NotificationGroupTypeCalls{
		tdCommon: tdCommon{Type: "notificationGroupTypeCalls"},
	}

	return &notificationGroupTypeCallsTemp
}

// GetNotificationGroupTypeEnum return the enum type of this object
func (notificationGroupTypeCalls *NotificationGroupTypeCalls) GetNotificationGroupTypeEnum() NotificationGroupTypeEnum {
	return NotificationGroupTypeCallsType
}
