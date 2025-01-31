package tdlib

import (
	"encoding/json"
)

// PassportElementError Contains the description of an error in a Telegram Passport element
type PassportElementError struct {
	tdCommon
	Type    PassportElementType        `json:"type"`    // Type of the Telegram Passport element which has the error
	Message string                     `json:"message"` // Error message
	Source  PassportElementErrorSource `json:"source"`  // Error source
}

// MessageType return the string telegram-type of PassportElementError
func (passportElementError *PassportElementError) MessageType() string {
	return "passportElementError"
}

// NewPassportElementError creates a new PassportElementError
//
// @param typeParam Type of the Telegram Passport element which has the error
// @param message Error message
// @param source Error source
func NewPassportElementError(typeParam PassportElementType, message string, source PassportElementErrorSource) *PassportElementError {
	passportElementErrorTemp := PassportElementError{
		tdCommon: tdCommon{Type: "passportElementError"},
		Type:     typeParam,
		Message:  message,
		Source:   source,
	}

	return &passportElementErrorTemp
}

// UnmarshalJSON unmarshal to json
func (passportElementError *PassportElementError) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	tempObj := struct {
		tdCommon
		Message string `json:"message"` // Error message

	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	passportElementError.tdCommon = tempObj.tdCommon
	passportElementError.Message = tempObj.Message

	fieldType, _ := unmarshalPassportElementType(objMap["type"])
	passportElementError.Type = fieldType

	fieldSource, _ := unmarshalPassportElementErrorSource(objMap["source"])
	passportElementError.Source = fieldSource

	return nil
}
