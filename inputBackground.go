package tdlib

import (
	"encoding/json"
	"fmt"
)

// InputBackground Contains information about background to set
type InputBackground interface {
	GetInputBackgroundEnum() InputBackgroundEnum
}

// InputBackgroundEnum Alias for abstract InputBackground 'Sub-Classes', used as constant-enum here
type InputBackgroundEnum string

// InputBackground enums
const (
	InputBackgroundLocalType  InputBackgroundEnum = "inputBackgroundLocal"
	InputBackgroundRemoteType InputBackgroundEnum = "inputBackgroundRemote"
)

func unmarshalInputBackground(rawMsg *json.RawMessage) (InputBackground, error) {

	if rawMsg == nil {
		return nil, nil
	}
	var objMap map[string]interface{}
	err := json.Unmarshal(*rawMsg, &objMap)
	if err != nil {
		return nil, err
	}

	switch InputBackgroundEnum(objMap["@type"].(string)) {
	case InputBackgroundLocalType:
		var inputBackgroundLocal InputBackgroundLocal
		err := json.Unmarshal(*rawMsg, &inputBackgroundLocal)
		return &inputBackgroundLocal, err

	case InputBackgroundRemoteType:
		var inputBackgroundRemote InputBackgroundRemote
		err := json.Unmarshal(*rawMsg, &inputBackgroundRemote)
		return &inputBackgroundRemote, err

	default:
		return nil, fmt.Errorf("Error UnMarshaling, unknown type:" + objMap["@type"].(string))
	}
}

// InputBackgroundLocal A background from a local file
type InputBackgroundLocal struct {
	tdCommon
	Background InputFile `json:"background"` // Background file to use. Only inputFileLocal and inputFileGenerated are supported. The file must be in JPEG format for wallpapers and in PNG format for patterns
}

// MessageType return the string telegram-type of InputBackgroundLocal
func (inputBackgroundLocal *InputBackgroundLocal) MessageType() string {
	return "inputBackgroundLocal"
}

// NewInputBackgroundLocal creates a new InputBackgroundLocal
//
// @param background Background file to use. Only inputFileLocal and inputFileGenerated are supported. The file must be in JPEG format for wallpapers and in PNG format for patterns
func NewInputBackgroundLocal(background InputFile) *InputBackgroundLocal {
	inputBackgroundLocalTemp := InputBackgroundLocal{
		tdCommon:   tdCommon{Type: "inputBackgroundLocal"},
		Background: background,
	}

	return &inputBackgroundLocalTemp
}

// UnmarshalJSON unmarshal to json
func (inputBackgroundLocal *InputBackgroundLocal) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	tempObj := struct {
		tdCommon
	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	inputBackgroundLocal.tdCommon = tempObj.tdCommon

	fieldBackground, _ := unmarshalInputFile(objMap["background"])
	inputBackgroundLocal.Background = fieldBackground

	return nil
}

// GetInputBackgroundEnum return the enum type of this object
func (inputBackgroundLocal *InputBackgroundLocal) GetInputBackgroundEnum() InputBackgroundEnum {
	return InputBackgroundLocalType
}

// InputBackgroundRemote A background from the server
type InputBackgroundRemote struct {
	tdCommon
	BackgroundId JSONInt64 `json:"background_id"` // The background identifier
}

// MessageType return the string telegram-type of InputBackgroundRemote
func (inputBackgroundRemote *InputBackgroundRemote) MessageType() string {
	return "inputBackgroundRemote"
}

// NewInputBackgroundRemote creates a new InputBackgroundRemote
//
// @param backgroundId The background identifier
func NewInputBackgroundRemote(backgroundId JSONInt64) *InputBackgroundRemote {
	inputBackgroundRemoteTemp := InputBackgroundRemote{
		tdCommon:     tdCommon{Type: "inputBackgroundRemote"},
		BackgroundId: backgroundId,
	}

	return &inputBackgroundRemoteTemp
}

// GetInputBackgroundEnum return the enum type of this object
func (inputBackgroundRemote *InputBackgroundRemote) GetInputBackgroundEnum() InputBackgroundEnum {
	return InputBackgroundRemoteType
}
