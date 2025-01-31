package tdlib

import (
	"encoding/json"
)

// InlineKeyboardButton Represents a single button in an inline keyboard
type InlineKeyboardButton struct {
	tdCommon
	Text string                   `json:"text"` // Text of the button
	Type InlineKeyboardButtonType `json:"type"` // Type of the button
}

// MessageType return the string telegram-type of InlineKeyboardButton
func (inlineKeyboardButton *InlineKeyboardButton) MessageType() string {
	return "inlineKeyboardButton"
}

// NewInlineKeyboardButton creates a new InlineKeyboardButton
//
// @param text Text of the button
// @param typeParam Type of the button
func NewInlineKeyboardButton(text string, typeParam InlineKeyboardButtonType) *InlineKeyboardButton {
	inlineKeyboardButtonTemp := InlineKeyboardButton{
		tdCommon: tdCommon{Type: "inlineKeyboardButton"},
		Text:     text,
		Type:     typeParam,
	}

	return &inlineKeyboardButtonTemp
}

// UnmarshalJSON unmarshal to json
func (inlineKeyboardButton *InlineKeyboardButton) UnmarshalJSON(b []byte) error {
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

	inlineKeyboardButton.tdCommon = tempObj.tdCommon
	inlineKeyboardButton.Text = tempObj.Text

	fieldType, _ := unmarshalInlineKeyboardButtonType(objMap["type"])
	inlineKeyboardButton.Type = fieldType

	return nil
}
