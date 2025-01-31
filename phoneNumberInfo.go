package tdlib

import (
	"encoding/json"
	"fmt"
)

// PhoneNumberInfo Contains information about a phone number
type PhoneNumberInfo struct {
	tdCommon
	Country              *CountryInfo `json:"country"`                // Information about the country to which the phone number belongs; may be null
	CountryCallingCode   string       `json:"country_calling_code"`   // The part of the phone number denoting country calling code or its part
	FormattedPhoneNumber string       `json:"formatted_phone_number"` // The phone number without country calling code formatted accordingly to local rules. Expected digits are returned as '-', but even more digits might be entered by the user
}

// MessageType return the string telegram-type of PhoneNumberInfo
func (phoneNumberInfo *PhoneNumberInfo) MessageType() string {
	return "phoneNumberInfo"
}

// NewPhoneNumberInfo creates a new PhoneNumberInfo
//
// @param country Information about the country to which the phone number belongs; may be null
// @param countryCallingCode The part of the phone number denoting country calling code or its part
// @param formattedPhoneNumber The phone number without country calling code formatted accordingly to local rules. Expected digits are returned as '-', but even more digits might be entered by the user
func NewPhoneNumberInfo(country *CountryInfo, countryCallingCode string, formattedPhoneNumber string) *PhoneNumberInfo {
	phoneNumberInfoTemp := PhoneNumberInfo{
		tdCommon:             tdCommon{Type: "phoneNumberInfo"},
		Country:              country,
		CountryCallingCode:   countryCallingCode,
		FormattedPhoneNumber: formattedPhoneNumber,
	}

	return &phoneNumberInfoTemp
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
