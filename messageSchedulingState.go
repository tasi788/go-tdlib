package tdlib

import (
	"encoding/json"
	"fmt"
)

// MessageSchedulingState Contains information about the time when a scheduled message will be sent
type MessageSchedulingState interface {
	GetMessageSchedulingStateEnum() MessageSchedulingStateEnum
}

// MessageSchedulingStateEnum Alias for abstract MessageSchedulingState 'Sub-Classes', used as constant-enum here
type MessageSchedulingStateEnum string

// MessageSchedulingState enums
const (
	MessageSchedulingStateSendAtDateType     MessageSchedulingStateEnum = "messageSchedulingStateSendAtDate"
	MessageSchedulingStateSendWhenOnlineType MessageSchedulingStateEnum = "messageSchedulingStateSendWhenOnline"
)

func unmarshalMessageSchedulingState(rawMsg *json.RawMessage) (MessageSchedulingState, error) {

	if rawMsg == nil {
		return nil, nil
	}
	var objMap map[string]interface{}
	err := json.Unmarshal(*rawMsg, &objMap)
	if err != nil {
		return nil, err
	}

	switch MessageSchedulingStateEnum(objMap["@type"].(string)) {
	case MessageSchedulingStateSendAtDateType:
		var messageSchedulingStateSendAtDate MessageSchedulingStateSendAtDate
		err := json.Unmarshal(*rawMsg, &messageSchedulingStateSendAtDate)
		return &messageSchedulingStateSendAtDate, err

	case MessageSchedulingStateSendWhenOnlineType:
		var messageSchedulingStateSendWhenOnline MessageSchedulingStateSendWhenOnline
		err := json.Unmarshal(*rawMsg, &messageSchedulingStateSendWhenOnline)
		return &messageSchedulingStateSendWhenOnline, err

	default:
		return nil, fmt.Errorf("Error UnMarshaling, unknown type:" + objMap["@type"].(string))
	}
}

// MessageSchedulingStateSendAtDate The message will be sent at the specified date
type MessageSchedulingStateSendAtDate struct {
	tdCommon
	SendDate int32 `json:"send_date"` // Date the message will be sent. The date must be within 367 days in the future
}

// MessageType return the string telegram-type of MessageSchedulingStateSendAtDate
func (messageSchedulingStateSendAtDate *MessageSchedulingStateSendAtDate) MessageType() string {
	return "messageSchedulingStateSendAtDate"
}

// NewMessageSchedulingStateSendAtDate creates a new MessageSchedulingStateSendAtDate
//
// @param sendDate Date the message will be sent. The date must be within 367 days in the future
func NewMessageSchedulingStateSendAtDate(sendDate int32) *MessageSchedulingStateSendAtDate {
	messageSchedulingStateSendAtDateTemp := MessageSchedulingStateSendAtDate{
		tdCommon: tdCommon{Type: "messageSchedulingStateSendAtDate"},
		SendDate: sendDate,
	}

	return &messageSchedulingStateSendAtDateTemp
}

// GetMessageSchedulingStateEnum return the enum type of this object
func (messageSchedulingStateSendAtDate *MessageSchedulingStateSendAtDate) GetMessageSchedulingStateEnum() MessageSchedulingStateEnum {
	return MessageSchedulingStateSendAtDateType
}

// MessageSchedulingStateSendWhenOnline The message will be sent when the peer will be online. Applicable to private chats only and when the exact online status of the peer is known
type MessageSchedulingStateSendWhenOnline struct {
	tdCommon
}

// MessageType return the string telegram-type of MessageSchedulingStateSendWhenOnline
func (messageSchedulingStateSendWhenOnline *MessageSchedulingStateSendWhenOnline) MessageType() string {
	return "messageSchedulingStateSendWhenOnline"
}

// NewMessageSchedulingStateSendWhenOnline creates a new MessageSchedulingStateSendWhenOnline
//
func NewMessageSchedulingStateSendWhenOnline() *MessageSchedulingStateSendWhenOnline {
	messageSchedulingStateSendWhenOnlineTemp := MessageSchedulingStateSendWhenOnline{
		tdCommon: tdCommon{Type: "messageSchedulingStateSendWhenOnline"},
	}

	return &messageSchedulingStateSendWhenOnlineTemp
}

// GetMessageSchedulingStateEnum return the enum type of this object
func (messageSchedulingStateSendWhenOnline *MessageSchedulingStateSendWhenOnline) GetMessageSchedulingStateEnum() MessageSchedulingStateEnum {
	return MessageSchedulingStateSendWhenOnlineType
}
