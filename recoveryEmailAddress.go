package tdlib

import (
	"encoding/json"
	"fmt"
)

// RecoveryEmailAddress Contains information about the current recovery email address
type RecoveryEmailAddress struct {
	tdCommon
	RecoveryEmailAddress string `json:"recovery_email_address"` // Recovery email address
}

// MessageType return the string telegram-type of RecoveryEmailAddress
func (recoveryEmailAddress *RecoveryEmailAddress) MessageType() string {
	return "recoveryEmailAddress"
}

// NewRecoveryEmailAddress creates a new RecoveryEmailAddress
//
// @param recoveryEmailAddress Recovery email address
func NewRecoveryEmailAddress(recoveryEmailAddress string) *RecoveryEmailAddress {
	recoveryEmailAddressTemp := RecoveryEmailAddress{
		tdCommon:             tdCommon{Type: "recoveryEmailAddress"},
		RecoveryEmailAddress: recoveryEmailAddress,
	}

	return &recoveryEmailAddressTemp
}

// GetRecoveryEmailAddress Returns a 2-step verification recovery email address that was previously set up. This method can be used to verify a password provided by the user
// @param password The password for the current user
func (client *Client) GetRecoveryEmailAddress(password string) (*RecoveryEmailAddress, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":    "getRecoveryEmailAddress",
		"password": password,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var recoveryEmailAddress RecoveryEmailAddress
	err = json.Unmarshal(result.Raw, &recoveryEmailAddress)
	return &recoveryEmailAddress, err
}
