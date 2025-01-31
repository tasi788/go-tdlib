package tdlib

import (
	"encoding/json"
	"fmt"
)

// InputMessageContent The content of a message to send
type InputMessageContent interface {
	GetInputMessageContentEnum() InputMessageContentEnum
}

// InputMessageContentEnum Alias for abstract InputMessageContent 'Sub-Classes', used as constant-enum here
type InputMessageContentEnum string

// InputMessageContent enums
const (
	InputMessageTextType      InputMessageContentEnum = "inputMessageText"
	InputMessageAnimationType InputMessageContentEnum = "inputMessageAnimation"
	InputMessageAudioType     InputMessageContentEnum = "inputMessageAudio"
	InputMessageDocumentType  InputMessageContentEnum = "inputMessageDocument"
	InputMessagePhotoType     InputMessageContentEnum = "inputMessagePhoto"
	InputMessageStickerType   InputMessageContentEnum = "inputMessageSticker"
	InputMessageVideoType     InputMessageContentEnum = "inputMessageVideo"
	InputMessageVideoNoteType InputMessageContentEnum = "inputMessageVideoNote"
	InputMessageVoiceNoteType InputMessageContentEnum = "inputMessageVoiceNote"
	InputMessageLocationType  InputMessageContentEnum = "inputMessageLocation"
	InputMessageVenueType     InputMessageContentEnum = "inputMessageVenue"
	InputMessageContactType   InputMessageContentEnum = "inputMessageContact"
	InputMessageDiceType      InputMessageContentEnum = "inputMessageDice"
	InputMessageGameType      InputMessageContentEnum = "inputMessageGame"
	InputMessageInvoiceType   InputMessageContentEnum = "inputMessageInvoice"
	InputMessagePollType      InputMessageContentEnum = "inputMessagePoll"
	InputMessageForwardedType InputMessageContentEnum = "inputMessageForwarded"
)

func unmarshalInputMessageContent(rawMsg *json.RawMessage) (InputMessageContent, error) {

	if rawMsg == nil {
		return nil, nil
	}
	var objMap map[string]interface{}
	err := json.Unmarshal(*rawMsg, &objMap)
	if err != nil {
		return nil, err
	}

	switch InputMessageContentEnum(objMap["@type"].(string)) {
	case InputMessageTextType:
		var inputMessageText InputMessageText
		err := json.Unmarshal(*rawMsg, &inputMessageText)
		return &inputMessageText, err

	case InputMessageAnimationType:
		var inputMessageAnimation InputMessageAnimation
		err := json.Unmarshal(*rawMsg, &inputMessageAnimation)
		return &inputMessageAnimation, err

	case InputMessageAudioType:
		var inputMessageAudio InputMessageAudio
		err := json.Unmarshal(*rawMsg, &inputMessageAudio)
		return &inputMessageAudio, err

	case InputMessageDocumentType:
		var inputMessageDocument InputMessageDocument
		err := json.Unmarshal(*rawMsg, &inputMessageDocument)
		return &inputMessageDocument, err

	case InputMessagePhotoType:
		var inputMessagePhoto InputMessagePhoto
		err := json.Unmarshal(*rawMsg, &inputMessagePhoto)
		return &inputMessagePhoto, err

	case InputMessageStickerType:
		var inputMessageSticker InputMessageSticker
		err := json.Unmarshal(*rawMsg, &inputMessageSticker)
		return &inputMessageSticker, err

	case InputMessageVideoType:
		var inputMessageVideo InputMessageVideo
		err := json.Unmarshal(*rawMsg, &inputMessageVideo)
		return &inputMessageVideo, err

	case InputMessageVideoNoteType:
		var inputMessageVideoNote InputMessageVideoNote
		err := json.Unmarshal(*rawMsg, &inputMessageVideoNote)
		return &inputMessageVideoNote, err

	case InputMessageVoiceNoteType:
		var inputMessageVoiceNote InputMessageVoiceNote
		err := json.Unmarshal(*rawMsg, &inputMessageVoiceNote)
		return &inputMessageVoiceNote, err

	case InputMessageLocationType:
		var inputMessageLocation InputMessageLocation
		err := json.Unmarshal(*rawMsg, &inputMessageLocation)
		return &inputMessageLocation, err

	case InputMessageVenueType:
		var inputMessageVenue InputMessageVenue
		err := json.Unmarshal(*rawMsg, &inputMessageVenue)
		return &inputMessageVenue, err

	case InputMessageContactType:
		var inputMessageContact InputMessageContact
		err := json.Unmarshal(*rawMsg, &inputMessageContact)
		return &inputMessageContact, err

	case InputMessageDiceType:
		var inputMessageDice InputMessageDice
		err := json.Unmarshal(*rawMsg, &inputMessageDice)
		return &inputMessageDice, err

	case InputMessageGameType:
		var inputMessageGame InputMessageGame
		err := json.Unmarshal(*rawMsg, &inputMessageGame)
		return &inputMessageGame, err

	case InputMessageInvoiceType:
		var inputMessageInvoice InputMessageInvoice
		err := json.Unmarshal(*rawMsg, &inputMessageInvoice)
		return &inputMessageInvoice, err

	case InputMessagePollType:
		var inputMessagePoll InputMessagePoll
		err := json.Unmarshal(*rawMsg, &inputMessagePoll)
		return &inputMessagePoll, err

	case InputMessageForwardedType:
		var inputMessageForwarded InputMessageForwarded
		err := json.Unmarshal(*rawMsg, &inputMessageForwarded)
		return &inputMessageForwarded, err

	default:
		return nil, fmt.Errorf("Error UnMarshaling, unknown type:" + objMap["@type"].(string))
	}
}

// InputMessageText A text message
type InputMessageText struct {
	tdCommon
	Text                  *FormattedText `json:"text"`                     // Formatted text to be sent; 1-GetOption("message_text_length_max") characters. Only Bold, Italic, Underline, Strikethrough, Spoiler, Code, Pre, PreCode, TextUrl and MentionName entities are allowed to be specified manually
	DisableWebPagePreview bool           `json:"disable_web_page_preview"` // True, if rich web page previews for URLs in the message text must be disabled
	ClearDraft            bool           `json:"clear_draft"`              // True, if a chat message draft must be deleted
}

// MessageType return the string telegram-type of InputMessageText
func (inputMessageText *InputMessageText) MessageType() string {
	return "inputMessageText"
}

// NewInputMessageText creates a new InputMessageText
//
// @param text Formatted text to be sent; 1-GetOption("message_text_length_max") characters. Only Bold, Italic, Underline, Strikethrough, Spoiler, Code, Pre, PreCode, TextUrl and MentionName entities are allowed to be specified manually
// @param disableWebPagePreview True, if rich web page previews for URLs in the message text must be disabled
// @param clearDraft True, if a chat message draft must be deleted
func NewInputMessageText(text *FormattedText, disableWebPagePreview bool, clearDraft bool) *InputMessageText {
	inputMessageTextTemp := InputMessageText{
		tdCommon:              tdCommon{Type: "inputMessageText"},
		Text:                  text,
		DisableWebPagePreview: disableWebPagePreview,
		ClearDraft:            clearDraft,
	}

	return &inputMessageTextTemp
}

// GetInputMessageContentEnum return the enum type of this object
func (inputMessageText *InputMessageText) GetInputMessageContentEnum() InputMessageContentEnum {
	return InputMessageTextType
}

// InputMessageAnimation An animation message (GIF-style).
type InputMessageAnimation struct {
	tdCommon
	Animation           InputFile       `json:"animation"`              // Animation file to be sent
	Thumbnail           *InputThumbnail `json:"thumbnail"`              // Animation thumbnail; pass null to skip thumbnail uploading
	AddedStickerFileIds []int32         `json:"added_sticker_file_ids"` // File identifiers of the stickers added to the animation, if applicable
	Duration            int32           `json:"duration"`               // Duration of the animation, in seconds
	Width               int32           `json:"width"`                  // Width of the animation; may be replaced by the server
	Height              int32           `json:"height"`                 // Height of the animation; may be replaced by the server
	Caption             *FormattedText  `json:"caption"`                // Animation caption; pass null to use an empty caption; 0-GetOption("message_caption_length_max") characters
}

// MessageType return the string telegram-type of InputMessageAnimation
func (inputMessageAnimation *InputMessageAnimation) MessageType() string {
	return "inputMessageAnimation"
}

// NewInputMessageAnimation creates a new InputMessageAnimation
//
// @param animation Animation file to be sent
// @param thumbnail Animation thumbnail; pass null to skip thumbnail uploading
// @param addedStickerFileIds File identifiers of the stickers added to the animation, if applicable
// @param duration Duration of the animation, in seconds
// @param width Width of the animation; may be replaced by the server
// @param height Height of the animation; may be replaced by the server
// @param caption Animation caption; pass null to use an empty caption; 0-GetOption("message_caption_length_max") characters
func NewInputMessageAnimation(animation InputFile, thumbnail *InputThumbnail, addedStickerFileIds []int32, duration int32, width int32, height int32, caption *FormattedText) *InputMessageAnimation {
	inputMessageAnimationTemp := InputMessageAnimation{
		tdCommon:            tdCommon{Type: "inputMessageAnimation"},
		Animation:           animation,
		Thumbnail:           thumbnail,
		AddedStickerFileIds: addedStickerFileIds,
		Duration:            duration,
		Width:               width,
		Height:              height,
		Caption:             caption,
	}

	return &inputMessageAnimationTemp
}

// UnmarshalJSON unmarshal to json
func (inputMessageAnimation *InputMessageAnimation) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	tempObj := struct {
		tdCommon
		Thumbnail           *InputThumbnail `json:"thumbnail"`              // Animation thumbnail; pass null to skip thumbnail uploading
		AddedStickerFileIds []int32         `json:"added_sticker_file_ids"` // File identifiers of the stickers added to the animation, if applicable
		Duration            int32           `json:"duration"`               // Duration of the animation, in seconds
		Width               int32           `json:"width"`                  // Width of the animation; may be replaced by the server
		Height              int32           `json:"height"`                 // Height of the animation; may be replaced by the server
		Caption             *FormattedText  `json:"caption"`                // Animation caption; pass null to use an empty caption; 0-GetOption("message_caption_length_max") characters
	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	inputMessageAnimation.tdCommon = tempObj.tdCommon
	inputMessageAnimation.Thumbnail = tempObj.Thumbnail
	inputMessageAnimation.AddedStickerFileIds = tempObj.AddedStickerFileIds
	inputMessageAnimation.Duration = tempObj.Duration
	inputMessageAnimation.Width = tempObj.Width
	inputMessageAnimation.Height = tempObj.Height
	inputMessageAnimation.Caption = tempObj.Caption

	fieldAnimation, _ := unmarshalInputFile(objMap["animation"])
	inputMessageAnimation.Animation = fieldAnimation

	return nil
}

// GetInputMessageContentEnum return the enum type of this object
func (inputMessageAnimation *InputMessageAnimation) GetInputMessageContentEnum() InputMessageContentEnum {
	return InputMessageAnimationType
}

// InputMessageAudio An audio message
type InputMessageAudio struct {
	tdCommon
	Audio               InputFile       `json:"audio"`                 // Audio file to be sent
	AlbumCoverThumbnail *InputThumbnail `json:"album_cover_thumbnail"` // Thumbnail of the cover for the album; pass null to skip thumbnail uploading
	Duration            int32           `json:"duration"`              // Duration of the audio, in seconds; may be replaced by the server
	Title               string          `json:"title"`                 // Title of the audio; 0-64 characters; may be replaced by the server
	Performer           string          `json:"performer"`             // Performer of the audio; 0-64 characters, may be replaced by the server
	Caption             *FormattedText  `json:"caption"`               // Audio caption; pass null to use an empty caption; 0-GetOption("message_caption_length_max") characters
}

// MessageType return the string telegram-type of InputMessageAudio
func (inputMessageAudio *InputMessageAudio) MessageType() string {
	return "inputMessageAudio"
}

// NewInputMessageAudio creates a new InputMessageAudio
//
// @param audio Audio file to be sent
// @param albumCoverThumbnail Thumbnail of the cover for the album; pass null to skip thumbnail uploading
// @param duration Duration of the audio, in seconds; may be replaced by the server
// @param title Title of the audio; 0-64 characters; may be replaced by the server
// @param performer Performer of the audio; 0-64 characters, may be replaced by the server
// @param caption Audio caption; pass null to use an empty caption; 0-GetOption("message_caption_length_max") characters
func NewInputMessageAudio(audio InputFile, albumCoverThumbnail *InputThumbnail, duration int32, title string, performer string, caption *FormattedText) *InputMessageAudio {
	inputMessageAudioTemp := InputMessageAudio{
		tdCommon:            tdCommon{Type: "inputMessageAudio"},
		Audio:               audio,
		AlbumCoverThumbnail: albumCoverThumbnail,
		Duration:            duration,
		Title:               title,
		Performer:           performer,
		Caption:             caption,
	}

	return &inputMessageAudioTemp
}

// UnmarshalJSON unmarshal to json
func (inputMessageAudio *InputMessageAudio) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	tempObj := struct {
		tdCommon
		AlbumCoverThumbnail *InputThumbnail `json:"album_cover_thumbnail"` // Thumbnail of the cover for the album; pass null to skip thumbnail uploading
		Duration            int32           `json:"duration"`              // Duration of the audio, in seconds; may be replaced by the server
		Title               string          `json:"title"`                 // Title of the audio; 0-64 characters; may be replaced by the server
		Performer           string          `json:"performer"`             // Performer of the audio; 0-64 characters, may be replaced by the server
		Caption             *FormattedText  `json:"caption"`               // Audio caption; pass null to use an empty caption; 0-GetOption("message_caption_length_max") characters
	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	inputMessageAudio.tdCommon = tempObj.tdCommon
	inputMessageAudio.AlbumCoverThumbnail = tempObj.AlbumCoverThumbnail
	inputMessageAudio.Duration = tempObj.Duration
	inputMessageAudio.Title = tempObj.Title
	inputMessageAudio.Performer = tempObj.Performer
	inputMessageAudio.Caption = tempObj.Caption

	fieldAudio, _ := unmarshalInputFile(objMap["audio"])
	inputMessageAudio.Audio = fieldAudio

	return nil
}

// GetInputMessageContentEnum return the enum type of this object
func (inputMessageAudio *InputMessageAudio) GetInputMessageContentEnum() InputMessageContentEnum {
	return InputMessageAudioType
}

// InputMessageDocument A document message (general file)
type InputMessageDocument struct {
	tdCommon
	Document                    InputFile       `json:"document"`                       // Document to be sent
	Thumbnail                   *InputThumbnail `json:"thumbnail"`                      // Document thumbnail; pass null to skip thumbnail uploading
	DisableContentTypeDetection bool            `json:"disable_content_type_detection"` // If true, automatic file type detection will be disabled and the document will be always sent as file. Always true for files sent to secret chats
	Caption                     *FormattedText  `json:"caption"`                        // Document caption; pass null to use an empty caption; 0-GetOption("message_caption_length_max") characters
}

// MessageType return the string telegram-type of InputMessageDocument
func (inputMessageDocument *InputMessageDocument) MessageType() string {
	return "inputMessageDocument"
}

// NewInputMessageDocument creates a new InputMessageDocument
//
// @param document Document to be sent
// @param thumbnail Document thumbnail; pass null to skip thumbnail uploading
// @param disableContentTypeDetection If true, automatic file type detection will be disabled and the document will be always sent as file. Always true for files sent to secret chats
// @param caption Document caption; pass null to use an empty caption; 0-GetOption("message_caption_length_max") characters
func NewInputMessageDocument(document InputFile, thumbnail *InputThumbnail, disableContentTypeDetection bool, caption *FormattedText) *InputMessageDocument {
	inputMessageDocumentTemp := InputMessageDocument{
		tdCommon:                    tdCommon{Type: "inputMessageDocument"},
		Document:                    document,
		Thumbnail:                   thumbnail,
		DisableContentTypeDetection: disableContentTypeDetection,
		Caption:                     caption,
	}

	return &inputMessageDocumentTemp
}

// UnmarshalJSON unmarshal to json
func (inputMessageDocument *InputMessageDocument) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	tempObj := struct {
		tdCommon
		Thumbnail                   *InputThumbnail `json:"thumbnail"`                      // Document thumbnail; pass null to skip thumbnail uploading
		DisableContentTypeDetection bool            `json:"disable_content_type_detection"` // If true, automatic file type detection will be disabled and the document will be always sent as file. Always true for files sent to secret chats
		Caption                     *FormattedText  `json:"caption"`                        // Document caption; pass null to use an empty caption; 0-GetOption("message_caption_length_max") characters
	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	inputMessageDocument.tdCommon = tempObj.tdCommon
	inputMessageDocument.Thumbnail = tempObj.Thumbnail
	inputMessageDocument.DisableContentTypeDetection = tempObj.DisableContentTypeDetection
	inputMessageDocument.Caption = tempObj.Caption

	fieldDocument, _ := unmarshalInputFile(objMap["document"])
	inputMessageDocument.Document = fieldDocument

	return nil
}

// GetInputMessageContentEnum return the enum type of this object
func (inputMessageDocument *InputMessageDocument) GetInputMessageContentEnum() InputMessageContentEnum {
	return InputMessageDocumentType
}

// InputMessagePhoto A photo message
type InputMessagePhoto struct {
	tdCommon
	Photo               InputFile       `json:"photo"`                  // Photo to send
	Thumbnail           *InputThumbnail `json:"thumbnail"`              // Photo thumbnail to be sent; pass null to skip thumbnail uploading. The thumbnail is sent to the other party only in secret chats
	AddedStickerFileIds []int32         `json:"added_sticker_file_ids"` // File identifiers of the stickers added to the photo, if applicable
	Width               int32           `json:"width"`                  // Photo width
	Height              int32           `json:"height"`                 // Photo height
	Caption             *FormattedText  `json:"caption"`                // Photo caption; pass null to use an empty caption; 0-GetOption("message_caption_length_max") characters
	Ttl                 int32           `json:"ttl"`                    // Photo TTL (Time To Live), in seconds (0-60). A non-zero TTL can be specified only in private chats
}

// MessageType return the string telegram-type of InputMessagePhoto
func (inputMessagePhoto *InputMessagePhoto) MessageType() string {
	return "inputMessagePhoto"
}

// NewInputMessagePhoto creates a new InputMessagePhoto
//
// @param photo Photo to send
// @param thumbnail Photo thumbnail to be sent; pass null to skip thumbnail uploading. The thumbnail is sent to the other party only in secret chats
// @param addedStickerFileIds File identifiers of the stickers added to the photo, if applicable
// @param width Photo width
// @param height Photo height
// @param caption Photo caption; pass null to use an empty caption; 0-GetOption("message_caption_length_max") characters
// @param ttl Photo TTL (Time To Live), in seconds (0-60). A non-zero TTL can be specified only in private chats
func NewInputMessagePhoto(photo InputFile, thumbnail *InputThumbnail, addedStickerFileIds []int32, width int32, height int32, caption *FormattedText, ttl int32) *InputMessagePhoto {
	inputMessagePhotoTemp := InputMessagePhoto{
		tdCommon:            tdCommon{Type: "inputMessagePhoto"},
		Photo:               photo,
		Thumbnail:           thumbnail,
		AddedStickerFileIds: addedStickerFileIds,
		Width:               width,
		Height:              height,
		Caption:             caption,
		Ttl:                 ttl,
	}

	return &inputMessagePhotoTemp
}

// UnmarshalJSON unmarshal to json
func (inputMessagePhoto *InputMessagePhoto) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	tempObj := struct {
		tdCommon
		Thumbnail           *InputThumbnail `json:"thumbnail"`              // Photo thumbnail to be sent; pass null to skip thumbnail uploading. The thumbnail is sent to the other party only in secret chats
		AddedStickerFileIds []int32         `json:"added_sticker_file_ids"` // File identifiers of the stickers added to the photo, if applicable
		Width               int32           `json:"width"`                  // Photo width
		Height              int32           `json:"height"`                 // Photo height
		Caption             *FormattedText  `json:"caption"`                // Photo caption; pass null to use an empty caption; 0-GetOption("message_caption_length_max") characters
		Ttl                 int32           `json:"ttl"`                    // Photo TTL (Time To Live), in seconds (0-60). A non-zero TTL can be specified only in private chats
	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	inputMessagePhoto.tdCommon = tempObj.tdCommon
	inputMessagePhoto.Thumbnail = tempObj.Thumbnail
	inputMessagePhoto.AddedStickerFileIds = tempObj.AddedStickerFileIds
	inputMessagePhoto.Width = tempObj.Width
	inputMessagePhoto.Height = tempObj.Height
	inputMessagePhoto.Caption = tempObj.Caption
	inputMessagePhoto.Ttl = tempObj.Ttl

	fieldPhoto, _ := unmarshalInputFile(objMap["photo"])
	inputMessagePhoto.Photo = fieldPhoto

	return nil
}

// GetInputMessageContentEnum return the enum type of this object
func (inputMessagePhoto *InputMessagePhoto) GetInputMessageContentEnum() InputMessageContentEnum {
	return InputMessagePhotoType
}

// InputMessageSticker A sticker message
type InputMessageSticker struct {
	tdCommon
	Sticker   InputFile       `json:"sticker"`   // Sticker to be sent
	Thumbnail *InputThumbnail `json:"thumbnail"` // Sticker thumbnail; pass null to skip thumbnail uploading
	Width     int32           `json:"width"`     // Sticker width
	Height    int32           `json:"height"`    // Sticker height
	Emoji     string          `json:"emoji"`     // Emoji used to choose the sticker
}

// MessageType return the string telegram-type of InputMessageSticker
func (inputMessageSticker *InputMessageSticker) MessageType() string {
	return "inputMessageSticker"
}

// NewInputMessageSticker creates a new InputMessageSticker
//
// @param sticker Sticker to be sent
// @param thumbnail Sticker thumbnail; pass null to skip thumbnail uploading
// @param width Sticker width
// @param height Sticker height
// @param emoji Emoji used to choose the sticker
func NewInputMessageSticker(sticker InputFile, thumbnail *InputThumbnail, width int32, height int32, emoji string) *InputMessageSticker {
	inputMessageStickerTemp := InputMessageSticker{
		tdCommon:  tdCommon{Type: "inputMessageSticker"},
		Sticker:   sticker,
		Thumbnail: thumbnail,
		Width:     width,
		Height:    height,
		Emoji:     emoji,
	}

	return &inputMessageStickerTemp
}

// UnmarshalJSON unmarshal to json
func (inputMessageSticker *InputMessageSticker) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	tempObj := struct {
		tdCommon
		Thumbnail *InputThumbnail `json:"thumbnail"` // Sticker thumbnail; pass null to skip thumbnail uploading
		Width     int32           `json:"width"`     // Sticker width
		Height    int32           `json:"height"`    // Sticker height
		Emoji     string          `json:"emoji"`     // Emoji used to choose the sticker
	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	inputMessageSticker.tdCommon = tempObj.tdCommon
	inputMessageSticker.Thumbnail = tempObj.Thumbnail
	inputMessageSticker.Width = tempObj.Width
	inputMessageSticker.Height = tempObj.Height
	inputMessageSticker.Emoji = tempObj.Emoji

	fieldSticker, _ := unmarshalInputFile(objMap["sticker"])
	inputMessageSticker.Sticker = fieldSticker

	return nil
}

// GetInputMessageContentEnum return the enum type of this object
func (inputMessageSticker *InputMessageSticker) GetInputMessageContentEnum() InputMessageContentEnum {
	return InputMessageStickerType
}

// InputMessageVideo A video message
type InputMessageVideo struct {
	tdCommon
	Video               InputFile       `json:"video"`                  // Video to be sent
	Thumbnail           *InputThumbnail `json:"thumbnail"`              // Video thumbnail; pass null to skip thumbnail uploading
	AddedStickerFileIds []int32         `json:"added_sticker_file_ids"` // File identifiers of the stickers added to the video, if applicable
	Duration            int32           `json:"duration"`               // Duration of the video, in seconds
	Width               int32           `json:"width"`                  // Video width
	Height              int32           `json:"height"`                 // Video height
	SupportsStreaming   bool            `json:"supports_streaming"`     // True, if the video is supposed to be streamed
	Caption             *FormattedText  `json:"caption"`                // Video caption; pass null to use an empty caption; 0-GetOption("message_caption_length_max") characters
	Ttl                 int32           `json:"ttl"`                    // Video TTL (Time To Live), in seconds (0-60). A non-zero TTL can be specified only in private chats
}

// MessageType return the string telegram-type of InputMessageVideo
func (inputMessageVideo *InputMessageVideo) MessageType() string {
	return "inputMessageVideo"
}

// NewInputMessageVideo creates a new InputMessageVideo
//
// @param video Video to be sent
// @param thumbnail Video thumbnail; pass null to skip thumbnail uploading
// @param addedStickerFileIds File identifiers of the stickers added to the video, if applicable
// @param duration Duration of the video, in seconds
// @param width Video width
// @param height Video height
// @param supportsStreaming True, if the video is supposed to be streamed
// @param caption Video caption; pass null to use an empty caption; 0-GetOption("message_caption_length_max") characters
// @param ttl Video TTL (Time To Live), in seconds (0-60). A non-zero TTL can be specified only in private chats
func NewInputMessageVideo(video InputFile, thumbnail *InputThumbnail, addedStickerFileIds []int32, duration int32, width int32, height int32, supportsStreaming bool, caption *FormattedText, ttl int32) *InputMessageVideo {
	inputMessageVideoTemp := InputMessageVideo{
		tdCommon:            tdCommon{Type: "inputMessageVideo"},
		Video:               video,
		Thumbnail:           thumbnail,
		AddedStickerFileIds: addedStickerFileIds,
		Duration:            duration,
		Width:               width,
		Height:              height,
		SupportsStreaming:   supportsStreaming,
		Caption:             caption,
		Ttl:                 ttl,
	}

	return &inputMessageVideoTemp
}

// UnmarshalJSON unmarshal to json
func (inputMessageVideo *InputMessageVideo) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	tempObj := struct {
		tdCommon
		Thumbnail           *InputThumbnail `json:"thumbnail"`              // Video thumbnail; pass null to skip thumbnail uploading
		AddedStickerFileIds []int32         `json:"added_sticker_file_ids"` // File identifiers of the stickers added to the video, if applicable
		Duration            int32           `json:"duration"`               // Duration of the video, in seconds
		Width               int32           `json:"width"`                  // Video width
		Height              int32           `json:"height"`                 // Video height
		SupportsStreaming   bool            `json:"supports_streaming"`     // True, if the video is supposed to be streamed
		Caption             *FormattedText  `json:"caption"`                // Video caption; pass null to use an empty caption; 0-GetOption("message_caption_length_max") characters
		Ttl                 int32           `json:"ttl"`                    // Video TTL (Time To Live), in seconds (0-60). A non-zero TTL can be specified only in private chats
	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	inputMessageVideo.tdCommon = tempObj.tdCommon
	inputMessageVideo.Thumbnail = tempObj.Thumbnail
	inputMessageVideo.AddedStickerFileIds = tempObj.AddedStickerFileIds
	inputMessageVideo.Duration = tempObj.Duration
	inputMessageVideo.Width = tempObj.Width
	inputMessageVideo.Height = tempObj.Height
	inputMessageVideo.SupportsStreaming = tempObj.SupportsStreaming
	inputMessageVideo.Caption = tempObj.Caption
	inputMessageVideo.Ttl = tempObj.Ttl

	fieldVideo, _ := unmarshalInputFile(objMap["video"])
	inputMessageVideo.Video = fieldVideo

	return nil
}

// GetInputMessageContentEnum return the enum type of this object
func (inputMessageVideo *InputMessageVideo) GetInputMessageContentEnum() InputMessageContentEnum {
	return InputMessageVideoType
}

// InputMessageVideoNote A video note message
type InputMessageVideoNote struct {
	tdCommon
	VideoNote InputFile       `json:"video_note"` // Video note to be sent
	Thumbnail *InputThumbnail `json:"thumbnail"`  // Video thumbnail; pass null to skip thumbnail uploading
	Duration  int32           `json:"duration"`   // Duration of the video, in seconds
	Length    int32           `json:"length"`     // Video width and height; must be positive and not greater than 640
}

// MessageType return the string telegram-type of InputMessageVideoNote
func (inputMessageVideoNote *InputMessageVideoNote) MessageType() string {
	return "inputMessageVideoNote"
}

// NewInputMessageVideoNote creates a new InputMessageVideoNote
//
// @param videoNote Video note to be sent
// @param thumbnail Video thumbnail; pass null to skip thumbnail uploading
// @param duration Duration of the video, in seconds
// @param length Video width and height; must be positive and not greater than 640
func NewInputMessageVideoNote(videoNote InputFile, thumbnail *InputThumbnail, duration int32, length int32) *InputMessageVideoNote {
	inputMessageVideoNoteTemp := InputMessageVideoNote{
		tdCommon:  tdCommon{Type: "inputMessageVideoNote"},
		VideoNote: videoNote,
		Thumbnail: thumbnail,
		Duration:  duration,
		Length:    length,
	}

	return &inputMessageVideoNoteTemp
}

// UnmarshalJSON unmarshal to json
func (inputMessageVideoNote *InputMessageVideoNote) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	tempObj := struct {
		tdCommon
		Thumbnail *InputThumbnail `json:"thumbnail"` // Video thumbnail; pass null to skip thumbnail uploading
		Duration  int32           `json:"duration"`  // Duration of the video, in seconds
		Length    int32           `json:"length"`    // Video width and height; must be positive and not greater than 640
	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	inputMessageVideoNote.tdCommon = tempObj.tdCommon
	inputMessageVideoNote.Thumbnail = tempObj.Thumbnail
	inputMessageVideoNote.Duration = tempObj.Duration
	inputMessageVideoNote.Length = tempObj.Length

	fieldVideoNote, _ := unmarshalInputFile(objMap["video_note"])
	inputMessageVideoNote.VideoNote = fieldVideoNote

	return nil
}

// GetInputMessageContentEnum return the enum type of this object
func (inputMessageVideoNote *InputMessageVideoNote) GetInputMessageContentEnum() InputMessageContentEnum {
	return InputMessageVideoNoteType
}

// InputMessageVoiceNote A voice note message
type InputMessageVoiceNote struct {
	tdCommon
	VoiceNote InputFile      `json:"voice_note"` // Voice note to be sent
	Duration  int32          `json:"duration"`   // Duration of the voice note, in seconds
	Waveform  []byte         `json:"waveform"`   // Waveform representation of the voice note, in 5-bit format
	Caption   *FormattedText `json:"caption"`    // Voice note caption; pass null to use an empty caption; 0-GetOption("message_caption_length_max") characters
}

// MessageType return the string telegram-type of InputMessageVoiceNote
func (inputMessageVoiceNote *InputMessageVoiceNote) MessageType() string {
	return "inputMessageVoiceNote"
}

// NewInputMessageVoiceNote creates a new InputMessageVoiceNote
//
// @param voiceNote Voice note to be sent
// @param duration Duration of the voice note, in seconds
// @param waveform Waveform representation of the voice note, in 5-bit format
// @param caption Voice note caption; pass null to use an empty caption; 0-GetOption("message_caption_length_max") characters
func NewInputMessageVoiceNote(voiceNote InputFile, duration int32, waveform []byte, caption *FormattedText) *InputMessageVoiceNote {
	inputMessageVoiceNoteTemp := InputMessageVoiceNote{
		tdCommon:  tdCommon{Type: "inputMessageVoiceNote"},
		VoiceNote: voiceNote,
		Duration:  duration,
		Waveform:  waveform,
		Caption:   caption,
	}

	return &inputMessageVoiceNoteTemp
}

// UnmarshalJSON unmarshal to json
func (inputMessageVoiceNote *InputMessageVoiceNote) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	tempObj := struct {
		tdCommon
		Duration int32          `json:"duration"` // Duration of the voice note, in seconds
		Waveform []byte         `json:"waveform"` // Waveform representation of the voice note, in 5-bit format
		Caption  *FormattedText `json:"caption"`  // Voice note caption; pass null to use an empty caption; 0-GetOption("message_caption_length_max") characters
	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	inputMessageVoiceNote.tdCommon = tempObj.tdCommon
	inputMessageVoiceNote.Duration = tempObj.Duration
	inputMessageVoiceNote.Waveform = tempObj.Waveform
	inputMessageVoiceNote.Caption = tempObj.Caption

	fieldVoiceNote, _ := unmarshalInputFile(objMap["voice_note"])
	inputMessageVoiceNote.VoiceNote = fieldVoiceNote

	return nil
}

// GetInputMessageContentEnum return the enum type of this object
func (inputMessageVoiceNote *InputMessageVoiceNote) GetInputMessageContentEnum() InputMessageContentEnum {
	return InputMessageVoiceNoteType
}

// InputMessageLocation A message with a location
type InputMessageLocation struct {
	tdCommon
	Location             *Location `json:"location"`               // Location to be sent
	LivePeriod           int32     `json:"live_period"`            // Period for which the location can be updated, in seconds; must be between 60 and 86400 for a live location and 0 otherwise
	Heading              int32     `json:"heading"`                // For live locations, a direction in which the location moves, in degrees; 1-360. Pass 0 if unknown
	ProximityAlertRadius int32     `json:"proximity_alert_radius"` // For live locations, a maximum distance to another chat member for proximity alerts, in meters (0-100000). Pass 0 if the notification is disabled. Can't be enabled in channels and Saved Messages
}

// MessageType return the string telegram-type of InputMessageLocation
func (inputMessageLocation *InputMessageLocation) MessageType() string {
	return "inputMessageLocation"
}

// NewInputMessageLocation creates a new InputMessageLocation
//
// @param location Location to be sent
// @param livePeriod Period for which the location can be updated, in seconds; must be between 60 and 86400 for a live location and 0 otherwise
// @param heading For live locations, a direction in which the location moves, in degrees; 1-360. Pass 0 if unknown
// @param proximityAlertRadius For live locations, a maximum distance to another chat member for proximity alerts, in meters (0-100000). Pass 0 if the notification is disabled. Can't be enabled in channels and Saved Messages
func NewInputMessageLocation(location *Location, livePeriod int32, heading int32, proximityAlertRadius int32) *InputMessageLocation {
	inputMessageLocationTemp := InputMessageLocation{
		tdCommon:             tdCommon{Type: "inputMessageLocation"},
		Location:             location,
		LivePeriod:           livePeriod,
		Heading:              heading,
		ProximityAlertRadius: proximityAlertRadius,
	}

	return &inputMessageLocationTemp
}

// GetInputMessageContentEnum return the enum type of this object
func (inputMessageLocation *InputMessageLocation) GetInputMessageContentEnum() InputMessageContentEnum {
	return InputMessageLocationType
}

// InputMessageVenue A message with information about a venue
type InputMessageVenue struct {
	tdCommon
	Venue *Venue `json:"venue"` // Venue to send
}

// MessageType return the string telegram-type of InputMessageVenue
func (inputMessageVenue *InputMessageVenue) MessageType() string {
	return "inputMessageVenue"
}

// NewInputMessageVenue creates a new InputMessageVenue
//
// @param venue Venue to send
func NewInputMessageVenue(venue *Venue) *InputMessageVenue {
	inputMessageVenueTemp := InputMessageVenue{
		tdCommon: tdCommon{Type: "inputMessageVenue"},
		Venue:    venue,
	}

	return &inputMessageVenueTemp
}

// GetInputMessageContentEnum return the enum type of this object
func (inputMessageVenue *InputMessageVenue) GetInputMessageContentEnum() InputMessageContentEnum {
	return InputMessageVenueType
}

// InputMessageContact A message containing a user contact
type InputMessageContact struct {
	tdCommon
	Contact *Contact `json:"contact"` // Contact to send
}

// MessageType return the string telegram-type of InputMessageContact
func (inputMessageContact *InputMessageContact) MessageType() string {
	return "inputMessageContact"
}

// NewInputMessageContact creates a new InputMessageContact
//
// @param contact Contact to send
func NewInputMessageContact(contact *Contact) *InputMessageContact {
	inputMessageContactTemp := InputMessageContact{
		tdCommon: tdCommon{Type: "inputMessageContact"},
		Contact:  contact,
	}

	return &inputMessageContactTemp
}

// GetInputMessageContentEnum return the enum type of this object
func (inputMessageContact *InputMessageContact) GetInputMessageContentEnum() InputMessageContentEnum {
	return InputMessageContactType
}

// InputMessageDice A dice message
type InputMessageDice struct {
	tdCommon
	Emoji      string `json:"emoji"`       // Emoji on which the dice throw animation is based
	ClearDraft bool   `json:"clear_draft"` // True, if the chat message draft must be deleted
}

// MessageType return the string telegram-type of InputMessageDice
func (inputMessageDice *InputMessageDice) MessageType() string {
	return "inputMessageDice"
}

// NewInputMessageDice creates a new InputMessageDice
//
// @param emoji Emoji on which the dice throw animation is based
// @param clearDraft True, if the chat message draft must be deleted
func NewInputMessageDice(emoji string, clearDraft bool) *InputMessageDice {
	inputMessageDiceTemp := InputMessageDice{
		tdCommon:   tdCommon{Type: "inputMessageDice"},
		Emoji:      emoji,
		ClearDraft: clearDraft,
	}

	return &inputMessageDiceTemp
}

// GetInputMessageContentEnum return the enum type of this object
func (inputMessageDice *InputMessageDice) GetInputMessageContentEnum() InputMessageContentEnum {
	return InputMessageDiceType
}

// InputMessageGame A message with a game; not supported for channels or secret chats
type InputMessageGame struct {
	tdCommon
	BotUserId     int64  `json:"bot_user_id"`     // User identifier of the bot that owns the game
	GameShortName string `json:"game_short_name"` // Short name of the game
}

// MessageType return the string telegram-type of InputMessageGame
func (inputMessageGame *InputMessageGame) MessageType() string {
	return "inputMessageGame"
}

// NewInputMessageGame creates a new InputMessageGame
//
// @param botUserId User identifier of the bot that owns the game
// @param gameShortName Short name of the game
func NewInputMessageGame(botUserId int64, gameShortName string) *InputMessageGame {
	inputMessageGameTemp := InputMessageGame{
		tdCommon:      tdCommon{Type: "inputMessageGame"},
		BotUserId:     botUserId,
		GameShortName: gameShortName,
	}

	return &inputMessageGameTemp
}

// GetInputMessageContentEnum return the enum type of this object
func (inputMessageGame *InputMessageGame) GetInputMessageContentEnum() InputMessageContentEnum {
	return InputMessageGameType
}

// InputMessageInvoice A message with an invoice; can be used only by bots
type InputMessageInvoice struct {
	tdCommon
	Invoice        *Invoice `json:"invoice"`         // Invoice
	Title          string   `json:"title"`           // Product title; 1-32 characters
	Description    string   `json:"description"`     // Product description; 0-255 characters
	PhotoUrl       string   `json:"photo_url"`       // Product photo URL; optional
	PhotoSize      int32    `json:"photo_size"`      // Product photo size
	PhotoWidth     int32    `json:"photo_width"`     // Product photo width
	PhotoHeight    int32    `json:"photo_height"`    // Product photo height
	Payload        []byte   `json:"payload"`         // The invoice payload
	ProviderToken  string   `json:"provider_token"`  // Payment provider token
	ProviderData   string   `json:"provider_data"`   // JSON-encoded data about the invoice, which will be shared with the payment provider
	StartParameter string   `json:"start_parameter"` // Unique invoice bot deep link parameter for the generation of this invoice. If empty, it would be possible to pay directly from forwards of the invoice message
}

// MessageType return the string telegram-type of InputMessageInvoice
func (inputMessageInvoice *InputMessageInvoice) MessageType() string {
	return "inputMessageInvoice"
}

// NewInputMessageInvoice creates a new InputMessageInvoice
//
// @param invoice Invoice
// @param title Product title; 1-32 characters
// @param description Product description; 0-255 characters
// @param photoUrl Product photo URL; optional
// @param photoSize Product photo size
// @param photoWidth Product photo width
// @param photoHeight Product photo height
// @param payload The invoice payload
// @param providerToken Payment provider token
// @param providerData JSON-encoded data about the invoice, which will be shared with the payment provider
// @param startParameter Unique invoice bot deep link parameter for the generation of this invoice. If empty, it would be possible to pay directly from forwards of the invoice message
func NewInputMessageInvoice(invoice *Invoice, title string, description string, photoUrl string, photoSize int32, photoWidth int32, photoHeight int32, payload []byte, providerToken string, providerData string, startParameter string) *InputMessageInvoice {
	inputMessageInvoiceTemp := InputMessageInvoice{
		tdCommon:       tdCommon{Type: "inputMessageInvoice"},
		Invoice:        invoice,
		Title:          title,
		Description:    description,
		PhotoUrl:       photoUrl,
		PhotoSize:      photoSize,
		PhotoWidth:     photoWidth,
		PhotoHeight:    photoHeight,
		Payload:        payload,
		ProviderToken:  providerToken,
		ProviderData:   providerData,
		StartParameter: startParameter,
	}

	return &inputMessageInvoiceTemp
}

// GetInputMessageContentEnum return the enum type of this object
func (inputMessageInvoice *InputMessageInvoice) GetInputMessageContentEnum() InputMessageContentEnum {
	return InputMessageInvoiceType
}

// InputMessagePoll A message with a poll. Polls can't be sent to secret chats. Polls can be sent only to a private chat with a bot
type InputMessagePoll struct {
	tdCommon
	Question    string   `json:"question"`     // Poll question; 1-255 characters (up to 300 characters for bots)
	Options     []string `json:"options"`      // List of poll answer options, 2-10 strings 1-100 characters each
	IsAnonymous bool     `json:"is_anonymous"` // True, if the poll voters are anonymous. Non-anonymous polls can't be sent or forwarded to channels
	Type        PollType `json:"type"`         // Type of the poll
	OpenPeriod  int32    `json:"open_period"`  // Amount of time the poll will be active after creation, in seconds; for bots only
	CloseDate   int32    `json:"close_date"`   // Point in time (Unix timestamp) when the poll will automatically be closed; for bots only
	IsClosed    bool     `json:"is_closed"`    // True, if the poll needs to be sent already closed; for bots only
}

// MessageType return the string telegram-type of InputMessagePoll
func (inputMessagePoll *InputMessagePoll) MessageType() string {
	return "inputMessagePoll"
}

// NewInputMessagePoll creates a new InputMessagePoll
//
// @param question Poll question; 1-255 characters (up to 300 characters for bots)
// @param options List of poll answer options, 2-10 strings 1-100 characters each
// @param isAnonymous True, if the poll voters are anonymous. Non-anonymous polls can't be sent or forwarded to channels
// @param typeParam Type of the poll
// @param openPeriod Amount of time the poll will be active after creation, in seconds; for bots only
// @param closeDate Point in time (Unix timestamp) when the poll will automatically be closed; for bots only
// @param isClosed True, if the poll needs to be sent already closed; for bots only
func NewInputMessagePoll(question string, options []string, isAnonymous bool, typeParam PollType, openPeriod int32, closeDate int32, isClosed bool) *InputMessagePoll {
	inputMessagePollTemp := InputMessagePoll{
		tdCommon:    tdCommon{Type: "inputMessagePoll"},
		Question:    question,
		Options:     options,
		IsAnonymous: isAnonymous,
		Type:        typeParam,
		OpenPeriod:  openPeriod,
		CloseDate:   closeDate,
		IsClosed:    isClosed,
	}

	return &inputMessagePollTemp
}

// UnmarshalJSON unmarshal to json
func (inputMessagePoll *InputMessagePoll) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	tempObj := struct {
		tdCommon
		Question    string   `json:"question"`     // Poll question; 1-255 characters (up to 300 characters for bots)
		Options     []string `json:"options"`      // List of poll answer options, 2-10 strings 1-100 characters each
		IsAnonymous bool     `json:"is_anonymous"` // True, if the poll voters are anonymous. Non-anonymous polls can't be sent or forwarded to channels
		OpenPeriod  int32    `json:"open_period"`  // Amount of time the poll will be active after creation, in seconds; for bots only
		CloseDate   int32    `json:"close_date"`   // Point in time (Unix timestamp) when the poll will automatically be closed; for bots only
		IsClosed    bool     `json:"is_closed"`    // True, if the poll needs to be sent already closed; for bots only
	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	inputMessagePoll.tdCommon = tempObj.tdCommon
	inputMessagePoll.Question = tempObj.Question
	inputMessagePoll.Options = tempObj.Options
	inputMessagePoll.IsAnonymous = tempObj.IsAnonymous
	inputMessagePoll.OpenPeriod = tempObj.OpenPeriod
	inputMessagePoll.CloseDate = tempObj.CloseDate
	inputMessagePoll.IsClosed = tempObj.IsClosed

	fieldType, _ := unmarshalPollType(objMap["type"])
	inputMessagePoll.Type = fieldType

	return nil
}

// GetInputMessageContentEnum return the enum type of this object
func (inputMessagePoll *InputMessagePoll) GetInputMessageContentEnum() InputMessageContentEnum {
	return InputMessagePollType
}

// InputMessageForwarded A forwarded message
type InputMessageForwarded struct {
	tdCommon
	FromChatId  int64               `json:"from_chat_id"`  // Identifier for the chat this forwarded message came from
	MessageId   int64               `json:"message_id"`    // Identifier of the message to forward
	InGameShare bool                `json:"in_game_share"` // True, if a game message is being shared from a launched game; applies only to game messages
	CopyOptions *MessageCopyOptions `json:"copy_options"`  // Options to be used to copy content of the message without reference to the original sender; pass null to forward the message as usual
}

// MessageType return the string telegram-type of InputMessageForwarded
func (inputMessageForwarded *InputMessageForwarded) MessageType() string {
	return "inputMessageForwarded"
}

// NewInputMessageForwarded creates a new InputMessageForwarded
//
// @param fromChatId Identifier for the chat this forwarded message came from
// @param messageId Identifier of the message to forward
// @param inGameShare True, if a game message is being shared from a launched game; applies only to game messages
// @param copyOptions Options to be used to copy content of the message without reference to the original sender; pass null to forward the message as usual
func NewInputMessageForwarded(fromChatId int64, messageId int64, inGameShare bool, copyOptions *MessageCopyOptions) *InputMessageForwarded {
	inputMessageForwardedTemp := InputMessageForwarded{
		tdCommon:    tdCommon{Type: "inputMessageForwarded"},
		FromChatId:  fromChatId,
		MessageId:   messageId,
		InGameShare: inGameShare,
		CopyOptions: copyOptions,
	}

	return &inputMessageForwardedTemp
}

// GetInputMessageContentEnum return the enum type of this object
func (inputMessageForwarded *InputMessageForwarded) GetInputMessageContentEnum() InputMessageContentEnum {
	return InputMessageForwardedType
}
