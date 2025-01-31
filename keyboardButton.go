package tdlib

import (
	"encoding/json"
)

// KeyboardButton Represents a single button in a bot keyboard
type KeyboardButton struct {
	tdCommon
	Text string             `json:"text"` // Text of the button
	Type KeyboardButtonType `json:"type"` // Type of the button
}

// MessageType return the string telegram-type of KeyboardButton
func (keyboardButton *KeyboardButton) MessageType() string {
	return "keyboardButton"
}

// NewKeyboardButton creates a new KeyboardButton
//
// @param text Text of the button
// @param typeParam Type of the button
func NewKeyboardButton(text string, typeParam KeyboardButtonType) *KeyboardButton {
	keyboardButtonTemp := KeyboardButton{
		tdCommon: tdCommon{Type: "keyboardButton"},
		Text:     text,
		Type:     typeParam,
	}

	return &keyboardButtonTemp
}

// UnmarshalJSON unmarshal to json
func (keyboardButton *KeyboardButton) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	tempObj := struct {
		tdCommon
		Text string `json:"text"` // Text of the button

	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	keyboardButton.tdCommon = tempObj.tdCommon
	keyboardButton.Text = tempObj.Text

	fieldType, _ := unmarshalKeyboardButtonType(objMap["type"])
	keyboardButton.Type = fieldType

	return nil
}
