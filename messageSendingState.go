package tdlib

import (
	"encoding/json"
	"fmt"
)

// MessageSendingState Contains information about the sending state of the message
type MessageSendingState interface {
	GetMessageSendingStateEnum() MessageSendingStateEnum
}

// MessageSendingStateEnum Alias for abstract MessageSendingState 'Sub-Classes', used as constant-enum here
type MessageSendingStateEnum string

// MessageSendingState enums
const (
	MessageSendingStatePendingType MessageSendingStateEnum = "messageSendingStatePending"
	MessageSendingStateFailedType  MessageSendingStateEnum = "messageSendingStateFailed"
)

func unmarshalMessageSendingState(rawMsg *json.RawMessage) (MessageSendingState, error) {

	if rawMsg == nil {
		return nil, nil
	}
	var objMap map[string]interface{}
	err := json.Unmarshal(*rawMsg, &objMap)
	if err != nil {
		return nil, err
	}

	switch MessageSendingStateEnum(objMap["@type"].(string)) {
	case MessageSendingStatePendingType:
		var messageSendingStatePending MessageSendingStatePending
		err := json.Unmarshal(*rawMsg, &messageSendingStatePending)
		return &messageSendingStatePending, err

	case MessageSendingStateFailedType:
		var messageSendingStateFailed MessageSendingStateFailed
		err := json.Unmarshal(*rawMsg, &messageSendingStateFailed)
		return &messageSendingStateFailed, err

	default:
		return nil, fmt.Errorf("Error UnMarshaling, unknown type:" + objMap["@type"].(string))
	}
}

// MessageSendingStatePending The message is being sent now, but has not yet been delivered to the server
type MessageSendingStatePending struct {
	tdCommon
}

// MessageType return the string telegram-type of MessageSendingStatePending
func (messageSendingStatePending *MessageSendingStatePending) MessageType() string {
	return "messageSendingStatePending"
}

// NewMessageSendingStatePending creates a new MessageSendingStatePending
//
func NewMessageSendingStatePending() *MessageSendingStatePending {
	messageSendingStatePendingTemp := MessageSendingStatePending{
		tdCommon: tdCommon{Type: "messageSendingStatePending"},
	}

	return &messageSendingStatePendingTemp
}

// GetMessageSendingStateEnum return the enum type of this object
func (messageSendingStatePending *MessageSendingStatePending) GetMessageSendingStateEnum() MessageSendingStateEnum {
	return MessageSendingStatePendingType
}

// MessageSendingStateFailed The message failed to be sent
type MessageSendingStateFailed struct {
	tdCommon
	ErrorCode         int32   `json:"error_code"`          // An error code; 0 if unknown
	ErrorMessage      string  `json:"error_message"`       // Error message
	CanRetry          bool    `json:"can_retry"`           // True, if the message can be re-sent
	NeedAnotherSender bool    `json:"need_another_sender"` // True, if the message can be re-sent only on behalf of a different sender
	RetryAfter        float64 `json:"retry_after"`         // Time left before the message can be re-sent, in seconds. No update is sent when this field changes
}

// MessageType return the string telegram-type of MessageSendingStateFailed
func (messageSendingStateFailed *MessageSendingStateFailed) MessageType() string {
	return "messageSendingStateFailed"
}

// NewMessageSendingStateFailed creates a new MessageSendingStateFailed
//
// @param errorCode An error code; 0 if unknown
// @param errorMessage Error message
// @param canRetry True, if the message can be re-sent
// @param needAnotherSender True, if the message can be re-sent only on behalf of a different sender
// @param retryAfter Time left before the message can be re-sent, in seconds. No update is sent when this field changes
func NewMessageSendingStateFailed(errorCode int32, errorMessage string, canRetry bool, needAnotherSender bool, retryAfter float64) *MessageSendingStateFailed {
	messageSendingStateFailedTemp := MessageSendingStateFailed{
		tdCommon:          tdCommon{Type: "messageSendingStateFailed"},
		ErrorCode:         errorCode,
		ErrorMessage:      errorMessage,
		CanRetry:          canRetry,
		NeedAnotherSender: needAnotherSender,
		RetryAfter:        retryAfter,
	}

	return &messageSendingStateFailedTemp
}

// GetMessageSendingStateEnum return the enum type of this object
func (messageSendingStateFailed *MessageSendingStateFailed) GetMessageSendingStateEnum() MessageSendingStateEnum {
	return MessageSendingStateFailedType
}
