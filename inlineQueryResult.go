package tdlib

import (
	"encoding/json"
	"fmt"
)

// InlineQueryResult Represents a single result of an inline query
type InlineQueryResult interface {
	GetInlineQueryResultEnum() InlineQueryResultEnum
}

// InlineQueryResultEnum Alias for abstract InlineQueryResult 'Sub-Classes', used as constant-enum here
type InlineQueryResultEnum string

// InlineQueryResult enums
const (
	InlineQueryResultArticleType   InlineQueryResultEnum = "inlineQueryResultArticle"
	InlineQueryResultContactType   InlineQueryResultEnum = "inlineQueryResultContact"
	InlineQueryResultLocationType  InlineQueryResultEnum = "inlineQueryResultLocation"
	InlineQueryResultVenueType     InlineQueryResultEnum = "inlineQueryResultVenue"
	InlineQueryResultGameType      InlineQueryResultEnum = "inlineQueryResultGame"
	InlineQueryResultAnimationType InlineQueryResultEnum = "inlineQueryResultAnimation"
	InlineQueryResultAudioType     InlineQueryResultEnum = "inlineQueryResultAudio"
	InlineQueryResultDocumentType  InlineQueryResultEnum = "inlineQueryResultDocument"
	InlineQueryResultPhotoType     InlineQueryResultEnum = "inlineQueryResultPhoto"
	InlineQueryResultStickerType   InlineQueryResultEnum = "inlineQueryResultSticker"
	InlineQueryResultVideoType     InlineQueryResultEnum = "inlineQueryResultVideo"
	InlineQueryResultVoiceNoteType InlineQueryResultEnum = "inlineQueryResultVoiceNote"
)

func unmarshalInlineQueryResult(rawMsg *json.RawMessage) (InlineQueryResult, error) {

	if rawMsg == nil {
		return nil, nil
	}
	var objMap map[string]interface{}
	err := json.Unmarshal(*rawMsg, &objMap)
	if err != nil {
		return nil, err
	}

	switch InlineQueryResultEnum(objMap["@type"].(string)) {
	case InlineQueryResultArticleType:
		var inlineQueryResultArticle InlineQueryResultArticle
		err := json.Unmarshal(*rawMsg, &inlineQueryResultArticle)
		return &inlineQueryResultArticle, err

	case InlineQueryResultContactType:
		var inlineQueryResultContact InlineQueryResultContact
		err := json.Unmarshal(*rawMsg, &inlineQueryResultContact)
		return &inlineQueryResultContact, err

	case InlineQueryResultLocationType:
		var inlineQueryResultLocation InlineQueryResultLocation
		err := json.Unmarshal(*rawMsg, &inlineQueryResultLocation)
		return &inlineQueryResultLocation, err

	case InlineQueryResultVenueType:
		var inlineQueryResultVenue InlineQueryResultVenue
		err := json.Unmarshal(*rawMsg, &inlineQueryResultVenue)
		return &inlineQueryResultVenue, err

	case InlineQueryResultGameType:
		var inlineQueryResultGame InlineQueryResultGame
		err := json.Unmarshal(*rawMsg, &inlineQueryResultGame)
		return &inlineQueryResultGame, err

	case InlineQueryResultAnimationType:
		var inlineQueryResultAnimation InlineQueryResultAnimation
		err := json.Unmarshal(*rawMsg, &inlineQueryResultAnimation)
		return &inlineQueryResultAnimation, err

	case InlineQueryResultAudioType:
		var inlineQueryResultAudio InlineQueryResultAudio
		err := json.Unmarshal(*rawMsg, &inlineQueryResultAudio)
		return &inlineQueryResultAudio, err

	case InlineQueryResultDocumentType:
		var inlineQueryResultDocument InlineQueryResultDocument
		err := json.Unmarshal(*rawMsg, &inlineQueryResultDocument)
		return &inlineQueryResultDocument, err

	case InlineQueryResultPhotoType:
		var inlineQueryResultPhoto InlineQueryResultPhoto
		err := json.Unmarshal(*rawMsg, &inlineQueryResultPhoto)
		return &inlineQueryResultPhoto, err

	case InlineQueryResultStickerType:
		var inlineQueryResultSticker InlineQueryResultSticker
		err := json.Unmarshal(*rawMsg, &inlineQueryResultSticker)
		return &inlineQueryResultSticker, err

	case InlineQueryResultVideoType:
		var inlineQueryResultVideo InlineQueryResultVideo
		err := json.Unmarshal(*rawMsg, &inlineQueryResultVideo)
		return &inlineQueryResultVideo, err

	case InlineQueryResultVoiceNoteType:
		var inlineQueryResultVoiceNote InlineQueryResultVoiceNote
		err := json.Unmarshal(*rawMsg, &inlineQueryResultVoiceNote)
		return &inlineQueryResultVoiceNote, err

	default:
		return nil, fmt.Errorf("Error UnMarshaling, unknown type:" + objMap["@type"].(string))
	}
}

// InlineQueryResultArticle Represents a link to an article or web page
type InlineQueryResultArticle struct {
	tdCommon
	Id          string     `json:"id"`          // Unique identifier of the query result
	Url         string     `json:"url"`         // URL of the result, if it exists
	HideUrl     bool       `json:"hide_url"`    // True, if the URL must be not shown
	Title       string     `json:"title"`       // Title of the result
	Description string     `json:"description"` // A short description of the result
	Thumbnail   *Thumbnail `json:"thumbnail"`   // Result thumbnail in JPEG format; may be null
}

// MessageType return the string telegram-type of InlineQueryResultArticle
func (inlineQueryResultArticle *InlineQueryResultArticle) MessageType() string {
	return "inlineQueryResultArticle"
}

// NewInlineQueryResultArticle creates a new InlineQueryResultArticle
//
// @param id Unique identifier of the query result
// @param url URL of the result, if it exists
// @param hideUrl True, if the URL must be not shown
// @param title Title of the result
// @param description A short description of the result
// @param thumbnail Result thumbnail in JPEG format; may be null
func NewInlineQueryResultArticle(id string, url string, hideUrl bool, title string, description string, thumbnail *Thumbnail) *InlineQueryResultArticle {
	inlineQueryResultArticleTemp := InlineQueryResultArticle{
		tdCommon:    tdCommon{Type: "inlineQueryResultArticle"},
		Id:          id,
		Url:         url,
		HideUrl:     hideUrl,
		Title:       title,
		Description: description,
		Thumbnail:   thumbnail,
	}

	return &inlineQueryResultArticleTemp
}

// GetInlineQueryResultEnum return the enum type of this object
func (inlineQueryResultArticle *InlineQueryResultArticle) GetInlineQueryResultEnum() InlineQueryResultEnum {
	return InlineQueryResultArticleType
}

// InlineQueryResultContact Represents a user contact
type InlineQueryResultContact struct {
	tdCommon
	Id        string     `json:"id"`        // Unique identifier of the query result
	Contact   *Contact   `json:"contact"`   // A user contact
	Thumbnail *Thumbnail `json:"thumbnail"` // Result thumbnail in JPEG format; may be null
}

// MessageType return the string telegram-type of InlineQueryResultContact
func (inlineQueryResultContact *InlineQueryResultContact) MessageType() string {
	return "inlineQueryResultContact"
}

// NewInlineQueryResultContact creates a new InlineQueryResultContact
//
// @param id Unique identifier of the query result
// @param contact A user contact
// @param thumbnail Result thumbnail in JPEG format; may be null
func NewInlineQueryResultContact(id string, contact *Contact, thumbnail *Thumbnail) *InlineQueryResultContact {
	inlineQueryResultContactTemp := InlineQueryResultContact{
		tdCommon:  tdCommon{Type: "inlineQueryResultContact"},
		Id:        id,
		Contact:   contact,
		Thumbnail: thumbnail,
	}

	return &inlineQueryResultContactTemp
}

// GetInlineQueryResultEnum return the enum type of this object
func (inlineQueryResultContact *InlineQueryResultContact) GetInlineQueryResultEnum() InlineQueryResultEnum {
	return InlineQueryResultContactType
}

// InlineQueryResultLocation Represents a point on the map
type InlineQueryResultLocation struct {
	tdCommon
	Id        string     `json:"id"`        // Unique identifier of the query result
	Location  *Location  `json:"location"`  // Location result
	Title     string     `json:"title"`     // Title of the result
	Thumbnail *Thumbnail `json:"thumbnail"` // Result thumbnail in JPEG format; may be null
}

// MessageType return the string telegram-type of InlineQueryResultLocation
func (inlineQueryResultLocation *InlineQueryResultLocation) MessageType() string {
	return "inlineQueryResultLocation"
}

// NewInlineQueryResultLocation creates a new InlineQueryResultLocation
//
// @param id Unique identifier of the query result
// @param location Location result
// @param title Title of the result
// @param thumbnail Result thumbnail in JPEG format; may be null
func NewInlineQueryResultLocation(id string, location *Location, title string, thumbnail *Thumbnail) *InlineQueryResultLocation {
	inlineQueryResultLocationTemp := InlineQueryResultLocation{
		tdCommon:  tdCommon{Type: "inlineQueryResultLocation"},
		Id:        id,
		Location:  location,
		Title:     title,
		Thumbnail: thumbnail,
	}

	return &inlineQueryResultLocationTemp
}

// GetInlineQueryResultEnum return the enum type of this object
func (inlineQueryResultLocation *InlineQueryResultLocation) GetInlineQueryResultEnum() InlineQueryResultEnum {
	return InlineQueryResultLocationType
}

// InlineQueryResultVenue Represents information about a venue
type InlineQueryResultVenue struct {
	tdCommon
	Id        string     `json:"id"`        // Unique identifier of the query result
	Venue     *Venue     `json:"venue"`     // Venue result
	Thumbnail *Thumbnail `json:"thumbnail"` // Result thumbnail in JPEG format; may be null
}

// MessageType return the string telegram-type of InlineQueryResultVenue
func (inlineQueryResultVenue *InlineQueryResultVenue) MessageType() string {
	return "inlineQueryResultVenue"
}

// NewInlineQueryResultVenue creates a new InlineQueryResultVenue
//
// @param id Unique identifier of the query result
// @param venue Venue result
// @param thumbnail Result thumbnail in JPEG format; may be null
func NewInlineQueryResultVenue(id string, venue *Venue, thumbnail *Thumbnail) *InlineQueryResultVenue {
	inlineQueryResultVenueTemp := InlineQueryResultVenue{
		tdCommon:  tdCommon{Type: "inlineQueryResultVenue"},
		Id:        id,
		Venue:     venue,
		Thumbnail: thumbnail,
	}

	return &inlineQueryResultVenueTemp
}

// GetInlineQueryResultEnum return the enum type of this object
func (inlineQueryResultVenue *InlineQueryResultVenue) GetInlineQueryResultEnum() InlineQueryResultEnum {
	return InlineQueryResultVenueType
}

// InlineQueryResultGame Represents information about a game
type InlineQueryResultGame struct {
	tdCommon
	Id   string `json:"id"`   // Unique identifier of the query result
	Game *Game  `json:"game"` // Game result
}

// MessageType return the string telegram-type of InlineQueryResultGame
func (inlineQueryResultGame *InlineQueryResultGame) MessageType() string {
	return "inlineQueryResultGame"
}

// NewInlineQueryResultGame creates a new InlineQueryResultGame
//
// @param id Unique identifier of the query result
// @param game Game result
func NewInlineQueryResultGame(id string, game *Game) *InlineQueryResultGame {
	inlineQueryResultGameTemp := InlineQueryResultGame{
		tdCommon: tdCommon{Type: "inlineQueryResultGame"},
		Id:       id,
		Game:     game,
	}

	return &inlineQueryResultGameTemp
}

// GetInlineQueryResultEnum return the enum type of this object
func (inlineQueryResultGame *InlineQueryResultGame) GetInlineQueryResultEnum() InlineQueryResultEnum {
	return InlineQueryResultGameType
}

// InlineQueryResultAnimation Represents an animation file
type InlineQueryResultAnimation struct {
	tdCommon
	Id        string     `json:"id"`        // Unique identifier of the query result
	Animation *Animation `json:"animation"` // Animation file
	Title     string     `json:"title"`     // Animation title
}

// MessageType return the string telegram-type of InlineQueryResultAnimation
func (inlineQueryResultAnimation *InlineQueryResultAnimation) MessageType() string {
	return "inlineQueryResultAnimation"
}

// NewInlineQueryResultAnimation creates a new InlineQueryResultAnimation
//
// @param id Unique identifier of the query result
// @param animation Animation file
// @param title Animation title
func NewInlineQueryResultAnimation(id string, animation *Animation, title string) *InlineQueryResultAnimation {
	inlineQueryResultAnimationTemp := InlineQueryResultAnimation{
		tdCommon:  tdCommon{Type: "inlineQueryResultAnimation"},
		Id:        id,
		Animation: animation,
		Title:     title,
	}

	return &inlineQueryResultAnimationTemp
}

// GetInlineQueryResultEnum return the enum type of this object
func (inlineQueryResultAnimation *InlineQueryResultAnimation) GetInlineQueryResultEnum() InlineQueryResultEnum {
	return InlineQueryResultAnimationType
}

// InlineQueryResultAudio Represents an audio file
type InlineQueryResultAudio struct {
	tdCommon
	Id    string `json:"id"`    // Unique identifier of the query result
	Audio *Audio `json:"audio"` // Audio file
}

// MessageType return the string telegram-type of InlineQueryResultAudio
func (inlineQueryResultAudio *InlineQueryResultAudio) MessageType() string {
	return "inlineQueryResultAudio"
}

// NewInlineQueryResultAudio creates a new InlineQueryResultAudio
//
// @param id Unique identifier of the query result
// @param audio Audio file
func NewInlineQueryResultAudio(id string, audio *Audio) *InlineQueryResultAudio {
	inlineQueryResultAudioTemp := InlineQueryResultAudio{
		tdCommon: tdCommon{Type: "inlineQueryResultAudio"},
		Id:       id,
		Audio:    audio,
	}

	return &inlineQueryResultAudioTemp
}

// GetInlineQueryResultEnum return the enum type of this object
func (inlineQueryResultAudio *InlineQueryResultAudio) GetInlineQueryResultEnum() InlineQueryResultEnum {
	return InlineQueryResultAudioType
}

// InlineQueryResultDocument Represents a document
type InlineQueryResultDocument struct {
	tdCommon
	Id          string    `json:"id"`          // Unique identifier of the query result
	Document    *Document `json:"document"`    // Document
	Title       string    `json:"title"`       // Document title
	Description string    `json:"description"` // Document description
}

// MessageType return the string telegram-type of InlineQueryResultDocument
func (inlineQueryResultDocument *InlineQueryResultDocument) MessageType() string {
	return "inlineQueryResultDocument"
}

// NewInlineQueryResultDocument creates a new InlineQueryResultDocument
//
// @param id Unique identifier of the query result
// @param document Document
// @param title Document title
// @param description Document description
func NewInlineQueryResultDocument(id string, document *Document, title string, description string) *InlineQueryResultDocument {
	inlineQueryResultDocumentTemp := InlineQueryResultDocument{
		tdCommon:    tdCommon{Type: "inlineQueryResultDocument"},
		Id:          id,
		Document:    document,
		Title:       title,
		Description: description,
	}

	return &inlineQueryResultDocumentTemp
}

// GetInlineQueryResultEnum return the enum type of this object
func (inlineQueryResultDocument *InlineQueryResultDocument) GetInlineQueryResultEnum() InlineQueryResultEnum {
	return InlineQueryResultDocumentType
}

// InlineQueryResultPhoto Represents a photo
type InlineQueryResultPhoto struct {
	tdCommon
	Id          string `json:"id"`          // Unique identifier of the query result
	Photo       *Photo `json:"photo"`       // Photo
	Title       string `json:"title"`       // Title of the result, if known
	Description string `json:"description"` // A short description of the result, if known
}

// MessageType return the string telegram-type of InlineQueryResultPhoto
func (inlineQueryResultPhoto *InlineQueryResultPhoto) MessageType() string {
	return "inlineQueryResultPhoto"
}

// NewInlineQueryResultPhoto creates a new InlineQueryResultPhoto
//
// @param id Unique identifier of the query result
// @param photo Photo
// @param title Title of the result, if known
// @param description A short description of the result, if known
func NewInlineQueryResultPhoto(id string, photo *Photo, title string, description string) *InlineQueryResultPhoto {
	inlineQueryResultPhotoTemp := InlineQueryResultPhoto{
		tdCommon:    tdCommon{Type: "inlineQueryResultPhoto"},
		Id:          id,
		Photo:       photo,
		Title:       title,
		Description: description,
	}

	return &inlineQueryResultPhotoTemp
}

// GetInlineQueryResultEnum return the enum type of this object
func (inlineQueryResultPhoto *InlineQueryResultPhoto) GetInlineQueryResultEnum() InlineQueryResultEnum {
	return InlineQueryResultPhotoType
}

// InlineQueryResultSticker Represents a sticker
type InlineQueryResultSticker struct {
	tdCommon
	Id      string   `json:"id"`      // Unique identifier of the query result
	Sticker *Sticker `json:"sticker"` // Sticker
}

// MessageType return the string telegram-type of InlineQueryResultSticker
func (inlineQueryResultSticker *InlineQueryResultSticker) MessageType() string {
	return "inlineQueryResultSticker"
}

// NewInlineQueryResultSticker creates a new InlineQueryResultSticker
//
// @param id Unique identifier of the query result
// @param sticker Sticker
func NewInlineQueryResultSticker(id string, sticker *Sticker) *InlineQueryResultSticker {
	inlineQueryResultStickerTemp := InlineQueryResultSticker{
		tdCommon: tdCommon{Type: "inlineQueryResultSticker"},
		Id:       id,
		Sticker:  sticker,
	}

	return &inlineQueryResultStickerTemp
}

// GetInlineQueryResultEnum return the enum type of this object
func (inlineQueryResultSticker *InlineQueryResultSticker) GetInlineQueryResultEnum() InlineQueryResultEnum {
	return InlineQueryResultStickerType
}

// InlineQueryResultVideo Represents a video
type InlineQueryResultVideo struct {
	tdCommon
	Id          string `json:"id"`          // Unique identifier of the query result
	Video       *Video `json:"video"`       // Video
	Title       string `json:"title"`       // Title of the video
	Description string `json:"description"` // Description of the video
}

// MessageType return the string telegram-type of InlineQueryResultVideo
func (inlineQueryResultVideo *InlineQueryResultVideo) MessageType() string {
	return "inlineQueryResultVideo"
}

// NewInlineQueryResultVideo creates a new InlineQueryResultVideo
//
// @param id Unique identifier of the query result
// @param video Video
// @param title Title of the video
// @param description Description of the video
func NewInlineQueryResultVideo(id string, video *Video, title string, description string) *InlineQueryResultVideo {
	inlineQueryResultVideoTemp := InlineQueryResultVideo{
		tdCommon:    tdCommon{Type: "inlineQueryResultVideo"},
		Id:          id,
		Video:       video,
		Title:       title,
		Description: description,
	}

	return &inlineQueryResultVideoTemp
}

// GetInlineQueryResultEnum return the enum type of this object
func (inlineQueryResultVideo *InlineQueryResultVideo) GetInlineQueryResultEnum() InlineQueryResultEnum {
	return InlineQueryResultVideoType
}

// InlineQueryResultVoiceNote Represents a voice note
type InlineQueryResultVoiceNote struct {
	tdCommon
	Id        string     `json:"id"`         // Unique identifier of the query result
	VoiceNote *VoiceNote `json:"voice_note"` // Voice note
	Title     string     `json:"title"`      // Title of the voice note
}

// MessageType return the string telegram-type of InlineQueryResultVoiceNote
func (inlineQueryResultVoiceNote *InlineQueryResultVoiceNote) MessageType() string {
	return "inlineQueryResultVoiceNote"
}

// NewInlineQueryResultVoiceNote creates a new InlineQueryResultVoiceNote
//
// @param id Unique identifier of the query result
// @param voiceNote Voice note
// @param title Title of the voice note
func NewInlineQueryResultVoiceNote(id string, voiceNote *VoiceNote, title string) *InlineQueryResultVoiceNote {
	inlineQueryResultVoiceNoteTemp := InlineQueryResultVoiceNote{
		tdCommon:  tdCommon{Type: "inlineQueryResultVoiceNote"},
		Id:        id,
		VoiceNote: voiceNote,
		Title:     title,
	}

	return &inlineQueryResultVoiceNoteTemp
}

// GetInlineQueryResultEnum return the enum type of this object
func (inlineQueryResultVoiceNote *InlineQueryResultVoiceNote) GetInlineQueryResultEnum() InlineQueryResultEnum {
	return InlineQueryResultVoiceNoteType
}
