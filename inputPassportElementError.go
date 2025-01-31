package tdlib

import (
	"encoding/json"
)

// InputPassportElementError Contains the description of an error in a Telegram Passport element; for bots only
type InputPassportElementError struct {
	tdCommon
	Type    PassportElementType             `json:"type"`    // Type of Telegram Passport element that has the error
	Message string                          `json:"message"` // Error message
	Source  InputPassportElementErrorSource `json:"source"`  // Error source
}

// MessageType return the string telegram-type of InputPassportElementError
func (inputPassportElementError *InputPassportElementError) MessageType() string {
	return "inputPassportElementError"
}

// NewInputPassportElementError creates a new InputPassportElementError
//
// @param typeParam Type of Telegram Passport element that has the error
// @param message Error message
// @param source Error source
func NewInputPassportElementError(typeParam PassportElementType, message string, source InputPassportElementErrorSource) *InputPassportElementError {
	inputPassportElementErrorTemp := InputPassportElementError{
		tdCommon: tdCommon{Type: "inputPassportElementError"},
		Type:     typeParam,
		Message:  message,
		Source:   source,
	}

	return &inputPassportElementErrorTemp
}

// UnmarshalJSON unmarshal to json
func (inputPassportElementError *InputPassportElementError) UnmarshalJSON(b []byte) error {
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

	inputPassportElementError.tdCommon = tempObj.tdCommon
	inputPassportElementError.Message = tempObj.Message

	fieldType, _ := unmarshalPassportElementType(objMap["type"])
	inputPassportElementError.Type = fieldType

	fieldSource, _ := unmarshalInputPassportElementErrorSource(objMap["source"])
	inputPassportElementError.Source = fieldSource

	return nil
}
