package tdlib

import (
	"encoding/json"
	"fmt"
)

// TemporaryPasswordState Returns information about the availability of a temporary password, which can be used for payments
type TemporaryPasswordState struct {
	tdCommon
	HasPassword bool  `json:"has_password"` // True, if a temporary password is available
	ValidFor    int32 `json:"valid_for"`    // Time left before the temporary password expires, in seconds
}

// MessageType return the string telegram-type of TemporaryPasswordState
func (temporaryPasswordState *TemporaryPasswordState) MessageType() string {
	return "temporaryPasswordState"
}

// NewTemporaryPasswordState creates a new TemporaryPasswordState
//
// @param hasPassword True, if a temporary password is available
// @param validFor Time left before the temporary password expires, in seconds
func NewTemporaryPasswordState(hasPassword bool, validFor int32) *TemporaryPasswordState {
	temporaryPasswordStateTemp := TemporaryPasswordState{
		tdCommon:    tdCommon{Type: "temporaryPasswordState"},
		HasPassword: hasPassword,
		ValidFor:    validFor,
	}

	return &temporaryPasswordStateTemp
}

// CreateTemporaryPassword Creates a new temporary password for processing payments
// @param password Persistent user password
// @param validFor Time during which the temporary password will be valid, in seconds; must be between 60 and 86400
func (client *Client) CreateTemporaryPassword(password string, validFor int32) (*TemporaryPasswordState, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":     "createTemporaryPassword",
		"password":  password,
		"valid_for": validFor,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var temporaryPasswordState TemporaryPasswordState
	err = json.Unmarshal(result.Raw, &temporaryPasswordState)
	return &temporaryPasswordState, err
}

// GetTemporaryPasswordState Returns information about the current temporary password
func (client *Client) GetTemporaryPasswordState() (*TemporaryPasswordState, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "getTemporaryPasswordState",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var temporaryPasswordState TemporaryPasswordState
	err = json.Unmarshal(result.Raw, &temporaryPasswordState)
	return &temporaryPasswordState, err
}
