package tdlib

// CountryInfo Contains information about a country
type CountryInfo struct {
	tdCommon
	CountryCode  string   `json:"country_code"`  // A two-letter ISO 3166-1 alpha-2 country code
	Name         string   `json:"name"`          // Native name of the country
	EnglishName  string   `json:"english_name"`  // English name of the country
	IsHidden     bool     `json:"is_hidden"`     // True, if the country must be hidden from the list of all countries
	CallingCodes []string `json:"calling_codes"` // List of country calling codes
}

// MessageType return the string telegram-type of CountryInfo
func (countryInfo *CountryInfo) MessageType() string {
	return "countryInfo"
}

// NewCountryInfo creates a new CountryInfo
//
// @param countryCode A two-letter ISO 3166-1 alpha-2 country code
// @param name Native name of the country
// @param englishName English name of the country
// @param isHidden True, if the country must be hidden from the list of all countries
// @param callingCodes List of country calling codes
func NewCountryInfo(countryCode string, name string, englishName string, isHidden bool, callingCodes []string) *CountryInfo {
	countryInfoTemp := CountryInfo{
		tdCommon:     tdCommon{Type: "countryInfo"},
		CountryCode:  countryCode,
		Name:         name,
		EnglishName:  englishName,
		IsHidden:     isHidden,
		CallingCodes: callingCodes,
	}

	return &countryInfoTemp
}
