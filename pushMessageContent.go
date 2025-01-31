package tdlib

import (
	"encoding/json"
	"fmt"
)

// PushMessageContent Contains content of a push message notification
type PushMessageContent interface {
	GetPushMessageContentEnum() PushMessageContentEnum
}

// PushMessageContentEnum Alias for abstract PushMessageContent 'Sub-Classes', used as constant-enum here
type PushMessageContentEnum string

// PushMessageContent enums
const (
	PushMessageContentHiddenType               PushMessageContentEnum = "pushMessageContentHidden"
	PushMessageContentAnimationType            PushMessageContentEnum = "pushMessageContentAnimation"
	PushMessageContentAudioType                PushMessageContentEnum = "pushMessageContentAudio"
	PushMessageContentContactType              PushMessageContentEnum = "pushMessageContentContact"
	PushMessageContentContactRegisteredType    PushMessageContentEnum = "pushMessageContentContactRegistered"
	PushMessageContentDocumentType             PushMessageContentEnum = "pushMessageContentDocument"
	PushMessageContentGameType                 PushMessageContentEnum = "pushMessageContentGame"
	PushMessageContentGameScoreType            PushMessageContentEnum = "pushMessageContentGameScore"
	PushMessageContentInvoiceType              PushMessageContentEnum = "pushMessageContentInvoice"
	PushMessageContentLocationType             PushMessageContentEnum = "pushMessageContentLocation"
	PushMessageContentPhotoType                PushMessageContentEnum = "pushMessageContentPhoto"
	PushMessageContentPollType                 PushMessageContentEnum = "pushMessageContentPoll"
	PushMessageContentScreenshotTakenType      PushMessageContentEnum = "pushMessageContentScreenshotTaken"
	PushMessageContentStickerType              PushMessageContentEnum = "pushMessageContentSticker"
	PushMessageContentTextType                 PushMessageContentEnum = "pushMessageContentText"
	PushMessageContentVideoType                PushMessageContentEnum = "pushMessageContentVideo"
	PushMessageContentVideoNoteType            PushMessageContentEnum = "pushMessageContentVideoNote"
	PushMessageContentVoiceNoteType            PushMessageContentEnum = "pushMessageContentVoiceNote"
	PushMessageContentBasicGroupChatCreateType PushMessageContentEnum = "pushMessageContentBasicGroupChatCreate"
	PushMessageContentChatAddMembersType       PushMessageContentEnum = "pushMessageContentChatAddMembers"
	PushMessageContentChatChangePhotoType      PushMessageContentEnum = "pushMessageContentChatChangePhoto"
	PushMessageContentChatChangeTitleType      PushMessageContentEnum = "pushMessageContentChatChangeTitle"
	PushMessageContentChatSetThemeType         PushMessageContentEnum = "pushMessageContentChatSetTheme"
	PushMessageContentChatDeleteMemberType     PushMessageContentEnum = "pushMessageContentChatDeleteMember"
	PushMessageContentChatJoinByLinkType       PushMessageContentEnum = "pushMessageContentChatJoinByLink"
	PushMessageContentChatJoinByRequestType    PushMessageContentEnum = "pushMessageContentChatJoinByRequest"
	PushMessageContentMessageForwardsType      PushMessageContentEnum = "pushMessageContentMessageForwards"
	PushMessageContentMediaAlbumType           PushMessageContentEnum = "pushMessageContentMediaAlbum"
)

func unmarshalPushMessageContent(rawMsg *json.RawMessage) (PushMessageContent, error) {

	if rawMsg == nil {
		return nil, nil
	}
	var objMap map[string]interface{}
	err := json.Unmarshal(*rawMsg, &objMap)
	if err != nil {
		return nil, err
	}

	switch PushMessageContentEnum(objMap["@type"].(string)) {
	case PushMessageContentHiddenType:
		var pushMessageContentHidden PushMessageContentHidden
		err := json.Unmarshal(*rawMsg, &pushMessageContentHidden)
		return &pushMessageContentHidden, err

	case PushMessageContentAnimationType:
		var pushMessageContentAnimation PushMessageContentAnimation
		err := json.Unmarshal(*rawMsg, &pushMessageContentAnimation)
		return &pushMessageContentAnimation, err

	case PushMessageContentAudioType:
		var pushMessageContentAudio PushMessageContentAudio
		err := json.Unmarshal(*rawMsg, &pushMessageContentAudio)
		return &pushMessageContentAudio, err

	case PushMessageContentContactType:
		var pushMessageContentContact PushMessageContentContact
		err := json.Unmarshal(*rawMsg, &pushMessageContentContact)
		return &pushMessageContentContact, err

	case PushMessageContentContactRegisteredType:
		var pushMessageContentContactRegistered PushMessageContentContactRegistered
		err := json.Unmarshal(*rawMsg, &pushMessageContentContactRegistered)
		return &pushMessageContentContactRegistered, err

	case PushMessageContentDocumentType:
		var pushMessageContentDocument PushMessageContentDocument
		err := json.Unmarshal(*rawMsg, &pushMessageContentDocument)
		return &pushMessageContentDocument, err

	case PushMessageContentGameType:
		var pushMessageContentGame PushMessageContentGame
		err := json.Unmarshal(*rawMsg, &pushMessageContentGame)
		return &pushMessageContentGame, err

	case PushMessageContentGameScoreType:
		var pushMessageContentGameScore PushMessageContentGameScore
		err := json.Unmarshal(*rawMsg, &pushMessageContentGameScore)
		return &pushMessageContentGameScore, err

	case PushMessageContentInvoiceType:
		var pushMessageContentInvoice PushMessageContentInvoice
		err := json.Unmarshal(*rawMsg, &pushMessageContentInvoice)
		return &pushMessageContentInvoice, err

	case PushMessageContentLocationType:
		var pushMessageContentLocation PushMessageContentLocation
		err := json.Unmarshal(*rawMsg, &pushMessageContentLocation)
		return &pushMessageContentLocation, err

	case PushMessageContentPhotoType:
		var pushMessageContentPhoto PushMessageContentPhoto
		err := json.Unmarshal(*rawMsg, &pushMessageContentPhoto)
		return &pushMessageContentPhoto, err

	case PushMessageContentPollType:
		var pushMessageContentPoll PushMessageContentPoll
		err := json.Unmarshal(*rawMsg, &pushMessageContentPoll)
		return &pushMessageContentPoll, err

	case PushMessageContentScreenshotTakenType:
		var pushMessageContentScreenshotTaken PushMessageContentScreenshotTaken
		err := json.Unmarshal(*rawMsg, &pushMessageContentScreenshotTaken)
		return &pushMessageContentScreenshotTaken, err

	case PushMessageContentStickerType:
		var pushMessageContentSticker PushMessageContentSticker
		err := json.Unmarshal(*rawMsg, &pushMessageContentSticker)
		return &pushMessageContentSticker, err

	case PushMessageContentTextType:
		var pushMessageContentText PushMessageContentText
		err := json.Unmarshal(*rawMsg, &pushMessageContentText)
		return &pushMessageContentText, err

	case PushMessageContentVideoType:
		var pushMessageContentVideo PushMessageContentVideo
		err := json.Unmarshal(*rawMsg, &pushMessageContentVideo)
		return &pushMessageContentVideo, err

	case PushMessageContentVideoNoteType:
		var pushMessageContentVideoNote PushMessageContentVideoNote
		err := json.Unmarshal(*rawMsg, &pushMessageContentVideoNote)
		return &pushMessageContentVideoNote, err

	case PushMessageContentVoiceNoteType:
		var pushMessageContentVoiceNote PushMessageContentVoiceNote
		err := json.Unmarshal(*rawMsg, &pushMessageContentVoiceNote)
		return &pushMessageContentVoiceNote, err

	case PushMessageContentBasicGroupChatCreateType:
		var pushMessageContentBasicGroupChatCreate PushMessageContentBasicGroupChatCreate
		err := json.Unmarshal(*rawMsg, &pushMessageContentBasicGroupChatCreate)
		return &pushMessageContentBasicGroupChatCreate, err

	case PushMessageContentChatAddMembersType:
		var pushMessageContentChatAddMembers PushMessageContentChatAddMembers
		err := json.Unmarshal(*rawMsg, &pushMessageContentChatAddMembers)
		return &pushMessageContentChatAddMembers, err

	case PushMessageContentChatChangePhotoType:
		var pushMessageContentChatChangePhoto PushMessageContentChatChangePhoto
		err := json.Unmarshal(*rawMsg, &pushMessageContentChatChangePhoto)
		return &pushMessageContentChatChangePhoto, err

	case PushMessageContentChatChangeTitleType:
		var pushMessageContentChatChangeTitle PushMessageContentChatChangeTitle
		err := json.Unmarshal(*rawMsg, &pushMessageContentChatChangeTitle)
		return &pushMessageContentChatChangeTitle, err

	case PushMessageContentChatSetThemeType:
		var pushMessageContentChatSetTheme PushMessageContentChatSetTheme
		err := json.Unmarshal(*rawMsg, &pushMessageContentChatSetTheme)
		return &pushMessageContentChatSetTheme, err

	case PushMessageContentChatDeleteMemberType:
		var pushMessageContentChatDeleteMember PushMessageContentChatDeleteMember
		err := json.Unmarshal(*rawMsg, &pushMessageContentChatDeleteMember)
		return &pushMessageContentChatDeleteMember, err

	case PushMessageContentChatJoinByLinkType:
		var pushMessageContentChatJoinByLink PushMessageContentChatJoinByLink
		err := json.Unmarshal(*rawMsg, &pushMessageContentChatJoinByLink)
		return &pushMessageContentChatJoinByLink, err

	case PushMessageContentChatJoinByRequestType:
		var pushMessageContentChatJoinByRequest PushMessageContentChatJoinByRequest
		err := json.Unmarshal(*rawMsg, &pushMessageContentChatJoinByRequest)
		return &pushMessageContentChatJoinByRequest, err

	case PushMessageContentMessageForwardsType:
		var pushMessageContentMessageForwards PushMessageContentMessageForwards
		err := json.Unmarshal(*rawMsg, &pushMessageContentMessageForwards)
		return &pushMessageContentMessageForwards, err

	case PushMessageContentMediaAlbumType:
		var pushMessageContentMediaAlbum PushMessageContentMediaAlbum
		err := json.Unmarshal(*rawMsg, &pushMessageContentMediaAlbum)
		return &pushMessageContentMediaAlbum, err

	default:
		return nil, fmt.Errorf("Error UnMarshaling, unknown type:" + objMap["@type"].(string))
	}
}

// PushMessageContentHidden A general message with hidden content
type PushMessageContentHidden struct {
	tdCommon
	IsPinned bool `json:"is_pinned"` // True, if the message is a pinned message with the specified content
}

// MessageType return the string telegram-type of PushMessageContentHidden
func (pushMessageContentHidden *PushMessageContentHidden) MessageType() string {
	return "pushMessageContentHidden"
}

// NewPushMessageContentHidden creates a new PushMessageContentHidden
//
// @param isPinned True, if the message is a pinned message with the specified content
func NewPushMessageContentHidden(isPinned bool) *PushMessageContentHidden {
	pushMessageContentHiddenTemp := PushMessageContentHidden{
		tdCommon: tdCommon{Type: "pushMessageContentHidden"},
		IsPinned: isPinned,
	}

	return &pushMessageContentHiddenTemp
}

// GetPushMessageContentEnum return the enum type of this object
func (pushMessageContentHidden *PushMessageContentHidden) GetPushMessageContentEnum() PushMessageContentEnum {
	return PushMessageContentHiddenType
}

// PushMessageContentAnimation An animation message (GIF-style).
type PushMessageContentAnimation struct {
	tdCommon
	Animation *Animation `json:"animation"` // Message content; may be null
	Caption   string     `json:"caption"`   // Animation caption
	IsPinned  bool       `json:"is_pinned"` // True, if the message is a pinned message with the specified content
}

// MessageType return the string telegram-type of PushMessageContentAnimation
func (pushMessageContentAnimation *PushMessageContentAnimation) MessageType() string {
	return "pushMessageContentAnimation"
}

// NewPushMessageContentAnimation creates a new PushMessageContentAnimation
//
// @param animation Message content; may be null
// @param caption Animation caption
// @param isPinned True, if the message is a pinned message with the specified content
func NewPushMessageContentAnimation(animation *Animation, caption string, isPinned bool) *PushMessageContentAnimation {
	pushMessageContentAnimationTemp := PushMessageContentAnimation{
		tdCommon:  tdCommon{Type: "pushMessageContentAnimation"},
		Animation: animation,
		Caption:   caption,
		IsPinned:  isPinned,
	}

	return &pushMessageContentAnimationTemp
}

// GetPushMessageContentEnum return the enum type of this object
func (pushMessageContentAnimation *PushMessageContentAnimation) GetPushMessageContentEnum() PushMessageContentEnum {
	return PushMessageContentAnimationType
}

// PushMessageContentAudio An audio message
type PushMessageContentAudio struct {
	tdCommon
	Audio    *Audio `json:"audio"`     // Message content; may be null
	IsPinned bool   `json:"is_pinned"` // True, if the message is a pinned message with the specified content
}

// MessageType return the string telegram-type of PushMessageContentAudio
func (pushMessageContentAudio *PushMessageContentAudio) MessageType() string {
	return "pushMessageContentAudio"
}

// NewPushMessageContentAudio creates a new PushMessageContentAudio
//
// @param audio Message content; may be null
// @param isPinned True, if the message is a pinned message with the specified content
func NewPushMessageContentAudio(audio *Audio, isPinned bool) *PushMessageContentAudio {
	pushMessageContentAudioTemp := PushMessageContentAudio{
		tdCommon: tdCommon{Type: "pushMessageContentAudio"},
		Audio:    audio,
		IsPinned: isPinned,
	}

	return &pushMessageContentAudioTemp
}

// GetPushMessageContentEnum return the enum type of this object
func (pushMessageContentAudio *PushMessageContentAudio) GetPushMessageContentEnum() PushMessageContentEnum {
	return PushMessageContentAudioType
}

// PushMessageContentContact A message with a user contact
type PushMessageContentContact struct {
	tdCommon
	Name     string `json:"name"`      // Contact's name
	IsPinned bool   `json:"is_pinned"` // True, if the message is a pinned message with the specified content
}

// MessageType return the string telegram-type of PushMessageContentContact
func (pushMessageContentContact *PushMessageContentContact) MessageType() string {
	return "pushMessageContentContact"
}

// NewPushMessageContentContact creates a new PushMessageContentContact
//
// @param name Contact's name
// @param isPinned True, if the message is a pinned message with the specified content
func NewPushMessageContentContact(name string, isPinned bool) *PushMessageContentContact {
	pushMessageContentContactTemp := PushMessageContentContact{
		tdCommon: tdCommon{Type: "pushMessageContentContact"},
		Name:     name,
		IsPinned: isPinned,
	}

	return &pushMessageContentContactTemp
}

// GetPushMessageContentEnum return the enum type of this object
func (pushMessageContentContact *PushMessageContentContact) GetPushMessageContentEnum() PushMessageContentEnum {
	return PushMessageContentContactType
}

// PushMessageContentContactRegistered A contact has registered with Telegram
type PushMessageContentContactRegistered struct {
	tdCommon
}

// MessageType return the string telegram-type of PushMessageContentContactRegistered
func (pushMessageContentContactRegistered *PushMessageContentContactRegistered) MessageType() string {
	return "pushMessageContentContactRegistered"
}

// NewPushMessageContentContactRegistered creates a new PushMessageContentContactRegistered
//
func NewPushMessageContentContactRegistered() *PushMessageContentContactRegistered {
	pushMessageContentContactRegisteredTemp := PushMessageContentContactRegistered{
		tdCommon: tdCommon{Type: "pushMessageContentContactRegistered"},
	}

	return &pushMessageContentContactRegisteredTemp
}

// GetPushMessageContentEnum return the enum type of this object
func (pushMessageContentContactRegistered *PushMessageContentContactRegistered) GetPushMessageContentEnum() PushMessageContentEnum {
	return PushMessageContentContactRegisteredType
}

// PushMessageContentDocument A document message (a general file)
type PushMessageContentDocument struct {
	tdCommon
	Document *Document `json:"document"`  // Message content; may be null
	IsPinned bool      `json:"is_pinned"` // True, if the message is a pinned message with the specified content
}

// MessageType return the string telegram-type of PushMessageContentDocument
func (pushMessageContentDocument *PushMessageContentDocument) MessageType() string {
	return "pushMessageContentDocument"
}

// NewPushMessageContentDocument creates a new PushMessageContentDocument
//
// @param document Message content; may be null
// @param isPinned True, if the message is a pinned message with the specified content
func NewPushMessageContentDocument(document *Document, isPinned bool) *PushMessageContentDocument {
	pushMessageContentDocumentTemp := PushMessageContentDocument{
		tdCommon: tdCommon{Type: "pushMessageContentDocument"},
		Document: document,
		IsPinned: isPinned,
	}

	return &pushMessageContentDocumentTemp
}

// GetPushMessageContentEnum return the enum type of this object
func (pushMessageContentDocument *PushMessageContentDocument) GetPushMessageContentEnum() PushMessageContentEnum {
	return PushMessageContentDocumentType
}

// PushMessageContentGame A message with a game
type PushMessageContentGame struct {
	tdCommon
	Title    string `json:"title"`     // Game title, empty for pinned game message
	IsPinned bool   `json:"is_pinned"` // True, if the message is a pinned message with the specified content
}

// MessageType return the string telegram-type of PushMessageContentGame
func (pushMessageContentGame *PushMessageContentGame) MessageType() string {
	return "pushMessageContentGame"
}

// NewPushMessageContentGame creates a new PushMessageContentGame
//
// @param title Game title, empty for pinned game message
// @param isPinned True, if the message is a pinned message with the specified content
func NewPushMessageContentGame(title string, isPinned bool) *PushMessageContentGame {
	pushMessageContentGameTemp := PushMessageContentGame{
		tdCommon: tdCommon{Type: "pushMessageContentGame"},
		Title:    title,
		IsPinned: isPinned,
	}

	return &pushMessageContentGameTemp
}

// GetPushMessageContentEnum return the enum type of this object
func (pushMessageContentGame *PushMessageContentGame) GetPushMessageContentEnum() PushMessageContentEnum {
	return PushMessageContentGameType
}

// PushMessageContentGameScore A new high score was achieved in a game
type PushMessageContentGameScore struct {
	tdCommon
	Title    string `json:"title"`     // Game title, empty for pinned message
	Score    int32  `json:"score"`     // New score, 0 for pinned message
	IsPinned bool   `json:"is_pinned"` // True, if the message is a pinned message with the specified content
}

// MessageType return the string telegram-type of PushMessageContentGameScore
func (pushMessageContentGameScore *PushMessageContentGameScore) MessageType() string {
	return "pushMessageContentGameScore"
}

// NewPushMessageContentGameScore creates a new PushMessageContentGameScore
//
// @param title Game title, empty for pinned message
// @param score New score, 0 for pinned message
// @param isPinned True, if the message is a pinned message with the specified content
func NewPushMessageContentGameScore(title string, score int32, isPinned bool) *PushMessageContentGameScore {
	pushMessageContentGameScoreTemp := PushMessageContentGameScore{
		tdCommon: tdCommon{Type: "pushMessageContentGameScore"},
		Title:    title,
		Score:    score,
		IsPinned: isPinned,
	}

	return &pushMessageContentGameScoreTemp
}

// GetPushMessageContentEnum return the enum type of this object
func (pushMessageContentGameScore *PushMessageContentGameScore) GetPushMessageContentEnum() PushMessageContentEnum {
	return PushMessageContentGameScoreType
}

// PushMessageContentInvoice A message with an invoice from a bot
type PushMessageContentInvoice struct {
	tdCommon
	Price    string `json:"price"`     // Product price
	IsPinned bool   `json:"is_pinned"` // True, if the message is a pinned message with the specified content
}

// MessageType return the string telegram-type of PushMessageContentInvoice
func (pushMessageContentInvoice *PushMessageContentInvoice) MessageType() string {
	return "pushMessageContentInvoice"
}

// NewPushMessageContentInvoice creates a new PushMessageContentInvoice
//
// @param price Product price
// @param isPinned True, if the message is a pinned message with the specified content
func NewPushMessageContentInvoice(price string, isPinned bool) *PushMessageContentInvoice {
	pushMessageContentInvoiceTemp := PushMessageContentInvoice{
		tdCommon: tdCommon{Type: "pushMessageContentInvoice"},
		Price:    price,
		IsPinned: isPinned,
	}

	return &pushMessageContentInvoiceTemp
}

// GetPushMessageContentEnum return the enum type of this object
func (pushMessageContentInvoice *PushMessageContentInvoice) GetPushMessageContentEnum() PushMessageContentEnum {
	return PushMessageContentInvoiceType
}

// PushMessageContentLocation A message with a location
type PushMessageContentLocation struct {
	tdCommon
	IsLive   bool `json:"is_live"`   // True, if the location is live
	IsPinned bool `json:"is_pinned"` // True, if the message is a pinned message with the specified content
}

// MessageType return the string telegram-type of PushMessageContentLocation
func (pushMessageContentLocation *PushMessageContentLocation) MessageType() string {
	return "pushMessageContentLocation"
}

// NewPushMessageContentLocation creates a new PushMessageContentLocation
//
// @param isLive True, if the location is live
// @param isPinned True, if the message is a pinned message with the specified content
func NewPushMessageContentLocation(isLive bool, isPinned bool) *PushMessageContentLocation {
	pushMessageContentLocationTemp := PushMessageContentLocation{
		tdCommon: tdCommon{Type: "pushMessageContentLocation"},
		IsLive:   isLive,
		IsPinned: isPinned,
	}

	return &pushMessageContentLocationTemp
}

// GetPushMessageContentEnum return the enum type of this object
func (pushMessageContentLocation *PushMessageContentLocation) GetPushMessageContentEnum() PushMessageContentEnum {
	return PushMessageContentLocationType
}

// PushMessageContentPhoto A photo message
type PushMessageContentPhoto struct {
	tdCommon
	Photo    *Photo `json:"photo"`     // Message content; may be null
	Caption  string `json:"caption"`   // Photo caption
	IsSecret bool   `json:"is_secret"` // True, if the photo is secret
	IsPinned bool   `json:"is_pinned"` // True, if the message is a pinned message with the specified content
}

// MessageType return the string telegram-type of PushMessageContentPhoto
func (pushMessageContentPhoto *PushMessageContentPhoto) MessageType() string {
	return "pushMessageContentPhoto"
}

// NewPushMessageContentPhoto creates a new PushMessageContentPhoto
//
// @param photo Message content; may be null
// @param caption Photo caption
// @param isSecret True, if the photo is secret
// @param isPinned True, if the message is a pinned message with the specified content
func NewPushMessageContentPhoto(photo *Photo, caption string, isSecret bool, isPinned bool) *PushMessageContentPhoto {
	pushMessageContentPhotoTemp := PushMessageContentPhoto{
		tdCommon: tdCommon{Type: "pushMessageContentPhoto"},
		Photo:    photo,
		Caption:  caption,
		IsSecret: isSecret,
		IsPinned: isPinned,
	}

	return &pushMessageContentPhotoTemp
}

// GetPushMessageContentEnum return the enum type of this object
func (pushMessageContentPhoto *PushMessageContentPhoto) GetPushMessageContentEnum() PushMessageContentEnum {
	return PushMessageContentPhotoType
}

// PushMessageContentPoll A message with a poll
type PushMessageContentPoll struct {
	tdCommon
	Question  string `json:"question"`   // Poll question
	IsRegular bool   `json:"is_regular"` // True, if the poll is regular and not in quiz mode
	IsPinned  bool   `json:"is_pinned"`  // True, if the message is a pinned message with the specified content
}

// MessageType return the string telegram-type of PushMessageContentPoll
func (pushMessageContentPoll *PushMessageContentPoll) MessageType() string {
	return "pushMessageContentPoll"
}

// NewPushMessageContentPoll creates a new PushMessageContentPoll
//
// @param question Poll question
// @param isRegular True, if the poll is regular and not in quiz mode
// @param isPinned True, if the message is a pinned message with the specified content
func NewPushMessageContentPoll(question string, isRegular bool, isPinned bool) *PushMessageContentPoll {
	pushMessageContentPollTemp := PushMessageContentPoll{
		tdCommon:  tdCommon{Type: "pushMessageContentPoll"},
		Question:  question,
		IsRegular: isRegular,
		IsPinned:  isPinned,
	}

	return &pushMessageContentPollTemp
}

// GetPushMessageContentEnum return the enum type of this object
func (pushMessageContentPoll *PushMessageContentPoll) GetPushMessageContentEnum() PushMessageContentEnum {
	return PushMessageContentPollType
}

// PushMessageContentScreenshotTaken A screenshot of a message in the chat has been taken
type PushMessageContentScreenshotTaken struct {
	tdCommon
}

// MessageType return the string telegram-type of PushMessageContentScreenshotTaken
func (pushMessageContentScreenshotTaken *PushMessageContentScreenshotTaken) MessageType() string {
	return "pushMessageContentScreenshotTaken"
}

// NewPushMessageContentScreenshotTaken creates a new PushMessageContentScreenshotTaken
//
func NewPushMessageContentScreenshotTaken() *PushMessageContentScreenshotTaken {
	pushMessageContentScreenshotTakenTemp := PushMessageContentScreenshotTaken{
		tdCommon: tdCommon{Type: "pushMessageContentScreenshotTaken"},
	}

	return &pushMessageContentScreenshotTakenTemp
}

// GetPushMessageContentEnum return the enum type of this object
func (pushMessageContentScreenshotTaken *PushMessageContentScreenshotTaken) GetPushMessageContentEnum() PushMessageContentEnum {
	return PushMessageContentScreenshotTakenType
}

// PushMessageContentSticker A message with a sticker
type PushMessageContentSticker struct {
	tdCommon
	Sticker  *Sticker `json:"sticker"`   // Message content; may be null
	Emoji    string   `json:"emoji"`     // Emoji corresponding to the sticker; may be empty
	IsPinned bool     `json:"is_pinned"` // True, if the message is a pinned message with the specified content
}

// MessageType return the string telegram-type of PushMessageContentSticker
func (pushMessageContentSticker *PushMessageContentSticker) MessageType() string {
	return "pushMessageContentSticker"
}

// NewPushMessageContentSticker creates a new PushMessageContentSticker
//
// @param sticker Message content; may be null
// @param emoji Emoji corresponding to the sticker; may be empty
// @param isPinned True, if the message is a pinned message with the specified content
func NewPushMessageContentSticker(sticker *Sticker, emoji string, isPinned bool) *PushMessageContentSticker {
	pushMessageContentStickerTemp := PushMessageContentSticker{
		tdCommon: tdCommon{Type: "pushMessageContentSticker"},
		Sticker:  sticker,
		Emoji:    emoji,
		IsPinned: isPinned,
	}

	return &pushMessageContentStickerTemp
}

// GetPushMessageContentEnum return the enum type of this object
func (pushMessageContentSticker *PushMessageContentSticker) GetPushMessageContentEnum() PushMessageContentEnum {
	return PushMessageContentStickerType
}

// PushMessageContentText A text message
type PushMessageContentText struct {
	tdCommon
	Text     string `json:"text"`      // Message text
	IsPinned bool   `json:"is_pinned"` // True, if the message is a pinned message with the specified content
}

// MessageType return the string telegram-type of PushMessageContentText
func (pushMessageContentText *PushMessageContentText) MessageType() string {
	return "pushMessageContentText"
}

// NewPushMessageContentText creates a new PushMessageContentText
//
// @param text Message text
// @param isPinned True, if the message is a pinned message with the specified content
func NewPushMessageContentText(text string, isPinned bool) *PushMessageContentText {
	pushMessageContentTextTemp := PushMessageContentText{
		tdCommon: tdCommon{Type: "pushMessageContentText"},
		Text:     text,
		IsPinned: isPinned,
	}

	return &pushMessageContentTextTemp
}

// GetPushMessageContentEnum return the enum type of this object
func (pushMessageContentText *PushMessageContentText) GetPushMessageContentEnum() PushMessageContentEnum {
	return PushMessageContentTextType
}

// PushMessageContentVideo A video message
type PushMessageContentVideo struct {
	tdCommon
	Video    *Video `json:"video"`     // Message content; may be null
	Caption  string `json:"caption"`   // Video caption
	IsSecret bool   `json:"is_secret"` // True, if the video is secret
	IsPinned bool   `json:"is_pinned"` // True, if the message is a pinned message with the specified content
}

// MessageType return the string telegram-type of PushMessageContentVideo
func (pushMessageContentVideo *PushMessageContentVideo) MessageType() string {
	return "pushMessageContentVideo"
}

// NewPushMessageContentVideo creates a new PushMessageContentVideo
//
// @param video Message content; may be null
// @param caption Video caption
// @param isSecret True, if the video is secret
// @param isPinned True, if the message is a pinned message with the specified content
func NewPushMessageContentVideo(video *Video, caption string, isSecret bool, isPinned bool) *PushMessageContentVideo {
	pushMessageContentVideoTemp := PushMessageContentVideo{
		tdCommon: tdCommon{Type: "pushMessageContentVideo"},
		Video:    video,
		Caption:  caption,
		IsSecret: isSecret,
		IsPinned: isPinned,
	}

	return &pushMessageContentVideoTemp
}

// GetPushMessageContentEnum return the enum type of this object
func (pushMessageContentVideo *PushMessageContentVideo) GetPushMessageContentEnum() PushMessageContentEnum {
	return PushMessageContentVideoType
}

// PushMessageContentVideoNote A video note message
type PushMessageContentVideoNote struct {
	tdCommon
	VideoNote *VideoNote `json:"video_note"` // Message content; may be null
	IsPinned  bool       `json:"is_pinned"`  // True, if the message is a pinned message with the specified content
}

// MessageType return the string telegram-type of PushMessageContentVideoNote
func (pushMessageContentVideoNote *PushMessageContentVideoNote) MessageType() string {
	return "pushMessageContentVideoNote"
}

// NewPushMessageContentVideoNote creates a new PushMessageContentVideoNote
//
// @param videoNote Message content; may be null
// @param isPinned True, if the message is a pinned message with the specified content
func NewPushMessageContentVideoNote(videoNote *VideoNote, isPinned bool) *PushMessageContentVideoNote {
	pushMessageContentVideoNoteTemp := PushMessageContentVideoNote{
		tdCommon:  tdCommon{Type: "pushMessageContentVideoNote"},
		VideoNote: videoNote,
		IsPinned:  isPinned,
	}

	return &pushMessageContentVideoNoteTemp
}

// GetPushMessageContentEnum return the enum type of this object
func (pushMessageContentVideoNote *PushMessageContentVideoNote) GetPushMessageContentEnum() PushMessageContentEnum {
	return PushMessageContentVideoNoteType
}

// PushMessageContentVoiceNote A voice note message
type PushMessageContentVoiceNote struct {
	tdCommon
	VoiceNote *VoiceNote `json:"voice_note"` // Message content; may be null
	IsPinned  bool       `json:"is_pinned"`  // True, if the message is a pinned message with the specified content
}

// MessageType return the string telegram-type of PushMessageContentVoiceNote
func (pushMessageContentVoiceNote *PushMessageContentVoiceNote) MessageType() string {
	return "pushMessageContentVoiceNote"
}

// NewPushMessageContentVoiceNote creates a new PushMessageContentVoiceNote
//
// @param voiceNote Message content; may be null
// @param isPinned True, if the message is a pinned message with the specified content
func NewPushMessageContentVoiceNote(voiceNote *VoiceNote, isPinned bool) *PushMessageContentVoiceNote {
	pushMessageContentVoiceNoteTemp := PushMessageContentVoiceNote{
		tdCommon:  tdCommon{Type: "pushMessageContentVoiceNote"},
		VoiceNote: voiceNote,
		IsPinned:  isPinned,
	}

	return &pushMessageContentVoiceNoteTemp
}

// GetPushMessageContentEnum return the enum type of this object
func (pushMessageContentVoiceNote *PushMessageContentVoiceNote) GetPushMessageContentEnum() PushMessageContentEnum {
	return PushMessageContentVoiceNoteType
}

// PushMessageContentBasicGroupChatCreate A newly created basic group
type PushMessageContentBasicGroupChatCreate struct {
	tdCommon
}

// MessageType return the string telegram-type of PushMessageContentBasicGroupChatCreate
func (pushMessageContentBasicGroupChatCreate *PushMessageContentBasicGroupChatCreate) MessageType() string {
	return "pushMessageContentBasicGroupChatCreate"
}

// NewPushMessageContentBasicGroupChatCreate creates a new PushMessageContentBasicGroupChatCreate
//
func NewPushMessageContentBasicGroupChatCreate() *PushMessageContentBasicGroupChatCreate {
	pushMessageContentBasicGroupChatCreateTemp := PushMessageContentBasicGroupChatCreate{
		tdCommon: tdCommon{Type: "pushMessageContentBasicGroupChatCreate"},
	}

	return &pushMessageContentBasicGroupChatCreateTemp
}

// GetPushMessageContentEnum return the enum type of this object
func (pushMessageContentBasicGroupChatCreate *PushMessageContentBasicGroupChatCreate) GetPushMessageContentEnum() PushMessageContentEnum {
	return PushMessageContentBasicGroupChatCreateType
}

// PushMessageContentChatAddMembers New chat members were invited to a group
type PushMessageContentChatAddMembers struct {
	tdCommon
	MemberName    string `json:"member_name"`     // Name of the added member
	IsCurrentUser bool   `json:"is_current_user"` // True, if the current user was added to the group
	IsReturned    bool   `json:"is_returned"`     // True, if the user has returned to the group themselves
}

// MessageType return the string telegram-type of PushMessageContentChatAddMembers
func (pushMessageContentChatAddMembers *PushMessageContentChatAddMembers) MessageType() string {
	return "pushMessageContentChatAddMembers"
}

// NewPushMessageContentChatAddMembers creates a new PushMessageContentChatAddMembers
//
// @param memberName Name of the added member
// @param isCurrentUser True, if the current user was added to the group
// @param isReturned True, if the user has returned to the group themselves
func NewPushMessageContentChatAddMembers(memberName string, isCurrentUser bool, isReturned bool) *PushMessageContentChatAddMembers {
	pushMessageContentChatAddMembersTemp := PushMessageContentChatAddMembers{
		tdCommon:      tdCommon{Type: "pushMessageContentChatAddMembers"},
		MemberName:    memberName,
		IsCurrentUser: isCurrentUser,
		IsReturned:    isReturned,
	}

	return &pushMessageContentChatAddMembersTemp
}

// GetPushMessageContentEnum return the enum type of this object
func (pushMessageContentChatAddMembers *PushMessageContentChatAddMembers) GetPushMessageContentEnum() PushMessageContentEnum {
	return PushMessageContentChatAddMembersType
}

// PushMessageContentChatChangePhoto A chat photo was edited
type PushMessageContentChatChangePhoto struct {
	tdCommon
}

// MessageType return the string telegram-type of PushMessageContentChatChangePhoto
func (pushMessageContentChatChangePhoto *PushMessageContentChatChangePhoto) MessageType() string {
	return "pushMessageContentChatChangePhoto"
}

// NewPushMessageContentChatChangePhoto creates a new PushMessageContentChatChangePhoto
//
func NewPushMessageContentChatChangePhoto() *PushMessageContentChatChangePhoto {
	pushMessageContentChatChangePhotoTemp := PushMessageContentChatChangePhoto{
		tdCommon: tdCommon{Type: "pushMessageContentChatChangePhoto"},
	}

	return &pushMessageContentChatChangePhotoTemp
}

// GetPushMessageContentEnum return the enum type of this object
func (pushMessageContentChatChangePhoto *PushMessageContentChatChangePhoto) GetPushMessageContentEnum() PushMessageContentEnum {
	return PushMessageContentChatChangePhotoType
}

// PushMessageContentChatChangeTitle A chat title was edited
type PushMessageContentChatChangeTitle struct {
	tdCommon
	Title string `json:"title"` // New chat title
}

// MessageType return the string telegram-type of PushMessageContentChatChangeTitle
func (pushMessageContentChatChangeTitle *PushMessageContentChatChangeTitle) MessageType() string {
	return "pushMessageContentChatChangeTitle"
}

// NewPushMessageContentChatChangeTitle creates a new PushMessageContentChatChangeTitle
//
// @param title New chat title
func NewPushMessageContentChatChangeTitle(title string) *PushMessageContentChatChangeTitle {
	pushMessageContentChatChangeTitleTemp := PushMessageContentChatChangeTitle{
		tdCommon: tdCommon{Type: "pushMessageContentChatChangeTitle"},
		Title:    title,
	}

	return &pushMessageContentChatChangeTitleTemp
}

// GetPushMessageContentEnum return the enum type of this object
func (pushMessageContentChatChangeTitle *PushMessageContentChatChangeTitle) GetPushMessageContentEnum() PushMessageContentEnum {
	return PushMessageContentChatChangeTitleType
}

// PushMessageContentChatSetTheme A chat theme was edited
type PushMessageContentChatSetTheme struct {
	tdCommon
	ThemeName string `json:"theme_name"` // If non-empty, name of a new theme, set for the chat. Otherwise chat theme was reset to the default one
}

// MessageType return the string telegram-type of PushMessageContentChatSetTheme
func (pushMessageContentChatSetTheme *PushMessageContentChatSetTheme) MessageType() string {
	return "pushMessageContentChatSetTheme"
}

// NewPushMessageContentChatSetTheme creates a new PushMessageContentChatSetTheme
//
// @param themeName If non-empty, name of a new theme, set for the chat. Otherwise chat theme was reset to the default one
func NewPushMessageContentChatSetTheme(themeName string) *PushMessageContentChatSetTheme {
	pushMessageContentChatSetThemeTemp := PushMessageContentChatSetTheme{
		tdCommon:  tdCommon{Type: "pushMessageContentChatSetTheme"},
		ThemeName: themeName,
	}

	return &pushMessageContentChatSetThemeTemp
}

// GetPushMessageContentEnum return the enum type of this object
func (pushMessageContentChatSetTheme *PushMessageContentChatSetTheme) GetPushMessageContentEnum() PushMessageContentEnum {
	return PushMessageContentChatSetThemeType
}

// PushMessageContentChatDeleteMember A chat member was deleted
type PushMessageContentChatDeleteMember struct {
	tdCommon
	MemberName    string `json:"member_name"`     // Name of the deleted member
	IsCurrentUser bool   `json:"is_current_user"` // True, if the current user was deleted from the group
	IsLeft        bool   `json:"is_left"`         // True, if the user has left the group themselves
}

// MessageType return the string telegram-type of PushMessageContentChatDeleteMember
func (pushMessageContentChatDeleteMember *PushMessageContentChatDeleteMember) MessageType() string {
	return "pushMessageContentChatDeleteMember"
}

// NewPushMessageContentChatDeleteMember creates a new PushMessageContentChatDeleteMember
//
// @param memberName Name of the deleted member
// @param isCurrentUser True, if the current user was deleted from the group
// @param isLeft True, if the user has left the group themselves
func NewPushMessageContentChatDeleteMember(memberName string, isCurrentUser bool, isLeft bool) *PushMessageContentChatDeleteMember {
	pushMessageContentChatDeleteMemberTemp := PushMessageContentChatDeleteMember{
		tdCommon:      tdCommon{Type: "pushMessageContentChatDeleteMember"},
		MemberName:    memberName,
		IsCurrentUser: isCurrentUser,
		IsLeft:        isLeft,
	}

	return &pushMessageContentChatDeleteMemberTemp
}

// GetPushMessageContentEnum return the enum type of this object
func (pushMessageContentChatDeleteMember *PushMessageContentChatDeleteMember) GetPushMessageContentEnum() PushMessageContentEnum {
	return PushMessageContentChatDeleteMemberType
}

// PushMessageContentChatJoinByLink A new member joined the chat via an invite link
type PushMessageContentChatJoinByLink struct {
	tdCommon
}

// MessageType return the string telegram-type of PushMessageContentChatJoinByLink
func (pushMessageContentChatJoinByLink *PushMessageContentChatJoinByLink) MessageType() string {
	return "pushMessageContentChatJoinByLink"
}

// NewPushMessageContentChatJoinByLink creates a new PushMessageContentChatJoinByLink
//
func NewPushMessageContentChatJoinByLink() *PushMessageContentChatJoinByLink {
	pushMessageContentChatJoinByLinkTemp := PushMessageContentChatJoinByLink{
		tdCommon: tdCommon{Type: "pushMessageContentChatJoinByLink"},
	}

	return &pushMessageContentChatJoinByLinkTemp
}

// GetPushMessageContentEnum return the enum type of this object
func (pushMessageContentChatJoinByLink *PushMessageContentChatJoinByLink) GetPushMessageContentEnum() PushMessageContentEnum {
	return PushMessageContentChatJoinByLinkType
}

// PushMessageContentChatJoinByRequest A new member was accepted to the chat by an administrator
type PushMessageContentChatJoinByRequest struct {
	tdCommon
}

// MessageType return the string telegram-type of PushMessageContentChatJoinByRequest
func (pushMessageContentChatJoinByRequest *PushMessageContentChatJoinByRequest) MessageType() string {
	return "pushMessageContentChatJoinByRequest"
}

// NewPushMessageContentChatJoinByRequest creates a new PushMessageContentChatJoinByRequest
//
func NewPushMessageContentChatJoinByRequest() *PushMessageContentChatJoinByRequest {
	pushMessageContentChatJoinByRequestTemp := PushMessageContentChatJoinByRequest{
		tdCommon: tdCommon{Type: "pushMessageContentChatJoinByRequest"},
	}

	return &pushMessageContentChatJoinByRequestTemp
}

// GetPushMessageContentEnum return the enum type of this object
func (pushMessageContentChatJoinByRequest *PushMessageContentChatJoinByRequest) GetPushMessageContentEnum() PushMessageContentEnum {
	return PushMessageContentChatJoinByRequestType
}

// PushMessageContentMessageForwards A forwarded messages
type PushMessageContentMessageForwards struct {
	tdCommon
	TotalCount int32 `json:"total_count"` // Number of forwarded messages
}

// MessageType return the string telegram-type of PushMessageContentMessageForwards
func (pushMessageContentMessageForwards *PushMessageContentMessageForwards) MessageType() string {
	return "pushMessageContentMessageForwards"
}

// NewPushMessageContentMessageForwards creates a new PushMessageContentMessageForwards
//
// @param totalCount Number of forwarded messages
func NewPushMessageContentMessageForwards(totalCount int32) *PushMessageContentMessageForwards {
	pushMessageContentMessageForwardsTemp := PushMessageContentMessageForwards{
		tdCommon:   tdCommon{Type: "pushMessageContentMessageForwards"},
		TotalCount: totalCount,
	}

	return &pushMessageContentMessageForwardsTemp
}

// GetPushMessageContentEnum return the enum type of this object
func (pushMessageContentMessageForwards *PushMessageContentMessageForwards) GetPushMessageContentEnum() PushMessageContentEnum {
	return PushMessageContentMessageForwardsType
}

// PushMessageContentMediaAlbum A media album
type PushMessageContentMediaAlbum struct {
	tdCommon
	TotalCount   int32 `json:"total_count"`   // Number of messages in the album
	HasPhotos    bool  `json:"has_photos"`    // True, if the album has at least one photo
	HasVideos    bool  `json:"has_videos"`    // True, if the album has at least one video
	HasAudios    bool  `json:"has_audios"`    // True, if the album has at least one audio file
	HasDocuments bool  `json:"has_documents"` // True, if the album has at least one document
}

// MessageType return the string telegram-type of PushMessageContentMediaAlbum
func (pushMessageContentMediaAlbum *PushMessageContentMediaAlbum) MessageType() string {
	return "pushMessageContentMediaAlbum"
}

// NewPushMessageContentMediaAlbum creates a new PushMessageContentMediaAlbum
//
// @param totalCount Number of messages in the album
// @param hasPhotos True, if the album has at least one photo
// @param hasVideos True, if the album has at least one video
// @param hasAudios True, if the album has at least one audio file
// @param hasDocuments True, if the album has at least one document
func NewPushMessageContentMediaAlbum(totalCount int32, hasPhotos bool, hasVideos bool, hasAudios bool, hasDocuments bool) *PushMessageContentMediaAlbum {
	pushMessageContentMediaAlbumTemp := PushMessageContentMediaAlbum{
		tdCommon:     tdCommon{Type: "pushMessageContentMediaAlbum"},
		TotalCount:   totalCount,
		HasPhotos:    hasPhotos,
		HasVideos:    hasVideos,
		HasAudios:    hasAudios,
		HasDocuments: hasDocuments,
	}

	return &pushMessageContentMediaAlbumTemp
}

// GetPushMessageContentEnum return the enum type of this object
func (pushMessageContentMediaAlbum *PushMessageContentMediaAlbum) GetPushMessageContentEnum() PushMessageContentEnum {
	return PushMessageContentMediaAlbumType
}
