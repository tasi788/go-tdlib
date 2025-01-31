package tdlib

import (
	"encoding/json"
	"fmt"
)

// PassportElementErrorSource Contains the description of an error in a Telegram Passport element
type PassportElementErrorSource interface {
	GetPassportElementErrorSourceEnum() PassportElementErrorSourceEnum
}

// PassportElementErrorSourceEnum Alias for abstract PassportElementErrorSource 'Sub-Classes', used as constant-enum here
type PassportElementErrorSourceEnum string

// PassportElementErrorSource enums
const (
	PassportElementErrorSourceUnspecifiedType      PassportElementErrorSourceEnum = "passportElementErrorSourceUnspecified"
	PassportElementErrorSourceDataFieldType        PassportElementErrorSourceEnum = "passportElementErrorSourceDataField"
	PassportElementErrorSourceFrontSideType        PassportElementErrorSourceEnum = "passportElementErrorSourceFrontSide"
	PassportElementErrorSourceReverseSideType      PassportElementErrorSourceEnum = "passportElementErrorSourceReverseSide"
	PassportElementErrorSourceSelfieType           PassportElementErrorSourceEnum = "passportElementErrorSourceSelfie"
	PassportElementErrorSourceTranslationFileType  PassportElementErrorSourceEnum = "passportElementErrorSourceTranslationFile"
	PassportElementErrorSourceTranslationFilesType PassportElementErrorSourceEnum = "passportElementErrorSourceTranslationFiles"
	PassportElementErrorSourceFileType             PassportElementErrorSourceEnum = "passportElementErrorSourceFile"
	PassportElementErrorSourceFilesType            PassportElementErrorSourceEnum = "passportElementErrorSourceFiles"
)

func unmarshalPassportElementErrorSource(rawMsg *json.RawMessage) (PassportElementErrorSource, error) {

	if rawMsg == nil {
		return nil, nil
	}
	var objMap map[string]interface{}
	err := json.Unmarshal(*rawMsg, &objMap)
	if err != nil {
		return nil, err
	}

	switch PassportElementErrorSourceEnum(objMap["@type"].(string)) {
	case PassportElementErrorSourceUnspecifiedType:
		var passportElementErrorSourceUnspecified PassportElementErrorSourceUnspecified
		err := json.Unmarshal(*rawMsg, &passportElementErrorSourceUnspecified)
		return &passportElementErrorSourceUnspecified, err

	case PassportElementErrorSourceDataFieldType:
		var passportElementErrorSourceDataField PassportElementErrorSourceDataField
		err := json.Unmarshal(*rawMsg, &passportElementErrorSourceDataField)
		return &passportElementErrorSourceDataField, err

	case PassportElementErrorSourceFrontSideType:
		var passportElementErrorSourceFrontSide PassportElementErrorSourceFrontSide
		err := json.Unmarshal(*rawMsg, &passportElementErrorSourceFrontSide)
		return &passportElementErrorSourceFrontSide, err

	case PassportElementErrorSourceReverseSideType:
		var passportElementErrorSourceReverseSide PassportElementErrorSourceReverseSide
		err := json.Unmarshal(*rawMsg, &passportElementErrorSourceReverseSide)
		return &passportElementErrorSourceReverseSide, err

	case PassportElementErrorSourceSelfieType:
		var passportElementErrorSourceSelfie PassportElementErrorSourceSelfie
		err := json.Unmarshal(*rawMsg, &passportElementErrorSourceSelfie)
		return &passportElementErrorSourceSelfie, err

	case PassportElementErrorSourceTranslationFileType:
		var passportElementErrorSourceTranslationFile PassportElementErrorSourceTranslationFile
		err := json.Unmarshal(*rawMsg, &passportElementErrorSourceTranslationFile)
		return &passportElementErrorSourceTranslationFile, err

	case PassportElementErrorSourceTranslationFilesType:
		var passportElementErrorSourceTranslationFiles PassportElementErrorSourceTranslationFiles
		err := json.Unmarshal(*rawMsg, &passportElementErrorSourceTranslationFiles)
		return &passportElementErrorSourceTranslationFiles, err

	case PassportElementErrorSourceFileType:
		var passportElementErrorSourceFile PassportElementErrorSourceFile
		err := json.Unmarshal(*rawMsg, &passportElementErrorSourceFile)
		return &passportElementErrorSourceFile, err

	case PassportElementErrorSourceFilesType:
		var passportElementErrorSourceFiles PassportElementErrorSourceFiles
		err := json.Unmarshal(*rawMsg, &passportElementErrorSourceFiles)
		return &passportElementErrorSourceFiles, err

	default:
		return nil, fmt.Errorf("Error UnMarshaling, unknown type:" + objMap["@type"].(string))
	}
}

// PassportElementErrorSourceUnspecified The element contains an error in an unspecified place. The error will be considered resolved when new data is added
type PassportElementErrorSourceUnspecified struct {
	tdCommon
}

// MessageType return the string telegram-type of PassportElementErrorSourceUnspecified
func (passportElementErrorSourceUnspecified *PassportElementErrorSourceUnspecified) MessageType() string {
	return "passportElementErrorSourceUnspecified"
}

// NewPassportElementErrorSourceUnspecified creates a new PassportElementErrorSourceUnspecified
//
func NewPassportElementErrorSourceUnspecified() *PassportElementErrorSourceUnspecified {
	passportElementErrorSourceUnspecifiedTemp := PassportElementErrorSourceUnspecified{
		tdCommon: tdCommon{Type: "passportElementErrorSourceUnspecified"},
	}

	return &passportElementErrorSourceUnspecifiedTemp
}

// GetPassportElementErrorSourceEnum return the enum type of this object
func (passportElementErrorSourceUnspecified *PassportElementErrorSourceUnspecified) GetPassportElementErrorSourceEnum() PassportElementErrorSourceEnum {
	return PassportElementErrorSourceUnspecifiedType
}

// PassportElementErrorSourceDataField One of the data fields contains an error. The error will be considered resolved when the value of the field changes
type PassportElementErrorSourceDataField struct {
	tdCommon
	FieldName string `json:"field_name"` // Field name
}

// MessageType return the string telegram-type of PassportElementErrorSourceDataField
func (passportElementErrorSourceDataField *PassportElementErrorSourceDataField) MessageType() string {
	return "passportElementErrorSourceDataField"
}

// NewPassportElementErrorSourceDataField creates a new PassportElementErrorSourceDataField
//
// @param fieldName Field name
func NewPassportElementErrorSourceDataField(fieldName string) *PassportElementErrorSourceDataField {
	passportElementErrorSourceDataFieldTemp := PassportElementErrorSourceDataField{
		tdCommon:  tdCommon{Type: "passportElementErrorSourceDataField"},
		FieldName: fieldName,
	}

	return &passportElementErrorSourceDataFieldTemp
}

// GetPassportElementErrorSourceEnum return the enum type of this object
func (passportElementErrorSourceDataField *PassportElementErrorSourceDataField) GetPassportElementErrorSourceEnum() PassportElementErrorSourceEnum {
	return PassportElementErrorSourceDataFieldType
}

// PassportElementErrorSourceFrontSide The front side of the document contains an error. The error will be considered resolved when the file with the front side changes
type PassportElementErrorSourceFrontSide struct {
	tdCommon
}

// MessageType return the string telegram-type of PassportElementErrorSourceFrontSide
func (passportElementErrorSourceFrontSide *PassportElementErrorSourceFrontSide) MessageType() string {
	return "passportElementErrorSourceFrontSide"
}

// NewPassportElementErrorSourceFrontSide creates a new PassportElementErrorSourceFrontSide
//
func NewPassportElementErrorSourceFrontSide() *PassportElementErrorSourceFrontSide {
	passportElementErrorSourceFrontSideTemp := PassportElementErrorSourceFrontSide{
		tdCommon: tdCommon{Type: "passportElementErrorSourceFrontSide"},
	}

	return &passportElementErrorSourceFrontSideTemp
}

// GetPassportElementErrorSourceEnum return the enum type of this object
func (passportElementErrorSourceFrontSide *PassportElementErrorSourceFrontSide) GetPassportElementErrorSourceEnum() PassportElementErrorSourceEnum {
	return PassportElementErrorSourceFrontSideType
}

// PassportElementErrorSourceReverseSide The reverse side of the document contains an error. The error will be considered resolved when the file with the reverse side changes
type PassportElementErrorSourceReverseSide struct {
	tdCommon
}

// MessageType return the string telegram-type of PassportElementErrorSourceReverseSide
func (passportElementErrorSourceReverseSide *PassportElementErrorSourceReverseSide) MessageType() string {
	return "passportElementErrorSourceReverseSide"
}

// NewPassportElementErrorSourceReverseSide creates a new PassportElementErrorSourceReverseSide
//
func NewPassportElementErrorSourceReverseSide() *PassportElementErrorSourceReverseSide {
	passportElementErrorSourceReverseSideTemp := PassportElementErrorSourceReverseSide{
		tdCommon: tdCommon{Type: "passportElementErrorSourceReverseSide"},
	}

	return &passportElementErrorSourceReverseSideTemp
}

// GetPassportElementErrorSourceEnum return the enum type of this object
func (passportElementErrorSourceReverseSide *PassportElementErrorSourceReverseSide) GetPassportElementErrorSourceEnum() PassportElementErrorSourceEnum {
	return PassportElementErrorSourceReverseSideType
}

// PassportElementErrorSourceSelfie The selfie with the document contains an error. The error will be considered resolved when the file with the selfie changes
type PassportElementErrorSourceSelfie struct {
	tdCommon
}

// MessageType return the string telegram-type of PassportElementErrorSourceSelfie
func (passportElementErrorSourceSelfie *PassportElementErrorSourceSelfie) MessageType() string {
	return "passportElementErrorSourceSelfie"
}

// NewPassportElementErrorSourceSelfie creates a new PassportElementErrorSourceSelfie
//
func NewPassportElementErrorSourceSelfie() *PassportElementErrorSourceSelfie {
	passportElementErrorSourceSelfieTemp := PassportElementErrorSourceSelfie{
		tdCommon: tdCommon{Type: "passportElementErrorSourceSelfie"},
	}

	return &passportElementErrorSourceSelfieTemp
}

// GetPassportElementErrorSourceEnum return the enum type of this object
func (passportElementErrorSourceSelfie *PassportElementErrorSourceSelfie) GetPassportElementErrorSourceEnum() PassportElementErrorSourceEnum {
	return PassportElementErrorSourceSelfieType
}

// PassportElementErrorSourceTranslationFile One of files with the translation of the document contains an error. The error will be considered resolved when the file changes
type PassportElementErrorSourceTranslationFile struct {
	tdCommon
	FileIndex int32 `json:"file_index"` // Index of a file with the error
}

// MessageType return the string telegram-type of PassportElementErrorSourceTranslationFile
func (passportElementErrorSourceTranslationFile *PassportElementErrorSourceTranslationFile) MessageType() string {
	return "passportElementErrorSourceTranslationFile"
}

// NewPassportElementErrorSourceTranslationFile creates a new PassportElementErrorSourceTranslationFile
//
// @param fileIndex Index of a file with the error
func NewPassportElementErrorSourceTranslationFile(fileIndex int32) *PassportElementErrorSourceTranslationFile {
	passportElementErrorSourceTranslationFileTemp := PassportElementErrorSourceTranslationFile{
		tdCommon:  tdCommon{Type: "passportElementErrorSourceTranslationFile"},
		FileIndex: fileIndex,
	}

	return &passportElementErrorSourceTranslationFileTemp
}

// GetPassportElementErrorSourceEnum return the enum type of this object
func (passportElementErrorSourceTranslationFile *PassportElementErrorSourceTranslationFile) GetPassportElementErrorSourceEnum() PassportElementErrorSourceEnum {
	return PassportElementErrorSourceTranslationFileType
}

// PassportElementErrorSourceTranslationFiles The translation of the document contains an error. The error will be considered resolved when the list of translation files changes
type PassportElementErrorSourceTranslationFiles struct {
	tdCommon
}

// MessageType return the string telegram-type of PassportElementErrorSourceTranslationFiles
func (passportElementErrorSourceTranslationFiles *PassportElementErrorSourceTranslationFiles) MessageType() string {
	return "passportElementErrorSourceTranslationFiles"
}

// NewPassportElementErrorSourceTranslationFiles creates a new PassportElementErrorSourceTranslationFiles
//
func NewPassportElementErrorSourceTranslationFiles() *PassportElementErrorSourceTranslationFiles {
	passportElementErrorSourceTranslationFilesTemp := PassportElementErrorSourceTranslationFiles{
		tdCommon: tdCommon{Type: "passportElementErrorSourceTranslationFiles"},
	}

	return &passportElementErrorSourceTranslationFilesTemp
}

// GetPassportElementErrorSourceEnum return the enum type of this object
func (passportElementErrorSourceTranslationFiles *PassportElementErrorSourceTranslationFiles) GetPassportElementErrorSourceEnum() PassportElementErrorSourceEnum {
	return PassportElementErrorSourceTranslationFilesType
}

// PassportElementErrorSourceFile The file contains an error. The error will be considered resolved when the file changes
type PassportElementErrorSourceFile struct {
	tdCommon
	FileIndex int32 `json:"file_index"` // Index of a file with the error
}

// MessageType return the string telegram-type of PassportElementErrorSourceFile
func (passportElementErrorSourceFile *PassportElementErrorSourceFile) MessageType() string {
	return "passportElementErrorSourceFile"
}

// NewPassportElementErrorSourceFile creates a new PassportElementErrorSourceFile
//
// @param fileIndex Index of a file with the error
func NewPassportElementErrorSourceFile(fileIndex int32) *PassportElementErrorSourceFile {
	passportElementErrorSourceFileTemp := PassportElementErrorSourceFile{
		tdCommon:  tdCommon{Type: "passportElementErrorSourceFile"},
		FileIndex: fileIndex,
	}

	return &passportElementErrorSourceFileTemp
}

// GetPassportElementErrorSourceEnum return the enum type of this object
func (passportElementErrorSourceFile *PassportElementErrorSourceFile) GetPassportElementErrorSourceEnum() PassportElementErrorSourceEnum {
	return PassportElementErrorSourceFileType
}

// PassportElementErrorSourceFiles The list of attached files contains an error. The error will be considered resolved when the list of files changes
type PassportElementErrorSourceFiles struct {
	tdCommon
}

// MessageType return the string telegram-type of PassportElementErrorSourceFiles
func (passportElementErrorSourceFiles *PassportElementErrorSourceFiles) MessageType() string {
	return "passportElementErrorSourceFiles"
}

// NewPassportElementErrorSourceFiles creates a new PassportElementErrorSourceFiles
//
func NewPassportElementErrorSourceFiles() *PassportElementErrorSourceFiles {
	passportElementErrorSourceFilesTemp := PassportElementErrorSourceFiles{
		tdCommon: tdCommon{Type: "passportElementErrorSourceFiles"},
	}

	return &passportElementErrorSourceFilesTemp
}

// GetPassportElementErrorSourceEnum return the enum type of this object
func (passportElementErrorSourceFiles *PassportElementErrorSourceFiles) GetPassportElementErrorSourceEnum() PassportElementErrorSourceEnum {
	return PassportElementErrorSourceFilesType
}
