package tdlib

import (
	"encoding/json"
)

// NotificationGroup Describes a group of notifications
type NotificationGroup struct {
	tdCommon
	Id            int32                 `json:"id"`            // Unique persistent auto-incremented from 1 identifier of the notification group
	Type          NotificationGroupType `json:"type"`          // Type of the group
	ChatId        int64                 `json:"chat_id"`       // Identifier of a chat to which all notifications in the group belong
	TotalCount    int32                 `json:"total_count"`   // Total number of active notifications in the group
	Notifications []Notification        `json:"notifications"` // The list of active notifications
}

// MessageType return the string telegram-type of NotificationGroup
func (notificationGroup *NotificationGroup) MessageType() string {
	return "notificationGroup"
}

// NewNotificationGroup creates a new NotificationGroup
//
// @param id Unique persistent auto-incremented from 1 identifier of the notification group
// @param typeParam Type of the group
// @param chatId Identifier of a chat to which all notifications in the group belong
// @param totalCount Total number of active notifications in the group
// @param notifications The list of active notifications
func NewNotificationGroup(id int32, typeParam NotificationGroupType, chatId int64, totalCount int32, notifications []Notification) *NotificationGroup {
	notificationGroupTemp := NotificationGroup{
		tdCommon:      tdCommon{Type: "notificationGroup"},
		Id:            id,
		Type:          typeParam,
		ChatId:        chatId,
		TotalCount:    totalCount,
		Notifications: notifications,
	}

	return &notificationGroupTemp
}

// UnmarshalJSON unmarshal to json
func (notificationGroup *NotificationGroup) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	tempObj := struct {
		tdCommon
		Id            int32          `json:"id"`            // Unique persistent auto-incremented from 1 identifier of the notification group
		ChatId        int64          `json:"chat_id"`       // Identifier of a chat to which all notifications in the group belong
		TotalCount    int32          `json:"total_count"`   // Total number of active notifications in the group
		Notifications []Notification `json:"notifications"` // The list of active notifications
	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	notificationGroup.tdCommon = tempObj.tdCommon
	notificationGroup.Id = tempObj.Id
	notificationGroup.ChatId = tempObj.ChatId
	notificationGroup.TotalCount = tempObj.TotalCount
	notificationGroup.Notifications = tempObj.Notifications

	fieldType, _ := unmarshalNotificationGroupType(objMap["type"])
	notificationGroup.Type = fieldType

	return nil
}
