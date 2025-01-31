package tdlib

import (
	"encoding/json"
	"fmt"
)

// MaskPoint Part of the face, relative to which a mask is placed
type MaskPoint interface {
	GetMaskPointEnum() MaskPointEnum
}

// MaskPointEnum Alias for abstract MaskPoint 'Sub-Classes', used as constant-enum here
type MaskPointEnum string

// MaskPoint enums
const (
	MaskPointForeheadType MaskPointEnum = "maskPointForehead"
	MaskPointEyesType     MaskPointEnum = "maskPointEyes"
	MaskPointMouthType    MaskPointEnum = "maskPointMouth"
	MaskPointChinType     MaskPointEnum = "maskPointChin"
)

func unmarshalMaskPoint(rawMsg *json.RawMessage) (MaskPoint, error) {

	if rawMsg == nil {
		return nil, nil
	}
	var objMap map[string]interface{}
	err := json.Unmarshal(*rawMsg, &objMap)
	if err != nil {
		return nil, err
	}

	switch MaskPointEnum(objMap["@type"].(string)) {
	case MaskPointForeheadType:
		var maskPointForehead MaskPointForehead
		err := json.Unmarshal(*rawMsg, &maskPointForehead)
		return &maskPointForehead, err

	case MaskPointEyesType:
		var maskPointEyes MaskPointEyes
		err := json.Unmarshal(*rawMsg, &maskPointEyes)
		return &maskPointEyes, err

	case MaskPointMouthType:
		var maskPointMouth MaskPointMouth
		err := json.Unmarshal(*rawMsg, &maskPointMouth)
		return &maskPointMouth, err

	case MaskPointChinType:
		var maskPointChin MaskPointChin
		err := json.Unmarshal(*rawMsg, &maskPointChin)
		return &maskPointChin, err

	default:
		return nil, fmt.Errorf("Error UnMarshaling, unknown type:" + objMap["@type"].(string))
	}
}

// MaskPointForehead The mask is placed relatively to the forehead
type MaskPointForehead struct {
	tdCommon
}

// MessageType return the string telegram-type of MaskPointForehead
func (maskPointForehead *MaskPointForehead) MessageType() string {
	return "maskPointForehead"
}

// NewMaskPointForehead creates a new MaskPointForehead
//
func NewMaskPointForehead() *MaskPointForehead {
	maskPointForeheadTemp := MaskPointForehead{
		tdCommon: tdCommon{Type: "maskPointForehead"},
	}

	return &maskPointForeheadTemp
}

// GetMaskPointEnum return the enum type of this object
func (maskPointForehead *MaskPointForehead) GetMaskPointEnum() MaskPointEnum {
	return MaskPointForeheadType
}

// MaskPointEyes The mask is placed relatively to the eyes
type MaskPointEyes struct {
	tdCommon
}

// MessageType return the string telegram-type of MaskPointEyes
func (maskPointEyes *MaskPointEyes) MessageType() string {
	return "maskPointEyes"
}

// NewMaskPointEyes creates a new MaskPointEyes
//
func NewMaskPointEyes() *MaskPointEyes {
	maskPointEyesTemp := MaskPointEyes{
		tdCommon: tdCommon{Type: "maskPointEyes"},
	}

	return &maskPointEyesTemp
}

// GetMaskPointEnum return the enum type of this object
func (maskPointEyes *MaskPointEyes) GetMaskPointEnum() MaskPointEnum {
	return MaskPointEyesType
}

// MaskPointMouth The mask is placed relatively to the mouth
type MaskPointMouth struct {
	tdCommon
}

// MessageType return the string telegram-type of MaskPointMouth
func (maskPointMouth *MaskPointMouth) MessageType() string {
	return "maskPointMouth"
}

// NewMaskPointMouth creates a new MaskPointMouth
//
func NewMaskPointMouth() *MaskPointMouth {
	maskPointMouthTemp := MaskPointMouth{
		tdCommon: tdCommon{Type: "maskPointMouth"},
	}

	return &maskPointMouthTemp
}

// GetMaskPointEnum return the enum type of this object
func (maskPointMouth *MaskPointMouth) GetMaskPointEnum() MaskPointEnum {
	return MaskPointMouthType
}

// MaskPointChin The mask is placed relatively to the chin
type MaskPointChin struct {
	tdCommon
}

// MessageType return the string telegram-type of MaskPointChin
func (maskPointChin *MaskPointChin) MessageType() string {
	return "maskPointChin"
}

// NewMaskPointChin creates a new MaskPointChin
//
func NewMaskPointChin() *MaskPointChin {
	maskPointChinTemp := MaskPointChin{
		tdCommon: tdCommon{Type: "maskPointChin"},
	}

	return &maskPointChinTemp
}

// GetMaskPointEnum return the enum type of this object
func (maskPointChin *MaskPointChin) GetMaskPointEnum() MaskPointEnum {
	return MaskPointChinType
}
