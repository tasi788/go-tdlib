package tdlib

import (
	"encoding/json"
	"fmt"
)

// Animations Represents a list of animations
type Animations struct {
	tdCommon
	Animations []Animation `json:"animations"` // List of animations
}

// MessageType return the string telegram-type of Animations
func (animations *Animations) MessageType() string {
	return "animations"
}

// NewAnimations creates a new Animations
//
// @param animations List of animations
func NewAnimations(animations []Animation) *Animations {
	animationsTemp := Animations{
		tdCommon:   tdCommon{Type: "animations"},
		Animations: animations,
	}

	return &animationsTemp
}

// GetSavedAnimations Returns saved animations
func (client *Client) GetSavedAnimations() (*Animations, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "getSavedAnimations",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var animations Animations
	err = json.Unmarshal(result.Raw, &animations)
	return &animations, err
}
