package tdlib

import (
	"encoding/json"
	"fmt"
)

// EmailAddressAuthenticationCodeInfo Information about the email address authentication code that was sent
type EmailAddressAuthenticationCodeInfo struct {
	tdCommon
	EmailAddressPattern string `json:"email_address_pattern"` // Pattern of the email address to which an authentication code was sent
	Length              int32  `json:"length"`                // Length of the code; 0 if unknown
}

// MessageType return the string telegram-type of EmailAddressAuthenticationCodeInfo
func (emailAddressAuthenticationCodeInfo *EmailAddressAuthenticationCodeInfo) MessageType() string {
	return "emailAddressAuthenticationCodeInfo"
}

// NewEmailAddressAuthenticationCodeInfo creates a new EmailAddressAuthenticationCodeInfo
//
// @param emailAddressPattern Pattern of the email address to which an authentication code was sent
// @param length Length of the code; 0 if unknown
func NewEmailAddressAuthenticationCodeInfo(emailAddressPattern string, length int32) *EmailAddressAuthenticationCodeInfo {
	emailAddressAuthenticationCodeInfoTemp := EmailAddressAuthenticationCodeInfo{
		tdCommon:            tdCommon{Type: "emailAddressAuthenticationCodeInfo"},
		EmailAddressPattern: emailAddressPattern,
		Length:              length,
	}

	return &emailAddressAuthenticationCodeInfoTemp
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
