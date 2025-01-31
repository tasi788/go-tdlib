package tdlib

import (
	"encoding/json"
	"fmt"
)

// VectorPathCommand Represents a vector path command
type VectorPathCommand interface {
	GetVectorPathCommandEnum() VectorPathCommandEnum
}

// VectorPathCommandEnum Alias for abstract VectorPathCommand 'Sub-Classes', used as constant-enum here
type VectorPathCommandEnum string

// VectorPathCommand enums
const (
	VectorPathCommandLineType             VectorPathCommandEnum = "vectorPathCommandLine"
	VectorPathCommandCubicBezierCurveType VectorPathCommandEnum = "vectorPathCommandCubicBezierCurve"
)

func unmarshalVectorPathCommand(rawMsg *json.RawMessage) (VectorPathCommand, error) {

	if rawMsg == nil {
		return nil, nil
	}
	var objMap map[string]interface{}
	err := json.Unmarshal(*rawMsg, &objMap)
	if err != nil {
		return nil, err
	}

	switch VectorPathCommandEnum(objMap["@type"].(string)) {
	case VectorPathCommandLineType:
		var vectorPathCommandLine VectorPathCommandLine
		err := json.Unmarshal(*rawMsg, &vectorPathCommandLine)
		return &vectorPathCommandLine, err

	case VectorPathCommandCubicBezierCurveType:
		var vectorPathCommandCubicBezierCurve VectorPathCommandCubicBezierCurve
		err := json.Unmarshal(*rawMsg, &vectorPathCommandCubicBezierCurve)
		return &vectorPathCommandCubicBezierCurve, err

	default:
		return nil, fmt.Errorf("Error UnMarshaling, unknown type:" + objMap["@type"].(string))
	}
}

// VectorPathCommandLine A straight line to a given point
type VectorPathCommandLine struct {
	tdCommon
	EndPoint *Point `json:"end_point"` // The end point of the straight line
}

// MessageType return the string telegram-type of VectorPathCommandLine
func (vectorPathCommandLine *VectorPathCommandLine) MessageType() string {
	return "vectorPathCommandLine"
}

// NewVectorPathCommandLine creates a new VectorPathCommandLine
//
// @param endPoint The end point of the straight line
func NewVectorPathCommandLine(endPoint *Point) *VectorPathCommandLine {
	vectorPathCommandLineTemp := VectorPathCommandLine{
		tdCommon: tdCommon{Type: "vectorPathCommandLine"},
		EndPoint: endPoint,
	}

	return &vectorPathCommandLineTemp
}

// GetVectorPathCommandEnum return the enum type of this object
func (vectorPathCommandLine *VectorPathCommandLine) GetVectorPathCommandEnum() VectorPathCommandEnum {
	return VectorPathCommandLineType
}

// VectorPathCommandCubicBezierCurve A cubic Bézier curve to a given point
type VectorPathCommandCubicBezierCurve struct {
	tdCommon
	StartControlPoint *Point `json:"start_control_point"` // The start control point of the curve
	EndControlPoint   *Point `json:"end_control_point"`   // The end control point of the curve
	EndPoint          *Point `json:"end_point"`           // The end point of the curve
}

// MessageType return the string telegram-type of VectorPathCommandCubicBezierCurve
func (vectorPathCommandCubicBezierCurve *VectorPathCommandCubicBezierCurve) MessageType() string {
	return "vectorPathCommandCubicBezierCurve"
}

// NewVectorPathCommandCubicBezierCurve creates a new VectorPathCommandCubicBezierCurve
//
// @param startControlPoint The start control point of the curve
// @param endControlPoint The end control point of the curve
// @param endPoint The end point of the curve
func NewVectorPathCommandCubicBezierCurve(startControlPoint *Point, endControlPoint *Point, endPoint *Point) *VectorPathCommandCubicBezierCurve {
	vectorPathCommandCubicBezierCurveTemp := VectorPathCommandCubicBezierCurve{
		tdCommon:          tdCommon{Type: "vectorPathCommandCubicBezierCurve"},
		StartControlPoint: startControlPoint,
		EndControlPoint:   endControlPoint,
		EndPoint:          endPoint,
	}

	return &vectorPathCommandCubicBezierCurveTemp
}

// GetVectorPathCommandEnum return the enum type of this object
func (vectorPathCommandCubicBezierCurve *VectorPathCommandCubicBezierCurve) GetVectorPathCommandEnum() VectorPathCommandEnum {
	return VectorPathCommandCubicBezierCurveType
}
