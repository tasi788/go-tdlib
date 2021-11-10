package tdlib

// SavedCredentials Contains information about saved card credentials
type SavedCredentials struct {
	tdCommon
	Id    string `json:"id"`    // Unique identifier of the saved credentials
	Title string `json:"title"` // Title of the saved credentials
}

// MessageType return the string telegram-type of SavedCredentials
func (savedCredentials *SavedCredentials) MessageType() string {
	return "savedCredentials"
}

// NewSavedCredentials creates a new SavedCredentials
//
// @param id Unique identifier of the saved credentials
// @param title Title of the saved credentials
func NewSavedCredentials(id string, title string) *SavedCredentials {
	savedCredentialsTemp := SavedCredentials{
		tdCommon: tdCommon{Type: "savedCredentials"},
		Id:       id,
		Title:    title,
	}

	return &savedCredentialsTemp
}
