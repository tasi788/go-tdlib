package tdlib

import (
	"encoding/json"
	"fmt"
)

// ReplyMarkup Contains a description of a custom keyboard and actions that can be done with it to quickly reply to bots
type ReplyMarkup interface {
	GetReplyMarkupEnum() ReplyMarkupEnum
}

// ReplyMarkupEnum Alias for abstract ReplyMarkup 'Sub-Classes', used as constant-enum here
type ReplyMarkupEnum string

// ReplyMarkup enums
const (
	ReplyMarkupRemoveKeyboardType ReplyMarkupEnum = "replyMarkupRemoveKeyboard"
	ReplyMarkupForceReplyType     ReplyMarkupEnum = "replyMarkupForceReply"
	ReplyMarkupShowKeyboardType   ReplyMarkupEnum = "replyMarkupShowKeyboard"
	ReplyMarkupInlineKeyboardType ReplyMarkupEnum = "replyMarkupInlineKeyboard"
)

func unmarshalReplyMarkup(rawMsg *json.RawMessage) (ReplyMarkup, error) {

	if rawMsg == nil {
		return nil, nil
	}
	var objMap map[string]interface{}
	err := json.Unmarshal(*rawMsg, &objMap)
	if err != nil {
		return nil, err
	}

	switch ReplyMarkupEnum(objMap["@type"].(string)) {
	case ReplyMarkupRemoveKeyboardType:
		var replyMarkupRemoveKeyboard ReplyMarkupRemoveKeyboard
		err := json.Unmarshal(*rawMsg, &replyMarkupRemoveKeyboard)
		return &replyMarkupRemoveKeyboard, err

	case ReplyMarkupForceReplyType:
		var replyMarkupForceReply ReplyMarkupForceReply
		err := json.Unmarshal(*rawMsg, &replyMarkupForceReply)
		return &replyMarkupForceReply, err

	case ReplyMarkupShowKeyboardType:
		var replyMarkupShowKeyboard ReplyMarkupShowKeyboard
		err := json.Unmarshal(*rawMsg, &replyMarkupShowKeyboard)
		return &replyMarkupShowKeyboard, err

	case ReplyMarkupInlineKeyboardType:
		var replyMarkupInlineKeyboard ReplyMarkupInlineKeyboard
		err := json.Unmarshal(*rawMsg, &replyMarkupInlineKeyboard)
		return &replyMarkupInlineKeyboard, err

	default:
		return nil, fmt.Errorf("Error UnMarshaling, unknown type:" + objMap["@type"].(string))
	}
}

// ReplyMarkupRemoveKeyboard Instructs application to remove the keyboard once this message has been received. This kind of keyboard can't be received in an incoming message; instead, UpdateChatReplyMarkup with message_id == 0 will be sent
type ReplyMarkupRemoveKeyboard struct {
	tdCommon
	IsPersonal bool `json:"is_personal"` // True, if the keyboard is removed only for the mentioned users or the target user of a reply
}

// MessageType return the string telegram-type of ReplyMarkupRemoveKeyboard
func (replyMarkupRemoveKeyboard *ReplyMarkupRemoveKeyboard) MessageType() string {
	return "replyMarkupRemoveKeyboard"
}

// NewReplyMarkupRemoveKeyboard creates a new ReplyMarkupRemoveKeyboard
//
// @param isPersonal True, if the keyboard is removed only for the mentioned users or the target user of a reply
func NewReplyMarkupRemoveKeyboard(isPersonal bool) *ReplyMarkupRemoveKeyboard {
	replyMarkupRemoveKeyboardTemp := ReplyMarkupRemoveKeyboard{
		tdCommon:   tdCommon{Type: "replyMarkupRemoveKeyboard"},
		IsPersonal: isPersonal,
	}

	return &replyMarkupRemoveKeyboardTemp
}

// GetReplyMarkupEnum return the enum type of this object
func (replyMarkupRemoveKeyboard *ReplyMarkupRemoveKeyboard) GetReplyMarkupEnum() ReplyMarkupEnum {
	return ReplyMarkupRemoveKeyboardType
}

// ReplyMarkupForceReply Instructs application to force a reply to this message
type ReplyMarkupForceReply struct {
	tdCommon
	IsPersonal            bool   `json:"is_personal"`             // True, if a forced reply must automatically be shown to the current user. For outgoing messages, specify true to show the forced reply only for the mentioned users and for the target user of a reply
	InputFieldPlaceholder string `json:"input_field_placeholder"` // If non-empty, the placeholder to be shown in the input field when the reply is active; 0-64 characters
}

// MessageType return the string telegram-type of ReplyMarkupForceReply
func (replyMarkupForceReply *ReplyMarkupForceReply) MessageType() string {
	return "replyMarkupForceReply"
}

// NewReplyMarkupForceReply creates a new ReplyMarkupForceReply
//
// @param isPersonal True, if a forced reply must automatically be shown to the current user. For outgoing messages, specify true to show the forced reply only for the mentioned users and for the target user of a reply
// @param inputFieldPlaceholder If non-empty, the placeholder to be shown in the input field when the reply is active; 0-64 characters
func NewReplyMarkupForceReply(isPersonal bool, inputFieldPlaceholder string) *ReplyMarkupForceReply {
	replyMarkupForceReplyTemp := ReplyMarkupForceReply{
		tdCommon:              tdCommon{Type: "replyMarkupForceReply"},
		IsPersonal:            isPersonal,
		InputFieldPlaceholder: inputFieldPlaceholder,
	}

	return &replyMarkupForceReplyTemp
}

// GetReplyMarkupEnum return the enum type of this object
func (replyMarkupForceReply *ReplyMarkupForceReply) GetReplyMarkupEnum() ReplyMarkupEnum {
	return ReplyMarkupForceReplyType
}

// ReplyMarkupShowKeyboard Contains a custom keyboard layout to quickly reply to bots
type ReplyMarkupShowKeyboard struct {
	tdCommon
	Rows                  [][]KeyboardButton `json:"rows"`                    // A list of rows of bot keyboard buttons
	ResizeKeyboard        bool               `json:"resize_keyboard"`         // True, if the application needs to resize the keyboard vertically
	OneTime               bool               `json:"one_time"`                // True, if the application needs to hide the keyboard after use
	IsPersonal            bool               `json:"is_personal"`             // True, if the keyboard must automatically be shown to the current user. For outgoing messages, specify true to show the keyboard only for the mentioned users and for the target user of a reply
	InputFieldPlaceholder string             `json:"input_field_placeholder"` // If non-empty, the placeholder to be shown in the input field when the keyboard is active; 0-64 characters
}

// MessageType return the string telegram-type of ReplyMarkupShowKeyboard
func (replyMarkupShowKeyboard *ReplyMarkupShowKeyboard) MessageType() string {
	return "replyMarkupShowKeyboard"
}

// NewReplyMarkupShowKeyboard creates a new ReplyMarkupShowKeyboard
//
// @param rows A list of rows of bot keyboard buttons
// @param resizeKeyboard True, if the application needs to resize the keyboard vertically
// @param oneTime True, if the application needs to hide the keyboard after use
// @param isPersonal True, if the keyboard must automatically be shown to the current user. For outgoing messages, specify true to show the keyboard only for the mentioned users and for the target user of a reply
// @param inputFieldPlaceholder If non-empty, the placeholder to be shown in the input field when the keyboard is active; 0-64 characters
func NewReplyMarkupShowKeyboard(rows [][]KeyboardButton, resizeKeyboard bool, oneTime bool, isPersonal bool, inputFieldPlaceholder string) *ReplyMarkupShowKeyboard {
	replyMarkupShowKeyboardTemp := ReplyMarkupShowKeyboard{
		tdCommon:              tdCommon{Type: "replyMarkupShowKeyboard"},
		Rows:                  rows,
		ResizeKeyboard:        resizeKeyboard,
		OneTime:               oneTime,
		IsPersonal:            isPersonal,
		InputFieldPlaceholder: inputFieldPlaceholder,
	}

	return &replyMarkupShowKeyboardTemp
}

// GetReplyMarkupEnum return the enum type of this object
func (replyMarkupShowKeyboard *ReplyMarkupShowKeyboard) GetReplyMarkupEnum() ReplyMarkupEnum {
	return ReplyMarkupShowKeyboardType
}

// ReplyMarkupInlineKeyboard Contains an inline keyboard layout
type ReplyMarkupInlineKeyboard struct {
	tdCommon
	Rows [][]InlineKeyboardButton `json:"rows"` // A list of rows of inline keyboard buttons
}

// MessageType return the string telegram-type of ReplyMarkupInlineKeyboard
func (replyMarkupInlineKeyboard *ReplyMarkupInlineKeyboard) MessageType() string {
	return "replyMarkupInlineKeyboard"
}

// NewReplyMarkupInlineKeyboard creates a new ReplyMarkupInlineKeyboard
//
// @param rows A list of rows of inline keyboard buttons
func NewReplyMarkupInlineKeyboard(rows [][]InlineKeyboardButton) *ReplyMarkupInlineKeyboard {
	replyMarkupInlineKeyboardTemp := ReplyMarkupInlineKeyboard{
		tdCommon: tdCommon{Type: "replyMarkupInlineKeyboard"},
		Rows:     rows,
	}

	return &replyMarkupInlineKeyboardTemp
}

// GetReplyMarkupEnum return the enum type of this object
func (replyMarkupInlineKeyboard *ReplyMarkupInlineKeyboard) GetReplyMarkupEnum() ReplyMarkupEnum {
	return ReplyMarkupInlineKeyboardType
}
