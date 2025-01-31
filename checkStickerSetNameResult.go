package tdlib

import (
	"encoding/json"
	"fmt"
)

// CheckStickerSetNameResult Represents result of checking whether a name can be used for a new sticker set
type CheckStickerSetNameResult interface {
	GetCheckStickerSetNameResultEnum() CheckStickerSetNameResultEnum
}

// CheckStickerSetNameResultEnum Alias for abstract CheckStickerSetNameResult 'Sub-Classes', used as constant-enum here
type CheckStickerSetNameResultEnum string

// CheckStickerSetNameResult enums
const (
	CheckStickerSetNameResultOkType           CheckStickerSetNameResultEnum = "checkStickerSetNameResultOk"
	CheckStickerSetNameResultNameInvalidType  CheckStickerSetNameResultEnum = "checkStickerSetNameResultNameInvalid"
	CheckStickerSetNameResultNameOccupiedType CheckStickerSetNameResultEnum = "checkStickerSetNameResultNameOccupied"
)

func unmarshalCheckStickerSetNameResult(rawMsg *json.RawMessage) (CheckStickerSetNameResult, error) {

	if rawMsg == nil {
		return nil, nil
	}
	var objMap map[string]interface{}
	err := json.Unmarshal(*rawMsg, &objMap)
	if err != nil {
		return nil, err
	}

	switch CheckStickerSetNameResultEnum(objMap["@type"].(string)) {
	case CheckStickerSetNameResultOkType:
		var checkStickerSetNameResultOk CheckStickerSetNameResultOk
		err := json.Unmarshal(*rawMsg, &checkStickerSetNameResultOk)
		return &checkStickerSetNameResultOk, err

	case CheckStickerSetNameResultNameInvalidType:
		var checkStickerSetNameResultNameInvalid CheckStickerSetNameResultNameInvalid
		err := json.Unmarshal(*rawMsg, &checkStickerSetNameResultNameInvalid)
		return &checkStickerSetNameResultNameInvalid, err

	case CheckStickerSetNameResultNameOccupiedType:
		var checkStickerSetNameResultNameOccupied CheckStickerSetNameResultNameOccupied
		err := json.Unmarshal(*rawMsg, &checkStickerSetNameResultNameOccupied)
		return &checkStickerSetNameResultNameOccupied, err

	default:
		return nil, fmt.Errorf("Error UnMarshaling, unknown type:" + objMap["@type"].(string))
	}
}

// CheckStickerSetNameResultOk The name can be set
type CheckStickerSetNameResultOk struct {
	tdCommon
}

// MessageType return the string telegram-type of CheckStickerSetNameResultOk
func (checkStickerSetNameResultOk *CheckStickerSetNameResultOk) MessageType() string {
	return "checkStickerSetNameResultOk"
}

// NewCheckStickerSetNameResultOk creates a new CheckStickerSetNameResultOk
//
func NewCheckStickerSetNameResultOk() *CheckStickerSetNameResultOk {
	checkStickerSetNameResultOkTemp := CheckStickerSetNameResultOk{
		tdCommon: tdCommon{Type: "checkStickerSetNameResultOk"},
	}

	return &checkStickerSetNameResultOkTemp
}

// GetCheckStickerSetNameResultEnum return the enum type of this object
func (checkStickerSetNameResultOk *CheckStickerSetNameResultOk) GetCheckStickerSetNameResultEnum() CheckStickerSetNameResultEnum {
	return CheckStickerSetNameResultOkType
}

// CheckStickerSetNameResultNameInvalid The name is invalid
type CheckStickerSetNameResultNameInvalid struct {
	tdCommon
}

// MessageType return the string telegram-type of CheckStickerSetNameResultNameInvalid
func (checkStickerSetNameResultNameInvalid *CheckStickerSetNameResultNameInvalid) MessageType() string {
	return "checkStickerSetNameResultNameInvalid"
}

// NewCheckStickerSetNameResultNameInvalid creates a new CheckStickerSetNameResultNameInvalid
//
func NewCheckStickerSetNameResultNameInvalid() *CheckStickerSetNameResultNameInvalid {
	checkStickerSetNameResultNameInvalidTemp := CheckStickerSetNameResultNameInvalid{
		tdCommon: tdCommon{Type: "checkStickerSetNameResultNameInvalid"},
	}

	return &checkStickerSetNameResultNameInvalidTemp
}

// GetCheckStickerSetNameResultEnum return the enum type of this object
func (checkStickerSetNameResultNameInvalid *CheckStickerSetNameResultNameInvalid) GetCheckStickerSetNameResultEnum() CheckStickerSetNameResultEnum {
	return CheckStickerSetNameResultNameInvalidType
}

// CheckStickerSetNameResultNameOccupied The name is occupied
type CheckStickerSetNameResultNameOccupied struct {
	tdCommon
}

// MessageType return the string telegram-type of CheckStickerSetNameResultNameOccupied
func (checkStickerSetNameResultNameOccupied *CheckStickerSetNameResultNameOccupied) MessageType() string {
	return "checkStickerSetNameResultNameOccupied"
}

// NewCheckStickerSetNameResultNameOccupied creates a new CheckStickerSetNameResultNameOccupied
//
func NewCheckStickerSetNameResultNameOccupied() *CheckStickerSetNameResultNameOccupied {
	checkStickerSetNameResultNameOccupiedTemp := CheckStickerSetNameResultNameOccupied{
		tdCommon: tdCommon{Type: "checkStickerSetNameResultNameOccupied"},
	}

	return &checkStickerSetNameResultNameOccupiedTemp
}

// GetCheckStickerSetNameResultEnum return the enum type of this object
func (checkStickerSetNameResultNameOccupied *CheckStickerSetNameResultNameOccupied) GetCheckStickerSetNameResultEnum() CheckStickerSetNameResultEnum {
	return CheckStickerSetNameResultNameOccupiedType
}

// CheckStickerSetName Checks whether a name can be used for a new sticker set
// @param name Name to be checked
func (client *Client) CheckStickerSetName(name string) (CheckStickerSetNameResult, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "checkStickerSetName",
		"name":  name,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	switch CheckStickerSetNameResultEnum(result.Data["@type"].(string)) {

	case CheckStickerSetNameResultOkType:
		var checkStickerSetNameResult CheckStickerSetNameResultOk
		err = json.Unmarshal(result.Raw, &checkStickerSetNameResult)
		return &checkStickerSetNameResult, err

	case CheckStickerSetNameResultNameInvalidType:
		var checkStickerSetNameResult CheckStickerSetNameResultNameInvalid
		err = json.Unmarshal(result.Raw, &checkStickerSetNameResult)
		return &checkStickerSetNameResult, err

	case CheckStickerSetNameResultNameOccupiedType:
		var checkStickerSetNameResult CheckStickerSetNameResultNameOccupied
		err = json.Unmarshal(result.Raw, &checkStickerSetNameResult)
		return &checkStickerSetNameResult, err

	default:
		return nil, fmt.Errorf("Invalid type")
	}
}
