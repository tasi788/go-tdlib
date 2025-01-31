package tdlib

import (
	"encoding/json"
	"fmt"
)

// AuthorizationState Represents the current authorization state of the TDLib client
type AuthorizationState interface {
	GetAuthorizationStateEnum() AuthorizationStateEnum
}

// AuthorizationStateEnum Alias for abstract AuthorizationState 'Sub-Classes', used as constant-enum here
type AuthorizationStateEnum string

// AuthorizationState enums
const (
	AuthorizationStateWaitTdlibParametersType         AuthorizationStateEnum = "authorizationStateWaitTdlibParameters"
	AuthorizationStateWaitEncryptionKeyType           AuthorizationStateEnum = "authorizationStateWaitEncryptionKey"
	AuthorizationStateWaitPhoneNumberType             AuthorizationStateEnum = "authorizationStateWaitPhoneNumber"
	AuthorizationStateWaitCodeType                    AuthorizationStateEnum = "authorizationStateWaitCode"
	AuthorizationStateWaitOtherDeviceConfirmationType AuthorizationStateEnum = "authorizationStateWaitOtherDeviceConfirmation"
	AuthorizationStateWaitRegistrationType            AuthorizationStateEnum = "authorizationStateWaitRegistration"
	AuthorizationStateWaitPasswordType                AuthorizationStateEnum = "authorizationStateWaitPassword"
	AuthorizationStateReadyType                       AuthorizationStateEnum = "authorizationStateReady"
	AuthorizationStateLoggingOutType                  AuthorizationStateEnum = "authorizationStateLoggingOut"
	AuthorizationStateClosingType                     AuthorizationStateEnum = "authorizationStateClosing"
	AuthorizationStateClosedType                      AuthorizationStateEnum = "authorizationStateClosed"
)

func unmarshalAuthorizationState(rawMsg *json.RawMessage) (AuthorizationState, error) {

	if rawMsg == nil {
		return nil, nil
	}
	var objMap map[string]interface{}
	err := json.Unmarshal(*rawMsg, &objMap)
	if err != nil {
		return nil, err
	}

	switch AuthorizationStateEnum(objMap["@type"].(string)) {
	case AuthorizationStateWaitTdlibParametersType:
		var authorizationStateWaitTdlibParameters AuthorizationStateWaitTdlibParameters
		err := json.Unmarshal(*rawMsg, &authorizationStateWaitTdlibParameters)
		return &authorizationStateWaitTdlibParameters, err

	case AuthorizationStateWaitEncryptionKeyType:
		var authorizationStateWaitEncryptionKey AuthorizationStateWaitEncryptionKey
		err := json.Unmarshal(*rawMsg, &authorizationStateWaitEncryptionKey)
		return &authorizationStateWaitEncryptionKey, err

	case AuthorizationStateWaitPhoneNumberType:
		var authorizationStateWaitPhoneNumber AuthorizationStateWaitPhoneNumber
		err := json.Unmarshal(*rawMsg, &authorizationStateWaitPhoneNumber)
		return &authorizationStateWaitPhoneNumber, err

	case AuthorizationStateWaitCodeType:
		var authorizationStateWaitCode AuthorizationStateWaitCode
		err := json.Unmarshal(*rawMsg, &authorizationStateWaitCode)
		return &authorizationStateWaitCode, err

	case AuthorizationStateWaitOtherDeviceConfirmationType:
		var authorizationStateWaitOtherDeviceConfirmation AuthorizationStateWaitOtherDeviceConfirmation
		err := json.Unmarshal(*rawMsg, &authorizationStateWaitOtherDeviceConfirmation)
		return &authorizationStateWaitOtherDeviceConfirmation, err

	case AuthorizationStateWaitRegistrationType:
		var authorizationStateWaitRegistration AuthorizationStateWaitRegistration
		err := json.Unmarshal(*rawMsg, &authorizationStateWaitRegistration)
		return &authorizationStateWaitRegistration, err

	case AuthorizationStateWaitPasswordType:
		var authorizationStateWaitPassword AuthorizationStateWaitPassword
		err := json.Unmarshal(*rawMsg, &authorizationStateWaitPassword)
		return &authorizationStateWaitPassword, err

	case AuthorizationStateReadyType:
		var authorizationStateReady AuthorizationStateReady
		err := json.Unmarshal(*rawMsg, &authorizationStateReady)
		return &authorizationStateReady, err

	case AuthorizationStateLoggingOutType:
		var authorizationStateLoggingOut AuthorizationStateLoggingOut
		err := json.Unmarshal(*rawMsg, &authorizationStateLoggingOut)
		return &authorizationStateLoggingOut, err

	case AuthorizationStateClosingType:
		var authorizationStateClosing AuthorizationStateClosing
		err := json.Unmarshal(*rawMsg, &authorizationStateClosing)
		return &authorizationStateClosing, err

	case AuthorizationStateClosedType:
		var authorizationStateClosed AuthorizationStateClosed
		err := json.Unmarshal(*rawMsg, &authorizationStateClosed)
		return &authorizationStateClosed, err

	default:
		return nil, fmt.Errorf("Error UnMarshaling, unknown type:" + objMap["@type"].(string))
	}
}

// AuthorizationStateWaitTdlibParameters TDLib needs TdlibParameters for initialization
type AuthorizationStateWaitTdlibParameters struct {
	tdCommon
}

// MessageType return the string telegram-type of AuthorizationStateWaitTdlibParameters
func (authorizationStateWaitTdlibParameters *AuthorizationStateWaitTdlibParameters) MessageType() string {
	return "authorizationStateWaitTdlibParameters"
}

// NewAuthorizationStateWaitTdlibParameters creates a new AuthorizationStateWaitTdlibParameters
//
func NewAuthorizationStateWaitTdlibParameters() *AuthorizationStateWaitTdlibParameters {
	authorizationStateWaitTdlibParametersTemp := AuthorizationStateWaitTdlibParameters{
		tdCommon: tdCommon{Type: "authorizationStateWaitTdlibParameters"},
	}

	return &authorizationStateWaitTdlibParametersTemp
}

// GetAuthorizationStateEnum return the enum type of this object
func (authorizationStateWaitTdlibParameters *AuthorizationStateWaitTdlibParameters) GetAuthorizationStateEnum() AuthorizationStateEnum {
	return AuthorizationStateWaitTdlibParametersType
}

// AuthorizationStateWaitEncryptionKey TDLib needs an encryption key to decrypt the local database
type AuthorizationStateWaitEncryptionKey struct {
	tdCommon
	IsEncrypted bool `json:"is_encrypted"` // True, if the database is currently encrypted
}

// MessageType return the string telegram-type of AuthorizationStateWaitEncryptionKey
func (authorizationStateWaitEncryptionKey *AuthorizationStateWaitEncryptionKey) MessageType() string {
	return "authorizationStateWaitEncryptionKey"
}

// NewAuthorizationStateWaitEncryptionKey creates a new AuthorizationStateWaitEncryptionKey
//
// @param isEncrypted True, if the database is currently encrypted
func NewAuthorizationStateWaitEncryptionKey(isEncrypted bool) *AuthorizationStateWaitEncryptionKey {
	authorizationStateWaitEncryptionKeyTemp := AuthorizationStateWaitEncryptionKey{
		tdCommon:    tdCommon{Type: "authorizationStateWaitEncryptionKey"},
		IsEncrypted: isEncrypted,
	}

	return &authorizationStateWaitEncryptionKeyTemp
}

// GetAuthorizationStateEnum return the enum type of this object
func (authorizationStateWaitEncryptionKey *AuthorizationStateWaitEncryptionKey) GetAuthorizationStateEnum() AuthorizationStateEnum {
	return AuthorizationStateWaitEncryptionKeyType
}

// AuthorizationStateWaitPhoneNumber TDLib needs the user's phone number to authorize. Call `setAuthenticationPhoneNumber` to provide the phone number, or use `requestQrCodeAuthentication`, or `checkAuthenticationBotToken` for other authentication options
type AuthorizationStateWaitPhoneNumber struct {
	tdCommon
}

// MessageType return the string telegram-type of AuthorizationStateWaitPhoneNumber
func (authorizationStateWaitPhoneNumber *AuthorizationStateWaitPhoneNumber) MessageType() string {
	return "authorizationStateWaitPhoneNumber"
}

// NewAuthorizationStateWaitPhoneNumber creates a new AuthorizationStateWaitPhoneNumber
//
func NewAuthorizationStateWaitPhoneNumber() *AuthorizationStateWaitPhoneNumber {
	authorizationStateWaitPhoneNumberTemp := AuthorizationStateWaitPhoneNumber{
		tdCommon: tdCommon{Type: "authorizationStateWaitPhoneNumber"},
	}

	return &authorizationStateWaitPhoneNumberTemp
}

// GetAuthorizationStateEnum return the enum type of this object
func (authorizationStateWaitPhoneNumber *AuthorizationStateWaitPhoneNumber) GetAuthorizationStateEnum() AuthorizationStateEnum {
	return AuthorizationStateWaitPhoneNumberType
}

// AuthorizationStateWaitCode TDLib needs the user's authentication code to authorize
type AuthorizationStateWaitCode struct {
	tdCommon
	CodeInfo *AuthenticationCodeInfo `json:"code_info"` // Information about the authorization code that was sent
}

// MessageType return the string telegram-type of AuthorizationStateWaitCode
func (authorizationStateWaitCode *AuthorizationStateWaitCode) MessageType() string {
	return "authorizationStateWaitCode"
}

// NewAuthorizationStateWaitCode creates a new AuthorizationStateWaitCode
//
// @param codeInfo Information about the authorization code that was sent
func NewAuthorizationStateWaitCode(codeInfo *AuthenticationCodeInfo) *AuthorizationStateWaitCode {
	authorizationStateWaitCodeTemp := AuthorizationStateWaitCode{
		tdCommon: tdCommon{Type: "authorizationStateWaitCode"},
		CodeInfo: codeInfo,
	}

	return &authorizationStateWaitCodeTemp
}

// GetAuthorizationStateEnum return the enum type of this object
func (authorizationStateWaitCode *AuthorizationStateWaitCode) GetAuthorizationStateEnum() AuthorizationStateEnum {
	return AuthorizationStateWaitCodeType
}

// AuthorizationStateWaitOtherDeviceConfirmation The user needs to confirm authorization on another logged in device by scanning a QR code with the provided link
type AuthorizationStateWaitOtherDeviceConfirmation struct {
	tdCommon
	Link string `json:"link"` // A tg:// URL for the QR code. The link will be updated frequently
}

// MessageType return the string telegram-type of AuthorizationStateWaitOtherDeviceConfirmation
func (authorizationStateWaitOtherDeviceConfirmation *AuthorizationStateWaitOtherDeviceConfirmation) MessageType() string {
	return "authorizationStateWaitOtherDeviceConfirmation"
}

// NewAuthorizationStateWaitOtherDeviceConfirmation creates a new AuthorizationStateWaitOtherDeviceConfirmation
//
// @param link A tg:// URL for the QR code. The link will be updated frequently
func NewAuthorizationStateWaitOtherDeviceConfirmation(link string) *AuthorizationStateWaitOtherDeviceConfirmation {
	authorizationStateWaitOtherDeviceConfirmationTemp := AuthorizationStateWaitOtherDeviceConfirmation{
		tdCommon: tdCommon{Type: "authorizationStateWaitOtherDeviceConfirmation"},
		Link:     link,
	}

	return &authorizationStateWaitOtherDeviceConfirmationTemp
}

// GetAuthorizationStateEnum return the enum type of this object
func (authorizationStateWaitOtherDeviceConfirmation *AuthorizationStateWaitOtherDeviceConfirmation) GetAuthorizationStateEnum() AuthorizationStateEnum {
	return AuthorizationStateWaitOtherDeviceConfirmationType
}

// AuthorizationStateWaitRegistration The user is unregistered and need to accept terms of service and enter their first name and last name to finish registration
type AuthorizationStateWaitRegistration struct {
	tdCommon
	TermsOfService *TermsOfService `json:"terms_of_service"` // Telegram terms of service
}

// MessageType return the string telegram-type of AuthorizationStateWaitRegistration
func (authorizationStateWaitRegistration *AuthorizationStateWaitRegistration) MessageType() string {
	return "authorizationStateWaitRegistration"
}

// NewAuthorizationStateWaitRegistration creates a new AuthorizationStateWaitRegistration
//
// @param termsOfService Telegram terms of service
func NewAuthorizationStateWaitRegistration(termsOfService *TermsOfService) *AuthorizationStateWaitRegistration {
	authorizationStateWaitRegistrationTemp := AuthorizationStateWaitRegistration{
		tdCommon:       tdCommon{Type: "authorizationStateWaitRegistration"},
		TermsOfService: termsOfService,
	}

	return &authorizationStateWaitRegistrationTemp
}

// GetAuthorizationStateEnum return the enum type of this object
func (authorizationStateWaitRegistration *AuthorizationStateWaitRegistration) GetAuthorizationStateEnum() AuthorizationStateEnum {
	return AuthorizationStateWaitRegistrationType
}

// AuthorizationStateWaitPassword The user has been authorized, but needs to enter a password to start using the application
type AuthorizationStateWaitPassword struct {
	tdCommon
	PasswordHint                string `json:"password_hint"`                  // Hint for the password; may be empty
	HasRecoveryEmailAddress     bool   `json:"has_recovery_email_address"`     // True, if a recovery email address has been set up
	RecoveryEmailAddressPattern string `json:"recovery_email_address_pattern"` // Pattern of the email address to which the recovery email was sent; empty until a recovery email has been sent
}

// MessageType return the string telegram-type of AuthorizationStateWaitPassword
func (authorizationStateWaitPassword *AuthorizationStateWaitPassword) MessageType() string {
	return "authorizationStateWaitPassword"
}

// NewAuthorizationStateWaitPassword creates a new AuthorizationStateWaitPassword
//
// @param passwordHint Hint for the password; may be empty
// @param hasRecoveryEmailAddress True, if a recovery email address has been set up
// @param recoveryEmailAddressPattern Pattern of the email address to which the recovery email was sent; empty until a recovery email has been sent
func NewAuthorizationStateWaitPassword(passwordHint string, hasRecoveryEmailAddress bool, recoveryEmailAddressPattern string) *AuthorizationStateWaitPassword {
	authorizationStateWaitPasswordTemp := AuthorizationStateWaitPassword{
		tdCommon:                    tdCommon{Type: "authorizationStateWaitPassword"},
		PasswordHint:                passwordHint,
		HasRecoveryEmailAddress:     hasRecoveryEmailAddress,
		RecoveryEmailAddressPattern: recoveryEmailAddressPattern,
	}

	return &authorizationStateWaitPasswordTemp
}

// GetAuthorizationStateEnum return the enum type of this object
func (authorizationStateWaitPassword *AuthorizationStateWaitPassword) GetAuthorizationStateEnum() AuthorizationStateEnum {
	return AuthorizationStateWaitPasswordType
}

// AuthorizationStateReady The user has been successfully authorized. TDLib is now ready to answer queries
type AuthorizationStateReady struct {
	tdCommon
}

// MessageType return the string telegram-type of AuthorizationStateReady
func (authorizationStateReady *AuthorizationStateReady) MessageType() string {
	return "authorizationStateReady"
}

// NewAuthorizationStateReady creates a new AuthorizationStateReady
//
func NewAuthorizationStateReady() *AuthorizationStateReady {
	authorizationStateReadyTemp := AuthorizationStateReady{
		tdCommon: tdCommon{Type: "authorizationStateReady"},
	}

	return &authorizationStateReadyTemp
}

// GetAuthorizationStateEnum return the enum type of this object
func (authorizationStateReady *AuthorizationStateReady) GetAuthorizationStateEnum() AuthorizationStateEnum {
	return AuthorizationStateReadyType
}

// AuthorizationStateLoggingOut The user is currently logging out
type AuthorizationStateLoggingOut struct {
	tdCommon
}

// MessageType return the string telegram-type of AuthorizationStateLoggingOut
func (authorizationStateLoggingOut *AuthorizationStateLoggingOut) MessageType() string {
	return "authorizationStateLoggingOut"
}

// NewAuthorizationStateLoggingOut creates a new AuthorizationStateLoggingOut
//
func NewAuthorizationStateLoggingOut() *AuthorizationStateLoggingOut {
	authorizationStateLoggingOutTemp := AuthorizationStateLoggingOut{
		tdCommon: tdCommon{Type: "authorizationStateLoggingOut"},
	}

	return &authorizationStateLoggingOutTemp
}

// GetAuthorizationStateEnum return the enum type of this object
func (authorizationStateLoggingOut *AuthorizationStateLoggingOut) GetAuthorizationStateEnum() AuthorizationStateEnum {
	return AuthorizationStateLoggingOutType
}

// AuthorizationStateClosing TDLib is closing, all subsequent queries will be answered with the error 500. Note that closing TDLib can take a while. All resources will be freed only after authorizationStateClosed has been received
type AuthorizationStateClosing struct {
	tdCommon
}

// MessageType return the string telegram-type of AuthorizationStateClosing
func (authorizationStateClosing *AuthorizationStateClosing) MessageType() string {
	return "authorizationStateClosing"
}

// NewAuthorizationStateClosing creates a new AuthorizationStateClosing
//
func NewAuthorizationStateClosing() *AuthorizationStateClosing {
	authorizationStateClosingTemp := AuthorizationStateClosing{
		tdCommon: tdCommon{Type: "authorizationStateClosing"},
	}

	return &authorizationStateClosingTemp
}

// GetAuthorizationStateEnum return the enum type of this object
func (authorizationStateClosing *AuthorizationStateClosing) GetAuthorizationStateEnum() AuthorizationStateEnum {
	return AuthorizationStateClosingType
}

// AuthorizationStateClosed TDLib client is in its final state. All databases are closed and all resources are released. No other updates will be received after this. All queries will be responded to with error code 500. To continue working, one must create a new instance of the TDLib client
type AuthorizationStateClosed struct {
	tdCommon
}

// MessageType return the string telegram-type of AuthorizationStateClosed
func (authorizationStateClosed *AuthorizationStateClosed) MessageType() string {
	return "authorizationStateClosed"
}

// NewAuthorizationStateClosed creates a new AuthorizationStateClosed
//
func NewAuthorizationStateClosed() *AuthorizationStateClosed {
	authorizationStateClosedTemp := AuthorizationStateClosed{
		tdCommon: tdCommon{Type: "authorizationStateClosed"},
	}

	return &authorizationStateClosedTemp
}

// GetAuthorizationStateEnum return the enum type of this object
func (authorizationStateClosed *AuthorizationStateClosed) GetAuthorizationStateEnum() AuthorizationStateEnum {
	return AuthorizationStateClosedType
}

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
