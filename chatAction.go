package tdlib

import (
	"encoding/json"
	"fmt"
)

// ChatAction Describes the different types of activity in a chat
type ChatAction interface {
	GetChatActionEnum() ChatActionEnum
}

// ChatActionEnum Alias for abstract ChatAction 'Sub-Classes', used as constant-enum here
type ChatActionEnum string

// ChatAction enums
const (
	ChatActionTypingType             ChatActionEnum = "chatActionTyping"
	ChatActionRecordingVideoType     ChatActionEnum = "chatActionRecordingVideo"
	ChatActionUploadingVideoType     ChatActionEnum = "chatActionUploadingVideo"
	ChatActionRecordingVoiceNoteType ChatActionEnum = "chatActionRecordingVoiceNote"
	ChatActionUploadingVoiceNoteType ChatActionEnum = "chatActionUploadingVoiceNote"
	ChatActionUploadingPhotoType     ChatActionEnum = "chatActionUploadingPhoto"
	ChatActionUploadingDocumentType  ChatActionEnum = "chatActionUploadingDocument"
	ChatActionChoosingStickerType    ChatActionEnum = "chatActionChoosingSticker"
	ChatActionChoosingLocationType   ChatActionEnum = "chatActionChoosingLocation"
	ChatActionChoosingContactType    ChatActionEnum = "chatActionChoosingContact"
	ChatActionStartPlayingGameType   ChatActionEnum = "chatActionStartPlayingGame"
	ChatActionRecordingVideoNoteType ChatActionEnum = "chatActionRecordingVideoNote"
	ChatActionUploadingVideoNoteType ChatActionEnum = "chatActionUploadingVideoNote"
	ChatActionWatchingAnimationsType ChatActionEnum = "chatActionWatchingAnimations"
	ChatActionCancelType             ChatActionEnum = "chatActionCancel"
)

func unmarshalChatAction(rawMsg *json.RawMessage) (ChatAction, error) {

	if rawMsg == nil {
		return nil, nil
	}
	var objMap map[string]interface{}
	err := json.Unmarshal(*rawMsg, &objMap)
	if err != nil {
		return nil, err
	}

	switch ChatActionEnum(objMap["@type"].(string)) {
	case ChatActionTypingType:
		var chatActionTyping ChatActionTyping
		err := json.Unmarshal(*rawMsg, &chatActionTyping)
		return &chatActionTyping, err

	case ChatActionRecordingVideoType:
		var chatActionRecordingVideo ChatActionRecordingVideo
		err := json.Unmarshal(*rawMsg, &chatActionRecordingVideo)
		return &chatActionRecordingVideo, err

	case ChatActionUploadingVideoType:
		var chatActionUploadingVideo ChatActionUploadingVideo
		err := json.Unmarshal(*rawMsg, &chatActionUploadingVideo)
		return &chatActionUploadingVideo, err

	case ChatActionRecordingVoiceNoteType:
		var chatActionRecordingVoiceNote ChatActionRecordingVoiceNote
		err := json.Unmarshal(*rawMsg, &chatActionRecordingVoiceNote)
		return &chatActionRecordingVoiceNote, err

	case ChatActionUploadingVoiceNoteType:
		var chatActionUploadingVoiceNote ChatActionUploadingVoiceNote
		err := json.Unmarshal(*rawMsg, &chatActionUploadingVoiceNote)
		return &chatActionUploadingVoiceNote, err

	case ChatActionUploadingPhotoType:
		var chatActionUploadingPhoto ChatActionUploadingPhoto
		err := json.Unmarshal(*rawMsg, &chatActionUploadingPhoto)
		return &chatActionUploadingPhoto, err

	case ChatActionUploadingDocumentType:
		var chatActionUploadingDocument ChatActionUploadingDocument
		err := json.Unmarshal(*rawMsg, &chatActionUploadingDocument)
		return &chatActionUploadingDocument, err

	case ChatActionChoosingStickerType:
		var chatActionChoosingSticker ChatActionChoosingSticker
		err := json.Unmarshal(*rawMsg, &chatActionChoosingSticker)
		return &chatActionChoosingSticker, err

	case ChatActionChoosingLocationType:
		var chatActionChoosingLocation ChatActionChoosingLocation
		err := json.Unmarshal(*rawMsg, &chatActionChoosingLocation)
		return &chatActionChoosingLocation, err

	case ChatActionChoosingContactType:
		var chatActionChoosingContact ChatActionChoosingContact
		err := json.Unmarshal(*rawMsg, &chatActionChoosingContact)
		return &chatActionChoosingContact, err

	case ChatActionStartPlayingGameType:
		var chatActionStartPlayingGame ChatActionStartPlayingGame
		err := json.Unmarshal(*rawMsg, &chatActionStartPlayingGame)
		return &chatActionStartPlayingGame, err

	case ChatActionRecordingVideoNoteType:
		var chatActionRecordingVideoNote ChatActionRecordingVideoNote
		err := json.Unmarshal(*rawMsg, &chatActionRecordingVideoNote)
		return &chatActionRecordingVideoNote, err

	case ChatActionUploadingVideoNoteType:
		var chatActionUploadingVideoNote ChatActionUploadingVideoNote
		err := json.Unmarshal(*rawMsg, &chatActionUploadingVideoNote)
		return &chatActionUploadingVideoNote, err

	case ChatActionWatchingAnimationsType:
		var chatActionWatchingAnimations ChatActionWatchingAnimations
		err := json.Unmarshal(*rawMsg, &chatActionWatchingAnimations)
		return &chatActionWatchingAnimations, err

	case ChatActionCancelType:
		var chatActionCancel ChatActionCancel
		err := json.Unmarshal(*rawMsg, &chatActionCancel)
		return &chatActionCancel, err

	default:
		return nil, fmt.Errorf("Error UnMarshaling, unknown type:" + objMap["@type"].(string))
	}
}

// ChatActionTyping The user is typing a message
type ChatActionTyping struct {
	tdCommon
}

// MessageType return the string telegram-type of ChatActionTyping
func (chatActionTyping *ChatActionTyping) MessageType() string {
	return "chatActionTyping"
}

// NewChatActionTyping creates a new ChatActionTyping
//
func NewChatActionTyping() *ChatActionTyping {
	chatActionTypingTemp := ChatActionTyping{
		tdCommon: tdCommon{Type: "chatActionTyping"},
	}

	return &chatActionTypingTemp
}

// GetChatActionEnum return the enum type of this object
func (chatActionTyping *ChatActionTyping) GetChatActionEnum() ChatActionEnum {
	return ChatActionTypingType
}

// ChatActionRecordingVideo The user is recording a video
type ChatActionRecordingVideo struct {
	tdCommon
}

// MessageType return the string telegram-type of ChatActionRecordingVideo
func (chatActionRecordingVideo *ChatActionRecordingVideo) MessageType() string {
	return "chatActionRecordingVideo"
}

// NewChatActionRecordingVideo creates a new ChatActionRecordingVideo
//
func NewChatActionRecordingVideo() *ChatActionRecordingVideo {
	chatActionRecordingVideoTemp := ChatActionRecordingVideo{
		tdCommon: tdCommon{Type: "chatActionRecordingVideo"},
	}

	return &chatActionRecordingVideoTemp
}

// GetChatActionEnum return the enum type of this object
func (chatActionRecordingVideo *ChatActionRecordingVideo) GetChatActionEnum() ChatActionEnum {
	return ChatActionRecordingVideoType
}

// ChatActionUploadingVideo The user is uploading a video
type ChatActionUploadingVideo struct {
	tdCommon
	Progress int32 `json:"progress"` // Upload progress, as a percentage
}

// MessageType return the string telegram-type of ChatActionUploadingVideo
func (chatActionUploadingVideo *ChatActionUploadingVideo) MessageType() string {
	return "chatActionUploadingVideo"
}

// NewChatActionUploadingVideo creates a new ChatActionUploadingVideo
//
// @param progress Upload progress, as a percentage
func NewChatActionUploadingVideo(progress int32) *ChatActionUploadingVideo {
	chatActionUploadingVideoTemp := ChatActionUploadingVideo{
		tdCommon: tdCommon{Type: "chatActionUploadingVideo"},
		Progress: progress,
	}

	return &chatActionUploadingVideoTemp
}

// GetChatActionEnum return the enum type of this object
func (chatActionUploadingVideo *ChatActionUploadingVideo) GetChatActionEnum() ChatActionEnum {
	return ChatActionUploadingVideoType
}

// ChatActionRecordingVoiceNote The user is recording a voice note
type ChatActionRecordingVoiceNote struct {
	tdCommon
}

// MessageType return the string telegram-type of ChatActionRecordingVoiceNote
func (chatActionRecordingVoiceNote *ChatActionRecordingVoiceNote) MessageType() string {
	return "chatActionRecordingVoiceNote"
}

// NewChatActionRecordingVoiceNote creates a new ChatActionRecordingVoiceNote
//
func NewChatActionRecordingVoiceNote() *ChatActionRecordingVoiceNote {
	chatActionRecordingVoiceNoteTemp := ChatActionRecordingVoiceNote{
		tdCommon: tdCommon{Type: "chatActionRecordingVoiceNote"},
	}

	return &chatActionRecordingVoiceNoteTemp
}

// GetChatActionEnum return the enum type of this object
func (chatActionRecordingVoiceNote *ChatActionRecordingVoiceNote) GetChatActionEnum() ChatActionEnum {
	return ChatActionRecordingVoiceNoteType
}

// ChatActionUploadingVoiceNote The user is uploading a voice note
type ChatActionUploadingVoiceNote struct {
	tdCommon
	Progress int32 `json:"progress"` // Upload progress, as a percentage
}

// MessageType return the string telegram-type of ChatActionUploadingVoiceNote
func (chatActionUploadingVoiceNote *ChatActionUploadingVoiceNote) MessageType() string {
	return "chatActionUploadingVoiceNote"
}

// NewChatActionUploadingVoiceNote creates a new ChatActionUploadingVoiceNote
//
// @param progress Upload progress, as a percentage
func NewChatActionUploadingVoiceNote(progress int32) *ChatActionUploadingVoiceNote {
	chatActionUploadingVoiceNoteTemp := ChatActionUploadingVoiceNote{
		tdCommon: tdCommon{Type: "chatActionUploadingVoiceNote"},
		Progress: progress,
	}

	return &chatActionUploadingVoiceNoteTemp
}

// GetChatActionEnum return the enum type of this object
func (chatActionUploadingVoiceNote *ChatActionUploadingVoiceNote) GetChatActionEnum() ChatActionEnum {
	return ChatActionUploadingVoiceNoteType
}

// ChatActionUploadingPhoto The user is uploading a photo
type ChatActionUploadingPhoto struct {
	tdCommon
	Progress int32 `json:"progress"` // Upload progress, as a percentage
}

// MessageType return the string telegram-type of ChatActionUploadingPhoto
func (chatActionUploadingPhoto *ChatActionUploadingPhoto) MessageType() string {
	return "chatActionUploadingPhoto"
}

// NewChatActionUploadingPhoto creates a new ChatActionUploadingPhoto
//
// @param progress Upload progress, as a percentage
func NewChatActionUploadingPhoto(progress int32) *ChatActionUploadingPhoto {
	chatActionUploadingPhotoTemp := ChatActionUploadingPhoto{
		tdCommon: tdCommon{Type: "chatActionUploadingPhoto"},
		Progress: progress,
	}

	return &chatActionUploadingPhotoTemp
}

// GetChatActionEnum return the enum type of this object
func (chatActionUploadingPhoto *ChatActionUploadingPhoto) GetChatActionEnum() ChatActionEnum {
	return ChatActionUploadingPhotoType
}

// ChatActionUploadingDocument The user is uploading a document
type ChatActionUploadingDocument struct {
	tdCommon
	Progress int32 `json:"progress"` // Upload progress, as a percentage
}

// MessageType return the string telegram-type of ChatActionUploadingDocument
func (chatActionUploadingDocument *ChatActionUploadingDocument) MessageType() string {
	return "chatActionUploadingDocument"
}

// NewChatActionUploadingDocument creates a new ChatActionUploadingDocument
//
// @param progress Upload progress, as a percentage
func NewChatActionUploadingDocument(progress int32) *ChatActionUploadingDocument {
	chatActionUploadingDocumentTemp := ChatActionUploadingDocument{
		tdCommon: tdCommon{Type: "chatActionUploadingDocument"},
		Progress: progress,
	}

	return &chatActionUploadingDocumentTemp
}

// GetChatActionEnum return the enum type of this object
func (chatActionUploadingDocument *ChatActionUploadingDocument) GetChatActionEnum() ChatActionEnum {
	return ChatActionUploadingDocumentType
}

// ChatActionChoosingSticker The user is picking a sticker to send
type ChatActionChoosingSticker struct {
	tdCommon
}

// MessageType return the string telegram-type of ChatActionChoosingSticker
func (chatActionChoosingSticker *ChatActionChoosingSticker) MessageType() string {
	return "chatActionChoosingSticker"
}

// NewChatActionChoosingSticker creates a new ChatActionChoosingSticker
//
func NewChatActionChoosingSticker() *ChatActionChoosingSticker {
	chatActionChoosingStickerTemp := ChatActionChoosingSticker{
		tdCommon: tdCommon{Type: "chatActionChoosingSticker"},
	}

	return &chatActionChoosingStickerTemp
}

// GetChatActionEnum return the enum type of this object
func (chatActionChoosingSticker *ChatActionChoosingSticker) GetChatActionEnum() ChatActionEnum {
	return ChatActionChoosingStickerType
}

// ChatActionChoosingLocation The user is picking a location or venue to send
type ChatActionChoosingLocation struct {
	tdCommon
}

// MessageType return the string telegram-type of ChatActionChoosingLocation
func (chatActionChoosingLocation *ChatActionChoosingLocation) MessageType() string {
	return "chatActionChoosingLocation"
}

// NewChatActionChoosingLocation creates a new ChatActionChoosingLocation
//
func NewChatActionChoosingLocation() *ChatActionChoosingLocation {
	chatActionChoosingLocationTemp := ChatActionChoosingLocation{
		tdCommon: tdCommon{Type: "chatActionChoosingLocation"},
	}

	return &chatActionChoosingLocationTemp
}

// GetChatActionEnum return the enum type of this object
func (chatActionChoosingLocation *ChatActionChoosingLocation) GetChatActionEnum() ChatActionEnum {
	return ChatActionChoosingLocationType
}

// ChatActionChoosingContact The user is picking a contact to send
type ChatActionChoosingContact struct {
	tdCommon
}

// MessageType return the string telegram-type of ChatActionChoosingContact
func (chatActionChoosingContact *ChatActionChoosingContact) MessageType() string {
	return "chatActionChoosingContact"
}

// NewChatActionChoosingContact creates a new ChatActionChoosingContact
//
func NewChatActionChoosingContact() *ChatActionChoosingContact {
	chatActionChoosingContactTemp := ChatActionChoosingContact{
		tdCommon: tdCommon{Type: "chatActionChoosingContact"},
	}

	return &chatActionChoosingContactTemp
}

// GetChatActionEnum return the enum type of this object
func (chatActionChoosingContact *ChatActionChoosingContact) GetChatActionEnum() ChatActionEnum {
	return ChatActionChoosingContactType
}

// ChatActionStartPlayingGame The user has started to play a game
type ChatActionStartPlayingGame struct {
	tdCommon
}

// MessageType return the string telegram-type of ChatActionStartPlayingGame
func (chatActionStartPlayingGame *ChatActionStartPlayingGame) MessageType() string {
	return "chatActionStartPlayingGame"
}

// NewChatActionStartPlayingGame creates a new ChatActionStartPlayingGame
//
func NewChatActionStartPlayingGame() *ChatActionStartPlayingGame {
	chatActionStartPlayingGameTemp := ChatActionStartPlayingGame{
		tdCommon: tdCommon{Type: "chatActionStartPlayingGame"},
	}

	return &chatActionStartPlayingGameTemp
}

// GetChatActionEnum return the enum type of this object
func (chatActionStartPlayingGame *ChatActionStartPlayingGame) GetChatActionEnum() ChatActionEnum {
	return ChatActionStartPlayingGameType
}

// ChatActionRecordingVideoNote The user is recording a video note
type ChatActionRecordingVideoNote struct {
	tdCommon
}

// MessageType return the string telegram-type of ChatActionRecordingVideoNote
func (chatActionRecordingVideoNote *ChatActionRecordingVideoNote) MessageType() string {
	return "chatActionRecordingVideoNote"
}

// NewChatActionRecordingVideoNote creates a new ChatActionRecordingVideoNote
//
func NewChatActionRecordingVideoNote() *ChatActionRecordingVideoNote {
	chatActionRecordingVideoNoteTemp := ChatActionRecordingVideoNote{
		tdCommon: tdCommon{Type: "chatActionRecordingVideoNote"},
	}

	return &chatActionRecordingVideoNoteTemp
}

// GetChatActionEnum return the enum type of this object
func (chatActionRecordingVideoNote *ChatActionRecordingVideoNote) GetChatActionEnum() ChatActionEnum {
	return ChatActionRecordingVideoNoteType
}

// ChatActionUploadingVideoNote The user is uploading a video note
type ChatActionUploadingVideoNote struct {
	tdCommon
	Progress int32 `json:"progress"` // Upload progress, as a percentage
}

// MessageType return the string telegram-type of ChatActionUploadingVideoNote
func (chatActionUploadingVideoNote *ChatActionUploadingVideoNote) MessageType() string {
	return "chatActionUploadingVideoNote"
}

// NewChatActionUploadingVideoNote creates a new ChatActionUploadingVideoNote
//
// @param progress Upload progress, as a percentage
func NewChatActionUploadingVideoNote(progress int32) *ChatActionUploadingVideoNote {
	chatActionUploadingVideoNoteTemp := ChatActionUploadingVideoNote{
		tdCommon: tdCommon{Type: "chatActionUploadingVideoNote"},
		Progress: progress,
	}

	return &chatActionUploadingVideoNoteTemp
}

// GetChatActionEnum return the enum type of this object
func (chatActionUploadingVideoNote *ChatActionUploadingVideoNote) GetChatActionEnum() ChatActionEnum {
	return ChatActionUploadingVideoNoteType
}

// ChatActionWatchingAnimations The user is watching animations sent by the other party by clicking on an animated emoji
type ChatActionWatchingAnimations struct {
	tdCommon
	Emoji string `json:"emoji"` // The animated emoji
}

// MessageType return the string telegram-type of ChatActionWatchingAnimations
func (chatActionWatchingAnimations *ChatActionWatchingAnimations) MessageType() string {
	return "chatActionWatchingAnimations"
}

// NewChatActionWatchingAnimations creates a new ChatActionWatchingAnimations
//
// @param emoji The animated emoji
func NewChatActionWatchingAnimations(emoji string) *ChatActionWatchingAnimations {
	chatActionWatchingAnimationsTemp := ChatActionWatchingAnimations{
		tdCommon: tdCommon{Type: "chatActionWatchingAnimations"},
		Emoji:    emoji,
	}

	return &chatActionWatchingAnimationsTemp
}

// GetChatActionEnum return the enum type of this object
func (chatActionWatchingAnimations *ChatActionWatchingAnimations) GetChatActionEnum() ChatActionEnum {
	return ChatActionWatchingAnimationsType
}

// ChatActionCancel The user has canceled the previous action
type ChatActionCancel struct {
	tdCommon
}

// MessageType return the string telegram-type of ChatActionCancel
func (chatActionCancel *ChatActionCancel) MessageType() string {
	return "chatActionCancel"
}

// NewChatActionCancel creates a new ChatActionCancel
//
func NewChatActionCancel() *ChatActionCancel {
	chatActionCancelTemp := ChatActionCancel{
		tdCommon: tdCommon{Type: "chatActionCancel"},
	}

	return &chatActionCancelTemp
}

// GetChatActionEnum return the enum type of this object
func (chatActionCancel *ChatActionCancel) GetChatActionEnum() ChatActionEnum {
	return ChatActionCancelType
}
