package tdlib

import (
	"encoding/json"
	"fmt"
)

// InputCredentials Contains information about the payment method chosen by the user
type InputCredentials interface {
	GetInputCredentialsEnum() InputCredentialsEnum
}

// InputCredentialsEnum Alias for abstract InputCredentials 'Sub-Classes', used as constant-enum here
type InputCredentialsEnum string

// InputCredentials enums
const (
	InputCredentialsSavedType     InputCredentialsEnum = "inputCredentialsSaved"
	InputCredentialsNewType       InputCredentialsEnum = "inputCredentialsNew"
	InputCredentialsApplePayType  InputCredentialsEnum = "inputCredentialsApplePay"
	InputCredentialsGooglePayType InputCredentialsEnum = "inputCredentialsGooglePay"
)

func unmarshalInputCredentials(rawMsg *json.RawMessage) (InputCredentials, error) {

	if rawMsg == nil {
		return nil, nil
	}
	var objMap map[string]interface{}
	err := json.Unmarshal(*rawMsg, &objMap)
	if err != nil {
		return nil, err
	}

	switch InputCredentialsEnum(objMap["@type"].(string)) {
	case InputCredentialsSavedType:
		var inputCredentialsSaved InputCredentialsSaved
		err := json.Unmarshal(*rawMsg, &inputCredentialsSaved)
		return &inputCredentialsSaved, err

	case InputCredentialsNewType:
		var inputCredentialsNew InputCredentialsNew
		err := json.Unmarshal(*rawMsg, &inputCredentialsNew)
		return &inputCredentialsNew, err

	case InputCredentialsApplePayType:
		var inputCredentialsApplePay InputCredentialsApplePay
		err := json.Unmarshal(*rawMsg, &inputCredentialsApplePay)
		return &inputCredentialsApplePay, err

	case InputCredentialsGooglePayType:
		var inputCredentialsGooglePay InputCredentialsGooglePay
		err := json.Unmarshal(*rawMsg, &inputCredentialsGooglePay)
		return &inputCredentialsGooglePay, err

	default:
		return nil, fmt.Errorf("Error UnMarshaling, unknown type:" + objMap["@type"].(string))
	}
}

// InputCredentialsSaved Applies if a user chooses some previously saved payment credentials. To use their previously saved credentials, the user must have a valid temporary password
type InputCredentialsSaved struct {
	tdCommon
	SavedCredentialsId string `json:"saved_credentials_id"` // Identifier of the saved credentials
}

// MessageType return the string telegram-type of InputCredentialsSaved
func (inputCredentialsSaved *InputCredentialsSaved) MessageType() string {
	return "inputCredentialsSaved"
}

// NewInputCredentialsSaved creates a new InputCredentialsSaved
//
// @param savedCredentialsId Identifier of the saved credentials
func NewInputCredentialsSaved(savedCredentialsId string) *InputCredentialsSaved {
	inputCredentialsSavedTemp := InputCredentialsSaved{
		tdCommon:           tdCommon{Type: "inputCredentialsSaved"},
		SavedCredentialsId: savedCredentialsId,
	}

	return &inputCredentialsSavedTemp
}

// GetInputCredentialsEnum return the enum type of this object
func (inputCredentialsSaved *InputCredentialsSaved) GetInputCredentialsEnum() InputCredentialsEnum {
	return InputCredentialsSavedType
}

// InputCredentialsNew Applies if a user enters new credentials on a payment provider website
type InputCredentialsNew struct {
	tdCommon
	Data      string `json:"data"`       // JSON-encoded data with the credential identifier from the payment provider
	AllowSave bool   `json:"allow_save"` // True, if the credential identifier can be saved on the server side
}

// MessageType return the string telegram-type of InputCredentialsNew
func (inputCredentialsNew *InputCredentialsNew) MessageType() string {
	return "inputCredentialsNew"
}

// NewInputCredentialsNew creates a new InputCredentialsNew
//
// @param data JSON-encoded data with the credential identifier from the payment provider
// @param allowSave True, if the credential identifier can be saved on the server side
func NewInputCredentialsNew(data string, allowSave bool) *InputCredentialsNew {
	inputCredentialsNewTemp := InputCredentialsNew{
		tdCommon:  tdCommon{Type: "inputCredentialsNew"},
		Data:      data,
		AllowSave: allowSave,
	}

	return &inputCredentialsNewTemp
}

// GetInputCredentialsEnum return the enum type of this object
func (inputCredentialsNew *InputCredentialsNew) GetInputCredentialsEnum() InputCredentialsEnum {
	return InputCredentialsNewType
}

// InputCredentialsApplePay Applies if a user enters new credentials using Apple Pay
type InputCredentialsApplePay struct {
	tdCommon
	Data string `json:"data"` // JSON-encoded data with the credential identifier
}

// MessageType return the string telegram-type of InputCredentialsApplePay
func (inputCredentialsApplePay *InputCredentialsApplePay) MessageType() string {
	return "inputCredentialsApplePay"
}

// NewInputCredentialsApplePay creates a new InputCredentialsApplePay
//
// @param data JSON-encoded data with the credential identifier
func NewInputCredentialsApplePay(data string) *InputCredentialsApplePay {
	inputCredentialsApplePayTemp := InputCredentialsApplePay{
		tdCommon: tdCommon{Type: "inputCredentialsApplePay"},
		Data:     data,
	}

	return &inputCredentialsApplePayTemp
}

// GetInputCredentialsEnum return the enum type of this object
func (inputCredentialsApplePay *InputCredentialsApplePay) GetInputCredentialsEnum() InputCredentialsEnum {
	return InputCredentialsApplePayType
}

// InputCredentialsGooglePay Applies if a user enters new credentials using Google Pay
type InputCredentialsGooglePay struct {
	tdCommon
	Data string `json:"data"` // JSON-encoded data with the credential identifier
}

// MessageType return the string telegram-type of InputCredentialsGooglePay
func (inputCredentialsGooglePay *InputCredentialsGooglePay) MessageType() string {
	return "inputCredentialsGooglePay"
}

// NewInputCredentialsGooglePay creates a new InputCredentialsGooglePay
//
// @param data JSON-encoded data with the credential identifier
func NewInputCredentialsGooglePay(data string) *InputCredentialsGooglePay {
	inputCredentialsGooglePayTemp := InputCredentialsGooglePay{
		tdCommon: tdCommon{Type: "inputCredentialsGooglePay"},
		Data:     data,
	}

	return &inputCredentialsGooglePayTemp
}

// GetInputCredentialsEnum return the enum type of this object
func (inputCredentialsGooglePay *InputCredentialsGooglePay) GetInputCredentialsEnum() InputCredentialsEnum {
	return InputCredentialsGooglePayType
}
