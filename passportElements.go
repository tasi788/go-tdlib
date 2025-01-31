package tdlib

import (
	"encoding/json"
	"fmt"
)

// PassportElements Contains information about saved Telegram Passport elements
type PassportElements struct {
	tdCommon
	Elements []PassportElement `json:"elements"` // Telegram Passport elements
}

// MessageType return the string telegram-type of PassportElements
func (passportElements *PassportElements) MessageType() string {
	return "passportElements"
}

// NewPassportElements creates a new PassportElements
//
// @param elements Telegram Passport elements
func NewPassportElements(elements []PassportElement) *PassportElements {
	passportElementsTemp := PassportElements{
		tdCommon: tdCommon{Type: "passportElements"},
		Elements: elements,
	}

	return &passportElementsTemp
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
