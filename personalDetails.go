package tdlib

// PersonalDetails Contains the user's personal details
type PersonalDetails struct {
	tdCommon
	FirstName            string `json:"first_name"`             // First name of the user written in English; 1-255 characters
	MiddleName           string `json:"middle_name"`            // Middle name of the user written in English; 0-255 characters
	LastName             string `json:"last_name"`              // Last name of the user written in English; 1-255 characters
	NativeFirstName      string `json:"native_first_name"`      // Native first name of the user; 1-255 characters
	NativeMiddleName     string `json:"native_middle_name"`     // Native middle name of the user; 0-255 characters
	NativeLastName       string `json:"native_last_name"`       // Native last name of the user; 1-255 characters
	Birthdate            *Date  `json:"birthdate"`              // Birthdate of the user
	Gender               string `json:"gender"`                 // Gender of the user, "male" or "female"
	CountryCode          string `json:"country_code"`           // A two-letter ISO 3166-1 alpha-2 country code of the user's country
	ResidenceCountryCode string `json:"residence_country_code"` // A two-letter ISO 3166-1 alpha-2 country code of the user's residence country
}

// MessageType return the string telegram-type of PersonalDetails
func (personalDetails *PersonalDetails) MessageType() string {
	return "personalDetails"
}

// NewPersonalDetails creates a new PersonalDetails
//
// @param firstName First name of the user written in English; 1-255 characters
// @param middleName Middle name of the user written in English; 0-255 characters
// @param lastName Last name of the user written in English; 1-255 characters
// @param nativeFirstName Native first name of the user; 1-255 characters
// @param nativeMiddleName Native middle name of the user; 0-255 characters
// @param nativeLastName Native last name of the user; 1-255 characters
// @param birthdate Birthdate of the user
// @param gender Gender of the user, "male" or "female"
// @param countryCode A two-letter ISO 3166-1 alpha-2 country code of the user's country
// @param residenceCountryCode A two-letter ISO 3166-1 alpha-2 country code of the user's residence country
func NewPersonalDetails(firstName string, middleName string, lastName string, nativeFirstName string, nativeMiddleName string, nativeLastName string, birthdate *Date, gender string, countryCode string, residenceCountryCode string) *PersonalDetails {
	personalDetailsTemp := PersonalDetails{
		tdCommon:             tdCommon{Type: "personalDetails"},
		FirstName:            firstName,
		MiddleName:           middleName,
		LastName:             lastName,
		NativeFirstName:      nativeFirstName,
		NativeMiddleName:     nativeMiddleName,
		NativeLastName:       nativeLastName,
		Birthdate:            birthdate,
		Gender:               gender,
		CountryCode:          countryCode,
		ResidenceCountryCode: residenceCountryCode,
	}

	return &personalDetailsTemp
}
