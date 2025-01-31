package tdlib

// TermsOfService Contains Telegram terms of service
type TermsOfService struct {
	tdCommon
	Text       *FormattedText `json:"text"`         // Text of the terms of service
	MinUserAge int32          `json:"min_user_age"` // The minimum age of a user to be able to accept the terms; 0 if any
	ShowPopup  bool           `json:"show_popup"`   // True, if a blocking popup with terms of service must be shown to the user
}

// MessageType return the string telegram-type of TermsOfService
func (termsOfService *TermsOfService) MessageType() string {
	return "termsOfService"
}

// NewTermsOfService creates a new TermsOfService
//
// @param text Text of the terms of service
// @param minUserAge The minimum age of a user to be able to accept the terms; 0 if any
// @param showPopup True, if a blocking popup with terms of service must be shown to the user
func NewTermsOfService(text *FormattedText, minUserAge int32, showPopup bool) *TermsOfService {
	termsOfServiceTemp := TermsOfService{
		tdCommon:   tdCommon{Type: "termsOfService"},
		Text:       text,
		MinUserAge: minUserAge,
		ShowPopup:  showPopup,
	}

	return &termsOfServiceTemp
}
