package tdlib

import (
	"encoding/json"
	"fmt"
)

// CallState Describes the current call state
type CallState interface {
	GetCallStateEnum() CallStateEnum
}

// CallStateEnum Alias for abstract CallState 'Sub-Classes', used as constant-enum here
type CallStateEnum string

// CallState enums
const (
	CallStatePendingType        CallStateEnum = "callStatePending"
	CallStateExchangingKeysType CallStateEnum = "callStateExchangingKeys"
	CallStateReadyType          CallStateEnum = "callStateReady"
	CallStateHangingUpType      CallStateEnum = "callStateHangingUp"
	CallStateDiscardedType      CallStateEnum = "callStateDiscarded"
	CallStateErrorType          CallStateEnum = "callStateError"
)

func unmarshalCallState(rawMsg *json.RawMessage) (CallState, error) {

	if rawMsg == nil {
		return nil, nil
	}
	var objMap map[string]interface{}
	err := json.Unmarshal(*rawMsg, &objMap)
	if err != nil {
		return nil, err
	}

	switch CallStateEnum(objMap["@type"].(string)) {
	case CallStatePendingType:
		var callStatePending CallStatePending
		err := json.Unmarshal(*rawMsg, &callStatePending)
		return &callStatePending, err

	case CallStateExchangingKeysType:
		var callStateExchangingKeys CallStateExchangingKeys
		err := json.Unmarshal(*rawMsg, &callStateExchangingKeys)
		return &callStateExchangingKeys, err

	case CallStateReadyType:
		var callStateReady CallStateReady
		err := json.Unmarshal(*rawMsg, &callStateReady)
		return &callStateReady, err

	case CallStateHangingUpType:
		var callStateHangingUp CallStateHangingUp
		err := json.Unmarshal(*rawMsg, &callStateHangingUp)
		return &callStateHangingUp, err

	case CallStateDiscardedType:
		var callStateDiscarded CallStateDiscarded
		err := json.Unmarshal(*rawMsg, &callStateDiscarded)
		return &callStateDiscarded, err

	case CallStateErrorType:
		var callStateError CallStateError
		err := json.Unmarshal(*rawMsg, &callStateError)
		return &callStateError, err

	default:
		return nil, fmt.Errorf("Error UnMarshaling, unknown type:" + objMap["@type"].(string))
	}
}

// CallStatePending The call is pending, waiting to be accepted by a user
type CallStatePending struct {
	tdCommon
	IsCreated  bool `json:"is_created"`  // True, if the call has already been created by the server
	IsReceived bool `json:"is_received"` // True, if the call has already been received by the other party
}

// MessageType return the string telegram-type of CallStatePending
func (callStatePending *CallStatePending) MessageType() string {
	return "callStatePending"
}

// NewCallStatePending creates a new CallStatePending
//
// @param isCreated True, if the call has already been created by the server
// @param isReceived True, if the call has already been received by the other party
func NewCallStatePending(isCreated bool, isReceived bool) *CallStatePending {
	callStatePendingTemp := CallStatePending{
		tdCommon:   tdCommon{Type: "callStatePending"},
		IsCreated:  isCreated,
		IsReceived: isReceived,
	}

	return &callStatePendingTemp
}

// GetCallStateEnum return the enum type of this object
func (callStatePending *CallStatePending) GetCallStateEnum() CallStateEnum {
	return CallStatePendingType
}

// CallStateExchangingKeys The call has been answered and encryption keys are being exchanged
type CallStateExchangingKeys struct {
	tdCommon
}

// MessageType return the string telegram-type of CallStateExchangingKeys
func (callStateExchangingKeys *CallStateExchangingKeys) MessageType() string {
	return "callStateExchangingKeys"
}

// NewCallStateExchangingKeys creates a new CallStateExchangingKeys
//
func NewCallStateExchangingKeys() *CallStateExchangingKeys {
	callStateExchangingKeysTemp := CallStateExchangingKeys{
		tdCommon: tdCommon{Type: "callStateExchangingKeys"},
	}

	return &callStateExchangingKeysTemp
}

// GetCallStateEnum return the enum type of this object
func (callStateExchangingKeys *CallStateExchangingKeys) GetCallStateEnum() CallStateEnum {
	return CallStateExchangingKeysType
}

// CallStateReady The call is ready to use
type CallStateReady struct {
	tdCommon
	Protocol      *CallProtocol `json:"protocol"`       // Call protocols supported by the peer
	Servers       []CallServer  `json:"servers"`        // List of available call servers
	Config        string        `json:"config"`         // A JSON-encoded call config
	EncryptionKey []byte        `json:"encryption_key"` // Call encryption key
	Emojis        []string      `json:"emojis"`         // Encryption key emojis fingerprint
	AllowP2p      bool          `json:"allow_p2p"`      // True, if peer-to-peer connection is allowed by users privacy settings
}

// MessageType return the string telegram-type of CallStateReady
func (callStateReady *CallStateReady) MessageType() string {
	return "callStateReady"
}

// NewCallStateReady creates a new CallStateReady
//
// @param protocol Call protocols supported by the peer
// @param servers List of available call servers
// @param config A JSON-encoded call config
// @param encryptionKey Call encryption key
// @param emojis Encryption key emojis fingerprint
// @param allowP2p True, if peer-to-peer connection is allowed by users privacy settings
func NewCallStateReady(protocol *CallProtocol, servers []CallServer, config string, encryptionKey []byte, emojis []string, allowP2p bool) *CallStateReady {
	callStateReadyTemp := CallStateReady{
		tdCommon:      tdCommon{Type: "callStateReady"},
		Protocol:      protocol,
		Servers:       servers,
		Config:        config,
		EncryptionKey: encryptionKey,
		Emojis:        emojis,
		AllowP2p:      allowP2p,
	}

	return &callStateReadyTemp
}

// GetCallStateEnum return the enum type of this object
func (callStateReady *CallStateReady) GetCallStateEnum() CallStateEnum {
	return CallStateReadyType
}

// CallStateHangingUp The call is hanging up after discardCall has been called
type CallStateHangingUp struct {
	tdCommon
}

// MessageType return the string telegram-type of CallStateHangingUp
func (callStateHangingUp *CallStateHangingUp) MessageType() string {
	return "callStateHangingUp"
}

// NewCallStateHangingUp creates a new CallStateHangingUp
//
func NewCallStateHangingUp() *CallStateHangingUp {
	callStateHangingUpTemp := CallStateHangingUp{
		tdCommon: tdCommon{Type: "callStateHangingUp"},
	}

	return &callStateHangingUpTemp
}

// GetCallStateEnum return the enum type of this object
func (callStateHangingUp *CallStateHangingUp) GetCallStateEnum() CallStateEnum {
	return CallStateHangingUpType
}

// CallStateDiscarded The call has ended successfully
type CallStateDiscarded struct {
	tdCommon
	Reason               CallDiscardReason `json:"reason"`                 // The reason, why the call has ended
	NeedRating           bool              `json:"need_rating"`            // True, if the call rating must be sent to the server
	NeedDebugInformation bool              `json:"need_debug_information"` // True, if the call debug information must be sent to the server
}

// MessageType return the string telegram-type of CallStateDiscarded
func (callStateDiscarded *CallStateDiscarded) MessageType() string {
	return "callStateDiscarded"
}

// NewCallStateDiscarded creates a new CallStateDiscarded
//
// @param reason The reason, why the call has ended
// @param needRating True, if the call rating must be sent to the server
// @param needDebugInformation True, if the call debug information must be sent to the server
func NewCallStateDiscarded(reason CallDiscardReason, needRating bool, needDebugInformation bool) *CallStateDiscarded {
	callStateDiscardedTemp := CallStateDiscarded{
		tdCommon:             tdCommon{Type: "callStateDiscarded"},
		Reason:               reason,
		NeedRating:           needRating,
		NeedDebugInformation: needDebugInformation,
	}

	return &callStateDiscardedTemp
}

// UnmarshalJSON unmarshal to json
func (callStateDiscarded *CallStateDiscarded) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	tempObj := struct {
		tdCommon
		NeedRating           bool `json:"need_rating"`            // True, if the call rating must be sent to the server
		NeedDebugInformation bool `json:"need_debug_information"` // True, if the call debug information must be sent to the server
	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	callStateDiscarded.tdCommon = tempObj.tdCommon
	callStateDiscarded.NeedRating = tempObj.NeedRating
	callStateDiscarded.NeedDebugInformation = tempObj.NeedDebugInformation

	fieldReason, _ := unmarshalCallDiscardReason(objMap["reason"])
	callStateDiscarded.Reason = fieldReason

	return nil
}

// GetCallStateEnum return the enum type of this object
func (callStateDiscarded *CallStateDiscarded) GetCallStateEnum() CallStateEnum {
	return CallStateDiscardedType
}

// CallStateError The call has ended with an error
type CallStateError struct {
	tdCommon
	Error *Error `json:"error"` // Error. An error with the code 4005000 will be returned if an outgoing call is missed because of an expired timeout
}

// MessageType return the string telegram-type of CallStateError
func (callStateError *CallStateError) MessageType() string {
	return "callStateError"
}

// NewCallStateError creates a new CallStateError
//
// @param error Error. An error with the code 4005000 will be returned if an outgoing call is missed because of an expired timeout
func NewCallStateError(error *Error) *CallStateError {
	callStateErrorTemp := CallStateError{
		tdCommon: tdCommon{Type: "callStateError"},
		Error:    error,
	}

	return &callStateErrorTemp
}

// GetCallStateEnum return the enum type of this object
func (callStateError *CallStateError) GetCallStateEnum() CallStateEnum {
	return CallStateErrorType
}
