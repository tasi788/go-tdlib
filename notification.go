package tdlib

import (
	"encoding/json"
)

// Notification Contains information about a notification
type Notification struct {
	tdCommon
	Id       int32            `json:"id"`        // Unique persistent identifier of this notification
	Date     int32            `json:"date"`      // Notification date
	IsSilent bool             `json:"is_silent"` // True, if the notification was initially silent
	Type     NotificationType `json:"type"`      // Notification type
}

// MessageType return the string telegram-type of Notification
func (notification *Notification) MessageType() string {
	return "notification"
}

// NewNotification creates a new Notification
//
// @param id Unique persistent identifier of this notification
// @param date Notification date
// @param isSilent True, if the notification was initially silent
// @param typeParam Notification type
func NewNotification(id int32, date int32, isSilent bool, typeParam NotificationType) *Notification {
	notificationTemp := Notification{
		tdCommon: tdCommon{Type: "notification"},
		Id:       id,
		Date:     date,
		IsSilent: isSilent,
		Type:     typeParam,
	}

	return &notificationTemp
}

// UnmarshalJSON unmarshal to json
func (notification *Notification) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	tempObj := struct {
		tdCommon
		Id       int32 `json:"id"`        // Unique persistent identifier of this notification
		Date     int32 `json:"date"`      // Notification date
		IsSilent bool  `json:"is_silent"` // True, if the notification was initially silent

	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	notification.tdCommon = tempObj.tdCommon
	notification.Id = tempObj.Id
	notification.Date = tempObj.Date
	notification.IsSilent = tempObj.IsSilent

	fieldType, _ := unmarshalNotificationType(objMap["type"])
	notification.Type = fieldType

	return nil
}
