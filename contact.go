package tdlib

// Contact Describes a user contact
type Contact struct {
	tdCommon
	PhoneNumber string `json:"phone_number"` // Phone number of the user
	FirstName   string `json:"first_name"`   // First name of the user; 1-255 characters in length
	LastName    string `json:"last_name"`    // Last name of the user
	Vcard       string `json:"vcard"`        // Additional data about the user in a form of vCard; 0-2048 bytes in length
	UserId      int64  `json:"user_id"`      // Identifier of the user, if known; otherwise 0
}

// MessageType return the string telegram-type of Contact
func (contact *Contact) MessageType() string {
	return "contact"
}

// NewContact creates a new Contact
//
// @param phoneNumber Phone number of the user
// @param firstName First name of the user; 1-255 characters in length
// @param lastName Last name of the user
// @param vcard Additional data about the user in a form of vCard; 0-2048 bytes in length
// @param userId Identifier of the user, if known; otherwise 0
func NewContact(phoneNumber string, firstName string, lastName string, vcard string, userId int64) *Contact {
	contactTemp := Contact{
		tdCommon:    tdCommon{Type: "contact"},
		PhoneNumber: phoneNumber,
		FirstName:   firstName,
		LastName:    lastName,
		Vcard:       vcard,
		UserId:      userId,
	}

	return &contactTemp
}
