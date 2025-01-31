package tdlib

import (
	"encoding/json"
	"fmt"
)

// UserPrivacySettingRules A list of privacy rules. Rules are matched in the specified order. The first matched rule defines the privacy setting for a given user. If no rule matches, the action is not allowed
type UserPrivacySettingRules struct {
	tdCommon
	Rules []UserPrivacySettingRule `json:"rules"` // A list of rules
}

// MessageType return the string telegram-type of UserPrivacySettingRules
func (userPrivacySettingRules *UserPrivacySettingRules) MessageType() string {
	return "userPrivacySettingRules"
}

// NewUserPrivacySettingRules creates a new UserPrivacySettingRules
//
// @param rules A list of rules
func NewUserPrivacySettingRules(rules []UserPrivacySettingRule) *UserPrivacySettingRules {
	userPrivacySettingRulesTemp := UserPrivacySettingRules{
		tdCommon: tdCommon{Type: "userPrivacySettingRules"},
		Rules:    rules,
	}

	return &userPrivacySettingRulesTemp
}

// GetUserPrivacySettingRules Returns the current privacy settings
// @param setting The privacy setting
func (client *Client) GetUserPrivacySettingRules(setting UserPrivacySetting) (*UserPrivacySettingRules, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "getUserPrivacySettingRules",
		"setting": setting,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var userPrivacySettingRules UserPrivacySettingRules
	err = json.Unmarshal(result.Raw, &userPrivacySettingRules)
	return &userPrivacySettingRules, err
}
