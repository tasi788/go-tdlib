package tdlib

import (
	"encoding/json"
)

// MessageSendOptions Options to be used when a message is sent
type MessageSendOptions struct {
	tdCommon
	DisableNotification bool                   `json:"disable_notification"` // Pass true to disable notification for the message
	FromBackground      bool                   `json:"from_background"`      // Pass true if the message is sent from the background
	ProtectContent      bool                   `json:"protect_content"`      // Pass true if the content of the message must be protected from forwarding and saving; for bots only
	SchedulingState     MessageSchedulingState `json:"scheduling_state"`     // Message scheduling state; pass null to send message immediately. Messages sent to a secret chat, live location messages and self-destructing messages can't be scheduled
}

// MessageType return the string telegram-type of MessageSendOptions
func (messageSendOptions *MessageSendOptions) MessageType() string {
	return "messageSendOptions"
}

// NewMessageSendOptions creates a new MessageSendOptions
//
// @param disableNotification Pass true to disable notification for the message
// @param fromBackground Pass true if the message is sent from the background
// @param protectContent Pass true if the content of the message must be protected from forwarding and saving; for bots only
// @param schedulingState Message scheduling state; pass null to send message immediately. Messages sent to a secret chat, live location messages and self-destructing messages can't be scheduled
func NewMessageSendOptions(disableNotification bool, fromBackground bool, protectContent bool, schedulingState MessageSchedulingState) *MessageSendOptions {
	messageSendOptionsTemp := MessageSendOptions{
		tdCommon:            tdCommon{Type: "messageSendOptions"},
		DisableNotification: disableNotification,
		FromBackground:      fromBackground,
		ProtectContent:      protectContent,
		SchedulingState:     schedulingState,
	}

	return &messageSendOptionsTemp
}

// UnmarshalJSON unmarshal to json
func (messageSendOptions *MessageSendOptions) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	tempObj := struct {
		tdCommon
		DisableNotification bool `json:"disable_notification"` // Pass true to disable notification for the message
		FromBackground      bool `json:"from_background"`      // Pass true if the message is sent from the background
		ProtectContent      bool `json:"protect_content"`      // Pass true if the content of the message must be protected from forwarding and saving; for bots only

	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	messageSendOptions.tdCommon = tempObj.tdCommon
	messageSendOptions.DisableNotification = tempObj.DisableNotification
	messageSendOptions.FromBackground = tempObj.FromBackground
	messageSendOptions.ProtectContent = tempObj.ProtectContent

	fieldSchedulingState, _ := unmarshalMessageSchedulingState(objMap["scheduling_state"])
	messageSendOptions.SchedulingState = fieldSchedulingState

	return nil
}
