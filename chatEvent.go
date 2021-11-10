package tdlib

import (
	"encoding/json"
)

// ChatEvent Represents a chat event
type ChatEvent struct {
	tdCommon
	Id     JSONInt64       `json:"id"`      // Chat event identifier
	Date   int32           `json:"date"`    // Point in time (Unix timestamp) when the event happened
	UserId int64           `json:"user_id"` // Identifier of the user who performed the action that triggered the event
	Action ChatEventAction `json:"action"`  // Action performed by the user
}

// MessageType return the string telegram-type of ChatEvent
func (chatEvent *ChatEvent) MessageType() string {
	return "chatEvent"
}

// NewChatEvent creates a new ChatEvent
//
// @param id Chat event identifier
// @param date Point in time (Unix timestamp) when the event happened
// @param userId Identifier of the user who performed the action that triggered the event
// @param action Action performed by the user
func NewChatEvent(id JSONInt64, date int32, userId int64, action ChatEventAction) *ChatEvent {
	chatEventTemp := ChatEvent{
		tdCommon: tdCommon{Type: "chatEvent"},
		Id:       id,
		Date:     date,
		UserId:   userId,
		Action:   action,
	}

	return &chatEventTemp
}

// UnmarshalJSON unmarshal to json
func (chatEvent *ChatEvent) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	tempObj := struct {
		tdCommon
		Id     JSONInt64 `json:"id"`      // Chat event identifier
		Date   int32     `json:"date"`    // Point in time (Unix timestamp) when the event happened
		UserId int64     `json:"user_id"` // Identifier of the user who performed the action that triggered the event

	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	chatEvent.tdCommon = tempObj.tdCommon
	chatEvent.Id = tempObj.Id
	chatEvent.Date = tempObj.Date
	chatEvent.UserId = tempObj.UserId

	fieldAction, _ := unmarshalChatEventAction(objMap["action"])
	chatEvent.Action = fieldAction

	return nil
}
