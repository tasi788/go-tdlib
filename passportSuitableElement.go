package tdlib

import (
	"encoding/json"
)

// PassportSuitableElement Contains information about a Telegram Passport element that was requested by a service
type PassportSuitableElement struct {
	tdCommon
	Type                  PassportElementType `json:"type"`                    // Type of the element
	IsSelfieRequired      bool                `json:"is_selfie_required"`      // True, if a selfie is required with the identity document
	IsTranslationRequired bool                `json:"is_translation_required"` // True, if a certified English translation is required with the document
	IsNativeNameRequired  bool                `json:"is_native_name_required"` // True, if personal details must include the user's name in the language of their country of residence
}

// MessageType return the string telegram-type of PassportSuitableElement
func (passportSuitableElement *PassportSuitableElement) MessageType() string {
	return "passportSuitableElement"
}

// NewPassportSuitableElement creates a new PassportSuitableElement
//
// @param typeParam Type of the element
// @param isSelfieRequired True, if a selfie is required with the identity document
// @param isTranslationRequired True, if a certified English translation is required with the document
// @param isNativeNameRequired True, if personal details must include the user's name in the language of their country of residence
func NewPassportSuitableElement(typeParam PassportElementType, isSelfieRequired bool, isTranslationRequired bool, isNativeNameRequired bool) *PassportSuitableElement {
	passportSuitableElementTemp := PassportSuitableElement{
		tdCommon:              tdCommon{Type: "passportSuitableElement"},
		Type:                  typeParam,
		IsSelfieRequired:      isSelfieRequired,
		IsTranslationRequired: isTranslationRequired,
		IsNativeNameRequired:  isNativeNameRequired,
	}

	return &passportSuitableElementTemp
}

// UnmarshalJSON unmarshal to json
func (passportSuitableElement *PassportSuitableElement) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	tempObj := struct {
		tdCommon
		IsSelfieRequired      bool `json:"is_selfie_required"`      // True, if a selfie is required with the identity document
		IsTranslationRequired bool `json:"is_translation_required"` // True, if a certified English translation is required with the document
		IsNativeNameRequired  bool `json:"is_native_name_required"` // True, if personal details must include the user's name in the language of their country of residence
	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	passportSuitableElement.tdCommon = tempObj.tdCommon
	passportSuitableElement.IsSelfieRequired = tempObj.IsSelfieRequired
	passportSuitableElement.IsTranslationRequired = tempObj.IsTranslationRequired
	passportSuitableElement.IsNativeNameRequired = tempObj.IsNativeNameRequired

	fieldType, _ := unmarshalPassportElementType(objMap["type"])
	passportSuitableElement.Type = fieldType

	return nil
}
