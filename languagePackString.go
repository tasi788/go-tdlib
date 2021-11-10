package tdlib

import (
	"encoding/json"
)

// LanguagePackString Represents one language pack string
type LanguagePackString struct {
	tdCommon
	Key   string                  `json:"key"`   // String key
	Value LanguagePackStringValue `json:"value"` // String value; pass null if the string needs to be taken from the built-in English language pack
}

// MessageType return the string telegram-type of LanguagePackString
func (languagePackString *LanguagePackString) MessageType() string {
	return "languagePackString"
}

// NewLanguagePackString creates a new LanguagePackString
//
// @param key String key
// @param value String value; pass null if the string needs to be taken from the built-in English language pack
func NewLanguagePackString(key string, value LanguagePackStringValue) *LanguagePackString {
	languagePackStringTemp := LanguagePackString{
		tdCommon: tdCommon{Type: "languagePackString"},
		Key:      key,
		Value:    value,
	}

	return &languagePackStringTemp
}

// UnmarshalJSON unmarshal to json
func (languagePackString *LanguagePackString) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	tempObj := struct {
		tdCommon
		Key string `json:"key"` // String key

	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	languagePackString.tdCommon = tempObj.tdCommon
	languagePackString.Key = tempObj.Key

	fieldValue, _ := unmarshalLanguagePackStringValue(objMap["value"])
	languagePackString.Value = fieldValue

	return nil
}
