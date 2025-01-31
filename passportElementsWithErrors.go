package tdlib

import (
	"encoding/json"
	"fmt"
)

// PassportElementsWithErrors Contains information about a Telegram Passport elements and corresponding errors
type PassportElementsWithErrors struct {
	tdCommon
	Elements []PassportElement      `json:"elements"` // Telegram Passport elements
	Errors   []PassportElementError `json:"errors"`   // Errors in the elements that are already available
}

// MessageType return the string telegram-type of PassportElementsWithErrors
func (passportElementsWithErrors *PassportElementsWithErrors) MessageType() string {
	return "passportElementsWithErrors"
}

// NewPassportElementsWithErrors creates a new PassportElementsWithErrors
//
// @param elements Telegram Passport elements
// @param errors Errors in the elements that are already available
func NewPassportElementsWithErrors(elements []PassportElement, errors []PassportElementError) *PassportElementsWithErrors {
	passportElementsWithErrorsTemp := PassportElementsWithErrors{
		tdCommon: tdCommon{Type: "passportElementsWithErrors"},
		Elements: elements,
		Errors:   errors,
	}

	return &passportElementsWithErrorsTemp
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
