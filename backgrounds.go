package tdlib

import (
	"encoding/json"
	"fmt"
)

// Backgrounds Contains a list of backgrounds
type Backgrounds struct {
	tdCommon
	Backgrounds []Background `json:"backgrounds"` // A list of backgrounds
}

// MessageType return the string telegram-type of Backgrounds
func (backgrounds *Backgrounds) MessageType() string {
	return "backgrounds"
}

// NewBackgrounds creates a new Backgrounds
//
// @param backgrounds A list of backgrounds
func NewBackgrounds(backgrounds []Background) *Backgrounds {
	backgroundsTemp := Backgrounds{
		tdCommon:    tdCommon{Type: "backgrounds"},
		Backgrounds: backgrounds,
	}

	return &backgroundsTemp
}

// GetBackgrounds Returns backgrounds installed by the user
// @param forDarkTheme True, if the backgrounds must be ordered for dark theme
func (client *Client) GetBackgrounds(forDarkTheme bool) (*Backgrounds, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":          "getBackgrounds",
		"for_dark_theme": forDarkTheme,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var backgrounds Backgrounds
	err = json.Unmarshal(result.Raw, &backgrounds)
	return &backgrounds, err
}
