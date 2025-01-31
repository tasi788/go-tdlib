package tdlib

import (
	"encoding/json"
	"fmt"
)

// CallDiscardReason Describes the reason why a call was discarded
type CallDiscardReason interface {
	GetCallDiscardReasonEnum() CallDiscardReasonEnum
}

// CallDiscardReasonEnum Alias for abstract CallDiscardReason 'Sub-Classes', used as constant-enum here
type CallDiscardReasonEnum string

// CallDiscardReason enums
const (
	CallDiscardReasonEmptyType        CallDiscardReasonEnum = "callDiscardReasonEmpty"
	CallDiscardReasonMissedType       CallDiscardReasonEnum = "callDiscardReasonMissed"
	CallDiscardReasonDeclinedType     CallDiscardReasonEnum = "callDiscardReasonDeclined"
	CallDiscardReasonDisconnectedType CallDiscardReasonEnum = "callDiscardReasonDisconnected"
	CallDiscardReasonHungUpType       CallDiscardReasonEnum = "callDiscardReasonHungUp"
)

func unmarshalCallDiscardReason(rawMsg *json.RawMessage) (CallDiscardReason, error) {

	if rawMsg == nil {
		return nil, nil
	}
	var objMap map[string]interface{}
	err := json.Unmarshal(*rawMsg, &objMap)
	if err != nil {
		return nil, err
	}

	switch CallDiscardReasonEnum(objMap["@type"].(string)) {
	case CallDiscardReasonEmptyType:
		var callDiscardReasonEmpty CallDiscardReasonEmpty
		err := json.Unmarshal(*rawMsg, &callDiscardReasonEmpty)
		return &callDiscardReasonEmpty, err

	case CallDiscardReasonMissedType:
		var callDiscardReasonMissed CallDiscardReasonMissed
		err := json.Unmarshal(*rawMsg, &callDiscardReasonMissed)
		return &callDiscardReasonMissed, err

	case CallDiscardReasonDeclinedType:
		var callDiscardReasonDeclined CallDiscardReasonDeclined
		err := json.Unmarshal(*rawMsg, &callDiscardReasonDeclined)
		return &callDiscardReasonDeclined, err

	case CallDiscardReasonDisconnectedType:
		var callDiscardReasonDisconnected CallDiscardReasonDisconnected
		err := json.Unmarshal(*rawMsg, &callDiscardReasonDisconnected)
		return &callDiscardReasonDisconnected, err

	case CallDiscardReasonHungUpType:
		var callDiscardReasonHungUp CallDiscardReasonHungUp
		err := json.Unmarshal(*rawMsg, &callDiscardReasonHungUp)
		return &callDiscardReasonHungUp, err

	default:
		return nil, fmt.Errorf("Error UnMarshaling, unknown type:" + objMap["@type"].(string))
	}
}

// CallDiscardReasonEmpty The call wasn't discarded, or the reason is unknown
type CallDiscardReasonEmpty struct {
	tdCommon
}

// MessageType return the string telegram-type of CallDiscardReasonEmpty
func (callDiscardReasonEmpty *CallDiscardReasonEmpty) MessageType() string {
	return "callDiscardReasonEmpty"
}

// NewCallDiscardReasonEmpty creates a new CallDiscardReasonEmpty
//
func NewCallDiscardReasonEmpty() *CallDiscardReasonEmpty {
	callDiscardReasonEmptyTemp := CallDiscardReasonEmpty{
		tdCommon: tdCommon{Type: "callDiscardReasonEmpty"},
	}

	return &callDiscardReasonEmptyTemp
}

// GetCallDiscardReasonEnum return the enum type of this object
func (callDiscardReasonEmpty *CallDiscardReasonEmpty) GetCallDiscardReasonEnum() CallDiscardReasonEnum {
	return CallDiscardReasonEmptyType
}

// CallDiscardReasonMissed The call was ended before the conversation started. It was canceled by the caller or missed by the other party
type CallDiscardReasonMissed struct {
	tdCommon
}

// MessageType return the string telegram-type of CallDiscardReasonMissed
func (callDiscardReasonMissed *CallDiscardReasonMissed) MessageType() string {
	return "callDiscardReasonMissed"
}

// NewCallDiscardReasonMissed creates a new CallDiscardReasonMissed
//
func NewCallDiscardReasonMissed() *CallDiscardReasonMissed {
	callDiscardReasonMissedTemp := CallDiscardReasonMissed{
		tdCommon: tdCommon{Type: "callDiscardReasonMissed"},
	}

	return &callDiscardReasonMissedTemp
}

// GetCallDiscardReasonEnum return the enum type of this object
func (callDiscardReasonMissed *CallDiscardReasonMissed) GetCallDiscardReasonEnum() CallDiscardReasonEnum {
	return CallDiscardReasonMissedType
}

// CallDiscardReasonDeclined The call was ended before the conversation started. It was declined by the other party
type CallDiscardReasonDeclined struct {
	tdCommon
}

// MessageType return the string telegram-type of CallDiscardReasonDeclined
func (callDiscardReasonDeclined *CallDiscardReasonDeclined) MessageType() string {
	return "callDiscardReasonDeclined"
}

// NewCallDiscardReasonDeclined creates a new CallDiscardReasonDeclined
//
func NewCallDiscardReasonDeclined() *CallDiscardReasonDeclined {
	callDiscardReasonDeclinedTemp := CallDiscardReasonDeclined{
		tdCommon: tdCommon{Type: "callDiscardReasonDeclined"},
	}

	return &callDiscardReasonDeclinedTemp
}

// GetCallDiscardReasonEnum return the enum type of this object
func (callDiscardReasonDeclined *CallDiscardReasonDeclined) GetCallDiscardReasonEnum() CallDiscardReasonEnum {
	return CallDiscardReasonDeclinedType
}

// CallDiscardReasonDisconnected The call was ended during the conversation because the users were disconnected
type CallDiscardReasonDisconnected struct {
	tdCommon
}

// MessageType return the string telegram-type of CallDiscardReasonDisconnected
func (callDiscardReasonDisconnected *CallDiscardReasonDisconnected) MessageType() string {
	return "callDiscardReasonDisconnected"
}

// NewCallDiscardReasonDisconnected creates a new CallDiscardReasonDisconnected
//
func NewCallDiscardReasonDisconnected() *CallDiscardReasonDisconnected {
	callDiscardReasonDisconnectedTemp := CallDiscardReasonDisconnected{
		tdCommon: tdCommon{Type: "callDiscardReasonDisconnected"},
	}

	return &callDiscardReasonDisconnectedTemp
}

// GetCallDiscardReasonEnum return the enum type of this object
func (callDiscardReasonDisconnected *CallDiscardReasonDisconnected) GetCallDiscardReasonEnum() CallDiscardReasonEnum {
	return CallDiscardReasonDisconnectedType
}

// CallDiscardReasonHungUp The call was ended because one of the parties hung up
type CallDiscardReasonHungUp struct {
	tdCommon
}

// MessageType return the string telegram-type of CallDiscardReasonHungUp
func (callDiscardReasonHungUp *CallDiscardReasonHungUp) MessageType() string {
	return "callDiscardReasonHungUp"
}

// NewCallDiscardReasonHungUp creates a new CallDiscardReasonHungUp
//
func NewCallDiscardReasonHungUp() *CallDiscardReasonHungUp {
	callDiscardReasonHungUpTemp := CallDiscardReasonHungUp{
		tdCommon: tdCommon{Type: "callDiscardReasonHungUp"},
	}

	return &callDiscardReasonHungUpTemp
}

// GetCallDiscardReasonEnum return the enum type of this object
func (callDiscardReasonHungUp *CallDiscardReasonHungUp) GetCallDiscardReasonEnum() CallDiscardReasonEnum {
	return CallDiscardReasonHungUpType
}
