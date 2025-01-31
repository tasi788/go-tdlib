package tdlib

import (
	"encoding/json"
	"fmt"
)

// LanguagePackInfo Contains information about a language pack
type LanguagePackInfo struct {
	tdCommon
	Id                    string `json:"id"`                      // Unique language pack identifier
	BaseLanguagePackId    string `json:"base_language_pack_id"`   // Identifier of a base language pack; may be empty. If a string is missed in the language pack, then it must be fetched from base language pack. Unsupported in custom language packs
	Name                  string `json:"name"`                    // Language name
	NativeName            string `json:"native_name"`             // Name of the language in that language
	PluralCode            string `json:"plural_code"`             // A language code to be used to apply plural forms. See https://www.unicode.org/cldr/charts/latest/supplemental/language_plural_rules.html for more info
	IsOfficial            bool   `json:"is_official"`             // True, if the language pack is official
	IsRtl                 bool   `json:"is_rtl"`                  // True, if the language pack strings are RTL
	IsBeta                bool   `json:"is_beta"`                 // True, if the language pack is a beta language pack
	IsInstalled           bool   `json:"is_installed"`            // True, if the language pack is installed by the current user
	TotalStringCount      int32  `json:"total_string_count"`      // Total number of non-deleted strings from the language pack
	TranslatedStringCount int32  `json:"translated_string_count"` // Total number of translated strings from the language pack
	LocalStringCount      int32  `json:"local_string_count"`      // Total number of non-deleted strings from the language pack available locally
	TranslationUrl        string `json:"translation_url"`         // Link to language translation interface; empty for custom local language packs
}

// MessageType return the string telegram-type of LanguagePackInfo
func (languagePackInfo *LanguagePackInfo) MessageType() string {
	return "languagePackInfo"
}

// NewLanguagePackInfo creates a new LanguagePackInfo
//
// @param id Unique language pack identifier
// @param baseLanguagePackId Identifier of a base language pack; may be empty. If a string is missed in the language pack, then it must be fetched from base language pack. Unsupported in custom language packs
// @param name Language name
// @param nativeName Name of the language in that language
// @param pluralCode A language code to be used to apply plural forms. See https://www.unicode.org/cldr/charts/latest/supplemental/language_plural_rules.html for more info
// @param isOfficial True, if the language pack is official
// @param isRtl True, if the language pack strings are RTL
// @param isBeta True, if the language pack is a beta language pack
// @param isInstalled True, if the language pack is installed by the current user
// @param totalStringCount Total number of non-deleted strings from the language pack
// @param translatedStringCount Total number of translated strings from the language pack
// @param localStringCount Total number of non-deleted strings from the language pack available locally
// @param translationUrl Link to language translation interface; empty for custom local language packs
func NewLanguagePackInfo(id string, baseLanguagePackId string, name string, nativeName string, pluralCode string, isOfficial bool, isRtl bool, isBeta bool, isInstalled bool, totalStringCount int32, translatedStringCount int32, localStringCount int32, translationUrl string) *LanguagePackInfo {
	languagePackInfoTemp := LanguagePackInfo{
		tdCommon:              tdCommon{Type: "languagePackInfo"},
		Id:                    id,
		BaseLanguagePackId:    baseLanguagePackId,
		Name:                  name,
		NativeName:            nativeName,
		PluralCode:            pluralCode,
		IsOfficial:            isOfficial,
		IsRtl:                 isRtl,
		IsBeta:                isBeta,
		IsInstalled:           isInstalled,
		TotalStringCount:      totalStringCount,
		TranslatedStringCount: translatedStringCount,
		LocalStringCount:      localStringCount,
		TranslationUrl:        translationUrl,
	}

	return &languagePackInfoTemp
}

// GetLanguagePackInfo Returns information about a language pack. Returned language pack identifier may be different from a provided one. Can be called before authorization
// @param languagePackId Language pack identifier
func (client *Client) GetLanguagePackInfo(languagePackId string) (*LanguagePackInfo, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":            "getLanguagePackInfo",
		"language_pack_id": languagePackId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var languagePackInfo LanguagePackInfo
	err = json.Unmarshal(result.Raw, &languagePackInfo)
	return &languagePackInfo, err
}
