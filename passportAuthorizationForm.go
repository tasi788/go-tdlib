package tdlib

import (
	"encoding/json"
	"fmt"
)

// PassportAuthorizationForm Contains information about a Telegram Passport authorization form that was requested
type PassportAuthorizationForm struct {
	tdCommon
	Id               int32                     `json:"id"`                 // Unique identifier of the authorization form
	RequiredElements []PassportRequiredElement `json:"required_elements"`  // Telegram Passport elements that must be provided to complete the form
	PrivacyPolicyUrl string                    `json:"privacy_policy_url"` // URL for the privacy policy of the service; may be empty
}

// MessageType return the string telegram-type of PassportAuthorizationForm
func (passportAuthorizationForm *PassportAuthorizationForm) MessageType() string {
	return "passportAuthorizationForm"
}

// NewPassportAuthorizationForm creates a new PassportAuthorizationForm
//
// @param id Unique identifier of the authorization form
// @param requiredElements Telegram Passport elements that must be provided to complete the form
// @param privacyPolicyUrl URL for the privacy policy of the service; may be empty
func NewPassportAuthorizationForm(id int32, requiredElements []PassportRequiredElement, privacyPolicyUrl string) *PassportAuthorizationForm {
	passportAuthorizationFormTemp := PassportAuthorizationForm{
		tdCommon:         tdCommon{Type: "passportAuthorizationForm"},
		Id:               id,
		RequiredElements: requiredElements,
		PrivacyPolicyUrl: privacyPolicyUrl,
	}

	return &passportAuthorizationFormTemp
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
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var passportAuthorizationForm PassportAuthorizationForm
	err = json.Unmarshal(result.Raw, &passportAuthorizationForm)
	return &passportAuthorizationForm, err
}
