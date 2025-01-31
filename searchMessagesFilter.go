package tdlib

import (
	"encoding/json"
	"fmt"
)

// SearchMessagesFilter Represents a filter for message search results
type SearchMessagesFilter interface {
	GetSearchMessagesFilterEnum() SearchMessagesFilterEnum
}

// SearchMessagesFilterEnum Alias for abstract SearchMessagesFilter 'Sub-Classes', used as constant-enum here
type SearchMessagesFilterEnum string

// SearchMessagesFilter enums
const (
	SearchMessagesFilterEmptyType             SearchMessagesFilterEnum = "searchMessagesFilterEmpty"
	SearchMessagesFilterAnimationType         SearchMessagesFilterEnum = "searchMessagesFilterAnimation"
	SearchMessagesFilterAudioType             SearchMessagesFilterEnum = "searchMessagesFilterAudio"
	SearchMessagesFilterDocumentType          SearchMessagesFilterEnum = "searchMessagesFilterDocument"
	SearchMessagesFilterPhotoType             SearchMessagesFilterEnum = "searchMessagesFilterPhoto"
	SearchMessagesFilterVideoType             SearchMessagesFilterEnum = "searchMessagesFilterVideo"
	SearchMessagesFilterVoiceNoteType         SearchMessagesFilterEnum = "searchMessagesFilterVoiceNote"
	SearchMessagesFilterPhotoAndVideoType     SearchMessagesFilterEnum = "searchMessagesFilterPhotoAndVideo"
	SearchMessagesFilterUrlType               SearchMessagesFilterEnum = "searchMessagesFilterUrl"
	SearchMessagesFilterChatPhotoType         SearchMessagesFilterEnum = "searchMessagesFilterChatPhoto"
	SearchMessagesFilterVideoNoteType         SearchMessagesFilterEnum = "searchMessagesFilterVideoNote"
	SearchMessagesFilterVoiceAndVideoNoteType SearchMessagesFilterEnum = "searchMessagesFilterVoiceAndVideoNote"
	SearchMessagesFilterMentionType           SearchMessagesFilterEnum = "searchMessagesFilterMention"
	SearchMessagesFilterUnreadMentionType     SearchMessagesFilterEnum = "searchMessagesFilterUnreadMention"
	SearchMessagesFilterFailedToSendType      SearchMessagesFilterEnum = "searchMessagesFilterFailedToSend"
	SearchMessagesFilterPinnedType            SearchMessagesFilterEnum = "searchMessagesFilterPinned"
)

func unmarshalSearchMessagesFilter(rawMsg *json.RawMessage) (SearchMessagesFilter, error) {

	if rawMsg == nil {
		return nil, nil
	}
	var objMap map[string]interface{}
	err := json.Unmarshal(*rawMsg, &objMap)
	if err != nil {
		return nil, err
	}

	switch SearchMessagesFilterEnum(objMap["@type"].(string)) {
	case SearchMessagesFilterEmptyType:
		var searchMessagesFilterEmpty SearchMessagesFilterEmpty
		err := json.Unmarshal(*rawMsg, &searchMessagesFilterEmpty)
		return &searchMessagesFilterEmpty, err

	case SearchMessagesFilterAnimationType:
		var searchMessagesFilterAnimation SearchMessagesFilterAnimation
		err := json.Unmarshal(*rawMsg, &searchMessagesFilterAnimation)
		return &searchMessagesFilterAnimation, err

	case SearchMessagesFilterAudioType:
		var searchMessagesFilterAudio SearchMessagesFilterAudio
		err := json.Unmarshal(*rawMsg, &searchMessagesFilterAudio)
		return &searchMessagesFilterAudio, err

	case SearchMessagesFilterDocumentType:
		var searchMessagesFilterDocument SearchMessagesFilterDocument
		err := json.Unmarshal(*rawMsg, &searchMessagesFilterDocument)
		return &searchMessagesFilterDocument, err

	case SearchMessagesFilterPhotoType:
		var searchMessagesFilterPhoto SearchMessagesFilterPhoto
		err := json.Unmarshal(*rawMsg, &searchMessagesFilterPhoto)
		return &searchMessagesFilterPhoto, err

	case SearchMessagesFilterVideoType:
		var searchMessagesFilterVideo SearchMessagesFilterVideo
		err := json.Unmarshal(*rawMsg, &searchMessagesFilterVideo)
		return &searchMessagesFilterVideo, err

	case SearchMessagesFilterVoiceNoteType:
		var searchMessagesFilterVoiceNote SearchMessagesFilterVoiceNote
		err := json.Unmarshal(*rawMsg, &searchMessagesFilterVoiceNote)
		return &searchMessagesFilterVoiceNote, err

	case SearchMessagesFilterPhotoAndVideoType:
		var searchMessagesFilterPhotoAndVideo SearchMessagesFilterPhotoAndVideo
		err := json.Unmarshal(*rawMsg, &searchMessagesFilterPhotoAndVideo)
		return &searchMessagesFilterPhotoAndVideo, err

	case SearchMessagesFilterUrlType:
		var searchMessagesFilterUrl SearchMessagesFilterUrl
		err := json.Unmarshal(*rawMsg, &searchMessagesFilterUrl)
		return &searchMessagesFilterUrl, err

	case SearchMessagesFilterChatPhotoType:
		var searchMessagesFilterChatPhoto SearchMessagesFilterChatPhoto
		err := json.Unmarshal(*rawMsg, &searchMessagesFilterChatPhoto)
		return &searchMessagesFilterChatPhoto, err

	case SearchMessagesFilterVideoNoteType:
		var searchMessagesFilterVideoNote SearchMessagesFilterVideoNote
		err := json.Unmarshal(*rawMsg, &searchMessagesFilterVideoNote)
		return &searchMessagesFilterVideoNote, err

	case SearchMessagesFilterVoiceAndVideoNoteType:
		var searchMessagesFilterVoiceAndVideoNote SearchMessagesFilterVoiceAndVideoNote
		err := json.Unmarshal(*rawMsg, &searchMessagesFilterVoiceAndVideoNote)
		return &searchMessagesFilterVoiceAndVideoNote, err

	case SearchMessagesFilterMentionType:
		var searchMessagesFilterMention SearchMessagesFilterMention
		err := json.Unmarshal(*rawMsg, &searchMessagesFilterMention)
		return &searchMessagesFilterMention, err

	case SearchMessagesFilterUnreadMentionType:
		var searchMessagesFilterUnreadMention SearchMessagesFilterUnreadMention
		err := json.Unmarshal(*rawMsg, &searchMessagesFilterUnreadMention)
		return &searchMessagesFilterUnreadMention, err

	case SearchMessagesFilterFailedToSendType:
		var searchMessagesFilterFailedToSend SearchMessagesFilterFailedToSend
		err := json.Unmarshal(*rawMsg, &searchMessagesFilterFailedToSend)
		return &searchMessagesFilterFailedToSend, err

	case SearchMessagesFilterPinnedType:
		var searchMessagesFilterPinned SearchMessagesFilterPinned
		err := json.Unmarshal(*rawMsg, &searchMessagesFilterPinned)
		return &searchMessagesFilterPinned, err

	default:
		return nil, fmt.Errorf("Error UnMarshaling, unknown type:" + objMap["@type"].(string))
	}
}

// SearchMessagesFilterEmpty Returns all found messages, no filter is applied
type SearchMessagesFilterEmpty struct {
	tdCommon
}

// MessageType return the string telegram-type of SearchMessagesFilterEmpty
func (searchMessagesFilterEmpty *SearchMessagesFilterEmpty) MessageType() string {
	return "searchMessagesFilterEmpty"
}

// NewSearchMessagesFilterEmpty creates a new SearchMessagesFilterEmpty
//
func NewSearchMessagesFilterEmpty() *SearchMessagesFilterEmpty {
	searchMessagesFilterEmptyTemp := SearchMessagesFilterEmpty{
		tdCommon: tdCommon{Type: "searchMessagesFilterEmpty"},
	}

	return &searchMessagesFilterEmptyTemp
}

// GetSearchMessagesFilterEnum return the enum type of this object
func (searchMessagesFilterEmpty *SearchMessagesFilterEmpty) GetSearchMessagesFilterEnum() SearchMessagesFilterEnum {
	return SearchMessagesFilterEmptyType
}

// SearchMessagesFilterAnimation Returns only animation messages
type SearchMessagesFilterAnimation struct {
	tdCommon
}

// MessageType return the string telegram-type of SearchMessagesFilterAnimation
func (searchMessagesFilterAnimation *SearchMessagesFilterAnimation) MessageType() string {
	return "searchMessagesFilterAnimation"
}

// NewSearchMessagesFilterAnimation creates a new SearchMessagesFilterAnimation
//
func NewSearchMessagesFilterAnimation() *SearchMessagesFilterAnimation {
	searchMessagesFilterAnimationTemp := SearchMessagesFilterAnimation{
		tdCommon: tdCommon{Type: "searchMessagesFilterAnimation"},
	}

	return &searchMessagesFilterAnimationTemp
}

// GetSearchMessagesFilterEnum return the enum type of this object
func (searchMessagesFilterAnimation *SearchMessagesFilterAnimation) GetSearchMessagesFilterEnum() SearchMessagesFilterEnum {
	return SearchMessagesFilterAnimationType
}

// SearchMessagesFilterAudio Returns only audio messages
type SearchMessagesFilterAudio struct {
	tdCommon
}

// MessageType return the string telegram-type of SearchMessagesFilterAudio
func (searchMessagesFilterAudio *SearchMessagesFilterAudio) MessageType() string {
	return "searchMessagesFilterAudio"
}

// NewSearchMessagesFilterAudio creates a new SearchMessagesFilterAudio
//
func NewSearchMessagesFilterAudio() *SearchMessagesFilterAudio {
	searchMessagesFilterAudioTemp := SearchMessagesFilterAudio{
		tdCommon: tdCommon{Type: "searchMessagesFilterAudio"},
	}

	return &searchMessagesFilterAudioTemp
}

// GetSearchMessagesFilterEnum return the enum type of this object
func (searchMessagesFilterAudio *SearchMessagesFilterAudio) GetSearchMessagesFilterEnum() SearchMessagesFilterEnum {
	return SearchMessagesFilterAudioType
}

// SearchMessagesFilterDocument Returns only document messages
type SearchMessagesFilterDocument struct {
	tdCommon
}

// MessageType return the string telegram-type of SearchMessagesFilterDocument
func (searchMessagesFilterDocument *SearchMessagesFilterDocument) MessageType() string {
	return "searchMessagesFilterDocument"
}

// NewSearchMessagesFilterDocument creates a new SearchMessagesFilterDocument
//
func NewSearchMessagesFilterDocument() *SearchMessagesFilterDocument {
	searchMessagesFilterDocumentTemp := SearchMessagesFilterDocument{
		tdCommon: tdCommon{Type: "searchMessagesFilterDocument"},
	}

	return &searchMessagesFilterDocumentTemp
}

// GetSearchMessagesFilterEnum return the enum type of this object
func (searchMessagesFilterDocument *SearchMessagesFilterDocument) GetSearchMessagesFilterEnum() SearchMessagesFilterEnum {
	return SearchMessagesFilterDocumentType
}

// SearchMessagesFilterPhoto Returns only photo messages
type SearchMessagesFilterPhoto struct {
	tdCommon
}

// MessageType return the string telegram-type of SearchMessagesFilterPhoto
func (searchMessagesFilterPhoto *SearchMessagesFilterPhoto) MessageType() string {
	return "searchMessagesFilterPhoto"
}

// NewSearchMessagesFilterPhoto creates a new SearchMessagesFilterPhoto
//
func NewSearchMessagesFilterPhoto() *SearchMessagesFilterPhoto {
	searchMessagesFilterPhotoTemp := SearchMessagesFilterPhoto{
		tdCommon: tdCommon{Type: "searchMessagesFilterPhoto"},
	}

	return &searchMessagesFilterPhotoTemp
}

// GetSearchMessagesFilterEnum return the enum type of this object
func (searchMessagesFilterPhoto *SearchMessagesFilterPhoto) GetSearchMessagesFilterEnum() SearchMessagesFilterEnum {
	return SearchMessagesFilterPhotoType
}

// SearchMessagesFilterVideo Returns only video messages
type SearchMessagesFilterVideo struct {
	tdCommon
}

// MessageType return the string telegram-type of SearchMessagesFilterVideo
func (searchMessagesFilterVideo *SearchMessagesFilterVideo) MessageType() string {
	return "searchMessagesFilterVideo"
}

// NewSearchMessagesFilterVideo creates a new SearchMessagesFilterVideo
//
func NewSearchMessagesFilterVideo() *SearchMessagesFilterVideo {
	searchMessagesFilterVideoTemp := SearchMessagesFilterVideo{
		tdCommon: tdCommon{Type: "searchMessagesFilterVideo"},
	}

	return &searchMessagesFilterVideoTemp
}

// GetSearchMessagesFilterEnum return the enum type of this object
func (searchMessagesFilterVideo *SearchMessagesFilterVideo) GetSearchMessagesFilterEnum() SearchMessagesFilterEnum {
	return SearchMessagesFilterVideoType
}

// SearchMessagesFilterVoiceNote Returns only voice note messages
type SearchMessagesFilterVoiceNote struct {
	tdCommon
}

// MessageType return the string telegram-type of SearchMessagesFilterVoiceNote
func (searchMessagesFilterVoiceNote *SearchMessagesFilterVoiceNote) MessageType() string {
	return "searchMessagesFilterVoiceNote"
}

// NewSearchMessagesFilterVoiceNote creates a new SearchMessagesFilterVoiceNote
//
func NewSearchMessagesFilterVoiceNote() *SearchMessagesFilterVoiceNote {
	searchMessagesFilterVoiceNoteTemp := SearchMessagesFilterVoiceNote{
		tdCommon: tdCommon{Type: "searchMessagesFilterVoiceNote"},
	}

	return &searchMessagesFilterVoiceNoteTemp
}

// GetSearchMessagesFilterEnum return the enum type of this object
func (searchMessagesFilterVoiceNote *SearchMessagesFilterVoiceNote) GetSearchMessagesFilterEnum() SearchMessagesFilterEnum {
	return SearchMessagesFilterVoiceNoteType
}

// SearchMessagesFilterPhotoAndVideo Returns only photo and video messages
type SearchMessagesFilterPhotoAndVideo struct {
	tdCommon
}

// MessageType return the string telegram-type of SearchMessagesFilterPhotoAndVideo
func (searchMessagesFilterPhotoAndVideo *SearchMessagesFilterPhotoAndVideo) MessageType() string {
	return "searchMessagesFilterPhotoAndVideo"
}

// NewSearchMessagesFilterPhotoAndVideo creates a new SearchMessagesFilterPhotoAndVideo
//
func NewSearchMessagesFilterPhotoAndVideo() *SearchMessagesFilterPhotoAndVideo {
	searchMessagesFilterPhotoAndVideoTemp := SearchMessagesFilterPhotoAndVideo{
		tdCommon: tdCommon{Type: "searchMessagesFilterPhotoAndVideo"},
	}

	return &searchMessagesFilterPhotoAndVideoTemp
}

// GetSearchMessagesFilterEnum return the enum type of this object
func (searchMessagesFilterPhotoAndVideo *SearchMessagesFilterPhotoAndVideo) GetSearchMessagesFilterEnum() SearchMessagesFilterEnum {
	return SearchMessagesFilterPhotoAndVideoType
}

// SearchMessagesFilterUrl Returns only messages containing URLs
type SearchMessagesFilterUrl struct {
	tdCommon
}

// MessageType return the string telegram-type of SearchMessagesFilterUrl
func (searchMessagesFilterUrl *SearchMessagesFilterUrl) MessageType() string {
	return "searchMessagesFilterUrl"
}

// NewSearchMessagesFilterUrl creates a new SearchMessagesFilterUrl
//
func NewSearchMessagesFilterUrl() *SearchMessagesFilterUrl {
	searchMessagesFilterUrlTemp := SearchMessagesFilterUrl{
		tdCommon: tdCommon{Type: "searchMessagesFilterUrl"},
	}

	return &searchMessagesFilterUrlTemp
}

// GetSearchMessagesFilterEnum return the enum type of this object
func (searchMessagesFilterUrl *SearchMessagesFilterUrl) GetSearchMessagesFilterEnum() SearchMessagesFilterEnum {
	return SearchMessagesFilterUrlType
}

// SearchMessagesFilterChatPhoto Returns only messages containing chat photos
type SearchMessagesFilterChatPhoto struct {
	tdCommon
}

// MessageType return the string telegram-type of SearchMessagesFilterChatPhoto
func (searchMessagesFilterChatPhoto *SearchMessagesFilterChatPhoto) MessageType() string {
	return "searchMessagesFilterChatPhoto"
}

// NewSearchMessagesFilterChatPhoto creates a new SearchMessagesFilterChatPhoto
//
func NewSearchMessagesFilterChatPhoto() *SearchMessagesFilterChatPhoto {
	searchMessagesFilterChatPhotoTemp := SearchMessagesFilterChatPhoto{
		tdCommon: tdCommon{Type: "searchMessagesFilterChatPhoto"},
	}

	return &searchMessagesFilterChatPhotoTemp
}

// GetSearchMessagesFilterEnum return the enum type of this object
func (searchMessagesFilterChatPhoto *SearchMessagesFilterChatPhoto) GetSearchMessagesFilterEnum() SearchMessagesFilterEnum {
	return SearchMessagesFilterChatPhotoType
}

// SearchMessagesFilterVideoNote Returns only video note messages
type SearchMessagesFilterVideoNote struct {
	tdCommon
}

// MessageType return the string telegram-type of SearchMessagesFilterVideoNote
func (searchMessagesFilterVideoNote *SearchMessagesFilterVideoNote) MessageType() string {
	return "searchMessagesFilterVideoNote"
}

// NewSearchMessagesFilterVideoNote creates a new SearchMessagesFilterVideoNote
//
func NewSearchMessagesFilterVideoNote() *SearchMessagesFilterVideoNote {
	searchMessagesFilterVideoNoteTemp := SearchMessagesFilterVideoNote{
		tdCommon: tdCommon{Type: "searchMessagesFilterVideoNote"},
	}

	return &searchMessagesFilterVideoNoteTemp
}

// GetSearchMessagesFilterEnum return the enum type of this object
func (searchMessagesFilterVideoNote *SearchMessagesFilterVideoNote) GetSearchMessagesFilterEnum() SearchMessagesFilterEnum {
	return SearchMessagesFilterVideoNoteType
}

// SearchMessagesFilterVoiceAndVideoNote Returns only voice and video note messages
type SearchMessagesFilterVoiceAndVideoNote struct {
	tdCommon
}

// MessageType return the string telegram-type of SearchMessagesFilterVoiceAndVideoNote
func (searchMessagesFilterVoiceAndVideoNote *SearchMessagesFilterVoiceAndVideoNote) MessageType() string {
	return "searchMessagesFilterVoiceAndVideoNote"
}

// NewSearchMessagesFilterVoiceAndVideoNote creates a new SearchMessagesFilterVoiceAndVideoNote
//
func NewSearchMessagesFilterVoiceAndVideoNote() *SearchMessagesFilterVoiceAndVideoNote {
	searchMessagesFilterVoiceAndVideoNoteTemp := SearchMessagesFilterVoiceAndVideoNote{
		tdCommon: tdCommon{Type: "searchMessagesFilterVoiceAndVideoNote"},
	}

	return &searchMessagesFilterVoiceAndVideoNoteTemp
}

// GetSearchMessagesFilterEnum return the enum type of this object
func (searchMessagesFilterVoiceAndVideoNote *SearchMessagesFilterVoiceAndVideoNote) GetSearchMessagesFilterEnum() SearchMessagesFilterEnum {
	return SearchMessagesFilterVoiceAndVideoNoteType
}

// SearchMessagesFilterMention Returns only messages with mentions of the current user, or messages that are replies to their messages
type SearchMessagesFilterMention struct {
	tdCommon
}

// MessageType return the string telegram-type of SearchMessagesFilterMention
func (searchMessagesFilterMention *SearchMessagesFilterMention) MessageType() string {
	return "searchMessagesFilterMention"
}

// NewSearchMessagesFilterMention creates a new SearchMessagesFilterMention
//
func NewSearchMessagesFilterMention() *SearchMessagesFilterMention {
	searchMessagesFilterMentionTemp := SearchMessagesFilterMention{
		tdCommon: tdCommon{Type: "searchMessagesFilterMention"},
	}

	return &searchMessagesFilterMentionTemp
}

// GetSearchMessagesFilterEnum return the enum type of this object
func (searchMessagesFilterMention *SearchMessagesFilterMention) GetSearchMessagesFilterEnum() SearchMessagesFilterEnum {
	return SearchMessagesFilterMentionType
}

// SearchMessagesFilterUnreadMention Returns only messages with unread mentions of the current user, or messages that are replies to their messages. When using this filter the results can't be additionally filtered by a query, a message thread or by the sending user
type SearchMessagesFilterUnreadMention struct {
	tdCommon
}

// MessageType return the string telegram-type of SearchMessagesFilterUnreadMention
func (searchMessagesFilterUnreadMention *SearchMessagesFilterUnreadMention) MessageType() string {
	return "searchMessagesFilterUnreadMention"
}

// NewSearchMessagesFilterUnreadMention creates a new SearchMessagesFilterUnreadMention
//
func NewSearchMessagesFilterUnreadMention() *SearchMessagesFilterUnreadMention {
	searchMessagesFilterUnreadMentionTemp := SearchMessagesFilterUnreadMention{
		tdCommon: tdCommon{Type: "searchMessagesFilterUnreadMention"},
	}

	return &searchMessagesFilterUnreadMentionTemp
}

// GetSearchMessagesFilterEnum return the enum type of this object
func (searchMessagesFilterUnreadMention *SearchMessagesFilterUnreadMention) GetSearchMessagesFilterEnum() SearchMessagesFilterEnum {
	return SearchMessagesFilterUnreadMentionType
}

// SearchMessagesFilterFailedToSend Returns only failed to send messages. This filter can be used only if the message database is used
type SearchMessagesFilterFailedToSend struct {
	tdCommon
}

// MessageType return the string telegram-type of SearchMessagesFilterFailedToSend
func (searchMessagesFilterFailedToSend *SearchMessagesFilterFailedToSend) MessageType() string {
	return "searchMessagesFilterFailedToSend"
}

// NewSearchMessagesFilterFailedToSend creates a new SearchMessagesFilterFailedToSend
//
func NewSearchMessagesFilterFailedToSend() *SearchMessagesFilterFailedToSend {
	searchMessagesFilterFailedToSendTemp := SearchMessagesFilterFailedToSend{
		tdCommon: tdCommon{Type: "searchMessagesFilterFailedToSend"},
	}

	return &searchMessagesFilterFailedToSendTemp
}

// GetSearchMessagesFilterEnum return the enum type of this object
func (searchMessagesFilterFailedToSend *SearchMessagesFilterFailedToSend) GetSearchMessagesFilterEnum() SearchMessagesFilterEnum {
	return SearchMessagesFilterFailedToSendType
}

// SearchMessagesFilterPinned Returns only pinned messages
type SearchMessagesFilterPinned struct {
	tdCommon
}

// MessageType return the string telegram-type of SearchMessagesFilterPinned
func (searchMessagesFilterPinned *SearchMessagesFilterPinned) MessageType() string {
	return "searchMessagesFilterPinned"
}

// NewSearchMessagesFilterPinned creates a new SearchMessagesFilterPinned
//
func NewSearchMessagesFilterPinned() *SearchMessagesFilterPinned {
	searchMessagesFilterPinnedTemp := SearchMessagesFilterPinned{
		tdCommon: tdCommon{Type: "searchMessagesFilterPinned"},
	}

	return &searchMessagesFilterPinnedTemp
}

// GetSearchMessagesFilterEnum return the enum type of this object
func (searchMessagesFilterPinned *SearchMessagesFilterPinned) GetSearchMessagesFilterEnum() SearchMessagesFilterEnum {
	return SearchMessagesFilterPinnedType
}
