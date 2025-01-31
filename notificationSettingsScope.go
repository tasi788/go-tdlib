package tdlib

import (
	"encoding/json"
	"fmt"
)

// NotificationSettingsScope Describes the types of chats to which notification settings are relevant
type NotificationSettingsScope interface {
	GetNotificationSettingsScopeEnum() NotificationSettingsScopeEnum
}

// NotificationSettingsScopeEnum Alias for abstract NotificationSettingsScope 'Sub-Classes', used as constant-enum here
type NotificationSettingsScopeEnum string

// NotificationSettingsScope enums
const (
	NotificationSettingsScopePrivateChatsType NotificationSettingsScopeEnum = "notificationSettingsScopePrivateChats"
	NotificationSettingsScopeGroupChatsType   NotificationSettingsScopeEnum = "notificationSettingsScopeGroupChats"
	NotificationSettingsScopeChannelChatsType NotificationSettingsScopeEnum = "notificationSettingsScopeChannelChats"
)

func unmarshalNotificationSettingsScope(rawMsg *json.RawMessage) (NotificationSettingsScope, error) {

	if rawMsg == nil {
		return nil, nil
	}
	var objMap map[string]interface{}
	err := json.Unmarshal(*rawMsg, &objMap)
	if err != nil {
		return nil, err
	}

	switch NotificationSettingsScopeEnum(objMap["@type"].(string)) {
	case NotificationSettingsScopePrivateChatsType:
		var notificationSettingsScopePrivateChats NotificationSettingsScopePrivateChats
		err := json.Unmarshal(*rawMsg, &notificationSettingsScopePrivateChats)
		return &notificationSettingsScopePrivateChats, err

	case NotificationSettingsScopeGroupChatsType:
		var notificationSettingsScopeGroupChats NotificationSettingsScopeGroupChats
		err := json.Unmarshal(*rawMsg, &notificationSettingsScopeGroupChats)
		return &notificationSettingsScopeGroupChats, err

	case NotificationSettingsScopeChannelChatsType:
		var notificationSettingsScopeChannelChats NotificationSettingsScopeChannelChats
		err := json.Unmarshal(*rawMsg, &notificationSettingsScopeChannelChats)
		return &notificationSettingsScopeChannelChats, err

	default:
		return nil, fmt.Errorf("Error UnMarshaling, unknown type:" + objMap["@type"].(string))
	}
}

// NotificationSettingsScopePrivateChats Notification settings applied to all private and secret chats when the corresponding chat setting has a default value
type NotificationSettingsScopePrivateChats struct {
	tdCommon
}

// MessageType return the string telegram-type of NotificationSettingsScopePrivateChats
func (notificationSettingsScopePrivateChats *NotificationSettingsScopePrivateChats) MessageType() string {
	return "notificationSettingsScopePrivateChats"
}

// NewNotificationSettingsScopePrivateChats creates a new NotificationSettingsScopePrivateChats
//
func NewNotificationSettingsScopePrivateChats() *NotificationSettingsScopePrivateChats {
	notificationSettingsScopePrivateChatsTemp := NotificationSettingsScopePrivateChats{
		tdCommon: tdCommon{Type: "notificationSettingsScopePrivateChats"},
	}

	return &notificationSettingsScopePrivateChatsTemp
}

// GetNotificationSettingsScopeEnum return the enum type of this object
func (notificationSettingsScopePrivateChats *NotificationSettingsScopePrivateChats) GetNotificationSettingsScopeEnum() NotificationSettingsScopeEnum {
	return NotificationSettingsScopePrivateChatsType
}

// NotificationSettingsScopeGroupChats Notification settings applied to all basic groups and supergroups when the corresponding chat setting has a default value
type NotificationSettingsScopeGroupChats struct {
	tdCommon
}

// MessageType return the string telegram-type of NotificationSettingsScopeGroupChats
func (notificationSettingsScopeGroupChats *NotificationSettingsScopeGroupChats) MessageType() string {
	return "notificationSettingsScopeGroupChats"
}

// NewNotificationSettingsScopeGroupChats creates a new NotificationSettingsScopeGroupChats
//
func NewNotificationSettingsScopeGroupChats() *NotificationSettingsScopeGroupChats {
	notificationSettingsScopeGroupChatsTemp := NotificationSettingsScopeGroupChats{
		tdCommon: tdCommon{Type: "notificationSettingsScopeGroupChats"},
	}

	return &notificationSettingsScopeGroupChatsTemp
}

// GetNotificationSettingsScopeEnum return the enum type of this object
func (notificationSettingsScopeGroupChats *NotificationSettingsScopeGroupChats) GetNotificationSettingsScopeEnum() NotificationSettingsScopeEnum {
	return NotificationSettingsScopeGroupChatsType
}

// NotificationSettingsScopeChannelChats Notification settings applied to all channels when the corresponding chat setting has a default value
type NotificationSettingsScopeChannelChats struct {
	tdCommon
}

// MessageType return the string telegram-type of NotificationSettingsScopeChannelChats
func (notificationSettingsScopeChannelChats *NotificationSettingsScopeChannelChats) MessageType() string {
	return "notificationSettingsScopeChannelChats"
}

// NewNotificationSettingsScopeChannelChats creates a new NotificationSettingsScopeChannelChats
//
func NewNotificationSettingsScopeChannelChats() *NotificationSettingsScopeChannelChats {
	notificationSettingsScopeChannelChatsTemp := NotificationSettingsScopeChannelChats{
		tdCommon: tdCommon{Type: "notificationSettingsScopeChannelChats"},
	}

	return &notificationSettingsScopeChannelChatsTemp
}

// GetNotificationSettingsScopeEnum return the enum type of this object
func (notificationSettingsScopeChannelChats *NotificationSettingsScopeChannelChats) GetNotificationSettingsScopeEnum() NotificationSettingsScopeEnum {
	return NotificationSettingsScopeChannelChatsType
}
