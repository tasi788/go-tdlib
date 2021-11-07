package tdlib

import (
	"encoding/json"
	"fmt"
)

// GetAuthorizationState Returns the current authorization state; this is an offline request. For informational purposes only. Use updateAuthorizationState instead to maintain the current authorization state. Can be called before initialization
func (client *Client) GetAuthorizationState() (AuthorizationState, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "getAuthorizationState",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	switch AuthorizationStateEnum(result.Data["@type"].(string)) {

	case AuthorizationStateWaitTdlibParametersType:
		var authorizationState AuthorizationStateWaitTdlibParameters
		err = json.Unmarshal(result.Raw, &authorizationState)
		return &authorizationState, err

	case AuthorizationStateWaitEncryptionKeyType:
		var authorizationState AuthorizationStateWaitEncryptionKey
		err = json.Unmarshal(result.Raw, &authorizationState)
		return &authorizationState, err

	case AuthorizationStateWaitPhoneNumberType:
		var authorizationState AuthorizationStateWaitPhoneNumber
		err = json.Unmarshal(result.Raw, &authorizationState)
		return &authorizationState, err

	case AuthorizationStateWaitCodeType:
		var authorizationState AuthorizationStateWaitCode
		err = json.Unmarshal(result.Raw, &authorizationState)
		return &authorizationState, err

	case AuthorizationStateWaitOtherDeviceConfirmationType:
		var authorizationState AuthorizationStateWaitOtherDeviceConfirmation
		err = json.Unmarshal(result.Raw, &authorizationState)
		return &authorizationState, err

	case AuthorizationStateWaitRegistrationType:
		var authorizationState AuthorizationStateWaitRegistration
		err = json.Unmarshal(result.Raw, &authorizationState)
		return &authorizationState, err

	case AuthorizationStateWaitPasswordType:
		var authorizationState AuthorizationStateWaitPassword
		err = json.Unmarshal(result.Raw, &authorizationState)
		return &authorizationState, err

	case AuthorizationStateReadyType:
		var authorizationState AuthorizationStateReady
		err = json.Unmarshal(result.Raw, &authorizationState)
		return &authorizationState, err

	case AuthorizationStateLoggingOutType:
		var authorizationState AuthorizationStateLoggingOut
		err = json.Unmarshal(result.Raw, &authorizationState)
		return &authorizationState, err

	case AuthorizationStateClosingType:
		var authorizationState AuthorizationStateClosing
		err = json.Unmarshal(result.Raw, &authorizationState)
		return &authorizationState, err

	case AuthorizationStateClosedType:
		var authorizationState AuthorizationStateClosed
		err = json.Unmarshal(result.Raw, &authorizationState)
		return &authorizationState, err

	default:
		return nil, fmt.Errorf("Invalid type")
	}
}

// SetTdlibParameters Sets the parameters for TDLib initialization. Works only when the current authorization state is authorizationStateWaitTdlibParameters
// @param parameters Parameters for TDLib initialization
func (client *Client) SetTdlibParameters(parameters *TdlibParameters) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":      "setTdlibParameters",
		"parameters": parameters,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// CheckDatabaseEncryptionKey Checks the database encryption key for correctness. Works only when the current authorization state is authorizationStateWaitEncryptionKey
// @param encryptionKey Encryption key to check or set up
func (client *Client) CheckDatabaseEncryptionKey(encryptionKey []byte) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":          "checkDatabaseEncryptionKey",
		"encryption_key": encryptionKey,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// SetAuthenticationPhoneNumber Sets the phone number of the user and sends an authentication code to the user. Works only when the current authorization state is authorizationStateWaitPhoneNumber,
// @param phoneNumber The phone number of the user, in international format
// @param settings Settings for the authentication of the user's phone number; pass null to use default settings
func (client *Client) SetAuthenticationPhoneNumber(phoneNumber string, settings *PhoneNumberAuthenticationSettings) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":        "setAuthenticationPhoneNumber",
		"phone_number": phoneNumber,
		"settings":     settings,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// ResendAuthenticationCode Re-sends an authentication code to the user. Works only when the current authorization state is authorizationStateWaitCode, the next_code_type of the result is not null and the server-specified timeout has passed
func (client *Client) ResendAuthenticationCode() (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "resendAuthenticationCode",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// CheckAuthenticationCode Checks the authentication code. Works only when the current authorization state is authorizationStateWaitCode
// @param code The verification code received via SMS, Telegram message, phone call, or flash call
func (client *Client) CheckAuthenticationCode(code string) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "checkAuthenticationCode",
		"code":  code,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// RequestQrCodeAuthentication Requests QR code authentication by scanning a QR code on another logged in device. Works only when the current authorization state is authorizationStateWaitPhoneNumber,
// @param otherUserIds List of user identifiers of other users currently using the application
func (client *Client) RequestQrCodeAuthentication(otherUserIds []int64) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":          "requestQrCodeAuthentication",
		"other_user_ids": otherUserIds,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// RegisterUser Finishes user registration. Works only when the current authorization state is authorizationStateWaitRegistration
// @param firstName The first name of the user; 1-64 characters
// @param lastName The last name of the user; 0-64 characters
func (client *Client) RegisterUser(firstName string, lastName string) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":      "registerUser",
		"first_name": firstName,
		"last_name":  lastName,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// CheckAuthenticationPassword Checks the authentication password for correctness. Works only when the current authorization state is authorizationStateWaitPassword
// @param password The password to check
func (client *Client) CheckAuthenticationPassword(password string) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":    "checkAuthenticationPassword",
		"password": password,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// RequestAuthenticationPasswordRecovery Requests to send a password recovery code to an email address that was previously set up. Works only when the current authorization state is authorizationStateWaitPassword
func (client *Client) RequestAuthenticationPasswordRecovery() (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "requestAuthenticationPasswordRecovery",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// CheckAuthenticationPasswordRecoveryCode Checks whether a password recovery code sent to an email address is valid. Works only when the current authorization state is authorizationStateWaitPassword
// @param recoveryCode Recovery code to check
func (client *Client) CheckAuthenticationPasswordRecoveryCode(recoveryCode string) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":         "checkAuthenticationPasswordRecoveryCode",
		"recovery_code": recoveryCode,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// RecoverAuthenticationPassword Recovers the password with a password recovery code sent to an email address that was previously set up. Works only when the current authorization state is authorizationStateWaitPassword
// @param recoveryCode Recovery code to check
// @param newPassword New password of the user; may be empty to remove the password
// @param newHint New password hint; may be empty
func (client *Client) RecoverAuthenticationPassword(recoveryCode string, newPassword string, newHint string) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":         "recoverAuthenticationPassword",
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

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// CheckAuthenticationBotToken Checks the authentication token of a bot; to log in as a bot. Works only when the current authorization state is authorizationStateWaitPhoneNumber. Can be used instead of setAuthenticationPhoneNumber and checkAuthenticationCode to log in
// @param token The bot token
func (client *Client) CheckAuthenticationBotToken(token string) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "checkAuthenticationBotToken",
		"token": token,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var okDummy Ok
	err = json.Unmarshal(result.Raw, &okDummy)
	return &okDummy, err

}

// LogOut Closes the TDLib instance after a proper logout. Requires an available network connection. All local data will be destroyed. After the logout completes, updateAuthorizationState with authorizationStateClosed will be sent
func (client *Client) LogOut() (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "logOut",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// Close Closes the TDLib instance. All databases will be flushed to disk and properly closed. After the close completes, updateAuthorizationState with authorizationStateClosed will be sent. Can be called before initialization
func (client *Client) Close() (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "close",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// Destroy Closes the TDLib instance, destroying all local data without a proper logout. The current user session will remain in the list of all active sessions. All local data will be destroyed. After the destruction completes updateAuthorizationState with authorizationStateClosed will be sent. Can be called before authorization
func (client *Client) Destroy() (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "destroy",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// ConfirmQrCodeAuthentication Confirms QR code authentication on another device. Returns created session on success
// @param link A link from a QR code. The link must be scanned by the in-app camera
func (client *Client) ConfirmQrCodeAuthentication(link string) (*Session, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "confirmQrCodeAuthentication",
		"link":  link,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var session Session
	err = json.Unmarshal(result.Raw, &session)
	return &session, err

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

// SetDatabaseEncryptionKey Changes the database encryption key. Usually the encryption key is never changed and is stored in some OS keychain
// @param newEncryptionKey New encryption key
func (client *Client) SetDatabaseEncryptionKey(newEncryptionKey []byte) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":              "setDatabaseEncryptionKey",
		"new_encryption_key": newEncryptionKey,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

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

// SetRecoveryEmailAddress Changes the 2-step verification recovery email address of the user. If a new recovery email address is specified, then the change will not be applied until the new recovery email address is confirmed.
// @param password
// @param newRecoveryEmailAddress
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
// @param code Verification code
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

// RequestPasswordRecovery Requests to send a 2-step verification password recovery code to an email address that was previously set up
func (client *Client) RequestPasswordRecovery() (*EmailAddressAuthenticationCodeInfo, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "requestPasswordRecovery",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var emailAddressAuthenticationCodeInfo EmailAddressAuthenticationCodeInfo
	err = json.Unmarshal(result.Raw, &emailAddressAuthenticationCodeInfo)
	return &emailAddressAuthenticationCodeInfo, err

}

// CheckPasswordRecoveryCode Checks whether a 2-step verification password recovery code sent to an email address is valid
// @param recoveryCode Recovery code to check
func (client *Client) CheckPasswordRecoveryCode(recoveryCode string) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":         "checkPasswordRecoveryCode",
		"recovery_code": recoveryCode,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

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

// ResetPassword Removes 2-step verification password without previous password and access to recovery email address. The password can't be reset immediately and the request needs to be repeated after the specified time
func (client *Client) ResetPassword() (ResetPasswordResult, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "resetPassword",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	switch ResetPasswordResultEnum(result.Data["@type"].(string)) {

	case ResetPasswordResultOkType:
		var resetPasswordResult ResetPasswordResultOk
		err = json.Unmarshal(result.Raw, &resetPasswordResult)
		return &resetPasswordResult, err

	case ResetPasswordResultPendingType:
		var resetPasswordResult ResetPasswordResultPending
		err = json.Unmarshal(result.Raw, &resetPasswordResult)
		return &resetPasswordResult, err

	case ResetPasswordResultDeclinedType:
		var resetPasswordResult ResetPasswordResultDeclined
		err = json.Unmarshal(result.Raw, &resetPasswordResult)
		return &resetPasswordResult, err

	default:
		return nil, fmt.Errorf("Invalid type")
	}
}

// CancelPasswordReset Cancels reset of 2-step verification password. The method can be called if passwordState.pending_reset_date > 0
func (client *Client) CancelPasswordReset() (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "cancelPasswordReset",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

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

// GetMe Returns the current user
func (client *Client) GetMe() (*User, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "getMe",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var user User
	err = json.Unmarshal(result.Raw, &user)
	return &user, err

}

// GetUser Returns information about a user by their identifier. This is an offline request if the current user is not a bot
// @param userId User identifier
func (client *Client) GetUser(userId int64) (*User, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "getUser",
		"user_id": userId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var userDummy User
	err = json.Unmarshal(result.Raw, &userDummy)
	return &userDummy, err

}

// GetUserFullInfo Returns full information about a user by their identifier
// @param userId User identifier
func (client *Client) GetUserFullInfo(userId int64) (*UserFullInfo, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "getUserFullInfo",
		"user_id": userId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var userFullInfo UserFullInfo
	err = json.Unmarshal(result.Raw, &userFullInfo)
	return &userFullInfo, err

}

// GetBasicGroup Returns information about a basic group by its identifier. This is an offline request if the current user is not a bot
// @param basicGroupId Basic group identifier
func (client *Client) GetBasicGroup(basicGroupId int64) (*BasicGroup, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":          "getBasicGroup",
		"basic_group_id": basicGroupId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var basicGroupDummy BasicGroup
	err = json.Unmarshal(result.Raw, &basicGroupDummy)
	return &basicGroupDummy, err

}

// GetBasicGroupFullInfo Returns full information about a basic group by its identifier
// @param basicGroupId Basic group identifier
func (client *Client) GetBasicGroupFullInfo(basicGroupId int64) (*BasicGroupFullInfo, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":          "getBasicGroupFullInfo",
		"basic_group_id": basicGroupId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var basicGroupFullInfo BasicGroupFullInfo
	err = json.Unmarshal(result.Raw, &basicGroupFullInfo)
	return &basicGroupFullInfo, err

}

// GetSupergroup Returns information about a supergroup or a channel by its identifier. This is an offline request if the current user is not a bot
// @param supergroupId Supergroup or channel identifier
func (client *Client) GetSupergroup(supergroupId int64) (*Supergroup, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":         "getSupergroup",
		"supergroup_id": supergroupId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var supergroupDummy Supergroup
	err = json.Unmarshal(result.Raw, &supergroupDummy)
	return &supergroupDummy, err

}

// GetSupergroupFullInfo Returns full information about a supergroup or a channel by its identifier, cached for up to 1 minute
// @param supergroupId Supergroup or channel identifier
func (client *Client) GetSupergroupFullInfo(supergroupId int64) (*SupergroupFullInfo, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":         "getSupergroupFullInfo",
		"supergroup_id": supergroupId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var supergroupFullInfo SupergroupFullInfo
	err = json.Unmarshal(result.Raw, &supergroupFullInfo)
	return &supergroupFullInfo, err

}

// GetSecretChat Returns information about a secret chat by its identifier. This is an offline request
// @param secretChatId Secret chat identifier
func (client *Client) GetSecretChat(secretChatId int32) (*SecretChat, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":          "getSecretChat",
		"secret_chat_id": secretChatId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var secretChatDummy SecretChat
	err = json.Unmarshal(result.Raw, &secretChatDummy)
	return &secretChatDummy, err

}

// GetChat Returns information about a chat by its identifier, this is an offline request if the current user is not a bot
// @param chatId Chat identifier
func (client *Client) GetChat(chatId int64) (*Chat, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "getChat",
		"chat_id": chatId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var chatDummy Chat
	err = json.Unmarshal(result.Raw, &chatDummy)
	return &chatDummy, err

}

// GetMessage Returns information about a message
// @param chatId Identifier of the chat the message belongs to
// @param messageId Identifier of the message to get
func (client *Client) GetMessage(chatId int64, messageId int64) (*Message, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":      "getMessage",
		"chat_id":    chatId,
		"message_id": messageId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var messageDummy Message
	err = json.Unmarshal(result.Raw, &messageDummy)
	return &messageDummy, err

}

// GetMessageLocally Returns information about a message, if it is available locally without sending network request. This is an offline request
// @param chatId Identifier of the chat the message belongs to
// @param messageId Identifier of the message to get
func (client *Client) GetMessageLocally(chatId int64, messageId int64) (*Message, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":      "getMessageLocally",
		"chat_id":    chatId,
		"message_id": messageId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var messageDummy Message
	err = json.Unmarshal(result.Raw, &messageDummy)
	return &messageDummy, err

}

// GetRepliedMessage Returns information about a message that is replied by a given message. Also returns the pinned message, the game message, and the invoice message for messages of the types messagePinMessage, messageGameScore, and messagePaymentSuccessful respectively
// @param chatId Identifier of the chat the message belongs to
// @param messageId Identifier of the message reply to which to get
func (client *Client) GetRepliedMessage(chatId int64, messageId int64) (*Message, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":      "getRepliedMessage",
		"chat_id":    chatId,
		"message_id": messageId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var messageDummy Message
	err = json.Unmarshal(result.Raw, &messageDummy)
	return &messageDummy, err

}

// GetChatPinnedMessage Returns information about a newest pinned message in the chat
// @param chatId Identifier of the chat the message belongs to
func (client *Client) GetChatPinnedMessage(chatId int64) (*Message, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "getChatPinnedMessage",
		"chat_id": chatId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var message Message
	err = json.Unmarshal(result.Raw, &message)
	return &message, err

}

// GetCallbackQueryMessage Returns information about a message with the callback button that originated a callback query; for bots only
// @param chatId Identifier of the chat the message belongs to
// @param messageId Message identifier
// @param callbackQueryId Identifier of the callback query
func (client *Client) GetCallbackQueryMessage(chatId int64, messageId int64, callbackQueryId JSONInt64) (*Message, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":             "getCallbackQueryMessage",
		"chat_id":           chatId,
		"message_id":        messageId,
		"callback_query_id": callbackQueryId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var messageDummy Message
	err = json.Unmarshal(result.Raw, &messageDummy)
	return &messageDummy, err

}

// GetMessages Returns information about messages. If a message is not found, returns null on the corresponding position of the result
// @param chatId Identifier of the chat the messages belong to
// @param messageIds Identifiers of the messages to get
func (client *Client) GetMessages(chatId int64, messageIds []int64) (*Messages, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":       "getMessages",
		"chat_id":     chatId,
		"message_ids": messageIds,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var messages Messages
	err = json.Unmarshal(result.Raw, &messages)
	return &messages, err

}

// GetMessageThread Returns information about a message thread. Can be used only if message.can_get_message_thread == true
// @param chatId Chat identifier
// @param messageId Identifier of the message
func (client *Client) GetMessageThread(chatId int64, messageId int64) (*MessageThreadInfo, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":      "getMessageThread",
		"chat_id":    chatId,
		"message_id": messageId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var messageThreadInfo MessageThreadInfo
	err = json.Unmarshal(result.Raw, &messageThreadInfo)
	return &messageThreadInfo, err

}

// GetMessageViewers Returns viewers of a recent outgoing message in a basic group or a supergroup chat. For video notes and voice notes only users, opened content of the message, are returned. The method can be called if message.can_get_viewers == true
// @param chatId Chat identifier
// @param messageId Identifier of the message
func (client *Client) GetMessageViewers(chatId int64, messageId int64) (*Users, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":      "getMessageViewers",
		"chat_id":    chatId,
		"message_id": messageId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var users Users
	err = json.Unmarshal(result.Raw, &users)
	return &users, err

}

// GetFile Returns information about a file; this is an offline request
// @param fileId Identifier of the file to get
func (client *Client) GetFile(fileId int32) (*File, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "getFile",
		"file_id": fileId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var fileDummy File
	err = json.Unmarshal(result.Raw, &fileDummy)
	return &fileDummy, err

}

// GetRemoteFile Returns information about a file by its remote ID; this is an offline request. Can be used to register a URL as a file for further uploading, or sending as a message. Even the request succeeds, the file can be used only if it is still accessible to the user.
// @param remoteFileId Remote identifier of the file to get
// @param fileType File type; pass null if unknown
func (client *Client) GetRemoteFile(remoteFileId string, fileType FileType) (*File, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":          "getRemoteFile",
		"remote_file_id": remoteFileId,
		"file_type":      fileType,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var fileDummy File
	err = json.Unmarshal(result.Raw, &fileDummy)
	return &fileDummy, err

}

// LoadChats Loads more chats from a chat list. The loaded chats and their positions in the chat list will be sent through updates. Chats are sorted by the pair (chat.position.order, chat.id) in descending order. Returns a 404 error if all chats has been loaded
// @param chatList The chat list in which to load chats; pass null to load chats from the main chat list
// @param limit The maximum number of chats to be loaded. For optimal performance, the number of loaded chats is chosen by TDLib and can be smaller than the specified limit, even if the end of the list is not reached
func (client *Client) LoadChats(chatList ChatList, limit int32) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":     "loadChats",
		"chat_list": chatList,
		"limit":     limit,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// GetChats Returns an ordered list of chats from the beginning of a chat list. For informational purposes only. Use loadChats and updates processing instead to maintain chat lists in a consistent state
// @param chatList The chat list in which to return chats; pass null to get chats from the main chat list
// @param limit The maximum number of chats to be returned
func (client *Client) GetChats(chatList ChatList, limit int32) (*Chats, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":     "getChats",
		"chat_list": chatList,
		"limit":     limit,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var chats Chats
	err = json.Unmarshal(result.Raw, &chats)
	return &chats, err

}

// SearchPublicChat Searches a public chat by its username. Currently only private chats, supergroups and channels can be public. Returns the chat if found; otherwise an error is returned
// @param username Username to be resolved
func (client *Client) SearchPublicChat(username string) (*Chat, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":    "searchPublicChat",
		"username": username,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var chat Chat
	err = json.Unmarshal(result.Raw, &chat)
	return &chat, err

}

// SearchPublicChats Searches public chats by looking for specified query in their username and title. Currently only private chats, supergroups and channels can be public. Returns a meaningful number of results.
// @param query
func (client *Client) SearchPublicChats(query string) (*Chats, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "searchPublicChats",
		"query": query,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var chats Chats
	err = json.Unmarshal(result.Raw, &chats)
	return &chats, err

}

// SearchChats Searches for the specified query in the title and username of already known chats, this is an offline request. Returns chats in the order seen in the main chat list
// @param query Query to search for. If the query is empty, returns up to 50 recently found chats
// @param limit The maximum number of chats to be returned
func (client *Client) SearchChats(query string, limit int32) (*Chats, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "searchChats",
		"query": query,
		"limit": limit,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var chats Chats
	err = json.Unmarshal(result.Raw, &chats)
	return &chats, err

}

// SearchChatsOnServer Searches for the specified query in the title and username of already known chats via request to the server. Returns chats in the order seen in the main chat list
// @param query Query to search for
// @param limit The maximum number of chats to be returned
func (client *Client) SearchChatsOnServer(query string, limit int32) (*Chats, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "searchChatsOnServer",
		"query": query,
		"limit": limit,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var chats Chats
	err = json.Unmarshal(result.Raw, &chats)
	return &chats, err

}

// SearchChatsNearby Returns a list of users and location-based supergroups nearby. The list of users nearby will be updated for 60 seconds after the request by the updates updateUsersNearby. The request must be sent again every 25 seconds with adjusted location to not miss new chats
// @param location Current user location
func (client *Client) SearchChatsNearby(location *Location) (*ChatsNearby, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":    "searchChatsNearby",
		"location": location,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var chatsNearby ChatsNearby
	err = json.Unmarshal(result.Raw, &chatsNearby)
	return &chatsNearby, err

}

// GetTopChats Returns a list of frequently used chats. Supported only if the chat info database is enabled
// @param category Category of chats to be returned
// @param limit The maximum number of chats to be returned; up to 30
func (client *Client) GetTopChats(category TopChatCategory, limit int32) (*Chats, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":    "getTopChats",
		"category": category,
		"limit":    limit,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var chats Chats
	err = json.Unmarshal(result.Raw, &chats)
	return &chats, err

}

// RemoveTopChat Removes a chat from the list of frequently used chats. Supported only if the chat info database is enabled
// @param category Category of frequently used chats
// @param chatId Chat identifier
func (client *Client) RemoveTopChat(category TopChatCategory, chatId int64) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":    "removeTopChat",
		"category": category,
		"chat_id":  chatId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// AddRecentlyFoundChat Adds a chat to the list of recently found chats. The chat is added to the beginning of the list. If the chat is already in the list, it will be removed from the list first
// @param chatId Identifier of the chat to add
func (client *Client) AddRecentlyFoundChat(chatId int64) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "addRecentlyFoundChat",
		"chat_id": chatId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// RemoveRecentlyFoundChat Removes a chat from the list of recently found chats
// @param chatId Identifier of the chat to be removed
func (client *Client) RemoveRecentlyFoundChat(chatId int64) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "removeRecentlyFoundChat",
		"chat_id": chatId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// ClearRecentlyFoundChats Clears the list of recently found chats
func (client *Client) ClearRecentlyFoundChats() (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "clearRecentlyFoundChats",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// GetRecentlyOpenedChats Returns recently opened chats, this is an offline request. Returns chats in the order of last opening
// @param limit The maximum number of chats to be returned
func (client *Client) GetRecentlyOpenedChats(limit int32) (*Chats, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "getRecentlyOpenedChats",
		"limit": limit,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var chats Chats
	err = json.Unmarshal(result.Raw, &chats)
	return &chats, err

}

// CheckChatUsername Checks whether a username can be set for a chat
// @param chatId Chat identifier; must be identifier of a supergroup chat, or a channel chat, or a private chat with self, or zero if the chat is being created
// @param username Username to be checked
func (client *Client) CheckChatUsername(chatId int64, username string) (CheckChatUsernameResult, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":    "checkChatUsername",
		"chat_id":  chatId,
		"username": username,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	switch CheckChatUsernameResultEnum(result.Data["@type"].(string)) {

	case CheckChatUsernameResultOkType:
		var checkChatUsernameResult CheckChatUsernameResultOk
		err = json.Unmarshal(result.Raw, &checkChatUsernameResult)
		return &checkChatUsernameResult, err

	case CheckChatUsernameResultUsernameInvalidType:
		var checkChatUsernameResult CheckChatUsernameResultUsernameInvalid
		err = json.Unmarshal(result.Raw, &checkChatUsernameResult)
		return &checkChatUsernameResult, err

	case CheckChatUsernameResultUsernameOccupiedType:
		var checkChatUsernameResult CheckChatUsernameResultUsernameOccupied
		err = json.Unmarshal(result.Raw, &checkChatUsernameResult)
		return &checkChatUsernameResult, err

	case CheckChatUsernameResultPublicChatsTooMuchType:
		var checkChatUsernameResult CheckChatUsernameResultPublicChatsTooMuch
		err = json.Unmarshal(result.Raw, &checkChatUsernameResult)
		return &checkChatUsernameResult, err

	case CheckChatUsernameResultPublicGroupsUnavailableType:
		var checkChatUsernameResult CheckChatUsernameResultPublicGroupsUnavailable
		err = json.Unmarshal(result.Raw, &checkChatUsernameResult)
		return &checkChatUsernameResult, err

	default:
		return nil, fmt.Errorf("Invalid type")
	}
}

// GetCreatedPublicChats Returns a list of public chats of the specified type, owned by the user
// @param typeParam Type of the public chats to return
func (client *Client) GetCreatedPublicChats(typeParam PublicChatType) (*Chats, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "getCreatedPublicChats",
		"type":  typeParam,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var chats Chats
	err = json.Unmarshal(result.Raw, &chats)
	return &chats, err

}

// CheckCreatedPublicChatsLimit Checks whether the maximum number of owned public chats has been reached. Returns corresponding error if the limit was reached
// @param typeParam Type of the public chats, for which to check the limit
func (client *Client) CheckCreatedPublicChatsLimit(typeParam PublicChatType) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "checkCreatedPublicChatsLimit",
		"type":  typeParam,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// GetSuitableDiscussionChats Returns a list of basic group and supergroup chats, which can be used as a discussion group for a channel. Returned basic group chats must be first upgraded to supergroups before they can be set as a discussion group. To set a returned supergroup as a discussion group, access to its old messages must be enabled using toggleSupergroupIsAllHistoryAvailable first
func (client *Client) GetSuitableDiscussionChats() (*Chats, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "getSuitableDiscussionChats",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var chats Chats
	err = json.Unmarshal(result.Raw, &chats)
	return &chats, err

}

// GetInactiveSupergroupChats Returns a list of recently inactive supergroups and channels. Can be used when user reaches limit on the number of joined supergroups and channels and receives CHANNELS_TOO_MUCH error
func (client *Client) GetInactiveSupergroupChats() (*Chats, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "getInactiveSupergroupChats",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var chats Chats
	err = json.Unmarshal(result.Raw, &chats)
	return &chats, err

}

// GetGroupsInCommon Returns a list of common group chats with a given user. Chats are sorted by their type and creation date
// @param userId User identifier
// @param offsetChatId Chat identifier starting from which to return chats; use 0 for the first request
// @param limit The maximum number of chats to be returned; up to 100
func (client *Client) GetGroupsInCommon(userId int64, offsetChatId int64, limit int32) (*Chats, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":          "getGroupsInCommon",
		"user_id":        userId,
		"offset_chat_id": offsetChatId,
		"limit":          limit,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var chats Chats
	err = json.Unmarshal(result.Raw, &chats)
	return &chats, err

}

// GetChatHistory Returns messages in a chat. The messages are returned in a reverse chronological order (i.e., in order of decreasing message_id).
// @param chatId Chat identifier
// @param fromMessageId Identifier of the message starting from which history must be fetched; use 0 to get results from the last message
// @param offset Specify 0 to get results from exactly the from_message_id or a negative offset up to 99 to get additionally some newer messages
// @param limit The maximum number of messages to be returned; must be positive and can't be greater than 100. If the offset is negative, the limit must be greater than or equal to -offset. For optimal performance, the number of returned messages is chosen by TDLib and can be smaller than the specified limit
// @param onlyLocal If true, returns only messages that are available locally without sending network requests
func (client *Client) GetChatHistory(chatId int64, fromMessageId int64, offset int32, limit int32, onlyLocal bool) (*Messages, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":           "getChatHistory",
		"chat_id":         chatId,
		"from_message_id": fromMessageId,
		"offset":          offset,
		"limit":           limit,
		"only_local":      onlyLocal,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var messages Messages
	err = json.Unmarshal(result.Raw, &messages)
	return &messages, err

}

// GetMessageThreadHistory Returns messages in a message thread of a message. Can be used only if message.can_get_message_thread == true. Message thread of a channel message is in the channel's linked supergroup.
// @param chatId Chat identifier
// @param messageId Message identifier, which thread history needs to be returned
// @param fromMessageId Identifier of the message starting from which history must be fetched; use 0 to get results from the last message
// @param offset Specify 0 to get results from exactly the from_message_id or a negative offset up to 99 to get additionally some newer messages
// @param limit The maximum number of messages to be returned; must be positive and can't be greater than 100. If the offset is negative, the limit must be greater than or equal to -offset. For optimal performance, the number of returned messages is chosen by TDLib and can be smaller than the specified limit
func (client *Client) GetMessageThreadHistory(chatId int64, messageId int64, fromMessageId int64, offset int32, limit int32) (*Messages, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":           "getMessageThreadHistory",
		"chat_id":         chatId,
		"message_id":      messageId,
		"from_message_id": fromMessageId,
		"offset":          offset,
		"limit":           limit,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var messages Messages
	err = json.Unmarshal(result.Raw, &messages)
	return &messages, err

}

// DeleteChatHistory Deletes all messages in the chat. Use chat.can_be_deleted_only_for_self and chat.can_be_deleted_for_all_users fields to find whether and how the method can be applied to the chat
// @param chatId Chat identifier
// @param removeFromChatList Pass true if the chat needs to be removed from the chat list
// @param revoke Pass true to try to delete chat history for all users
func (client *Client) DeleteChatHistory(chatId int64, removeFromChatList bool, revoke bool) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":                 "deleteChatHistory",
		"chat_id":               chatId,
		"remove_from_chat_list": removeFromChatList,
		"revoke":                revoke,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var okDummy Ok
	err = json.Unmarshal(result.Raw, &okDummy)
	return &okDummy, err

}

// DeleteChat Deletes a chat along with all messages in the corresponding chat for all chat members; requires owner privileges. For group chats this will release the username and remove all members. Chats with more than 1000 members can't be deleted using this method
// @param chatId Chat identifier
func (client *Client) DeleteChat(chatId int64) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "deleteChat",
		"chat_id": chatId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// SearchChatMessages Searches for messages with given words in the chat. Returns the results in reverse chronological order, i.e. in order of decreasing message_id. Cannot be used in secret chats with a non-empty query
// @param chatId Identifier of the chat in which to search messages
// @param query Query to search for
// @param sender Sender of messages to search for; pass null to search for messages from any sender. Not supported in secret chats
// @param fromMessageId Identifier of the message starting from which history must be fetched; use 0 to get results from the last message
// @param offset Specify 0 to get results from exactly the from_message_id or a negative offset to get the specified message and some newer messages
// @param limit The maximum number of messages to be returned; must be positive and can't be greater than 100. If the offset is negative, the limit must be greater than -offset. For optimal performance, the number of returned messages is chosen by TDLib and can be smaller than the specified limit
// @param filter Additional filter for messages to search; pass null to search for all messages
// @param messageThreadId If not 0, only messages in the specified thread will be returned; supergroups only
func (client *Client) SearchChatMessages(chatId int64, query string, sender MessageSender, fromMessageId int64, offset int32, limit int32, filter SearchMessagesFilter, messageThreadId int64) (*Messages, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":             "searchChatMessages",
		"chat_id":           chatId,
		"query":             query,
		"sender":            sender,
		"from_message_id":   fromMessageId,
		"offset":            offset,
		"limit":             limit,
		"filter":            filter,
		"message_thread_id": messageThreadId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var messages Messages
	err = json.Unmarshal(result.Raw, &messages)
	return &messages, err

}

// SearchMessages Searches for messages in all chats except secret chats. Returns the results in reverse chronological order (i.e., in order of decreasing (date, chat_id, message_id)).
// @param chatList Chat list in which to search messages; pass null to search in all chats regardless of their chat list. Only Main and Archive chat lists are supported
// @param query Query to search for
// @param offsetDate The date of the message starting from which the results need to be fetched. Use 0 or any date in the future to get results from the last message
// @param offsetChatId The chat identifier of the last found message, or 0 for the first request
// @param offsetMessageId The message identifier of the last found message, or 0 for the first request
// @param limit The maximum number of messages to be returned; up to 100. For optimal performance, the number of returned messages is chosen by TDLib and can be smaller than the specified limit
// @param filter Additional filter for messages to search; pass null to search for all messages. Filters searchMessagesFilterCall, searchMessagesFilterMissedCall, searchMessagesFilterMention, searchMessagesFilterUnreadMention, searchMessagesFilterFailedToSend and searchMessagesFilterPinned are unsupported in this function
// @param minDate If not 0, the minimum date of the messages to return
// @param maxDate If not 0, the maximum date of the messages to return
func (client *Client) SearchMessages(chatList ChatList, query string, offsetDate int32, offsetChatId int64, offsetMessageId int64, limit int32, filter SearchMessagesFilter, minDate int32, maxDate int32) (*Messages, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":             "searchMessages",
		"chat_list":         chatList,
		"query":             query,
		"offset_date":       offsetDate,
		"offset_chat_id":    offsetChatId,
		"offset_message_id": offsetMessageId,
		"limit":             limit,
		"filter":            filter,
		"min_date":          minDate,
		"max_date":          maxDate,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var messages Messages
	err = json.Unmarshal(result.Raw, &messages)
	return &messages, err

}

// SearchSecretMessages Searches for messages in secret chats. Returns the results in reverse chronological order. For optimal performance, the number of returned messages is chosen by TDLib
// @param chatId Identifier of the chat in which to search. Specify 0 to search in all secret chats
// @param query Query to search for. If empty, searchChatMessages must be used instead
// @param offset Offset of the first entry to return as received from the previous request; use empty string to get first chunk of results
// @param limit The maximum number of messages to be returned; up to 100. For optimal performance, the number of returned messages is chosen by TDLib and can be smaller than the specified limit
// @param filter Additional filter for messages to search; pass null to search for all messages
func (client *Client) SearchSecretMessages(chatId int64, query string, offset string, limit int32, filter SearchMessagesFilter) (*FoundMessages, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "searchSecretMessages",
		"chat_id": chatId,
		"query":   query,
		"offset":  offset,
		"limit":   limit,
		"filter":  filter,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var foundMessages FoundMessages
	err = json.Unmarshal(result.Raw, &foundMessages)
	return &foundMessages, err

}

// SearchCallMessages Searches for call messages. Returns the results in reverse chronological order (i. e., in order of decreasing message_id). For optimal performance, the number of returned messages is chosen by TDLib
// @param fromMessageId Identifier of the message from which to search; use 0 to get results from the last message
// @param limit The maximum number of messages to be returned; up to 100. For optimal performance, the number of returned messages is chosen by TDLib and can be smaller than the specified limit
// @param onlyMissed If true, returns only messages with missed calls
func (client *Client) SearchCallMessages(fromMessageId int64, limit int32, onlyMissed bool) (*Messages, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":           "searchCallMessages",
		"from_message_id": fromMessageId,
		"limit":           limit,
		"only_missed":     onlyMissed,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var messages Messages
	err = json.Unmarshal(result.Raw, &messages)
	return &messages, err

}

// DeleteAllCallMessages Deletes all call messages
// @param revoke Pass true to delete the messages for all users
func (client *Client) DeleteAllCallMessages(revoke bool) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":  "deleteAllCallMessages",
		"revoke": revoke,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var okDummy Ok
	err = json.Unmarshal(result.Raw, &okDummy)
	return &okDummy, err

}

// SearchChatRecentLocationMessages Returns information about the recent locations of chat members that were sent to the chat. Returns up to 1 location message per user
// @param chatId Chat identifier
// @param limit The maximum number of messages to be returned
func (client *Client) SearchChatRecentLocationMessages(chatId int64, limit int32) (*Messages, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "searchChatRecentLocationMessages",
		"chat_id": chatId,
		"limit":   limit,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var messages Messages
	err = json.Unmarshal(result.Raw, &messages)
	return &messages, err

}

// GetActiveLiveLocationMessages Returns all active live locations that need to be updated by the application. The list is persistent across application restarts only if the message database is used
func (client *Client) GetActiveLiveLocationMessages() (*Messages, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "getActiveLiveLocationMessages",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var messages Messages
	err = json.Unmarshal(result.Raw, &messages)
	return &messages, err

}

// GetChatMessageByDate Returns the last message sent in a chat no later than the specified date
// @param chatId Chat identifier
// @param date Point in time (Unix timestamp) relative to which to search for messages
func (client *Client) GetChatMessageByDate(chatId int64, date int32) (*Message, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "getChatMessageByDate",
		"chat_id": chatId,
		"date":    date,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var message Message
	err = json.Unmarshal(result.Raw, &message)
	return &message, err

}

// GetChatSparseMessagePositions Returns sparse positions of messages of the specified type in the chat to be used for shared media scroll implementation. Returns the results in reverse chronological order (i.e., in order of decreasing message_id).
// @param chatId Identifier of the chat in which to return information about message positions
// @param filter Filter for message content. Filters searchMessagesFilterEmpty, searchMessagesFilterCall, searchMessagesFilterMissedCall, searchMessagesFilterMention and searchMessagesFilterUnreadMention are unsupported in this function
// @param fromMessageId The message identifier from which to return information about message positions
// @param limit The expected number of message positions to be returned; 50-2000. A smaller number of positions can be returned, if there are not enough appropriate messages
func (client *Client) GetChatSparseMessagePositions(chatId int64, filter SearchMessagesFilter, fromMessageId int64, limit int32) (*MessagePositions, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":           "getChatSparseMessagePositions",
		"chat_id":         chatId,
		"filter":          filter,
		"from_message_id": fromMessageId,
		"limit":           limit,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var messagePositions MessagePositions
	err = json.Unmarshal(result.Raw, &messagePositions)
	return &messagePositions, err

}

// GetChatMessageCalendar Returns information about the next messages of the specified type in the chat splitted by days. Returns the results in reverse chronological order. Can return partial result for the last returned day. Behavior of this method depends on the value of the option "utc_time_offset"
// @param chatId Identifier of the chat in which to return information about messages
// @param filter Filter for message content. Filters searchMessagesFilterEmpty, searchMessagesFilterCall, searchMessagesFilterMissedCall, searchMessagesFilterMention and searchMessagesFilterUnreadMention are unsupported in this function
// @param fromMessageId The message identifier from which to return information about messages; use 0 to get results from the last message
func (client *Client) GetChatMessageCalendar(chatId int64, filter SearchMessagesFilter, fromMessageId int64) (*MessageCalendar, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":           "getChatMessageCalendar",
		"chat_id":         chatId,
		"filter":          filter,
		"from_message_id": fromMessageId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var messageCalendar MessageCalendar
	err = json.Unmarshal(result.Raw, &messageCalendar)
	return &messageCalendar, err

}

// GetChatMessageCount Returns approximate number of messages of the specified type in the chat
// @param chatId Identifier of the chat in which to count messages
// @param filter Filter for message content; searchMessagesFilterEmpty is unsupported in this function
// @param returnLocal If true, returns count that is available locally without sending network requests, returning -1 if the number of messages is unknown
func (client *Client) GetChatMessageCount(chatId int64, filter SearchMessagesFilter, returnLocal bool) (*Count, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":        "getChatMessageCount",
		"chat_id":      chatId,
		"filter":       filter,
		"return_local": returnLocal,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var count Count
	err = json.Unmarshal(result.Raw, &count)
	return &count, err

}

// GetChatScheduledMessages Returns all scheduled messages in a chat. The messages are returned in a reverse chronological order (i.e., in order of decreasing message_id)
// @param chatId Chat identifier
func (client *Client) GetChatScheduledMessages(chatId int64) (*Messages, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "getChatScheduledMessages",
		"chat_id": chatId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var messages Messages
	err = json.Unmarshal(result.Raw, &messages)
	return &messages, err

}

// GetMessagePublicForwards Returns forwarded copies of a channel message to different public channels. For optimal performance, the number of returned messages is chosen by TDLib
// @param chatId Chat identifier of the message
// @param messageId Message identifier
// @param offset Offset of the first entry to return as received from the previous request; use empty string to get first chunk of results
// @param limit The maximum number of messages to be returned; must be positive and can't be greater than 100. For optimal performance, the number of returned messages is chosen by TDLib and can be smaller than the specified limit
func (client *Client) GetMessagePublicForwards(chatId int64, messageId int64, offset string, limit int32) (*FoundMessages, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":      "getMessagePublicForwards",
		"chat_id":    chatId,
		"message_id": messageId,
		"offset":     offset,
		"limit":      limit,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var foundMessages FoundMessages
	err = json.Unmarshal(result.Raw, &foundMessages)
	return &foundMessages, err

}

// GetChatSponsoredMessages Returns sponsored messages to be shown in a chat; for channel chats only
// @param chatId Identifier of the chat
func (client *Client) GetChatSponsoredMessages(chatId int64) (*SponsoredMessages, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "getChatSponsoredMessages",
		"chat_id": chatId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var sponsoredMessages SponsoredMessages
	err = json.Unmarshal(result.Raw, &sponsoredMessages)
	return &sponsoredMessages, err

}

// ViewSponsoredMessage Informs TDLib that a sponsored message was viewed by the user
// @param chatId Identifier of the chat with the sponsored message
// @param sponsoredMessageId The identifier of the sponsored message being viewed
func (client *Client) ViewSponsoredMessage(chatId int64, sponsoredMessageId int32) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":                "viewSponsoredMessage",
		"chat_id":              chatId,
		"sponsored_message_id": sponsoredMessageId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// RemoveNotification Removes an active notification from notification list. Needs to be called only if the notification is removed by the current user
// @param notificationGroupId Identifier of notification group to which the notification belongs
// @param notificationId Identifier of removed notification
func (client *Client) RemoveNotification(notificationGroupId int32, notificationId int32) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":                 "removeNotification",
		"notification_group_id": notificationGroupId,
		"notification_id":       notificationId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// RemoveNotificationGroup Removes a group of active notifications. Needs to be called only if the notification group is removed by the current user
// @param notificationGroupId Notification group identifier
// @param maxNotificationId The maximum identifier of removed notifications
func (client *Client) RemoveNotificationGroup(notificationGroupId int32, maxNotificationId int32) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":                 "removeNotificationGroup",
		"notification_group_id": notificationGroupId,
		"max_notification_id":   maxNotificationId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// GetMessageLink Returns an HTTPS link to a message in a chat. Available only for already sent messages in supergroups and channels, or if message.can_get_media_timestamp_links and a media timestamp link is generated. This is an offline request
// @param chatId Identifier of the chat to which the message belongs
// @param messageId Identifier of the message
// @param mediaTimestamp If not 0, timestamp from which the video/audio/video note/voice note playing must start, in seconds. The media can be in the message content or in its web page preview
// @param forAlbum Pass true to create a link for the whole media album
// @param forComment Pass true to create a link to the message as a channel post comment, or from a message thread
func (client *Client) GetMessageLink(chatId int64, messageId int64, mediaTimestamp int32, forAlbum bool, forComment bool) (*MessageLink, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":           "getMessageLink",
		"chat_id":         chatId,
		"message_id":      messageId,
		"media_timestamp": mediaTimestamp,
		"for_album":       forAlbum,
		"for_comment":     forComment,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var messageLink MessageLink
	err = json.Unmarshal(result.Raw, &messageLink)
	return &messageLink, err

}

// GetMessageEmbeddingCode Returns an HTML code for embedding the message. Available only for messages in supergroups and channels with a username
// @param chatId Identifier of the chat to which the message belongs
// @param messageId Identifier of the message
// @param forAlbum Pass true to return an HTML code for embedding of the whole media album
func (client *Client) GetMessageEmbeddingCode(chatId int64, messageId int64, forAlbum bool) (*Text, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":      "getMessageEmbeddingCode",
		"chat_id":    chatId,
		"message_id": messageId,
		"for_album":  forAlbum,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var text Text
	err = json.Unmarshal(result.Raw, &text)
	return &text, err

}

// GetMessageLinkInfo Returns information about a public or private message link. Can be called for any internal link of the type internalLinkTypeMessage
// @param url The message link
func (client *Client) GetMessageLinkInfo(url string) (*MessageLinkInfo, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "getMessageLinkInfo",
		"url":   url,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var messageLinkInfo MessageLinkInfo
	err = json.Unmarshal(result.Raw, &messageLinkInfo)
	return &messageLinkInfo, err

}

// SendMessage Sends a message. Returns the sent message
// @param chatId Target chat
// @param messageThreadId If not 0, a message thread identifier in which the message will be sent
// @param replyToMessageId Identifier of the message to reply to or 0
// @param options Options to be used to send the message; pass null to use default options
// @param replyMarkup Markup for replying to the message; pass null if none; for bots only
// @param inputMessageContent The content of the message to be sent
func (client *Client) SendMessage(chatId int64, messageThreadId int64, replyToMessageId int64, options *MessageSendOptions, replyMarkup ReplyMarkup, inputMessageContent InputMessageContent) (*Message, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":                 "sendMessage",
		"chat_id":               chatId,
		"message_thread_id":     messageThreadId,
		"reply_to_message_id":   replyToMessageId,
		"options":               options,
		"reply_markup":          replyMarkup,
		"input_message_content": inputMessageContent,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var messageDummy Message
	err = json.Unmarshal(result.Raw, &messageDummy)
	return &messageDummy, err

}

// SendMessageAlbum Sends 2-10 messages grouped together into an album. Currently only audio, document, photo and video messages can be grouped into an album. Documents and audio files can be only grouped in an album with messages of the same type. Returns sent messages
// @param chatId Target chat
// @param messageThreadId If not 0, a message thread identifier in which the messages will be sent
// @param replyToMessageId Identifier of a message to reply to or 0
// @param options Options to be used to send the messages; pass null to use default options
// @param inputMessageContents Contents of messages to be sent. At most 10 messages can be added to an album
func (client *Client) SendMessageAlbum(chatId int64, messageThreadId int64, replyToMessageId int64, options *MessageSendOptions, inputMessageContents []InputMessageContent) (*Messages, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":                  "sendMessageAlbum",
		"chat_id":                chatId,
		"message_thread_id":      messageThreadId,
		"reply_to_message_id":    replyToMessageId,
		"options":                options,
		"input_message_contents": inputMessageContents,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var messages Messages
	err = json.Unmarshal(result.Raw, &messages)
	return &messages, err

}

// SendBotStartMessage Invites a bot to a chat (if it is not yet a member) and sends it the /start command. Bots can't be invited to a private chat other than the chat with the bot. Bots can't be invited to channels (although they can be added as admins) and secret chats. Returns the sent message
// @param botUserId Identifier of the bot
// @param chatId Identifier of the target chat
// @param parameter A hidden parameter sent to the bot for deep linking purposes (https://core.telegram.org/bots#deep-linking)
func (client *Client) SendBotStartMessage(botUserId int64, chatId int64, parameter string) (*Message, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":       "sendBotStartMessage",
		"bot_user_id": botUserId,
		"chat_id":     chatId,
		"parameter":   parameter,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var message Message
	err = json.Unmarshal(result.Raw, &message)
	return &message, err

}

// SendInlineQueryResultMessage Sends the result of an inline query as a message. Returns the sent message. Always clears a chat draft message
// @param chatId Target chat
// @param messageThreadId If not 0, a message thread identifier in which the message will be sent
// @param replyToMessageId Identifier of a message to reply to or 0
// @param options Options to be used to send the message; pass null to use default options
// @param queryId Identifier of the inline query
// @param resultId Identifier of the inline result
// @param hideViaBot If true, there will be no mention of a bot, via which the message is sent. Can be used only for bots GetOption("animation_search_bot_username"), GetOption("photo_search_bot_username") and GetOption("venue_search_bot_username")
func (client *Client) SendInlineQueryResultMessage(chatId int64, messageThreadId int64, replyToMessageId int64, options *MessageSendOptions, queryId JSONInt64, resultId string, hideViaBot bool) (*Message, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":               "sendInlineQueryResultMessage",
		"chat_id":             chatId,
		"message_thread_id":   messageThreadId,
		"reply_to_message_id": replyToMessageId,
		"options":             options,
		"query_id":            queryId,
		"result_id":           resultId,
		"hide_via_bot":        hideViaBot,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var messageDummy Message
	err = json.Unmarshal(result.Raw, &messageDummy)
	return &messageDummy, err

}

// ForwardMessages Forwards previously sent messages. Returns the forwarded messages in the same order as the message identifiers passed in message_ids. If a message can't be forwarded, null will be returned instead of the message
// @param chatId Identifier of the chat to which to forward messages
// @param fromChatId Identifier of the chat from which to forward messages
// @param messageIds Identifiers of the messages to forward. Message identifiers must be in a strictly increasing order. At most 100 messages can be forwarded simultaneously
// @param options Options to be used to send the messages; pass null to use default options
// @param sendCopy If true, content of the messages will be copied without reference to the original sender. Always true if the messages are forwarded to a secret chat or are local
// @param removeCaption If true, media caption of message copies will be removed. Ignored if send_copy is false
// @param onlyPreview If true, messages will not be forwarded and instead fake messages will be returned
func (client *Client) ForwardMessages(chatId int64, fromChatId int64, messageIds []int64, options *MessageSendOptions, sendCopy bool, removeCaption bool, onlyPreview bool) (*Messages, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":          "forwardMessages",
		"chat_id":        chatId,
		"from_chat_id":   fromChatId,
		"message_ids":    messageIds,
		"options":        options,
		"send_copy":      sendCopy,
		"remove_caption": removeCaption,
		"only_preview":   onlyPreview,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var messages Messages
	err = json.Unmarshal(result.Raw, &messages)
	return &messages, err

}

// ResendMessages Resends messages which failed to send. Can be called only for messages for which messageSendingStateFailed.can_retry is true and after specified in messageSendingStateFailed.retry_after time passed.
// @param chatId Identifier of the chat to send messages
// @param messageIds Identifiers of the messages to resend. Message identifiers must be in a strictly increasing order
func (client *Client) ResendMessages(chatId int64, messageIds []int64) (*Messages, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":       "resendMessages",
		"chat_id":     chatId,
		"message_ids": messageIds,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var messages Messages
	err = json.Unmarshal(result.Raw, &messages)
	return &messages, err

}

// SendChatScreenshotTakenNotification Sends a notification about a screenshot taken in a chat. Supported only in private and secret chats
// @param chatId Chat identifier
func (client *Client) SendChatScreenshotTakenNotification(chatId int64) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "sendChatScreenshotTakenNotification",
		"chat_id": chatId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// AddLocalMessage Adds a local message to a chat. The message is persistent across application restarts only if the message database is used. Returns the added message
// @param chatId Target chat
// @param sender The sender of the message
// @param replyToMessageId Identifier of the message to reply to or 0
// @param disableNotification Pass true to disable notification for the message
// @param inputMessageContent The content of the message to be added
func (client *Client) AddLocalMessage(chatId int64, sender MessageSender, replyToMessageId int64, disableNotification bool, inputMessageContent InputMessageContent) (*Message, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":                 "addLocalMessage",
		"chat_id":               chatId,
		"sender":                sender,
		"reply_to_message_id":   replyToMessageId,
		"disable_notification":  disableNotification,
		"input_message_content": inputMessageContent,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var messageDummy Message
	err = json.Unmarshal(result.Raw, &messageDummy)
	return &messageDummy, err

}

// DeleteMessages Deletes messages
// @param chatId Chat identifier
// @param messageIds Identifiers of the messages to be deleted
// @param revoke Pass true to try to delete messages for all chat members. Always true for supergroups, channels and secret chats
func (client *Client) DeleteMessages(chatId int64, messageIds []int64, revoke bool) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":       "deleteMessages",
		"chat_id":     chatId,
		"message_ids": messageIds,
		"revoke":      revoke,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var okDummy Ok
	err = json.Unmarshal(result.Raw, &okDummy)
	return &okDummy, err

}

// DeleteChatMessagesFromUser Deletes all messages sent by the specified user to a chat. Supported only for supergroups; requires can_delete_messages administrator privileges
// @param chatId Chat identifier
// @param userId User identifier
func (client *Client) DeleteChatMessagesFromUser(chatId int64, userId int64) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "deleteChatMessagesFromUser",
		"chat_id": chatId,
		"user_id": userId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// DeleteChatMessagesByDate Deletes all messages between the specified dates in a chat. Supported only for private chats and basic groups. Messages sent in the last 30 seconds will not be deleted
// @param chatId Chat identifier
// @param minDate The minimum date of the messages to delete
// @param maxDate The maximum date of the messages to delete
// @param revoke Pass true to try to delete chat messages for all users; private chats only
func (client *Client) DeleteChatMessagesByDate(chatId int64, minDate int32, maxDate int32, revoke bool) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":    "deleteChatMessagesByDate",
		"chat_id":  chatId,
		"min_date": minDate,
		"max_date": maxDate,
		"revoke":   revoke,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var okDummy Ok
	err = json.Unmarshal(result.Raw, &okDummy)
	return &okDummy, err

}

// EditMessageText Edits the text of a message (or a text of a game message). Returns the edited message after the edit is completed on the server side
// @param chatId The chat the message belongs to
// @param messageId Identifier of the message
// @param replyMarkup The new message reply markup; pass null if none; for bots only
// @param inputMessageContent New text content of the message. Must be of type inputMessageText
func (client *Client) EditMessageText(chatId int64, messageId int64, replyMarkup ReplyMarkup, inputMessageContent InputMessageContent) (*Message, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":                 "editMessageText",
		"chat_id":               chatId,
		"message_id":            messageId,
		"reply_markup":          replyMarkup,
		"input_message_content": inputMessageContent,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var messageDummy Message
	err = json.Unmarshal(result.Raw, &messageDummy)
	return &messageDummy, err

}

// EditMessageLiveLocation Edits the message content of a live location. Messages can be edited for a limited period of time specified in the live location. Returns the edited message after the edit is completed on the server side
// @param chatId The chat the message belongs to
// @param messageId Identifier of the message
// @param replyMarkup The new message reply markup; pass null if none; for bots only
// @param location New location content of the message; pass null to stop sharing the live location
// @param heading The new direction in which the location moves, in degrees; 1-360. Pass 0 if unknown
// @param proximityAlertRadius The new maximum distance for proximity alerts, in meters (0-100000). Pass 0 if the notification is disabled
func (client *Client) EditMessageLiveLocation(chatId int64, messageId int64, replyMarkup ReplyMarkup, location *Location, heading int32, proximityAlertRadius int32) (*Message, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":                  "editMessageLiveLocation",
		"chat_id":                chatId,
		"message_id":             messageId,
		"reply_markup":           replyMarkup,
		"location":               location,
		"heading":                heading,
		"proximity_alert_radius": proximityAlertRadius,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var messageDummy Message
	err = json.Unmarshal(result.Raw, &messageDummy)
	return &messageDummy, err

}

// EditMessageMedia Edits the content of a message with an animation, an audio, a document, a photo or a video, including message caption. If only the caption needs to be edited, use editMessageCaption instead.
// @param chatId The chat the message belongs to
// @param messageId Identifier of the message
// @param replyMarkup The new message reply markup; pass null if none; for bots only
// @param inputMessageContent New content of the message. Must be one of the following types: inputMessageAnimation, inputMessageAudio, inputMessageDocument, inputMessagePhoto or inputMessageVideo
func (client *Client) EditMessageMedia(chatId int64, messageId int64, replyMarkup ReplyMarkup, inputMessageContent InputMessageContent) (*Message, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":                 "editMessageMedia",
		"chat_id":               chatId,
		"message_id":            messageId,
		"reply_markup":          replyMarkup,
		"input_message_content": inputMessageContent,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var messageDummy Message
	err = json.Unmarshal(result.Raw, &messageDummy)
	return &messageDummy, err

}

// EditMessageCaption Edits the message content caption. Returns the edited message after the edit is completed on the server side
// @param chatId The chat the message belongs to
// @param messageId Identifier of the message
// @param replyMarkup The new message reply markup; pass null if none; for bots only
// @param caption New message content caption; 0-GetOption("message_caption_length_max") characters; pass null to remove caption
func (client *Client) EditMessageCaption(chatId int64, messageId int64, replyMarkup ReplyMarkup, caption *FormattedText) (*Message, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":        "editMessageCaption",
		"chat_id":      chatId,
		"message_id":   messageId,
		"reply_markup": replyMarkup,
		"caption":      caption,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var messageDummy Message
	err = json.Unmarshal(result.Raw, &messageDummy)
	return &messageDummy, err

}

// EditMessageReplyMarkup Edits the message reply markup; for bots only. Returns the edited message after the edit is completed on the server side
// @param chatId The chat the message belongs to
// @param messageId Identifier of the message
// @param replyMarkup The new message reply markup; pass null if none
func (client *Client) EditMessageReplyMarkup(chatId int64, messageId int64, replyMarkup ReplyMarkup) (*Message, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":        "editMessageReplyMarkup",
		"chat_id":      chatId,
		"message_id":   messageId,
		"reply_markup": replyMarkup,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var messageDummy Message
	err = json.Unmarshal(result.Raw, &messageDummy)
	return &messageDummy, err

}

// EditInlineMessageText Edits the text of an inline text or game message sent via a bot; for bots only
// @param inlineMessageId Inline message identifier
// @param replyMarkup The new message reply markup; pass null if none
// @param inputMessageContent New text content of the message. Must be of type inputMessageText
func (client *Client) EditInlineMessageText(inlineMessageId string, replyMarkup ReplyMarkup, inputMessageContent InputMessageContent) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":                 "editInlineMessageText",
		"inline_message_id":     inlineMessageId,
		"reply_markup":          replyMarkup,
		"input_message_content": inputMessageContent,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// EditInlineMessageLiveLocation Edits the content of a live location in an inline message sent via a bot; for bots only
// @param inlineMessageId Inline message identifier
// @param replyMarkup The new message reply markup; pass null if none
// @param location New location content of the message; pass null to stop sharing the live location
// @param heading The new direction in which the location moves, in degrees; 1-360. Pass 0 if unknown
// @param proximityAlertRadius The new maximum distance for proximity alerts, in meters (0-100000). Pass 0 if the notification is disabled
func (client *Client) EditInlineMessageLiveLocation(inlineMessageId string, replyMarkup ReplyMarkup, location *Location, heading int32, proximityAlertRadius int32) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":                  "editInlineMessageLiveLocation",
		"inline_message_id":      inlineMessageId,
		"reply_markup":           replyMarkup,
		"location":               location,
		"heading":                heading,
		"proximity_alert_radius": proximityAlertRadius,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// EditInlineMessageMedia Edits the content of a message with an animation, an audio, a document, a photo or a video in an inline message sent via a bot; for bots only
// @param inlineMessageId Inline message identifier
// @param replyMarkup The new message reply markup; pass null if none; for bots only
// @param inputMessageContent New content of the message. Must be one of the following types: inputMessageAnimation, inputMessageAudio, inputMessageDocument, inputMessagePhoto or inputMessageVideo
func (client *Client) EditInlineMessageMedia(inlineMessageId string, replyMarkup ReplyMarkup, inputMessageContent InputMessageContent) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":                 "editInlineMessageMedia",
		"inline_message_id":     inlineMessageId,
		"reply_markup":          replyMarkup,
		"input_message_content": inputMessageContent,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// EditInlineMessageCaption Edits the caption of an inline message sent via a bot; for bots only
// @param inlineMessageId Inline message identifier
// @param replyMarkup The new message reply markup; pass null if none
// @param caption New message content caption; pass null to remove caption; 0-GetOption("message_caption_length_max") characters
func (client *Client) EditInlineMessageCaption(inlineMessageId string, replyMarkup ReplyMarkup, caption *FormattedText) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":             "editInlineMessageCaption",
		"inline_message_id": inlineMessageId,
		"reply_markup":      replyMarkup,
		"caption":           caption,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// EditInlineMessageReplyMarkup Edits the reply markup of an inline message sent via a bot; for bots only
// @param inlineMessageId Inline message identifier
// @param replyMarkup The new message reply markup; pass null if none
func (client *Client) EditInlineMessageReplyMarkup(inlineMessageId string, replyMarkup ReplyMarkup) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":             "editInlineMessageReplyMarkup",
		"inline_message_id": inlineMessageId,
		"reply_markup":      replyMarkup,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// EditMessageSchedulingState Edits the time when a scheduled message will be sent. Scheduling state of all messages in the same album or forwarded together with the message will be also changed
// @param chatId The chat the message belongs to
// @param messageId Identifier of the message
// @param schedulingState The new message scheduling state; pass null to send the message immediately
func (client *Client) EditMessageSchedulingState(chatId int64, messageId int64, schedulingState MessageSchedulingState) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":            "editMessageSchedulingState",
		"chat_id":          chatId,
		"message_id":       messageId,
		"scheduling_state": schedulingState,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// GetTextEntities Returns all entities (mentions, hashtags, cashtags, bot commands, bank card numbers, URLs, and email addresses) contained in the text. Can be called synchronously
// @param text The text in which to look for entites
func (client *Client) GetTextEntities(text string) (*TextEntities, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "getTextEntities",
		"text":  text,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var textEntities TextEntities
	err = json.Unmarshal(result.Raw, &textEntities)
	return &textEntities, err

}

// ParseTextEntities Parses Bold, Italic, Underline, Strikethrough, Code, Pre, PreCode, TextUrl and MentionName entities contained in the text. Can be called synchronously
// @param text The text to parse
// @param parseMode Text parse mode
func (client *Client) ParseTextEntities(text string, parseMode TextParseMode) (*FormattedText, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":      "parseTextEntities",
		"text":       text,
		"parse_mode": parseMode,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var formattedText FormattedText
	err = json.Unmarshal(result.Raw, &formattedText)
	return &formattedText, err

}

// ParseMarkdown Parses Markdown entities in a human-friendly format, ignoring markup errors. Can be called synchronously
// @param text The text to parse. For example, "__italic__ ~~strikethrough~~ **bold** `code` ```pre``` __[italic__ text_url](telegram.org) __italic**bold italic__bold**"
func (client *Client) ParseMarkdown(text *FormattedText) (*FormattedText, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "parseMarkdown",
		"text":  text,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var formattedText FormattedText
	err = json.Unmarshal(result.Raw, &formattedText)
	return &formattedText, err

}

// GetMarkdownText Replaces text entities with Markdown formatting in a human-friendly format. Entities that can't be represented in Markdown unambiguously are kept as is. Can be called synchronously
// @param text The text
func (client *Client) GetMarkdownText(text *FormattedText) (*FormattedText, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "getMarkdownText",
		"text":  text,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var formattedText FormattedText
	err = json.Unmarshal(result.Raw, &formattedText)
	return &formattedText, err

}

// GetFileMimeType Returns the MIME type of a file, guessed by its extension. Returns an empty string on failure. Can be called synchronously
// @param fileName The name of the file or path to the file
func (client *Client) GetFileMimeType(fileName string) (*Text, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":     "getFileMimeType",
		"file_name": fileName,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var text Text
	err = json.Unmarshal(result.Raw, &text)
	return &text, err

}

// GetFileExtension Returns the extension of a file, guessed by its MIME type. Returns an empty string on failure. Can be called synchronously
// @param mimeType The MIME type of the file
func (client *Client) GetFileExtension(mimeType string) (*Text, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":     "getFileExtension",
		"mime_type": mimeType,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var text Text
	err = json.Unmarshal(result.Raw, &text)
	return &text, err

}

// CleanFileName Removes potentially dangerous characters from the name of a file. The encoding of the file name is supposed to be UTF-8. Returns an empty string on failure. Can be called synchronously
// @param fileName File name or path to the file
func (client *Client) CleanFileName(fileName string) (*Text, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":     "cleanFileName",
		"file_name": fileName,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var text Text
	err = json.Unmarshal(result.Raw, &text)
	return &text, err

}

// GetLanguagePackString Returns a string stored in the local database from the specified localization target and language pack by its key. Returns a 404 error if the string is not found. Can be called synchronously
// @param languagePackDatabasePath Path to the language pack database in which strings are stored
// @param localizationTarget Localization target to which the language pack belongs
// @param languagePackId Language pack identifier
// @param key Language pack key of the string to be returned
func (client *Client) GetLanguagePackString(languagePackDatabasePath string, localizationTarget string, languagePackId string, key string) (LanguagePackStringValue, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":                       "getLanguagePackString",
		"language_pack_database_path": languagePackDatabasePath,
		"localization_target":         localizationTarget,
		"language_pack_id":            languagePackId,
		"key":                         key,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	switch LanguagePackStringValueEnum(result.Data["@type"].(string)) {

	case LanguagePackStringValueOrdinaryType:
		var languagePackStringValue LanguagePackStringValueOrdinary
		err = json.Unmarshal(result.Raw, &languagePackStringValue)
		return &languagePackStringValue, err

	case LanguagePackStringValuePluralizedType:
		var languagePackStringValue LanguagePackStringValuePluralized
		err = json.Unmarshal(result.Raw, &languagePackStringValue)
		return &languagePackStringValue, err

	case LanguagePackStringValueDeletedType:
		var languagePackStringValue LanguagePackStringValueDeleted
		err = json.Unmarshal(result.Raw, &languagePackStringValue)
		return &languagePackStringValue, err

	default:
		return nil, fmt.Errorf("Invalid type")
	}
}

// GetJsonValue Converts a JSON-serialized string to corresponding JsonValue object. Can be called synchronously
// @param json The JSON-serialized string
func (client *Client) GetJsonValue(jValue string) (JsonValue, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "getJsonValue",
		"json":  jValue,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	switch JsonValueEnum(result.Data["@type"].(string)) {

	case JsonValueNullType:
		var jsonValue JsonValueNull
		err = json.Unmarshal(result.Raw, &jsonValue)
		return &jsonValue, err

	case JsonValueBooleanType:
		var jsonValue JsonValueBoolean
		err = json.Unmarshal(result.Raw, &jsonValue)
		return &jsonValue, err

	case JsonValueNumberType:
		var jsonValue JsonValueNumber
		err = json.Unmarshal(result.Raw, &jsonValue)
		return &jsonValue, err

	case JsonValueStringType:
		var jsonValue JsonValueString
		err = json.Unmarshal(result.Raw, &jsonValue)
		return &jsonValue, err

	case JsonValueArrayType:
		var jsonValue JsonValueArray
		err = json.Unmarshal(result.Raw, &jsonValue)
		return &jsonValue, err

	case JsonValueObjectType:
		var jsonValue JsonValueObject
		err = json.Unmarshal(result.Raw, &jsonValue)
		return &jsonValue, err

	default:
		return nil, fmt.Errorf("Invalid type")
	}
}

// GetJsonString Converts a JsonValue object to corresponding JSON-serialized string. Can be called synchronously
// @param jsonValue The JsonValue object
func (client *Client) GetJsonString(jsonValue JsonValue) (*Text, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":      "getJsonString",
		"json_value": jsonValue,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var text Text
	err = json.Unmarshal(result.Raw, &text)
	return &text, err

}

// SetPollAnswer Changes the user answer to a poll. A poll in quiz mode can be answered only once
// @param chatId Identifier of the chat to which the poll belongs
// @param messageId Identifier of the message containing the poll
// @param optionIds 0-based identifiers of answer options, chosen by the user. User can choose more than 1 answer option only is the poll allows multiple answers
func (client *Client) SetPollAnswer(chatId int64, messageId int64, optionIds []int32) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":      "setPollAnswer",
		"chat_id":    chatId,
		"message_id": messageId,
		"option_ids": optionIds,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// GetPollVoters Returns users voted for the specified option in a non-anonymous polls. For optimal performance, the number of returned users is chosen by TDLib
// @param chatId Identifier of the chat to which the poll belongs
// @param messageId Identifier of the message containing the poll
// @param optionId 0-based identifier of the answer option
// @param offset Number of users to skip in the result; must be non-negative
// @param limit The maximum number of users to be returned; must be positive and can't be greater than 50. For optimal performance, the number of returned users is chosen by TDLib and can be smaller than the specified limit, even if the end of the voter list has not been reached
func (client *Client) GetPollVoters(chatId int64, messageId int64, optionId int32, offset int32, limit int32) (*Users, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":      "getPollVoters",
		"chat_id":    chatId,
		"message_id": messageId,
		"option_id":  optionId,
		"offset":     offset,
		"limit":      limit,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var users Users
	err = json.Unmarshal(result.Raw, &users)
	return &users, err

}

// StopPoll Stops a poll. A poll in a message can be stopped when the message has can_be_edited flag set
// @param chatId Identifier of the chat to which the poll belongs
// @param messageId Identifier of the message containing the poll
// @param replyMarkup The new message reply markup; pass null if none; for bots only
func (client *Client) StopPoll(chatId int64, messageId int64, replyMarkup ReplyMarkup) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":        "stopPoll",
		"chat_id":      chatId,
		"message_id":   messageId,
		"reply_markup": replyMarkup,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// HideSuggestedAction Hides a suggested action
// @param action Suggested action to hide
func (client *Client) HideSuggestedAction(action SuggestedAction) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":  "hideSuggestedAction",
		"action": action,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// GetLoginUrlInfo Returns information about a button of type inlineKeyboardButtonTypeLoginUrl. The method needs to be called when the user presses the button
// @param chatId Chat identifier of the message with the button
// @param messageId Message identifier of the message with the button
// @param buttonId Button identifier
func (client *Client) GetLoginUrlInfo(chatId int64, messageId int64, buttonId int64) (LoginUrlInfo, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":      "getLoginUrlInfo",
		"chat_id":    chatId,
		"message_id": messageId,
		"button_id":  buttonId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	switch LoginUrlInfoEnum(result.Data["@type"].(string)) {

	case LoginUrlInfoOpenType:
		var loginUrlInfo LoginUrlInfoOpen
		err = json.Unmarshal(result.Raw, &loginUrlInfo)
		return &loginUrlInfo, err

	case LoginUrlInfoRequestConfirmationType:
		var loginUrlInfo LoginUrlInfoRequestConfirmation
		err = json.Unmarshal(result.Raw, &loginUrlInfo)
		return &loginUrlInfo, err

	default:
		return nil, fmt.Errorf("Invalid type")
	}
}

// GetLoginUrl Returns an HTTP URL which can be used to automatically authorize the user on a website after clicking an inline button of type inlineKeyboardButtonTypeLoginUrl.
// @param chatId Chat identifier of the message with the button
// @param messageId Message identifier of the message with the button
// @param buttonId Button identifier
// @param allowWriteAccess True, if the user allowed the bot to send them messages
func (client *Client) GetLoginUrl(chatId int64, messageId int64, buttonId int64, allowWriteAccess bool) (*HttpUrl, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":              "getLoginUrl",
		"chat_id":            chatId,
		"message_id":         messageId,
		"button_id":          buttonId,
		"allow_write_access": allowWriteAccess,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var httpUrl HttpUrl
	err = json.Unmarshal(result.Raw, &httpUrl)
	return &httpUrl, err

}

// GetInlineQueryResults Sends an inline query to a bot and returns its results. Returns an error with code 502 if the bot fails to answer the query before the query timeout expires
// @param botUserId The identifier of the target bot
// @param chatId Identifier of the chat where the query was sent
// @param userLocation Location of the user; pass null if unknown or the bot doesn't need user's location
// @param query Text of the query
// @param offset Offset of the first entry to return
func (client *Client) GetInlineQueryResults(botUserId int64, chatId int64, userLocation *Location, query string, offset string) (*InlineQueryResults, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":         "getInlineQueryResults",
		"bot_user_id":   botUserId,
		"chat_id":       chatId,
		"user_location": userLocation,
		"query":         query,
		"offset":        offset,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var inlineQueryResults InlineQueryResults
	err = json.Unmarshal(result.Raw, &inlineQueryResults)
	return &inlineQueryResults, err

}

// AnswerInlineQuery Sets the result of an inline query; for bots only
// @param inlineQueryId Identifier of the inline query
// @param isPersonal True, if the result of the query can be cached for the specified user
// @param results The results of the query
// @param cacheTime Allowed time to cache the results of the query, in seconds
// @param nextOffset Offset for the next inline query; pass an empty string if there are no more results
// @param switchPmText If non-empty, this text must be shown on the button that opens a private chat with the bot and sends a start message to the bot with the parameter switch_pm_parameter
// @param switchPmParameter The parameter for the bot start message
func (client *Client) AnswerInlineQuery(inlineQueryId JSONInt64, isPersonal bool, results []InputInlineQueryResult, cacheTime int32, nextOffset string, switchPmText string, switchPmParameter string) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":               "answerInlineQuery",
		"inline_query_id":     inlineQueryId,
		"is_personal":         isPersonal,
		"results":             results,
		"cache_time":          cacheTime,
		"next_offset":         nextOffset,
		"switch_pm_text":      switchPmText,
		"switch_pm_parameter": switchPmParameter,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// GetCallbackQueryAnswer Sends a callback query to a bot and returns an answer. Returns an error with code 502 if the bot fails to answer the query before the query timeout expires
// @param chatId Identifier of the chat with the message
// @param messageId Identifier of the message from which the query originated
// @param payload Query payload
func (client *Client) GetCallbackQueryAnswer(chatId int64, messageId int64, payload CallbackQueryPayload) (*CallbackQueryAnswer, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":      "getCallbackQueryAnswer",
		"chat_id":    chatId,
		"message_id": messageId,
		"payload":    payload,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var callbackQueryAnswer CallbackQueryAnswer
	err = json.Unmarshal(result.Raw, &callbackQueryAnswer)
	return &callbackQueryAnswer, err

}

// AnswerCallbackQuery Sets the result of a callback query; for bots only
// @param callbackQueryId Identifier of the callback query
// @param text Text of the answer
// @param showAlert If true, an alert must be shown to the user instead of a toast notification
// @param url URL to be opened
// @param cacheTime Time during which the result of the query can be cached, in seconds
func (client *Client) AnswerCallbackQuery(callbackQueryId JSONInt64, text string, showAlert bool, url string, cacheTime int32) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":             "answerCallbackQuery",
		"callback_query_id": callbackQueryId,
		"text":              text,
		"show_alert":        showAlert,
		"url":               url,
		"cache_time":        cacheTime,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// AnswerShippingQuery Sets the result of a shipping query; for bots only
// @param shippingQueryId Identifier of the shipping query
// @param shippingOptions Available shipping options
// @param errorMessage An error message, empty on success
func (client *Client) AnswerShippingQuery(shippingQueryId JSONInt64, shippingOptions []ShippingOption, errorMessage string) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":             "answerShippingQuery",
		"shipping_query_id": shippingQueryId,
		"shipping_options":  shippingOptions,
		"error_message":     errorMessage,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// AnswerPreCheckoutQuery Sets the result of a pre-checkout query; for bots only
// @param preCheckoutQueryId Identifier of the pre-checkout query
// @param errorMessage An error message, empty on success
func (client *Client) AnswerPreCheckoutQuery(preCheckoutQueryId JSONInt64, errorMessage string) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":                 "answerPreCheckoutQuery",
		"pre_checkout_query_id": preCheckoutQueryId,
		"error_message":         errorMessage,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// SetGameScore Updates the game score of the specified user in the game; for bots only
// @param chatId The chat to which the message with the game belongs
// @param messageId Identifier of the message
// @param editMessage True, if the message needs to be edited
// @param userId User identifier
// @param score The new score
// @param force Pass true to update the score even if it decreases. If the score is 0, the user will be deleted from the high score table
func (client *Client) SetGameScore(chatId int64, messageId int64, editMessage bool, userId int64, score int32, force bool) (*Message, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":        "setGameScore",
		"chat_id":      chatId,
		"message_id":   messageId,
		"edit_message": editMessage,
		"user_id":      userId,
		"score":        score,
		"force":        force,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var messageDummy Message
	err = json.Unmarshal(result.Raw, &messageDummy)
	return &messageDummy, err

}

// SetInlineGameScore Updates the game score of the specified user in a game; for bots only
// @param inlineMessageId Inline message identifier
// @param editMessage True, if the message needs to be edited
// @param userId User identifier
// @param score The new score
// @param force Pass true to update the score even if it decreases. If the score is 0, the user will be deleted from the high score table
func (client *Client) SetInlineGameScore(inlineMessageId string, editMessage bool, userId int64, score int32, force bool) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":             "setInlineGameScore",
		"inline_message_id": inlineMessageId,
		"edit_message":      editMessage,
		"user_id":           userId,
		"score":             score,
		"force":             force,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// GetGameHighScores Returns the high scores for a game and some part of the high score table in the range of the specified user; for bots only
// @param chatId The chat that contains the message with the game
// @param messageId Identifier of the message
// @param userId User identifier
func (client *Client) GetGameHighScores(chatId int64, messageId int64, userId int64) (*GameHighScores, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":      "getGameHighScores",
		"chat_id":    chatId,
		"message_id": messageId,
		"user_id":    userId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var gameHighScores GameHighScores
	err = json.Unmarshal(result.Raw, &gameHighScores)
	return &gameHighScores, err

}

// GetInlineGameHighScores Returns game high scores and some part of the high score table in the range of the specified user; for bots only
// @param inlineMessageId Inline message identifier
// @param userId User identifier
func (client *Client) GetInlineGameHighScores(inlineMessageId string, userId int64) (*GameHighScores, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":             "getInlineGameHighScores",
		"inline_message_id": inlineMessageId,
		"user_id":           userId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var gameHighScores GameHighScores
	err = json.Unmarshal(result.Raw, &gameHighScores)
	return &gameHighScores, err

}

// DeleteChatReplyMarkup Deletes the default reply markup from a chat. Must be called after a one-time keyboard or a ForceReply reply markup has been used. UpdateChatReplyMarkup will be sent if the reply markup is changed
// @param chatId Chat identifier
// @param messageId The message identifier of the used keyboard
func (client *Client) DeleteChatReplyMarkup(chatId int64, messageId int64) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":      "deleteChatReplyMarkup",
		"chat_id":    chatId,
		"message_id": messageId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// SendChatAction Sends a notification about user activity in a chat
// @param chatId Chat identifier
// @param messageThreadId If not 0, a message thread identifier in which the action was performed
// @param action The action description; pass null to cancel the currently active action
func (client *Client) SendChatAction(chatId int64, messageThreadId int64, action ChatAction) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":             "sendChatAction",
		"chat_id":           chatId,
		"message_thread_id": messageThreadId,
		"action":            action,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// OpenChat Informs TDLib that the chat is opened by the user. Many useful activities depend on the chat being opened or closed (e.g., in supergroups and channels all updates are received only for opened chats)
// @param chatId Chat identifier
func (client *Client) OpenChat(chatId int64) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "openChat",
		"chat_id": chatId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// CloseChat Informs TDLib that the chat is closed by the user. Many useful activities depend on the chat being opened or closed
// @param chatId Chat identifier
func (client *Client) CloseChat(chatId int64) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "closeChat",
		"chat_id": chatId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// ViewMessages Informs TDLib that messages are being viewed by the user. Many useful activities depend on whether the messages are currently being viewed or not (e.g., marking messages as read, incrementing a view counter, updating a view counter, removing deleted messages in supergroups and channels)
// @param chatId Chat identifier
// @param messageThreadId If not 0, a message thread identifier in which the messages are being viewed
// @param messageIds The identifiers of the messages being viewed
// @param forceRead True, if messages in closed chats must be marked as read by the request
func (client *Client) ViewMessages(chatId int64, messageThreadId int64, messageIds []int64, forceRead bool) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":             "viewMessages",
		"chat_id":           chatId,
		"message_thread_id": messageThreadId,
		"message_ids":       messageIds,
		"force_read":        forceRead,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// OpenMessageContent Informs TDLib that the message content has been opened (e.g., the user has opened a photo, video, document, location or venue, or has listened to an audio file or voice note message). An updateMessageContentOpened update will be generated if something has changed
// @param chatId Chat identifier of the message
// @param messageId Identifier of the message with the opened content
func (client *Client) OpenMessageContent(chatId int64, messageId int64) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":      "openMessageContent",
		"chat_id":    chatId,
		"message_id": messageId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// ClickAnimatedEmojiMessage Informs TDLib that a message with an animated emoji was clicked by the user. Returns a big animated sticker to be played or a 404 error if usual animation needs to be played
// @param chatId Chat identifier of the message
// @param messageId Identifier of the clicked message
func (client *Client) ClickAnimatedEmojiMessage(chatId int64, messageId int64) (*Sticker, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":      "clickAnimatedEmojiMessage",
		"chat_id":    chatId,
		"message_id": messageId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var sticker Sticker
	err = json.Unmarshal(result.Raw, &sticker)
	return &sticker, err

}

// GetInternalLinkType Returns information about the type of an internal link. Returns a 404 error if the link is not internal. Can be called before authorization
// @param link The link
func (client *Client) GetInternalLinkType(link string) (InternalLinkType, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "getInternalLinkType",
		"link":  link,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	switch InternalLinkTypeEnum(result.Data["@type"].(string)) {

	case InternalLinkTypeActiveSessionsType:
		var internalLinkType InternalLinkTypeActiveSessions
		err = json.Unmarshal(result.Raw, &internalLinkType)
		return &internalLinkType, err

	case InternalLinkTypeAuthenticationCodeType:
		var internalLinkType InternalLinkTypeAuthenticationCode
		err = json.Unmarshal(result.Raw, &internalLinkType)
		return &internalLinkType, err

	case InternalLinkTypeBackgroundType:
		var internalLinkType InternalLinkTypeBackground
		err = json.Unmarshal(result.Raw, &internalLinkType)
		return &internalLinkType, err

	case InternalLinkTypeBotStartType:
		var internalLinkType InternalLinkTypeBotStart
		err = json.Unmarshal(result.Raw, &internalLinkType)
		return &internalLinkType, err

	case InternalLinkTypeBotStartInGroupType:
		var internalLinkType InternalLinkTypeBotStartInGroup
		err = json.Unmarshal(result.Raw, &internalLinkType)
		return &internalLinkType, err

	case InternalLinkTypeChangePhoneNumberType:
		var internalLinkType InternalLinkTypeChangePhoneNumber
		err = json.Unmarshal(result.Raw, &internalLinkType)
		return &internalLinkType, err

	case InternalLinkTypeChatInviteType:
		var internalLinkType InternalLinkTypeChatInvite
		err = json.Unmarshal(result.Raw, &internalLinkType)
		return &internalLinkType, err

	case InternalLinkTypeFilterSettingsType:
		var internalLinkType InternalLinkTypeFilterSettings
		err = json.Unmarshal(result.Raw, &internalLinkType)
		return &internalLinkType, err

	case InternalLinkTypeGameType:
		var internalLinkType InternalLinkTypeGame
		err = json.Unmarshal(result.Raw, &internalLinkType)
		return &internalLinkType, err

	case InternalLinkTypeLanguagePackType:
		var internalLinkType InternalLinkTypeLanguagePack
		err = json.Unmarshal(result.Raw, &internalLinkType)
		return &internalLinkType, err

	case InternalLinkTypeMessageType:
		var internalLinkType InternalLinkTypeMessage
		err = json.Unmarshal(result.Raw, &internalLinkType)
		return &internalLinkType, err

	case InternalLinkTypeMessageDraftType:
		var internalLinkType InternalLinkTypeMessageDraft
		err = json.Unmarshal(result.Raw, &internalLinkType)
		return &internalLinkType, err

	case InternalLinkTypePassportDataRequestType:
		var internalLinkType InternalLinkTypePassportDataRequest
		err = json.Unmarshal(result.Raw, &internalLinkType)
		return &internalLinkType, err

	case InternalLinkTypePhoneNumberConfirmationType:
		var internalLinkType InternalLinkTypePhoneNumberConfirmation
		err = json.Unmarshal(result.Raw, &internalLinkType)
		return &internalLinkType, err

	case InternalLinkTypeProxyType:
		var internalLinkType InternalLinkTypeProxy
		err = json.Unmarshal(result.Raw, &internalLinkType)
		return &internalLinkType, err

	case InternalLinkTypePublicChatType:
		var internalLinkType InternalLinkTypePublicChat
		err = json.Unmarshal(result.Raw, &internalLinkType)
		return &internalLinkType, err

	case InternalLinkTypeQrCodeAuthenticationType:
		var internalLinkType InternalLinkTypeQrCodeAuthentication
		err = json.Unmarshal(result.Raw, &internalLinkType)
		return &internalLinkType, err

	case InternalLinkTypeSettingsType:
		var internalLinkType InternalLinkTypeSettings
		err = json.Unmarshal(result.Raw, &internalLinkType)
		return &internalLinkType, err

	case InternalLinkTypeStickerSetType:
		var internalLinkType InternalLinkTypeStickerSet
		err = json.Unmarshal(result.Raw, &internalLinkType)
		return &internalLinkType, err

	case InternalLinkTypeThemeType:
		var internalLinkType InternalLinkTypeTheme
		err = json.Unmarshal(result.Raw, &internalLinkType)
		return &internalLinkType, err

	case InternalLinkTypeThemeSettingsType:
		var internalLinkType InternalLinkTypeThemeSettings
		err = json.Unmarshal(result.Raw, &internalLinkType)
		return &internalLinkType, err

	case InternalLinkTypeUnknownDeepLinkType:
		var internalLinkType InternalLinkTypeUnknownDeepLink
		err = json.Unmarshal(result.Raw, &internalLinkType)
		return &internalLinkType, err

	case InternalLinkTypeUnsupportedProxyType:
		var internalLinkType InternalLinkTypeUnsupportedProxy
		err = json.Unmarshal(result.Raw, &internalLinkType)
		return &internalLinkType, err

	case InternalLinkTypeVideoChatType:
		var internalLinkType InternalLinkTypeVideoChat
		err = json.Unmarshal(result.Raw, &internalLinkType)
		return &internalLinkType, err

	default:
		return nil, fmt.Errorf("Invalid type")
	}
}

// GetExternalLinkInfo Returns information about an action to be done when the current user clicks an external link. Don't use this method for links from secret chats if web page preview is disabled in secret chats
// @param link The link
func (client *Client) GetExternalLinkInfo(link string) (LoginUrlInfo, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "getExternalLinkInfo",
		"link":  link,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	switch LoginUrlInfoEnum(result.Data["@type"].(string)) {

	case LoginUrlInfoOpenType:
		var loginUrlInfo LoginUrlInfoOpen
		err = json.Unmarshal(result.Raw, &loginUrlInfo)
		return &loginUrlInfo, err

	case LoginUrlInfoRequestConfirmationType:
		var loginUrlInfo LoginUrlInfoRequestConfirmation
		err = json.Unmarshal(result.Raw, &loginUrlInfo)
		return &loginUrlInfo, err

	default:
		return nil, fmt.Errorf("Invalid type")
	}
}

// GetExternalLink Returns an HTTP URL which can be used to automatically authorize the current user on a website after clicking an HTTP link. Use the method getExternalLinkInfo to find whether a prior user confirmation is needed
// @param link The HTTP link
// @param allowWriteAccess True, if the current user allowed the bot, returned in getExternalLinkInfo, to send them messages
func (client *Client) GetExternalLink(link string, allowWriteAccess bool) (*HttpUrl, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":              "getExternalLink",
		"link":               link,
		"allow_write_access": allowWriteAccess,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var httpUrl HttpUrl
	err = json.Unmarshal(result.Raw, &httpUrl)
	return &httpUrl, err

}

// ReadAllChatMentions Marks all mentions in a chat as read
// @param chatId Chat identifier
func (client *Client) ReadAllChatMentions(chatId int64) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "readAllChatMentions",
		"chat_id": chatId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// CreatePrivateChat Returns an existing chat corresponding to a given user
// @param userId User identifier
// @param force If true, the chat will be created without network request. In this case all information about the chat except its type, title and photo can be incorrect
func (client *Client) CreatePrivateChat(userId int64, force bool) (*Chat, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "createPrivateChat",
		"user_id": userId,
		"force":   force,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var chat Chat
	err = json.Unmarshal(result.Raw, &chat)
	return &chat, err

}

// CreateBasicGroupChat Returns an existing chat corresponding to a known basic group
// @param basicGroupId Basic group identifier
// @param force If true, the chat will be created without network request. In this case all information about the chat except its type, title and photo can be incorrect
func (client *Client) CreateBasicGroupChat(basicGroupId int64, force bool) (*Chat, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":          "createBasicGroupChat",
		"basic_group_id": basicGroupId,
		"force":          force,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var chat Chat
	err = json.Unmarshal(result.Raw, &chat)
	return &chat, err

}

// CreateSupergroupChat Returns an existing chat corresponding to a known supergroup or channel
// @param supergroupId Supergroup or channel identifier
// @param force If true, the chat will be created without network request. In this case all information about the chat except its type, title and photo can be incorrect
func (client *Client) CreateSupergroupChat(supergroupId int64, force bool) (*Chat, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":         "createSupergroupChat",
		"supergroup_id": supergroupId,
		"force":         force,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var chat Chat
	err = json.Unmarshal(result.Raw, &chat)
	return &chat, err

}

// CreateSecretChat Returns an existing chat corresponding to a known secret chat
// @param secretChatId Secret chat identifier
func (client *Client) CreateSecretChat(secretChatId int32) (*Chat, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":          "createSecretChat",
		"secret_chat_id": secretChatId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var chatDummy Chat
	err = json.Unmarshal(result.Raw, &chatDummy)
	return &chatDummy, err

}

// CreateNewBasicGroupChat Creates a new basic group and sends a corresponding messageBasicGroupChatCreate. Returns the newly created chat
// @param userIds Identifiers of users to be added to the basic group
// @param title Title of the new basic group; 1-128 characters
func (client *Client) CreateNewBasicGroupChat(userIds []int64, title string) (*Chat, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":    "createNewBasicGroupChat",
		"user_ids": userIds,
		"title":    title,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var chat Chat
	err = json.Unmarshal(result.Raw, &chat)
	return &chat, err

}

// CreateNewSupergroupChat Creates a new supergroup or channel and sends a corresponding messageSupergroupChatCreate. Returns the newly created chat
// @param title Title of the new chat; 1-128 characters
// @param isChannel True, if a channel chat needs to be created
// @param description
// @param location Chat location if a location-based supergroup is being created; pass null to create an ordinary supergroup chat
// @param forImport True, if the supergroup is created for importing messages using importMessage
func (client *Client) CreateNewSupergroupChat(title string, isChannel bool, description string, location *ChatLocation, forImport bool) (*Chat, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":       "createNewSupergroupChat",
		"title":       title,
		"is_channel":  isChannel,
		"description": description,
		"location":    location,
		"for_import":  forImport,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var chat Chat
	err = json.Unmarshal(result.Raw, &chat)
	return &chat, err

}

// CreateNewSecretChat Creates a new secret chat. Returns the newly created chat
// @param userId Identifier of the target user
func (client *Client) CreateNewSecretChat(userId int64) (*Chat, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "createNewSecretChat",
		"user_id": userId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var chat Chat
	err = json.Unmarshal(result.Raw, &chat)
	return &chat, err

}

// UpgradeBasicGroupChatToSupergroupChat Creates a new supergroup from an existing basic group and sends a corresponding messageChatUpgradeTo and messageChatUpgradeFrom; requires creator privileges. Deactivates the original basic group
// @param chatId Identifier of the chat to upgrade
func (client *Client) UpgradeBasicGroupChatToSupergroupChat(chatId int64) (*Chat, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "upgradeBasicGroupChatToSupergroupChat",
		"chat_id": chatId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var chatDummy Chat
	err = json.Unmarshal(result.Raw, &chatDummy)
	return &chatDummy, err

}

// GetChatListsToAddChat Returns chat lists to which the chat can be added. This is an offline request
// @param chatId Chat identifier
func (client *Client) GetChatListsToAddChat(chatId int64) (*ChatLists, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "getChatListsToAddChat",
		"chat_id": chatId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var chatLists ChatLists
	err = json.Unmarshal(result.Raw, &chatLists)
	return &chatLists, err

}

// AddChatToList Adds a chat to a chat list. A chat can't be simultaneously in Main and Archive chat lists, so it is automatically removed from another one if needed
// @param chatId Chat identifier
// @param chatList The chat list. Use getChatListsToAddChat to get suitable chat lists
func (client *Client) AddChatToList(chatId int64, chatList ChatList) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":     "addChatToList",
		"chat_id":   chatId,
		"chat_list": chatList,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// GetChatFilter Returns information about a chat filter by its identifier
// @param chatFilterId Chat filter identifier
func (client *Client) GetChatFilter(chatFilterId int32) (*ChatFilter, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":          "getChatFilter",
		"chat_filter_id": chatFilterId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var chatFilterDummy ChatFilter
	err = json.Unmarshal(result.Raw, &chatFilterDummy)
	return &chatFilterDummy, err

}

// CreateChatFilter Creates new chat filter. Returns information about the created chat filter
// @param filter Chat filter
func (client *Client) CreateChatFilter(filter *ChatFilter) (*ChatFilterInfo, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":  "createChatFilter",
		"filter": filter,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var chatFilterInfo ChatFilterInfo
	err = json.Unmarshal(result.Raw, &chatFilterInfo)
	return &chatFilterInfo, err

}

// EditChatFilter Edits existing chat filter. Returns information about the edited chat filter
// @param chatFilterId Chat filter identifier
// @param filter The edited chat filter
func (client *Client) EditChatFilter(chatFilterId int32, filter *ChatFilter) (*ChatFilterInfo, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":          "editChatFilter",
		"chat_filter_id": chatFilterId,
		"filter":         filter,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var chatFilterInfo ChatFilterInfo
	err = json.Unmarshal(result.Raw, &chatFilterInfo)
	return &chatFilterInfo, err

}

// DeleteChatFilter Deletes existing chat filter
// @param chatFilterId Chat filter identifier
func (client *Client) DeleteChatFilter(chatFilterId int32) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":          "deleteChatFilter",
		"chat_filter_id": chatFilterId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// ReorderChatFilters Changes the order of chat filters
// @param chatFilterIds Identifiers of chat filters in the new correct order
func (client *Client) ReorderChatFilters(chatFilterIds []int32) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":           "reorderChatFilters",
		"chat_filter_ids": chatFilterIds,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// GetRecommendedChatFilters Returns recommended chat filters for the current user
func (client *Client) GetRecommendedChatFilters() (*RecommendedChatFilters, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "getRecommendedChatFilters",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var recommendedChatFilters RecommendedChatFilters
	err = json.Unmarshal(result.Raw, &recommendedChatFilters)
	return &recommendedChatFilters, err

}

// GetChatFilterDefaultIconName Returns default icon name for a filter. Can be called synchronously
// @param filter Chat filter
func (client *Client) GetChatFilterDefaultIconName(filter *ChatFilter) (*Text, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":  "getChatFilterDefaultIconName",
		"filter": filter,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var text Text
	err = json.Unmarshal(result.Raw, &text)
	return &text, err

}

// SetChatTitle Changes the chat title. Supported only for basic groups, supergroups and channels. Requires can_change_info administrator right
// @param chatId Chat identifier
// @param title New title of the chat; 1-128 characters
func (client *Client) SetChatTitle(chatId int64, title string) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "setChatTitle",
		"chat_id": chatId,
		"title":   title,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// SetChatPhoto Changes the photo of a chat. Supported only for basic groups, supergroups and channels. Requires can_change_info administrator right
// @param chatId Chat identifier
// @param photo New chat photo; pass null to delete the chat photo
func (client *Client) SetChatPhoto(chatId int64, photo InputChatPhoto) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "setChatPhoto",
		"chat_id": chatId,
		"photo":   photo,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// SetChatMessageTtlSetting Changes the message TTL setting (sets a new self-destruct timer) in a chat. Requires can_delete_messages administrator right in basic groups, supergroups and channels
// @param chatId Chat identifier
// @param ttl New TTL value, in seconds; must be one of 0, 86400, 7 * 86400, or 31 * 86400 unless the chat is secret
func (client *Client) SetChatMessageTtlSetting(chatId int64, ttl int32) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "setChatMessageTtlSetting",
		"chat_id": chatId,
		"ttl":     ttl,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// SetChatPermissions Changes the chat members permissions. Supported only for basic groups and supergroups. Requires can_restrict_members administrator right
// @param chatId Chat identifier
// @param permissions New non-administrator members permissions in the chat
func (client *Client) SetChatPermissions(chatId int64, permissions *ChatPermissions) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":       "setChatPermissions",
		"chat_id":     chatId,
		"permissions": permissions,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// SetChatTheme Changes the chat theme. Supported only in private and secret chats
// @param chatId Chat identifier
// @param themeName Name of the new chat theme; pass an empty string to return the default theme
func (client *Client) SetChatTheme(chatId int64, themeName string) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":      "setChatTheme",
		"chat_id":    chatId,
		"theme_name": themeName,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// SetChatDraftMessage Changes the draft message in a chat
// @param chatId Chat identifier
// @param messageThreadId If not 0, a message thread identifier in which the draft was changed
// @param draftMessage New draft message; pass null to remove the draft
func (client *Client) SetChatDraftMessage(chatId int64, messageThreadId int64, draftMessage *DraftMessage) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":             "setChatDraftMessage",
		"chat_id":           chatId,
		"message_thread_id": messageThreadId,
		"draft_message":     draftMessage,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// SetChatNotificationSettings Changes the notification settings of a chat. Notification settings of a chat with the current user (Saved Messages) can't be changed
// @param chatId Chat identifier
// @param notificationSettings New notification settings for the chat. If the chat is muted for more than 1 week, it is considered to be muted forever
func (client *Client) SetChatNotificationSettings(chatId int64, notificationSettings *ChatNotificationSettings) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":                 "setChatNotificationSettings",
		"chat_id":               chatId,
		"notification_settings": notificationSettings,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// ToggleChatIsMarkedAsUnread Changes the marked as unread state of a chat
// @param chatId Chat identifier
// @param isMarkedAsUnread New value of is_marked_as_unread
func (client *Client) ToggleChatIsMarkedAsUnread(chatId int64, isMarkedAsUnread bool) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":               "toggleChatIsMarkedAsUnread",
		"chat_id":             chatId,
		"is_marked_as_unread": isMarkedAsUnread,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// ToggleChatDefaultDisableNotification Changes the value of the default disable_notification parameter, used when a message is sent to a chat
// @param chatId Chat identifier
// @param defaultDisableNotification New value of default_disable_notification
func (client *Client) ToggleChatDefaultDisableNotification(chatId int64, defaultDisableNotification bool) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":                        "toggleChatDefaultDisableNotification",
		"chat_id":                      chatId,
		"default_disable_notification": defaultDisableNotification,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// SetChatClientData Changes application-specific data associated with a chat
// @param chatId Chat identifier
// @param clientData New value of client_data
func (client *Client) SetChatClientData(chatId int64, clientData string) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":       "setChatClientData",
		"chat_id":     chatId,
		"client_data": clientData,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// SetChatDescription Changes information about a chat. Available for basic groups, supergroups, and channels. Requires can_change_info administrator right
// @param chatId Identifier of the chat
// @param description
func (client *Client) SetChatDescription(chatId int64, description string) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":       "setChatDescription",
		"chat_id":     chatId,
		"description": description,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// SetChatDiscussionGroup Changes the discussion group of a channel chat; requires can_change_info administrator right in the channel if it is specified
// @param chatId Identifier of the channel chat. Pass 0 to remove a link from the supergroup passed in the second argument to a linked channel chat (requires can_pin_messages rights in the supergroup)
// @param discussionChatId Identifier of a new channel's discussion group. Use 0 to remove the discussion group.
func (client *Client) SetChatDiscussionGroup(chatId int64, discussionChatId int64) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":              "setChatDiscussionGroup",
		"chat_id":            chatId,
		"discussion_chat_id": discussionChatId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// SetChatLocation Changes the location of a chat. Available only for some location-based supergroups, use supergroupFullInfo.can_set_location to check whether the method is allowed to use
// @param chatId Chat identifier
// @param location New location for the chat; must be valid and not null
func (client *Client) SetChatLocation(chatId int64, location *ChatLocation) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":    "setChatLocation",
		"chat_id":  chatId,
		"location": location,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// SetChatSlowModeDelay Changes the slow mode delay of a chat. Available only for supergroups; requires can_restrict_members rights
// @param chatId Chat identifier
// @param slowModeDelay New slow mode delay for the chat, in seconds; must be one of 0, 10, 30, 60, 300, 900, 3600
func (client *Client) SetChatSlowModeDelay(chatId int64, slowModeDelay int32) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":           "setChatSlowModeDelay",
		"chat_id":         chatId,
		"slow_mode_delay": slowModeDelay,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// PinChatMessage Pins a message in a chat; requires can_pin_messages rights or can_edit_messages rights in the channel
// @param chatId Identifier of the chat
// @param messageId Identifier of the new pinned message
// @param disableNotification True, if there must be no notification about the pinned message. Notifications are always disabled in channels and private chats
// @param onlyForSelf True, if the message needs to be pinned for one side only; private chats only
func (client *Client) PinChatMessage(chatId int64, messageId int64, disableNotification bool, onlyForSelf bool) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":                "pinChatMessage",
		"chat_id":              chatId,
		"message_id":           messageId,
		"disable_notification": disableNotification,
		"only_for_self":        onlyForSelf,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// UnpinChatMessage Removes a pinned message from a chat; requires can_pin_messages rights in the group or can_edit_messages rights in the channel
// @param chatId Identifier of the chat
// @param messageId Identifier of the removed pinned message
func (client *Client) UnpinChatMessage(chatId int64, messageId int64) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":      "unpinChatMessage",
		"chat_id":    chatId,
		"message_id": messageId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// UnpinAllChatMessages Removes all pinned messages from a chat; requires can_pin_messages rights in the group or can_edit_messages rights in the channel
// @param chatId Identifier of the chat
func (client *Client) UnpinAllChatMessages(chatId int64) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "unpinAllChatMessages",
		"chat_id": chatId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// JoinChat Adds the current user as a new member to a chat. Private and secret chats can't be joined using this method
// @param chatId Chat identifier
func (client *Client) JoinChat(chatId int64) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "joinChat",
		"chat_id": chatId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// LeaveChat Removes the current user from chat members. Private and secret chats can't be left using this method
// @param chatId Chat identifier
func (client *Client) LeaveChat(chatId int64) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "leaveChat",
		"chat_id": chatId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// AddChatMember Adds a new member to a chat. Members can't be added to private or secret chats
// @param chatId Chat identifier
// @param userId Identifier of the user
// @param forwardLimit The number of earlier messages from the chat to be forwarded to the new member; up to 100. Ignored for supergroups and channels, or if the added user is a bot
func (client *Client) AddChatMember(chatId int64, userId int64, forwardLimit int32) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":         "addChatMember",
		"chat_id":       chatId,
		"user_id":       userId,
		"forward_limit": forwardLimit,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// AddChatMembers Adds multiple new members to a chat. Currently this method is only available for supergroups and channels. This method can't be used to join a chat. Members can't be added to a channel if it has more than 200 members
// @param chatId Chat identifier
// @param userIds Identifiers of the users to be added to the chat. The maximum number of added users is 20 for supergroups and 100 for channels
func (client *Client) AddChatMembers(chatId int64, userIds []int64) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":    "addChatMembers",
		"chat_id":  chatId,
		"user_ids": userIds,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// SetChatMemberStatus Changes the status of a chat member, needs appropriate privileges. This function is currently not suitable for transferring chat ownership; use transferChatOwnership instead. Use addChatMember or banChatMember if some additional parameters needs to be passed
// @param chatId Chat identifier
// @param memberId Member identifier. Chats can be only banned and unbanned in supergroups and channels
// @param status The new status of the member in the chat
func (client *Client) SetChatMemberStatus(chatId int64, memberId MessageSender, status ChatMemberStatus) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":     "setChatMemberStatus",
		"chat_id":   chatId,
		"member_id": memberId,
		"status":    status,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// BanChatMember Bans a member in a chat. Members can't be banned in private or secret chats. In supergroups and channels, the user will not be able to return to the group on their own using invite links, etc., unless unbanned first
// @param chatId Chat identifier
// @param memberId Member identifier
// @param bannedUntilDate Point in time (Unix timestamp) when the user will be unbanned; 0 if never. If the user is banned for more than 366 days or for less than 30 seconds from the current time, the user is considered to be banned forever. Ignored in basic groups
// @param revokeMessages Pass true to delete all messages in the chat for the user that is being removed. Always true for supergroups and channels
func (client *Client) BanChatMember(chatId int64, memberId MessageSender, bannedUntilDate int32, revokeMessages bool) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":             "banChatMember",
		"chat_id":           chatId,
		"member_id":         memberId,
		"banned_until_date": bannedUntilDate,
		"revoke_messages":   revokeMessages,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var okDummy Ok
	err = json.Unmarshal(result.Raw, &okDummy)
	return &okDummy, err

}

// CanTransferOwnership Checks whether the current session can be used to transfer a chat ownership to another user
func (client *Client) CanTransferOwnership() (CanTransferOwnershipResult, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "canTransferOwnership",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	switch CanTransferOwnershipResultEnum(result.Data["@type"].(string)) {

	case CanTransferOwnershipResultOkType:
		var canTransferOwnershipResult CanTransferOwnershipResultOk
		err = json.Unmarshal(result.Raw, &canTransferOwnershipResult)
		return &canTransferOwnershipResult, err

	case CanTransferOwnershipResultPasswordNeededType:
		var canTransferOwnershipResult CanTransferOwnershipResultPasswordNeeded
		err = json.Unmarshal(result.Raw, &canTransferOwnershipResult)
		return &canTransferOwnershipResult, err

	case CanTransferOwnershipResultPasswordTooFreshType:
		var canTransferOwnershipResult CanTransferOwnershipResultPasswordTooFresh
		err = json.Unmarshal(result.Raw, &canTransferOwnershipResult)
		return &canTransferOwnershipResult, err

	case CanTransferOwnershipResultSessionTooFreshType:
		var canTransferOwnershipResult CanTransferOwnershipResultSessionTooFresh
		err = json.Unmarshal(result.Raw, &canTransferOwnershipResult)
		return &canTransferOwnershipResult, err

	default:
		return nil, fmt.Errorf("Invalid type")
	}
}

// TransferChatOwnership Changes the owner of a chat. The current user must be a current owner of the chat. Use the method canTransferOwnership to check whether the ownership can be transferred from the current session. Available only for supergroups and channel chats
// @param chatId Chat identifier
// @param userId Identifier of the user to which transfer the ownership. The ownership can't be transferred to a bot or to a deleted user
// @param password The password of the current user
func (client *Client) TransferChatOwnership(chatId int64, userId int64, password string) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":    "transferChatOwnership",
		"chat_id":  chatId,
		"user_id":  userId,
		"password": password,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// GetChatMember Returns information about a single member of a chat
// @param chatId Chat identifier
// @param memberId Member identifier
func (client *Client) GetChatMember(chatId int64, memberId MessageSender) (*ChatMember, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":     "getChatMember",
		"chat_id":   chatId,
		"member_id": memberId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var chatMember ChatMember
	err = json.Unmarshal(result.Raw, &chatMember)
	return &chatMember, err

}

// SearchChatMembers Searches for a specified query in the first name, last name and username of the members of a specified chat. Requires administrator rights in channels
// @param chatId Chat identifier
// @param query Query to search for
// @param limit The maximum number of users to be returned; up to 200
// @param filter The type of users to search for; pass null to search among all chat members
func (client *Client) SearchChatMembers(chatId int64, query string, limit int32, filter ChatMembersFilter) (*ChatMembers, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "searchChatMembers",
		"chat_id": chatId,
		"query":   query,
		"limit":   limit,
		"filter":  filter,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var chatMembers ChatMembers
	err = json.Unmarshal(result.Raw, &chatMembers)
	return &chatMembers, err

}

// GetChatAdministrators Returns a list of administrators of the chat with their custom titles
// @param chatId Chat identifier
func (client *Client) GetChatAdministrators(chatId int64) (*ChatAdministrators, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "getChatAdministrators",
		"chat_id": chatId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var chatAdministrators ChatAdministrators
	err = json.Unmarshal(result.Raw, &chatAdministrators)
	return &chatAdministrators, err

}

// ClearAllDraftMessages Clears draft messages in all chats
// @param excludeSecretChats If true, local draft messages in secret chats will not be cleared
func (client *Client) ClearAllDraftMessages(excludeSecretChats bool) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":                "clearAllDraftMessages",
		"exclude_secret_chats": excludeSecretChats,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// GetChatNotificationSettingsExceptions Returns list of chats with non-default notification settings
// @param scope If specified, only chats from the scope will be returned; pass null to return chats from all scopes
// @param compareSound If true, also chats with non-default sound will be returned
func (client *Client) GetChatNotificationSettingsExceptions(scope NotificationSettingsScope, compareSound bool) (*Chats, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":         "getChatNotificationSettingsExceptions",
		"scope":         scope,
		"compare_sound": compareSound,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var chats Chats
	err = json.Unmarshal(result.Raw, &chats)
	return &chats, err

}

// GetScopeNotificationSettings Returns the notification settings for chats of a given type
// @param scope Types of chats for which to return the notification settings information
func (client *Client) GetScopeNotificationSettings(scope NotificationSettingsScope) (*ScopeNotificationSettings, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "getScopeNotificationSettings",
		"scope": scope,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var scopeNotificationSettings ScopeNotificationSettings
	err = json.Unmarshal(result.Raw, &scopeNotificationSettings)
	return &scopeNotificationSettings, err

}

// SetScopeNotificationSettings Changes notification settings for chats of a given type
// @param scope Types of chats for which to change the notification settings
// @param notificationSettings The new notification settings for the given scope
func (client *Client) SetScopeNotificationSettings(scope NotificationSettingsScope, notificationSettings *ScopeNotificationSettings) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":                 "setScopeNotificationSettings",
		"scope":                 scope,
		"notification_settings": notificationSettings,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// ResetAllNotificationSettings Resets all notification settings to their default values. By default, all chats are unmuted, the sound is set to "default" and message previews are shown
func (client *Client) ResetAllNotificationSettings() (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "resetAllNotificationSettings",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// ToggleChatIsPinned Changes the pinned state of a chat. There can be up to GetOption("pinned_chat_count_max")/GetOption("pinned_archived_chat_count_max") pinned non-secret chats and the same number of secret chats in the main/arhive chat list
// @param chatList Chat list in which to change the pinned state of the chat
// @param chatId Chat identifier
// @param isPinned True, if the chat is pinned
func (client *Client) ToggleChatIsPinned(chatList ChatList, chatId int64, isPinned bool) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":     "toggleChatIsPinned",
		"chat_list": chatList,
		"chat_id":   chatId,
		"is_pinned": isPinned,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// SetPinnedChats Changes the order of pinned chats
// @param chatList Chat list in which to change the order of pinned chats
// @param chatIds The new list of pinned chats
func (client *Client) SetPinnedChats(chatList ChatList, chatIds []int64) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":     "setPinnedChats",
		"chat_list": chatList,
		"chat_ids":  chatIds,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// DownloadFile Downloads a file from the cloud. Download progress and completion of the download will be notified through updateFile updates
// @param fileId Identifier of the file to download
// @param priority Priority of the download (1-32). The higher the priority, the earlier the file will be downloaded. If the priorities of two files are equal, then the last one for which downloadFile was called will be downloaded first
// @param offset The starting position from which the file needs to be downloaded
// @param limit Number of bytes which need to be downloaded starting from the "offset" position before the download will be automatically canceled; use 0 to download without a limit
// @param synchronous If false, this request returns file state just after the download has been started. If true, this request returns file state only after
func (client *Client) DownloadFile(fileId int32, priority int32, offset int32, limit int32, synchronous bool) (*File, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":       "downloadFile",
		"file_id":     fileId,
		"priority":    priority,
		"offset":      offset,
		"limit":       limit,
		"synchronous": synchronous,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var fileDummy File
	err = json.Unmarshal(result.Raw, &fileDummy)
	return &fileDummy, err

}

// GetFileDownloadedPrefixSize Returns file downloaded prefix size from a given offset, in bytes
// @param fileId Identifier of the file
// @param offset Offset from which downloaded prefix size needs to be calculated
func (client *Client) GetFileDownloadedPrefixSize(fileId int32, offset int32) (*Count, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "getFileDownloadedPrefixSize",
		"file_id": fileId,
		"offset":  offset,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var count Count
	err = json.Unmarshal(result.Raw, &count)
	return &count, err

}

// CancelDownloadFile Stops the downloading of a file. If a file has already been downloaded, does nothing
// @param fileId Identifier of a file to stop downloading
// @param onlyIfPending Pass true to stop downloading only if it hasn't been started, i.e. request hasn't been sent to server
func (client *Client) CancelDownloadFile(fileId int32, onlyIfPending bool) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":           "cancelDownloadFile",
		"file_id":         fileId,
		"only_if_pending": onlyIfPending,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// GetSuggestedFileName Returns suggested name for saving a file in a given directory
// @param fileId Identifier of the file
// @param directory Directory in which the file is supposed to be saved
func (client *Client) GetSuggestedFileName(fileId int32, directory string) (*Text, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":     "getSuggestedFileName",
		"file_id":   fileId,
		"directory": directory,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var text Text
	err = json.Unmarshal(result.Raw, &text)
	return &text, err

}

// UploadFile Asynchronously uploads a file to the cloud without sending it in a message. updateFile will be used to notify about upload progress and successful completion of the upload. The file will not have a persistent remote identifier until it will be sent in a message
// @param file File to upload
// @param fileType File type; pass null if unknown
// @param priority Priority of the upload (1-32). The higher the priority, the earlier the file will be uploaded. If the priorities of two files are equal, then the first one for which uploadFile was called will be uploaded first
func (client *Client) UploadFile(file InputFile, fileType FileType, priority int32) (*File, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":     "uploadFile",
		"file":      file,
		"file_type": fileType,
		"priority":  priority,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var fileDummy File
	err = json.Unmarshal(result.Raw, &fileDummy)
	return &fileDummy, err

}

// CancelUploadFile Stops the uploading of a file. Supported only for files uploaded by using uploadFile. For other files the behavior is undefined
// @param fileId Identifier of the file to stop uploading
func (client *Client) CancelUploadFile(fileId int32) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "cancelUploadFile",
		"file_id": fileId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// WriteGeneratedFilePart Writes a part of a generated file. This method is intended to be used only if the application has no direct access to TDLib's file system, because it is usually slower than a direct write to the destination file
// @param generationId The identifier of the generation process
// @param offset The offset from which to write the data to the file
// @param data The data to write
func (client *Client) WriteGeneratedFilePart(generationId JSONInt64, offset int32, data []byte) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":         "writeGeneratedFilePart",
		"generation_id": generationId,
		"offset":        offset,
		"data":          data,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// SetFileGenerationProgress Informs TDLib on a file generation progress
// @param generationId The identifier of the generation process
// @param expectedSize Expected size of the generated file, in bytes; 0 if unknown
// @param localPrefixSize The number of bytes already generated
func (client *Client) SetFileGenerationProgress(generationId JSONInt64, expectedSize int32, localPrefixSize int32) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":             "setFileGenerationProgress",
		"generation_id":     generationId,
		"expected_size":     expectedSize,
		"local_prefix_size": localPrefixSize,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// FinishFileGeneration Finishes the file generation
// @param generationId The identifier of the generation process
// @param error If passed, the file generation has failed and must be terminated; pass null if the file generation succeeded
func (client *Client) FinishFileGeneration(generationId JSONInt64, error *Error) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":         "finishFileGeneration",
		"generation_id": generationId,
		"error":         error,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// ReadFilePart Reads a part of a file from the TDLib file cache and returns read bytes. This method is intended to be used only if the application has no direct access to TDLib's file system, because it is usually slower than a direct read from the file
// @param fileId Identifier of the file. The file must be located in the TDLib file cache
// @param offset The offset from which to read the file
// @param count Number of bytes to read. An error will be returned if there are not enough bytes available in the file from the specified position. Pass 0 to read all available data from the specified position
func (client *Client) ReadFilePart(fileId int32, offset int32, count int32) (*FilePart, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "readFilePart",
		"file_id": fileId,
		"offset":  offset,
		"count":   count,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var filePart FilePart
	err = json.Unmarshal(result.Raw, &filePart)
	return &filePart, err

}

// DeleteFile Deletes a file from the TDLib file cache
// @param fileId Identifier of the file to delete
func (client *Client) DeleteFile(fileId int32) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "deleteFile",
		"file_id": fileId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// GetMessageFileType Returns information about a file with messages exported from another app
// @param messageFileHead Beginning of the message file; up to 100 first lines
func (client *Client) GetMessageFileType(messageFileHead string) (MessageFileType, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":             "getMessageFileType",
		"message_file_head": messageFileHead,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	switch MessageFileTypeEnum(result.Data["@type"].(string)) {

	case MessageFileTypePrivateType:
		var messageFileType MessageFileTypePrivate
		err = json.Unmarshal(result.Raw, &messageFileType)
		return &messageFileType, err

	case MessageFileTypeGroupType:
		var messageFileType MessageFileTypeGroup
		err = json.Unmarshal(result.Raw, &messageFileType)
		return &messageFileType, err

	case MessageFileTypeUnknownType:
		var messageFileType MessageFileTypeUnknown
		err = json.Unmarshal(result.Raw, &messageFileType)
		return &messageFileType, err

	default:
		return nil, fmt.Errorf("Invalid type")
	}
}

// GetMessageImportConfirmationText Returns a confirmation text to be shown to the user before starting message import
// @param chatId Identifier of a chat to which the messages will be imported. It must be an identifier of a private chat with a mutual contact or an identifier of a supergroup chat with can_change_info administrator right
func (client *Client) GetMessageImportConfirmationText(chatId int64) (*Text, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "getMessageImportConfirmationText",
		"chat_id": chatId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var text Text
	err = json.Unmarshal(result.Raw, &text)
	return &text, err

}

// ImportMessages Imports messages exported from another app
// @param chatId Identifier of a chat to which the messages will be imported. It must be an identifier of a private chat with a mutual contact or an identifier of a supergroup chat with can_change_info administrator right
// @param messageFile File with messages to import. Only inputFileLocal and inputFileGenerated are supported. The file must not be previously uploaded
// @param attachedFiles Files used in the imported messages. Only inputFileLocal and inputFileGenerated are supported. The files must not be previously uploaded
func (client *Client) ImportMessages(chatId int64, messageFile InputFile, attachedFiles []InputFile) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":          "importMessages",
		"chat_id":        chatId,
		"message_file":   messageFile,
		"attached_files": attachedFiles,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// ReplacePrimaryChatInviteLink Replaces current primary invite link for a chat with a new primary invite link. Available for basic groups, supergroups, and channels. Requires administrator privileges and can_invite_users right
// @param chatId Chat identifier
func (client *Client) ReplacePrimaryChatInviteLink(chatId int64) (*ChatInviteLink, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "replacePrimaryChatInviteLink",
		"chat_id": chatId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var chatInviteLink ChatInviteLink
	err = json.Unmarshal(result.Raw, &chatInviteLink)
	return &chatInviteLink, err

}

// CreateChatInviteLink Creates a new invite link for a chat. Available for basic groups, supergroups, and channels. Requires administrator privileges and can_invite_users right in the chat
// @param chatId Chat identifier
// @param name Invite link name; 0-32 characters
// @param expireDate Point in time (Unix timestamp) when the link will expire; pass 0 if never
// @param memberLimit The maximum number of chat members that can join the chat by the link simultaneously; 0-99999; pass 0 if not limited
// @param createsJoinRequest True, if the link only creates join request. If true, member_limit must not be specified
func (client *Client) CreateChatInviteLink(chatId int64, name string, expireDate int32, memberLimit int32, createsJoinRequest bool) (*ChatInviteLink, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":                "createChatInviteLink",
		"chat_id":              chatId,
		"name":                 name,
		"expire_date":          expireDate,
		"member_limit":         memberLimit,
		"creates_join_request": createsJoinRequest,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var chatInviteLink ChatInviteLink
	err = json.Unmarshal(result.Raw, &chatInviteLink)
	return &chatInviteLink, err

}

// EditChatInviteLink Edits a non-primary invite link for a chat. Available for basic groups, supergroups, and channels. Requires administrator privileges and can_invite_users right in the chat for own links and owner privileges for other links
// @param chatId Chat identifier
// @param inviteLink Invite link to be edited
// @param name Invite link name; 0-32 characters
// @param expireDate Point in time (Unix timestamp) when the link will expire; pass 0 if never
// @param memberLimit The maximum number of chat members that can join the chat by the link simultaneously; 0-99999; pass 0 if not limited
// @param createsJoinRequest True, if the link only creates join request. If true, member_limit must not be specified
func (client *Client) EditChatInviteLink(chatId int64, inviteLink string, name string, expireDate int32, memberLimit int32, createsJoinRequest bool) (*ChatInviteLink, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":                "editChatInviteLink",
		"chat_id":              chatId,
		"invite_link":          inviteLink,
		"name":                 name,
		"expire_date":          expireDate,
		"member_limit":         memberLimit,
		"creates_join_request": createsJoinRequest,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var chatInviteLink ChatInviteLink
	err = json.Unmarshal(result.Raw, &chatInviteLink)
	return &chatInviteLink, err

}

// GetChatInviteLink Returns information about an invite link. Requires administrator privileges and can_invite_users right in the chat to get own links and owner privileges to get other links
// @param chatId Chat identifier
// @param inviteLink Invite link to get
func (client *Client) GetChatInviteLink(chatId int64, inviteLink string) (*ChatInviteLink, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":       "getChatInviteLink",
		"chat_id":     chatId,
		"invite_link": inviteLink,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var chatInviteLink ChatInviteLink
	err = json.Unmarshal(result.Raw, &chatInviteLink)
	return &chatInviteLink, err

}

// GetChatInviteLinkCounts Returns list of chat administrators with number of their invite links. Requires owner privileges in the chat
// @param chatId Chat identifier
func (client *Client) GetChatInviteLinkCounts(chatId int64) (*ChatInviteLinkCounts, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "getChatInviteLinkCounts",
		"chat_id": chatId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var chatInviteLinkCounts ChatInviteLinkCounts
	err = json.Unmarshal(result.Raw, &chatInviteLinkCounts)
	return &chatInviteLinkCounts, err

}

// GetChatInviteLinks Returns invite links for a chat created by specified administrator. Requires administrator privileges and can_invite_users right in the chat to get own links and owner privileges to get other links
// @param chatId Chat identifier
// @param creatorUserId User identifier of a chat administrator. Must be an identifier of the current user for non-owner
// @param isRevoked Pass true if revoked links needs to be returned instead of active or expired
// @param offsetDate Creation date of an invite link starting after which to return invite links; use 0 to get results from the beginning
// @param offsetInviteLink Invite link starting after which to return invite links; use empty string to get results from the beginning
// @param limit The maximum number of invite links to return; up to 100
func (client *Client) GetChatInviteLinks(chatId int64, creatorUserId int64, isRevoked bool, offsetDate int32, offsetInviteLink string, limit int32) (*ChatInviteLinks, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":              "getChatInviteLinks",
		"chat_id":            chatId,
		"creator_user_id":    creatorUserId,
		"is_revoked":         isRevoked,
		"offset_date":        offsetDate,
		"offset_invite_link": offsetInviteLink,
		"limit":              limit,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var chatInviteLinks ChatInviteLinks
	err = json.Unmarshal(result.Raw, &chatInviteLinks)
	return &chatInviteLinks, err

}

// GetChatInviteLinkMembers Returns chat members joined a chat by an invite link. Requires administrator privileges and can_invite_users right in the chat for own links and owner privileges for other links
// @param chatId Chat identifier
// @param inviteLink Invite link for which to return chat members
// @param offsetMember A chat member from which to return next chat members; pass null to get results from the beginning
// @param limit The maximum number of chat members to return; up to 100
func (client *Client) GetChatInviteLinkMembers(chatId int64, inviteLink string, offsetMember *ChatInviteLinkMember, limit int32) (*ChatInviteLinkMembers, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":         "getChatInviteLinkMembers",
		"chat_id":       chatId,
		"invite_link":   inviteLink,
		"offset_member": offsetMember,
		"limit":         limit,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var chatInviteLinkMembers ChatInviteLinkMembers
	err = json.Unmarshal(result.Raw, &chatInviteLinkMembers)
	return &chatInviteLinkMembers, err

}

// RevokeChatInviteLink Revokes invite link for a chat. Available for basic groups, supergroups, and channels. Requires administrator privileges and can_invite_users right in the chat for own links and owner privileges for other links.
// @param chatId Chat identifier
// @param inviteLink Invite link to be revoked
func (client *Client) RevokeChatInviteLink(chatId int64, inviteLink string) (*ChatInviteLinks, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":       "revokeChatInviteLink",
		"chat_id":     chatId,
		"invite_link": inviteLink,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var chatInviteLinks ChatInviteLinks
	err = json.Unmarshal(result.Raw, &chatInviteLinks)
	return &chatInviteLinks, err

}

// DeleteRevokedChatInviteLink Deletes revoked chat invite links. Requires administrator privileges and can_invite_users right in the chat for own links and owner privileges for other links
// @param chatId Chat identifier
// @param inviteLink Invite link to revoke
func (client *Client) DeleteRevokedChatInviteLink(chatId int64, inviteLink string) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":       "deleteRevokedChatInviteLink",
		"chat_id":     chatId,
		"invite_link": inviteLink,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// DeleteAllRevokedChatInviteLinks Deletes all revoked chat invite links created by a given chat administrator. Requires administrator privileges and can_invite_users right in the chat for own links and owner privileges for other links
// @param chatId Chat identifier
// @param creatorUserId User identifier of a chat administrator, which links will be deleted. Must be an identifier of the current user for non-owner
func (client *Client) DeleteAllRevokedChatInviteLinks(chatId int64, creatorUserId int64) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":           "deleteAllRevokedChatInviteLinks",
		"chat_id":         chatId,
		"creator_user_id": creatorUserId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// CheckChatInviteLink Checks the validity of an invite link for a chat and returns information about the corresponding chat
// @param inviteLink Invite link to be checked
func (client *Client) CheckChatInviteLink(inviteLink string) (*ChatInviteLinkInfo, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":       "checkChatInviteLink",
		"invite_link": inviteLink,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var chatInviteLinkInfo ChatInviteLinkInfo
	err = json.Unmarshal(result.Raw, &chatInviteLinkInfo)
	return &chatInviteLinkInfo, err

}

// JoinChatByInviteLink Uses an invite link to add the current user to the chat if possible
// @param inviteLink Invite link to use
func (client *Client) JoinChatByInviteLink(inviteLink string) (*Chat, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":       "joinChatByInviteLink",
		"invite_link": inviteLink,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var chat Chat
	err = json.Unmarshal(result.Raw, &chat)
	return &chat, err

}

// GetChatJoinRequests Returns pending join requests in a chat
// @param chatId Chat identifier
// @param inviteLink Invite link for which to return join requests. If empty, all join requests will be returned. Requires administrator privileges and can_invite_users right in the chat for own links and owner privileges for other links
// @param query A query to search for in the first names, last names and usernames of the users to return
// @param offsetRequest A chat join request from which to return next requests; pass null to get results from the beginning
// @param limit The maximum number of chat join requests to return
func (client *Client) GetChatJoinRequests(chatId int64, inviteLink string, query string, offsetRequest *ChatJoinRequest, limit int32) (*ChatJoinRequests, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":          "getChatJoinRequests",
		"chat_id":        chatId,
		"invite_link":    inviteLink,
		"query":          query,
		"offset_request": offsetRequest,
		"limit":          limit,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var chatJoinRequests ChatJoinRequests
	err = json.Unmarshal(result.Raw, &chatJoinRequests)
	return &chatJoinRequests, err

}

// ApproveChatJoinRequest Approves pending join request in a chat
// @param chatId Chat identifier
// @param userId Identifier of the user, which request will be approved
func (client *Client) ApproveChatJoinRequest(chatId int64, userId int64) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "approveChatJoinRequest",
		"chat_id": chatId,
		"user_id": userId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// DeclineChatJoinRequest Declines pending join request in a chat
// @param chatId Chat identifier
// @param userId Identifier of the user, which request will be declined
func (client *Client) DeclineChatJoinRequest(chatId int64, userId int64) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "declineChatJoinRequest",
		"chat_id": chatId,
		"user_id": userId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// CreateCall Creates a new call
// @param userId Identifier of the user to be called
// @param protocol The call protocols supported by the application
// @param isVideo True, if a video call needs to be created
func (client *Client) CreateCall(userId int64, protocol *CallProtocol, isVideo bool) (*CallId, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":    "createCall",
		"user_id":  userId,
		"protocol": protocol,
		"is_video": isVideo,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var callId CallId
	err = json.Unmarshal(result.Raw, &callId)
	return &callId, err

}

// AcceptCall Accepts an incoming call
// @param callId Call identifier
// @param protocol The call protocols supported by the application
func (client *Client) AcceptCall(callId int32, protocol *CallProtocol) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":    "acceptCall",
		"call_id":  callId,
		"protocol": protocol,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// SendCallSignalingData Sends call signaling data
// @param callId Call identifier
// @param data The data
func (client *Client) SendCallSignalingData(callId int32, data []byte) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "sendCallSignalingData",
		"call_id": callId,
		"data":    data,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// DiscardCall Discards a call
// @param callId Call identifier
// @param isDisconnected True, if the user was disconnected
// @param duration The call duration, in seconds
// @param isVideo True, if the call was a video call
// @param connectionId Identifier of the connection used during the call
func (client *Client) DiscardCall(callId int32, isDisconnected bool, duration int32, isVideo bool, connectionId JSONInt64) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":           "discardCall",
		"call_id":         callId,
		"is_disconnected": isDisconnected,
		"duration":        duration,
		"is_video":        isVideo,
		"connection_id":   connectionId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// SendCallRating Sends a call rating
// @param callId Call identifier
// @param rating Call rating; 1-5
// @param comment An optional user comment if the rating is less than 5
// @param problems List of the exact types of problems with the call, specified by the user
func (client *Client) SendCallRating(callId int32, rating int32, comment string, problems []CallProblem) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":    "sendCallRating",
		"call_id":  callId,
		"rating":   rating,
		"comment":  comment,
		"problems": problems,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// SendCallDebugInformation Sends debug information for a call
// @param callId Call identifier
// @param debugInformation Debug information in application-specific format
func (client *Client) SendCallDebugInformation(callId int32, debugInformation string) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":             "sendCallDebugInformation",
		"call_id":           callId,
		"debug_information": debugInformation,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// GetVideoChatAvailableParticipants Returns list of participant identifiers, which can be used to join video chats in a chat
// @param chatId Chat identifier
func (client *Client) GetVideoChatAvailableParticipants(chatId int64) (*MessageSenders, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "getVideoChatAvailableParticipants",
		"chat_id": chatId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var messageSenders MessageSenders
	err = json.Unmarshal(result.Raw, &messageSenders)
	return &messageSenders, err

}

// SetVideoChatDefaultParticipant Changes default participant identifier, which can be used to join video chats in a chat
// @param chatId Chat identifier
// @param defaultParticipantId Default group call participant identifier to join the video chats
func (client *Client) SetVideoChatDefaultParticipant(chatId int64, defaultParticipantId MessageSender) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":                  "setVideoChatDefaultParticipant",
		"chat_id":                chatId,
		"default_participant_id": defaultParticipantId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// CreateVideoChat Creates a video chat (a group call bound to a chat). Available only for basic groups, supergroups and channels; requires can_manage_video_chats rights
// @param chatId Chat identifier, in which the video chat will be created
// @param title Group call title; if empty, chat title will be used
// @param startDate Point in time (Unix timestamp) when the group call is supposed to be started by an administrator; 0 to start the video chat immediately. The date must be at least 10 seconds and at most 8 days in the future
func (client *Client) CreateVideoChat(chatId int64, title string, startDate int32) (*GroupCallId, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":      "createVideoChat",
		"chat_id":    chatId,
		"title":      title,
		"start_date": startDate,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var groupCallId GroupCallId
	err = json.Unmarshal(result.Raw, &groupCallId)
	return &groupCallId, err

}

// GetGroupCall Returns information about a group call
// @param groupCallId Group call identifier
func (client *Client) GetGroupCall(groupCallId int32) (*GroupCall, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":         "getGroupCall",
		"group_call_id": groupCallId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var groupCallDummy GroupCall
	err = json.Unmarshal(result.Raw, &groupCallDummy)
	return &groupCallDummy, err

}

// StartScheduledGroupCall Starts a scheduled group call
// @param groupCallId Group call identifier
func (client *Client) StartScheduledGroupCall(groupCallId int32) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":         "startScheduledGroupCall",
		"group_call_id": groupCallId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// ToggleGroupCallEnabledStartNotification Toggles whether the current user will receive a notification when the group call will start; scheduled group calls only
// @param groupCallId Group call identifier
// @param enabledStartNotification New value of the enabled_start_notification setting
func (client *Client) ToggleGroupCallEnabledStartNotification(groupCallId int32, enabledStartNotification bool) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":                      "toggleGroupCallEnabledStartNotification",
		"group_call_id":              groupCallId,
		"enabled_start_notification": enabledStartNotification,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// JoinGroupCall Joins an active group call. Returns join response payload for tgcalls
// @param groupCallId Group call identifier
// @param participantId Identifier of a group call participant, which will be used to join the call; pass null to join as self; video chats only
// @param audioSourceId Caller audio channel synchronization source identifier; received from tgcalls
// @param payload Group call join payload; received from tgcalls
// @param isMuted True, if the user's microphone is muted
// @param isMyVideoEnabled True, if the user's video is enabled
// @param inviteHash If non-empty, invite hash to be used to join the group call without being muted by administrators
func (client *Client) JoinGroupCall(groupCallId int32, participantId MessageSender, audioSourceId int32, payload string, isMuted bool, isMyVideoEnabled bool, inviteHash string) (*Text, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":               "joinGroupCall",
		"group_call_id":       groupCallId,
		"participant_id":      participantId,
		"audio_source_id":     audioSourceId,
		"payload":             payload,
		"is_muted":            isMuted,
		"is_my_video_enabled": isMyVideoEnabled,
		"invite_hash":         inviteHash,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var text Text
	err = json.Unmarshal(result.Raw, &text)
	return &text, err

}

// StartGroupCallScreenSharing Starts screen sharing in a joined group call. Returns join response payload for tgcalls
// @param groupCallId Group call identifier
// @param audioSourceId Screen sharing audio channel synchronization source identifier; received from tgcalls
// @param payload Group call join payload; received from tgcalls
func (client *Client) StartGroupCallScreenSharing(groupCallId int32, audioSourceId int32, payload string) (*Text, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":           "startGroupCallScreenSharing",
		"group_call_id":   groupCallId,
		"audio_source_id": audioSourceId,
		"payload":         payload,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var text Text
	err = json.Unmarshal(result.Raw, &text)
	return &text, err

}

// ToggleGroupCallScreenSharingIsPaused Pauses or unpauses screen sharing in a joined group call
// @param groupCallId Group call identifier
// @param isPaused True if screen sharing is paused
func (client *Client) ToggleGroupCallScreenSharingIsPaused(groupCallId int32, isPaused bool) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":         "toggleGroupCallScreenSharingIsPaused",
		"group_call_id": groupCallId,
		"is_paused":     isPaused,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// EndGroupCallScreenSharing Ends screen sharing in a joined group call
// @param groupCallId Group call identifier
func (client *Client) EndGroupCallScreenSharing(groupCallId int32) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":         "endGroupCallScreenSharing",
		"group_call_id": groupCallId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// SetGroupCallTitle Sets group call title. Requires groupCall.can_be_managed group call flag
// @param groupCallId Group call identifier
// @param title New group call title; 1-64 characters
func (client *Client) SetGroupCallTitle(groupCallId int32, title string) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":         "setGroupCallTitle",
		"group_call_id": groupCallId,
		"title":         title,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// ToggleGroupCallMuteNewParticipants Toggles whether new participants of a group call can be unmuted only by administrators of the group call. Requires groupCall.can_toggle_mute_new_participants group call flag
// @param groupCallId Group call identifier
// @param muteNewParticipants New value of the mute_new_participants setting
func (client *Client) ToggleGroupCallMuteNewParticipants(groupCallId int32, muteNewParticipants bool) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":                 "toggleGroupCallMuteNewParticipants",
		"group_call_id":         groupCallId,
		"mute_new_participants": muteNewParticipants,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// RevokeGroupCallInviteLink Revokes invite link for a group call. Requires groupCall.can_be_managed group call flag
// @param groupCallId Group call identifier
func (client *Client) RevokeGroupCallInviteLink(groupCallId int32) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":         "revokeGroupCallInviteLink",
		"group_call_id": groupCallId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// InviteGroupCallParticipants Invites users to an active group call. Sends a service message of type messageInviteToGroupCall for video chats
// @param groupCallId Group call identifier
// @param userIds User identifiers. At most 10 users can be invited simultaneously
func (client *Client) InviteGroupCallParticipants(groupCallId int32, userIds []int64) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":         "inviteGroupCallParticipants",
		"group_call_id": groupCallId,
		"user_ids":      userIds,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// GetGroupCallInviteLink Returns invite link to a video chat in a public chat
// @param groupCallId Group call identifier
// @param canSelfUnmute Pass true if the invite link needs to contain an invite hash, passing which to joinGroupCall would allow the invited user to unmute themselves. Requires groupCall.can_be_managed group call flag
func (client *Client) GetGroupCallInviteLink(groupCallId int32, canSelfUnmute bool) (*HttpUrl, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":           "getGroupCallInviteLink",
		"group_call_id":   groupCallId,
		"can_self_unmute": canSelfUnmute,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var httpUrl HttpUrl
	err = json.Unmarshal(result.Raw, &httpUrl)
	return &httpUrl, err

}

// StartGroupCallRecording Starts recording of an active group call. Requires groupCall.can_be_managed group call flag
// @param groupCallId Group call identifier
// @param title Group call recording title; 0-64 characters
// @param recordVideo Pass true to record a video file instead of an audio file
// @param usePortraitOrientation Pass true to use portrait orientation for video instead of landscape one
func (client *Client) StartGroupCallRecording(groupCallId int32, title string, recordVideo bool, usePortraitOrientation bool) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":                    "startGroupCallRecording",
		"group_call_id":            groupCallId,
		"title":                    title,
		"record_video":             recordVideo,
		"use_portrait_orientation": usePortraitOrientation,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// EndGroupCallRecording Ends recording of an active group call. Requires groupCall.can_be_managed group call flag
// @param groupCallId Group call identifier
func (client *Client) EndGroupCallRecording(groupCallId int32) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":         "endGroupCallRecording",
		"group_call_id": groupCallId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// ToggleGroupCallIsMyVideoPaused Toggles whether current user's video is paused
// @param groupCallId Group call identifier
// @param isMyVideoPaused Pass true if the current user's video is paused
func (client *Client) ToggleGroupCallIsMyVideoPaused(groupCallId int32, isMyVideoPaused bool) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":              "toggleGroupCallIsMyVideoPaused",
		"group_call_id":      groupCallId,
		"is_my_video_paused": isMyVideoPaused,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// ToggleGroupCallIsMyVideoEnabled Toggles whether current user's video is enabled
// @param groupCallId Group call identifier
// @param isMyVideoEnabled Pass true if the current user's video is enabled
func (client *Client) ToggleGroupCallIsMyVideoEnabled(groupCallId int32, isMyVideoEnabled bool) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":               "toggleGroupCallIsMyVideoEnabled",
		"group_call_id":       groupCallId,
		"is_my_video_enabled": isMyVideoEnabled,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// SetGroupCallParticipantIsSpeaking Informs TDLib that speaking state of a participant of an active group has changed
// @param groupCallId Group call identifier
// @param audioSource Group call participant's synchronization audio source identifier, or 0 for the current user
// @param isSpeaking True, if the user is speaking
func (client *Client) SetGroupCallParticipantIsSpeaking(groupCallId int32, audioSource int32, isSpeaking bool) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":         "setGroupCallParticipantIsSpeaking",
		"group_call_id": groupCallId,
		"audio_source":  audioSource,
		"is_speaking":   isSpeaking,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// ToggleGroupCallParticipantIsMuted Toggles whether a participant of an active group call is muted, unmuted, or allowed to unmute themselves
// @param groupCallId Group call identifier
// @param participantId Participant identifier
// @param isMuted Pass true if the user must be muted and false otherwise
func (client *Client) ToggleGroupCallParticipantIsMuted(groupCallId int32, participantId MessageSender, isMuted bool) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":          "toggleGroupCallParticipantIsMuted",
		"group_call_id":  groupCallId,
		"participant_id": participantId,
		"is_muted":       isMuted,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// SetGroupCallParticipantVolumeLevel Changes volume level of a participant of an active group call. If the current user can manage the group call, then the participant's volume level will be changed for all users with the default volume level
// @param groupCallId Group call identifier
// @param participantId Participant identifier
// @param volumeLevel New participant's volume level; 1-20000 in hundreds of percents
func (client *Client) SetGroupCallParticipantVolumeLevel(groupCallId int32, participantId MessageSender, volumeLevel int32) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":          "setGroupCallParticipantVolumeLevel",
		"group_call_id":  groupCallId,
		"participant_id": participantId,
		"volume_level":   volumeLevel,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// ToggleGroupCallParticipantIsHandRaised Toggles whether a group call participant hand is rased
// @param groupCallId Group call identifier
// @param participantId Participant identifier
// @param isHandRaised Pass true if the user's hand needs to be raised. Only self hand can be raised. Requires groupCall.can_be_managed group call flag to lower other's hand
func (client *Client) ToggleGroupCallParticipantIsHandRaised(groupCallId int32, participantId MessageSender, isHandRaised bool) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":          "toggleGroupCallParticipantIsHandRaised",
		"group_call_id":  groupCallId,
		"participant_id": participantId,
		"is_hand_raised": isHandRaised,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// LoadGroupCallParticipants Loads more participants of a group call. The loaded participants will be received through updates. Use the field groupCall.loaded_all_participants to check whether all participants has already been loaded
// @param groupCallId Group call identifier. The group call must be previously received through getGroupCall and must be joined or being joined
// @param limit The maximum number of participants to load; up to 100
func (client *Client) LoadGroupCallParticipants(groupCallId int32, limit int32) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":         "loadGroupCallParticipants",
		"group_call_id": groupCallId,
		"limit":         limit,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// LeaveGroupCall Leaves a group call
// @param groupCallId Group call identifier
func (client *Client) LeaveGroupCall(groupCallId int32) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":         "leaveGroupCall",
		"group_call_id": groupCallId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// DiscardGroupCall Discards a group call. Requires groupCall.can_be_managed
// @param groupCallId Group call identifier
func (client *Client) DiscardGroupCall(groupCallId int32) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":         "discardGroupCall",
		"group_call_id": groupCallId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// GetGroupCallStreamSegment Returns a file with a segment of a group call stream in a modified OGG format for audio or MPEG-4 format for video
// @param groupCallId Group call identifier
// @param timeOffset Point in time when the stream segment begins; Unix timestamp in milliseconds
// @param scale Segment duration scale; 0-1. Segment's duration is 1000/(2**scale) milliseconds
// @param channelId Identifier of an audio/video channel to get as received from tgcalls
// @param videoQuality Video quality as received from tgcalls; pass null to get the worst available quality
func (client *Client) GetGroupCallStreamSegment(groupCallId int32, timeOffset int64, scale int32, channelId int32, videoQuality GroupCallVideoQuality) (*FilePart, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":         "getGroupCallStreamSegment",
		"group_call_id": groupCallId,
		"time_offset":   timeOffset,
		"scale":         scale,
		"channel_id":    channelId,
		"video_quality": videoQuality,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var filePart FilePart
	err = json.Unmarshal(result.Raw, &filePart)
	return &filePart, err

}

// ToggleMessageSenderIsBlocked Changes the block state of a message sender. Currently, only users and supergroup chats can be blocked
// @param sender Message Sender
// @param isBlocked New value of is_blocked
func (client *Client) ToggleMessageSenderIsBlocked(sender MessageSender, isBlocked bool) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":      "toggleMessageSenderIsBlocked",
		"sender":     sender,
		"is_blocked": isBlocked,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// BlockMessageSenderFromReplies Blocks an original sender of a message in the Replies chat
// @param messageId The identifier of an incoming message in the Replies chat
// @param deleteMessage Pass true if the message must be deleted
// @param deleteAllMessages Pass true if all messages from the same sender must be deleted
// @param reportSpam Pass true if the sender must be reported to the Telegram moderators
func (client *Client) BlockMessageSenderFromReplies(messageId int64, deleteMessage bool, deleteAllMessages bool, reportSpam bool) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":               "blockMessageSenderFromReplies",
		"message_id":          messageId,
		"delete_message":      deleteMessage,
		"delete_all_messages": deleteAllMessages,
		"report_spam":         reportSpam,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// GetBlockedMessageSenders Returns users and chats that were blocked by the current user
// @param offset Number of users and chats to skip in the result; must be non-negative
// @param limit The maximum number of users and chats to return; up to 100
func (client *Client) GetBlockedMessageSenders(offset int32, limit int32) (*MessageSenders, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":  "getBlockedMessageSenders",
		"offset": offset,
		"limit":  limit,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var messageSenders MessageSenders
	err = json.Unmarshal(result.Raw, &messageSenders)
	return &messageSenders, err

}

// AddContact Adds a user to the contact list or edits an existing contact by their user identifier
// @param contact The contact to add or edit; phone number can be empty and needs to be specified only if known, vCard is ignored
// @param sharePhoneNumber True, if the new contact needs to be allowed to see current user's phone number. A corresponding rule to userPrivacySettingShowPhoneNumber will be added if needed. Use the field userFullInfo.need_phone_number_privacy_exception to check whether the current user needs to be asked to share their phone number
func (client *Client) AddContact(contact *Contact, sharePhoneNumber bool) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":              "addContact",
		"contact":            contact,
		"share_phone_number": sharePhoneNumber,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// ImportContacts Adds new contacts or edits existing contacts by their phone numbers; contacts' user identifiers are ignored
// @param contacts The list of contacts to import or edit; contacts' vCard are ignored and are not imported
func (client *Client) ImportContacts(contacts []Contact) (*ImportedContacts, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":    "importContacts",
		"contacts": contacts,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var importedContacts ImportedContacts
	err = json.Unmarshal(result.Raw, &importedContacts)
	return &importedContacts, err

}

// GetContacts Returns all user contacts
func (client *Client) GetContacts() (*Users, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "getContacts",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var users Users
	err = json.Unmarshal(result.Raw, &users)
	return &users, err

}

// SearchContacts Searches for the specified query in the first names, last names and usernames of the known user contacts
// @param query Query to search for; may be empty to return all contacts
// @param limit The maximum number of users to be returned
func (client *Client) SearchContacts(query string, limit int32) (*Users, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "searchContacts",
		"query": query,
		"limit": limit,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var users Users
	err = json.Unmarshal(result.Raw, &users)
	return &users, err

}

// RemoveContacts Removes users from the contact list
// @param userIds Identifiers of users to be deleted
func (client *Client) RemoveContacts(userIds []int64) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":    "removeContacts",
		"user_ids": userIds,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// GetImportedContactCount Returns the total number of imported contacts
func (client *Client) GetImportedContactCount() (*Count, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "getImportedContactCount",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var count Count
	err = json.Unmarshal(result.Raw, &count)
	return &count, err

}

// ChangeImportedContacts Changes imported contacts using the list of contacts saved on the device. Imports newly added contacts and, if at least the file database is enabled, deletes recently deleted contacts.
// @param contacts
func (client *Client) ChangeImportedContacts(contacts []Contact) (*ImportedContacts, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":    "changeImportedContacts",
		"contacts": contacts,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var importedContacts ImportedContacts
	err = json.Unmarshal(result.Raw, &importedContacts)
	return &importedContacts, err

}

// ClearImportedContacts Clears all imported contacts, contact list remains unchanged
func (client *Client) ClearImportedContacts() (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "clearImportedContacts",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// SharePhoneNumber Shares the phone number of the current user with a mutual contact. Supposed to be called when the user clicks on chatActionBarSharePhoneNumber
// @param userId Identifier of the user with whom to share the phone number. The user must be a mutual contact
func (client *Client) SharePhoneNumber(userId int64) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "sharePhoneNumber",
		"user_id": userId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// GetUserProfilePhotos Returns the profile photos of a user. The result of this query may be outdated: some photos might have been deleted already
// @param userId User identifier
// @param offset The number of photos to skip; must be non-negative
// @param limit The maximum number of photos to be returned; up to 100
func (client *Client) GetUserProfilePhotos(userId int64, offset int32, limit int32) (*ChatPhotos, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "getUserProfilePhotos",
		"user_id": userId,
		"offset":  offset,
		"limit":   limit,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var chatPhotos ChatPhotos
	err = json.Unmarshal(result.Raw, &chatPhotos)
	return &chatPhotos, err

}

// GetStickers Returns stickers from the installed sticker sets that correspond to a given emoji. If the emoji is non-empty, favorite and recently used stickers may also be returned
// @param emoji String representation of emoji. If empty, returns all known installed stickers
// @param limit The maximum number of stickers to be returned
func (client *Client) GetStickers(emoji string, limit int32) (*Stickers, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "getStickers",
		"emoji": emoji,
		"limit": limit,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var stickers Stickers
	err = json.Unmarshal(result.Raw, &stickers)
	return &stickers, err

}

// SearchStickers Searches for stickers from public sticker sets that correspond to a given emoji
// @param emoji String representation of emoji; must be non-empty
// @param limit The maximum number of stickers to be returned
func (client *Client) SearchStickers(emoji string, limit int32) (*Stickers, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "searchStickers",
		"emoji": emoji,
		"limit": limit,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var stickers Stickers
	err = json.Unmarshal(result.Raw, &stickers)
	return &stickers, err

}

// GetInstalledStickerSets Returns a list of installed sticker sets
// @param isMasks Pass true to return mask sticker sets; pass false to return ordinary sticker sets
func (client *Client) GetInstalledStickerSets(isMasks bool) (*StickerSets, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":    "getInstalledStickerSets",
		"is_masks": isMasks,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var stickerSets StickerSets
	err = json.Unmarshal(result.Raw, &stickerSets)
	return &stickerSets, err

}

// GetArchivedStickerSets Returns a list of archived sticker sets
// @param isMasks Pass true to return mask stickers sets; pass false to return ordinary sticker sets
// @param offsetStickerSetId Identifier of the sticker set from which to return the result
// @param limit The maximum number of sticker sets to return; up to 100
func (client *Client) GetArchivedStickerSets(isMasks bool, offsetStickerSetId JSONInt64, limit int32) (*StickerSets, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":                 "getArchivedStickerSets",
		"is_masks":              isMasks,
		"offset_sticker_set_id": offsetStickerSetId,
		"limit":                 limit,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var stickerSets StickerSets
	err = json.Unmarshal(result.Raw, &stickerSets)
	return &stickerSets, err

}

// GetTrendingStickerSets Returns a list of trending sticker sets. For optimal performance, the number of returned sticker sets is chosen by TDLib
// @param offset The offset from which to return the sticker sets; must be non-negative
// @param limit The maximum number of sticker sets to be returned; up to 100. For optimal performance, the number of returned sticker sets is chosen by TDLib and can be smaller than the specified limit, even if the end of the list has not been reached
func (client *Client) GetTrendingStickerSets(offset int32, limit int32) (*StickerSets, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":  "getTrendingStickerSets",
		"offset": offset,
		"limit":  limit,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var stickerSets StickerSets
	err = json.Unmarshal(result.Raw, &stickerSets)
	return &stickerSets, err

}

// GetAttachedStickerSets Returns a list of sticker sets attached to a file. Currently only photos and videos can have attached sticker sets
// @param fileId File identifier
func (client *Client) GetAttachedStickerSets(fileId int32) (*StickerSets, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "getAttachedStickerSets",
		"file_id": fileId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var stickerSets StickerSets
	err = json.Unmarshal(result.Raw, &stickerSets)
	return &stickerSets, err

}

// GetStickerSet Returns information about a sticker set by its identifier
// @param setId Identifier of the sticker set
func (client *Client) GetStickerSet(setId JSONInt64) (*StickerSet, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":  "getStickerSet",
		"set_id": setId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var stickerSet StickerSet
	err = json.Unmarshal(result.Raw, &stickerSet)
	return &stickerSet, err

}

// SearchStickerSet Searches for a sticker set by its name
// @param name Name of the sticker set
func (client *Client) SearchStickerSet(name string) (*StickerSet, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "searchStickerSet",
		"name":  name,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var stickerSet StickerSet
	err = json.Unmarshal(result.Raw, &stickerSet)
	return &stickerSet, err

}

// SearchInstalledStickerSets Searches for installed sticker sets by looking for specified query in their title and name
// @param isMasks Pass true to return mask sticker sets; pass false to return ordinary sticker sets
// @param query Query to search for
// @param limit The maximum number of sticker sets to return
func (client *Client) SearchInstalledStickerSets(isMasks bool, query string, limit int32) (*StickerSets, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":    "searchInstalledStickerSets",
		"is_masks": isMasks,
		"query":    query,
		"limit":    limit,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var stickerSets StickerSets
	err = json.Unmarshal(result.Raw, &stickerSets)
	return &stickerSets, err

}

// SearchStickerSets Searches for ordinary sticker sets by looking for specified query in their title and name. Excludes installed sticker sets from the results
// @param query Query to search for
func (client *Client) SearchStickerSets(query string) (*StickerSets, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "searchStickerSets",
		"query": query,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var stickerSets StickerSets
	err = json.Unmarshal(result.Raw, &stickerSets)
	return &stickerSets, err

}

// ChangeStickerSet Installs/uninstalls or activates/archives a sticker set
// @param setId Identifier of the sticker set
// @param isInstalled The new value of is_installed
// @param isArchived The new value of is_archived. A sticker set can't be installed and archived simultaneously
func (client *Client) ChangeStickerSet(setId JSONInt64, isInstalled bool, isArchived bool) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":        "changeStickerSet",
		"set_id":       setId,
		"is_installed": isInstalled,
		"is_archived":  isArchived,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// ViewTrendingStickerSets Informs the server that some trending sticker sets have been viewed by the user
// @param stickerSetIds Identifiers of viewed trending sticker sets
func (client *Client) ViewTrendingStickerSets(stickerSetIds []JSONInt64) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":           "viewTrendingStickerSets",
		"sticker_set_ids": stickerSetIds,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// ReorderInstalledStickerSets Changes the order of installed sticker sets
// @param isMasks Pass true to change the order of mask sticker sets; pass false to change the order of ordinary sticker sets
// @param stickerSetIds Identifiers of installed sticker sets in the new correct order
func (client *Client) ReorderInstalledStickerSets(isMasks bool, stickerSetIds []JSONInt64) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":           "reorderInstalledStickerSets",
		"is_masks":        isMasks,
		"sticker_set_ids": stickerSetIds,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// GetRecentStickers Returns a list of recently used stickers
// @param isAttached Pass true to return stickers and masks that were recently attached to photos or video files; pass false to return recently sent stickers
func (client *Client) GetRecentStickers(isAttached bool) (*Stickers, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":       "getRecentStickers",
		"is_attached": isAttached,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var stickers Stickers
	err = json.Unmarshal(result.Raw, &stickers)
	return &stickers, err

}

// AddRecentSticker Manually adds a new sticker to the list of recently used stickers. The new sticker is added to the top of the list. If the sticker was already in the list, it is removed from the list first. Only stickers belonging to a sticker set can be added to this list
// @param isAttached Pass true to add the sticker to the list of stickers recently attached to photo or video files; pass false to add the sticker to the list of recently sent stickers
// @param sticker Sticker file to add
func (client *Client) AddRecentSticker(isAttached bool, sticker InputFile) (*Stickers, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":       "addRecentSticker",
		"is_attached": isAttached,
		"sticker":     sticker,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var stickers Stickers
	err = json.Unmarshal(result.Raw, &stickers)
	return &stickers, err

}

// RemoveRecentSticker Removes a sticker from the list of recently used stickers
// @param isAttached Pass true to remove the sticker from the list of stickers recently attached to photo or video files; pass false to remove the sticker from the list of recently sent stickers
// @param sticker Sticker file to delete
func (client *Client) RemoveRecentSticker(isAttached bool, sticker InputFile) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":       "removeRecentSticker",
		"is_attached": isAttached,
		"sticker":     sticker,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// ClearRecentStickers Clears the list of recently used stickers
// @param isAttached Pass true to clear the list of stickers recently attached to photo or video files; pass false to clear the list of recently sent stickers
func (client *Client) ClearRecentStickers(isAttached bool) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":       "clearRecentStickers",
		"is_attached": isAttached,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// GetFavoriteStickers Returns favorite stickers
func (client *Client) GetFavoriteStickers() (*Stickers, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "getFavoriteStickers",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var stickers Stickers
	err = json.Unmarshal(result.Raw, &stickers)
	return &stickers, err

}

// AddFavoriteSticker Adds a new sticker to the list of favorite stickers. The new sticker is added to the top of the list. If the sticker was already in the list, it is removed from the list first. Only stickers belonging to a sticker set can be added to this list
// @param sticker Sticker file to add
func (client *Client) AddFavoriteSticker(sticker InputFile) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "addFavoriteSticker",
		"sticker": sticker,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// RemoveFavoriteSticker Removes a sticker from the list of favorite stickers
// @param sticker Sticker file to delete from the list
func (client *Client) RemoveFavoriteSticker(sticker InputFile) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "removeFavoriteSticker",
		"sticker": sticker,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// GetStickerEmojis Returns emoji corresponding to a sticker. The list is only for informational purposes, because a sticker is always sent with a fixed emoji from the corresponding Sticker object
// @param sticker Sticker file identifier
func (client *Client) GetStickerEmojis(sticker InputFile) (*Emojis, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "getStickerEmojis",
		"sticker": sticker,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var emojis Emojis
	err = json.Unmarshal(result.Raw, &emojis)
	return &emojis, err

}

// SearchEmojis Searches for emojis by keywords. Supported only if the file database is enabled
// @param text Text to search for
// @param exactMatch True, if only emojis, which exactly match text needs to be returned
// @param inputLanguageCodes List of possible IETF language tags of the user's input language; may be empty if unknown
func (client *Client) SearchEmojis(text string, exactMatch bool, inputLanguageCodes []string) (*Emojis, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":                "searchEmojis",
		"text":                 text,
		"exact_match":          exactMatch,
		"input_language_codes": inputLanguageCodes,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var emojis Emojis
	err = json.Unmarshal(result.Raw, &emojis)
	return &emojis, err

}

// GetAnimatedEmoji Returns an animated emoji corresponding to a given emoji. Returns a 404 error if the emoji has no animated emoji
// @param emoji The emoji
func (client *Client) GetAnimatedEmoji(emoji string) (*AnimatedEmoji, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "getAnimatedEmoji",
		"emoji": emoji,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var animatedEmoji AnimatedEmoji
	err = json.Unmarshal(result.Raw, &animatedEmoji)
	return &animatedEmoji, err

}

// GetEmojiSuggestionsUrl Returns an HTTP URL which can be used to automatically log in to the translation platform and suggest new emoji replacements. The URL will be valid for 30 seconds after generation
// @param languageCode Language code for which the emoji replacements will be suggested
func (client *Client) GetEmojiSuggestionsUrl(languageCode string) (*HttpUrl, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":         "getEmojiSuggestionsUrl",
		"language_code": languageCode,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var httpUrl HttpUrl
	err = json.Unmarshal(result.Raw, &httpUrl)
	return &httpUrl, err

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

// AddSavedAnimation Manually adds a new animation to the list of saved animations. The new animation is added to the beginning of the list. If the animation was already in the list, it is removed first. Only non-secret video animations with MIME type "video/mp4" can be added to the list
// @param animation The animation file to be added. Only animations known to the server (i.e., successfully sent via a message) can be added to the list
func (client *Client) AddSavedAnimation(animation InputFile) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":     "addSavedAnimation",
		"animation": animation,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// RemoveSavedAnimation Removes an animation from the list of saved animations
// @param animation Animation file to be removed
func (client *Client) RemoveSavedAnimation(animation InputFile) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":     "removeSavedAnimation",
		"animation": animation,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// GetRecentInlineBots Returns up to 20 recently used inline bots in the order of their last usage
func (client *Client) GetRecentInlineBots() (*Users, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "getRecentInlineBots",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var users Users
	err = json.Unmarshal(result.Raw, &users)
	return &users, err

}

// SearchHashtags Searches for recently used hashtags by their prefix
// @param prefix Hashtag prefix to search for
// @param limit The maximum number of hashtags to be returned
func (client *Client) SearchHashtags(prefix string, limit int32) (*Hashtags, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":  "searchHashtags",
		"prefix": prefix,
		"limit":  limit,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var hashtags Hashtags
	err = json.Unmarshal(result.Raw, &hashtags)
	return &hashtags, err

}

// RemoveRecentHashtag Removes a hashtag from the list of recently used hashtags
// @param hashtag Hashtag to delete
func (client *Client) RemoveRecentHashtag(hashtag string) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "removeRecentHashtag",
		"hashtag": hashtag,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// GetWebPagePreview Returns a web page preview by the text of the message. Do not call this function too often. Returns a 404 error if the web page has no preview
// @param text Message text with formatting
func (client *Client) GetWebPagePreview(text *FormattedText) (*WebPage, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "getWebPagePreview",
		"text":  text,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var webPage WebPage
	err = json.Unmarshal(result.Raw, &webPage)
	return &webPage, err

}

// GetWebPageInstantView Returns an instant view version of a web page if available. Returns a 404 error if the web page has no instant view page
// @param url The web page URL
// @param forceFull If true, the full instant view for the web page will be returned
func (client *Client) GetWebPageInstantView(url string, forceFull bool) (*WebPageInstantView, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":      "getWebPageInstantView",
		"url":        url,
		"force_full": forceFull,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var webPageInstantView WebPageInstantView
	err = json.Unmarshal(result.Raw, &webPageInstantView)
	return &webPageInstantView, err

}

// SetProfilePhoto Changes a profile photo for the current user
// @param photo Profile photo to set
func (client *Client) SetProfilePhoto(photo InputChatPhoto) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "setProfilePhoto",
		"photo": photo,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// DeleteProfilePhoto Deletes a profile photo
// @param profilePhotoId Identifier of the profile photo to delete
func (client *Client) DeleteProfilePhoto(profilePhotoId JSONInt64) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":            "deleteProfilePhoto",
		"profile_photo_id": profilePhotoId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// SetName Changes the first and last name of the current user
// @param firstName The new value of the first name for the current user; 1-64 characters
// @param lastName The new value of the optional last name for the current user; 0-64 characters
func (client *Client) SetName(firstName string, lastName string) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":      "setName",
		"first_name": firstName,
		"last_name":  lastName,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// SetBio Changes the bio of the current user
// @param bio The new value of the user bio; 0-70 characters without line feeds
func (client *Client) SetBio(bio string) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "setBio",
		"bio":   bio,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// SetUsername Changes the username of the current user
// @param username The new value of the username. Use an empty string to remove the username
func (client *Client) SetUsername(username string) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":    "setUsername",
		"username": username,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// SetLocation Changes the location of the current user. Needs to be called if GetOption("is_location_visible") is true and location changes for more than 1 kilometer
// @param location The new location of the user
func (client *Client) SetLocation(location *Location) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":    "setLocation",
		"location": location,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// ChangePhoneNumber Changes the phone number of the user and sends an authentication code to the user's new phone number. On success, returns information about the sent code
// @param phoneNumber The new phone number of the user in international format
// @param settings Settings for the authentication of the user's phone number; pass null to use default settings
func (client *Client) ChangePhoneNumber(phoneNumber string, settings *PhoneNumberAuthenticationSettings) (*AuthenticationCodeInfo, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":        "changePhoneNumber",
		"phone_number": phoneNumber,
		"settings":     settings,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var authenticationCodeInfo AuthenticationCodeInfo
	err = json.Unmarshal(result.Raw, &authenticationCodeInfo)
	return &authenticationCodeInfo, err

}

// ResendChangePhoneNumberCode Re-sends the authentication code sent to confirm a new phone number for the current user. Works only if the previously received authenticationCodeInfo next_code_type was not null and the server-specified timeout has passed
func (client *Client) ResendChangePhoneNumberCode() (*AuthenticationCodeInfo, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "resendChangePhoneNumberCode",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var authenticationCodeInfo AuthenticationCodeInfo
	err = json.Unmarshal(result.Raw, &authenticationCodeInfo)
	return &authenticationCodeInfo, err

}

// CheckChangePhoneNumberCode Checks the authentication code sent to confirm a new phone number of the user
// @param code Verification code received by SMS, phone call or flash call
func (client *Client) CheckChangePhoneNumberCode(code string) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "checkChangePhoneNumberCode",
		"code":  code,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// SetCommands Sets the list of commands supported by the bot for the given user scope and language; for bots only
// @param scope The scope to which the commands are relevant; pass null to change commands in the default bot command scope
// @param languageCode A two-letter ISO 639-1 country code. If empty, the commands will be applied to all users from the given scope, for which language there are no dedicated commands
// @param commands List of the bot's commands
func (client *Client) SetCommands(scope BotCommandScope, languageCode string, commands []BotCommand) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":         "setCommands",
		"scope":         scope,
		"language_code": languageCode,
		"commands":      commands,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// DeleteCommands Deletes commands supported by the bot for the given user scope and language; for bots only
// @param scope The scope to which the commands are relevant; pass null to delete commands in the default bot command scope
// @param languageCode A two-letter ISO 639-1 country code or an empty string
func (client *Client) DeleteCommands(scope BotCommandScope, languageCode string) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":         "deleteCommands",
		"scope":         scope,
		"language_code": languageCode,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// GetCommands Returns the list of commands supported by the bot for the given user scope and language; for bots only
// @param scope The scope to which the commands are relevant; pass null to get commands in the default bot command scope
// @param languageCode A two-letter ISO 639-1 country code or an empty string
func (client *Client) GetCommands(scope BotCommandScope, languageCode string) (*BotCommands, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":         "getCommands",
		"scope":         scope,
		"language_code": languageCode,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var botCommands BotCommands
	err = json.Unmarshal(result.Raw, &botCommands)
	return &botCommands, err

}

// GetActiveSessions Returns all active sessions of the current user
func (client *Client) GetActiveSessions() (*Sessions, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "getActiveSessions",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var sessions Sessions
	err = json.Unmarshal(result.Raw, &sessions)
	return &sessions, err

}

// TerminateSession Terminates a session of the current user
// @param sessionId Session identifier
func (client *Client) TerminateSession(sessionId JSONInt64) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":      "terminateSession",
		"session_id": sessionId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// TerminateAllOtherSessions Terminates all other sessions of the current user
func (client *Client) TerminateAllOtherSessions() (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "terminateAllOtherSessions",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// GetConnectedWebsites Returns all website where the current user used Telegram to log in
func (client *Client) GetConnectedWebsites() (*ConnectedWebsites, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "getConnectedWebsites",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var connectedWebsites ConnectedWebsites
	err = json.Unmarshal(result.Raw, &connectedWebsites)
	return &connectedWebsites, err

}

// DisconnectWebsite Disconnects website from the current user's Telegram account
// @param websiteId Website identifier
func (client *Client) DisconnectWebsite(websiteId JSONInt64) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":      "disconnectWebsite",
		"website_id": websiteId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// DisconnectAllWebsites Disconnects all websites from the current user's Telegram account
func (client *Client) DisconnectAllWebsites() (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "disconnectAllWebsites",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// SetSupergroupUsername Changes the username of a supergroup or channel, requires owner privileges in the supergroup or channel
// @param supergroupId Identifier of the supergroup or channel
// @param username New value of the username. Use an empty string to remove the username
func (client *Client) SetSupergroupUsername(supergroupId int64, username string) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":         "setSupergroupUsername",
		"supergroup_id": supergroupId,
		"username":      username,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// SetSupergroupStickerSet Changes the sticker set of a supergroup; requires can_change_info administrator right
// @param supergroupId Identifier of the supergroup
// @param stickerSetId New value of the supergroup sticker set identifier. Use 0 to remove the supergroup sticker set
func (client *Client) SetSupergroupStickerSet(supergroupId int64, stickerSetId JSONInt64) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":          "setSupergroupStickerSet",
		"supergroup_id":  supergroupId,
		"sticker_set_id": stickerSetId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// ToggleSupergroupSignMessages Toggles sender signatures messages sent in a channel; requires can_change_info administrator right
// @param supergroupId Identifier of the channel
// @param signMessages New value of sign_messages
func (client *Client) ToggleSupergroupSignMessages(supergroupId int64, signMessages bool) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":         "toggleSupergroupSignMessages",
		"supergroup_id": supergroupId,
		"sign_messages": signMessages,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// ToggleSupergroupIsAllHistoryAvailable Toggles whether the message history of a supergroup is available to new members; requires can_change_info administrator right
// @param supergroupId The identifier of the supergroup
// @param isAllHistoryAvailable The new value of is_all_history_available
func (client *Client) ToggleSupergroupIsAllHistoryAvailable(supergroupId int64, isAllHistoryAvailable bool) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":                    "toggleSupergroupIsAllHistoryAvailable",
		"supergroup_id":            supergroupId,
		"is_all_history_available": isAllHistoryAvailable,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// ToggleSupergroupIsBroadcastGroup Upgrades supergroup to a broadcast group; requires owner privileges in the supergroup
// @param supergroupId Identifier of the supergroup
func (client *Client) ToggleSupergroupIsBroadcastGroup(supergroupId int64) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":         "toggleSupergroupIsBroadcastGroup",
		"supergroup_id": supergroupId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// ReportSupergroupSpam Reports some messages from a user in a supergroup as spam; requires administrator rights in the supergroup
// @param supergroupId Supergroup identifier
// @param userId User identifier
// @param messageIds Identifiers of messages sent in the supergroup by the user. This list must be non-empty
func (client *Client) ReportSupergroupSpam(supergroupId int64, userId int64, messageIds []int64) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":         "reportSupergroupSpam",
		"supergroup_id": supergroupId,
		"user_id":       userId,
		"message_ids":   messageIds,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// GetSupergroupMembers Returns information about members or banned users in a supergroup or channel. Can be used only if supergroupFullInfo.can_get_members == true; additionally, administrator privileges may be required for some filters
// @param supergroupId Identifier of the supergroup or channel
// @param filter The type of users to return; pass null to use supergroupMembersFilterRecent
// @param offset Number of users to skip
// @param limit The maximum number of users be returned; up to 200
func (client *Client) GetSupergroupMembers(supergroupId int64, filter SupergroupMembersFilter, offset int32, limit int32) (*ChatMembers, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":         "getSupergroupMembers",
		"supergroup_id": supergroupId,
		"filter":        filter,
		"offset":        offset,
		"limit":         limit,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var chatMembers ChatMembers
	err = json.Unmarshal(result.Raw, &chatMembers)
	return &chatMembers, err

}

// CloseSecretChat Closes a secret chat, effectively transferring its state to secretChatStateClosed
// @param secretChatId Secret chat identifier
func (client *Client) CloseSecretChat(secretChatId int32) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":          "closeSecretChat",
		"secret_chat_id": secretChatId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// GetChatEventLog Returns a list of service actions taken by chat members and administrators in the last 48 hours. Available only for supergroups and channels. Requires administrator rights. Returns results in reverse chronological order (i. e., in order of decreasing event_id)
// @param chatId Chat identifier
// @param query Search query by which to filter events
// @param fromEventId Identifier of an event from which to return results. Use 0 to get results from the latest events
// @param limit The maximum number of events to return; up to 100
// @param filters The types of events to return; pass null to get chat events of all types
// @param userIds User identifiers by which to filter events. By default, events relating to all users will be returned
func (client *Client) GetChatEventLog(chatId int64, query string, fromEventId JSONInt64, limit int32, filters *ChatEventLogFilters, userIds []int64) (*ChatEvents, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":         "getChatEventLog",
		"chat_id":       chatId,
		"query":         query,
		"from_event_id": fromEventId,
		"limit":         limit,
		"filters":       filters,
		"user_ids":      userIds,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var chatEvents ChatEvents
	err = json.Unmarshal(result.Raw, &chatEvents)
	return &chatEvents, err

}

// GetPaymentForm Returns an invoice payment form. This method must be called when the user presses inlineKeyboardButtonBuy
// @param chatId Chat identifier of the Invoice message
// @param messageId Message identifier
// @param theme Preferred payment form theme; pass null to use the default theme
func (client *Client) GetPaymentForm(chatId int64, messageId int64, theme *PaymentFormTheme) (*PaymentForm, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":      "getPaymentForm",
		"chat_id":    chatId,
		"message_id": messageId,
		"theme":      theme,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var paymentForm PaymentForm
	err = json.Unmarshal(result.Raw, &paymentForm)
	return &paymentForm, err

}

// ValidateOrderInfo Validates the order information provided by a user and returns the available shipping options for a flexible invoice
// @param chatId Chat identifier of the Invoice message
// @param messageId Message identifier
// @param orderInfo The order information, provided by the user; pass null if empty
// @param allowSave True, if the order information can be saved
func (client *Client) ValidateOrderInfo(chatId int64, messageId int64, orderInfo *OrderInfo, allowSave bool) (*ValidatedOrderInfo, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":      "validateOrderInfo",
		"chat_id":    chatId,
		"message_id": messageId,
		"order_info": orderInfo,
		"allow_save": allowSave,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var validatedOrderInfo ValidatedOrderInfo
	err = json.Unmarshal(result.Raw, &validatedOrderInfo)
	return &validatedOrderInfo, err

}

// SendPaymentForm Sends a filled-out payment form to the bot for final verification
// @param chatId Chat identifier of the Invoice message
// @param messageId Message identifier
// @param paymentFormId Payment form identifier returned by getPaymentForm
// @param orderInfoId Identifier returned by validateOrderInfo, or an empty string
// @param shippingOptionId Identifier of a chosen shipping option, if applicable
// @param credentials The credentials chosen by user for payment
// @param tipAmount Chosen by the user amount of tip in the smallest units of the currency
func (client *Client) SendPaymentForm(chatId int64, messageId int64, paymentFormId JSONInt64, orderInfoId string, shippingOptionId string, credentials InputCredentials, tipAmount int64) (*PaymentResult, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":              "sendPaymentForm",
		"chat_id":            chatId,
		"message_id":         messageId,
		"payment_form_id":    paymentFormId,
		"order_info_id":      orderInfoId,
		"shipping_option_id": shippingOptionId,
		"credentials":        credentials,
		"tip_amount":         tipAmount,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var paymentResult PaymentResult
	err = json.Unmarshal(result.Raw, &paymentResult)
	return &paymentResult, err

}

// GetPaymentReceipt Returns information about a successful payment
// @param chatId Chat identifier of the PaymentSuccessful message
// @param messageId Message identifier
func (client *Client) GetPaymentReceipt(chatId int64, messageId int64) (*PaymentReceipt, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":      "getPaymentReceipt",
		"chat_id":    chatId,
		"message_id": messageId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var paymentReceipt PaymentReceipt
	err = json.Unmarshal(result.Raw, &paymentReceipt)
	return &paymentReceipt, err

}

// GetSavedOrderInfo Returns saved order info, if any
func (client *Client) GetSavedOrderInfo() (*OrderInfo, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "getSavedOrderInfo",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var orderInfo OrderInfo
	err = json.Unmarshal(result.Raw, &orderInfo)
	return &orderInfo, err

}

// DeleteSavedOrderInfo Deletes saved order info
func (client *Client) DeleteSavedOrderInfo() (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "deleteSavedOrderInfo",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// DeleteSavedCredentials Deletes saved credentials for all payment provider bots
func (client *Client) DeleteSavedCredentials() (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "deleteSavedCredentials",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// GetSupportUser Returns a user that can be contacted to get support
func (client *Client) GetSupportUser() (*User, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "getSupportUser",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var user User
	err = json.Unmarshal(result.Raw, &user)
	return &user, err

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

// GetBackgroundUrl Constructs a persistent HTTP URL for a background
// @param name Background name
// @param typeParam Background type
func (client *Client) GetBackgroundUrl(name string, typeParam BackgroundType) (*HttpUrl, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "getBackgroundUrl",
		"name":  name,
		"type":  typeParam,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var httpUrl HttpUrl
	err = json.Unmarshal(result.Raw, &httpUrl)
	return &httpUrl, err

}

// SearchBackground Searches for a background by its name
// @param name The name of the background
func (client *Client) SearchBackground(name string) (*Background, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "searchBackground",
		"name":  name,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var background Background
	err = json.Unmarshal(result.Raw, &background)
	return &background, err

}

// SetBackground Changes the background selected by the user; adds background to the list of installed backgrounds
// @param background The input background to use; pass null to create a new filled backgrounds or to remove the current background
// @param typeParam Background type; pass null to use the default type of the remote background or to remove the current background
// @param forDarkTheme True, if the background is chosen for dark theme
func (client *Client) SetBackground(background InputBackground, typeParam BackgroundType, forDarkTheme bool) (*Background, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":          "setBackground",
		"background":     background,
		"type":           typeParam,
		"for_dark_theme": forDarkTheme,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var backgroundDummy Background
	err = json.Unmarshal(result.Raw, &backgroundDummy)
	return &backgroundDummy, err

}

// RemoveBackground Removes background from the list of installed backgrounds
// @param backgroundId The background identifier
func (client *Client) RemoveBackground(backgroundId JSONInt64) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":         "removeBackground",
		"background_id": backgroundId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// ResetBackgrounds Resets list of installed backgrounds to its default value
func (client *Client) ResetBackgrounds() (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "resetBackgrounds",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

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

// GetLanguagePackStrings Returns strings from a language pack in the current localization target by their keys. Can be called before authorization
// @param languagePackId Language pack identifier of the strings to be returned
// @param keys Language pack keys of the strings to be returned; leave empty to request all available strings
func (client *Client) GetLanguagePackStrings(languagePackId string, keys []string) (*LanguagePackStrings, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":            "getLanguagePackStrings",
		"language_pack_id": languagePackId,
		"keys":             keys,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var languagePackStrings LanguagePackStrings
	err = json.Unmarshal(result.Raw, &languagePackStrings)
	return &languagePackStrings, err

}

// SynchronizeLanguagePack Fetches the latest versions of all strings from a language pack in the current localization target from the server. This method doesn't need to be called explicitly for the current used/base language packs. Can be called before authorization
// @param languagePackId Language pack identifier
func (client *Client) SynchronizeLanguagePack(languagePackId string) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":            "synchronizeLanguagePack",
		"language_pack_id": languagePackId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// AddCustomServerLanguagePack Adds a custom server language pack to the list of installed language packs in current localization target. Can be called before authorization
// @param languagePackId Identifier of a language pack to be added; may be different from a name that is used in an "https://t.me/setlanguage/" link
func (client *Client) AddCustomServerLanguagePack(languagePackId string) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":            "addCustomServerLanguagePack",
		"language_pack_id": languagePackId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// SetCustomLanguagePack Adds or changes a custom local language pack to the current localization target
// @param info Information about the language pack. Language pack ID must start with 'X', consist only of English letters, digits and hyphens, and must not exceed 64 characters. Can be called before authorization
// @param strings Strings of the new language pack
func (client *Client) SetCustomLanguagePack(info *LanguagePackInfo, strings []LanguagePackString) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "setCustomLanguagePack",
		"info":    info,
		"strings": strings,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// EditCustomLanguagePackInfo Edits information about a custom local language pack in the current localization target. Can be called before authorization
// @param info New information about the custom local language pack
func (client *Client) EditCustomLanguagePackInfo(info *LanguagePackInfo) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "editCustomLanguagePackInfo",
		"info":  info,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// SetCustomLanguagePackString Adds, edits or deletes a string in a custom local language pack. Can be called before authorization
// @param languagePackId Identifier of a previously added custom local language pack in the current localization target
// @param newString New language pack string
func (client *Client) SetCustomLanguagePackString(languagePackId string, newString *LanguagePackString) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":            "setCustomLanguagePackString",
		"language_pack_id": languagePackId,
		"new_string":       newString,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// DeleteLanguagePack Deletes all information about a language pack in the current localization target. The language pack which is currently in use (including base language pack) or is being synchronized can't be deleted. Can be called before authorization
// @param languagePackId Identifier of the language pack to delete
func (client *Client) DeleteLanguagePack(languagePackId string) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":            "deleteLanguagePack",
		"language_pack_id": languagePackId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// RegisterDevice Registers the currently used device for receiving push notifications. Returns a globally unique identifier of the push notification subscription
// @param deviceToken Device token
// @param otherUserIds List of user identifiers of other users currently using the application
func (client *Client) RegisterDevice(deviceToken DeviceToken, otherUserIds []int64) (*PushReceiverId, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":          "registerDevice",
		"device_token":   deviceToken,
		"other_user_ids": otherUserIds,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var pushReceiverId PushReceiverId
	err = json.Unmarshal(result.Raw, &pushReceiverId)
	return &pushReceiverId, err

}

// ProcessPushNotification Handles a push notification. Returns error with code 406 if the push notification is not supported and connection to the server is required to fetch new data. Can be called before authorization
// @param payload JSON-encoded push notification payload with all fields sent by the server, and "google.sent_time" and "google.notification.sound" fields added
func (client *Client) ProcessPushNotification(payload string) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "processPushNotification",
		"payload": payload,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// GetPushReceiverId Returns a globally unique push notification subscription identifier for identification of an account, which has received a push notification. Can be called synchronously
// @param payload JSON-encoded push notification payload
func (client *Client) GetPushReceiverId(payload string) (*PushReceiverId, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "getPushReceiverId",
		"payload": payload,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var pushReceiverId PushReceiverId
	err = json.Unmarshal(result.Raw, &pushReceiverId)
	return &pushReceiverId, err

}

// GetRecentlyVisitedTMeUrls Returns t.me URLs recently visited by a newly registered user
// @param referrer Google Play referrer to identify the user
func (client *Client) GetRecentlyVisitedTMeUrls(referrer string) (*TMeUrls, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":    "getRecentlyVisitedTMeUrls",
		"referrer": referrer,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var tMeUrls TMeUrls
	err = json.Unmarshal(result.Raw, &tMeUrls)
	return &tMeUrls, err

}

// SetUserPrivacySettingRules Changes user privacy settings
// @param setting The privacy setting
// @param rules The new privacy rules
func (client *Client) SetUserPrivacySettingRules(setting UserPrivacySetting, rules *UserPrivacySettingRules) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "setUserPrivacySettingRules",
		"setting": setting,
		"rules":   rules,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

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

// GetOption Returns the value of an option by its name. (Check the list of available options on https://core.telegram.org/tdlib/options.) Can be called before authorization
// @param name The name of the option
func (client *Client) GetOption(name string) (OptionValue, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "getOption",
		"name":  name,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	switch OptionValueEnum(result.Data["@type"].(string)) {

	case OptionValueBooleanType:
		var optionValue OptionValueBoolean
		err = json.Unmarshal(result.Raw, &optionValue)
		return &optionValue, err

	case OptionValueEmptyType:
		var optionValue OptionValueEmpty
		err = json.Unmarshal(result.Raw, &optionValue)
		return &optionValue, err

	case OptionValueIntegerType:
		var optionValue OptionValueInteger
		err = json.Unmarshal(result.Raw, &optionValue)
		return &optionValue, err

	case OptionValueStringType:
		var optionValue OptionValueString
		err = json.Unmarshal(result.Raw, &optionValue)
		return &optionValue, err

	default:
		return nil, fmt.Errorf("Invalid type")
	}
}

// SetOption Sets the value of an option. (Check the list of available options on https://core.telegram.org/tdlib/options.) Only writable options can be set. Can be called before authorization
// @param name The name of the option
// @param value The new value of the option; pass null to reset option value to a default value
func (client *Client) SetOption(name string, value OptionValue) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "setOption",
		"name":  name,
		"value": value,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// SetAccountTtl Changes the period of inactivity after which the account of the current user will automatically be deleted
// @param ttl New account TTL
func (client *Client) SetAccountTtl(ttl *AccountTtl) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "setAccountTtl",
		"ttl":   ttl,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// GetAccountTtl Returns the period of inactivity after which the account of the current user will automatically be deleted
func (client *Client) GetAccountTtl() (*AccountTtl, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "getAccountTtl",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var accountTtl AccountTtl
	err = json.Unmarshal(result.Raw, &accountTtl)
	return &accountTtl, err

}

// DeleteAccount Deletes the account of the current user, deleting all information associated with the user from the server. The phone number of the account can be used to create a new account. Can be called before authorization when the current authorization state is authorizationStateWaitPassword
// @param reason The reason why the account was deleted; optional
func (client *Client) DeleteAccount(reason string) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":  "deleteAccount",
		"reason": reason,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// RemoveChatActionBar Removes a chat action bar without any other action
// @param chatId Chat identifier
func (client *Client) RemoveChatActionBar(chatId int64) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "removeChatActionBar",
		"chat_id": chatId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// ReportChat Reports a chat to the Telegram moderators. A chat can be reported only from the chat action bar, or if this is a private chat with a bot, a private chat with a user sharing their location, a supergroup, or a channel, since other chats can't be checked by moderators
// @param chatId Chat identifier
// @param messageIds Identifiers of reported messages, if any
// @param reason The reason for reporting the chat
// @param text Additional report details; 0-1024 characters
func (client *Client) ReportChat(chatId int64, messageIds []int64, reason ChatReportReason, text string) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":       "reportChat",
		"chat_id":     chatId,
		"message_ids": messageIds,
		"reason":      reason,
		"text":        text,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// ReportChatPhoto Reports a chat photo to the Telegram moderators. A chat photo can be reported only if this is a private chat with a bot, a private chat with a user sharing their location, a supergroup, or a channel, since other chats can't be checked by moderators
// @param chatId Chat identifier
// @param fileId Identifier of the photo to report. Only full photos from chatPhoto can be reported
// @param reason The reason for reporting the chat photo
// @param text Additional report details; 0-1024 characters
func (client *Client) ReportChatPhoto(chatId int64, fileId int32, reason ChatReportReason, text string) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "reportChatPhoto",
		"chat_id": chatId,
		"file_id": fileId,
		"reason":  reason,
		"text":    text,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// GetChatStatistics Returns detailed statistics about a chat. Currently this method can be used only for supergroups and channels. Can be used only if supergroupFullInfo.can_get_statistics == true
// @param chatId Chat identifier
// @param isDark Pass true if a dark theme is used by the application
func (client *Client) GetChatStatistics(chatId int64, isDark bool) (ChatStatistics, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "getChatStatistics",
		"chat_id": chatId,
		"is_dark": isDark,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	switch ChatStatisticsEnum(result.Data["@type"].(string)) {

	case ChatStatisticsSupergroupType:
		var chatStatistics ChatStatisticsSupergroup
		err = json.Unmarshal(result.Raw, &chatStatistics)
		return &chatStatistics, err

	case ChatStatisticsChannelType:
		var chatStatistics ChatStatisticsChannel
		err = json.Unmarshal(result.Raw, &chatStatistics)
		return &chatStatistics, err

	default:
		return nil, fmt.Errorf("Invalid type")
	}
}

// GetMessageStatistics Returns detailed statistics about a message. Can be used only if message.can_get_statistics == true
// @param chatId Chat identifier
// @param messageId Message identifier
// @param isDark Pass true if a dark theme is used by the application
func (client *Client) GetMessageStatistics(chatId int64, messageId int64, isDark bool) (*MessageStatistics, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":      "getMessageStatistics",
		"chat_id":    chatId,
		"message_id": messageId,
		"is_dark":    isDark,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var messageStatistics MessageStatistics
	err = json.Unmarshal(result.Raw, &messageStatistics)
	return &messageStatistics, err

}

// GetStatisticalGraph Loads an asynchronous or a zoomed in statistical graph
// @param chatId Chat identifier
// @param token The token for graph loading
// @param x X-value for zoomed in graph or 0 otherwise
func (client *Client) GetStatisticalGraph(chatId int64, token string, x int64) (StatisticalGraph, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "getStatisticalGraph",
		"chat_id": chatId,
		"token":   token,
		"x":       x,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	switch StatisticalGraphEnum(result.Data["@type"].(string)) {

	case StatisticalGraphDataType:
		var statisticalGraph StatisticalGraphData
		err = json.Unmarshal(result.Raw, &statisticalGraph)
		return &statisticalGraph, err

	case StatisticalGraphAsyncType:
		var statisticalGraph StatisticalGraphAsync
		err = json.Unmarshal(result.Raw, &statisticalGraph)
		return &statisticalGraph, err

	case StatisticalGraphErrorType:
		var statisticalGraph StatisticalGraphError
		err = json.Unmarshal(result.Raw, &statisticalGraph)
		return &statisticalGraph, err

	default:
		return nil, fmt.Errorf("Invalid type")
	}
}

// GetStorageStatistics Returns storage usage statistics. Can be called before authorization
// @param chatLimit The maximum number of chats with the largest storage usage for which separate statistics need to be returned. All other chats will be grouped in entries with chat_id == 0. If the chat info database is not used, the chat_limit is ignored and is always set to 0
func (client *Client) GetStorageStatistics(chatLimit int32) (*StorageStatistics, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":      "getStorageStatistics",
		"chat_limit": chatLimit,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var storageStatistics StorageStatistics
	err = json.Unmarshal(result.Raw, &storageStatistics)
	return &storageStatistics, err

}

// GetStorageStatisticsFast Quickly returns approximate storage usage statistics. Can be called before authorization
func (client *Client) GetStorageStatisticsFast() (*StorageStatisticsFast, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "getStorageStatisticsFast",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var storageStatisticsFast StorageStatisticsFast
	err = json.Unmarshal(result.Raw, &storageStatisticsFast)
	return &storageStatisticsFast, err

}

// GetDatabaseStatistics Returns database statistics
func (client *Client) GetDatabaseStatistics() (*DatabaseStatistics, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "getDatabaseStatistics",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var databaseStatistics DatabaseStatistics
	err = json.Unmarshal(result.Raw, &databaseStatistics)
	return &databaseStatistics, err

}

// OptimizeStorage Optimizes storage usage, i.e. deletes some files and returns new storage usage statistics. Secret thumbnails can't be deleted
// @param size Limit on the total size of files after deletion, in bytes. Pass -1 to use the default limit
// @param ttl Limit on the time that has passed since the last time a file was accessed (or creation time for some filesystems). Pass -1 to use the default limit
// @param count Limit on the total count of files after deletion. Pass -1 to use the default limit
// @param immunityDelay The amount of time after the creation of a file during which it can't be deleted, in seconds. Pass -1 to use the default value
// @param fileTypes If non-empty, only files with the given types are considered. By default, all types except thumbnails, profile photos, stickers and wallpapers are deleted
// @param chatIds If non-empty, only files from the given chats are considered. Use 0 as chat identifier to delete files not belonging to any chat (e.g., profile photos)
// @param excludeChatIds If non-empty, files from the given chats are excluded. Use 0 as chat identifier to exclude all files not belonging to any chat (e.g., profile photos)
// @param returnDeletedFileStatistics Pass true if statistics about the files that were deleted must be returned instead of the whole storage usage statistics. Affects only returned statistics
// @param chatLimit Same as in getStorageStatistics. Affects only returned statistics
func (client *Client) OptimizeStorage(size int64, ttl int32, count int32, immunityDelay int32, fileTypes []FileType, chatIds []int64, excludeChatIds []int64, returnDeletedFileStatistics bool, chatLimit int32) (*StorageStatistics, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":                          "optimizeStorage",
		"size":                           size,
		"ttl":                            ttl,
		"count":                          count,
		"immunity_delay":                 immunityDelay,
		"file_types":                     fileTypes,
		"chat_ids":                       chatIds,
		"exclude_chat_ids":               excludeChatIds,
		"return_deleted_file_statistics": returnDeletedFileStatistics,
		"chat_limit":                     chatLimit,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var storageStatistics StorageStatistics
	err = json.Unmarshal(result.Raw, &storageStatistics)
	return &storageStatistics, err

}

// SetNetworkType Sets the current network type. Can be called before authorization. Calling this method forces all network connections to reopen, mitigating the delay in switching between different networks, so it must be called whenever the network is changed, even if the network type remains the same.
// @param typeParam
func (client *Client) SetNetworkType(typeParam NetworkType) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "setNetworkType",
		"type":  typeParam,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// GetNetworkStatistics Returns network data usage statistics. Can be called before authorization
// @param onlyCurrent If true, returns only data for the current library launch
func (client *Client) GetNetworkStatistics(onlyCurrent bool) (*NetworkStatistics, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":        "getNetworkStatistics",
		"only_current": onlyCurrent,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var networkStatistics NetworkStatistics
	err = json.Unmarshal(result.Raw, &networkStatistics)
	return &networkStatistics, err

}

// AddNetworkStatistics Adds the specified data to data usage statistics. Can be called before authorization
// @param entry The network statistics entry with the data to be added to statistics
func (client *Client) AddNetworkStatistics(entry NetworkStatisticsEntry) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "addNetworkStatistics",
		"entry": entry,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// ResetNetworkStatistics Resets all network data usage statistics to zero. Can be called before authorization
func (client *Client) ResetNetworkStatistics() (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "resetNetworkStatistics",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

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

// SetAutoDownloadSettings Sets auto-download settings
// @param settings New user auto-download settings
// @param typeParam Type of the network for which the new settings are relevant
func (client *Client) SetAutoDownloadSettings(settings *AutoDownloadSettings, typeParam NetworkType) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":    "setAutoDownloadSettings",
		"settings": settings,
		"type":     typeParam,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// GetBankCardInfo Returns information about a bank card
// @param bankCardNumber The bank card number
func (client *Client) GetBankCardInfo(bankCardNumber string) (*BankCardInfo, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":            "getBankCardInfo",
		"bank_card_number": bankCardNumber,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var bankCardInfo BankCardInfo
	err = json.Unmarshal(result.Raw, &bankCardInfo)
	return &bankCardInfo, err

}

// GetPassportElement Returns one of the available Telegram Passport elements
// @param typeParam Telegram Passport element type
// @param password Password of the current user
func (client *Client) GetPassportElement(typeParam PassportElementType, password string) (PassportElement, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":    "getPassportElement",
		"type":     typeParam,
		"password": password,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	switch PassportElementEnum(result.Data["@type"].(string)) {

	case PassportElementPersonalDetailsType:
		var passportElement PassportElementPersonalDetails
		err = json.Unmarshal(result.Raw, &passportElement)
		return &passportElement, err

	case PassportElementPassportType:
		var passportElement PassportElementPassport
		err = json.Unmarshal(result.Raw, &passportElement)
		return &passportElement, err

	case PassportElementDriverLicenseType:
		var passportElement PassportElementDriverLicense
		err = json.Unmarshal(result.Raw, &passportElement)
		return &passportElement, err

	case PassportElementIdentityCardType:
		var passportElement PassportElementIdentityCard
		err = json.Unmarshal(result.Raw, &passportElement)
		return &passportElement, err

	case PassportElementInternalPassportType:
		var passportElement PassportElementInternalPassport
		err = json.Unmarshal(result.Raw, &passportElement)
		return &passportElement, err

	case PassportElementAddressType:
		var passportElement PassportElementAddress
		err = json.Unmarshal(result.Raw, &passportElement)
		return &passportElement, err

	case PassportElementUtilityBillType:
		var passportElement PassportElementUtilityBill
		err = json.Unmarshal(result.Raw, &passportElement)
		return &passportElement, err

	case PassportElementBankStatementType:
		var passportElement PassportElementBankStatement
		err = json.Unmarshal(result.Raw, &passportElement)
		return &passportElement, err

	case PassportElementRentalAgreementType:
		var passportElement PassportElementRentalAgreement
		err = json.Unmarshal(result.Raw, &passportElement)
		return &passportElement, err

	case PassportElementPassportRegistrationType:
		var passportElement PassportElementPassportRegistration
		err = json.Unmarshal(result.Raw, &passportElement)
		return &passportElement, err

	case PassportElementTemporaryRegistrationType:
		var passportElement PassportElementTemporaryRegistration
		err = json.Unmarshal(result.Raw, &passportElement)
		return &passportElement, err

	case PassportElementPhoneNumberType:
		var passportElement PassportElementPhoneNumber
		err = json.Unmarshal(result.Raw, &passportElement)
		return &passportElement, err

	case PassportElementEmailAddressType:
		var passportElement PassportElementEmailAddress
		err = json.Unmarshal(result.Raw, &passportElement)
		return &passportElement, err

	default:
		return nil, fmt.Errorf("Invalid type")
	}
}

// GetAllPassportElements Returns all available Telegram Passport elements
// @param password Password of the current user
func (client *Client) GetAllPassportElements(password string) (*PassportElements, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":    "getAllPassportElements",
		"password": password,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var passportElements PassportElements
	err = json.Unmarshal(result.Raw, &passportElements)
	return &passportElements, err

}

// SetPassportElement Adds an element to the user's Telegram Passport. May return an error with a message "PHONE_VERIFICATION_NEEDED" or "EMAIL_VERIFICATION_NEEDED" if the chosen phone number or the chosen email address must be verified first
// @param element Input Telegram Passport element
// @param password Password of the current user
func (client *Client) SetPassportElement(element InputPassportElement, password string) (PassportElement, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":    "setPassportElement",
		"element":  element,
		"password": password,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	switch PassportElementEnum(result.Data["@type"].(string)) {

	case PassportElementPersonalDetailsType:
		var passportElement PassportElementPersonalDetails
		err = json.Unmarshal(result.Raw, &passportElement)
		return &passportElement, err

	case PassportElementPassportType:
		var passportElement PassportElementPassport
		err = json.Unmarshal(result.Raw, &passportElement)
		return &passportElement, err

	case PassportElementDriverLicenseType:
		var passportElement PassportElementDriverLicense
		err = json.Unmarshal(result.Raw, &passportElement)
		return &passportElement, err

	case PassportElementIdentityCardType:
		var passportElement PassportElementIdentityCard
		err = json.Unmarshal(result.Raw, &passportElement)
		return &passportElement, err

	case PassportElementInternalPassportType:
		var passportElement PassportElementInternalPassport
		err = json.Unmarshal(result.Raw, &passportElement)
		return &passportElement, err

	case PassportElementAddressType:
		var passportElement PassportElementAddress
		err = json.Unmarshal(result.Raw, &passportElement)
		return &passportElement, err

	case PassportElementUtilityBillType:
		var passportElement PassportElementUtilityBill
		err = json.Unmarshal(result.Raw, &passportElement)
		return &passportElement, err

	case PassportElementBankStatementType:
		var passportElement PassportElementBankStatement
		err = json.Unmarshal(result.Raw, &passportElement)
		return &passportElement, err

	case PassportElementRentalAgreementType:
		var passportElement PassportElementRentalAgreement
		err = json.Unmarshal(result.Raw, &passportElement)
		return &passportElement, err

	case PassportElementPassportRegistrationType:
		var passportElement PassportElementPassportRegistration
		err = json.Unmarshal(result.Raw, &passportElement)
		return &passportElement, err

	case PassportElementTemporaryRegistrationType:
		var passportElement PassportElementTemporaryRegistration
		err = json.Unmarshal(result.Raw, &passportElement)
		return &passportElement, err

	case PassportElementPhoneNumberType:
		var passportElement PassportElementPhoneNumber
		err = json.Unmarshal(result.Raw, &passportElement)
		return &passportElement, err

	case PassportElementEmailAddressType:
		var passportElement PassportElementEmailAddress
		err = json.Unmarshal(result.Raw, &passportElement)
		return &passportElement, err

	default:
		return nil, fmt.Errorf("Invalid type")
	}
}

// DeletePassportElement Deletes a Telegram Passport element
// @param typeParam Element type
func (client *Client) DeletePassportElement(typeParam PassportElementType) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "deletePassportElement",
		"type":  typeParam,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// SetPassportElementErrors Informs the user that some of the elements in their Telegram Passport contain errors; for bots only. The user will not be able to resend the elements, until the errors are fixed
// @param userId User identifier
// @param errors The errors
func (client *Client) SetPassportElementErrors(userId int64, errors []InputPassportElementError) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "setPassportElementErrors",
		"user_id": userId,
		"errors":  errors,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// GetPreferredCountryLanguage Returns an IETF language tag of the language preferred in the country, which must be used to fill native fields in Telegram Passport personal details. Returns a 404 error if unknown
// @param countryCode A two-letter ISO 3166-1 alpha-2 country code
func (client *Client) GetPreferredCountryLanguage(countryCode string) (*Text, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":        "getPreferredCountryLanguage",
		"country_code": countryCode,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var text Text
	err = json.Unmarshal(result.Raw, &text)
	return &text, err

}

// SendPhoneNumberVerificationCode Sends a code to verify a phone number to be added to a user's Telegram Passport
// @param phoneNumber The phone number of the user, in international format
// @param settings Settings for the authentication of the user's phone number; pass null to use default settings
func (client *Client) SendPhoneNumberVerificationCode(phoneNumber string, settings *PhoneNumberAuthenticationSettings) (*AuthenticationCodeInfo, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":        "sendPhoneNumberVerificationCode",
		"phone_number": phoneNumber,
		"settings":     settings,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var authenticationCodeInfo AuthenticationCodeInfo
	err = json.Unmarshal(result.Raw, &authenticationCodeInfo)
	return &authenticationCodeInfo, err

}

// ResendPhoneNumberVerificationCode Re-sends the code to verify a phone number to be added to a user's Telegram Passport
func (client *Client) ResendPhoneNumberVerificationCode() (*AuthenticationCodeInfo, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "resendPhoneNumberVerificationCode",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var authenticationCodeInfo AuthenticationCodeInfo
	err = json.Unmarshal(result.Raw, &authenticationCodeInfo)
	return &authenticationCodeInfo, err

}

// CheckPhoneNumberVerificationCode Checks the phone number verification code for Telegram Passport
// @param code Verification code
func (client *Client) CheckPhoneNumberVerificationCode(code string) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "checkPhoneNumberVerificationCode",
		"code":  code,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// SendEmailAddressVerificationCode Sends a code to verify an email address to be added to a user's Telegram Passport
// @param emailAddress Email address
func (client *Client) SendEmailAddressVerificationCode(emailAddress string) (*EmailAddressAuthenticationCodeInfo, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":         "sendEmailAddressVerificationCode",
		"email_address": emailAddress,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var emailAddressAuthenticationCodeInfo EmailAddressAuthenticationCodeInfo
	err = json.Unmarshal(result.Raw, &emailAddressAuthenticationCodeInfo)
	return &emailAddressAuthenticationCodeInfo, err

}

// ResendEmailAddressVerificationCode Re-sends the code to verify an email address to be added to a user's Telegram Passport
func (client *Client) ResendEmailAddressVerificationCode() (*EmailAddressAuthenticationCodeInfo, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "resendEmailAddressVerificationCode",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var emailAddressAuthenticationCodeInfo EmailAddressAuthenticationCodeInfo
	err = json.Unmarshal(result.Raw, &emailAddressAuthenticationCodeInfo)
	return &emailAddressAuthenticationCodeInfo, err

}

// CheckEmailAddressVerificationCode Checks the email address verification code for Telegram Passport
// @param code Verification code
func (client *Client) CheckEmailAddressVerificationCode(code string) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "checkEmailAddressVerificationCode",
		"code":  code,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// GetPassportAuthorizationForm Returns a Telegram Passport authorization form for sharing data with a service
// @param botUserId User identifier of the service's bot
// @param scope Telegram Passport element types requested by the service
// @param publicKey Service's public key
// @param nonce Unique request identifier provided by the service
func (client *Client) GetPassportAuthorizationForm(botUserId int64, scope string, publicKey string, nonce string) (*PassportAuthorizationForm, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":       "getPassportAuthorizationForm",
		"bot_user_id": botUserId,
		"scope":       scope,
		"public_key":  publicKey,
		"nonce":       nonce,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var passportAuthorizationForm PassportAuthorizationForm
	err = json.Unmarshal(result.Raw, &passportAuthorizationForm)
	return &passportAuthorizationForm, err

}

// GetPassportAuthorizationFormAvailableElements Returns already available Telegram Passport elements suitable for completing a Telegram Passport authorization form. Result can be received only once for each authorization form
// @param autorizationFormId Authorization form identifier
// @param password Password of the current user
func (client *Client) GetPassportAuthorizationFormAvailableElements(autorizationFormId int32, password string) (*PassportElementsWithErrors, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":                "getPassportAuthorizationFormAvailableElements",
		"autorization_form_id": autorizationFormId,
		"password":             password,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var passportElementsWithErrors PassportElementsWithErrors
	err = json.Unmarshal(result.Raw, &passportElementsWithErrors)
	return &passportElementsWithErrors, err

}

// SendPassportAuthorizationForm Sends a Telegram Passport authorization form, effectively sharing data with the service. This method must be called after getPassportAuthorizationFormAvailableElements if some previously available elements are going to be reused
// @param autorizationFormId Authorization form identifier
// @param typeParams Types of Telegram Passport elements chosen by user to complete the authorization form
func (client *Client) SendPassportAuthorizationForm(autorizationFormId int32, typeParams []PassportElementType) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":                "sendPassportAuthorizationForm",
		"autorization_form_id": autorizationFormId,
		"types":                typeParams,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// SendPhoneNumberConfirmationCode Sends phone number confirmation code to handle links of the type internalLinkTypePhoneNumberConfirmation
// @param hash Hash value from the link
// @param phoneNumber Phone number value from the link
// @param settings Settings for the authentication of the user's phone number; pass null to use default settings
func (client *Client) SendPhoneNumberConfirmationCode(hash string, phoneNumber string, settings *PhoneNumberAuthenticationSettings) (*AuthenticationCodeInfo, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":        "sendPhoneNumberConfirmationCode",
		"hash":         hash,
		"phone_number": phoneNumber,
		"settings":     settings,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var authenticationCodeInfo AuthenticationCodeInfo
	err = json.Unmarshal(result.Raw, &authenticationCodeInfo)
	return &authenticationCodeInfo, err

}

// ResendPhoneNumberConfirmationCode Resends phone number confirmation code
func (client *Client) ResendPhoneNumberConfirmationCode() (*AuthenticationCodeInfo, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "resendPhoneNumberConfirmationCode",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var authenticationCodeInfo AuthenticationCodeInfo
	err = json.Unmarshal(result.Raw, &authenticationCodeInfo)
	return &authenticationCodeInfo, err

}

// CheckPhoneNumberConfirmationCode Checks phone number confirmation code
// @param code The phone number confirmation code
func (client *Client) CheckPhoneNumberConfirmationCode(code string) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "checkPhoneNumberConfirmationCode",
		"code":  code,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// SetBotUpdatesStatus Informs the server about the number of pending bot updates if they haven't been processed for a long time; for bots only
// @param pendingUpdateCount The number of pending updates
// @param errorMessage The last error message
func (client *Client) SetBotUpdatesStatus(pendingUpdateCount int32, errorMessage string) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":                "setBotUpdatesStatus",
		"pending_update_count": pendingUpdateCount,
		"error_message":        errorMessage,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// UploadStickerFile Uploads a PNG image with a sticker; returns the uploaded file
// @param userId Sticker file owner; ignored for regular users
// @param sticker Sticker file to upload
func (client *Client) UploadStickerFile(userId int64, sticker InputSticker) (*File, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "uploadStickerFile",
		"user_id": userId,
		"sticker": sticker,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var file File
	err = json.Unmarshal(result.Raw, &file)
	return &file, err

}

// GetSuggestedStickerSetName Returns a suggested name for a new sticker set with a given title
// @param title Sticker set title; 1-64 characters
func (client *Client) GetSuggestedStickerSetName(title string) (*Text, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "getSuggestedStickerSetName",
		"title": title,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var text Text
	err = json.Unmarshal(result.Raw, &text)
	return &text, err

}

// CheckStickerSetName Checks whether a name can be used for a new sticker set
// @param name Name to be checked
func (client *Client) CheckStickerSetName(name string) (CheckStickerSetNameResult, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "checkStickerSetName",
		"name":  name,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	switch CheckStickerSetNameResultEnum(result.Data["@type"].(string)) {

	case CheckStickerSetNameResultOkType:
		var checkStickerSetNameResult CheckStickerSetNameResultOk
		err = json.Unmarshal(result.Raw, &checkStickerSetNameResult)
		return &checkStickerSetNameResult, err

	case CheckStickerSetNameResultNameInvalidType:
		var checkStickerSetNameResult CheckStickerSetNameResultNameInvalid
		err = json.Unmarshal(result.Raw, &checkStickerSetNameResult)
		return &checkStickerSetNameResult, err

	case CheckStickerSetNameResultNameOccupiedType:
		var checkStickerSetNameResult CheckStickerSetNameResultNameOccupied
		err = json.Unmarshal(result.Raw, &checkStickerSetNameResult)
		return &checkStickerSetNameResult, err

	default:
		return nil, fmt.Errorf("Invalid type")
	}
}

// CreateNewStickerSet Creates a new sticker set. Returns the newly created sticker set
// @param userId Sticker set owner; ignored for regular users
// @param title Sticker set title; 1-64 characters
// @param name Sticker set name. Can contain only English letters, digits and underscores. Must end with *"_by_<bot username>"* (*<bot_username>* is case insensitive) for bots; 1-64 characters
// @param isMasks True, if stickers are masks. Animated stickers can't be masks
// @param stickers List of stickers to be added to the set; must be non-empty. All stickers must be of the same type. For animated stickers, uploadStickerFile must be used before the sticker is shown
// @param source Source of the sticker set; may be empty if unknown
func (client *Client) CreateNewStickerSet(userId int64, title string, name string, isMasks bool, stickers []InputSticker, source string) (*StickerSet, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":    "createNewStickerSet",
		"user_id":  userId,
		"title":    title,
		"name":     name,
		"is_masks": isMasks,
		"stickers": stickers,
		"source":   source,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var stickerSet StickerSet
	err = json.Unmarshal(result.Raw, &stickerSet)
	return &stickerSet, err

}

// AddStickerToSet Adds a new sticker to a set; for bots only. Returns the sticker set
// @param userId Sticker set owner
// @param name Sticker set name
// @param sticker Sticker to add to the set
func (client *Client) AddStickerToSet(userId int64, name string, sticker InputSticker) (*StickerSet, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "addStickerToSet",
		"user_id": userId,
		"name":    name,
		"sticker": sticker,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var stickerSet StickerSet
	err = json.Unmarshal(result.Raw, &stickerSet)
	return &stickerSet, err

}

// SetStickerSetThumbnail Sets a sticker set thumbnail; for bots only. Returns the sticker set
// @param userId Sticker set owner
// @param name Sticker set name
// @param thumbnail Thumbnail to set in PNG or TGS format; pass null to remove the sticker set thumbnail. Animated thumbnail must be set for animated sticker sets and only for them
func (client *Client) SetStickerSetThumbnail(userId int64, name string, thumbnail InputFile) (*StickerSet, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":     "setStickerSetThumbnail",
		"user_id":   userId,
		"name":      name,
		"thumbnail": thumbnail,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var stickerSet StickerSet
	err = json.Unmarshal(result.Raw, &stickerSet)
	return &stickerSet, err

}

// SetStickerPositionInSet Changes the position of a sticker in the set to which it belongs; for bots only. The sticker set must have been created by the bot
// @param sticker Sticker
// @param position New position of the sticker in the set, zero-based
func (client *Client) SetStickerPositionInSet(sticker InputFile, position int32) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":    "setStickerPositionInSet",
		"sticker":  sticker,
		"position": position,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// RemoveStickerFromSet Removes a sticker from the set to which it belongs; for bots only. The sticker set must have been created by the bot
// @param sticker Sticker
func (client *Client) RemoveStickerFromSet(sticker InputFile) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "removeStickerFromSet",
		"sticker": sticker,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// GetMapThumbnailFile Returns information about a file with a map thumbnail in PNG format. Only map thumbnail files with size less than 1MB can be downloaded
// @param location Location of the map center
// @param zoom Map zoom level; 13-20
// @param width Map width in pixels before applying scale; 16-1024
// @param height Map height in pixels before applying scale; 16-1024
// @param scale Map scale; 1-3
// @param chatId Identifier of a chat, in which the thumbnail will be shown. Use 0 if unknown
func (client *Client) GetMapThumbnailFile(location *Location, zoom int32, width int32, height int32, scale int32, chatId int64) (*File, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":    "getMapThumbnailFile",
		"location": location,
		"zoom":     zoom,
		"width":    width,
		"height":   height,
		"scale":    scale,
		"chat_id":  chatId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var file File
	err = json.Unmarshal(result.Raw, &file)
	return &file, err

}

// AcceptTermsOfService Accepts Telegram terms of services
// @param termsOfServiceId Terms of service identifier
func (client *Client) AcceptTermsOfService(termsOfServiceId string) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":               "acceptTermsOfService",
		"terms_of_service_id": termsOfServiceId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// SendCustomRequest Sends a custom request; for bots only
// @param method The method name
// @param parameters JSON-serialized method parameters
func (client *Client) SendCustomRequest(method string, parameters string) (*CustomRequestResult, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":      "sendCustomRequest",
		"method":     method,
		"parameters": parameters,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var customRequestResult CustomRequestResult
	err = json.Unmarshal(result.Raw, &customRequestResult)
	return &customRequestResult, err

}

// AnswerCustomQuery Answers a custom query; for bots only
// @param customQueryId Identifier of a custom query
// @param data JSON-serialized answer to the query
func (client *Client) AnswerCustomQuery(customQueryId JSONInt64, data string) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":           "answerCustomQuery",
		"custom_query_id": customQueryId,
		"data":            data,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// SetAlarm Succeeds after a specified amount of time has passed. Can be called before initialization
// @param seconds Number of seconds before the function returns
func (client *Client) SetAlarm(seconds float64) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "setAlarm",
		"seconds": seconds,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// GetCountries Returns information about existing countries. Can be called before authorization
func (client *Client) GetCountries() (*Countries, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "getCountries",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var countries Countries
	err = json.Unmarshal(result.Raw, &countries)
	return &countries, err

}

// GetCountryCode Uses the current IP address to find the current country. Returns two-letter ISO 3166-1 alpha-2 country code. Can be called before authorization
func (client *Client) GetCountryCode() (*Text, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "getCountryCode",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var text Text
	err = json.Unmarshal(result.Raw, &text)
	return &text, err

}

// GetPhoneNumberInfo Returns information about a phone number by its prefix. Can be called before authorization
// @param phoneNumberPrefix The phone number prefix
func (client *Client) GetPhoneNumberInfo(phoneNumberPrefix string) (*PhoneNumberInfo, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":               "getPhoneNumberInfo",
		"phone_number_prefix": phoneNumberPrefix,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var phoneNumberInfo PhoneNumberInfo
	err = json.Unmarshal(result.Raw, &phoneNumberInfo)
	return &phoneNumberInfo, err

}

// GetPhoneNumberInfoSync Returns information about a phone number by its prefix synchronously. getCountries must be called at least once after changing localization to the specified language if properly localized country information is expected. Can be called synchronously
// @param languageCode A two-letter ISO 639-1 country code for country information localization
// @param phoneNumberPrefix The phone number prefix
func (client *Client) GetPhoneNumberInfoSync(languageCode string, phoneNumberPrefix string) (*PhoneNumberInfo, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":               "getPhoneNumberInfoSync",
		"language_code":       languageCode,
		"phone_number_prefix": phoneNumberPrefix,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var phoneNumberInfo PhoneNumberInfo
	err = json.Unmarshal(result.Raw, &phoneNumberInfo)
	return &phoneNumberInfo, err

}

// GetApplicationDownloadLink Returns the link for downloading official Telegram application to be used when the current user invites friends to Telegram
func (client *Client) GetApplicationDownloadLink() (*HttpUrl, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "getApplicationDownloadLink",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var httpUrl HttpUrl
	err = json.Unmarshal(result.Raw, &httpUrl)
	return &httpUrl, err

}

// GetDeepLinkInfo Returns information about a tg:// deep link. Use "tg://need_update_for_some_feature" or "tg:some_unsupported_feature" for testing. Returns a 404 error for unknown links. Can be called before authorization
// @param link The link
func (client *Client) GetDeepLinkInfo(link string) (*DeepLinkInfo, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "getDeepLinkInfo",
		"link":  link,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var deepLinkInfo DeepLinkInfo
	err = json.Unmarshal(result.Raw, &deepLinkInfo)
	return &deepLinkInfo, err

}

// GetApplicationConfig Returns application config, provided by the server. Can be called before authorization
func (client *Client) GetApplicationConfig() (JsonValue, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "getApplicationConfig",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	switch JsonValueEnum(result.Data["@type"].(string)) {

	case JsonValueNullType:
		var jsonValue JsonValueNull
		err = json.Unmarshal(result.Raw, &jsonValue)
		return &jsonValue, err

	case JsonValueBooleanType:
		var jsonValue JsonValueBoolean
		err = json.Unmarshal(result.Raw, &jsonValue)
		return &jsonValue, err

	case JsonValueNumberType:
		var jsonValue JsonValueNumber
		err = json.Unmarshal(result.Raw, &jsonValue)
		return &jsonValue, err

	case JsonValueStringType:
		var jsonValue JsonValueString
		err = json.Unmarshal(result.Raw, &jsonValue)
		return &jsonValue, err

	case JsonValueArrayType:
		var jsonValue JsonValueArray
		err = json.Unmarshal(result.Raw, &jsonValue)
		return &jsonValue, err

	case JsonValueObjectType:
		var jsonValue JsonValueObject
		err = json.Unmarshal(result.Raw, &jsonValue)
		return &jsonValue, err

	default:
		return nil, fmt.Errorf("Invalid type")
	}
}

// SaveApplicationLogEvent Saves application log event on the server. Can be called before authorization
// @param typeParam Event type
// @param chatId Optional chat identifier, associated with the event
// @param data The log event data
func (client *Client) SaveApplicationLogEvent(typeParam string, chatId int64, data JsonValue) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "saveApplicationLogEvent",
		"type":    typeParam,
		"chat_id": chatId,
		"data":    data,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// AddProxy Adds a proxy server for network requests. Can be called before authorization
// @param server Proxy server IP address
// @param port Proxy server port
// @param enable True, if the proxy needs to be enabled
// @param typeParam Proxy type
func (client *Client) AddProxy(server string, port int32, enable bool, typeParam ProxyType) (*Proxy, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":  "addProxy",
		"server": server,
		"port":   port,
		"enable": enable,
		"type":   typeParam,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var proxy Proxy
	err = json.Unmarshal(result.Raw, &proxy)
	return &proxy, err

}

// EditProxy Edits an existing proxy server for network requests. Can be called before authorization
// @param proxyId Proxy identifier
// @param server Proxy server IP address
// @param port Proxy server port
// @param enable True, if the proxy needs to be enabled
// @param typeParam Proxy type
func (client *Client) EditProxy(proxyId int32, server string, port int32, enable bool, typeParam ProxyType) (*Proxy, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":    "editProxy",
		"proxy_id": proxyId,
		"server":   server,
		"port":     port,
		"enable":   enable,
		"type":     typeParam,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var proxyDummy Proxy
	err = json.Unmarshal(result.Raw, &proxyDummy)
	return &proxyDummy, err

}

// EnableProxy Enables a proxy. Only one proxy can be enabled at a time. Can be called before authorization
// @param proxyId Proxy identifier
func (client *Client) EnableProxy(proxyId int32) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":    "enableProxy",
		"proxy_id": proxyId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// DisableProxy Disables the currently enabled proxy. Can be called before authorization
func (client *Client) DisableProxy() (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "disableProxy",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// RemoveProxy Removes a proxy server. Can be called before authorization
// @param proxyId Proxy identifier
func (client *Client) RemoveProxy(proxyId int32) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":    "removeProxy",
		"proxy_id": proxyId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// GetProxies Returns list of proxies that are currently set up. Can be called before authorization
func (client *Client) GetProxies() (*Proxies, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "getProxies",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var proxies Proxies
	err = json.Unmarshal(result.Raw, &proxies)
	return &proxies, err

}

// GetProxyLink Returns an HTTPS link, which can be used to add a proxy. Available only for SOCKS5 and MTProto proxies. Can be called before authorization
// @param proxyId Proxy identifier
func (client *Client) GetProxyLink(proxyId int32) (*HttpUrl, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":    "getProxyLink",
		"proxy_id": proxyId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var httpUrl HttpUrl
	err = json.Unmarshal(result.Raw, &httpUrl)
	return &httpUrl, err

}

// PingProxy Computes time needed to receive a response from a Telegram server through a proxy. Can be called before authorization
// @param proxyId Proxy identifier. Use 0 to ping a Telegram server without a proxy
func (client *Client) PingProxy(proxyId int32) (*Seconds, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":    "pingProxy",
		"proxy_id": proxyId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var seconds Seconds
	err = json.Unmarshal(result.Raw, &seconds)
	return &seconds, err

}

// SetLogStream Sets new log stream for internal logging of TDLib. Can be called synchronously
// @param logStream New log stream
func (client *Client) SetLogStream(logStream LogStream) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":      "setLogStream",
		"log_stream": logStream,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// GetLogStream Returns information about currently used log stream for internal logging of TDLib. Can be called synchronously
func (client *Client) GetLogStream() (LogStream, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "getLogStream",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	switch LogStreamEnum(result.Data["@type"].(string)) {

	case LogStreamDefaultType:
		var logStream LogStreamDefault
		err = json.Unmarshal(result.Raw, &logStream)
		return &logStream, err

	case LogStreamFileType:
		var logStream LogStreamFile
		err = json.Unmarshal(result.Raw, &logStream)
		return &logStream, err

	case LogStreamEmptyType:
		var logStream LogStreamEmpty
		err = json.Unmarshal(result.Raw, &logStream)
		return &logStream, err

	default:
		return nil, fmt.Errorf("Invalid type")
	}
}

// SetLogVerbosityLevel Sets the verbosity level of the internal logging of TDLib. Can be called synchronously
// @param newVerbosityLevel New value of the verbosity level for logging. Value 0 corresponds to fatal errors, value 1 corresponds to errors, value 2 corresponds to warnings and debug warnings, value 3 corresponds to informational, value 4 corresponds to debug, value 5 corresponds to verbose debug, value greater than 5 and up to 1023 can be used to enable even more logging
func (client *Client) SetLogVerbosityLevel(newVerbosityLevel int32) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":               "setLogVerbosityLevel",
		"new_verbosity_level": newVerbosityLevel,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// GetLogVerbosityLevel Returns current verbosity level of the internal logging of TDLib. Can be called synchronously
func (client *Client) GetLogVerbosityLevel() (*LogVerbosityLevel, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "getLogVerbosityLevel",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var logVerbosityLevel LogVerbosityLevel
	err = json.Unmarshal(result.Raw, &logVerbosityLevel)
	return &logVerbosityLevel, err

}

// GetLogTags Returns list of available TDLib internal log tags, for example, ["actor", "binlog", "connections", "notifications", "proxy"]. Can be called synchronously
func (client *Client) GetLogTags() (*LogTags, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "getLogTags",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var logTags LogTags
	err = json.Unmarshal(result.Raw, &logTags)
	return &logTags, err

}

// SetLogTagVerbosityLevel Sets the verbosity level for a specified TDLib internal log tag. Can be called synchronously
// @param tag Logging tag to change verbosity level
// @param newVerbosityLevel New verbosity level; 1-1024
func (client *Client) SetLogTagVerbosityLevel(tag string, newVerbosityLevel int32) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":               "setLogTagVerbosityLevel",
		"tag":                 tag,
		"new_verbosity_level": newVerbosityLevel,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// GetLogTagVerbosityLevel Returns current verbosity level for a specified TDLib internal log tag. Can be called synchronously
// @param tag Logging tag to change verbosity level
func (client *Client) GetLogTagVerbosityLevel(tag string) (*LogVerbosityLevel, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "getLogTagVerbosityLevel",
		"tag":   tag,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var logVerbosityLevel LogVerbosityLevel
	err = json.Unmarshal(result.Raw, &logVerbosityLevel)
	return &logVerbosityLevel, err

}

// AddLogMessage Adds a message to TDLib internal log. Can be called synchronously
// @param verbosityLevel The minimum verbosity level needed for the message to be logged; 0-1023
// @param text Text of a message to log
func (client *Client) AddLogMessage(verbosityLevel int32, text string) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":           "addLogMessage",
		"verbosity_level": verbosityLevel,
		"text":            text,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// TestCallEmpty Does nothing; for testing only. This is an offline method. Can be called before authorization
func (client *Client) TestCallEmpty() (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "testCallEmpty",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// TestCallString Returns the received string; for testing only. This is an offline method. Can be called before authorization
// @param x String to return
func (client *Client) TestCallString(x string) (*TestString, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "testCallString",
		"x":     x,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var testString TestString
	err = json.Unmarshal(result.Raw, &testString)
	return &testString, err

}

// TestCallBytes Returns the received bytes; for testing only. This is an offline method. Can be called before authorization
// @param x Bytes to return
func (client *Client) TestCallBytes(x []byte) (*TestBytes, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "testCallBytes",
		"x":     x,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var testBytes TestBytes
	err = json.Unmarshal(result.Raw, &testBytes)
	return &testBytes, err

}

// TestCallVectorInt Returns the received vector of numbers; for testing only. This is an offline method. Can be called before authorization
// @param x Vector of numbers to return
func (client *Client) TestCallVectorInt(x []int32) (*TestVectorInt, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "testCallVectorInt",
		"x":     x,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var testVectorInt TestVectorInt
	err = json.Unmarshal(result.Raw, &testVectorInt)
	return &testVectorInt, err

}

// TestCallVectorIntObject Returns the received vector of objects containing a number; for testing only. This is an offline method. Can be called before authorization
// @param x Vector of objects to return
func (client *Client) TestCallVectorIntObject(x []TestInt) (*TestVectorIntObject, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "testCallVectorIntObject",
		"x":     x,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var testVectorIntObject TestVectorIntObject
	err = json.Unmarshal(result.Raw, &testVectorIntObject)
	return &testVectorIntObject, err

}

// TestCallVectorString Returns the received vector of strings; for testing only. This is an offline method. Can be called before authorization
// @param x Vector of strings to return
func (client *Client) TestCallVectorString(x []string) (*TestVectorString, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "testCallVectorString",
		"x":     x,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var testVectorString TestVectorString
	err = json.Unmarshal(result.Raw, &testVectorString)
	return &testVectorString, err

}

// TestCallVectorStringObject Returns the received vector of objects containing a string; for testing only. This is an offline method. Can be called before authorization
// @param x Vector of objects to return
func (client *Client) TestCallVectorStringObject(x []TestString) (*TestVectorStringObject, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "testCallVectorStringObject",
		"x":     x,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var testVectorStringObject TestVectorStringObject
	err = json.Unmarshal(result.Raw, &testVectorStringObject)
	return &testVectorStringObject, err

}

// TestSquareInt Returns the squared received number; for testing only. This is an offline method. Can be called before authorization
// @param x Number to square
func (client *Client) TestSquareInt(x int32) (*TestInt, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "testSquareInt",
		"x":     x,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var testInt TestInt
	err = json.Unmarshal(result.Raw, &testInt)
	return &testInt, err

}

// TestNetwork Sends a simple network request to the Telegram servers; for testing only. Can be called before authorization
func (client *Client) TestNetwork() (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "testNetwork",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// TestProxy Sends a simple network request to the Telegram servers via proxy; for testing only. Can be called before authorization
// @param server Proxy server IP address
// @param port Proxy server port
// @param typeParam Proxy type
// @param dcId Identifier of a datacenter, with which to test connection
// @param timeout The maximum overall timeout for the request
func (client *Client) TestProxy(server string, port int32, typeParam ProxyType, dcId int32, timeout float64) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "testProxy",
		"server":  server,
		"port":    port,
		"type":    typeParam,
		"dc_id":   dcId,
		"timeout": timeout,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// TestGetDifference Forces an updates.getDifference call to the Telegram servers; for testing only
func (client *Client) TestGetDifference() (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "testGetDifference",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// TestUseUpdate Does nothing and ensures that the Update object is used; for testing only. This is an offline method. Can be called before authorization
func (client *Client) TestUseUpdate() (Update, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "testUseUpdate",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	switch UpdateEnum(result.Data["@type"].(string)) {

	case UpdateAuthorizationStateType:
		var update UpdateAuthorizationState
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateNewMessageType:
		var update UpdateNewMessage
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateMessageSendAcknowledgedType:
		var update UpdateMessageSendAcknowledged
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateMessageSendSucceededType:
		var update UpdateMessageSendSucceeded
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateMessageSendFailedType:
		var update UpdateMessageSendFailed
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateMessageContentType:
		var update UpdateMessageContent
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateMessageEditedType:
		var update UpdateMessageEdited
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateMessageIsPinnedType:
		var update UpdateMessageIsPinned
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateMessageInteractionInfoType:
		var update UpdateMessageInteractionInfo
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateMessageContentOpenedType:
		var update UpdateMessageContentOpened
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateMessageMentionReadType:
		var update UpdateMessageMentionRead
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateMessageLiveLocationViewedType:
		var update UpdateMessageLiveLocationViewed
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateNewChatType:
		var update UpdateNewChat
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateChatTitleType:
		var update UpdateChatTitle
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateChatPhotoType:
		var update UpdateChatPhoto
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateChatPermissionsType:
		var update UpdateChatPermissions
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateChatLastMessageType:
		var update UpdateChatLastMessage
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateChatPositionType:
		var update UpdateChatPosition
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateChatIsMarkedAsUnreadType:
		var update UpdateChatIsMarkedAsUnread
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateChatIsBlockedType:
		var update UpdateChatIsBlocked
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateChatHasScheduledMessagesType:
		var update UpdateChatHasScheduledMessages
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateChatVideoChatType:
		var update UpdateChatVideoChat
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateChatDefaultDisableNotificationType:
		var update UpdateChatDefaultDisableNotification
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateChatReadInboxType:
		var update UpdateChatReadInbox
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateChatReadOutboxType:
		var update UpdateChatReadOutbox
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateChatUnreadMentionCountType:
		var update UpdateChatUnreadMentionCount
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateChatNotificationSettingsType:
		var update UpdateChatNotificationSettings
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateScopeNotificationSettingsType:
		var update UpdateScopeNotificationSettings
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateChatMessageTtlSettingType:
		var update UpdateChatMessageTtlSetting
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateChatActionBarType:
		var update UpdateChatActionBar
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateChatThemeType:
		var update UpdateChatTheme
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateChatPendingJoinRequestsType:
		var update UpdateChatPendingJoinRequests
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateChatReplyMarkupType:
		var update UpdateChatReplyMarkup
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateChatDraftMessageType:
		var update UpdateChatDraftMessage
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateChatFiltersType:
		var update UpdateChatFilters
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateChatOnlineMemberCountType:
		var update UpdateChatOnlineMemberCount
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateNotificationType:
		var update UpdateNotification
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateNotificationGroupType:
		var update UpdateNotificationGroup
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateActiveNotificationsType:
		var update UpdateActiveNotifications
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateHavePendingNotificationsType:
		var update UpdateHavePendingNotifications
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateDeleteMessagesType:
		var update UpdateDeleteMessages
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateUserChatActionType:
		var update UpdateUserChatAction
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateUserStatusType:
		var update UpdateUserStatus
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateUserType:
		var update UpdateUser
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateBasicGroupType:
		var update UpdateBasicGroup
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateSupergroupType:
		var update UpdateSupergroup
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateSecretChatType:
		var update UpdateSecretChat
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateUserFullInfoType:
		var update UpdateUserFullInfo
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateBasicGroupFullInfoType:
		var update UpdateBasicGroupFullInfo
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateSupergroupFullInfoType:
		var update UpdateSupergroupFullInfo
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateServiceNotificationType:
		var update UpdateServiceNotification
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateFileType:
		var update UpdateFile
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateFileGenerationStartType:
		var update UpdateFileGenerationStart
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateFileGenerationStopType:
		var update UpdateFileGenerationStop
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateCallType:
		var update UpdateCall
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateGroupCallType:
		var update UpdateGroupCall
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateGroupCallParticipantType:
		var update UpdateGroupCallParticipant
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateNewCallSignalingDataType:
		var update UpdateNewCallSignalingData
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateUserPrivacySettingRulesType:
		var update UpdateUserPrivacySettingRules
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateUnreadMessageCountType:
		var update UpdateUnreadMessageCount
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateUnreadChatCountType:
		var update UpdateUnreadChatCount
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateOptionType:
		var update UpdateOption
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateStickerSetType:
		var update UpdateStickerSet
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateInstalledStickerSetsType:
		var update UpdateInstalledStickerSets
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateTrendingStickerSetsType:
		var update UpdateTrendingStickerSets
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateRecentStickersType:
		var update UpdateRecentStickers
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateFavoriteStickersType:
		var update UpdateFavoriteStickers
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateSavedAnimationsType:
		var update UpdateSavedAnimations
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateSelectedBackgroundType:
		var update UpdateSelectedBackground
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateChatThemesType:
		var update UpdateChatThemes
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateLanguagePackStringsType:
		var update UpdateLanguagePackStrings
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateConnectionStateType:
		var update UpdateConnectionState
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateTermsOfServiceType:
		var update UpdateTermsOfService
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateUsersNearbyType:
		var update UpdateUsersNearby
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateDiceEmojisType:
		var update UpdateDiceEmojis
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateAnimatedEmojiMessageClickedType:
		var update UpdateAnimatedEmojiMessageClicked
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateAnimationSearchParametersType:
		var update UpdateAnimationSearchParameters
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateSuggestedActionsType:
		var update UpdateSuggestedActions
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateNewInlineQueryType:
		var update UpdateNewInlineQuery
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateNewChosenInlineResultType:
		var update UpdateNewChosenInlineResult
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateNewCallbackQueryType:
		var update UpdateNewCallbackQuery
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateNewInlineCallbackQueryType:
		var update UpdateNewInlineCallbackQuery
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateNewShippingQueryType:
		var update UpdateNewShippingQuery
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateNewPreCheckoutQueryType:
		var update UpdateNewPreCheckoutQuery
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateNewCustomEventType:
		var update UpdateNewCustomEvent
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateNewCustomQueryType:
		var update UpdateNewCustomQuery
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdatePollType:
		var update UpdatePoll
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdatePollAnswerType:
		var update UpdatePollAnswer
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateChatMemberType:
		var update UpdateChatMember
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateNewChatJoinRequestType:
		var update UpdateNewChatJoinRequest
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	default:
		return nil, fmt.Errorf("Invalid type")
	}
}

// TestReturnError Returns the specified error and ensures that the Error object is used; for testing only. Can be called synchronously
// @param error The error to be returned
func (client *Client) TestReturnError(error *Error) (*Error, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "testReturnError",
		"error": error,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var errorDummy Error
	err = json.Unmarshal(result.Raw, &errorDummy)
	return &errorDummy, err

}
