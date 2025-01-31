package tdlib

import (
	"encoding/json"
	"fmt"
)

// BackgroundType Describes the type of a background
type BackgroundType interface {
	GetBackgroundTypeEnum() BackgroundTypeEnum
}

// BackgroundTypeEnum Alias for abstract BackgroundType 'Sub-Classes', used as constant-enum here
type BackgroundTypeEnum string

// BackgroundType enums
const (
	BackgroundTypeWallpaperType BackgroundTypeEnum = "backgroundTypeWallpaper"
	BackgroundTypePatternType   BackgroundTypeEnum = "backgroundTypePattern"
	BackgroundTypeFillType      BackgroundTypeEnum = "backgroundTypeFill"
)

func unmarshalBackgroundType(rawMsg *json.RawMessage) (BackgroundType, error) {

	if rawMsg == nil {
		return nil, nil
	}
	var objMap map[string]interface{}
	err := json.Unmarshal(*rawMsg, &objMap)
	if err != nil {
		return nil, err
	}

	switch BackgroundTypeEnum(objMap["@type"].(string)) {
	case BackgroundTypeWallpaperType:
		var backgroundTypeWallpaper BackgroundTypeWallpaper
		err := json.Unmarshal(*rawMsg, &backgroundTypeWallpaper)
		return &backgroundTypeWallpaper, err

	case BackgroundTypePatternType:
		var backgroundTypePattern BackgroundTypePattern
		err := json.Unmarshal(*rawMsg, &backgroundTypePattern)
		return &backgroundTypePattern, err

	case BackgroundTypeFillType:
		var backgroundTypeFill BackgroundTypeFill
		err := json.Unmarshal(*rawMsg, &backgroundTypeFill)
		return &backgroundTypeFill, err

	default:
		return nil, fmt.Errorf("Error UnMarshaling, unknown type:" + objMap["@type"].(string))
	}
}

// BackgroundTypeWallpaper A wallpaper in JPEG format
type BackgroundTypeWallpaper struct {
	tdCommon
	IsBlurred bool `json:"is_blurred"` // True, if the wallpaper must be downscaled to fit in 450x450 square and then box-blurred with radius 12
	IsMoving  bool `json:"is_moving"`  // True, if the background needs to be slightly moved when device is tilted
}

// MessageType return the string telegram-type of BackgroundTypeWallpaper
func (backgroundTypeWallpaper *BackgroundTypeWallpaper) MessageType() string {
	return "backgroundTypeWallpaper"
}

// NewBackgroundTypeWallpaper creates a new BackgroundTypeWallpaper
//
// @param isBlurred True, if the wallpaper must be downscaled to fit in 450x450 square and then box-blurred with radius 12
// @param isMoving True, if the background needs to be slightly moved when device is tilted
func NewBackgroundTypeWallpaper(isBlurred bool, isMoving bool) *BackgroundTypeWallpaper {
	backgroundTypeWallpaperTemp := BackgroundTypeWallpaper{
		tdCommon:  tdCommon{Type: "backgroundTypeWallpaper"},
		IsBlurred: isBlurred,
		IsMoving:  isMoving,
	}

	return &backgroundTypeWallpaperTemp
}

// GetBackgroundTypeEnum return the enum type of this object
func (backgroundTypeWallpaper *BackgroundTypeWallpaper) GetBackgroundTypeEnum() BackgroundTypeEnum {
	return BackgroundTypeWallpaperType
}

// BackgroundTypePattern A PNG or TGV (gzipped subset of SVG with MIME type "application/x-tgwallpattern") pattern to be combined with the background fill chosen by the user
type BackgroundTypePattern struct {
	tdCommon
	Fill       BackgroundFill `json:"fill"`        // Fill of the background
	Intensity  int32          `json:"intensity"`   // Intensity of the pattern when it is shown above the filled background; 0-100.
	IsInverted bool           `json:"is_inverted"` // True, if the background fill must be applied only to the pattern itself. All other pixels are black in this case. For dark themes only
	IsMoving   bool           `json:"is_moving"`   // True, if the background needs to be slightly moved when device is tilted
}

// MessageType return the string telegram-type of BackgroundTypePattern
func (backgroundTypePattern *BackgroundTypePattern) MessageType() string {
	return "backgroundTypePattern"
}

// NewBackgroundTypePattern creates a new BackgroundTypePattern
//
// @param fill Fill of the background
// @param intensity Intensity of the pattern when it is shown above the filled background; 0-100.
// @param isInverted True, if the background fill must be applied only to the pattern itself. All other pixels are black in this case. For dark themes only
// @param isMoving True, if the background needs to be slightly moved when device is tilted
func NewBackgroundTypePattern(fill BackgroundFill, intensity int32, isInverted bool, isMoving bool) *BackgroundTypePattern {
	backgroundTypePatternTemp := BackgroundTypePattern{
		tdCommon:   tdCommon{Type: "backgroundTypePattern"},
		Fill:       fill,
		Intensity:  intensity,
		IsInverted: isInverted,
		IsMoving:   isMoving,
	}

	return &backgroundTypePatternTemp
}

// UnmarshalJSON unmarshal to json
func (backgroundTypePattern *BackgroundTypePattern) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	tempObj := struct {
		tdCommon
		Intensity  int32 `json:"intensity"`   // Intensity of the pattern when it is shown above the filled background; 0-100.
		IsInverted bool  `json:"is_inverted"` // True, if the background fill must be applied only to the pattern itself. All other pixels are black in this case. For dark themes only
		IsMoving   bool  `json:"is_moving"`   // True, if the background needs to be slightly moved when device is tilted
	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	backgroundTypePattern.tdCommon = tempObj.tdCommon
	backgroundTypePattern.Intensity = tempObj.Intensity
	backgroundTypePattern.IsInverted = tempObj.IsInverted
	backgroundTypePattern.IsMoving = tempObj.IsMoving

	fieldFill, _ := unmarshalBackgroundFill(objMap["fill"])
	backgroundTypePattern.Fill = fieldFill

	return nil
}

// GetBackgroundTypeEnum return the enum type of this object
func (backgroundTypePattern *BackgroundTypePattern) GetBackgroundTypeEnum() BackgroundTypeEnum {
	return BackgroundTypePatternType
}

// BackgroundTypeFill A filled background
type BackgroundTypeFill struct {
	tdCommon
	Fill BackgroundFill `json:"fill"` // The background fill
}

// MessageType return the string telegram-type of BackgroundTypeFill
func (backgroundTypeFill *BackgroundTypeFill) MessageType() string {
	return "backgroundTypeFill"
}

// NewBackgroundTypeFill creates a new BackgroundTypeFill
//
// @param fill The background fill
func NewBackgroundTypeFill(fill BackgroundFill) *BackgroundTypeFill {
	backgroundTypeFillTemp := BackgroundTypeFill{
		tdCommon: tdCommon{Type: "backgroundTypeFill"},
		Fill:     fill,
	}

	return &backgroundTypeFillTemp
}

// UnmarshalJSON unmarshal to json
func (backgroundTypeFill *BackgroundTypeFill) UnmarshalJSON(b []byte) error {
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

	backgroundTypeFill.tdCommon = tempObj.tdCommon

	fieldFill, _ := unmarshalBackgroundFill(objMap["fill"])
	backgroundTypeFill.Fill = fieldFill

	return nil
}

// GetBackgroundTypeEnum return the enum type of this object
func (backgroundTypeFill *BackgroundTypeFill) GetBackgroundTypeEnum() BackgroundTypeEnum {
	return BackgroundTypeFillType
}
