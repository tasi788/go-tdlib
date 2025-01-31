package tdlib

// PersonalDocument A personal document, containing some information about a user
type PersonalDocument struct {
	tdCommon
	Files       []DatedFile `json:"files"`       // List of files containing the pages of the document
	Translation []DatedFile `json:"translation"` // List of files containing a certified English translation of the document
}

// MessageType return the string telegram-type of PersonalDocument
func (personalDocument *PersonalDocument) MessageType() string {
	return "personalDocument"
}

// NewPersonalDocument creates a new PersonalDocument
//
// @param files List of files containing the pages of the document
// @param translation List of files containing a certified English translation of the document
func NewPersonalDocument(files []DatedFile, translation []DatedFile) *PersonalDocument {
	personalDocumentTemp := PersonalDocument{
		tdCommon:    tdCommon{Type: "personalDocument"},
		Files:       files,
		Translation: translation,
	}

	return &personalDocumentTemp
}
