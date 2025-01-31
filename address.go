package tdlib

// Address Describes an address
type Address struct {
	tdCommon
	CountryCode string `json:"country_code"` // A two-letter ISO 3166-1 alpha-2 country code
	State       string `json:"state"`        // State, if applicable
	City        string `json:"city"`         // City
	StreetLine1 string `json:"street_line1"` // First line of the address
	StreetLine2 string `json:"street_line2"` // Second line of the address
	PostalCode  string `json:"postal_code"`  // Address postal code
}

// MessageType return the string telegram-type of Address
func (address *Address) MessageType() string {
	return "address"
}

// NewAddress creates a new Address
//
// @param countryCode A two-letter ISO 3166-1 alpha-2 country code
// @param state State, if applicable
// @param city City
// @param streetLine1 First line of the address
// @param streetLine2 Second line of the address
// @param postalCode Address postal code
func NewAddress(countryCode string, state string, city string, streetLine1 string, streetLine2 string, postalCode string) *Address {
	addressTemp := Address{
		tdCommon:    tdCommon{Type: "address"},
		CountryCode: countryCode,
		State:       state,
		City:        city,
		StreetLine1: streetLine1,
		StreetLine2: streetLine2,
		PostalCode:  postalCode,
	}

	return &addressTemp
}
