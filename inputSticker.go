package tdlib

import (
	"encoding/json"
	"fmt"
)

// InputSticker Describes a sticker that needs to be added to a sticker set
type InputSticker interface {
	GetInputStickerEnum() InputStickerEnum
}

// InputStickerEnum Alias for abstract InputSticker 'Sub-Classes', used as constant-enum here
type InputStickerEnum string

// InputSticker enums
const (
	InputStickerStaticType   InputStickerEnum = "inputStickerStatic"
	InputStickerAnimatedType InputStickerEnum = "inputStickerAnimated"
)

func unmarshalInputSticker(rawMsg *json.RawMessage) (InputSticker, error) {

	if rawMsg == nil {
		return nil, nil
	}
	var objMap map[string]interface{}
	err := json.Unmarshal(*rawMsg, &objMap)
	if err != nil {
		return nil, err
	}

	switch InputStickerEnum(objMap["@type"].(string)) {
	case InputStickerStaticType:
		var inputStickerStatic InputStickerStatic
		err := json.Unmarshal(*rawMsg, &inputStickerStatic)
		return &inputStickerStatic, err

	case InputStickerAnimatedType:
		var inputStickerAnimated InputStickerAnimated
		err := json.Unmarshal(*rawMsg, &inputStickerAnimated)
		return &inputStickerAnimated, err

	default:
		return nil, fmt.Errorf("Error UnMarshaling, unknown type:" + objMap["@type"].(string))
	}
}

// InputStickerStatic A static sticker in PNG format, which will be converted to WEBP server-side
type InputStickerStatic struct {
	tdCommon
	Sticker      InputFile     `json:"sticker"`       // PNG image with the sticker; must be up to 512 KB in size and fit in a 512x512 square
	Emojis       string        `json:"emojis"`        // Emojis corresponding to the sticker
	MaskPosition *MaskPosition `json:"mask_position"` // For masks, position where the mask is placed; pass null if unspecified
}

// MessageType return the string telegram-type of InputStickerStatic
func (inputStickerStatic *InputStickerStatic) MessageType() string {
	return "inputStickerStatic"
}

// NewInputStickerStatic creates a new InputStickerStatic
//
// @param sticker PNG image with the sticker; must be up to 512 KB in size and fit in a 512x512 square
// @param emojis Emojis corresponding to the sticker
// @param maskPosition For masks, position where the mask is placed; pass null if unspecified
func NewInputStickerStatic(sticker InputFile, emojis string, maskPosition *MaskPosition) *InputStickerStatic {
	inputStickerStaticTemp := InputStickerStatic{
		tdCommon:     tdCommon{Type: "inputStickerStatic"},
		Sticker:      sticker,
		Emojis:       emojis,
		MaskPosition: maskPosition,
	}

	return &inputStickerStaticTemp
}

// UnmarshalJSON unmarshal to json
func (inputStickerStatic *InputStickerStatic) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	tempObj := struct {
		tdCommon
		Emojis       string        `json:"emojis"`        // Emojis corresponding to the sticker
		MaskPosition *MaskPosition `json:"mask_position"` // For masks, position where the mask is placed; pass null if unspecified
	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	inputStickerStatic.tdCommon = tempObj.tdCommon
	inputStickerStatic.Emojis = tempObj.Emojis
	inputStickerStatic.MaskPosition = tempObj.MaskPosition

	fieldSticker, _ := unmarshalInputFile(objMap["sticker"])
	inputStickerStatic.Sticker = fieldSticker

	return nil
}

// GetInputStickerEnum return the enum type of this object
func (inputStickerStatic *InputStickerStatic) GetInputStickerEnum() InputStickerEnum {
	return InputStickerStaticType
}

// InputStickerAnimated An animated sticker in TGS format
type InputStickerAnimated struct {
	tdCommon
	Sticker InputFile `json:"sticker"` // File with the animated sticker. Only local or uploaded within a week files are supported. See https://core.telegram.org/animated_stickers#technical-requirements for technical requirements
	Emojis  string    `json:"emojis"`  // Emojis corresponding to the sticker
}

// MessageType return the string telegram-type of InputStickerAnimated
func (inputStickerAnimated *InputStickerAnimated) MessageType() string {
	return "inputStickerAnimated"
}

// NewInputStickerAnimated creates a new InputStickerAnimated
//
// @param sticker File with the animated sticker. Only local or uploaded within a week files are supported. See https://core.telegram.org/animated_stickers#technical-requirements for technical requirements
// @param emojis Emojis corresponding to the sticker
func NewInputStickerAnimated(sticker InputFile, emojis string) *InputStickerAnimated {
	inputStickerAnimatedTemp := InputStickerAnimated{
		tdCommon: tdCommon{Type: "inputStickerAnimated"},
		Sticker:  sticker,
		Emojis:   emojis,
	}

	return &inputStickerAnimatedTemp
}

// UnmarshalJSON unmarshal to json
func (inputStickerAnimated *InputStickerAnimated) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	tempObj := struct {
		tdCommon
		Emojis string `json:"emojis"` // Emojis corresponding to the sticker
	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	inputStickerAnimated.tdCommon = tempObj.tdCommon
	inputStickerAnimated.Emojis = tempObj.Emojis

	fieldSticker, _ := unmarshalInputFile(objMap["sticker"])
	inputStickerAnimated.Sticker = fieldSticker

	return nil
}

// GetInputStickerEnum return the enum type of this object
func (inputStickerAnimated *InputStickerAnimated) GetInputStickerEnum() InputStickerEnum {
	return InputStickerAnimatedType
}
