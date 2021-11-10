package tdlib

import (
	"encoding/json"
	"fmt"
)

// LoginUrlInfo Contains information about an inline button of type inlineKeyboardButtonTypeLoginUrl
type LoginUrlInfo interface {
	GetLoginUrlInfoEnum() LoginUrlInfoEnum
}

// LoginUrlInfoEnum Alias for abstract LoginUrlInfo 'Sub-Classes', used as constant-enum here
type LoginUrlInfoEnum string

// LoginUrlInfo enums
const (
	LoginUrlInfoOpenType                LoginUrlInfoEnum = "loginUrlInfoOpen"
	LoginUrlInfoRequestConfirmationType LoginUrlInfoEnum = "loginUrlInfoRequestConfirmation"
)

func unmarshalLoginUrlInfo(rawMsg *json.RawMessage) (LoginUrlInfo, error) {

	if rawMsg == nil {
		return nil, nil
	}
	var objMap map[string]interface{}
	err := json.Unmarshal(*rawMsg, &objMap)
	if err != nil {
		return nil, err
	}

	switch LoginUrlInfoEnum(objMap["@type"].(string)) {
	case LoginUrlInfoOpenType:
		var loginUrlInfoOpen LoginUrlInfoOpen
		err := json.Unmarshal(*rawMsg, &loginUrlInfoOpen)
		return &loginUrlInfoOpen, err

	case LoginUrlInfoRequestConfirmationType:
		var loginUrlInfoRequestConfirmation LoginUrlInfoRequestConfirmation
		err := json.Unmarshal(*rawMsg, &loginUrlInfoRequestConfirmation)
		return &loginUrlInfoRequestConfirmation, err

	default:
		return nil, fmt.Errorf("Error UnMarshaling, unknown type:" + objMap["@type"].(string))
	}
}

// LoginUrlInfoOpen An HTTP url needs to be open
type LoginUrlInfoOpen struct {
	tdCommon
	Url         string `json:"url"`          // The URL to open
	SkipConfirm bool   `json:"skip_confirm"` // True, if there is no need to show an ordinary open URL confirm
}

// MessageType return the string telegram-type of LoginUrlInfoOpen
func (loginUrlInfoOpen *LoginUrlInfoOpen) MessageType() string {
	return "loginUrlInfoOpen"
}

// NewLoginUrlInfoOpen creates a new LoginUrlInfoOpen
//
// @param url The URL to open
// @param skipConfirm True, if there is no need to show an ordinary open URL confirm
func NewLoginUrlInfoOpen(url string, skipConfirm bool) *LoginUrlInfoOpen {
	loginUrlInfoOpenTemp := LoginUrlInfoOpen{
		tdCommon:    tdCommon{Type: "loginUrlInfoOpen"},
		Url:         url,
		SkipConfirm: skipConfirm,
	}

	return &loginUrlInfoOpenTemp
}

// GetLoginUrlInfoEnum return the enum type of this object
func (loginUrlInfoOpen *LoginUrlInfoOpen) GetLoginUrlInfoEnum() LoginUrlInfoEnum {
	return LoginUrlInfoOpenType
}

// LoginUrlInfoRequestConfirmation An authorization confirmation dialog needs to be shown to the user
type LoginUrlInfoRequestConfirmation struct {
	tdCommon
	Url                string `json:"url"`                  // An HTTP URL to be opened
	Domain             string `json:"domain"`               // A domain of the URL
	BotUserId          int64  `json:"bot_user_id"`          // User identifier of a bot linked with the website
	RequestWriteAccess bool   `json:"request_write_access"` // True, if the user needs to be requested to give the permission to the bot to send them messages
}

// MessageType return the string telegram-type of LoginUrlInfoRequestConfirmation
func (loginUrlInfoRequestConfirmation *LoginUrlInfoRequestConfirmation) MessageType() string {
	return "loginUrlInfoRequestConfirmation"
}

// NewLoginUrlInfoRequestConfirmation creates a new LoginUrlInfoRequestConfirmation
//
// @param url An HTTP URL to be opened
// @param domain A domain of the URL
// @param botUserId User identifier of a bot linked with the website
// @param requestWriteAccess True, if the user needs to be requested to give the permission to the bot to send them messages
func NewLoginUrlInfoRequestConfirmation(url string, domain string, botUserId int64, requestWriteAccess bool) *LoginUrlInfoRequestConfirmation {
	loginUrlInfoRequestConfirmationTemp := LoginUrlInfoRequestConfirmation{
		tdCommon:           tdCommon{Type: "loginUrlInfoRequestConfirmation"},
		Url:                url,
		Domain:             domain,
		BotUserId:          botUserId,
		RequestWriteAccess: requestWriteAccess,
	}

	return &loginUrlInfoRequestConfirmationTemp
}

// GetLoginUrlInfoEnum return the enum type of this object
func (loginUrlInfoRequestConfirmation *LoginUrlInfoRequestConfirmation) GetLoginUrlInfoEnum() LoginUrlInfoEnum {
	return LoginUrlInfoRequestConfirmationType
}

// GetLoginUrlInfo Returns information about a button of type inlineKeyboardButtonTypeLoginUrl. The method needs to be called when the user presses the button
// @param chatId Chat identifier of the message with the button
// @param messageId Message identifier of the message with the button
// @param buttonId Button identifier
func (client *Client) GetLoginUrlInfo(chatId int64, messageId int64, buttonId int64) (LoginUrlInfo, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":      "getLoginUrlInfo",
		"chat_id":    chatId,
		"message_id": messageId,
		"button_id":  buttonId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	switch LoginUrlInfoEnum(result.Data["@type"].(string)) {

	case LoginUrlInfoOpenType:
		var loginUrlInfo LoginUrlInfoOpen
		err = json.Unmarshal(result.Raw, &loginUrlInfo)
		return &loginUrlInfo, err

	case LoginUrlInfoRequestConfirmationType:
		var loginUrlInfo LoginUrlInfoRequestConfirmation
		err = json.Unmarshal(result.Raw, &loginUrlInfo)
		return &loginUrlInfo, err

	default:
		return nil, fmt.Errorf("Invalid type")
	}
}

// GetExternalLinkInfo Returns information about an action to be done when the current user clicks an external link. Don't use this method for links from secret chats if web page preview is disabled in secret chats
// @param link The link
func (client *Client) GetExternalLinkInfo(link string) (LoginUrlInfo, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "getExternalLinkInfo",
		"link":  link,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	switch LoginUrlInfoEnum(result.Data["@type"].(string)) {

	case LoginUrlInfoOpenType:
		var loginUrlInfo LoginUrlInfoOpen
		err = json.Unmarshal(result.Raw, &loginUrlInfo)
		return &loginUrlInfo, err

	case LoginUrlInfoRequestConfirmationType:
		var loginUrlInfo LoginUrlInfoRequestConfirmation
		err = json.Unmarshal(result.Raw, &loginUrlInfo)
		return &loginUrlInfo, err

	default:
		return nil, fmt.Errorf("Invalid type")
	}
}
