package tdlib

import (
	"encoding/json"
	"fmt"
)

// WebPage Describes a web page preview
type WebPage struct {
	tdCommon
	Url                string         `json:"url"`                  // Original URL of the link
	DisplayUrl         string         `json:"display_url"`          // URL to display
	Type               string         `json:"type"`                 // Type of the web page. Can be: article, photo, audio, video, document, profile, app, or something else
	SiteName           string         `json:"site_name"`            // Short name of the site (e.g., Google Docs, App Store)
	Title              string         `json:"title"`                // Title of the content
	Description        *FormattedText `json:"description"`          // Description of the content
	Photo              *Photo         `json:"photo"`                // Image representing the content; may be null
	EmbedUrl           string         `json:"embed_url"`            // URL to show in the embedded preview
	EmbedType          string         `json:"embed_type"`           // MIME type of the embedded preview, (e.g., text/html or video/mp4)
	EmbedWidth         int32          `json:"embed_width"`          // Width of the embedded preview
	EmbedHeight        int32          `json:"embed_height"`         // Height of the embedded preview
	Duration           int32          `json:"duration"`             // Duration of the content, in seconds
	Author             string         `json:"author"`               // Author of the content
	Animation          *Animation     `json:"animation"`            // Preview of the content as an animation, if available; may be null
	Audio              *Audio         `json:"audio"`                // Preview of the content as an audio file, if available; may be null
	Document           *Document      `json:"document"`             // Preview of the content as a document, if available; may be null
	Sticker            *Sticker       `json:"sticker"`              // Preview of the content as a sticker for small WEBP files, if available; may be null
	Video              *Video         `json:"video"`                // Preview of the content as a video, if available; may be null
	VideoNote          *VideoNote     `json:"video_note"`           // Preview of the content as a video note, if available; may be null
	VoiceNote          *VoiceNote     `json:"voice_note"`           // Preview of the content as a voice note, if available; may be null
	InstantViewVersion int32          `json:"instant_view_version"` // Version of instant view, available for the web page (currently, can be 1 or 2), 0 if none
}

// MessageType return the string telegram-type of WebPage
func (webPage *WebPage) MessageType() string {
	return "webPage"
}

// NewWebPage creates a new WebPage
//
// @param url Original URL of the link
// @param displayUrl URL to display
// @param typeParam Type of the web page. Can be: article, photo, audio, video, document, profile, app, or something else
// @param siteName Short name of the site (e.g., Google Docs, App Store)
// @param title Title of the content
// @param description Description of the content
// @param photo Image representing the content; may be null
// @param embedUrl URL to show in the embedded preview
// @param embedType MIME type of the embedded preview, (e.g., text/html or video/mp4)
// @param embedWidth Width of the embedded preview
// @param embedHeight Height of the embedded preview
// @param duration Duration of the content, in seconds
// @param author Author of the content
// @param animation Preview of the content as an animation, if available; may be null
// @param audio Preview of the content as an audio file, if available; may be null
// @param document Preview of the content as a document, if available; may be null
// @param sticker Preview of the content as a sticker for small WEBP files, if available; may be null
// @param video Preview of the content as a video, if available; may be null
// @param videoNote Preview of the content as a video note, if available; may be null
// @param voiceNote Preview of the content as a voice note, if available; may be null
// @param instantViewVersion Version of instant view, available for the web page (currently, can be 1 or 2), 0 if none
func NewWebPage(url string, displayUrl string, typeParam string, siteName string, title string, description *FormattedText, photo *Photo, embedUrl string, embedType string, embedWidth int32, embedHeight int32, duration int32, author string, animation *Animation, audio *Audio, document *Document, sticker *Sticker, video *Video, videoNote *VideoNote, voiceNote *VoiceNote, instantViewVersion int32) *WebPage {
	webPageTemp := WebPage{
		tdCommon:           tdCommon{Type: "webPage"},
		Url:                url,
		DisplayUrl:         displayUrl,
		Type:               typeParam,
		SiteName:           siteName,
		Title:              title,
		Description:        description,
		Photo:              photo,
		EmbedUrl:           embedUrl,
		EmbedType:          embedType,
		EmbedWidth:         embedWidth,
		EmbedHeight:        embedHeight,
		Duration:           duration,
		Author:             author,
		Animation:          animation,
		Audio:              audio,
		Document:           document,
		Sticker:            sticker,
		Video:              video,
		VideoNote:          videoNote,
		VoiceNote:          voiceNote,
		InstantViewVersion: instantViewVersion,
	}

	return &webPageTemp
}

// GetWebPagePreview Returns a web page preview by the text of the message. Do not call this function too often. Returns a 404 error if the web page has no preview
// @param text Message text with formatting
func (client *Client) GetWebPagePreview(text *FormattedText) (*WebPage, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "getWebPagePreview",
		"text":  text,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var webPage WebPage
	err = json.Unmarshal(result.Raw, &webPage)
	return &webPage, err
}
