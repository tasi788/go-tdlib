package tdlib

import (
	"encoding/json"
	"fmt"
)

// BackgroundFill Describes a fill of a background
type BackgroundFill interface {
	GetBackgroundFillEnum() BackgroundFillEnum
}

// BackgroundFillEnum Alias for abstract BackgroundFill 'Sub-Classes', used as constant-enum here
type BackgroundFillEnum string

// BackgroundFill enums
const (
	BackgroundFillSolidType            BackgroundFillEnum = "backgroundFillSolid"
	BackgroundFillGradientType         BackgroundFillEnum = "backgroundFillGradient"
	BackgroundFillFreeformGradientType BackgroundFillEnum = "backgroundFillFreeformGradient"
)

func unmarshalBackgroundFill(rawMsg *json.RawMessage) (BackgroundFill, error) {

	if rawMsg == nil {
		return nil, nil
	}
	var objMap map[string]interface{}
	err := json.Unmarshal(*rawMsg, &objMap)
	if err != nil {
		return nil, err
	}

	switch BackgroundFillEnum(objMap["@type"].(string)) {
	case BackgroundFillSolidType:
		var backgroundFillSolid BackgroundFillSolid
		err := json.Unmarshal(*rawMsg, &backgroundFillSolid)
		return &backgroundFillSolid, err

	case BackgroundFillGradientType:
		var backgroundFillGradient BackgroundFillGradient
		err := json.Unmarshal(*rawMsg, &backgroundFillGradient)
		return &backgroundFillGradient, err

	case BackgroundFillFreeformGradientType:
		var backgroundFillFreeformGradient BackgroundFillFreeformGradient
		err := json.Unmarshal(*rawMsg, &backgroundFillFreeformGradient)
		return &backgroundFillFreeformGradient, err

	default:
		return nil, fmt.Errorf("Error UnMarshaling, unknown type:" + objMap["@type"].(string))
	}
}

// BackgroundFillSolid Describes a solid fill of a background
type BackgroundFillSolid struct {
	tdCommon
	Color int32 `json:"color"` // A color of the background in the RGB24 format
}

// MessageType return the string telegram-type of BackgroundFillSolid
func (backgroundFillSolid *BackgroundFillSolid) MessageType() string {
	return "backgroundFillSolid"
}

// NewBackgroundFillSolid creates a new BackgroundFillSolid
//
// @param color A color of the background in the RGB24 format
func NewBackgroundFillSolid(color int32) *BackgroundFillSolid {
	backgroundFillSolidTemp := BackgroundFillSolid{
		tdCommon: tdCommon{Type: "backgroundFillSolid"},
		Color:    color,
	}

	return &backgroundFillSolidTemp
}

// GetBackgroundFillEnum return the enum type of this object
func (backgroundFillSolid *BackgroundFillSolid) GetBackgroundFillEnum() BackgroundFillEnum {
	return BackgroundFillSolidType
}

// BackgroundFillGradient Describes a gradient fill of a background
type BackgroundFillGradient struct {
	tdCommon
	TopColor      int32 `json:"top_color"`      // A top color of the background in the RGB24 format
	BottomColor   int32 `json:"bottom_color"`   // A bottom color of the background in the RGB24 format
	RotationAngle int32 `json:"rotation_angle"` // Clockwise rotation angle of the gradient, in degrees; 0-359. Must be always divisible by 45
}

// MessageType return the string telegram-type of BackgroundFillGradient
func (backgroundFillGradient *BackgroundFillGradient) MessageType() string {
	return "backgroundFillGradient"
}

// NewBackgroundFillGradient creates a new BackgroundFillGradient
//
// @param topColor A top color of the background in the RGB24 format
// @param bottomColor A bottom color of the background in the RGB24 format
// @param rotationAngle Clockwise rotation angle of the gradient, in degrees; 0-359. Must be always divisible by 45
func NewBackgroundFillGradient(topColor int32, bottomColor int32, rotationAngle int32) *BackgroundFillGradient {
	backgroundFillGradientTemp := BackgroundFillGradient{
		tdCommon:      tdCommon{Type: "backgroundFillGradient"},
		TopColor:      topColor,
		BottomColor:   bottomColor,
		RotationAngle: rotationAngle,
	}

	return &backgroundFillGradientTemp
}

// GetBackgroundFillEnum return the enum type of this object
func (backgroundFillGradient *BackgroundFillGradient) GetBackgroundFillEnum() BackgroundFillEnum {
	return BackgroundFillGradientType
}

// BackgroundFillFreeformGradient Describes a freeform gradient fill of a background
type BackgroundFillFreeformGradient struct {
	tdCommon
	Colors []int32 `json:"colors"` // A list of 3 or 4 colors of the freeform gradients in the RGB24 format
}

// MessageType return the string telegram-type of BackgroundFillFreeformGradient
func (backgroundFillFreeformGradient *BackgroundFillFreeformGradient) MessageType() string {
	return "backgroundFillFreeformGradient"
}

// NewBackgroundFillFreeformGradient creates a new BackgroundFillFreeformGradient
//
// @param colors A list of 3 or 4 colors of the freeform gradients in the RGB24 format
func NewBackgroundFillFreeformGradient(colors []int32) *BackgroundFillFreeformGradient {
	backgroundFillFreeformGradientTemp := BackgroundFillFreeformGradient{
		tdCommon: tdCommon{Type: "backgroundFillFreeformGradient"},
		Colors:   colors,
	}

	return &backgroundFillFreeformGradientTemp
}

// GetBackgroundFillEnum return the enum type of this object
func (backgroundFillFreeformGradient *BackgroundFillFreeformGradient) GetBackgroundFillEnum() BackgroundFillEnum {
	return BackgroundFillFreeformGradientType
}
