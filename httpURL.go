package tdlib

import (
	"encoding/json"
	"fmt"
)

// HttpUrl Contains an HTTP URL
type HttpUrl struct {
	tdCommon
	Url string `json:"url"` // The URL
}

// MessageType return the string telegram-type of HttpUrl
func (httpUrl *HttpUrl) MessageType() string {
	return "httpUrl"
}

// NewHttpUrl creates a new HttpUrl
//
// @param url The URL
func NewHttpUrl(url string) *HttpUrl {
	httpUrlTemp := HttpUrl{
		tdCommon: tdCommon{Type: "httpUrl"},
		Url:      url,
	}

	return &httpUrlTemp
}

// GetLoginUrl Returns an HTTP URL which can be used to automatically authorize the user on a website after clicking an inline button of type inlineKeyboardButtonTypeLoginUrl. Use the method getLoginUrlInfo to find whether a prior user confirmation is needed. If an error is returned, then the button must be handled as an ordinary URL button
// @param chatId Chat identifier of the message with the button
// @param messageId Message identifier of the message with the button
// @param buttonId Button identifier
// @param allowWriteAccess True, if the user allowed the bot to send them messages
func (client *Client) GetLoginUrl(chatId int64, messageId int64, buttonId int64, allowWriteAccess bool) (*HttpUrl, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":              "getLoginUrl",
		"chat_id":            chatId,
		"message_id":         messageId,
		"button_id":          buttonId,
		"allow_write_access": allowWriteAccess,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var httpUrl HttpUrl
	err = json.Unmarshal(result.Raw, &httpUrl)
	return &httpUrl, err
}

// GetExternalLink Returns an HTTP URL which can be used to automatically authorize the current user on a website after clicking an HTTP link. Use the method getExternalLinkInfo to find whether a prior user confirmation is needed
// @param link The HTTP link
// @param allowWriteAccess True, if the current user allowed the bot, returned in getExternalLinkInfo, to send them messages
func (client *Client) GetExternalLink(link string, allowWriteAccess bool) (*HttpUrl, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":              "getExternalLink",
		"link":               link,
		"allow_write_access": allowWriteAccess,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var httpUrl HttpUrl
	err = json.Unmarshal(result.Raw, &httpUrl)
	return &httpUrl, err
}

// GetGroupCallInviteLink Returns invite link to a video chat in a public chat
// @param groupCallId Group call identifier
// @param canSelfUnmute Pass true if the invite link needs to contain an invite hash, passing which to joinGroupCall would allow the invited user to unmute themselves. Requires groupCall.can_be_managed group call flag
func (client *Client) GetGroupCallInviteLink(groupCallId int32, canSelfUnmute bool) (*HttpUrl, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":           "getGroupCallInviteLink",
		"group_call_id":   groupCallId,
		"can_self_unmute": canSelfUnmute,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var httpUrl HttpUrl
	err = json.Unmarshal(result.Raw, &httpUrl)
	return &httpUrl, err
}

// GetEmojiSuggestionsUrl Returns an HTTP URL which can be used to automatically log in to the translation platform and suggest new emoji replacements. The URL will be valid for 30 seconds after generation
// @param languageCode Language code for which the emoji replacements will be suggested
func (client *Client) GetEmojiSuggestionsUrl(languageCode string) (*HttpUrl, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":         "getEmojiSuggestionsUrl",
		"language_code": languageCode,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var httpUrl HttpUrl
	err = json.Unmarshal(result.Raw, &httpUrl)
	return &httpUrl, err
}

// GetBackgroundUrl Constructs a persistent HTTP URL for a background
// @param name Background name
// @param typeParam Background type
func (client *Client) GetBackgroundUrl(name string, typeParam BackgroundType) (*HttpUrl, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "getBackgroundUrl",
		"name":  name,
		"type":  typeParam,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var httpUrl HttpUrl
	err = json.Unmarshal(result.Raw, &httpUrl)
	return &httpUrl, err
}

// GetApplicationDownloadLink Returns the link for downloading official Telegram application to be used when the current user invites friends to Telegram
func (client *Client) GetApplicationDownloadLink() (*HttpUrl, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "getApplicationDownloadLink",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var httpUrl HttpUrl
	err = json.Unmarshal(result.Raw, &httpUrl)
	return &httpUrl, err
}

// GetProxyLink Returns an HTTPS link, which can be used to add a proxy. Available only for SOCKS5 and MTProto proxies. Can be called before authorization
// @param proxyId Proxy identifier
func (client *Client) GetProxyLink(proxyId int32) (*HttpUrl, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":    "getProxyLink",
		"proxy_id": proxyId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var httpUrl HttpUrl
	err = json.Unmarshal(result.Raw, &httpUrl)
	return &httpUrl, err
}
