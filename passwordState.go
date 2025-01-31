package tdlib

import (
	"encoding/json"
	"fmt"
)

// PasswordState Represents the current state of 2-step verification
type PasswordState struct {
	tdCommon
	HasPassword                  bool                                `json:"has_password"`                     // True, if a 2-step verification password is set
	PasswordHint                 string                              `json:"password_hint"`                    // Hint for the password; may be empty
	HasRecoveryEmailAddress      bool                                `json:"has_recovery_email_address"`       // True, if a recovery email is set
	HasPassportData              bool                                `json:"has_passport_data"`                // True, if some Telegram Passport elements were saved
	RecoveryEmailAddressCodeInfo *EmailAddressAuthenticationCodeInfo `json:"recovery_email_address_code_info"` // Information about the recovery email address to which the confirmation email was sent; may be null
	PendingResetDate             int32                               `json:"pending_reset_date"`               // If not 0, point in time (Unix timestamp) after which the password can be reset immediately using resetPassword
}

// MessageType return the string telegram-type of PasswordState
func (passwordState *PasswordState) MessageType() string {
	return "passwordState"
}

// NewPasswordState creates a new PasswordState
//
// @param hasPassword True, if a 2-step verification password is set
// @param passwordHint Hint for the password; may be empty
// @param hasRecoveryEmailAddress True, if a recovery email is set
// @param hasPassportData True, if some Telegram Passport elements were saved
// @param recoveryEmailAddressCodeInfo Information about the recovery email address to which the confirmation email was sent; may be null
// @param pendingResetDate If not 0, point in time (Unix timestamp) after which the password can be reset immediately using resetPassword
func NewPasswordState(hasPassword bool, passwordHint string, hasRecoveryEmailAddress bool, hasPassportData bool, recoveryEmailAddressCodeInfo *EmailAddressAuthenticationCodeInfo, pendingResetDate int32) *PasswordState {
	passwordStateTemp := PasswordState{
		tdCommon:                     tdCommon{Type: "passwordState"},
		HasPassword:                  hasPassword,
		PasswordHint:                 passwordHint,
		HasRecoveryEmailAddress:      hasRecoveryEmailAddress,
		HasPassportData:              hasPassportData,
		RecoveryEmailAddressCodeInfo: recoveryEmailAddressCodeInfo,
		PendingResetDate:             pendingResetDate,
	}

	return &passwordStateTemp
}

// GetPasswordState Returns the current state of 2-step verification
func (client *Client) GetPasswordState() (*PasswordState, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "getPasswordState",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var passwordState PasswordState
	err = json.Unmarshal(result.Raw, &passwordState)
	return &passwordState, err
}

// SetPassword Changes the password for the current user. If a new recovery email address is specified, then the change will not be applied until the new recovery email address is confirmed
// @param oldPassword Previous password of the user
// @param newPassword New password of the user; may be empty to remove the password
// @param newHint New password hint; may be empty
// @param setRecoveryEmailAddress Pass true if the recovery email address must be changed
// @param newRecoveryEmailAddress New recovery email address; may be empty
func (client *Client) SetPassword(oldPassword string, newPassword string, newHint string, setRecoveryEmailAddress bool, newRecoveryEmailAddress string) (*PasswordState, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":                      "setPassword",
		"old_password":               oldPassword,
		"new_password":               newPassword,
		"new_hint":                   newHint,
		"set_recovery_email_address": setRecoveryEmailAddress,
		"new_recovery_email_address": newRecoveryEmailAddress,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var passwordState PasswordState
	err = json.Unmarshal(result.Raw, &passwordState)
	return &passwordState, err
}

// SetRecoveryEmailAddress Changes the 2-step verification recovery email address of the user. If a new recovery email address is specified, then the change will not be applied until the new recovery email address is confirmed. If new_recovery_email_address is the same as the email address that is currently set up, this call succeeds immediately and aborts all other requests waiting for an email confirmation
// @param password Password of the current user
// @param newRecoveryEmailAddress New recovery email address
func (client *Client) SetRecoveryEmailAddress(password string, newRecoveryEmailAddress string) (*PasswordState, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":                      "setRecoveryEmailAddress",
		"password":                   password,
		"new_recovery_email_address": newRecoveryEmailAddress,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var passwordState PasswordState
	err = json.Unmarshal(result.Raw, &passwordState)
	return &passwordState, err
}

// CheckRecoveryEmailAddressCode Checks the 2-step verification recovery email address verification code
// @param code Verification code to check
func (client *Client) CheckRecoveryEmailAddressCode(code string) (*PasswordState, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "checkRecoveryEmailAddressCode",
		"code":  code,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var passwordState PasswordState
	err = json.Unmarshal(result.Raw, &passwordState)
	return &passwordState, err
}

// ResendRecoveryEmailAddressCode Resends the 2-step verification recovery email address verification code
func (client *Client) ResendRecoveryEmailAddressCode() (*PasswordState, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "resendRecoveryEmailAddressCode",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var passwordState PasswordState
	err = json.Unmarshal(result.Raw, &passwordState)
	return &passwordState, err
}

// RecoverPassword Recovers the 2-step verification password using a recovery code sent to an email address that was previously set up
// @param recoveryCode Recovery code to check
// @param newPassword New password of the user; may be empty to remove the password
// @param newHint New password hint; may be empty
func (client *Client) RecoverPassword(recoveryCode string, newPassword string, newHint string) (*PasswordState, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":         "recoverPassword",
		"recovery_code": recoveryCode,
		"new_password":  newPassword,
		"new_hint":      newHint,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var passwordState PasswordState
	err = json.Unmarshal(result.Raw, &passwordState)
	return &passwordState, err
}
