package tdlib

import (
	"encoding/json"
	"fmt"
)

// SecretChatState Describes the current secret chat state
type SecretChatState interface {
	GetSecretChatStateEnum() SecretChatStateEnum
}

// SecretChatStateEnum Alias for abstract SecretChatState 'Sub-Classes', used as constant-enum here
type SecretChatStateEnum string

// SecretChatState enums
const (
	SecretChatStatePendingType SecretChatStateEnum = "secretChatStatePending"
	SecretChatStateReadyType   SecretChatStateEnum = "secretChatStateReady"
	SecretChatStateClosedType  SecretChatStateEnum = "secretChatStateClosed"
)

func unmarshalSecretChatState(rawMsg *json.RawMessage) (SecretChatState, error) {

	if rawMsg == nil {
		return nil, nil
	}
	var objMap map[string]interface{}
	err := json.Unmarshal(*rawMsg, &objMap)
	if err != nil {
		return nil, err
	}

	switch SecretChatStateEnum(objMap["@type"].(string)) {
	case SecretChatStatePendingType:
		var secretChatStatePending SecretChatStatePending
		err := json.Unmarshal(*rawMsg, &secretChatStatePending)
		return &secretChatStatePending, err

	case SecretChatStateReadyType:
		var secretChatStateReady SecretChatStateReady
		err := json.Unmarshal(*rawMsg, &secretChatStateReady)
		return &secretChatStateReady, err

	case SecretChatStateClosedType:
		var secretChatStateClosed SecretChatStateClosed
		err := json.Unmarshal(*rawMsg, &secretChatStateClosed)
		return &secretChatStateClosed, err

	default:
		return nil, fmt.Errorf("Error UnMarshaling, unknown type:" + objMap["@type"].(string))
	}
}

// SecretChatStatePending The secret chat is not yet created; waiting for the other user to get online
type SecretChatStatePending struct {
	tdCommon
}

// MessageType return the string telegram-type of SecretChatStatePending
func (secretChatStatePending *SecretChatStatePending) MessageType() string {
	return "secretChatStatePending"
}

// NewSecretChatStatePending creates a new SecretChatStatePending
//
func NewSecretChatStatePending() *SecretChatStatePending {
	secretChatStatePendingTemp := SecretChatStatePending{
		tdCommon: tdCommon{Type: "secretChatStatePending"},
	}

	return &secretChatStatePendingTemp
}

// GetSecretChatStateEnum return the enum type of this object
func (secretChatStatePending *SecretChatStatePending) GetSecretChatStateEnum() SecretChatStateEnum {
	return SecretChatStatePendingType
}

// SecretChatStateReady The secret chat is ready to use
type SecretChatStateReady struct {
	tdCommon
}

// MessageType return the string telegram-type of SecretChatStateReady
func (secretChatStateReady *SecretChatStateReady) MessageType() string {
	return "secretChatStateReady"
}

// NewSecretChatStateReady creates a new SecretChatStateReady
//
func NewSecretChatStateReady() *SecretChatStateReady {
	secretChatStateReadyTemp := SecretChatStateReady{
		tdCommon: tdCommon{Type: "secretChatStateReady"},
	}

	return &secretChatStateReadyTemp
}

// GetSecretChatStateEnum return the enum type of this object
func (secretChatStateReady *SecretChatStateReady) GetSecretChatStateEnum() SecretChatStateEnum {
	return SecretChatStateReadyType
}

// SecretChatStateClosed The secret chat is closed
type SecretChatStateClosed struct {
	tdCommon
}

// MessageType return the string telegram-type of SecretChatStateClosed
func (secretChatStateClosed *SecretChatStateClosed) MessageType() string {
	return "secretChatStateClosed"
}

// NewSecretChatStateClosed creates a new SecretChatStateClosed
//
func NewSecretChatStateClosed() *SecretChatStateClosed {
	secretChatStateClosedTemp := SecretChatStateClosed{
		tdCommon: tdCommon{Type: "secretChatStateClosed"},
	}

	return &secretChatStateClosedTemp
}

// GetSecretChatStateEnum return the enum type of this object
func (secretChatStateClosed *SecretChatStateClosed) GetSecretChatStateEnum() SecretChatStateEnum {
	return SecretChatStateClosedType
}
