package tdlib

import (
	"encoding/json"
)

// ChatEvent Represents a chat event
type ChatEvent struct {
	tdCommon
	Id       JSONInt64       `json:"id"`        // Chat event identifier
	Date     int32           `json:"date"`      // Point in time (Unix timestamp) when the event happened
	MemberId MessageSender   `json:"member_id"` // Identifier of the user or chat who performed the action
	Action   ChatEventAction `json:"action"`    // The action
}

// MessageType return the string telegram-type of ChatEvent
func (chatEvent *ChatEvent) MessageType() string {
	return "chatEvent"
}

// NewChatEvent creates a new ChatEvent
//
// @param id Chat event identifier
// @param date Point in time (Unix timestamp) when the event happened
// @param memberId Identifier of the user or chat who performed the action
// @param action The action
func NewChatEvent(id JSONInt64, date int32, memberId MessageSender, action ChatEventAction) *ChatEvent {
	chatEventTemp := ChatEvent{
		tdCommon: tdCommon{Type: "chatEvent"},
		Id:       id,
		Date:     date,
		MemberId: memberId,
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
		Id   JSONInt64 `json:"id"`   // Chat event identifier
		Date int32     `json:"date"` // Point in time (Unix timestamp) when the event happened

	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	chatEvent.tdCommon = tempObj.tdCommon
	chatEvent.Id = tempObj.Id
	chatEvent.Date = tempObj.Date

	fieldMemberId, _ := unmarshalMessageSender(objMap["member_id"])
	chatEvent.MemberId = fieldMemberId

	fieldAction, _ := unmarshalChatEventAction(objMap["action"])
	chatEvent.Action = fieldAction

	return nil
}
