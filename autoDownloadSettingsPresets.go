package tdlib

import (
	"encoding/json"
	"fmt"
)

// AutoDownloadSettingsPresets Contains auto-download settings presets for the current user
type AutoDownloadSettingsPresets struct {
	tdCommon
	Low    *AutoDownloadSettings `json:"low"`    // Preset with lowest settings; supposed to be used by default when roaming
	Medium *AutoDownloadSettings `json:"medium"` // Preset with medium settings; supposed to be used by default when using mobile data
	High   *AutoDownloadSettings `json:"high"`   // Preset with highest settings; supposed to be used by default when connected on Wi-Fi
}

// MessageType return the string telegram-type of AutoDownloadSettingsPresets
func (autoDownloadSettingsPresets *AutoDownloadSettingsPresets) MessageType() string {
	return "autoDownloadSettingsPresets"
}

// NewAutoDownloadSettingsPresets creates a new AutoDownloadSettingsPresets
//
// @param low Preset with lowest settings; supposed to be used by default when roaming
// @param medium Preset with medium settings; supposed to be used by default when using mobile data
// @param high Preset with highest settings; supposed to be used by default when connected on Wi-Fi
func NewAutoDownloadSettingsPresets(low *AutoDownloadSettings, medium *AutoDownloadSettings, high *AutoDownloadSettings) *AutoDownloadSettingsPresets {
	autoDownloadSettingsPresetsTemp := AutoDownloadSettingsPresets{
		tdCommon: tdCommon{Type: "autoDownloadSettingsPresets"},
		Low:      low,
		Medium:   medium,
		High:     high,
	}

	return &autoDownloadSettingsPresetsTemp
}

// GetAutoDownloadSettingsPresets Returns auto-download settings presets for the current user
func (client *Client) GetAutoDownloadSettingsPresets() (*AutoDownloadSettingsPresets, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "getAutoDownloadSettingsPresets",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var autoDownloadSettingsPresets AutoDownloadSettingsPresets
	err = json.Unmarshal(result.Raw, &autoDownloadSettingsPresets)
	return &autoDownloadSettingsPresets, err
}
