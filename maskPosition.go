package tdlib

import (
	"encoding/json"
)

// MaskPosition Position on a photo where a mask is placed
type MaskPosition struct {
	tdCommon
	Point  MaskPoint `json:"point"`   // Part of the face, relative to which the mask is placed
	XShift float64   `json:"x_shift"` // Shift by X-axis measured in widths of the mask scaled to the face size, from left to right. (For example, -1.0 will place the mask just to the left of the default mask position)
	YShift float64   `json:"y_shift"` // Shift by Y-axis measured in heights of the mask scaled to the face size, from top to bottom. (For example, 1.0 will place the mask just below the default mask position)
	Scale  float64   `json:"scale"`   // Mask scaling coefficient. (For example, 2.0 means a doubled size)
}

// MessageType return the string telegram-type of MaskPosition
func (maskPosition *MaskPosition) MessageType() string {
	return "maskPosition"
}

// NewMaskPosition creates a new MaskPosition
//
// @param point Part of the face, relative to which the mask is placed
// @param xShift Shift by X-axis measured in widths of the mask scaled to the face size, from left to right. (For example, -1.0 will place the mask just to the left of the default mask position)
// @param yShift Shift by Y-axis measured in heights of the mask scaled to the face size, from top to bottom. (For example, 1.0 will place the mask just below the default mask position)
// @param scale Mask scaling coefficient. (For example, 2.0 means a doubled size)
func NewMaskPosition(point MaskPoint, xShift float64, yShift float64, scale float64) *MaskPosition {
	maskPositionTemp := MaskPosition{
		tdCommon: tdCommon{Type: "maskPosition"},
		Point:    point,
		XShift:   xShift,
		YShift:   yShift,
		Scale:    scale,
	}

	return &maskPositionTemp
}

// UnmarshalJSON unmarshal to json
func (maskPosition *MaskPosition) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	tempObj := struct {
		tdCommon
		XShift float64 `json:"x_shift"` // Shift by X-axis measured in widths of the mask scaled to the face size, from left to right. (For example, -1.0 will place the mask just to the left of the default mask position)
		YShift float64 `json:"y_shift"` // Shift by Y-axis measured in heights of the mask scaled to the face size, from top to bottom. (For example, 1.0 will place the mask just below the default mask position)
		Scale  float64 `json:"scale"`   // Mask scaling coefficient. (For example, 2.0 means a doubled size)
	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	maskPosition.tdCommon = tempObj.tdCommon
	maskPosition.XShift = tempObj.XShift
	maskPosition.YShift = tempObj.YShift
	maskPosition.Scale = tempObj.Scale

	fieldPoint, _ := unmarshalMaskPoint(objMap["point"])
	maskPosition.Point = fieldPoint

	return nil
}
