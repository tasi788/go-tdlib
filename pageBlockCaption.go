package tdlib

import (
	"encoding/json"
)

// PageBlockCaption Contains a caption of an instant view web page block, consisting of a text and a trailing credit
type PageBlockCaption struct {
	tdCommon
	Text   RichText `json:"text"`   // Content of the caption
	Credit RichText `json:"credit"` // Block credit (like HTML tag <cite>)
}

// MessageType return the string telegram-type of PageBlockCaption
func (pageBlockCaption *PageBlockCaption) MessageType() string {
	return "pageBlockCaption"
}

// NewPageBlockCaption creates a new PageBlockCaption
//
// @param text Content of the caption
// @param credit Block credit (like HTML tag <cite>)
func NewPageBlockCaption(text RichText, credit RichText) *PageBlockCaption {
	pageBlockCaptionTemp := PageBlockCaption{
		tdCommon: tdCommon{Type: "pageBlockCaption"},
		Text:     text,
		Credit:   credit,
	}

	return &pageBlockCaptionTemp
}

// UnmarshalJSON unmarshal to json
func (pageBlockCaption *PageBlockCaption) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	tempObj := struct {
		tdCommon
	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	pageBlockCaption.tdCommon = tempObj.tdCommon

	fieldText, _ := unmarshalRichText(objMap["text"])
	pageBlockCaption.Text = fieldText

	fieldCredit, _ := unmarshalRichText(objMap["credit"])
	pageBlockCaption.Credit = fieldCredit

	return nil
}
