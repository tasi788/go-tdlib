package tdlib

import (
	"encoding/json"
	"fmt"
)

// Updates Contains a list of updates
type Updates struct {
	tdCommon
	Updates []Update `json:"updates"` // List of updates
}

// MessageType return the string telegram-type of Updates
func (updates *Updates) MessageType() string {
	return "updates"
}

// NewUpdates creates a new Updates
//
// @param updates List of updates
func NewUpdates(updates []Update) *Updates {
	updatesTemp := Updates{
		tdCommon: tdCommon{Type: "updates"},
		Updates:  updates,
	}

	return &updatesTemp
}

// GetCurrentState Returns all updates needed to restore current TDLib state, i.e. all actual UpdateAuthorizationState/UpdateUser/UpdateNewChat and others. This is especially useful if TDLib is run in a separate process. Can be called before initialization
func (client *Client) GetCurrentState() (*Updates, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "getCurrentState",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var updates Updates
	err = json.Unmarshal(result.Raw, &updates)
	return &updates, err
}
