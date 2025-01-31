package tdlib

// InputPersonalDocument A personal document to be saved to Telegram Passport
type InputPersonalDocument struct {
	tdCommon
	Files       []InputFile `json:"files"`       // List of files containing the pages of the document
	Translation []InputFile `json:"translation"` // List of files containing a certified English translation of the document
}

// MessageType return the string telegram-type of InputPersonalDocument
func (inputPersonalDocument *InputPersonalDocument) MessageType() string {
	return "inputPersonalDocument"
}

// NewInputPersonalDocument creates a new InputPersonalDocument
//
// @param files List of files containing the pages of the document
// @param translation List of files containing a certified English translation of the document
func NewInputPersonalDocument(files []InputFile, translation []InputFile) *InputPersonalDocument {
	inputPersonalDocumentTemp := InputPersonalDocument{
		tdCommon:    tdCommon{Type: "inputPersonalDocument"},
		Files:       files,
		Translation: translation,
	}

	return &inputPersonalDocumentTemp
}
