package tdlib

import (
	"encoding/json"
	"fmt"
)

// KeyboardButtonType Describes a keyboard button type
type KeyboardButtonType interface {
	GetKeyboardButtonTypeEnum() KeyboardButtonTypeEnum
}

// KeyboardButtonTypeEnum Alias for abstract KeyboardButtonType 'Sub-Classes', used as constant-enum here
type KeyboardButtonTypeEnum string

// KeyboardButtonType enums
const (
	KeyboardButtonTypeTextType               KeyboardButtonTypeEnum = "keyboardButtonTypeText"
	KeyboardButtonTypeRequestPhoneNumberType KeyboardButtonTypeEnum = "keyboardButtonTypeRequestPhoneNumber"
	KeyboardButtonTypeRequestLocationType    KeyboardButtonTypeEnum = "keyboardButtonTypeRequestLocation"
	KeyboardButtonTypeRequestPollType        KeyboardButtonTypeEnum = "keyboardButtonTypeRequestPoll"
)

func unmarshalKeyboardButtonType(rawMsg *json.RawMessage) (KeyboardButtonType, error) {

	if rawMsg == nil {
		return nil, nil
	}
	var objMap map[string]interface{}
	err := json.Unmarshal(*rawMsg, &objMap)
	if err != nil {
		return nil, err
	}

	switch KeyboardButtonTypeEnum(objMap["@type"].(string)) {
	case KeyboardButtonTypeTextType:
		var keyboardButtonTypeText KeyboardButtonTypeText
		err := json.Unmarshal(*rawMsg, &keyboardButtonTypeText)
		return &keyboardButtonTypeText, err

	case KeyboardButtonTypeRequestPhoneNumberType:
		var keyboardButtonTypeRequestPhoneNumber KeyboardButtonTypeRequestPhoneNumber
		err := json.Unmarshal(*rawMsg, &keyboardButtonTypeRequestPhoneNumber)
		return &keyboardButtonTypeRequestPhoneNumber, err

	case KeyboardButtonTypeRequestLocationType:
		var keyboardButtonTypeRequestLocation KeyboardButtonTypeRequestLocation
		err := json.Unmarshal(*rawMsg, &keyboardButtonTypeRequestLocation)
		return &keyboardButtonTypeRequestLocation, err

	case KeyboardButtonTypeRequestPollType:
		var keyboardButtonTypeRequestPoll KeyboardButtonTypeRequestPoll
		err := json.Unmarshal(*rawMsg, &keyboardButtonTypeRequestPoll)
		return &keyboardButtonTypeRequestPoll, err

	default:
		return nil, fmt.Errorf("Error UnMarshaling, unknown type:" + objMap["@type"].(string))
	}
}

// KeyboardButtonTypeText A simple button, with text that must be sent when the button is pressed
type KeyboardButtonTypeText struct {
	tdCommon
}

// MessageType return the string telegram-type of KeyboardButtonTypeText
func (keyboardButtonTypeText *KeyboardButtonTypeText) MessageType() string {
	return "keyboardButtonTypeText"
}

// NewKeyboardButtonTypeText creates a new KeyboardButtonTypeText
//
func NewKeyboardButtonTypeText() *KeyboardButtonTypeText {
	keyboardButtonTypeTextTemp := KeyboardButtonTypeText{
		tdCommon: tdCommon{Type: "keyboardButtonTypeText"},
	}

	return &keyboardButtonTypeTextTemp
}

// GetKeyboardButtonTypeEnum return the enum type of this object
func (keyboardButtonTypeText *KeyboardButtonTypeText) GetKeyboardButtonTypeEnum() KeyboardButtonTypeEnum {
	return KeyboardButtonTypeTextType
}

// KeyboardButtonTypeRequestPhoneNumber A button that sends the user's phone number when pressed; available only in private chats
type KeyboardButtonTypeRequestPhoneNumber struct {
	tdCommon
}

// MessageType return the string telegram-type of KeyboardButtonTypeRequestPhoneNumber
func (keyboardButtonTypeRequestPhoneNumber *KeyboardButtonTypeRequestPhoneNumber) MessageType() string {
	return "keyboardButtonTypeRequestPhoneNumber"
}

// NewKeyboardButtonTypeRequestPhoneNumber creates a new KeyboardButtonTypeRequestPhoneNumber
//
func NewKeyboardButtonTypeRequestPhoneNumber() *KeyboardButtonTypeRequestPhoneNumber {
	keyboardButtonTypeRequestPhoneNumberTemp := KeyboardButtonTypeRequestPhoneNumber{
		tdCommon: tdCommon{Type: "keyboardButtonTypeRequestPhoneNumber"},
	}

	return &keyboardButtonTypeRequestPhoneNumberTemp
}

// GetKeyboardButtonTypeEnum return the enum type of this object
func (keyboardButtonTypeRequestPhoneNumber *KeyboardButtonTypeRequestPhoneNumber) GetKeyboardButtonTypeEnum() KeyboardButtonTypeEnum {
	return KeyboardButtonTypeRequestPhoneNumberType
}

// KeyboardButtonTypeRequestLocation A button that sends the user's location when pressed; available only in private chats
type KeyboardButtonTypeRequestLocation struct {
	tdCommon
}

// MessageType return the string telegram-type of KeyboardButtonTypeRequestLocation
func (keyboardButtonTypeRequestLocation *KeyboardButtonTypeRequestLocation) MessageType() string {
	return "keyboardButtonTypeRequestLocation"
}

// NewKeyboardButtonTypeRequestLocation creates a new KeyboardButtonTypeRequestLocation
//
func NewKeyboardButtonTypeRequestLocation() *KeyboardButtonTypeRequestLocation {
	keyboardButtonTypeRequestLocationTemp := KeyboardButtonTypeRequestLocation{
		tdCommon: tdCommon{Type: "keyboardButtonTypeRequestLocation"},
	}

	return &keyboardButtonTypeRequestLocationTemp
}

// GetKeyboardButtonTypeEnum return the enum type of this object
func (keyboardButtonTypeRequestLocation *KeyboardButtonTypeRequestLocation) GetKeyboardButtonTypeEnum() KeyboardButtonTypeEnum {
	return KeyboardButtonTypeRequestLocationType
}

// KeyboardButtonTypeRequestPoll A button that allows the user to create and send a poll when pressed; available only in private chats
type KeyboardButtonTypeRequestPoll struct {
	tdCommon
	ForceRegular bool `json:"force_regular"` // If true, only regular polls must be allowed to create
	ForceQuiz    bool `json:"force_quiz"`    // If true, only polls in quiz mode must be allowed to create
}

// MessageType return the string telegram-type of KeyboardButtonTypeRequestPoll
func (keyboardButtonTypeRequestPoll *KeyboardButtonTypeRequestPoll) MessageType() string {
	return "keyboardButtonTypeRequestPoll"
}

// NewKeyboardButtonTypeRequestPoll creates a new KeyboardButtonTypeRequestPoll
//
// @param forceRegular If true, only regular polls must be allowed to create
// @param forceQuiz If true, only polls in quiz mode must be allowed to create
func NewKeyboardButtonTypeRequestPoll(forceRegular bool, forceQuiz bool) *KeyboardButtonTypeRequestPoll {
	keyboardButtonTypeRequestPollTemp := KeyboardButtonTypeRequestPoll{
		tdCommon:     tdCommon{Type: "keyboardButtonTypeRequestPoll"},
		ForceRegular: forceRegular,
		ForceQuiz:    forceQuiz,
	}

	return &keyboardButtonTypeRequestPollTemp
}

// GetKeyboardButtonTypeEnum return the enum type of this object
func (keyboardButtonTypeRequestPoll *KeyboardButtonTypeRequestPoll) GetKeyboardButtonTypeEnum() KeyboardButtonTypeEnum {
	return KeyboardButtonTypeRequestPollType
}
