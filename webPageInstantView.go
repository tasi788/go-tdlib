package tdlib

import (
	"encoding/json"
	"fmt"
)

// WebPageInstantView Describes an instant view page for a web page
type WebPageInstantView struct {
	tdCommon
	PageBlocks   []PageBlock      `json:"page_blocks"`   // Content of the web page
	ViewCount    int32            `json:"view_count"`    // Number of the instant view views; 0 if unknown
	Version      int32            `json:"version"`       // Version of the instant view; currently, can be 1 or 2
	IsRtl        bool             `json:"is_rtl"`        // True, if the instant view must be shown from right to left
	IsFull       bool             `json:"is_full"`       // True, if the instant view contains the full page. A network request might be needed to get the full web page instant view
	FeedbackLink InternalLinkType `json:"feedback_link"` // An internal link to be opened to leave feedback about the instant view
}

// MessageType return the string telegram-type of WebPageInstantView
func (webPageInstantView *WebPageInstantView) MessageType() string {
	return "webPageInstantView"
}

// NewWebPageInstantView creates a new WebPageInstantView
//
// @param pageBlocks Content of the web page
// @param viewCount Number of the instant view views; 0 if unknown
// @param version Version of the instant view; currently, can be 1 or 2
// @param isRtl True, if the instant view must be shown from right to left
// @param isFull True, if the instant view contains the full page. A network request might be needed to get the full web page instant view
// @param feedbackLink An internal link to be opened to leave feedback about the instant view
func NewWebPageInstantView(pageBlocks []PageBlock, viewCount int32, version int32, isRtl bool, isFull bool, feedbackLink InternalLinkType) *WebPageInstantView {
	webPageInstantViewTemp := WebPageInstantView{
		tdCommon:     tdCommon{Type: "webPageInstantView"},
		PageBlocks:   pageBlocks,
		ViewCount:    viewCount,
		Version:      version,
		IsRtl:        isRtl,
		IsFull:       isFull,
		FeedbackLink: feedbackLink,
	}

	return &webPageInstantViewTemp
}

// UnmarshalJSON unmarshal to json
func (webPageInstantView *WebPageInstantView) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	tempObj := struct {
		tdCommon
		PageBlocks []PageBlock `json:"page_blocks"` // Content of the web page
		ViewCount  int32       `json:"view_count"`  // Number of the instant view views; 0 if unknown
		Version    int32       `json:"version"`     // Version of the instant view; currently, can be 1 or 2
		IsRtl      bool        `json:"is_rtl"`      // True, if the instant view must be shown from right to left
		IsFull     bool        `json:"is_full"`     // True, if the instant view contains the full page. A network request might be needed to get the full web page instant view

	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	webPageInstantView.tdCommon = tempObj.tdCommon
	webPageInstantView.PageBlocks = tempObj.PageBlocks
	webPageInstantView.ViewCount = tempObj.ViewCount
	webPageInstantView.Version = tempObj.Version
	webPageInstantView.IsRtl = tempObj.IsRtl
	webPageInstantView.IsFull = tempObj.IsFull

	fieldFeedbackLink, _ := unmarshalInternalLinkType(objMap["feedback_link"])
	webPageInstantView.FeedbackLink = fieldFeedbackLink

	return nil
}

// GetWebPageInstantView Returns an instant view version of a web page if available. Returns a 404 error if the web page has no instant view page
// @param url The web page URL
// @param forceFull If true, the full instant view for the web page will be returned
func (client *Client) GetWebPageInstantView(url string, forceFull bool) (*WebPageInstantView, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":      "getWebPageInstantView",
		"url":        url,
		"force_full": forceFull,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var webPageInstantView WebPageInstantView
	err = json.Unmarshal(result.Raw, &webPageInstantView)
	return &webPageInstantView, err
}
