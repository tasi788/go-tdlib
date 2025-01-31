package tdlib

import (
	"encoding/json"
)

// EncryptedPassportElement Contains information about an encrypted Telegram Passport element; for bots only
type EncryptedPassportElement struct {
	tdCommon
	Type        PassportElementType `json:"type"`         // Type of Telegram Passport element
	Data        []byte              `json:"data"`         // Encrypted JSON-encoded data about the user
	FrontSide   *DatedFile          `json:"front_side"`   // The front side of an identity document
	ReverseSide *DatedFile          `json:"reverse_side"` // The reverse side of an identity document; may be null
	Selfie      *DatedFile          `json:"selfie"`       // Selfie with the document; may be null
	Translation []DatedFile         `json:"translation"`  // List of files containing a certified English translation of the document
	Files       []DatedFile         `json:"files"`        // List of attached files
	Value       string              `json:"value"`        // Unencrypted data, phone number or email address
	Hash        string              `json:"hash"`         // Hash of the entire element
}

// MessageType return the string telegram-type of EncryptedPassportElement
func (encryptedPassportElement *EncryptedPassportElement) MessageType() string {
	return "encryptedPassportElement"
}

// NewEncryptedPassportElement creates a new EncryptedPassportElement
//
// @param typeParam Type of Telegram Passport element
// @param data Encrypted JSON-encoded data about the user
// @param frontSide The front side of an identity document
// @param reverseSide The reverse side of an identity document; may be null
// @param selfie Selfie with the document; may be null
// @param translation List of files containing a certified English translation of the document
// @param files List of attached files
// @param value Unencrypted data, phone number or email address
// @param hash Hash of the entire element
func NewEncryptedPassportElement(typeParam PassportElementType, data []byte, frontSide *DatedFile, reverseSide *DatedFile, selfie *DatedFile, translation []DatedFile, files []DatedFile, value string, hash string) *EncryptedPassportElement {
	encryptedPassportElementTemp := EncryptedPassportElement{
		tdCommon:    tdCommon{Type: "encryptedPassportElement"},
		Type:        typeParam,
		Data:        data,
		FrontSide:   frontSide,
		ReverseSide: reverseSide,
		Selfie:      selfie,
		Translation: translation,
		Files:       files,
		Value:       value,
		Hash:        hash,
	}

	return &encryptedPassportElementTemp
}

// UnmarshalJSON unmarshal to json
func (encryptedPassportElement *EncryptedPassportElement) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	tempObj := struct {
		tdCommon
		Data        []byte      `json:"data"`         // Encrypted JSON-encoded data about the user
		FrontSide   *DatedFile  `json:"front_side"`   // The front side of an identity document
		ReverseSide *DatedFile  `json:"reverse_side"` // The reverse side of an identity document; may be null
		Selfie      *DatedFile  `json:"selfie"`       // Selfie with the document; may be null
		Translation []DatedFile `json:"translation"`  // List of files containing a certified English translation of the document
		Files       []DatedFile `json:"files"`        // List of attached files
		Value       string      `json:"value"`        // Unencrypted data, phone number or email address
		Hash        string      `json:"hash"`         // Hash of the entire element
	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	encryptedPassportElement.tdCommon = tempObj.tdCommon
	encryptedPassportElement.Data = tempObj.Data
	encryptedPassportElement.FrontSide = tempObj.FrontSide
	encryptedPassportElement.ReverseSide = tempObj.ReverseSide
	encryptedPassportElement.Selfie = tempObj.Selfie
	encryptedPassportElement.Translation = tempObj.Translation
	encryptedPassportElement.Files = tempObj.Files
	encryptedPassportElement.Value = tempObj.Value
	encryptedPassportElement.Hash = tempObj.Hash

	fieldType, _ := unmarshalPassportElementType(objMap["type"])
	encryptedPassportElement.Type = fieldType

	return nil
}
