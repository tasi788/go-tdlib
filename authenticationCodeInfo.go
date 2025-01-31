package tdlib

import (
	"encoding/json"
	"fmt"
)

// AuthenticationCodeInfo Information about the authentication code that was sent
type AuthenticationCodeInfo struct {
	tdCommon
	PhoneNumber string                 `json:"phone_number"` // A phone number that is being authenticated
	Type        AuthenticationCodeType `json:"type"`         // The way the code was sent to the user
	NextType    AuthenticationCodeType `json:"next_type"`    // The way the next code will be sent to the user; may be null
	Timeout     int32                  `json:"timeout"`      // Timeout before the code can be re-sent, in seconds
}

// MessageType return the string telegram-type of AuthenticationCodeInfo
func (authenticationCodeInfo *AuthenticationCodeInfo) MessageType() string {
	return "authenticationCodeInfo"
}

// NewAuthenticationCodeInfo creates a new AuthenticationCodeInfo
//
// @param phoneNumber A phone number that is being authenticated
// @param typeParam The way the code was sent to the user
// @param nextType The way the next code will be sent to the user; may be null
// @param timeout Timeout before the code can be re-sent, in seconds
func NewAuthenticationCodeInfo(phoneNumber string, typeParam AuthenticationCodeType, nextType AuthenticationCodeType, timeout int32) *AuthenticationCodeInfo {
	authenticationCodeInfoTemp := AuthenticationCodeInfo{
		tdCommon:    tdCommon{Type: "authenticationCodeInfo"},
		PhoneNumber: phoneNumber,
		Type:        typeParam,
		NextType:    nextType,
		Timeout:     timeout,
	}

	return &authenticationCodeInfoTemp
}

// UnmarshalJSON unmarshal to json
func (authenticationCodeInfo *AuthenticationCodeInfo) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	tempObj := struct {
		tdCommon
		PhoneNumber string `json:"phone_number"` // A phone number that is being authenticated
		Timeout     int32  `json:"timeout"`      // Timeout before the code can be re-sent, in seconds
	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	authenticationCodeInfo.tdCommon = tempObj.tdCommon
	authenticationCodeInfo.PhoneNumber = tempObj.PhoneNumber
	authenticationCodeInfo.Timeout = tempObj.Timeout

	fieldType, _ := unmarshalAuthenticationCodeType(objMap["type"])
	authenticationCodeInfo.Type = fieldType

	fieldNextType, _ := unmarshalAuthenticationCodeType(objMap["next_type"])
	authenticationCodeInfo.NextType = fieldNextType

	return nil
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
