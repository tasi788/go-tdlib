package tdlib

import (
	"encoding/json"
)

// InputIdentityDocument An identity document to be saved to Telegram Passport
type InputIdentityDocument struct {
	tdCommon
	Number      string      `json:"number"`       // Document number; 1-24 characters
	ExpiryDate  *Date       `json:"expiry_date"`  // Document expiry date; pass null if not applicable
	FrontSide   InputFile   `json:"front_side"`   // Front side of the document
	ReverseSide InputFile   `json:"reverse_side"` // Reverse side of the document; only for driver license and identity card; pass null otherwise
	Selfie      InputFile   `json:"selfie"`       // Selfie with the document; pass null if unavailable
	Translation []InputFile `json:"translation"`  // List of files containing a certified English translation of the document
}

// MessageType return the string telegram-type of InputIdentityDocument
func (inputIdentityDocument *InputIdentityDocument) MessageType() string {
	return "inputIdentityDocument"
}

// NewInputIdentityDocument creates a new InputIdentityDocument
//
// @param number Document number; 1-24 characters
// @param expiryDate Document expiry date; pass null if not applicable
// @param frontSide Front side of the document
// @param reverseSide Reverse side of the document; only for driver license and identity card; pass null otherwise
// @param selfie Selfie with the document; pass null if unavailable
// @param translation List of files containing a certified English translation of the document
func NewInputIdentityDocument(number string, expiryDate *Date, frontSide InputFile, reverseSide InputFile, selfie InputFile, translation []InputFile) *InputIdentityDocument {
	inputIdentityDocumentTemp := InputIdentityDocument{
		tdCommon:    tdCommon{Type: "inputIdentityDocument"},
		Number:      number,
		ExpiryDate:  expiryDate,
		FrontSide:   frontSide,
		ReverseSide: reverseSide,
		Selfie:      selfie,
		Translation: translation,
	}

	return &inputIdentityDocumentTemp
}

// UnmarshalJSON unmarshal to json
func (inputIdentityDocument *InputIdentityDocument) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	tempObj := struct {
		tdCommon
		Number      string      `json:"number"`      // Document number; 1-24 characters
		ExpiryDate  *Date       `json:"expiry_date"` // Document expiry date; pass null if not applicable
		Translation []InputFile `json:"translation"` // List of files containing a certified English translation of the document
	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	inputIdentityDocument.tdCommon = tempObj.tdCommon
	inputIdentityDocument.Number = tempObj.Number
	inputIdentityDocument.ExpiryDate = tempObj.ExpiryDate
	inputIdentityDocument.Translation = tempObj.Translation

	fieldFrontSide, _ := unmarshalInputFile(objMap["front_side"])
	inputIdentityDocument.FrontSide = fieldFrontSide

	fieldReverseSide, _ := unmarshalInputFile(objMap["reverse_side"])
	inputIdentityDocument.ReverseSide = fieldReverseSide

	fieldSelfie, _ := unmarshalInputFile(objMap["selfie"])
	inputIdentityDocument.Selfie = fieldSelfie

	return nil
}
