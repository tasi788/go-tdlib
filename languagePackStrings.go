package tdlib

import (
	"encoding/json"
	"fmt"
)

// LanguagePackStrings Contains a list of language pack strings
type LanguagePackStrings struct {
	tdCommon
	Strings []LanguagePackString `json:"strings"` // A list of language pack strings
}

// MessageType return the string telegram-type of LanguagePackStrings
func (languagePackStrings *LanguagePackStrings) MessageType() string {
	return "languagePackStrings"
}

// NewLanguagePackStrings creates a new LanguagePackStrings
//
// @param strings A list of language pack strings
func NewLanguagePackStrings(strings []LanguagePackString) *LanguagePackStrings {
	languagePackStringsTemp := LanguagePackStrings{
		tdCommon: tdCommon{Type: "languagePackStrings"},
		Strings:  strings,
	}

	return &languagePackStringsTemp
}

// GetLanguagePackStrings Returns strings from a language pack in the current localization target by their keys. Can be called before authorization
// @param languagePackId Language pack identifier of the strings to be returned
// @param keys Language pack keys of the strings to be returned; leave empty to request all available strings
func (client *Client) GetLanguagePackStrings(languagePackId string, keys []string) (*LanguagePackStrings, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":            "getLanguagePackStrings",
		"language_pack_id": languagePackId,
		"keys":             keys,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var languagePackStrings LanguagePackStrings
	err = json.Unmarshal(result.Raw, &languagePackStrings)
	return &languagePackStrings, err
}
