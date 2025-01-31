package tdlib

import (
	"encoding/json"
	"fmt"
)

// Countries Contains information about countries
type Countries struct {
	tdCommon
	Countries []CountryInfo `json:"countries"` // The list of countries
}

// MessageType return the string telegram-type of Countries
func (countries *Countries) MessageType() string {
	return "countries"
}

// NewCountries creates a new Countries
//
// @param countries The list of countries
func NewCountries(countries []CountryInfo) *Countries {
	countriesTemp := Countries{
		tdCommon:  tdCommon{Type: "countries"},
		Countries: countries,
	}

	return &countriesTemp
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
