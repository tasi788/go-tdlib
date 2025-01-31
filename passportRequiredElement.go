package tdlib

// PassportRequiredElement Contains a description of the required Telegram Passport element that was requested by a service
type PassportRequiredElement struct {
	tdCommon
	SuitableElements []PassportSuitableElement `json:"suitable_elements"` // List of Telegram Passport elements any of which is enough to provide
}

// MessageType return the string telegram-type of PassportRequiredElement
func (passportRequiredElement *PassportRequiredElement) MessageType() string {
	return "passportRequiredElement"
}

// NewPassportRequiredElement creates a new PassportRequiredElement
//
// @param suitableElements List of Telegram Passport elements any of which is enough to provide
func NewPassportRequiredElement(suitableElements []PassportSuitableElement) *PassportRequiredElement {
	passportRequiredElementTemp := PassportRequiredElement{
		tdCommon:         tdCommon{Type: "passportRequiredElement"},
		SuitableElements: suitableElements,
	}

	return &passportRequiredElementTemp
}
