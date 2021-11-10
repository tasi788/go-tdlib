package tdlib

import (
	"encoding/json"
	"fmt"
)

// InputInlineQueryResult Represents a single result of an inline query; for bots only
type InputInlineQueryResult interface {
	GetInputInlineQueryResultEnum() InputInlineQueryResultEnum
}

// InputInlineQueryResultEnum Alias for abstract InputInlineQueryResult 'Sub-Classes', used as constant-enum here
type InputInlineQueryResultEnum string

// InputInlineQueryResult enums
const (
	InputInlineQueryResultAnimationType InputInlineQueryResultEnum = "inputInlineQueryResultAnimation"
	InputInlineQueryResultArticleType   InputInlineQueryResultEnum = "inputInlineQueryResultArticle"
	InputInlineQueryResultAudioType     InputInlineQueryResultEnum = "inputInlineQueryResultAudio"
	InputInlineQueryResultContactType   InputInlineQueryResultEnum = "inputInlineQueryResultContact"
	InputInlineQueryResultDocumentType  InputInlineQueryResultEnum = "inputInlineQueryResultDocument"
	InputInlineQueryResultGameType      InputInlineQueryResultEnum = "inputInlineQueryResultGame"
	InputInlineQueryResultLocationType  InputInlineQueryResultEnum = "inputInlineQueryResultLocation"
	InputInlineQueryResultPhotoType     InputInlineQueryResultEnum = "inputInlineQueryResultPhoto"
	InputInlineQueryResultStickerType   InputInlineQueryResultEnum = "inputInlineQueryResultSticker"
	InputInlineQueryResultVenueType     InputInlineQueryResultEnum = "inputInlineQueryResultVenue"
	InputInlineQueryResultVideoType     InputInlineQueryResultEnum = "inputInlineQueryResultVideo"
	InputInlineQueryResultVoiceNoteType InputInlineQueryResultEnum = "inputInlineQueryResultVoiceNote"
)

func unmarshalInputInlineQueryResult(rawMsg *json.RawMessage) (InputInlineQueryResult, error) {

	if rawMsg == nil {
		return nil, nil
	}
	var objMap map[string]interface{}
	err := json.Unmarshal(*rawMsg, &objMap)
	if err != nil {
		return nil, err
	}

	switch InputInlineQueryResultEnum(objMap["@type"].(string)) {
	case InputInlineQueryResultAnimationType:
		var inputInlineQueryResultAnimation InputInlineQueryResultAnimation
		err := json.Unmarshal(*rawMsg, &inputInlineQueryResultAnimation)
		return &inputInlineQueryResultAnimation, err

	case InputInlineQueryResultArticleType:
		var inputInlineQueryResultArticle InputInlineQueryResultArticle
		err := json.Unmarshal(*rawMsg, &inputInlineQueryResultArticle)
		return &inputInlineQueryResultArticle, err

	case InputInlineQueryResultAudioType:
		var inputInlineQueryResultAudio InputInlineQueryResultAudio
		err := json.Unmarshal(*rawMsg, &inputInlineQueryResultAudio)
		return &inputInlineQueryResultAudio, err

	case InputInlineQueryResultContactType:
		var inputInlineQueryResultContact InputInlineQueryResultContact
		err := json.Unmarshal(*rawMsg, &inputInlineQueryResultContact)
		return &inputInlineQueryResultContact, err

	case InputInlineQueryResultDocumentType:
		var inputInlineQueryResultDocument InputInlineQueryResultDocument
		err := json.Unmarshal(*rawMsg, &inputInlineQueryResultDocument)
		return &inputInlineQueryResultDocument, err

	case InputInlineQueryResultGameType:
		var inputInlineQueryResultGame InputInlineQueryResultGame
		err := json.Unmarshal(*rawMsg, &inputInlineQueryResultGame)
		return &inputInlineQueryResultGame, err

	case InputInlineQueryResultLocationType:
		var inputInlineQueryResultLocation InputInlineQueryResultLocation
		err := json.Unmarshal(*rawMsg, &inputInlineQueryResultLocation)
		return &inputInlineQueryResultLocation, err

	case InputInlineQueryResultPhotoType:
		var inputInlineQueryResultPhoto InputInlineQueryResultPhoto
		err := json.Unmarshal(*rawMsg, &inputInlineQueryResultPhoto)
		return &inputInlineQueryResultPhoto, err

	case InputInlineQueryResultStickerType:
		var inputInlineQueryResultSticker InputInlineQueryResultSticker
		err := json.Unmarshal(*rawMsg, &inputInlineQueryResultSticker)
		return &inputInlineQueryResultSticker, err

	case InputInlineQueryResultVenueType:
		var inputInlineQueryResultVenue InputInlineQueryResultVenue
		err := json.Unmarshal(*rawMsg, &inputInlineQueryResultVenue)
		return &inputInlineQueryResultVenue, err

	case InputInlineQueryResultVideoType:
		var inputInlineQueryResultVideo InputInlineQueryResultVideo
		err := json.Unmarshal(*rawMsg, &inputInlineQueryResultVideo)
		return &inputInlineQueryResultVideo, err

	case InputInlineQueryResultVoiceNoteType:
		var inputInlineQueryResultVoiceNote InputInlineQueryResultVoiceNote
		err := json.Unmarshal(*rawMsg, &inputInlineQueryResultVoiceNote)
		return &inputInlineQueryResultVoiceNote, err

	default:
		return nil, fmt.Errorf("Error UnMarshaling, unknown type:" + objMap["@type"].(string))
	}
}

// InputInlineQueryResultAnimation Represents a link to an animated GIF or an animated (i.e., without sound) H.264/MPEG-4 AVC video
type InputInlineQueryResultAnimation struct {
	tdCommon
	Id                  string              `json:"id"`                    // Unique identifier of the query result
	Title               string              `json:"title"`                 // Title of the query result
	ThumbnailUrl        string              `json:"thumbnail_url"`         // URL of the result thumbnail (JPEG, GIF, or MPEG4), if it exists
	ThumbnailMimeType   string              `json:"thumbnail_mime_type"`   // MIME type of the video thumbnail. If non-empty, must be one of "image/jpeg", "image/gif" and "video/mp4"
	VideoUrl            string              `json:"video_url"`             // The URL of the video file (file size must not exceed 1MB)
	VideoMimeType       string              `json:"video_mime_type"`       // MIME type of the video file. Must be one of "image/gif" and "video/mp4"
	VideoDuration       int32               `json:"video_duration"`        // Duration of the video, in seconds
	VideoWidth          int32               `json:"video_width"`           // Width of the video
	VideoHeight         int32               `json:"video_height"`          // Height of the video
	ReplyMarkup         ReplyMarkup         `json:"reply_markup"`          // The message reply markup; pass null if none. Must be of type replyMarkupInlineKeyboard or null
	InputMessageContent InputMessageContent `json:"input_message_content"` // The content of the message to be sent. Must be one of the following types: inputMessageText, inputMessageAnimation, inputMessageInvoice, inputMessageLocation, inputMessageVenue or inputMessageContact
}

// MessageType return the string telegram-type of InputInlineQueryResultAnimation
func (inputInlineQueryResultAnimation *InputInlineQueryResultAnimation) MessageType() string {
	return "inputInlineQueryResultAnimation"
}

// NewInputInlineQueryResultAnimation creates a new InputInlineQueryResultAnimation
//
// @param id Unique identifier of the query result
// @param title Title of the query result
// @param thumbnailUrl URL of the result thumbnail (JPEG, GIF, or MPEG4), if it exists
// @param thumbnailMimeType MIME type of the video thumbnail. If non-empty, must be one of "image/jpeg", "image/gif" and "video/mp4"
// @param videoUrl The URL of the video file (file size must not exceed 1MB)
// @param videoMimeType MIME type of the video file. Must be one of "image/gif" and "video/mp4"
// @param videoDuration Duration of the video, in seconds
// @param videoWidth Width of the video
// @param videoHeight Height of the video
// @param replyMarkup The message reply markup; pass null if none. Must be of type replyMarkupInlineKeyboard or null
// @param inputMessageContent The content of the message to be sent. Must be one of the following types: inputMessageText, inputMessageAnimation, inputMessageInvoice, inputMessageLocation, inputMessageVenue or inputMessageContact
func NewInputInlineQueryResultAnimation(id string, title string, thumbnailUrl string, thumbnailMimeType string, videoUrl string, videoMimeType string, videoDuration int32, videoWidth int32, videoHeight int32, replyMarkup ReplyMarkup, inputMessageContent InputMessageContent) *InputInlineQueryResultAnimation {
	inputInlineQueryResultAnimationTemp := InputInlineQueryResultAnimation{
		tdCommon:            tdCommon{Type: "inputInlineQueryResultAnimation"},
		Id:                  id,
		Title:               title,
		ThumbnailUrl:        thumbnailUrl,
		ThumbnailMimeType:   thumbnailMimeType,
		VideoUrl:            videoUrl,
		VideoMimeType:       videoMimeType,
		VideoDuration:       videoDuration,
		VideoWidth:          videoWidth,
		VideoHeight:         videoHeight,
		ReplyMarkup:         replyMarkup,
		InputMessageContent: inputMessageContent,
	}

	return &inputInlineQueryResultAnimationTemp
}

// UnmarshalJSON unmarshal to json
func (inputInlineQueryResultAnimation *InputInlineQueryResultAnimation) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	tempObj := struct {
		tdCommon
		Id                string `json:"id"`                  // Unique identifier of the query result
		Title             string `json:"title"`               // Title of the query result
		ThumbnailUrl      string `json:"thumbnail_url"`       // URL of the result thumbnail (JPEG, GIF, or MPEG4), if it exists
		ThumbnailMimeType string `json:"thumbnail_mime_type"` // MIME type of the video thumbnail. If non-empty, must be one of "image/jpeg", "image/gif" and "video/mp4"
		VideoUrl          string `json:"video_url"`           // The URL of the video file (file size must not exceed 1MB)
		VideoMimeType     string `json:"video_mime_type"`     // MIME type of the video file. Must be one of "image/gif" and "video/mp4"
		VideoDuration     int32  `json:"video_duration"`      // Duration of the video, in seconds
		VideoWidth        int32  `json:"video_width"`         // Width of the video
		VideoHeight       int32  `json:"video_height"`        // Height of the video

	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	inputInlineQueryResultAnimation.tdCommon = tempObj.tdCommon
	inputInlineQueryResultAnimation.Id = tempObj.Id
	inputInlineQueryResultAnimation.Title = tempObj.Title
	inputInlineQueryResultAnimation.ThumbnailUrl = tempObj.ThumbnailUrl
	inputInlineQueryResultAnimation.ThumbnailMimeType = tempObj.ThumbnailMimeType
	inputInlineQueryResultAnimation.VideoUrl = tempObj.VideoUrl
	inputInlineQueryResultAnimation.VideoMimeType = tempObj.VideoMimeType
	inputInlineQueryResultAnimation.VideoDuration = tempObj.VideoDuration
	inputInlineQueryResultAnimation.VideoWidth = tempObj.VideoWidth
	inputInlineQueryResultAnimation.VideoHeight = tempObj.VideoHeight

	fieldReplyMarkup, _ := unmarshalReplyMarkup(objMap["reply_markup"])
	inputInlineQueryResultAnimation.ReplyMarkup = fieldReplyMarkup

	fieldInputMessageContent, _ := unmarshalInputMessageContent(objMap["input_message_content"])
	inputInlineQueryResultAnimation.InputMessageContent = fieldInputMessageContent

	return nil
}

// GetInputInlineQueryResultEnum return the enum type of this object
func (inputInlineQueryResultAnimation *InputInlineQueryResultAnimation) GetInputInlineQueryResultEnum() InputInlineQueryResultEnum {
	return InputInlineQueryResultAnimationType
}

// InputInlineQueryResultArticle Represents a link to an article or web page
type InputInlineQueryResultArticle struct {
	tdCommon
	Id                  string              `json:"id"`                    // Unique identifier of the query result
	Url                 string              `json:"url"`                   // URL of the result, if it exists
	HideUrl             bool                `json:"hide_url"`              // True, if the URL must be not shown
	Title               string              `json:"title"`                 // Title of the result
	Description         string              `json:"description"`           // A short description of the result
	ThumbnailUrl        string              `json:"thumbnail_url"`         // URL of the result thumbnail, if it exists
	ThumbnailWidth      int32               `json:"thumbnail_width"`       // Thumbnail width, if known
	ThumbnailHeight     int32               `json:"thumbnail_height"`      // Thumbnail height, if known
	ReplyMarkup         ReplyMarkup         `json:"reply_markup"`          // The message reply markup; pass null if none. Must be of type replyMarkupInlineKeyboard or null
	InputMessageContent InputMessageContent `json:"input_message_content"` // The content of the message to be sent. Must be one of the following types: inputMessageText, inputMessageInvoice, inputMessageLocation, inputMessageVenue or inputMessageContact
}

// MessageType return the string telegram-type of InputInlineQueryResultArticle
func (inputInlineQueryResultArticle *InputInlineQueryResultArticle) MessageType() string {
	return "inputInlineQueryResultArticle"
}

// NewInputInlineQueryResultArticle creates a new InputInlineQueryResultArticle
//
// @param id Unique identifier of the query result
// @param url URL of the result, if it exists
// @param hideUrl True, if the URL must be not shown
// @param title Title of the result
// @param description A short description of the result
// @param thumbnailUrl URL of the result thumbnail, if it exists
// @param thumbnailWidth Thumbnail width, if known
// @param thumbnailHeight Thumbnail height, if known
// @param replyMarkup The message reply markup; pass null if none. Must be of type replyMarkupInlineKeyboard or null
// @param inputMessageContent The content of the message to be sent. Must be one of the following types: inputMessageText, inputMessageInvoice, inputMessageLocation, inputMessageVenue or inputMessageContact
func NewInputInlineQueryResultArticle(id string, url string, hideUrl bool, title string, description string, thumbnailUrl string, thumbnailWidth int32, thumbnailHeight int32, replyMarkup ReplyMarkup, inputMessageContent InputMessageContent) *InputInlineQueryResultArticle {
	inputInlineQueryResultArticleTemp := InputInlineQueryResultArticle{
		tdCommon:            tdCommon{Type: "inputInlineQueryResultArticle"},
		Id:                  id,
		Url:                 url,
		HideUrl:             hideUrl,
		Title:               title,
		Description:         description,
		ThumbnailUrl:        thumbnailUrl,
		ThumbnailWidth:      thumbnailWidth,
		ThumbnailHeight:     thumbnailHeight,
		ReplyMarkup:         replyMarkup,
		InputMessageContent: inputMessageContent,
	}

	return &inputInlineQueryResultArticleTemp
}

// UnmarshalJSON unmarshal to json
func (inputInlineQueryResultArticle *InputInlineQueryResultArticle) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	tempObj := struct {
		tdCommon
		Id              string `json:"id"`               // Unique identifier of the query result
		Url             string `json:"url"`              // URL of the result, if it exists
		HideUrl         bool   `json:"hide_url"`         // True, if the URL must be not shown
		Title           string `json:"title"`            // Title of the result
		Description     string `json:"description"`      // A short description of the result
		ThumbnailUrl    string `json:"thumbnail_url"`    // URL of the result thumbnail, if it exists
		ThumbnailWidth  int32  `json:"thumbnail_width"`  // Thumbnail width, if known
		ThumbnailHeight int32  `json:"thumbnail_height"` // Thumbnail height, if known

	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	inputInlineQueryResultArticle.tdCommon = tempObj.tdCommon
	inputInlineQueryResultArticle.Id = tempObj.Id
	inputInlineQueryResultArticle.Url = tempObj.Url
	inputInlineQueryResultArticle.HideUrl = tempObj.HideUrl
	inputInlineQueryResultArticle.Title = tempObj.Title
	inputInlineQueryResultArticle.Description = tempObj.Description
	inputInlineQueryResultArticle.ThumbnailUrl = tempObj.ThumbnailUrl
	inputInlineQueryResultArticle.ThumbnailWidth = tempObj.ThumbnailWidth
	inputInlineQueryResultArticle.ThumbnailHeight = tempObj.ThumbnailHeight

	fieldReplyMarkup, _ := unmarshalReplyMarkup(objMap["reply_markup"])
	inputInlineQueryResultArticle.ReplyMarkup = fieldReplyMarkup

	fieldInputMessageContent, _ := unmarshalInputMessageContent(objMap["input_message_content"])
	inputInlineQueryResultArticle.InputMessageContent = fieldInputMessageContent

	return nil
}

// GetInputInlineQueryResultEnum return the enum type of this object
func (inputInlineQueryResultArticle *InputInlineQueryResultArticle) GetInputInlineQueryResultEnum() InputInlineQueryResultEnum {
	return InputInlineQueryResultArticleType
}

// InputInlineQueryResultAudio Represents a link to an MP3 audio file
type InputInlineQueryResultAudio struct {
	tdCommon
	Id                  string              `json:"id"`                    // Unique identifier of the query result
	Title               string              `json:"title"`                 // Title of the audio file
	Performer           string              `json:"performer"`             // Performer of the audio file
	AudioUrl            string              `json:"audio_url"`             // The URL of the audio file
	AudioDuration       int32               `json:"audio_duration"`        // Audio file duration, in seconds
	ReplyMarkup         ReplyMarkup         `json:"reply_markup"`          // The message reply markup; pass null if none. Must be of type replyMarkupInlineKeyboard or null
	InputMessageContent InputMessageContent `json:"input_message_content"` // The content of the message to be sent. Must be one of the following types: inputMessageText, inputMessageAudio, inputMessageInvoice, inputMessageLocation, inputMessageVenue or inputMessageContact
}

// MessageType return the string telegram-type of InputInlineQueryResultAudio
func (inputInlineQueryResultAudio *InputInlineQueryResultAudio) MessageType() string {
	return "inputInlineQueryResultAudio"
}

// NewInputInlineQueryResultAudio creates a new InputInlineQueryResultAudio
//
// @param id Unique identifier of the query result
// @param title Title of the audio file
// @param performer Performer of the audio file
// @param audioUrl The URL of the audio file
// @param audioDuration Audio file duration, in seconds
// @param replyMarkup The message reply markup; pass null if none. Must be of type replyMarkupInlineKeyboard or null
// @param inputMessageContent The content of the message to be sent. Must be one of the following types: inputMessageText, inputMessageAudio, inputMessageInvoice, inputMessageLocation, inputMessageVenue or inputMessageContact
func NewInputInlineQueryResultAudio(id string, title string, performer string, audioUrl string, audioDuration int32, replyMarkup ReplyMarkup, inputMessageContent InputMessageContent) *InputInlineQueryResultAudio {
	inputInlineQueryResultAudioTemp := InputInlineQueryResultAudio{
		tdCommon:            tdCommon{Type: "inputInlineQueryResultAudio"},
		Id:                  id,
		Title:               title,
		Performer:           performer,
		AudioUrl:            audioUrl,
		AudioDuration:       audioDuration,
		ReplyMarkup:         replyMarkup,
		InputMessageContent: inputMessageContent,
	}

	return &inputInlineQueryResultAudioTemp
}

// UnmarshalJSON unmarshal to json
func (inputInlineQueryResultAudio *InputInlineQueryResultAudio) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	tempObj := struct {
		tdCommon
		Id            string `json:"id"`             // Unique identifier of the query result
		Title         string `json:"title"`          // Title of the audio file
		Performer     string `json:"performer"`      // Performer of the audio file
		AudioUrl      string `json:"audio_url"`      // The URL of the audio file
		AudioDuration int32  `json:"audio_duration"` // Audio file duration, in seconds

	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	inputInlineQueryResultAudio.tdCommon = tempObj.tdCommon
	inputInlineQueryResultAudio.Id = tempObj.Id
	inputInlineQueryResultAudio.Title = tempObj.Title
	inputInlineQueryResultAudio.Performer = tempObj.Performer
	inputInlineQueryResultAudio.AudioUrl = tempObj.AudioUrl
	inputInlineQueryResultAudio.AudioDuration = tempObj.AudioDuration

	fieldReplyMarkup, _ := unmarshalReplyMarkup(objMap["reply_markup"])
	inputInlineQueryResultAudio.ReplyMarkup = fieldReplyMarkup

	fieldInputMessageContent, _ := unmarshalInputMessageContent(objMap["input_message_content"])
	inputInlineQueryResultAudio.InputMessageContent = fieldInputMessageContent

	return nil
}

// GetInputInlineQueryResultEnum return the enum type of this object
func (inputInlineQueryResultAudio *InputInlineQueryResultAudio) GetInputInlineQueryResultEnum() InputInlineQueryResultEnum {
	return InputInlineQueryResultAudioType
}

// InputInlineQueryResultContact Represents a user contact
type InputInlineQueryResultContact struct {
	tdCommon
	Id                  string              `json:"id"`                    // Unique identifier of the query result
	Contact             *Contact            `json:"contact"`               // User contact
	ThumbnailUrl        string              `json:"thumbnail_url"`         // URL of the result thumbnail, if it exists
	ThumbnailWidth      int32               `json:"thumbnail_width"`       // Thumbnail width, if known
	ThumbnailHeight     int32               `json:"thumbnail_height"`      // Thumbnail height, if known
	ReplyMarkup         ReplyMarkup         `json:"reply_markup"`          // The message reply markup; pass null if none. Must be of type replyMarkupInlineKeyboard or null
	InputMessageContent InputMessageContent `json:"input_message_content"` // The content of the message to be sent. Must be one of the following types: inputMessageText, inputMessageInvoice, inputMessageLocation, inputMessageVenue or inputMessageContact
}

// MessageType return the string telegram-type of InputInlineQueryResultContact
func (inputInlineQueryResultContact *InputInlineQueryResultContact) MessageType() string {
	return "inputInlineQueryResultContact"
}

// NewInputInlineQueryResultContact creates a new InputInlineQueryResultContact
//
// @param id Unique identifier of the query result
// @param contact User contact
// @param thumbnailUrl URL of the result thumbnail, if it exists
// @param thumbnailWidth Thumbnail width, if known
// @param thumbnailHeight Thumbnail height, if known
// @param replyMarkup The message reply markup; pass null if none. Must be of type replyMarkupInlineKeyboard or null
// @param inputMessageContent The content of the message to be sent. Must be one of the following types: inputMessageText, inputMessageInvoice, inputMessageLocation, inputMessageVenue or inputMessageContact
func NewInputInlineQueryResultContact(id string, contact *Contact, thumbnailUrl string, thumbnailWidth int32, thumbnailHeight int32, replyMarkup ReplyMarkup, inputMessageContent InputMessageContent) *InputInlineQueryResultContact {
	inputInlineQueryResultContactTemp := InputInlineQueryResultContact{
		tdCommon:            tdCommon{Type: "inputInlineQueryResultContact"},
		Id:                  id,
		Contact:             contact,
		ThumbnailUrl:        thumbnailUrl,
		ThumbnailWidth:      thumbnailWidth,
		ThumbnailHeight:     thumbnailHeight,
		ReplyMarkup:         replyMarkup,
		InputMessageContent: inputMessageContent,
	}

	return &inputInlineQueryResultContactTemp
}

// UnmarshalJSON unmarshal to json
func (inputInlineQueryResultContact *InputInlineQueryResultContact) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	tempObj := struct {
		tdCommon
		Id              string   `json:"id"`               // Unique identifier of the query result
		Contact         *Contact `json:"contact"`          // User contact
		ThumbnailUrl    string   `json:"thumbnail_url"`    // URL of the result thumbnail, if it exists
		ThumbnailWidth  int32    `json:"thumbnail_width"`  // Thumbnail width, if known
		ThumbnailHeight int32    `json:"thumbnail_height"` // Thumbnail height, if known

	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	inputInlineQueryResultContact.tdCommon = tempObj.tdCommon
	inputInlineQueryResultContact.Id = tempObj.Id
	inputInlineQueryResultContact.Contact = tempObj.Contact
	inputInlineQueryResultContact.ThumbnailUrl = tempObj.ThumbnailUrl
	inputInlineQueryResultContact.ThumbnailWidth = tempObj.ThumbnailWidth
	inputInlineQueryResultContact.ThumbnailHeight = tempObj.ThumbnailHeight

	fieldReplyMarkup, _ := unmarshalReplyMarkup(objMap["reply_markup"])
	inputInlineQueryResultContact.ReplyMarkup = fieldReplyMarkup

	fieldInputMessageContent, _ := unmarshalInputMessageContent(objMap["input_message_content"])
	inputInlineQueryResultContact.InputMessageContent = fieldInputMessageContent

	return nil
}

// GetInputInlineQueryResultEnum return the enum type of this object
func (inputInlineQueryResultContact *InputInlineQueryResultContact) GetInputInlineQueryResultEnum() InputInlineQueryResultEnum {
	return InputInlineQueryResultContactType
}

// InputInlineQueryResultDocument Represents a link to a file
type InputInlineQueryResultDocument struct {
	tdCommon
	Id                  string              `json:"id"`                    // Unique identifier of the query result
	Title               string              `json:"title"`                 // Title of the resulting file
	Description         string              `json:"description"`           // Short description of the result, if known
	DocumentUrl         string              `json:"document_url"`          // URL of the file
	MimeType            string              `json:"mime_type"`             // MIME type of the file content; only "application/pdf" and "application/zip" are currently allowed
	ThumbnailUrl        string              `json:"thumbnail_url"`         // The URL of the file thumbnail, if it exists
	ThumbnailWidth      int32               `json:"thumbnail_width"`       // Width of the thumbnail
	ThumbnailHeight     int32               `json:"thumbnail_height"`      // Height of the thumbnail
	ReplyMarkup         ReplyMarkup         `json:"reply_markup"`          // The message reply markup; pass null if none. Must be of type replyMarkupInlineKeyboard or null
	InputMessageContent InputMessageContent `json:"input_message_content"` // The content of the message to be sent. Must be one of the following types: inputMessageText, inputMessageDocument, inputMessageInvoice, inputMessageLocation, inputMessageVenue or inputMessageContact
}

// MessageType return the string telegram-type of InputInlineQueryResultDocument
func (inputInlineQueryResultDocument *InputInlineQueryResultDocument) MessageType() string {
	return "inputInlineQueryResultDocument"
}

// NewInputInlineQueryResultDocument creates a new InputInlineQueryResultDocument
//
// @param id Unique identifier of the query result
// @param title Title of the resulting file
// @param description Short description of the result, if known
// @param documentUrl URL of the file
// @param mimeType MIME type of the file content; only "application/pdf" and "application/zip" are currently allowed
// @param thumbnailUrl The URL of the file thumbnail, if it exists
// @param thumbnailWidth Width of the thumbnail
// @param thumbnailHeight Height of the thumbnail
// @param replyMarkup The message reply markup; pass null if none. Must be of type replyMarkupInlineKeyboard or null
// @param inputMessageContent The content of the message to be sent. Must be one of the following types: inputMessageText, inputMessageDocument, inputMessageInvoice, inputMessageLocation, inputMessageVenue or inputMessageContact
func NewInputInlineQueryResultDocument(id string, title string, description string, documentUrl string, mimeType string, thumbnailUrl string, thumbnailWidth int32, thumbnailHeight int32, replyMarkup ReplyMarkup, inputMessageContent InputMessageContent) *InputInlineQueryResultDocument {
	inputInlineQueryResultDocumentTemp := InputInlineQueryResultDocument{
		tdCommon:            tdCommon{Type: "inputInlineQueryResultDocument"},
		Id:                  id,
		Title:               title,
		Description:         description,
		DocumentUrl:         documentUrl,
		MimeType:            mimeType,
		ThumbnailUrl:        thumbnailUrl,
		ThumbnailWidth:      thumbnailWidth,
		ThumbnailHeight:     thumbnailHeight,
		ReplyMarkup:         replyMarkup,
		InputMessageContent: inputMessageContent,
	}

	return &inputInlineQueryResultDocumentTemp
}

// UnmarshalJSON unmarshal to json
func (inputInlineQueryResultDocument *InputInlineQueryResultDocument) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	tempObj := struct {
		tdCommon
		Id              string `json:"id"`               // Unique identifier of the query result
		Title           string `json:"title"`            // Title of the resulting file
		Description     string `json:"description"`      // Short description of the result, if known
		DocumentUrl     string `json:"document_url"`     // URL of the file
		MimeType        string `json:"mime_type"`        // MIME type of the file content; only "application/pdf" and "application/zip" are currently allowed
		ThumbnailUrl    string `json:"thumbnail_url"`    // The URL of the file thumbnail, if it exists
		ThumbnailWidth  int32  `json:"thumbnail_width"`  // Width of the thumbnail
		ThumbnailHeight int32  `json:"thumbnail_height"` // Height of the thumbnail

	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	inputInlineQueryResultDocument.tdCommon = tempObj.tdCommon
	inputInlineQueryResultDocument.Id = tempObj.Id
	inputInlineQueryResultDocument.Title = tempObj.Title
	inputInlineQueryResultDocument.Description = tempObj.Description
	inputInlineQueryResultDocument.DocumentUrl = tempObj.DocumentUrl
	inputInlineQueryResultDocument.MimeType = tempObj.MimeType
	inputInlineQueryResultDocument.ThumbnailUrl = tempObj.ThumbnailUrl
	inputInlineQueryResultDocument.ThumbnailWidth = tempObj.ThumbnailWidth
	inputInlineQueryResultDocument.ThumbnailHeight = tempObj.ThumbnailHeight

	fieldReplyMarkup, _ := unmarshalReplyMarkup(objMap["reply_markup"])
	inputInlineQueryResultDocument.ReplyMarkup = fieldReplyMarkup

	fieldInputMessageContent, _ := unmarshalInputMessageContent(objMap["input_message_content"])
	inputInlineQueryResultDocument.InputMessageContent = fieldInputMessageContent

	return nil
}

// GetInputInlineQueryResultEnum return the enum type of this object
func (inputInlineQueryResultDocument *InputInlineQueryResultDocument) GetInputInlineQueryResultEnum() InputInlineQueryResultEnum {
	return InputInlineQueryResultDocumentType
}

// InputInlineQueryResultGame Represents a game
type InputInlineQueryResultGame struct {
	tdCommon
	Id            string      `json:"id"`              // Unique identifier of the query result
	GameShortName string      `json:"game_short_name"` // Short name of the game
	ReplyMarkup   ReplyMarkup `json:"reply_markup"`    // The message reply markup; pass null if none. Must be of type replyMarkupInlineKeyboard or null
}

// MessageType return the string telegram-type of InputInlineQueryResultGame
func (inputInlineQueryResultGame *InputInlineQueryResultGame) MessageType() string {
	return "inputInlineQueryResultGame"
}

// NewInputInlineQueryResultGame creates a new InputInlineQueryResultGame
//
// @param id Unique identifier of the query result
// @param gameShortName Short name of the game
// @param replyMarkup The message reply markup; pass null if none. Must be of type replyMarkupInlineKeyboard or null
func NewInputInlineQueryResultGame(id string, gameShortName string, replyMarkup ReplyMarkup) *InputInlineQueryResultGame {
	inputInlineQueryResultGameTemp := InputInlineQueryResultGame{
		tdCommon:      tdCommon{Type: "inputInlineQueryResultGame"},
		Id:            id,
		GameShortName: gameShortName,
		ReplyMarkup:   replyMarkup,
	}

	return &inputInlineQueryResultGameTemp
}

// UnmarshalJSON unmarshal to json
func (inputInlineQueryResultGame *InputInlineQueryResultGame) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	tempObj := struct {
		tdCommon
		Id            string `json:"id"`              // Unique identifier of the query result
		GameShortName string `json:"game_short_name"` // Short name of the game

	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	inputInlineQueryResultGame.tdCommon = tempObj.tdCommon
	inputInlineQueryResultGame.Id = tempObj.Id
	inputInlineQueryResultGame.GameShortName = tempObj.GameShortName

	fieldReplyMarkup, _ := unmarshalReplyMarkup(objMap["reply_markup"])
	inputInlineQueryResultGame.ReplyMarkup = fieldReplyMarkup

	return nil
}

// GetInputInlineQueryResultEnum return the enum type of this object
func (inputInlineQueryResultGame *InputInlineQueryResultGame) GetInputInlineQueryResultEnum() InputInlineQueryResultEnum {
	return InputInlineQueryResultGameType
}

// InputInlineQueryResultLocation Represents a point on the map
type InputInlineQueryResultLocation struct {
	tdCommon
	Id                  string              `json:"id"`                    // Unique identifier of the query result
	Location            *Location           `json:"location"`              // Location result
	LivePeriod          int32               `json:"live_period"`           // Amount of time relative to the message sent time until the location can be updated, in seconds
	Title               string              `json:"title"`                 // Title of the result
	ThumbnailUrl        string              `json:"thumbnail_url"`         // URL of the result thumbnail, if it exists
	ThumbnailWidth      int32               `json:"thumbnail_width"`       // Thumbnail width, if known
	ThumbnailHeight     int32               `json:"thumbnail_height"`      // Thumbnail height, if known
	ReplyMarkup         ReplyMarkup         `json:"reply_markup"`          // The message reply markup; pass null if none. Must be of type replyMarkupInlineKeyboard or null
	InputMessageContent InputMessageContent `json:"input_message_content"` // The content of the message to be sent. Must be one of the following types: inputMessageText, inputMessageInvoice, inputMessageLocation, inputMessageVenue or inputMessageContact
}

// MessageType return the string telegram-type of InputInlineQueryResultLocation
func (inputInlineQueryResultLocation *InputInlineQueryResultLocation) MessageType() string {
	return "inputInlineQueryResultLocation"
}

// NewInputInlineQueryResultLocation creates a new InputInlineQueryResultLocation
//
// @param id Unique identifier of the query result
// @param location Location result
// @param livePeriod Amount of time relative to the message sent time until the location can be updated, in seconds
// @param title Title of the result
// @param thumbnailUrl URL of the result thumbnail, if it exists
// @param thumbnailWidth Thumbnail width, if known
// @param thumbnailHeight Thumbnail height, if known
// @param replyMarkup The message reply markup; pass null if none. Must be of type replyMarkupInlineKeyboard or null
// @param inputMessageContent The content of the message to be sent. Must be one of the following types: inputMessageText, inputMessageInvoice, inputMessageLocation, inputMessageVenue or inputMessageContact
func NewInputInlineQueryResultLocation(id string, location *Location, livePeriod int32, title string, thumbnailUrl string, thumbnailWidth int32, thumbnailHeight int32, replyMarkup ReplyMarkup, inputMessageContent InputMessageContent) *InputInlineQueryResultLocation {
	inputInlineQueryResultLocationTemp := InputInlineQueryResultLocation{
		tdCommon:            tdCommon{Type: "inputInlineQueryResultLocation"},
		Id:                  id,
		Location:            location,
		LivePeriod:          livePeriod,
		Title:               title,
		ThumbnailUrl:        thumbnailUrl,
		ThumbnailWidth:      thumbnailWidth,
		ThumbnailHeight:     thumbnailHeight,
		ReplyMarkup:         replyMarkup,
		InputMessageContent: inputMessageContent,
	}

	return &inputInlineQueryResultLocationTemp
}

// UnmarshalJSON unmarshal to json
func (inputInlineQueryResultLocation *InputInlineQueryResultLocation) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	tempObj := struct {
		tdCommon
		Id              string    `json:"id"`               // Unique identifier of the query result
		Location        *Location `json:"location"`         // Location result
		LivePeriod      int32     `json:"live_period"`      // Amount of time relative to the message sent time until the location can be updated, in seconds
		Title           string    `json:"title"`            // Title of the result
		ThumbnailUrl    string    `json:"thumbnail_url"`    // URL of the result thumbnail, if it exists
		ThumbnailWidth  int32     `json:"thumbnail_width"`  // Thumbnail width, if known
		ThumbnailHeight int32     `json:"thumbnail_height"` // Thumbnail height, if known

	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	inputInlineQueryResultLocation.tdCommon = tempObj.tdCommon
	inputInlineQueryResultLocation.Id = tempObj.Id
	inputInlineQueryResultLocation.Location = tempObj.Location
	inputInlineQueryResultLocation.LivePeriod = tempObj.LivePeriod
	inputInlineQueryResultLocation.Title = tempObj.Title
	inputInlineQueryResultLocation.ThumbnailUrl = tempObj.ThumbnailUrl
	inputInlineQueryResultLocation.ThumbnailWidth = tempObj.ThumbnailWidth
	inputInlineQueryResultLocation.ThumbnailHeight = tempObj.ThumbnailHeight

	fieldReplyMarkup, _ := unmarshalReplyMarkup(objMap["reply_markup"])
	inputInlineQueryResultLocation.ReplyMarkup = fieldReplyMarkup

	fieldInputMessageContent, _ := unmarshalInputMessageContent(objMap["input_message_content"])
	inputInlineQueryResultLocation.InputMessageContent = fieldInputMessageContent

	return nil
}

// GetInputInlineQueryResultEnum return the enum type of this object
func (inputInlineQueryResultLocation *InputInlineQueryResultLocation) GetInputInlineQueryResultEnum() InputInlineQueryResultEnum {
	return InputInlineQueryResultLocationType
}

// InputInlineQueryResultPhoto Represents link to a JPEG image
type InputInlineQueryResultPhoto struct {
	tdCommon
	Id                  string              `json:"id"`                    // Unique identifier of the query result
	Title               string              `json:"title"`                 // Title of the result, if known
	Description         string              `json:"description"`           // A short description of the result, if known
	ThumbnailUrl        string              `json:"thumbnail_url"`         // URL of the photo thumbnail, if it exists
	PhotoUrl            string              `json:"photo_url"`             // The URL of the JPEG photo (photo size must not exceed 5MB)
	PhotoWidth          int32               `json:"photo_width"`           // Width of the photo
	PhotoHeight         int32               `json:"photo_height"`          // Height of the photo
	ReplyMarkup         ReplyMarkup         `json:"reply_markup"`          // The message reply markup; pass null if none. Must be of type replyMarkupInlineKeyboard or null
	InputMessageContent InputMessageContent `json:"input_message_content"` // The content of the message to be sent. Must be one of the following types: inputMessageText, inputMessagePhoto, inputMessageInvoice, inputMessageLocation, inputMessageVenue or inputMessageContact
}

// MessageType return the string telegram-type of InputInlineQueryResultPhoto
func (inputInlineQueryResultPhoto *InputInlineQueryResultPhoto) MessageType() string {
	return "inputInlineQueryResultPhoto"
}

// NewInputInlineQueryResultPhoto creates a new InputInlineQueryResultPhoto
//
// @param id Unique identifier of the query result
// @param title Title of the result, if known
// @param description A short description of the result, if known
// @param thumbnailUrl URL of the photo thumbnail, if it exists
// @param photoUrl The URL of the JPEG photo (photo size must not exceed 5MB)
// @param photoWidth Width of the photo
// @param photoHeight Height of the photo
// @param replyMarkup The message reply markup; pass null if none. Must be of type replyMarkupInlineKeyboard or null
// @param inputMessageContent The content of the message to be sent. Must be one of the following types: inputMessageText, inputMessagePhoto, inputMessageInvoice, inputMessageLocation, inputMessageVenue or inputMessageContact
func NewInputInlineQueryResultPhoto(id string, title string, description string, thumbnailUrl string, photoUrl string, photoWidth int32, photoHeight int32, replyMarkup ReplyMarkup, inputMessageContent InputMessageContent) *InputInlineQueryResultPhoto {
	inputInlineQueryResultPhotoTemp := InputInlineQueryResultPhoto{
		tdCommon:            tdCommon{Type: "inputInlineQueryResultPhoto"},
		Id:                  id,
		Title:               title,
		Description:         description,
		ThumbnailUrl:        thumbnailUrl,
		PhotoUrl:            photoUrl,
		PhotoWidth:          photoWidth,
		PhotoHeight:         photoHeight,
		ReplyMarkup:         replyMarkup,
		InputMessageContent: inputMessageContent,
	}

	return &inputInlineQueryResultPhotoTemp
}

// UnmarshalJSON unmarshal to json
func (inputInlineQueryResultPhoto *InputInlineQueryResultPhoto) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	tempObj := struct {
		tdCommon
		Id           string `json:"id"`            // Unique identifier of the query result
		Title        string `json:"title"`         // Title of the result, if known
		Description  string `json:"description"`   // A short description of the result, if known
		ThumbnailUrl string `json:"thumbnail_url"` // URL of the photo thumbnail, if it exists
		PhotoUrl     string `json:"photo_url"`     // The URL of the JPEG photo (photo size must not exceed 5MB)
		PhotoWidth   int32  `json:"photo_width"`   // Width of the photo
		PhotoHeight  int32  `json:"photo_height"`  // Height of the photo

	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	inputInlineQueryResultPhoto.tdCommon = tempObj.tdCommon
	inputInlineQueryResultPhoto.Id = tempObj.Id
	inputInlineQueryResultPhoto.Title = tempObj.Title
	inputInlineQueryResultPhoto.Description = tempObj.Description
	inputInlineQueryResultPhoto.ThumbnailUrl = tempObj.ThumbnailUrl
	inputInlineQueryResultPhoto.PhotoUrl = tempObj.PhotoUrl
	inputInlineQueryResultPhoto.PhotoWidth = tempObj.PhotoWidth
	inputInlineQueryResultPhoto.PhotoHeight = tempObj.PhotoHeight

	fieldReplyMarkup, _ := unmarshalReplyMarkup(objMap["reply_markup"])
	inputInlineQueryResultPhoto.ReplyMarkup = fieldReplyMarkup

	fieldInputMessageContent, _ := unmarshalInputMessageContent(objMap["input_message_content"])
	inputInlineQueryResultPhoto.InputMessageContent = fieldInputMessageContent

	return nil
}

// GetInputInlineQueryResultEnum return the enum type of this object
func (inputInlineQueryResultPhoto *InputInlineQueryResultPhoto) GetInputInlineQueryResultEnum() InputInlineQueryResultEnum {
	return InputInlineQueryResultPhotoType
}

// InputInlineQueryResultSticker Represents a link to a WEBP or TGS sticker
type InputInlineQueryResultSticker struct {
	tdCommon
	Id                  string              `json:"id"`                    // Unique identifier of the query result
	ThumbnailUrl        string              `json:"thumbnail_url"`         // URL of the sticker thumbnail, if it exists
	StickerUrl          string              `json:"sticker_url"`           // The URL of the WEBP or TGS sticker (sticker file size must not exceed 5MB)
	StickerWidth        int32               `json:"sticker_width"`         // Width of the sticker
	StickerHeight       int32               `json:"sticker_height"`        // Height of the sticker
	ReplyMarkup         ReplyMarkup         `json:"reply_markup"`          // The message reply markup; pass null if none. Must be of type replyMarkupInlineKeyboard or null
	InputMessageContent InputMessageContent `json:"input_message_content"` // The content of the message to be sent. Must be one of the following types: inputMessageText, inputMessageSticker, inputMessageInvoice, inputMessageLocation, inputMessageVenue or inputMessageContact
}

// MessageType return the string telegram-type of InputInlineQueryResultSticker
func (inputInlineQueryResultSticker *InputInlineQueryResultSticker) MessageType() string {
	return "inputInlineQueryResultSticker"
}

// NewInputInlineQueryResultSticker creates a new InputInlineQueryResultSticker
//
// @param id Unique identifier of the query result
// @param thumbnailUrl URL of the sticker thumbnail, if it exists
// @param stickerUrl The URL of the WEBP or TGS sticker (sticker file size must not exceed 5MB)
// @param stickerWidth Width of the sticker
// @param stickerHeight Height of the sticker
// @param replyMarkup The message reply markup; pass null if none. Must be of type replyMarkupInlineKeyboard or null
// @param inputMessageContent The content of the message to be sent. Must be one of the following types: inputMessageText, inputMessageSticker, inputMessageInvoice, inputMessageLocation, inputMessageVenue or inputMessageContact
func NewInputInlineQueryResultSticker(id string, thumbnailUrl string, stickerUrl string, stickerWidth int32, stickerHeight int32, replyMarkup ReplyMarkup, inputMessageContent InputMessageContent) *InputInlineQueryResultSticker {
	inputInlineQueryResultStickerTemp := InputInlineQueryResultSticker{
		tdCommon:            tdCommon{Type: "inputInlineQueryResultSticker"},
		Id:                  id,
		ThumbnailUrl:        thumbnailUrl,
		StickerUrl:          stickerUrl,
		StickerWidth:        stickerWidth,
		StickerHeight:       stickerHeight,
		ReplyMarkup:         replyMarkup,
		InputMessageContent: inputMessageContent,
	}

	return &inputInlineQueryResultStickerTemp
}

// UnmarshalJSON unmarshal to json
func (inputInlineQueryResultSticker *InputInlineQueryResultSticker) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	tempObj := struct {
		tdCommon
		Id            string `json:"id"`             // Unique identifier of the query result
		ThumbnailUrl  string `json:"thumbnail_url"`  // URL of the sticker thumbnail, if it exists
		StickerUrl    string `json:"sticker_url"`    // The URL of the WEBP or TGS sticker (sticker file size must not exceed 5MB)
		StickerWidth  int32  `json:"sticker_width"`  // Width of the sticker
		StickerHeight int32  `json:"sticker_height"` // Height of the sticker

	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	inputInlineQueryResultSticker.tdCommon = tempObj.tdCommon
	inputInlineQueryResultSticker.Id = tempObj.Id
	inputInlineQueryResultSticker.ThumbnailUrl = tempObj.ThumbnailUrl
	inputInlineQueryResultSticker.StickerUrl = tempObj.StickerUrl
	inputInlineQueryResultSticker.StickerWidth = tempObj.StickerWidth
	inputInlineQueryResultSticker.StickerHeight = tempObj.StickerHeight

	fieldReplyMarkup, _ := unmarshalReplyMarkup(objMap["reply_markup"])
	inputInlineQueryResultSticker.ReplyMarkup = fieldReplyMarkup

	fieldInputMessageContent, _ := unmarshalInputMessageContent(objMap["input_message_content"])
	inputInlineQueryResultSticker.InputMessageContent = fieldInputMessageContent

	return nil
}

// GetInputInlineQueryResultEnum return the enum type of this object
func (inputInlineQueryResultSticker *InputInlineQueryResultSticker) GetInputInlineQueryResultEnum() InputInlineQueryResultEnum {
	return InputInlineQueryResultStickerType
}

// InputInlineQueryResultVenue Represents information about a venue
type InputInlineQueryResultVenue struct {
	tdCommon
	Id                  string              `json:"id"`                    // Unique identifier of the query result
	Venue               *Venue              `json:"venue"`                 // Venue result
	ThumbnailUrl        string              `json:"thumbnail_url"`         // URL of the result thumbnail, if it exists
	ThumbnailWidth      int32               `json:"thumbnail_width"`       // Thumbnail width, if known
	ThumbnailHeight     int32               `json:"thumbnail_height"`      // Thumbnail height, if known
	ReplyMarkup         ReplyMarkup         `json:"reply_markup"`          // The message reply markup; pass null if none. Must be of type replyMarkupInlineKeyboard or null
	InputMessageContent InputMessageContent `json:"input_message_content"` // The content of the message to be sent. Must be one of the following types: inputMessageText, inputMessageInvoice, inputMessageLocation, inputMessageVenue or inputMessageContact
}

// MessageType return the string telegram-type of InputInlineQueryResultVenue
func (inputInlineQueryResultVenue *InputInlineQueryResultVenue) MessageType() string {
	return "inputInlineQueryResultVenue"
}

// NewInputInlineQueryResultVenue creates a new InputInlineQueryResultVenue
//
// @param id Unique identifier of the query result
// @param venue Venue result
// @param thumbnailUrl URL of the result thumbnail, if it exists
// @param thumbnailWidth Thumbnail width, if known
// @param thumbnailHeight Thumbnail height, if known
// @param replyMarkup The message reply markup; pass null if none. Must be of type replyMarkupInlineKeyboard or null
// @param inputMessageContent The content of the message to be sent. Must be one of the following types: inputMessageText, inputMessageInvoice, inputMessageLocation, inputMessageVenue or inputMessageContact
func NewInputInlineQueryResultVenue(id string, venue *Venue, thumbnailUrl string, thumbnailWidth int32, thumbnailHeight int32, replyMarkup ReplyMarkup, inputMessageContent InputMessageContent) *InputInlineQueryResultVenue {
	inputInlineQueryResultVenueTemp := InputInlineQueryResultVenue{
		tdCommon:            tdCommon{Type: "inputInlineQueryResultVenue"},
		Id:                  id,
		Venue:               venue,
		ThumbnailUrl:        thumbnailUrl,
		ThumbnailWidth:      thumbnailWidth,
		ThumbnailHeight:     thumbnailHeight,
		ReplyMarkup:         replyMarkup,
		InputMessageContent: inputMessageContent,
	}

	return &inputInlineQueryResultVenueTemp
}

// UnmarshalJSON unmarshal to json
func (inputInlineQueryResultVenue *InputInlineQueryResultVenue) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	tempObj := struct {
		tdCommon
		Id              string `json:"id"`               // Unique identifier of the query result
		Venue           *Venue `json:"venue"`            // Venue result
		ThumbnailUrl    string `json:"thumbnail_url"`    // URL of the result thumbnail, if it exists
		ThumbnailWidth  int32  `json:"thumbnail_width"`  // Thumbnail width, if known
		ThumbnailHeight int32  `json:"thumbnail_height"` // Thumbnail height, if known

	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	inputInlineQueryResultVenue.tdCommon = tempObj.tdCommon
	inputInlineQueryResultVenue.Id = tempObj.Id
	inputInlineQueryResultVenue.Venue = tempObj.Venue
	inputInlineQueryResultVenue.ThumbnailUrl = tempObj.ThumbnailUrl
	inputInlineQueryResultVenue.ThumbnailWidth = tempObj.ThumbnailWidth
	inputInlineQueryResultVenue.ThumbnailHeight = tempObj.ThumbnailHeight

	fieldReplyMarkup, _ := unmarshalReplyMarkup(objMap["reply_markup"])
	inputInlineQueryResultVenue.ReplyMarkup = fieldReplyMarkup

	fieldInputMessageContent, _ := unmarshalInputMessageContent(objMap["input_message_content"])
	inputInlineQueryResultVenue.InputMessageContent = fieldInputMessageContent

	return nil
}

// GetInputInlineQueryResultEnum return the enum type of this object
func (inputInlineQueryResultVenue *InputInlineQueryResultVenue) GetInputInlineQueryResultEnum() InputInlineQueryResultEnum {
	return InputInlineQueryResultVenueType
}

// InputInlineQueryResultVideo Represents a link to a page containing an embedded video player or a video file
type InputInlineQueryResultVideo struct {
	tdCommon
	Id                  string              `json:"id"`                    // Unique identifier of the query result
	Title               string              `json:"title"`                 // Title of the result
	Description         string              `json:"description"`           // A short description of the result, if known
	ThumbnailUrl        string              `json:"thumbnail_url"`         // The URL of the video thumbnail (JPEG), if it exists
	VideoUrl            string              `json:"video_url"`             // URL of the embedded video player or video file
	MimeType            string              `json:"mime_type"`             // MIME type of the content of the video URL, only "text/html" or "video/mp4" are currently supported
	VideoWidth          int32               `json:"video_width"`           // Width of the video
	VideoHeight         int32               `json:"video_height"`          // Height of the video
	VideoDuration       int32               `json:"video_duration"`        // Video duration, in seconds
	ReplyMarkup         ReplyMarkup         `json:"reply_markup"`          // The message reply markup; pass null if none. Must be of type replyMarkupInlineKeyboard or null
	InputMessageContent InputMessageContent `json:"input_message_content"` // The content of the message to be sent. Must be one of the following types: inputMessageText, inputMessageVideo, inputMessageInvoice, inputMessageLocation, inputMessageVenue or inputMessageContact
}

// MessageType return the string telegram-type of InputInlineQueryResultVideo
func (inputInlineQueryResultVideo *InputInlineQueryResultVideo) MessageType() string {
	return "inputInlineQueryResultVideo"
}

// NewInputInlineQueryResultVideo creates a new InputInlineQueryResultVideo
//
// @param id Unique identifier of the query result
// @param title Title of the result
// @param description A short description of the result, if known
// @param thumbnailUrl The URL of the video thumbnail (JPEG), if it exists
// @param videoUrl URL of the embedded video player or video file
// @param mimeType MIME type of the content of the video URL, only "text/html" or "video/mp4" are currently supported
// @param videoWidth Width of the video
// @param videoHeight Height of the video
// @param videoDuration Video duration, in seconds
// @param replyMarkup The message reply markup; pass null if none. Must be of type replyMarkupInlineKeyboard or null
// @param inputMessageContent The content of the message to be sent. Must be one of the following types: inputMessageText, inputMessageVideo, inputMessageInvoice, inputMessageLocation, inputMessageVenue or inputMessageContact
func NewInputInlineQueryResultVideo(id string, title string, description string, thumbnailUrl string, videoUrl string, mimeType string, videoWidth int32, videoHeight int32, videoDuration int32, replyMarkup ReplyMarkup, inputMessageContent InputMessageContent) *InputInlineQueryResultVideo {
	inputInlineQueryResultVideoTemp := InputInlineQueryResultVideo{
		tdCommon:            tdCommon{Type: "inputInlineQueryResultVideo"},
		Id:                  id,
		Title:               title,
		Description:         description,
		ThumbnailUrl:        thumbnailUrl,
		VideoUrl:            videoUrl,
		MimeType:            mimeType,
		VideoWidth:          videoWidth,
		VideoHeight:         videoHeight,
		VideoDuration:       videoDuration,
		ReplyMarkup:         replyMarkup,
		InputMessageContent: inputMessageContent,
	}

	return &inputInlineQueryResultVideoTemp
}

// UnmarshalJSON unmarshal to json
func (inputInlineQueryResultVideo *InputInlineQueryResultVideo) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	tempObj := struct {
		tdCommon
		Id            string `json:"id"`             // Unique identifier of the query result
		Title         string `json:"title"`          // Title of the result
		Description   string `json:"description"`    // A short description of the result, if known
		ThumbnailUrl  string `json:"thumbnail_url"`  // The URL of the video thumbnail (JPEG), if it exists
		VideoUrl      string `json:"video_url"`      // URL of the embedded video player or video file
		MimeType      string `json:"mime_type"`      // MIME type of the content of the video URL, only "text/html" or "video/mp4" are currently supported
		VideoWidth    int32  `json:"video_width"`    // Width of the video
		VideoHeight   int32  `json:"video_height"`   // Height of the video
		VideoDuration int32  `json:"video_duration"` // Video duration, in seconds

	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	inputInlineQueryResultVideo.tdCommon = tempObj.tdCommon
	inputInlineQueryResultVideo.Id = tempObj.Id
	inputInlineQueryResultVideo.Title = tempObj.Title
	inputInlineQueryResultVideo.Description = tempObj.Description
	inputInlineQueryResultVideo.ThumbnailUrl = tempObj.ThumbnailUrl
	inputInlineQueryResultVideo.VideoUrl = tempObj.VideoUrl
	inputInlineQueryResultVideo.MimeType = tempObj.MimeType
	inputInlineQueryResultVideo.VideoWidth = tempObj.VideoWidth
	inputInlineQueryResultVideo.VideoHeight = tempObj.VideoHeight
	inputInlineQueryResultVideo.VideoDuration = tempObj.VideoDuration

	fieldReplyMarkup, _ := unmarshalReplyMarkup(objMap["reply_markup"])
	inputInlineQueryResultVideo.ReplyMarkup = fieldReplyMarkup

	fieldInputMessageContent, _ := unmarshalInputMessageContent(objMap["input_message_content"])
	inputInlineQueryResultVideo.InputMessageContent = fieldInputMessageContent

	return nil
}

// GetInputInlineQueryResultEnum return the enum type of this object
func (inputInlineQueryResultVideo *InputInlineQueryResultVideo) GetInputInlineQueryResultEnum() InputInlineQueryResultEnum {
	return InputInlineQueryResultVideoType
}

// InputInlineQueryResultVoiceNote Represents a link to an opus-encoded audio file within an OGG container, single channel audio
type InputInlineQueryResultVoiceNote struct {
	tdCommon
	Id                  string              `json:"id"`                    // Unique identifier of the query result
	Title               string              `json:"title"`                 // Title of the voice note
	VoiceNoteUrl        string              `json:"voice_note_url"`        // The URL of the voice note file
	VoiceNoteDuration   int32               `json:"voice_note_duration"`   // Duration of the voice note, in seconds
	ReplyMarkup         ReplyMarkup         `json:"reply_markup"`          // The message reply markup; pass null if none. Must be of type replyMarkupInlineKeyboard or null
	InputMessageContent InputMessageContent `json:"input_message_content"` // The content of the message to be sent. Must be one of the following types: inputMessageText, inputMessageVoiceNote, inputMessageInvoice, inputMessageLocation, inputMessageVenue or inputMessageContact
}

// MessageType return the string telegram-type of InputInlineQueryResultVoiceNote
func (inputInlineQueryResultVoiceNote *InputInlineQueryResultVoiceNote) MessageType() string {
	return "inputInlineQueryResultVoiceNote"
}

// NewInputInlineQueryResultVoiceNote creates a new InputInlineQueryResultVoiceNote
//
// @param id Unique identifier of the query result
// @param title Title of the voice note
// @param voiceNoteUrl The URL of the voice note file
// @param voiceNoteDuration Duration of the voice note, in seconds
// @param replyMarkup The message reply markup; pass null if none. Must be of type replyMarkupInlineKeyboard or null
// @param inputMessageContent The content of the message to be sent. Must be one of the following types: inputMessageText, inputMessageVoiceNote, inputMessageInvoice, inputMessageLocation, inputMessageVenue or inputMessageContact
func NewInputInlineQueryResultVoiceNote(id string, title string, voiceNoteUrl string, voiceNoteDuration int32, replyMarkup ReplyMarkup, inputMessageContent InputMessageContent) *InputInlineQueryResultVoiceNote {
	inputInlineQueryResultVoiceNoteTemp := InputInlineQueryResultVoiceNote{
		tdCommon:            tdCommon{Type: "inputInlineQueryResultVoiceNote"},
		Id:                  id,
		Title:               title,
		VoiceNoteUrl:        voiceNoteUrl,
		VoiceNoteDuration:   voiceNoteDuration,
		ReplyMarkup:         replyMarkup,
		InputMessageContent: inputMessageContent,
	}

	return &inputInlineQueryResultVoiceNoteTemp
}

// UnmarshalJSON unmarshal to json
func (inputInlineQueryResultVoiceNote *InputInlineQueryResultVoiceNote) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	tempObj := struct {
		tdCommon
		Id                string `json:"id"`                  // Unique identifier of the query result
		Title             string `json:"title"`               // Title of the voice note
		VoiceNoteUrl      string `json:"voice_note_url"`      // The URL of the voice note file
		VoiceNoteDuration int32  `json:"voice_note_duration"` // Duration of the voice note, in seconds

	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	inputInlineQueryResultVoiceNote.tdCommon = tempObj.tdCommon
	inputInlineQueryResultVoiceNote.Id = tempObj.Id
	inputInlineQueryResultVoiceNote.Title = tempObj.Title
	inputInlineQueryResultVoiceNote.VoiceNoteUrl = tempObj.VoiceNoteUrl
	inputInlineQueryResultVoiceNote.VoiceNoteDuration = tempObj.VoiceNoteDuration

	fieldReplyMarkup, _ := unmarshalReplyMarkup(objMap["reply_markup"])
	inputInlineQueryResultVoiceNote.ReplyMarkup = fieldReplyMarkup

	fieldInputMessageContent, _ := unmarshalInputMessageContent(objMap["input_message_content"])
	inputInlineQueryResultVoiceNote.InputMessageContent = fieldInputMessageContent

	return nil
}

// GetInputInlineQueryResultEnum return the enum type of this object
func (inputInlineQueryResultVoiceNote *InputInlineQueryResultVoiceNote) GetInputInlineQueryResultEnum() InputInlineQueryResultEnum {
	return InputInlineQueryResultVoiceNoteType
}
