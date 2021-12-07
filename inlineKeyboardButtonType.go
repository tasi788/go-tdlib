package tdlib

import (
	"encoding/json"
	"fmt"
)

// InlineKeyboardButtonType Describes the type of an inline keyboard button
type InlineKeyboardButtonType interface {
	GetInlineKeyboardButtonTypeEnum() InlineKeyboardButtonTypeEnum
}

// InlineKeyboardButtonTypeEnum Alias for abstract InlineKeyboardButtonType 'Sub-Classes', used as constant-enum here
type InlineKeyboardButtonTypeEnum string

// InlineKeyboardButtonType enums
const (
	InlineKeyboardButtonTypeUrlType                  InlineKeyboardButtonTypeEnum = "inlineKeyboardButtonTypeUrl"
	InlineKeyboardButtonTypeLoginUrlType             InlineKeyboardButtonTypeEnum = "inlineKeyboardButtonTypeLoginUrl"
	InlineKeyboardButtonTypeCallbackType             InlineKeyboardButtonTypeEnum = "inlineKeyboardButtonTypeCallback"
	InlineKeyboardButtonTypeCallbackWithPasswordType InlineKeyboardButtonTypeEnum = "inlineKeyboardButtonTypeCallbackWithPassword"
	InlineKeyboardButtonTypeCallbackGameType         InlineKeyboardButtonTypeEnum = "inlineKeyboardButtonTypeCallbackGame"
	InlineKeyboardButtonTypeSwitchInlineType         InlineKeyboardButtonTypeEnum = "inlineKeyboardButtonTypeSwitchInline"
	InlineKeyboardButtonTypeBuyType                  InlineKeyboardButtonTypeEnum = "inlineKeyboardButtonTypeBuy"
	InlineKeyboardButtonTypeUserType                 InlineKeyboardButtonTypeEnum = "inlineKeyboardButtonTypeUser"
)

func unmarshalInlineKeyboardButtonType(rawMsg *json.RawMessage) (InlineKeyboardButtonType, error) {

	if rawMsg == nil {
		return nil, nil
	}
	var objMap map[string]interface{}
	err := json.Unmarshal(*rawMsg, &objMap)
	if err != nil {
		return nil, err
	}

	switch InlineKeyboardButtonTypeEnum(objMap["@type"].(string)) {
	case InlineKeyboardButtonTypeUrlType:
		var inlineKeyboardButtonTypeUrl InlineKeyboardButtonTypeUrl
		err := json.Unmarshal(*rawMsg, &inlineKeyboardButtonTypeUrl)
		return &inlineKeyboardButtonTypeUrl, err

	case InlineKeyboardButtonTypeLoginUrlType:
		var inlineKeyboardButtonTypeLoginUrl InlineKeyboardButtonTypeLoginUrl
		err := json.Unmarshal(*rawMsg, &inlineKeyboardButtonTypeLoginUrl)
		return &inlineKeyboardButtonTypeLoginUrl, err

	case InlineKeyboardButtonTypeCallbackType:
		var inlineKeyboardButtonTypeCallback InlineKeyboardButtonTypeCallback
		err := json.Unmarshal(*rawMsg, &inlineKeyboardButtonTypeCallback)
		return &inlineKeyboardButtonTypeCallback, err

	case InlineKeyboardButtonTypeCallbackWithPasswordType:
		var inlineKeyboardButtonTypeCallbackWithPassword InlineKeyboardButtonTypeCallbackWithPassword
		err := json.Unmarshal(*rawMsg, &inlineKeyboardButtonTypeCallbackWithPassword)
		return &inlineKeyboardButtonTypeCallbackWithPassword, err

	case InlineKeyboardButtonTypeCallbackGameType:
		var inlineKeyboardButtonTypeCallbackGame InlineKeyboardButtonTypeCallbackGame
		err := json.Unmarshal(*rawMsg, &inlineKeyboardButtonTypeCallbackGame)
		return &inlineKeyboardButtonTypeCallbackGame, err

	case InlineKeyboardButtonTypeSwitchInlineType:
		var inlineKeyboardButtonTypeSwitchInline InlineKeyboardButtonTypeSwitchInline
		err := json.Unmarshal(*rawMsg, &inlineKeyboardButtonTypeSwitchInline)
		return &inlineKeyboardButtonTypeSwitchInline, err

	case InlineKeyboardButtonTypeBuyType:
		var inlineKeyboardButtonTypeBuy InlineKeyboardButtonTypeBuy
		err := json.Unmarshal(*rawMsg, &inlineKeyboardButtonTypeBuy)
		return &inlineKeyboardButtonTypeBuy, err

	case InlineKeyboardButtonTypeUserType:
		var inlineKeyboardButtonTypeUser InlineKeyboardButtonTypeUser
		err := json.Unmarshal(*rawMsg, &inlineKeyboardButtonTypeUser)
		return &inlineKeyboardButtonTypeUser, err

	default:
		return nil, fmt.Errorf("Error UnMarshaling, unknown type:" + objMap["@type"].(string))
	}
}

// InlineKeyboardButtonTypeUrl A button that opens a specified URL
type InlineKeyboardButtonTypeUrl struct {
	tdCommon
	Url string `json:"url"` // HTTP or tg:// URL to open
}

// MessageType return the string telegram-type of InlineKeyboardButtonTypeUrl
func (inlineKeyboardButtonTypeUrl *InlineKeyboardButtonTypeUrl) MessageType() string {
	return "inlineKeyboardButtonTypeUrl"
}

// NewInlineKeyboardButtonTypeUrl creates a new InlineKeyboardButtonTypeUrl
//
// @param url HTTP or tg:// URL to open
func NewInlineKeyboardButtonTypeUrl(url string) *InlineKeyboardButtonTypeUrl {
	inlineKeyboardButtonTypeUrlTemp := InlineKeyboardButtonTypeUrl{
		tdCommon: tdCommon{Type: "inlineKeyboardButtonTypeUrl"},
		Url:      url,
	}

	return &inlineKeyboardButtonTypeUrlTemp
}

// GetInlineKeyboardButtonTypeEnum return the enum type of this object
func (inlineKeyboardButtonTypeUrl *InlineKeyboardButtonTypeUrl) GetInlineKeyboardButtonTypeEnum() InlineKeyboardButtonTypeEnum {
	return InlineKeyboardButtonTypeUrlType
}

// InlineKeyboardButtonTypeLoginUrl A button that opens a specified URL and automatically authorize the current user if allowed to do so
type InlineKeyboardButtonTypeLoginUrl struct {
	tdCommon
	Url         string `json:"url"`          // An HTTP URL to open
	Id          int64  `json:"id"`           // Unique button identifier
	ForwardText string `json:"forward_text"` // If non-empty, new text of the button in forwarded messages
}

// MessageType return the string telegram-type of InlineKeyboardButtonTypeLoginUrl
func (inlineKeyboardButtonTypeLoginUrl *InlineKeyboardButtonTypeLoginUrl) MessageType() string {
	return "inlineKeyboardButtonTypeLoginUrl"
}

// NewInlineKeyboardButtonTypeLoginUrl creates a new InlineKeyboardButtonTypeLoginUrl
//
// @param url An HTTP URL to open
// @param id Unique button identifier
// @param forwardText If non-empty, new text of the button in forwarded messages
func NewInlineKeyboardButtonTypeLoginUrl(url string, id int64, forwardText string) *InlineKeyboardButtonTypeLoginUrl {
	inlineKeyboardButtonTypeLoginUrlTemp := InlineKeyboardButtonTypeLoginUrl{
		tdCommon:    tdCommon{Type: "inlineKeyboardButtonTypeLoginUrl"},
		Url:         url,
		Id:          id,
		ForwardText: forwardText,
	}

	return &inlineKeyboardButtonTypeLoginUrlTemp
}

// GetInlineKeyboardButtonTypeEnum return the enum type of this object
func (inlineKeyboardButtonTypeLoginUrl *InlineKeyboardButtonTypeLoginUrl) GetInlineKeyboardButtonTypeEnum() InlineKeyboardButtonTypeEnum {
	return InlineKeyboardButtonTypeLoginUrlType
}

// InlineKeyboardButtonTypeCallback A button that sends a callback query to a bot
type InlineKeyboardButtonTypeCallback struct {
	tdCommon
	Data []byte `json:"data"` // Data to be sent to the bot via a callback query
}

// MessageType return the string telegram-type of InlineKeyboardButtonTypeCallback
func (inlineKeyboardButtonTypeCallback *InlineKeyboardButtonTypeCallback) MessageType() string {
	return "inlineKeyboardButtonTypeCallback"
}

// NewInlineKeyboardButtonTypeCallback creates a new InlineKeyboardButtonTypeCallback
//
// @param data Data to be sent to the bot via a callback query
func NewInlineKeyboardButtonTypeCallback(data []byte) *InlineKeyboardButtonTypeCallback {
	inlineKeyboardButtonTypeCallbackTemp := InlineKeyboardButtonTypeCallback{
		tdCommon: tdCommon{Type: "inlineKeyboardButtonTypeCallback"},
		Data:     data,
	}

	return &inlineKeyboardButtonTypeCallbackTemp
}

// GetInlineKeyboardButtonTypeEnum return the enum type of this object
func (inlineKeyboardButtonTypeCallback *InlineKeyboardButtonTypeCallback) GetInlineKeyboardButtonTypeEnum() InlineKeyboardButtonTypeEnum {
	return InlineKeyboardButtonTypeCallbackType
}

// InlineKeyboardButtonTypeCallbackWithPassword A button that asks for password of the current user and then sends a callback query to a bot
type InlineKeyboardButtonTypeCallbackWithPassword struct {
	tdCommon
	Data []byte `json:"data"` // Data to be sent to the bot via a callback query
}

// MessageType return the string telegram-type of InlineKeyboardButtonTypeCallbackWithPassword
func (inlineKeyboardButtonTypeCallbackWithPassword *InlineKeyboardButtonTypeCallbackWithPassword) MessageType() string {
	return "inlineKeyboardButtonTypeCallbackWithPassword"
}

// NewInlineKeyboardButtonTypeCallbackWithPassword creates a new InlineKeyboardButtonTypeCallbackWithPassword
//
// @param data Data to be sent to the bot via a callback query
func NewInlineKeyboardButtonTypeCallbackWithPassword(data []byte) *InlineKeyboardButtonTypeCallbackWithPassword {
	inlineKeyboardButtonTypeCallbackWithPasswordTemp := InlineKeyboardButtonTypeCallbackWithPassword{
		tdCommon: tdCommon{Type: "inlineKeyboardButtonTypeCallbackWithPassword"},
		Data:     data,
	}

	return &inlineKeyboardButtonTypeCallbackWithPasswordTemp
}

// GetInlineKeyboardButtonTypeEnum return the enum type of this object
func (inlineKeyboardButtonTypeCallbackWithPassword *InlineKeyboardButtonTypeCallbackWithPassword) GetInlineKeyboardButtonTypeEnum() InlineKeyboardButtonTypeEnum {
	return InlineKeyboardButtonTypeCallbackWithPasswordType
}

// InlineKeyboardButtonTypeCallbackGame A button with a game that sends a callback query to a bot. This button must be in the first column and row of the keyboard and can be attached only to a message with content of the type messageGame
type InlineKeyboardButtonTypeCallbackGame struct {
	tdCommon
}

// MessageType return the string telegram-type of InlineKeyboardButtonTypeCallbackGame
func (inlineKeyboardButtonTypeCallbackGame *InlineKeyboardButtonTypeCallbackGame) MessageType() string {
	return "inlineKeyboardButtonTypeCallbackGame"
}

// NewInlineKeyboardButtonTypeCallbackGame creates a new InlineKeyboardButtonTypeCallbackGame
//
func NewInlineKeyboardButtonTypeCallbackGame() *InlineKeyboardButtonTypeCallbackGame {
	inlineKeyboardButtonTypeCallbackGameTemp := InlineKeyboardButtonTypeCallbackGame{
		tdCommon: tdCommon{Type: "inlineKeyboardButtonTypeCallbackGame"},
	}

	return &inlineKeyboardButtonTypeCallbackGameTemp
}

// GetInlineKeyboardButtonTypeEnum return the enum type of this object
func (inlineKeyboardButtonTypeCallbackGame *InlineKeyboardButtonTypeCallbackGame) GetInlineKeyboardButtonTypeEnum() InlineKeyboardButtonTypeEnum {
	return InlineKeyboardButtonTypeCallbackGameType
}

// InlineKeyboardButtonTypeSwitchInline A button that forces an inline query to the bot to be inserted in the input field
type InlineKeyboardButtonTypeSwitchInline struct {
	tdCommon
	Query         string `json:"query"`           // Inline query to be sent to the bot
	InCurrentChat bool   `json:"in_current_chat"` // True, if the inline query must be sent from the current chat
}

// MessageType return the string telegram-type of InlineKeyboardButtonTypeSwitchInline
func (inlineKeyboardButtonTypeSwitchInline *InlineKeyboardButtonTypeSwitchInline) MessageType() string {
	return "inlineKeyboardButtonTypeSwitchInline"
}

// NewInlineKeyboardButtonTypeSwitchInline creates a new InlineKeyboardButtonTypeSwitchInline
//
// @param query Inline query to be sent to the bot
// @param inCurrentChat True, if the inline query must be sent from the current chat
func NewInlineKeyboardButtonTypeSwitchInline(query string, inCurrentChat bool) *InlineKeyboardButtonTypeSwitchInline {
	inlineKeyboardButtonTypeSwitchInlineTemp := InlineKeyboardButtonTypeSwitchInline{
		tdCommon:      tdCommon{Type: "inlineKeyboardButtonTypeSwitchInline"},
		Query:         query,
		InCurrentChat: inCurrentChat,
	}

	return &inlineKeyboardButtonTypeSwitchInlineTemp
}

// GetInlineKeyboardButtonTypeEnum return the enum type of this object
func (inlineKeyboardButtonTypeSwitchInline *InlineKeyboardButtonTypeSwitchInline) GetInlineKeyboardButtonTypeEnum() InlineKeyboardButtonTypeEnum {
	return InlineKeyboardButtonTypeSwitchInlineType
}

// InlineKeyboardButtonTypeBuy A button to buy something. This button must be in the first column and row of the keyboard and can be attached only to a message with content of the type messageInvoice
type InlineKeyboardButtonTypeBuy struct {
	tdCommon
}

// MessageType return the string telegram-type of InlineKeyboardButtonTypeBuy
func (inlineKeyboardButtonTypeBuy *InlineKeyboardButtonTypeBuy) MessageType() string {
	return "inlineKeyboardButtonTypeBuy"
}

// NewInlineKeyboardButtonTypeBuy creates a new InlineKeyboardButtonTypeBuy
//
func NewInlineKeyboardButtonTypeBuy() *InlineKeyboardButtonTypeBuy {
	inlineKeyboardButtonTypeBuyTemp := InlineKeyboardButtonTypeBuy{
		tdCommon: tdCommon{Type: "inlineKeyboardButtonTypeBuy"},
	}

	return &inlineKeyboardButtonTypeBuyTemp
}

// GetInlineKeyboardButtonTypeEnum return the enum type of this object
func (inlineKeyboardButtonTypeBuy *InlineKeyboardButtonTypeBuy) GetInlineKeyboardButtonTypeEnum() InlineKeyboardButtonTypeEnum {
	return InlineKeyboardButtonTypeBuyType
}

// InlineKeyboardButtonTypeUser A button to open a chat with a user
type InlineKeyboardButtonTypeUser struct {
	tdCommon
	UserId int64 `json:"user_id"` // User identifier
}

// MessageType return the string telegram-type of InlineKeyboardButtonTypeUser
func (inlineKeyboardButtonTypeUser *InlineKeyboardButtonTypeUser) MessageType() string {
	return "inlineKeyboardButtonTypeUser"
}

// NewInlineKeyboardButtonTypeUser creates a new InlineKeyboardButtonTypeUser
//
// @param userId User identifier
func NewInlineKeyboardButtonTypeUser(userId int64) *InlineKeyboardButtonTypeUser {
	inlineKeyboardButtonTypeUserTemp := InlineKeyboardButtonTypeUser{
		tdCommon: tdCommon{Type: "inlineKeyboardButtonTypeUser"},
		UserId:   userId,
	}

	return &inlineKeyboardButtonTypeUserTemp
}

// GetInlineKeyboardButtonTypeEnum return the enum type of this object
func (inlineKeyboardButtonTypeUser *InlineKeyboardButtonTypeUser) GetInlineKeyboardButtonTypeEnum() InlineKeyboardButtonTypeEnum {
	return InlineKeyboardButtonTypeUserType
}
