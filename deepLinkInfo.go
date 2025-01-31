package tdlib

import (
	"encoding/json"
	"fmt"
)

// DeepLinkInfo Contains information about a tg: deep link
type DeepLinkInfo struct {
	tdCommon
	Text                  *FormattedText `json:"text"`                    // Text to be shown to the user
	NeedUpdateApplication bool           `json:"need_update_application"` // True, if the user must be asked to update the application
}

// MessageType return the string telegram-type of DeepLinkInfo
func (deepLinkInfo *DeepLinkInfo) MessageType() string {
	return "deepLinkInfo"
}

// NewDeepLinkInfo creates a new DeepLinkInfo
//
// @param text Text to be shown to the user
// @param needUpdateApplication True, if the user must be asked to update the application
func NewDeepLinkInfo(text *FormattedText, needUpdateApplication bool) *DeepLinkInfo {
	deepLinkInfoTemp := DeepLinkInfo{
		tdCommon:              tdCommon{Type: "deepLinkInfo"},
		Text:                  text,
		NeedUpdateApplication: needUpdateApplication,
	}

	return &deepLinkInfoTemp
}

// GetDeepLinkInfo Returns information about a tg:// deep link. Use "tg://need_update_for_some_feature" or "tg:some_unsupported_feature" for testing. Returns a 404 error for unknown links. Can be called before authorization
// @param link The link
func (client *Client) GetDeepLinkInfo(link string) (*DeepLinkInfo, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "getDeepLinkInfo",
		"link":  link,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var deepLinkInfo DeepLinkInfo
	err = json.Unmarshal(result.Raw, &deepLinkInfo)
	return &deepLinkInfo, err
}
