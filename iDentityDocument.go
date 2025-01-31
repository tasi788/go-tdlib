package tdlib

// IdentityDocument An identity document
type IdentityDocument struct {
	tdCommon
	Number      string      `json:"number"`       // Document number; 1-24 characters
	ExpiryDate  *Date       `json:"expiry_date"`  // Document expiry date; may be null if not applicable
	FrontSide   *DatedFile  `json:"front_side"`   // Front side of the document
	ReverseSide *DatedFile  `json:"reverse_side"` // Reverse side of the document; only for driver license and identity card; may be null
	Selfie      *DatedFile  `json:"selfie"`       // Selfie with the document; may be null
	Translation []DatedFile `json:"translation"`  // List of files containing a certified English translation of the document
}

// MessageType return the string telegram-type of IdentityDocument
func (identityDocument *IdentityDocument) MessageType() string {
	return "identityDocument"
}

// NewIdentityDocument creates a new IdentityDocument
//
// @param number Document number; 1-24 characters
// @param expiryDate Document expiry date; may be null if not applicable
// @param frontSide Front side of the document
// @param reverseSide Reverse side of the document; only for driver license and identity card; may be null
// @param selfie Selfie with the document; may be null
// @param translation List of files containing a certified English translation of the document
func NewIdentityDocument(number string, expiryDate *Date, frontSide *DatedFile, reverseSide *DatedFile, selfie *DatedFile, translation []DatedFile) *IdentityDocument {
	identityDocumentTemp := IdentityDocument{
		tdCommon:    tdCommon{Type: "identityDocument"},
		Number:      number,
		ExpiryDate:  expiryDate,
		FrontSide:   frontSide,
		ReverseSide: reverseSide,
		Selfie:      selfie,
		Translation: translation,
	}

	return &identityDocumentTemp
}
