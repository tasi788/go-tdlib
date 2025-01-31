package tdlib

import (
	"encoding/json"
	"fmt"
)

// InputChatPhoto Describes a photo to be set as a user profile or chat photo
type InputChatPhoto interface {
	GetInputChatPhotoEnum() InputChatPhotoEnum
}

// InputChatPhotoEnum Alias for abstract InputChatPhoto 'Sub-Classes', used as constant-enum here
type InputChatPhotoEnum string

// InputChatPhoto enums
const (
	InputChatPhotoPreviousType  InputChatPhotoEnum = "inputChatPhotoPrevious"
	InputChatPhotoStaticType    InputChatPhotoEnum = "inputChatPhotoStatic"
	InputChatPhotoAnimationType InputChatPhotoEnum = "inputChatPhotoAnimation"
)

func unmarshalInputChatPhoto(rawMsg *json.RawMessage) (InputChatPhoto, error) {

	if rawMsg == nil {
		return nil, nil
	}
	var objMap map[string]interface{}
	err := json.Unmarshal(*rawMsg, &objMap)
	if err != nil {
		return nil, err
	}

	switch InputChatPhotoEnum(objMap["@type"].(string)) {
	case InputChatPhotoPreviousType:
		var inputChatPhotoPrevious InputChatPhotoPrevious
		err := json.Unmarshal(*rawMsg, &inputChatPhotoPrevious)
		return &inputChatPhotoPrevious, err

	case InputChatPhotoStaticType:
		var inputChatPhotoStatic InputChatPhotoStatic
		err := json.Unmarshal(*rawMsg, &inputChatPhotoStatic)
		return &inputChatPhotoStatic, err

	case InputChatPhotoAnimationType:
		var inputChatPhotoAnimation InputChatPhotoAnimation
		err := json.Unmarshal(*rawMsg, &inputChatPhotoAnimation)
		return &inputChatPhotoAnimation, err

	default:
		return nil, fmt.Errorf("Error UnMarshaling, unknown type:" + objMap["@type"].(string))
	}
}

// InputChatPhotoPrevious A previously used profile photo of the current user
type InputChatPhotoPrevious struct {
	tdCommon
	ChatPhotoId JSONInt64 `json:"chat_photo_id"` // Identifier of the current user's profile photo to reuse
}

// MessageType return the string telegram-type of InputChatPhotoPrevious
func (inputChatPhotoPrevious *InputChatPhotoPrevious) MessageType() string {
	return "inputChatPhotoPrevious"
}

// NewInputChatPhotoPrevious creates a new InputChatPhotoPrevious
//
// @param chatPhotoId Identifier of the current user's profile photo to reuse
func NewInputChatPhotoPrevious(chatPhotoId JSONInt64) *InputChatPhotoPrevious {
	inputChatPhotoPreviousTemp := InputChatPhotoPrevious{
		tdCommon:    tdCommon{Type: "inputChatPhotoPrevious"},
		ChatPhotoId: chatPhotoId,
	}

	return &inputChatPhotoPreviousTemp
}

// GetInputChatPhotoEnum return the enum type of this object
func (inputChatPhotoPrevious *InputChatPhotoPrevious) GetInputChatPhotoEnum() InputChatPhotoEnum {
	return InputChatPhotoPreviousType
}

// InputChatPhotoStatic A static photo in JPEG format
type InputChatPhotoStatic struct {
	tdCommon
	Photo InputFile `json:"photo"` // Photo to be set as profile photo. Only inputFileLocal and inputFileGenerated are allowed
}

// MessageType return the string telegram-type of InputChatPhotoStatic
func (inputChatPhotoStatic *InputChatPhotoStatic) MessageType() string {
	return "inputChatPhotoStatic"
}

// NewInputChatPhotoStatic creates a new InputChatPhotoStatic
//
// @param photo Photo to be set as profile photo. Only inputFileLocal and inputFileGenerated are allowed
func NewInputChatPhotoStatic(photo InputFile) *InputChatPhotoStatic {
	inputChatPhotoStaticTemp := InputChatPhotoStatic{
		tdCommon: tdCommon{Type: "inputChatPhotoStatic"},
		Photo:    photo,
	}

	return &inputChatPhotoStaticTemp
}

// UnmarshalJSON unmarshal to json
func (inputChatPhotoStatic *InputChatPhotoStatic) UnmarshalJSON(b []byte) error {
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

	inputChatPhotoStatic.tdCommon = tempObj.tdCommon

	fieldPhoto, _ := unmarshalInputFile(objMap["photo"])
	inputChatPhotoStatic.Photo = fieldPhoto

	return nil
}

// GetInputChatPhotoEnum return the enum type of this object
func (inputChatPhotoStatic *InputChatPhotoStatic) GetInputChatPhotoEnum() InputChatPhotoEnum {
	return InputChatPhotoStaticType
}

// InputChatPhotoAnimation An animation in MPEG4 format; must be square, at most 10 seconds long, have width between 160 and 800 and be at most 2MB in size
type InputChatPhotoAnimation struct {
	tdCommon
	Animation          InputFile `json:"animation"`            // Animation to be set as profile photo. Only inputFileLocal and inputFileGenerated are allowed
	MainFrameTimestamp float64   `json:"main_frame_timestamp"` // Timestamp of the frame, which will be used as static chat photo
}

// MessageType return the string telegram-type of InputChatPhotoAnimation
func (inputChatPhotoAnimation *InputChatPhotoAnimation) MessageType() string {
	return "inputChatPhotoAnimation"
}

// NewInputChatPhotoAnimation creates a new InputChatPhotoAnimation
//
// @param animation Animation to be set as profile photo. Only inputFileLocal and inputFileGenerated are allowed
// @param mainFrameTimestamp Timestamp of the frame, which will be used as static chat photo
func NewInputChatPhotoAnimation(animation InputFile, mainFrameTimestamp float64) *InputChatPhotoAnimation {
	inputChatPhotoAnimationTemp := InputChatPhotoAnimation{
		tdCommon:           tdCommon{Type: "inputChatPhotoAnimation"},
		Animation:          animation,
		MainFrameTimestamp: mainFrameTimestamp,
	}

	return &inputChatPhotoAnimationTemp
}

// UnmarshalJSON unmarshal to json
func (inputChatPhotoAnimation *InputChatPhotoAnimation) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	tempObj := struct {
		tdCommon
		MainFrameTimestamp float64 `json:"main_frame_timestamp"` // Timestamp of the frame, which will be used as static chat photo
	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	inputChatPhotoAnimation.tdCommon = tempObj.tdCommon
	inputChatPhotoAnimation.MainFrameTimestamp = tempObj.MainFrameTimestamp

	fieldAnimation, _ := unmarshalInputFile(objMap["animation"])
	inputChatPhotoAnimation.Animation = fieldAnimation

	return nil
}

// GetInputChatPhotoEnum return the enum type of this object
func (inputChatPhotoAnimation *InputChatPhotoAnimation) GetInputChatPhotoEnum() InputChatPhotoEnum {
	return InputChatPhotoAnimationType
}
