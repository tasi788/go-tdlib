package tdlib

import (
	"encoding/json"
	"fmt"
)

// TMeUrls Contains a list of t.me URLs
type TMeUrls struct {
	tdCommon
	Urls []TMeUrl `json:"urls"` // List of URLs
}

// MessageType return the string telegram-type of TMeUrls
func (tMeUrls *TMeUrls) MessageType() string {
	return "tMeUrls"
}

// NewTMeUrls creates a new TMeUrls
//
// @param urls List of URLs
func NewTMeUrls(urls []TMeUrl) *TMeUrls {
	tMeUrlsTemp := TMeUrls{
		tdCommon: tdCommon{Type: "tMeUrls"},
		Urls:     urls,
	}

	return &tMeUrlsTemp
}

// GetRecentlyVisitedTMeUrls Returns t.me URLs recently visited by a newly registered user
// @param referrer Google Play referrer to identify the user
func (client *Client) GetRecentlyVisitedTMeUrls(referrer string) (*TMeUrls, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":    "getRecentlyVisitedTMeUrls",
		"referrer": referrer,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var tMeUrls TMeUrls
	err = json.Unmarshal(result.Raw, &tMeUrls)
	return &tMeUrls, err
}
