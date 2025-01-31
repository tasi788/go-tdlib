package tdlib

import (
	"encoding/json"
	"fmt"
)

// InputFile Points to a file
type InputFile interface {
	GetInputFileEnum() InputFileEnum
}

// InputFileEnum Alias for abstract InputFile 'Sub-Classes', used as constant-enum here
type InputFileEnum string

// InputFile enums
const (
	InputFileIdType        InputFileEnum = "inputFileId"
	InputFileRemoteType    InputFileEnum = "inputFileRemote"
	InputFileLocalType     InputFileEnum = "inputFileLocal"
	InputFileGeneratedType InputFileEnum = "inputFileGenerated"
)

func unmarshalInputFile(rawMsg *json.RawMessage) (InputFile, error) {

	if rawMsg == nil {
		return nil, nil
	}
	var objMap map[string]interface{}
	err := json.Unmarshal(*rawMsg, &objMap)
	if err != nil {
		return nil, err
	}

	switch InputFileEnum(objMap["@type"].(string)) {
	case InputFileIdType:
		var inputFileId InputFileId
		err := json.Unmarshal(*rawMsg, &inputFileId)
		return &inputFileId, err

	case InputFileRemoteType:
		var inputFileRemote InputFileRemote
		err := json.Unmarshal(*rawMsg, &inputFileRemote)
		return &inputFileRemote, err

	case InputFileLocalType:
		var inputFileLocal InputFileLocal
		err := json.Unmarshal(*rawMsg, &inputFileLocal)
		return &inputFileLocal, err

	case InputFileGeneratedType:
		var inputFileGenerated InputFileGenerated
		err := json.Unmarshal(*rawMsg, &inputFileGenerated)
		return &inputFileGenerated, err

	default:
		return nil, fmt.Errorf("Error UnMarshaling, unknown type:" + objMap["@type"].(string))
	}
}

// InputFileId A file defined by its unique ID
type InputFileId struct {
	tdCommon
	Id int32 `json:"id"` // Unique file identifier
}

// MessageType return the string telegram-type of InputFileId
func (inputFileId *InputFileId) MessageType() string {
	return "inputFileId"
}

// NewInputFileId creates a new InputFileId
//
// @param id Unique file identifier
func NewInputFileId(id int32) *InputFileId {
	inputFileIdTemp := InputFileId{
		tdCommon: tdCommon{Type: "inputFileId"},
		Id:       id,
	}

	return &inputFileIdTemp
}

// GetInputFileEnum return the enum type of this object
func (inputFileId *InputFileId) GetInputFileEnum() InputFileEnum {
	return InputFileIdType
}

// InputFileRemote A file defined by its remote ID. The remote ID is guaranteed to be usable only if the corresponding file is still accessible to the user and known to TDLib. For example, if the file is from a message, then the message must be not deleted and accessible to the user. If the file database is disabled, then the corresponding object with the file must be preloaded by the application
type InputFileRemote struct {
	tdCommon
	Id string `json:"id"` // Remote file identifier
}

// MessageType return the string telegram-type of InputFileRemote
func (inputFileRemote *InputFileRemote) MessageType() string {
	return "inputFileRemote"
}

// NewInputFileRemote creates a new InputFileRemote
//
// @param id Remote file identifier
func NewInputFileRemote(id string) *InputFileRemote {
	inputFileRemoteTemp := InputFileRemote{
		tdCommon: tdCommon{Type: "inputFileRemote"},
		Id:       id,
	}

	return &inputFileRemoteTemp
}

// GetInputFileEnum return the enum type of this object
func (inputFileRemote *InputFileRemote) GetInputFileEnum() InputFileEnum {
	return InputFileRemoteType
}

// InputFileLocal A file defined by a local path
type InputFileLocal struct {
	tdCommon
	Path string `json:"path"` // Local path to the file
}

// MessageType return the string telegram-type of InputFileLocal
func (inputFileLocal *InputFileLocal) MessageType() string {
	return "inputFileLocal"
}

// NewInputFileLocal creates a new InputFileLocal
//
// @param path Local path to the file
func NewInputFileLocal(path string) *InputFileLocal {
	inputFileLocalTemp := InputFileLocal{
		tdCommon: tdCommon{Type: "inputFileLocal"},
		Path:     path,
	}

	return &inputFileLocalTemp
}

// GetInputFileEnum return the enum type of this object
func (inputFileLocal *InputFileLocal) GetInputFileEnum() InputFileEnum {
	return InputFileLocalType
}

// InputFileGenerated A file generated by the application
type InputFileGenerated struct {
	tdCommon
	OriginalPath string `json:"original_path"` // Local path to a file from which the file is generated; may be empty if there is no such file
	Conversion   string `json:"conversion"`    // String specifying the conversion applied to the original file; must be persistent across application restarts. Conversions beginning with '#' are reserved for internal TDLib usage
	ExpectedSize int32  `json:"expected_size"` // Expected size of the generated file, in bytes; 0 if unknown
}

// MessageType return the string telegram-type of InputFileGenerated
func (inputFileGenerated *InputFileGenerated) MessageType() string {
	return "inputFileGenerated"
}

// NewInputFileGenerated creates a new InputFileGenerated
//
// @param originalPath Local path to a file from which the file is generated; may be empty if there is no such file
// @param conversion String specifying the conversion applied to the original file; must be persistent across application restarts. Conversions beginning with '#' are reserved for internal TDLib usage
// @param expectedSize Expected size of the generated file, in bytes; 0 if unknown
func NewInputFileGenerated(originalPath string, conversion string, expectedSize int32) *InputFileGenerated {
	inputFileGeneratedTemp := InputFileGenerated{
		tdCommon:     tdCommon{Type: "inputFileGenerated"},
		OriginalPath: originalPath,
		Conversion:   conversion,
		ExpectedSize: expectedSize,
	}

	return &inputFileGeneratedTemp
}

// GetInputFileEnum return the enum type of this object
func (inputFileGenerated *InputFileGenerated) GetInputFileEnum() InputFileEnum {
	return InputFileGeneratedType
}
