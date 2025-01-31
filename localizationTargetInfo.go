package tdlib

import (
	"encoding/json"
	"fmt"
)

// LocalizationTargetInfo Contains information about the current localization target
type LocalizationTargetInfo struct {
	tdCommon
	LanguagePacks []LanguagePackInfo `json:"language_packs"` // List of available language packs for this application
}

// MessageType return the string telegram-type of LocalizationTargetInfo
func (localizationTargetInfo *LocalizationTargetInfo) MessageType() string {
	return "localizationTargetInfo"
}

// NewLocalizationTargetInfo creates a new LocalizationTargetInfo
//
// @param languagePacks List of available language packs for this application
func NewLocalizationTargetInfo(languagePacks []LanguagePackInfo) *LocalizationTargetInfo {
	localizationTargetInfoTemp := LocalizationTargetInfo{
		tdCommon:      tdCommon{Type: "localizationTargetInfo"},
		LanguagePacks: languagePacks,
	}

	return &localizationTargetInfoTemp
}

// GetLocalizationTargetInfo Returns information about the current localization target. This is an offline request if only_local is true. Can be called before authorization
// @param onlyLocal If true, returns only locally available information without sending network requests
func (client *Client) GetLocalizationTargetInfo(onlyLocal bool) (*LocalizationTargetInfo, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":      "getLocalizationTargetInfo",
		"only_local": onlyLocal,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var localizationTargetInfo LocalizationTargetInfo
	err = json.Unmarshal(result.Raw, &localizationTargetInfo)
	return &localizationTargetInfo, err
}
