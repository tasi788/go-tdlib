package tdlib

import (
	"encoding/json"
)

// TMeUrl Represents a URL linking to an internal Telegram entity
type TMeUrl struct {
	tdCommon
	Url  string     `json:"url"`  // URL
	Type TMeUrlType `json:"type"` // Type of the URL
}

// MessageType return the string telegram-type of TMeUrl
func (tMeUrl *TMeUrl) MessageType() string {
	return "tMeUrl"
}

// NewTMeUrl creates a new TMeUrl
//
// @param url URL
// @param typeParam Type of the URL
func NewTMeUrl(url string, typeParam TMeUrlType) *TMeUrl {
	tMeUrlTemp := TMeUrl{
		tdCommon: tdCommon{Type: "tMeUrl"},
		Url:      url,
		Type:     typeParam,
	}

	return &tMeUrlTemp
}

// UnmarshalJSON unmarshal to json
func (tMeUrl *TMeUrl) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	tempObj := struct {
		tdCommon
		Url string `json:"url"` // URL

	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	tMeUrl.tdCommon = tempObj.tdCommon
	tMeUrl.Url = tempObj.Url

	fieldType, _ := unmarshalTMeUrlType(objMap["type"])
	tMeUrl.Type = fieldType

	return nil
}
