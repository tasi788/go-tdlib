package tdlib

import (
	"encoding/json"
	"fmt"
)

// Hashtags Contains a list of hashtags
type Hashtags struct {
	tdCommon
	Hashtags []string `json:"hashtags"` // A list of hashtags
}

// MessageType return the string telegram-type of Hashtags
func (hashtags *Hashtags) MessageType() string {
	return "hashtags"
}

// NewHashtags creates a new Hashtags
//
// @param hashtags A list of hashtags
func NewHashtags(hashtags []string) *Hashtags {
	hashtagsTemp := Hashtags{
		tdCommon: tdCommon{Type: "hashtags"},
		Hashtags: hashtags,
	}

	return &hashtagsTemp
}

// SearchHashtags Searches for recently used hashtags by their prefix
// @param prefix Hashtag prefix to search for
// @param limit The maximum number of hashtags to be returned
func (client *Client) SearchHashtags(prefix string, limit int32) (*Hashtags, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":  "searchHashtags",
		"prefix": prefix,
		"limit":  limit,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var hashtags Hashtags
	err = json.Unmarshal(result.Raw, &hashtags)
	return &hashtags, err
}
