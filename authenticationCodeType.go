package tdlib

import (
	"encoding/json"
	"fmt"
)

// AuthenticationCodeType Provides information about the method by which an authentication code is delivered to the user
type AuthenticationCodeType interface {
	GetAuthenticationCodeTypeEnum() AuthenticationCodeTypeEnum
}

// AuthenticationCodeTypeEnum Alias for abstract AuthenticationCodeType 'Sub-Classes', used as constant-enum here
type AuthenticationCodeTypeEnum string

// AuthenticationCodeType enums
const (
	AuthenticationCodeTypeTelegramMessageType AuthenticationCodeTypeEnum = "authenticationCodeTypeTelegramMessage"
	AuthenticationCodeTypeSmsType             AuthenticationCodeTypeEnum = "authenticationCodeTypeSms"
	AuthenticationCodeTypeCallType            AuthenticationCodeTypeEnum = "authenticationCodeTypeCall"
	AuthenticationCodeTypeFlashCallType       AuthenticationCodeTypeEnum = "authenticationCodeTypeFlashCall"
	AuthenticationCodeTypeMissedCallType      AuthenticationCodeTypeEnum = "authenticationCodeTypeMissedCall"
)

func unmarshalAuthenticationCodeType(rawMsg *json.RawMessage) (AuthenticationCodeType, error) {

	if rawMsg == nil {
		return nil, nil
	}
	var objMap map[string]interface{}
	err := json.Unmarshal(*rawMsg, &objMap)
	if err != nil {
		return nil, err
	}

	switch AuthenticationCodeTypeEnum(objMap["@type"].(string)) {
	case AuthenticationCodeTypeTelegramMessageType:
		var authenticationCodeTypeTelegramMessage AuthenticationCodeTypeTelegramMessage
		err := json.Unmarshal(*rawMsg, &authenticationCodeTypeTelegramMessage)
		return &authenticationCodeTypeTelegramMessage, err

	case AuthenticationCodeTypeSmsType:
		var authenticationCodeTypeSms AuthenticationCodeTypeSms
		err := json.Unmarshal(*rawMsg, &authenticationCodeTypeSms)
		return &authenticationCodeTypeSms, err

	case AuthenticationCodeTypeCallType:
		var authenticationCodeTypeCall AuthenticationCodeTypeCall
		err := json.Unmarshal(*rawMsg, &authenticationCodeTypeCall)
		return &authenticationCodeTypeCall, err

	case AuthenticationCodeTypeFlashCallType:
		var authenticationCodeTypeFlashCall AuthenticationCodeTypeFlashCall
		err := json.Unmarshal(*rawMsg, &authenticationCodeTypeFlashCall)
		return &authenticationCodeTypeFlashCall, err

	case AuthenticationCodeTypeMissedCallType:
		var authenticationCodeTypeMissedCall AuthenticationCodeTypeMissedCall
		err := json.Unmarshal(*rawMsg, &authenticationCodeTypeMissedCall)
		return &authenticationCodeTypeMissedCall, err

	default:
		return nil, fmt.Errorf("Error UnMarshaling, unknown type:" + objMap["@type"].(string))
	}
}

// AuthenticationCodeTypeTelegramMessage An authentication code is delivered via a private Telegram message, which can be viewed from another active session
type AuthenticationCodeTypeTelegramMessage struct {
	tdCommon
	Length int32 `json:"length"` // Length of the code
}

// MessageType return the string telegram-type of AuthenticationCodeTypeTelegramMessage
func (authenticationCodeTypeTelegramMessage *AuthenticationCodeTypeTelegramMessage) MessageType() string {
	return "authenticationCodeTypeTelegramMessage"
}

// NewAuthenticationCodeTypeTelegramMessage creates a new AuthenticationCodeTypeTelegramMessage
//
// @param length Length of the code
func NewAuthenticationCodeTypeTelegramMessage(length int32) *AuthenticationCodeTypeTelegramMessage {
	authenticationCodeTypeTelegramMessageTemp := AuthenticationCodeTypeTelegramMessage{
		tdCommon: tdCommon{Type: "authenticationCodeTypeTelegramMessage"},
		Length:   length,
	}

	return &authenticationCodeTypeTelegramMessageTemp
}

// GetAuthenticationCodeTypeEnum return the enum type of this object
func (authenticationCodeTypeTelegramMessage *AuthenticationCodeTypeTelegramMessage) GetAuthenticationCodeTypeEnum() AuthenticationCodeTypeEnum {
	return AuthenticationCodeTypeTelegramMessageType
}

// AuthenticationCodeTypeSms An authentication code is delivered via an SMS message to the specified phone number
type AuthenticationCodeTypeSms struct {
	tdCommon
	Length int32 `json:"length"` // Length of the code
}

// MessageType return the string telegram-type of AuthenticationCodeTypeSms
func (authenticationCodeTypeSms *AuthenticationCodeTypeSms) MessageType() string {
	return "authenticationCodeTypeSms"
}

// NewAuthenticationCodeTypeSms creates a new AuthenticationCodeTypeSms
//
// @param length Length of the code
func NewAuthenticationCodeTypeSms(length int32) *AuthenticationCodeTypeSms {
	authenticationCodeTypeSmsTemp := AuthenticationCodeTypeSms{
		tdCommon: tdCommon{Type: "authenticationCodeTypeSms"},
		Length:   length,
	}

	return &authenticationCodeTypeSmsTemp
}

// GetAuthenticationCodeTypeEnum return the enum type of this object
func (authenticationCodeTypeSms *AuthenticationCodeTypeSms) GetAuthenticationCodeTypeEnum() AuthenticationCodeTypeEnum {
	return AuthenticationCodeTypeSmsType
}

// AuthenticationCodeTypeCall An authentication code is delivered via a phone call to the specified phone number
type AuthenticationCodeTypeCall struct {
	tdCommon
	Length int32 `json:"length"` // Length of the code
}

// MessageType return the string telegram-type of AuthenticationCodeTypeCall
func (authenticationCodeTypeCall *AuthenticationCodeTypeCall) MessageType() string {
	return "authenticationCodeTypeCall"
}

// NewAuthenticationCodeTypeCall creates a new AuthenticationCodeTypeCall
//
// @param length Length of the code
func NewAuthenticationCodeTypeCall(length int32) *AuthenticationCodeTypeCall {
	authenticationCodeTypeCallTemp := AuthenticationCodeTypeCall{
		tdCommon: tdCommon{Type: "authenticationCodeTypeCall"},
		Length:   length,
	}

	return &authenticationCodeTypeCallTemp
}

// GetAuthenticationCodeTypeEnum return the enum type of this object
func (authenticationCodeTypeCall *AuthenticationCodeTypeCall) GetAuthenticationCodeTypeEnum() AuthenticationCodeTypeEnum {
	return AuthenticationCodeTypeCallType
}

// AuthenticationCodeTypeFlashCall An authentication code is delivered by an immediately canceled call to the specified phone number. The phone number from which the call was made is the code that should be entered automatically
type AuthenticationCodeTypeFlashCall struct {
	tdCommon
	Pattern string `json:"pattern"` // Pattern of the phone number from which the call will be made
}

// MessageType return the string telegram-type of AuthenticationCodeTypeFlashCall
func (authenticationCodeTypeFlashCall *AuthenticationCodeTypeFlashCall) MessageType() string {
	return "authenticationCodeTypeFlashCall"
}

// NewAuthenticationCodeTypeFlashCall creates a new AuthenticationCodeTypeFlashCall
//
// @param pattern Pattern of the phone number from which the call will be made
func NewAuthenticationCodeTypeFlashCall(pattern string) *AuthenticationCodeTypeFlashCall {
	authenticationCodeTypeFlashCallTemp := AuthenticationCodeTypeFlashCall{
		tdCommon: tdCommon{Type: "authenticationCodeTypeFlashCall"},
		Pattern:  pattern,
	}

	return &authenticationCodeTypeFlashCallTemp
}

// GetAuthenticationCodeTypeEnum return the enum type of this object
func (authenticationCodeTypeFlashCall *AuthenticationCodeTypeFlashCall) GetAuthenticationCodeTypeEnum() AuthenticationCodeTypeEnum {
	return AuthenticationCodeTypeFlashCallType
}

// AuthenticationCodeTypeMissedCall An authentication code is delivered by an immediately canceled call to the specified phone number. The phone number from which the call was made is the code that should be entered manually by the user
type AuthenticationCodeTypeMissedCall struct {
	tdCommon
	PhoneNumberPrefix string `json:"phone_number_prefix"` // Prefix of the phone number from which the call will be made
	Length            int32  `json:"length"`              // Number of digits in the code, excluding the prefix
}

// MessageType return the string telegram-type of AuthenticationCodeTypeMissedCall
func (authenticationCodeTypeMissedCall *AuthenticationCodeTypeMissedCall) MessageType() string {
	return "authenticationCodeTypeMissedCall"
}

// NewAuthenticationCodeTypeMissedCall creates a new AuthenticationCodeTypeMissedCall
//
// @param phoneNumberPrefix Prefix of the phone number from which the call will be made
// @param length Number of digits in the code, excluding the prefix
func NewAuthenticationCodeTypeMissedCall(phoneNumberPrefix string, length int32) *AuthenticationCodeTypeMissedCall {
	authenticationCodeTypeMissedCallTemp := AuthenticationCodeTypeMissedCall{
		tdCommon:          tdCommon{Type: "authenticationCodeTypeMissedCall"},
		PhoneNumberPrefix: phoneNumberPrefix,
		Length:            length,
	}

	return &authenticationCodeTypeMissedCallTemp
}

// GetAuthenticationCodeTypeEnum return the enum type of this object
func (authenticationCodeTypeMissedCall *AuthenticationCodeTypeMissedCall) GetAuthenticationCodeTypeEnum() AuthenticationCodeTypeEnum {
	return AuthenticationCodeTypeMissedCallType
}
