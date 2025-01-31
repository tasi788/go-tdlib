package tdlib

import (
	"encoding/json"
	"fmt"
)

// InputPassportElementErrorSource Contains the description of an error in a Telegram Passport element; for bots only
type InputPassportElementErrorSource interface {
	GetInputPassportElementErrorSourceEnum() InputPassportElementErrorSourceEnum
}

// InputPassportElementErrorSourceEnum Alias for abstract InputPassportElementErrorSource 'Sub-Classes', used as constant-enum here
type InputPassportElementErrorSourceEnum string

// InputPassportElementErrorSource enums
const (
	InputPassportElementErrorSourceUnspecifiedType      InputPassportElementErrorSourceEnum = "inputPassportElementErrorSourceUnspecified"
	InputPassportElementErrorSourceDataFieldType        InputPassportElementErrorSourceEnum = "inputPassportElementErrorSourceDataField"
	InputPassportElementErrorSourceFrontSideType        InputPassportElementErrorSourceEnum = "inputPassportElementErrorSourceFrontSide"
	InputPassportElementErrorSourceReverseSideType      InputPassportElementErrorSourceEnum = "inputPassportElementErrorSourceReverseSide"
	InputPassportElementErrorSourceSelfieType           InputPassportElementErrorSourceEnum = "inputPassportElementErrorSourceSelfie"
	InputPassportElementErrorSourceTranslationFileType  InputPassportElementErrorSourceEnum = "inputPassportElementErrorSourceTranslationFile"
	InputPassportElementErrorSourceTranslationFilesType InputPassportElementErrorSourceEnum = "inputPassportElementErrorSourceTranslationFiles"
	InputPassportElementErrorSourceFileType             InputPassportElementErrorSourceEnum = "inputPassportElementErrorSourceFile"
	InputPassportElementErrorSourceFilesType            InputPassportElementErrorSourceEnum = "inputPassportElementErrorSourceFiles"
)

func unmarshalInputPassportElementErrorSource(rawMsg *json.RawMessage) (InputPassportElementErrorSource, error) {

	if rawMsg == nil {
		return nil, nil
	}
	var objMap map[string]interface{}
	err := json.Unmarshal(*rawMsg, &objMap)
	if err != nil {
		return nil, err
	}

	switch InputPassportElementErrorSourceEnum(objMap["@type"].(string)) {
	case InputPassportElementErrorSourceUnspecifiedType:
		var inputPassportElementErrorSourceUnspecified InputPassportElementErrorSourceUnspecified
		err := json.Unmarshal(*rawMsg, &inputPassportElementErrorSourceUnspecified)
		return &inputPassportElementErrorSourceUnspecified, err

	case InputPassportElementErrorSourceDataFieldType:
		var inputPassportElementErrorSourceDataField InputPassportElementErrorSourceDataField
		err := json.Unmarshal(*rawMsg, &inputPassportElementErrorSourceDataField)
		return &inputPassportElementErrorSourceDataField, err

	case InputPassportElementErrorSourceFrontSideType:
		var inputPassportElementErrorSourceFrontSide InputPassportElementErrorSourceFrontSide
		err := json.Unmarshal(*rawMsg, &inputPassportElementErrorSourceFrontSide)
		return &inputPassportElementErrorSourceFrontSide, err

	case InputPassportElementErrorSourceReverseSideType:
		var inputPassportElementErrorSourceReverseSide InputPassportElementErrorSourceReverseSide
		err := json.Unmarshal(*rawMsg, &inputPassportElementErrorSourceReverseSide)
		return &inputPassportElementErrorSourceReverseSide, err

	case InputPassportElementErrorSourceSelfieType:
		var inputPassportElementErrorSourceSelfie InputPassportElementErrorSourceSelfie
		err := json.Unmarshal(*rawMsg, &inputPassportElementErrorSourceSelfie)
		return &inputPassportElementErrorSourceSelfie, err

	case InputPassportElementErrorSourceTranslationFileType:
		var inputPassportElementErrorSourceTranslationFile InputPassportElementErrorSourceTranslationFile
		err := json.Unmarshal(*rawMsg, &inputPassportElementErrorSourceTranslationFile)
		return &inputPassportElementErrorSourceTranslationFile, err

	case InputPassportElementErrorSourceTranslationFilesType:
		var inputPassportElementErrorSourceTranslationFiles InputPassportElementErrorSourceTranslationFiles
		err := json.Unmarshal(*rawMsg, &inputPassportElementErrorSourceTranslationFiles)
		return &inputPassportElementErrorSourceTranslationFiles, err

	case InputPassportElementErrorSourceFileType:
		var inputPassportElementErrorSourceFile InputPassportElementErrorSourceFile
		err := json.Unmarshal(*rawMsg, &inputPassportElementErrorSourceFile)
		return &inputPassportElementErrorSourceFile, err

	case InputPassportElementErrorSourceFilesType:
		var inputPassportElementErrorSourceFiles InputPassportElementErrorSourceFiles
		err := json.Unmarshal(*rawMsg, &inputPassportElementErrorSourceFiles)
		return &inputPassportElementErrorSourceFiles, err

	default:
		return nil, fmt.Errorf("Error UnMarshaling, unknown type:" + objMap["@type"].(string))
	}
}

// InputPassportElementErrorSourceUnspecified The element contains an error in an unspecified place. The error will be considered resolved when new data is added
type InputPassportElementErrorSourceUnspecified struct {
	tdCommon
	ElementHash []byte `json:"element_hash"` // Current hash of the entire element
}

// MessageType return the string telegram-type of InputPassportElementErrorSourceUnspecified
func (inputPassportElementErrorSourceUnspecified *InputPassportElementErrorSourceUnspecified) MessageType() string {
	return "inputPassportElementErrorSourceUnspecified"
}

// NewInputPassportElementErrorSourceUnspecified creates a new InputPassportElementErrorSourceUnspecified
//
// @param elementHash Current hash of the entire element
func NewInputPassportElementErrorSourceUnspecified(elementHash []byte) *InputPassportElementErrorSourceUnspecified {
	inputPassportElementErrorSourceUnspecifiedTemp := InputPassportElementErrorSourceUnspecified{
		tdCommon:    tdCommon{Type: "inputPassportElementErrorSourceUnspecified"},
		ElementHash: elementHash,
	}

	return &inputPassportElementErrorSourceUnspecifiedTemp
}

// GetInputPassportElementErrorSourceEnum return the enum type of this object
func (inputPassportElementErrorSourceUnspecified *InputPassportElementErrorSourceUnspecified) GetInputPassportElementErrorSourceEnum() InputPassportElementErrorSourceEnum {
	return InputPassportElementErrorSourceUnspecifiedType
}

// InputPassportElementErrorSourceDataField A data field contains an error. The error is considered resolved when the field's value changes
type InputPassportElementErrorSourceDataField struct {
	tdCommon
	FieldName string `json:"field_name"` // Field name
	DataHash  []byte `json:"data_hash"`  // Current data hash
}

// MessageType return the string telegram-type of InputPassportElementErrorSourceDataField
func (inputPassportElementErrorSourceDataField *InputPassportElementErrorSourceDataField) MessageType() string {
	return "inputPassportElementErrorSourceDataField"
}

// NewInputPassportElementErrorSourceDataField creates a new InputPassportElementErrorSourceDataField
//
// @param fieldName Field name
// @param dataHash Current data hash
func NewInputPassportElementErrorSourceDataField(fieldName string, dataHash []byte) *InputPassportElementErrorSourceDataField {
	inputPassportElementErrorSourceDataFieldTemp := InputPassportElementErrorSourceDataField{
		tdCommon:  tdCommon{Type: "inputPassportElementErrorSourceDataField"},
		FieldName: fieldName,
		DataHash:  dataHash,
	}

	return &inputPassportElementErrorSourceDataFieldTemp
}

// GetInputPassportElementErrorSourceEnum return the enum type of this object
func (inputPassportElementErrorSourceDataField *InputPassportElementErrorSourceDataField) GetInputPassportElementErrorSourceEnum() InputPassportElementErrorSourceEnum {
	return InputPassportElementErrorSourceDataFieldType
}

// InputPassportElementErrorSourceFrontSide The front side of the document contains an error. The error is considered resolved when the file with the front side of the document changes
type InputPassportElementErrorSourceFrontSide struct {
	tdCommon
	FileHash []byte `json:"file_hash"` // Current hash of the file containing the front side
}

// MessageType return the string telegram-type of InputPassportElementErrorSourceFrontSide
func (inputPassportElementErrorSourceFrontSide *InputPassportElementErrorSourceFrontSide) MessageType() string {
	return "inputPassportElementErrorSourceFrontSide"
}

// NewInputPassportElementErrorSourceFrontSide creates a new InputPassportElementErrorSourceFrontSide
//
// @param fileHash Current hash of the file containing the front side
func NewInputPassportElementErrorSourceFrontSide(fileHash []byte) *InputPassportElementErrorSourceFrontSide {
	inputPassportElementErrorSourceFrontSideTemp := InputPassportElementErrorSourceFrontSide{
		tdCommon: tdCommon{Type: "inputPassportElementErrorSourceFrontSide"},
		FileHash: fileHash,
	}

	return &inputPassportElementErrorSourceFrontSideTemp
}

// GetInputPassportElementErrorSourceEnum return the enum type of this object
func (inputPassportElementErrorSourceFrontSide *InputPassportElementErrorSourceFrontSide) GetInputPassportElementErrorSourceEnum() InputPassportElementErrorSourceEnum {
	return InputPassportElementErrorSourceFrontSideType
}

// InputPassportElementErrorSourceReverseSide The reverse side of the document contains an error. The error is considered resolved when the file with the reverse side of the document changes
type InputPassportElementErrorSourceReverseSide struct {
	tdCommon
	FileHash []byte `json:"file_hash"` // Current hash of the file containing the reverse side
}

// MessageType return the string telegram-type of InputPassportElementErrorSourceReverseSide
func (inputPassportElementErrorSourceReverseSide *InputPassportElementErrorSourceReverseSide) MessageType() string {
	return "inputPassportElementErrorSourceReverseSide"
}

// NewInputPassportElementErrorSourceReverseSide creates a new InputPassportElementErrorSourceReverseSide
//
// @param fileHash Current hash of the file containing the reverse side
func NewInputPassportElementErrorSourceReverseSide(fileHash []byte) *InputPassportElementErrorSourceReverseSide {
	inputPassportElementErrorSourceReverseSideTemp := InputPassportElementErrorSourceReverseSide{
		tdCommon: tdCommon{Type: "inputPassportElementErrorSourceReverseSide"},
		FileHash: fileHash,
	}

	return &inputPassportElementErrorSourceReverseSideTemp
}

// GetInputPassportElementErrorSourceEnum return the enum type of this object
func (inputPassportElementErrorSourceReverseSide *InputPassportElementErrorSourceReverseSide) GetInputPassportElementErrorSourceEnum() InputPassportElementErrorSourceEnum {
	return InputPassportElementErrorSourceReverseSideType
}

// InputPassportElementErrorSourceSelfie The selfie contains an error. The error is considered resolved when the file with the selfie changes
type InputPassportElementErrorSourceSelfie struct {
	tdCommon
	FileHash []byte `json:"file_hash"` // Current hash of the file containing the selfie
}

// MessageType return the string telegram-type of InputPassportElementErrorSourceSelfie
func (inputPassportElementErrorSourceSelfie *InputPassportElementErrorSourceSelfie) MessageType() string {
	return "inputPassportElementErrorSourceSelfie"
}

// NewInputPassportElementErrorSourceSelfie creates a new InputPassportElementErrorSourceSelfie
//
// @param fileHash Current hash of the file containing the selfie
func NewInputPassportElementErrorSourceSelfie(fileHash []byte) *InputPassportElementErrorSourceSelfie {
	inputPassportElementErrorSourceSelfieTemp := InputPassportElementErrorSourceSelfie{
		tdCommon: tdCommon{Type: "inputPassportElementErrorSourceSelfie"},
		FileHash: fileHash,
	}

	return &inputPassportElementErrorSourceSelfieTemp
}

// GetInputPassportElementErrorSourceEnum return the enum type of this object
func (inputPassportElementErrorSourceSelfie *InputPassportElementErrorSourceSelfie) GetInputPassportElementErrorSourceEnum() InputPassportElementErrorSourceEnum {
	return InputPassportElementErrorSourceSelfieType
}

// InputPassportElementErrorSourceTranslationFile One of the files containing the translation of the document contains an error. The error is considered resolved when the file with the translation changes
type InputPassportElementErrorSourceTranslationFile struct {
	tdCommon
	FileHash []byte `json:"file_hash"` // Current hash of the file containing the translation
}

// MessageType return the string telegram-type of InputPassportElementErrorSourceTranslationFile
func (inputPassportElementErrorSourceTranslationFile *InputPassportElementErrorSourceTranslationFile) MessageType() string {
	return "inputPassportElementErrorSourceTranslationFile"
}

// NewInputPassportElementErrorSourceTranslationFile creates a new InputPassportElementErrorSourceTranslationFile
//
// @param fileHash Current hash of the file containing the translation
func NewInputPassportElementErrorSourceTranslationFile(fileHash []byte) *InputPassportElementErrorSourceTranslationFile {
	inputPassportElementErrorSourceTranslationFileTemp := InputPassportElementErrorSourceTranslationFile{
		tdCommon: tdCommon{Type: "inputPassportElementErrorSourceTranslationFile"},
		FileHash: fileHash,
	}

	return &inputPassportElementErrorSourceTranslationFileTemp
}

// GetInputPassportElementErrorSourceEnum return the enum type of this object
func (inputPassportElementErrorSourceTranslationFile *InputPassportElementErrorSourceTranslationFile) GetInputPassportElementErrorSourceEnum() InputPassportElementErrorSourceEnum {
	return InputPassportElementErrorSourceTranslationFileType
}

// InputPassportElementErrorSourceTranslationFiles The translation of the document contains an error. The error is considered resolved when the list of files changes
type InputPassportElementErrorSourceTranslationFiles struct {
	tdCommon
	FileHashes [][]byte `json:"file_hashes"` // Current hashes of all files with the translation
}

// MessageType return the string telegram-type of InputPassportElementErrorSourceTranslationFiles
func (inputPassportElementErrorSourceTranslationFiles *InputPassportElementErrorSourceTranslationFiles) MessageType() string {
	return "inputPassportElementErrorSourceTranslationFiles"
}

// NewInputPassportElementErrorSourceTranslationFiles creates a new InputPassportElementErrorSourceTranslationFiles
//
// @param fileHashes Current hashes of all files with the translation
func NewInputPassportElementErrorSourceTranslationFiles(fileHashes [][]byte) *InputPassportElementErrorSourceTranslationFiles {
	inputPassportElementErrorSourceTranslationFilesTemp := InputPassportElementErrorSourceTranslationFiles{
		tdCommon:   tdCommon{Type: "inputPassportElementErrorSourceTranslationFiles"},
		FileHashes: fileHashes,
	}

	return &inputPassportElementErrorSourceTranslationFilesTemp
}

// GetInputPassportElementErrorSourceEnum return the enum type of this object
func (inputPassportElementErrorSourceTranslationFiles *InputPassportElementErrorSourceTranslationFiles) GetInputPassportElementErrorSourceEnum() InputPassportElementErrorSourceEnum {
	return InputPassportElementErrorSourceTranslationFilesType
}

// InputPassportElementErrorSourceFile The file contains an error. The error is considered resolved when the file changes
type InputPassportElementErrorSourceFile struct {
	tdCommon
	FileHash []byte `json:"file_hash"` // Current hash of the file which has the error
}

// MessageType return the string telegram-type of InputPassportElementErrorSourceFile
func (inputPassportElementErrorSourceFile *InputPassportElementErrorSourceFile) MessageType() string {
	return "inputPassportElementErrorSourceFile"
}

// NewInputPassportElementErrorSourceFile creates a new InputPassportElementErrorSourceFile
//
// @param fileHash Current hash of the file which has the error
func NewInputPassportElementErrorSourceFile(fileHash []byte) *InputPassportElementErrorSourceFile {
	inputPassportElementErrorSourceFileTemp := InputPassportElementErrorSourceFile{
		tdCommon: tdCommon{Type: "inputPassportElementErrorSourceFile"},
		FileHash: fileHash,
	}

	return &inputPassportElementErrorSourceFileTemp
}

// GetInputPassportElementErrorSourceEnum return the enum type of this object
func (inputPassportElementErrorSourceFile *InputPassportElementErrorSourceFile) GetInputPassportElementErrorSourceEnum() InputPassportElementErrorSourceEnum {
	return InputPassportElementErrorSourceFileType
}

// InputPassportElementErrorSourceFiles The list of attached files contains an error. The error is considered resolved when the file list changes
type InputPassportElementErrorSourceFiles struct {
	tdCommon
	FileHashes [][]byte `json:"file_hashes"` // Current hashes of all attached files
}

// MessageType return the string telegram-type of InputPassportElementErrorSourceFiles
func (inputPassportElementErrorSourceFiles *InputPassportElementErrorSourceFiles) MessageType() string {
	return "inputPassportElementErrorSourceFiles"
}

// NewInputPassportElementErrorSourceFiles creates a new InputPassportElementErrorSourceFiles
//
// @param fileHashes Current hashes of all attached files
func NewInputPassportElementErrorSourceFiles(fileHashes [][]byte) *InputPassportElementErrorSourceFiles {
	inputPassportElementErrorSourceFilesTemp := InputPassportElementErrorSourceFiles{
		tdCommon:   tdCommon{Type: "inputPassportElementErrorSourceFiles"},
		FileHashes: fileHashes,
	}

	return &inputPassportElementErrorSourceFilesTemp
}

// GetInputPassportElementErrorSourceEnum return the enum type of this object
func (inputPassportElementErrorSourceFiles *InputPassportElementErrorSourceFiles) GetInputPassportElementErrorSourceEnum() InputPassportElementErrorSourceEnum {
	return InputPassportElementErrorSourceFilesType
}
