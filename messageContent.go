package tdlib

import (
	"encoding/json"
	"fmt"
)

// MessageContent Contains the content of a message
type MessageContent interface {
	GetMessageContentEnum() MessageContentEnum
}

// MessageContentEnum Alias for abstract MessageContent 'Sub-Classes', used as constant-enum here
type MessageContentEnum string

// MessageContent enums
const (
	MessageTextType                        MessageContentEnum = "messageText"
	MessageAnimationType                   MessageContentEnum = "messageAnimation"
	MessageAudioType                       MessageContentEnum = "messageAudio"
	MessageDocumentType                    MessageContentEnum = "messageDocument"
	MessagePhotoType                       MessageContentEnum = "messagePhoto"
	MessageExpiredPhotoType                MessageContentEnum = "messageExpiredPhoto"
	MessageStickerType                     MessageContentEnum = "messageSticker"
	MessageVideoType                       MessageContentEnum = "messageVideo"
	MessageExpiredVideoType                MessageContentEnum = "messageExpiredVideo"
	MessageVideoNoteType                   MessageContentEnum = "messageVideoNote"
	MessageVoiceNoteType                   MessageContentEnum = "messageVoiceNote"
	MessageLocationType                    MessageContentEnum = "messageLocation"
	MessageVenueType                       MessageContentEnum = "messageVenue"
	MessageContactType                     MessageContentEnum = "messageContact"
	MessageAnimatedEmojiType               MessageContentEnum = "messageAnimatedEmoji"
	MessageDiceType                        MessageContentEnum = "messageDice"
	MessageGameType                        MessageContentEnum = "messageGame"
	MessagePollType                        MessageContentEnum = "messagePoll"
	MessageInvoiceType                     MessageContentEnum = "messageInvoice"
	MessageCallType                        MessageContentEnum = "messageCall"
	MessageVideoChatScheduledType          MessageContentEnum = "messageVideoChatScheduled"
	MessageVideoChatStartedType            MessageContentEnum = "messageVideoChatStarted"
	MessageVideoChatEndedType              MessageContentEnum = "messageVideoChatEnded"
	MessageInviteVideoChatParticipantsType MessageContentEnum = "messageInviteVideoChatParticipants"
	MessageBasicGroupChatCreateType        MessageContentEnum = "messageBasicGroupChatCreate"
	MessageSupergroupChatCreateType        MessageContentEnum = "messageSupergroupChatCreate"
	MessageChatChangeTitleType             MessageContentEnum = "messageChatChangeTitle"
	MessageChatChangePhotoType             MessageContentEnum = "messageChatChangePhoto"
	MessageChatDeletePhotoType             MessageContentEnum = "messageChatDeletePhoto"
	MessageChatAddMembersType              MessageContentEnum = "messageChatAddMembers"
	MessageChatJoinByLinkType              MessageContentEnum = "messageChatJoinByLink"
	MessageChatJoinByRequestType           MessageContentEnum = "messageChatJoinByRequest"
	MessageChatDeleteMemberType            MessageContentEnum = "messageChatDeleteMember"
	MessageChatUpgradeToType               MessageContentEnum = "messageChatUpgradeTo"
	MessageChatUpgradeFromType             MessageContentEnum = "messageChatUpgradeFrom"
	MessagePinMessageType                  MessageContentEnum = "messagePinMessage"
	MessageScreenshotTakenType             MessageContentEnum = "messageScreenshotTaken"
	MessageChatSetThemeType                MessageContentEnum = "messageChatSetTheme"
	MessageChatSetTtlType                  MessageContentEnum = "messageChatSetTtl"
	MessageCustomServiceActionType         MessageContentEnum = "messageCustomServiceAction"
	MessageGameScoreType                   MessageContentEnum = "messageGameScore"
	MessagePaymentSuccessfulType           MessageContentEnum = "messagePaymentSuccessful"
	MessagePaymentSuccessfulBotType        MessageContentEnum = "messagePaymentSuccessfulBot"
	MessageContactRegisteredType           MessageContentEnum = "messageContactRegistered"
	MessageWebsiteConnectedType            MessageContentEnum = "messageWebsiteConnected"
	MessagePassportDataSentType            MessageContentEnum = "messagePassportDataSent"
	MessagePassportDataReceivedType        MessageContentEnum = "messagePassportDataReceived"
	MessageProximityAlertTriggeredType     MessageContentEnum = "messageProximityAlertTriggered"
	MessageUnsupportedType                 MessageContentEnum = "messageUnsupported"
)

func unmarshalMessageContent(rawMsg *json.RawMessage) (MessageContent, error) {

	if rawMsg == nil {
		return nil, nil
	}
	var objMap map[string]interface{}
	err := json.Unmarshal(*rawMsg, &objMap)
	if err != nil {
		return nil, err
	}

	switch MessageContentEnum(objMap["@type"].(string)) {
	case MessageTextType:
		var messageText MessageText
		err := json.Unmarshal(*rawMsg, &messageText)
		return &messageText, err

	case MessageAnimationType:
		var messageAnimation MessageAnimation
		err := json.Unmarshal(*rawMsg, &messageAnimation)
		return &messageAnimation, err

	case MessageAudioType:
		var messageAudio MessageAudio
		err := json.Unmarshal(*rawMsg, &messageAudio)
		return &messageAudio, err

	case MessageDocumentType:
		var messageDocument MessageDocument
		err := json.Unmarshal(*rawMsg, &messageDocument)
		return &messageDocument, err

	case MessagePhotoType:
		var messagePhoto MessagePhoto
		err := json.Unmarshal(*rawMsg, &messagePhoto)
		return &messagePhoto, err

	case MessageExpiredPhotoType:
		var messageExpiredPhoto MessageExpiredPhoto
		err := json.Unmarshal(*rawMsg, &messageExpiredPhoto)
		return &messageExpiredPhoto, err

	case MessageStickerType:
		var messageSticker MessageSticker
		err := json.Unmarshal(*rawMsg, &messageSticker)
		return &messageSticker, err

	case MessageVideoType:
		var messageVideo MessageVideo
		err := json.Unmarshal(*rawMsg, &messageVideo)
		return &messageVideo, err

	case MessageExpiredVideoType:
		var messageExpiredVideo MessageExpiredVideo
		err := json.Unmarshal(*rawMsg, &messageExpiredVideo)
		return &messageExpiredVideo, err

	case MessageVideoNoteType:
		var messageVideoNote MessageVideoNote
		err := json.Unmarshal(*rawMsg, &messageVideoNote)
		return &messageVideoNote, err

	case MessageVoiceNoteType:
		var messageVoiceNote MessageVoiceNote
		err := json.Unmarshal(*rawMsg, &messageVoiceNote)
		return &messageVoiceNote, err

	case MessageLocationType:
		var messageLocation MessageLocation
		err := json.Unmarshal(*rawMsg, &messageLocation)
		return &messageLocation, err

	case MessageVenueType:
		var messageVenue MessageVenue
		err := json.Unmarshal(*rawMsg, &messageVenue)
		return &messageVenue, err

	case MessageContactType:
		var messageContact MessageContact
		err := json.Unmarshal(*rawMsg, &messageContact)
		return &messageContact, err

	case MessageAnimatedEmojiType:
		var messageAnimatedEmoji MessageAnimatedEmoji
		err := json.Unmarshal(*rawMsg, &messageAnimatedEmoji)
		return &messageAnimatedEmoji, err

	case MessageDiceType:
		var messageDice MessageDice
		err := json.Unmarshal(*rawMsg, &messageDice)
		return &messageDice, err

	case MessageGameType:
		var messageGame MessageGame
		err := json.Unmarshal(*rawMsg, &messageGame)
		return &messageGame, err

	case MessagePollType:
		var messagePoll MessagePoll
		err := json.Unmarshal(*rawMsg, &messagePoll)
		return &messagePoll, err

	case MessageInvoiceType:
		var messageInvoice MessageInvoice
		err := json.Unmarshal(*rawMsg, &messageInvoice)
		return &messageInvoice, err

	case MessageCallType:
		var messageCall MessageCall
		err := json.Unmarshal(*rawMsg, &messageCall)
		return &messageCall, err

	case MessageVideoChatScheduledType:
		var messageVideoChatScheduled MessageVideoChatScheduled
		err := json.Unmarshal(*rawMsg, &messageVideoChatScheduled)
		return &messageVideoChatScheduled, err

	case MessageVideoChatStartedType:
		var messageVideoChatStarted MessageVideoChatStarted
		err := json.Unmarshal(*rawMsg, &messageVideoChatStarted)
		return &messageVideoChatStarted, err

	case MessageVideoChatEndedType:
		var messageVideoChatEnded MessageVideoChatEnded
		err := json.Unmarshal(*rawMsg, &messageVideoChatEnded)
		return &messageVideoChatEnded, err

	case MessageInviteVideoChatParticipantsType:
		var messageInviteVideoChatParticipants MessageInviteVideoChatParticipants
		err := json.Unmarshal(*rawMsg, &messageInviteVideoChatParticipants)
		return &messageInviteVideoChatParticipants, err

	case MessageBasicGroupChatCreateType:
		var messageBasicGroupChatCreate MessageBasicGroupChatCreate
		err := json.Unmarshal(*rawMsg, &messageBasicGroupChatCreate)
		return &messageBasicGroupChatCreate, err

	case MessageSupergroupChatCreateType:
		var messageSupergroupChatCreate MessageSupergroupChatCreate
		err := json.Unmarshal(*rawMsg, &messageSupergroupChatCreate)
		return &messageSupergroupChatCreate, err

	case MessageChatChangeTitleType:
		var messageChatChangeTitle MessageChatChangeTitle
		err := json.Unmarshal(*rawMsg, &messageChatChangeTitle)
		return &messageChatChangeTitle, err

	case MessageChatChangePhotoType:
		var messageChatChangePhoto MessageChatChangePhoto
		err := json.Unmarshal(*rawMsg, &messageChatChangePhoto)
		return &messageChatChangePhoto, err

	case MessageChatDeletePhotoType:
		var messageChatDeletePhoto MessageChatDeletePhoto
		err := json.Unmarshal(*rawMsg, &messageChatDeletePhoto)
		return &messageChatDeletePhoto, err

	case MessageChatAddMembersType:
		var messageChatAddMembers MessageChatAddMembers
		err := json.Unmarshal(*rawMsg, &messageChatAddMembers)
		return &messageChatAddMembers, err

	case MessageChatJoinByLinkType:
		var messageChatJoinByLink MessageChatJoinByLink
		err := json.Unmarshal(*rawMsg, &messageChatJoinByLink)
		return &messageChatJoinByLink, err

	case MessageChatJoinByRequestType:
		var messageChatJoinByRequest MessageChatJoinByRequest
		err := json.Unmarshal(*rawMsg, &messageChatJoinByRequest)
		return &messageChatJoinByRequest, err

	case MessageChatDeleteMemberType:
		var messageChatDeleteMember MessageChatDeleteMember
		err := json.Unmarshal(*rawMsg, &messageChatDeleteMember)
		return &messageChatDeleteMember, err

	case MessageChatUpgradeToType:
		var messageChatUpgradeTo MessageChatUpgradeTo
		err := json.Unmarshal(*rawMsg, &messageChatUpgradeTo)
		return &messageChatUpgradeTo, err

	case MessageChatUpgradeFromType:
		var messageChatUpgradeFrom MessageChatUpgradeFrom
		err := json.Unmarshal(*rawMsg, &messageChatUpgradeFrom)
		return &messageChatUpgradeFrom, err

	case MessagePinMessageType:
		var messagePinMessage MessagePinMessage
		err := json.Unmarshal(*rawMsg, &messagePinMessage)
		return &messagePinMessage, err

	case MessageScreenshotTakenType:
		var messageScreenshotTaken MessageScreenshotTaken
		err := json.Unmarshal(*rawMsg, &messageScreenshotTaken)
		return &messageScreenshotTaken, err

	case MessageChatSetThemeType:
		var messageChatSetTheme MessageChatSetTheme
		err := json.Unmarshal(*rawMsg, &messageChatSetTheme)
		return &messageChatSetTheme, err

	case MessageChatSetTtlType:
		var messageChatSetTtl MessageChatSetTtl
		err := json.Unmarshal(*rawMsg, &messageChatSetTtl)
		return &messageChatSetTtl, err

	case MessageCustomServiceActionType:
		var messageCustomServiceAction MessageCustomServiceAction
		err := json.Unmarshal(*rawMsg, &messageCustomServiceAction)
		return &messageCustomServiceAction, err

	case MessageGameScoreType:
		var messageGameScore MessageGameScore
		err := json.Unmarshal(*rawMsg, &messageGameScore)
		return &messageGameScore, err

	case MessagePaymentSuccessfulType:
		var messagePaymentSuccessful MessagePaymentSuccessful
		err := json.Unmarshal(*rawMsg, &messagePaymentSuccessful)
		return &messagePaymentSuccessful, err

	case MessagePaymentSuccessfulBotType:
		var messagePaymentSuccessfulBot MessagePaymentSuccessfulBot
		err := json.Unmarshal(*rawMsg, &messagePaymentSuccessfulBot)
		return &messagePaymentSuccessfulBot, err

	case MessageContactRegisteredType:
		var messageContactRegistered MessageContactRegistered
		err := json.Unmarshal(*rawMsg, &messageContactRegistered)
		return &messageContactRegistered, err

	case MessageWebsiteConnectedType:
		var messageWebsiteConnected MessageWebsiteConnected
		err := json.Unmarshal(*rawMsg, &messageWebsiteConnected)
		return &messageWebsiteConnected, err

	case MessagePassportDataSentType:
		var messagePassportDataSent MessagePassportDataSent
		err := json.Unmarshal(*rawMsg, &messagePassportDataSent)
		return &messagePassportDataSent, err

	case MessagePassportDataReceivedType:
		var messagePassportDataReceived MessagePassportDataReceived
		err := json.Unmarshal(*rawMsg, &messagePassportDataReceived)
		return &messagePassportDataReceived, err

	case MessageProximityAlertTriggeredType:
		var messageProximityAlertTriggered MessageProximityAlertTriggered
		err := json.Unmarshal(*rawMsg, &messageProximityAlertTriggered)
		return &messageProximityAlertTriggered, err

	case MessageUnsupportedType:
		var messageUnsupported MessageUnsupported
		err := json.Unmarshal(*rawMsg, &messageUnsupported)
		return &messageUnsupported, err

	default:
		return nil, fmt.Errorf("Error UnMarshaling, unknown type:" + objMap["@type"].(string))
	}
}

// MessageText A text message
type MessageText struct {
	tdCommon
	Text    *FormattedText `json:"text"`     // Text of the message
	WebPage *WebPage       `json:"web_page"` // A preview of the web page that's mentioned in the text; may be null
}

// MessageType return the string telegram-type of MessageText
func (messageText *MessageText) MessageType() string {
	return "messageText"
}

// NewMessageText creates a new MessageText
//
// @param text Text of the message
// @param webPage A preview of the web page that's mentioned in the text; may be null
func NewMessageText(text *FormattedText, webPage *WebPage) *MessageText {
	messageTextTemp := MessageText{
		tdCommon: tdCommon{Type: "messageText"},
		Text:     text,
		WebPage:  webPage,
	}

	return &messageTextTemp
}

// GetMessageContentEnum return the enum type of this object
func (messageText *MessageText) GetMessageContentEnum() MessageContentEnum {
	return MessageTextType
}

// MessageAnimation An animation message (GIF-style).
type MessageAnimation struct {
	tdCommon
	Animation *Animation     `json:"animation"` // The animation description
	Caption   *FormattedText `json:"caption"`   // Animation caption
	IsSecret  bool           `json:"is_secret"` // True, if the animation thumbnail must be blurred and the animation must be shown only while tapped
}

// MessageType return the string telegram-type of MessageAnimation
func (messageAnimation *MessageAnimation) MessageType() string {
	return "messageAnimation"
}

// NewMessageAnimation creates a new MessageAnimation
//
// @param animation The animation description
// @param caption Animation caption
// @param isSecret True, if the animation thumbnail must be blurred and the animation must be shown only while tapped
func NewMessageAnimation(animation *Animation, caption *FormattedText, isSecret bool) *MessageAnimation {
	messageAnimationTemp := MessageAnimation{
		tdCommon:  tdCommon{Type: "messageAnimation"},
		Animation: animation,
		Caption:   caption,
		IsSecret:  isSecret,
	}

	return &messageAnimationTemp
}

// GetMessageContentEnum return the enum type of this object
func (messageAnimation *MessageAnimation) GetMessageContentEnum() MessageContentEnum {
	return MessageAnimationType
}

// MessageAudio An audio message
type MessageAudio struct {
	tdCommon
	Audio   *Audio         `json:"audio"`   // The audio description
	Caption *FormattedText `json:"caption"` // Audio caption
}

// MessageType return the string telegram-type of MessageAudio
func (messageAudio *MessageAudio) MessageType() string {
	return "messageAudio"
}

// NewMessageAudio creates a new MessageAudio
//
// @param audio The audio description
// @param caption Audio caption
func NewMessageAudio(audio *Audio, caption *FormattedText) *MessageAudio {
	messageAudioTemp := MessageAudio{
		tdCommon: tdCommon{Type: "messageAudio"},
		Audio:    audio,
		Caption:  caption,
	}

	return &messageAudioTemp
}

// GetMessageContentEnum return the enum type of this object
func (messageAudio *MessageAudio) GetMessageContentEnum() MessageContentEnum {
	return MessageAudioType
}

// MessageDocument A document message (general file)
type MessageDocument struct {
	tdCommon
	Document *Document      `json:"document"` // The document description
	Caption  *FormattedText `json:"caption"`  // Document caption
}

// MessageType return the string telegram-type of MessageDocument
func (messageDocument *MessageDocument) MessageType() string {
	return "messageDocument"
}

// NewMessageDocument creates a new MessageDocument
//
// @param document The document description
// @param caption Document caption
func NewMessageDocument(document *Document, caption *FormattedText) *MessageDocument {
	messageDocumentTemp := MessageDocument{
		tdCommon: tdCommon{Type: "messageDocument"},
		Document: document,
		Caption:  caption,
	}

	return &messageDocumentTemp
}

// GetMessageContentEnum return the enum type of this object
func (messageDocument *MessageDocument) GetMessageContentEnum() MessageContentEnum {
	return MessageDocumentType
}

// MessagePhoto A photo message
type MessagePhoto struct {
	tdCommon
	Photo    *Photo         `json:"photo"`     // The photo description
	Caption  *FormattedText `json:"caption"`   // Photo caption
	IsSecret bool           `json:"is_secret"` // True, if the photo must be blurred and must be shown only while tapped
}

// MessageType return the string telegram-type of MessagePhoto
func (messagePhoto *MessagePhoto) MessageType() string {
	return "messagePhoto"
}

// NewMessagePhoto creates a new MessagePhoto
//
// @param photo The photo description
// @param caption Photo caption
// @param isSecret True, if the photo must be blurred and must be shown only while tapped
func NewMessagePhoto(photo *Photo, caption *FormattedText, isSecret bool) *MessagePhoto {
	messagePhotoTemp := MessagePhoto{
		tdCommon: tdCommon{Type: "messagePhoto"},
		Photo:    photo,
		Caption:  caption,
		IsSecret: isSecret,
	}

	return &messagePhotoTemp
}

// GetMessageContentEnum return the enum type of this object
func (messagePhoto *MessagePhoto) GetMessageContentEnum() MessageContentEnum {
	return MessagePhotoType
}

// MessageExpiredPhoto An expired photo message (self-destructed after TTL has elapsed)
type MessageExpiredPhoto struct {
	tdCommon
}

// MessageType return the string telegram-type of MessageExpiredPhoto
func (messageExpiredPhoto *MessageExpiredPhoto) MessageType() string {
	return "messageExpiredPhoto"
}

// NewMessageExpiredPhoto creates a new MessageExpiredPhoto
//
func NewMessageExpiredPhoto() *MessageExpiredPhoto {
	messageExpiredPhotoTemp := MessageExpiredPhoto{
		tdCommon: tdCommon{Type: "messageExpiredPhoto"},
	}

	return &messageExpiredPhotoTemp
}

// GetMessageContentEnum return the enum type of this object
func (messageExpiredPhoto *MessageExpiredPhoto) GetMessageContentEnum() MessageContentEnum {
	return MessageExpiredPhotoType
}

// MessageSticker A sticker message
type MessageSticker struct {
	tdCommon
	Sticker *Sticker `json:"sticker"` // The sticker description
}

// MessageType return the string telegram-type of MessageSticker
func (messageSticker *MessageSticker) MessageType() string {
	return "messageSticker"
}

// NewMessageSticker creates a new MessageSticker
//
// @param sticker The sticker description
func NewMessageSticker(sticker *Sticker) *MessageSticker {
	messageStickerTemp := MessageSticker{
		tdCommon: tdCommon{Type: "messageSticker"},
		Sticker:  sticker,
	}

	return &messageStickerTemp
}

// GetMessageContentEnum return the enum type of this object
func (messageSticker *MessageSticker) GetMessageContentEnum() MessageContentEnum {
	return MessageStickerType
}

// MessageVideo A video message
type MessageVideo struct {
	tdCommon
	Video    *Video         `json:"video"`     // The video description
	Caption  *FormattedText `json:"caption"`   // Video caption
	IsSecret bool           `json:"is_secret"` // True, if the video thumbnail must be blurred and the video must be shown only while tapped
}

// MessageType return the string telegram-type of MessageVideo
func (messageVideo *MessageVideo) MessageType() string {
	return "messageVideo"
}

// NewMessageVideo creates a new MessageVideo
//
// @param video The video description
// @param caption Video caption
// @param isSecret True, if the video thumbnail must be blurred and the video must be shown only while tapped
func NewMessageVideo(video *Video, caption *FormattedText, isSecret bool) *MessageVideo {
	messageVideoTemp := MessageVideo{
		tdCommon: tdCommon{Type: "messageVideo"},
		Video:    video,
		Caption:  caption,
		IsSecret: isSecret,
	}

	return &messageVideoTemp
}

// GetMessageContentEnum return the enum type of this object
func (messageVideo *MessageVideo) GetMessageContentEnum() MessageContentEnum {
	return MessageVideoType
}

// MessageExpiredVideo An expired video message (self-destructed after TTL has elapsed)
type MessageExpiredVideo struct {
	tdCommon
}

// MessageType return the string telegram-type of MessageExpiredVideo
func (messageExpiredVideo *MessageExpiredVideo) MessageType() string {
	return "messageExpiredVideo"
}

// NewMessageExpiredVideo creates a new MessageExpiredVideo
//
func NewMessageExpiredVideo() *MessageExpiredVideo {
	messageExpiredVideoTemp := MessageExpiredVideo{
		tdCommon: tdCommon{Type: "messageExpiredVideo"},
	}

	return &messageExpiredVideoTemp
}

// GetMessageContentEnum return the enum type of this object
func (messageExpiredVideo *MessageExpiredVideo) GetMessageContentEnum() MessageContentEnum {
	return MessageExpiredVideoType
}

// MessageVideoNote A video note message
type MessageVideoNote struct {
	tdCommon
	VideoNote *VideoNote `json:"video_note"` // The video note description
	IsViewed  bool       `json:"is_viewed"`  // True, if at least one of the recipients has viewed the video note
	IsSecret  bool       `json:"is_secret"`  // True, if the video note thumbnail must be blurred and the video note must be shown only while tapped
}

// MessageType return the string telegram-type of MessageVideoNote
func (messageVideoNote *MessageVideoNote) MessageType() string {
	return "messageVideoNote"
}

// NewMessageVideoNote creates a new MessageVideoNote
//
// @param videoNote The video note description
// @param isViewed True, if at least one of the recipients has viewed the video note
// @param isSecret True, if the video note thumbnail must be blurred and the video note must be shown only while tapped
func NewMessageVideoNote(videoNote *VideoNote, isViewed bool, isSecret bool) *MessageVideoNote {
	messageVideoNoteTemp := MessageVideoNote{
		tdCommon:  tdCommon{Type: "messageVideoNote"},
		VideoNote: videoNote,
		IsViewed:  isViewed,
		IsSecret:  isSecret,
	}

	return &messageVideoNoteTemp
}

// GetMessageContentEnum return the enum type of this object
func (messageVideoNote *MessageVideoNote) GetMessageContentEnum() MessageContentEnum {
	return MessageVideoNoteType
}

// MessageVoiceNote A voice note message
type MessageVoiceNote struct {
	tdCommon
	VoiceNote  *VoiceNote     `json:"voice_note"`  // The voice note description
	Caption    *FormattedText `json:"caption"`     // Voice note caption
	IsListened bool           `json:"is_listened"` // True, if at least one of the recipients has listened to the voice note
}

// MessageType return the string telegram-type of MessageVoiceNote
func (messageVoiceNote *MessageVoiceNote) MessageType() string {
	return "messageVoiceNote"
}

// NewMessageVoiceNote creates a new MessageVoiceNote
//
// @param voiceNote The voice note description
// @param caption Voice note caption
// @param isListened True, if at least one of the recipients has listened to the voice note
func NewMessageVoiceNote(voiceNote *VoiceNote, caption *FormattedText, isListened bool) *MessageVoiceNote {
	messageVoiceNoteTemp := MessageVoiceNote{
		tdCommon:   tdCommon{Type: "messageVoiceNote"},
		VoiceNote:  voiceNote,
		Caption:    caption,
		IsListened: isListened,
	}

	return &messageVoiceNoteTemp
}

// GetMessageContentEnum return the enum type of this object
func (messageVoiceNote *MessageVoiceNote) GetMessageContentEnum() MessageContentEnum {
	return MessageVoiceNoteType
}

// MessageLocation A message with a location
type MessageLocation struct {
	tdCommon
	Location             *Location `json:"location"`               // The location description
	LivePeriod           int32     `json:"live_period"`            // Time relative to the message send date, for which the location can be updated, in seconds
	ExpiresIn            int32     `json:"expires_in"`             // Left time for which the location can be updated, in seconds. updateMessageContent is not sent when this field changes
	Heading              int32     `json:"heading"`                // For live locations, a direction in which the location moves, in degrees; 1-360. If 0 the direction is unknown
	ProximityAlertRadius int32     `json:"proximity_alert_radius"` // For live locations, a maximum distance to another chat member for proximity alerts, in meters (0-100000). 0 if the notification is disabled. Available only for the message sender
}

// MessageType return the string telegram-type of MessageLocation
func (messageLocation *MessageLocation) MessageType() string {
	return "messageLocation"
}

// NewMessageLocation creates a new MessageLocation
//
// @param location The location description
// @param livePeriod Time relative to the message send date, for which the location can be updated, in seconds
// @param expiresIn Left time for which the location can be updated, in seconds. updateMessageContent is not sent when this field changes
// @param heading For live locations, a direction in which the location moves, in degrees; 1-360. If 0 the direction is unknown
// @param proximityAlertRadius For live locations, a maximum distance to another chat member for proximity alerts, in meters (0-100000). 0 if the notification is disabled. Available only for the message sender
func NewMessageLocation(location *Location, livePeriod int32, expiresIn int32, heading int32, proximityAlertRadius int32) *MessageLocation {
	messageLocationTemp := MessageLocation{
		tdCommon:             tdCommon{Type: "messageLocation"},
		Location:             location,
		LivePeriod:           livePeriod,
		ExpiresIn:            expiresIn,
		Heading:              heading,
		ProximityAlertRadius: proximityAlertRadius,
	}

	return &messageLocationTemp
}

// GetMessageContentEnum return the enum type of this object
func (messageLocation *MessageLocation) GetMessageContentEnum() MessageContentEnum {
	return MessageLocationType
}

// MessageVenue A message with information about a venue
type MessageVenue struct {
	tdCommon
	Venue *Venue `json:"venue"` // The venue description
}

// MessageType return the string telegram-type of MessageVenue
func (messageVenue *MessageVenue) MessageType() string {
	return "messageVenue"
}

// NewMessageVenue creates a new MessageVenue
//
// @param venue The venue description
func NewMessageVenue(venue *Venue) *MessageVenue {
	messageVenueTemp := MessageVenue{
		tdCommon: tdCommon{Type: "messageVenue"},
		Venue:    venue,
	}

	return &messageVenueTemp
}

// GetMessageContentEnum return the enum type of this object
func (messageVenue *MessageVenue) GetMessageContentEnum() MessageContentEnum {
	return MessageVenueType
}

// MessageContact A message with a user contact
type MessageContact struct {
	tdCommon
	Contact *Contact `json:"contact"` // The contact description
}

// MessageType return the string telegram-type of MessageContact
func (messageContact *MessageContact) MessageType() string {
	return "messageContact"
}

// NewMessageContact creates a new MessageContact
//
// @param contact The contact description
func NewMessageContact(contact *Contact) *MessageContact {
	messageContactTemp := MessageContact{
		tdCommon: tdCommon{Type: "messageContact"},
		Contact:  contact,
	}

	return &messageContactTemp
}

// GetMessageContentEnum return the enum type of this object
func (messageContact *MessageContact) GetMessageContentEnum() MessageContentEnum {
	return MessageContactType
}

// MessageAnimatedEmoji A message with an animated emoji
type MessageAnimatedEmoji struct {
	tdCommon
	AnimatedEmoji *AnimatedEmoji `json:"animated_emoji"` // The animated emoji
	Emoji         string         `json:"emoji"`          // The corresponding emoji
}

// MessageType return the string telegram-type of MessageAnimatedEmoji
func (messageAnimatedEmoji *MessageAnimatedEmoji) MessageType() string {
	return "messageAnimatedEmoji"
}

// NewMessageAnimatedEmoji creates a new MessageAnimatedEmoji
//
// @param animatedEmoji The animated emoji
// @param emoji The corresponding emoji
func NewMessageAnimatedEmoji(animatedEmoji *AnimatedEmoji, emoji string) *MessageAnimatedEmoji {
	messageAnimatedEmojiTemp := MessageAnimatedEmoji{
		tdCommon:      tdCommon{Type: "messageAnimatedEmoji"},
		AnimatedEmoji: animatedEmoji,
		Emoji:         emoji,
	}

	return &messageAnimatedEmojiTemp
}

// GetMessageContentEnum return the enum type of this object
func (messageAnimatedEmoji *MessageAnimatedEmoji) GetMessageContentEnum() MessageContentEnum {
	return MessageAnimatedEmojiType
}

// MessageDice A dice message. The dice value is randomly generated by the server
type MessageDice struct {
	tdCommon
	InitialState                DiceStickers `json:"initial_state"`                  // The animated stickers with the initial dice animation; may be null if unknown. updateMessageContent will be sent when the sticker became known
	FinalState                  DiceStickers `json:"final_state"`                    // The animated stickers with the final dice animation; may be null if unknown. updateMessageContent will be sent when the sticker became known
	Emoji                       string       `json:"emoji"`                          // Emoji on which the dice throw animation is based
	Value                       int32        `json:"value"`                          // The dice value. If the value is 0, the dice don't have final state yet
	SuccessAnimationFrameNumber int32        `json:"success_animation_frame_number"` // Number of frame after which a success animation like a shower of confetti needs to be shown on updateMessageSendSucceeded
}

// MessageType return the string telegram-type of MessageDice
func (messageDice *MessageDice) MessageType() string {
	return "messageDice"
}

// NewMessageDice creates a new MessageDice
//
// @param initialState The animated stickers with the initial dice animation; may be null if unknown. updateMessageContent will be sent when the sticker became known
// @param finalState The animated stickers with the final dice animation; may be null if unknown. updateMessageContent will be sent when the sticker became known
// @param emoji Emoji on which the dice throw animation is based
// @param value The dice value. If the value is 0, the dice don't have final state yet
// @param successAnimationFrameNumber Number of frame after which a success animation like a shower of confetti needs to be shown on updateMessageSendSucceeded
func NewMessageDice(initialState DiceStickers, finalState DiceStickers, emoji string, value int32, successAnimationFrameNumber int32) *MessageDice {
	messageDiceTemp := MessageDice{
		tdCommon:                    tdCommon{Type: "messageDice"},
		InitialState:                initialState,
		FinalState:                  finalState,
		Emoji:                       emoji,
		Value:                       value,
		SuccessAnimationFrameNumber: successAnimationFrameNumber,
	}

	return &messageDiceTemp
}

// UnmarshalJSON unmarshal to json
func (messageDice *MessageDice) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	tempObj := struct {
		tdCommon
		Emoji                       string `json:"emoji"`                          // Emoji on which the dice throw animation is based
		Value                       int32  `json:"value"`                          // The dice value. If the value is 0, the dice don't have final state yet
		SuccessAnimationFrameNumber int32  `json:"success_animation_frame_number"` // Number of frame after which a success animation like a shower of confetti needs to be shown on updateMessageSendSucceeded
	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	messageDice.tdCommon = tempObj.tdCommon
	messageDice.Emoji = tempObj.Emoji
	messageDice.Value = tempObj.Value
	messageDice.SuccessAnimationFrameNumber = tempObj.SuccessAnimationFrameNumber

	fieldInitialState, _ := unmarshalDiceStickers(objMap["initial_state"])
	messageDice.InitialState = fieldInitialState

	fieldFinalState, _ := unmarshalDiceStickers(objMap["final_state"])
	messageDice.FinalState = fieldFinalState

	return nil
}

// GetMessageContentEnum return the enum type of this object
func (messageDice *MessageDice) GetMessageContentEnum() MessageContentEnum {
	return MessageDiceType
}

// MessageGame A message with a game
type MessageGame struct {
	tdCommon
	Game *Game `json:"game"` // The game description
}

// MessageType return the string telegram-type of MessageGame
func (messageGame *MessageGame) MessageType() string {
	return "messageGame"
}

// NewMessageGame creates a new MessageGame
//
// @param game The game description
func NewMessageGame(game *Game) *MessageGame {
	messageGameTemp := MessageGame{
		tdCommon: tdCommon{Type: "messageGame"},
		Game:     game,
	}

	return &messageGameTemp
}

// GetMessageContentEnum return the enum type of this object
func (messageGame *MessageGame) GetMessageContentEnum() MessageContentEnum {
	return MessageGameType
}

// MessagePoll A message with a poll
type MessagePoll struct {
	tdCommon
	Poll *Poll `json:"poll"` // The poll description
}

// MessageType return the string telegram-type of MessagePoll
func (messagePoll *MessagePoll) MessageType() string {
	return "messagePoll"
}

// NewMessagePoll creates a new MessagePoll
//
// @param poll The poll description
func NewMessagePoll(poll *Poll) *MessagePoll {
	messagePollTemp := MessagePoll{
		tdCommon: tdCommon{Type: "messagePoll"},
		Poll:     poll,
	}

	return &messagePollTemp
}

// GetMessageContentEnum return the enum type of this object
func (messagePoll *MessagePoll) GetMessageContentEnum() MessageContentEnum {
	return MessagePollType
}

// MessageInvoice A message with an invoice from a bot
type MessageInvoice struct {
	tdCommon
	Title               string `json:"title"`                 // Product title
	Description         string `json:"description"`           // Product description
	Photo               *Photo `json:"photo"`                 // Product photo; may be null
	Currency            string `json:"currency"`              // Currency for the product price
	TotalAmount         int64  `json:"total_amount"`          // Product total price in the smallest units of the currency
	StartParameter      string `json:"start_parameter"`       // Unique invoice bot start_parameter. To share an invoice use the URL https://t.me/{bot_username}?start={start_parameter}
	IsTest              bool   `json:"is_test"`               // True, if the invoice is a test invoice
	NeedShippingAddress bool   `json:"need_shipping_address"` // True, if the shipping address must be specified
	ReceiptMessageId    int64  `json:"receipt_message_id"`    // The identifier of the message with the receipt, after the product has been purchased
}

// MessageType return the string telegram-type of MessageInvoice
func (messageInvoice *MessageInvoice) MessageType() string {
	return "messageInvoice"
}

// NewMessageInvoice creates a new MessageInvoice
//
// @param title Product title
// @param description Product description
// @param photo Product photo; may be null
// @param currency Currency for the product price
// @param totalAmount Product total price in the smallest units of the currency
// @param startParameter Unique invoice bot start_parameter. To share an invoice use the URL https://t.me/{bot_username}?start={start_parameter}
// @param isTest True, if the invoice is a test invoice
// @param needShippingAddress True, if the shipping address must be specified
// @param receiptMessageId The identifier of the message with the receipt, after the product has been purchased
func NewMessageInvoice(title string, description string, photo *Photo, currency string, totalAmount int64, startParameter string, isTest bool, needShippingAddress bool, receiptMessageId int64) *MessageInvoice {
	messageInvoiceTemp := MessageInvoice{
		tdCommon:            tdCommon{Type: "messageInvoice"},
		Title:               title,
		Description:         description,
		Photo:               photo,
		Currency:            currency,
		TotalAmount:         totalAmount,
		StartParameter:      startParameter,
		IsTest:              isTest,
		NeedShippingAddress: needShippingAddress,
		ReceiptMessageId:    receiptMessageId,
	}

	return &messageInvoiceTemp
}

// GetMessageContentEnum return the enum type of this object
func (messageInvoice *MessageInvoice) GetMessageContentEnum() MessageContentEnum {
	return MessageInvoiceType
}

// MessageCall A message with information about an ended call
type MessageCall struct {
	tdCommon
	IsVideo       bool              `json:"is_video"`       // True, if the call was a video call
	DiscardReason CallDiscardReason `json:"discard_reason"` // Reason why the call was discarded
	Duration      int32             `json:"duration"`       // Call duration, in seconds
}

// MessageType return the string telegram-type of MessageCall
func (messageCall *MessageCall) MessageType() string {
	return "messageCall"
}

// NewMessageCall creates a new MessageCall
//
// @param isVideo True, if the call was a video call
// @param discardReason Reason why the call was discarded
// @param duration Call duration, in seconds
func NewMessageCall(isVideo bool, discardReason CallDiscardReason, duration int32) *MessageCall {
	messageCallTemp := MessageCall{
		tdCommon:      tdCommon{Type: "messageCall"},
		IsVideo:       isVideo,
		DiscardReason: discardReason,
		Duration:      duration,
	}

	return &messageCallTemp
}

// UnmarshalJSON unmarshal to json
func (messageCall *MessageCall) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	tempObj := struct {
		tdCommon
		IsVideo  bool  `json:"is_video"` // True, if the call was a video call
		Duration int32 `json:"duration"` // Call duration, in seconds
	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	messageCall.tdCommon = tempObj.tdCommon
	messageCall.IsVideo = tempObj.IsVideo
	messageCall.Duration = tempObj.Duration

	fieldDiscardReason, _ := unmarshalCallDiscardReason(objMap["discard_reason"])
	messageCall.DiscardReason = fieldDiscardReason

	return nil
}

// GetMessageContentEnum return the enum type of this object
func (messageCall *MessageCall) GetMessageContentEnum() MessageContentEnum {
	return MessageCallType
}

// MessageVideoChatScheduled A new video chat was scheduled
type MessageVideoChatScheduled struct {
	tdCommon
	GroupCallId int32 `json:"group_call_id"` // Identifier of the video chat. The video chat can be received through the method getGroupCall
	StartDate   int32 `json:"start_date"`    // Point in time (Unix timestamp) when the group call is supposed to be started by an administrator
}

// MessageType return the string telegram-type of MessageVideoChatScheduled
func (messageVideoChatScheduled *MessageVideoChatScheduled) MessageType() string {
	return "messageVideoChatScheduled"
}

// NewMessageVideoChatScheduled creates a new MessageVideoChatScheduled
//
// @param groupCallId Identifier of the video chat. The video chat can be received through the method getGroupCall
// @param startDate Point in time (Unix timestamp) when the group call is supposed to be started by an administrator
func NewMessageVideoChatScheduled(groupCallId int32, startDate int32) *MessageVideoChatScheduled {
	messageVideoChatScheduledTemp := MessageVideoChatScheduled{
		tdCommon:    tdCommon{Type: "messageVideoChatScheduled"},
		GroupCallId: groupCallId,
		StartDate:   startDate,
	}

	return &messageVideoChatScheduledTemp
}

// GetMessageContentEnum return the enum type of this object
func (messageVideoChatScheduled *MessageVideoChatScheduled) GetMessageContentEnum() MessageContentEnum {
	return MessageVideoChatScheduledType
}

// MessageVideoChatStarted A newly created video chat
type MessageVideoChatStarted struct {
	tdCommon
	GroupCallId int32 `json:"group_call_id"` // Identifier of the video chat. The video chat can be received through the method getGroupCall
}

// MessageType return the string telegram-type of MessageVideoChatStarted
func (messageVideoChatStarted *MessageVideoChatStarted) MessageType() string {
	return "messageVideoChatStarted"
}

// NewMessageVideoChatStarted creates a new MessageVideoChatStarted
//
// @param groupCallId Identifier of the video chat. The video chat can be received through the method getGroupCall
func NewMessageVideoChatStarted(groupCallId int32) *MessageVideoChatStarted {
	messageVideoChatStartedTemp := MessageVideoChatStarted{
		tdCommon:    tdCommon{Type: "messageVideoChatStarted"},
		GroupCallId: groupCallId,
	}

	return &messageVideoChatStartedTemp
}

// GetMessageContentEnum return the enum type of this object
func (messageVideoChatStarted *MessageVideoChatStarted) GetMessageContentEnum() MessageContentEnum {
	return MessageVideoChatStartedType
}

// MessageVideoChatEnded A message with information about an ended video chat
type MessageVideoChatEnded struct {
	tdCommon
	Duration int32 `json:"duration"` // Call duration, in seconds
}

// MessageType return the string telegram-type of MessageVideoChatEnded
func (messageVideoChatEnded *MessageVideoChatEnded) MessageType() string {
	return "messageVideoChatEnded"
}

// NewMessageVideoChatEnded creates a new MessageVideoChatEnded
//
// @param duration Call duration, in seconds
func NewMessageVideoChatEnded(duration int32) *MessageVideoChatEnded {
	messageVideoChatEndedTemp := MessageVideoChatEnded{
		tdCommon: tdCommon{Type: "messageVideoChatEnded"},
		Duration: duration,
	}

	return &messageVideoChatEndedTemp
}

// GetMessageContentEnum return the enum type of this object
func (messageVideoChatEnded *MessageVideoChatEnded) GetMessageContentEnum() MessageContentEnum {
	return MessageVideoChatEndedType
}

// MessageInviteVideoChatParticipants A message with information about an invite to a video chat
type MessageInviteVideoChatParticipants struct {
	tdCommon
	GroupCallId int32   `json:"group_call_id"` // Identifier of the video chat. The video chat can be received through the method getGroupCall
	UserIds     []int64 `json:"user_ids"`      // Invited user identifiers
}

// MessageType return the string telegram-type of MessageInviteVideoChatParticipants
func (messageInviteVideoChatParticipants *MessageInviteVideoChatParticipants) MessageType() string {
	return "messageInviteVideoChatParticipants"
}

// NewMessageInviteVideoChatParticipants creates a new MessageInviteVideoChatParticipants
//
// @param groupCallId Identifier of the video chat. The video chat can be received through the method getGroupCall
// @param userIds Invited user identifiers
func NewMessageInviteVideoChatParticipants(groupCallId int32, userIds []int64) *MessageInviteVideoChatParticipants {
	messageInviteVideoChatParticipantsTemp := MessageInviteVideoChatParticipants{
		tdCommon:    tdCommon{Type: "messageInviteVideoChatParticipants"},
		GroupCallId: groupCallId,
		UserIds:     userIds,
	}

	return &messageInviteVideoChatParticipantsTemp
}

// GetMessageContentEnum return the enum type of this object
func (messageInviteVideoChatParticipants *MessageInviteVideoChatParticipants) GetMessageContentEnum() MessageContentEnum {
	return MessageInviteVideoChatParticipantsType
}

// MessageBasicGroupChatCreate A newly created basic group
type MessageBasicGroupChatCreate struct {
	tdCommon
	Title         string  `json:"title"`           // Title of the basic group
	MemberUserIds []int64 `json:"member_user_ids"` // User identifiers of members in the basic group
}

// MessageType return the string telegram-type of MessageBasicGroupChatCreate
func (messageBasicGroupChatCreate *MessageBasicGroupChatCreate) MessageType() string {
	return "messageBasicGroupChatCreate"
}

// NewMessageBasicGroupChatCreate creates a new MessageBasicGroupChatCreate
//
// @param title Title of the basic group
// @param memberUserIds User identifiers of members in the basic group
func NewMessageBasicGroupChatCreate(title string, memberUserIds []int64) *MessageBasicGroupChatCreate {
	messageBasicGroupChatCreateTemp := MessageBasicGroupChatCreate{
		tdCommon:      tdCommon{Type: "messageBasicGroupChatCreate"},
		Title:         title,
		MemberUserIds: memberUserIds,
	}

	return &messageBasicGroupChatCreateTemp
}

// GetMessageContentEnum return the enum type of this object
func (messageBasicGroupChatCreate *MessageBasicGroupChatCreate) GetMessageContentEnum() MessageContentEnum {
	return MessageBasicGroupChatCreateType
}

// MessageSupergroupChatCreate A newly created supergroup or channel
type MessageSupergroupChatCreate struct {
	tdCommon
	Title string `json:"title"` // Title of the supergroup or channel
}

// MessageType return the string telegram-type of MessageSupergroupChatCreate
func (messageSupergroupChatCreate *MessageSupergroupChatCreate) MessageType() string {
	return "messageSupergroupChatCreate"
}

// NewMessageSupergroupChatCreate creates a new MessageSupergroupChatCreate
//
// @param title Title of the supergroup or channel
func NewMessageSupergroupChatCreate(title string) *MessageSupergroupChatCreate {
	messageSupergroupChatCreateTemp := MessageSupergroupChatCreate{
		tdCommon: tdCommon{Type: "messageSupergroupChatCreate"},
		Title:    title,
	}

	return &messageSupergroupChatCreateTemp
}

// GetMessageContentEnum return the enum type of this object
func (messageSupergroupChatCreate *MessageSupergroupChatCreate) GetMessageContentEnum() MessageContentEnum {
	return MessageSupergroupChatCreateType
}

// MessageChatChangeTitle An updated chat title
type MessageChatChangeTitle struct {
	tdCommon
	Title string `json:"title"` // New chat title
}

// MessageType return the string telegram-type of MessageChatChangeTitle
func (messageChatChangeTitle *MessageChatChangeTitle) MessageType() string {
	return "messageChatChangeTitle"
}

// NewMessageChatChangeTitle creates a new MessageChatChangeTitle
//
// @param title New chat title
func NewMessageChatChangeTitle(title string) *MessageChatChangeTitle {
	messageChatChangeTitleTemp := MessageChatChangeTitle{
		tdCommon: tdCommon{Type: "messageChatChangeTitle"},
		Title:    title,
	}

	return &messageChatChangeTitleTemp
}

// GetMessageContentEnum return the enum type of this object
func (messageChatChangeTitle *MessageChatChangeTitle) GetMessageContentEnum() MessageContentEnum {
	return MessageChatChangeTitleType
}

// MessageChatChangePhoto An updated chat photo
type MessageChatChangePhoto struct {
	tdCommon
	Photo *ChatPhoto `json:"photo"` // New chat photo
}

// MessageType return the string telegram-type of MessageChatChangePhoto
func (messageChatChangePhoto *MessageChatChangePhoto) MessageType() string {
	return "messageChatChangePhoto"
}

// NewMessageChatChangePhoto creates a new MessageChatChangePhoto
//
// @param photo New chat photo
func NewMessageChatChangePhoto(photo *ChatPhoto) *MessageChatChangePhoto {
	messageChatChangePhotoTemp := MessageChatChangePhoto{
		tdCommon: tdCommon{Type: "messageChatChangePhoto"},
		Photo:    photo,
	}

	return &messageChatChangePhotoTemp
}

// GetMessageContentEnum return the enum type of this object
func (messageChatChangePhoto *MessageChatChangePhoto) GetMessageContentEnum() MessageContentEnum {
	return MessageChatChangePhotoType
}

// MessageChatDeletePhoto A deleted chat photo
type MessageChatDeletePhoto struct {
	tdCommon
}

// MessageType return the string telegram-type of MessageChatDeletePhoto
func (messageChatDeletePhoto *MessageChatDeletePhoto) MessageType() string {
	return "messageChatDeletePhoto"
}

// NewMessageChatDeletePhoto creates a new MessageChatDeletePhoto
//
func NewMessageChatDeletePhoto() *MessageChatDeletePhoto {
	messageChatDeletePhotoTemp := MessageChatDeletePhoto{
		tdCommon: tdCommon{Type: "messageChatDeletePhoto"},
	}

	return &messageChatDeletePhotoTemp
}

// GetMessageContentEnum return the enum type of this object
func (messageChatDeletePhoto *MessageChatDeletePhoto) GetMessageContentEnum() MessageContentEnum {
	return MessageChatDeletePhotoType
}

// MessageChatAddMembers New chat members were added
type MessageChatAddMembers struct {
	tdCommon
	MemberUserIds []int64 `json:"member_user_ids"` // User identifiers of the new members
}

// MessageType return the string telegram-type of MessageChatAddMembers
func (messageChatAddMembers *MessageChatAddMembers) MessageType() string {
	return "messageChatAddMembers"
}

// NewMessageChatAddMembers creates a new MessageChatAddMembers
//
// @param memberUserIds User identifiers of the new members
func NewMessageChatAddMembers(memberUserIds []int64) *MessageChatAddMembers {
	messageChatAddMembersTemp := MessageChatAddMembers{
		tdCommon:      tdCommon{Type: "messageChatAddMembers"},
		MemberUserIds: memberUserIds,
	}

	return &messageChatAddMembersTemp
}

// GetMessageContentEnum return the enum type of this object
func (messageChatAddMembers *MessageChatAddMembers) GetMessageContentEnum() MessageContentEnum {
	return MessageChatAddMembersType
}

// MessageChatJoinByLink A new member joined the chat via an invite link
type MessageChatJoinByLink struct {
	tdCommon
}

// MessageType return the string telegram-type of MessageChatJoinByLink
func (messageChatJoinByLink *MessageChatJoinByLink) MessageType() string {
	return "messageChatJoinByLink"
}

// NewMessageChatJoinByLink creates a new MessageChatJoinByLink
//
func NewMessageChatJoinByLink() *MessageChatJoinByLink {
	messageChatJoinByLinkTemp := MessageChatJoinByLink{
		tdCommon: tdCommon{Type: "messageChatJoinByLink"},
	}

	return &messageChatJoinByLinkTemp
}

// GetMessageContentEnum return the enum type of this object
func (messageChatJoinByLink *MessageChatJoinByLink) GetMessageContentEnum() MessageContentEnum {
	return MessageChatJoinByLinkType
}

// MessageChatJoinByRequest A new member was accepted to the chat by an administrator
type MessageChatJoinByRequest struct {
	tdCommon
}

// MessageType return the string telegram-type of MessageChatJoinByRequest
func (messageChatJoinByRequest *MessageChatJoinByRequest) MessageType() string {
	return "messageChatJoinByRequest"
}

// NewMessageChatJoinByRequest creates a new MessageChatJoinByRequest
//
func NewMessageChatJoinByRequest() *MessageChatJoinByRequest {
	messageChatJoinByRequestTemp := MessageChatJoinByRequest{
		tdCommon: tdCommon{Type: "messageChatJoinByRequest"},
	}

	return &messageChatJoinByRequestTemp
}

// GetMessageContentEnum return the enum type of this object
func (messageChatJoinByRequest *MessageChatJoinByRequest) GetMessageContentEnum() MessageContentEnum {
	return MessageChatJoinByRequestType
}

// MessageChatDeleteMember A chat member was deleted
type MessageChatDeleteMember struct {
	tdCommon
	UserId int64 `json:"user_id"` // User identifier of the deleted chat member
}

// MessageType return the string telegram-type of MessageChatDeleteMember
func (messageChatDeleteMember *MessageChatDeleteMember) MessageType() string {
	return "messageChatDeleteMember"
}

// NewMessageChatDeleteMember creates a new MessageChatDeleteMember
//
// @param userId User identifier of the deleted chat member
func NewMessageChatDeleteMember(userId int64) *MessageChatDeleteMember {
	messageChatDeleteMemberTemp := MessageChatDeleteMember{
		tdCommon: tdCommon{Type: "messageChatDeleteMember"},
		UserId:   userId,
	}

	return &messageChatDeleteMemberTemp
}

// GetMessageContentEnum return the enum type of this object
func (messageChatDeleteMember *MessageChatDeleteMember) GetMessageContentEnum() MessageContentEnum {
	return MessageChatDeleteMemberType
}

// MessageChatUpgradeTo A basic group was upgraded to a supergroup and was deactivated as the result
type MessageChatUpgradeTo struct {
	tdCommon
	SupergroupId int64 `json:"supergroup_id"` // Identifier of the supergroup to which the basic group was upgraded
}

// MessageType return the string telegram-type of MessageChatUpgradeTo
func (messageChatUpgradeTo *MessageChatUpgradeTo) MessageType() string {
	return "messageChatUpgradeTo"
}

// NewMessageChatUpgradeTo creates a new MessageChatUpgradeTo
//
// @param supergroupId Identifier of the supergroup to which the basic group was upgraded
func NewMessageChatUpgradeTo(supergroupId int64) *MessageChatUpgradeTo {
	messageChatUpgradeToTemp := MessageChatUpgradeTo{
		tdCommon:     tdCommon{Type: "messageChatUpgradeTo"},
		SupergroupId: supergroupId,
	}

	return &messageChatUpgradeToTemp
}

// GetMessageContentEnum return the enum type of this object
func (messageChatUpgradeTo *MessageChatUpgradeTo) GetMessageContentEnum() MessageContentEnum {
	return MessageChatUpgradeToType
}

// MessageChatUpgradeFrom A supergroup has been created from a basic group
type MessageChatUpgradeFrom struct {
	tdCommon
	Title        string `json:"title"`          // Title of the newly created supergroup
	BasicGroupId int64  `json:"basic_group_id"` // The identifier of the original basic group
}

// MessageType return the string telegram-type of MessageChatUpgradeFrom
func (messageChatUpgradeFrom *MessageChatUpgradeFrom) MessageType() string {
	return "messageChatUpgradeFrom"
}

// NewMessageChatUpgradeFrom creates a new MessageChatUpgradeFrom
//
// @param title Title of the newly created supergroup
// @param basicGroupId The identifier of the original basic group
func NewMessageChatUpgradeFrom(title string, basicGroupId int64) *MessageChatUpgradeFrom {
	messageChatUpgradeFromTemp := MessageChatUpgradeFrom{
		tdCommon:     tdCommon{Type: "messageChatUpgradeFrom"},
		Title:        title,
		BasicGroupId: basicGroupId,
	}

	return &messageChatUpgradeFromTemp
}

// GetMessageContentEnum return the enum type of this object
func (messageChatUpgradeFrom *MessageChatUpgradeFrom) GetMessageContentEnum() MessageContentEnum {
	return MessageChatUpgradeFromType
}

// MessagePinMessage A message has been pinned
type MessagePinMessage struct {
	tdCommon
	MessageId int64 `json:"message_id"` // Identifier of the pinned message, can be an identifier of a deleted message or 0
}

// MessageType return the string telegram-type of MessagePinMessage
func (messagePinMessage *MessagePinMessage) MessageType() string {
	return "messagePinMessage"
}

// NewMessagePinMessage creates a new MessagePinMessage
//
// @param messageId Identifier of the pinned message, can be an identifier of a deleted message or 0
func NewMessagePinMessage(messageId int64) *MessagePinMessage {
	messagePinMessageTemp := MessagePinMessage{
		tdCommon:  tdCommon{Type: "messagePinMessage"},
		MessageId: messageId,
	}

	return &messagePinMessageTemp
}

// GetMessageContentEnum return the enum type of this object
func (messagePinMessage *MessagePinMessage) GetMessageContentEnum() MessageContentEnum {
	return MessagePinMessageType
}

// MessageScreenshotTaken A screenshot of a message in the chat has been taken
type MessageScreenshotTaken struct {
	tdCommon
}

// MessageType return the string telegram-type of MessageScreenshotTaken
func (messageScreenshotTaken *MessageScreenshotTaken) MessageType() string {
	return "messageScreenshotTaken"
}

// NewMessageScreenshotTaken creates a new MessageScreenshotTaken
//
func NewMessageScreenshotTaken() *MessageScreenshotTaken {
	messageScreenshotTakenTemp := MessageScreenshotTaken{
		tdCommon: tdCommon{Type: "messageScreenshotTaken"},
	}

	return &messageScreenshotTakenTemp
}

// GetMessageContentEnum return the enum type of this object
func (messageScreenshotTaken *MessageScreenshotTaken) GetMessageContentEnum() MessageContentEnum {
	return MessageScreenshotTakenType
}

// MessageChatSetTheme A theme in the chat has been changed
type MessageChatSetTheme struct {
	tdCommon
	ThemeName string `json:"theme_name"` // If non-empty, name of a new theme, set for the chat. Otherwise chat theme was reset to the default one
}

// MessageType return the string telegram-type of MessageChatSetTheme
func (messageChatSetTheme *MessageChatSetTheme) MessageType() string {
	return "messageChatSetTheme"
}

// NewMessageChatSetTheme creates a new MessageChatSetTheme
//
// @param themeName If non-empty, name of a new theme, set for the chat. Otherwise chat theme was reset to the default one
func NewMessageChatSetTheme(themeName string) *MessageChatSetTheme {
	messageChatSetThemeTemp := MessageChatSetTheme{
		tdCommon:  tdCommon{Type: "messageChatSetTheme"},
		ThemeName: themeName,
	}

	return &messageChatSetThemeTemp
}

// GetMessageContentEnum return the enum type of this object
func (messageChatSetTheme *MessageChatSetTheme) GetMessageContentEnum() MessageContentEnum {
	return MessageChatSetThemeType
}

// MessageChatSetTtl The TTL (Time To Live) setting for messages in the chat has been changed
type MessageChatSetTtl struct {
	tdCommon
	Ttl int32 `json:"ttl"` // New message TTL
}

// MessageType return the string telegram-type of MessageChatSetTtl
func (messageChatSetTtl *MessageChatSetTtl) MessageType() string {
	return "messageChatSetTtl"
}

// NewMessageChatSetTtl creates a new MessageChatSetTtl
//
// @param ttl New message TTL
func NewMessageChatSetTtl(ttl int32) *MessageChatSetTtl {
	messageChatSetTtlTemp := MessageChatSetTtl{
		tdCommon: tdCommon{Type: "messageChatSetTtl"},
		Ttl:      ttl,
	}

	return &messageChatSetTtlTemp
}

// GetMessageContentEnum return the enum type of this object
func (messageChatSetTtl *MessageChatSetTtl) GetMessageContentEnum() MessageContentEnum {
	return MessageChatSetTtlType
}

// MessageCustomServiceAction A non-standard action has happened in the chat
type MessageCustomServiceAction struct {
	tdCommon
	Text string `json:"text"` // Message text to be shown in the chat
}

// MessageType return the string telegram-type of MessageCustomServiceAction
func (messageCustomServiceAction *MessageCustomServiceAction) MessageType() string {
	return "messageCustomServiceAction"
}

// NewMessageCustomServiceAction creates a new MessageCustomServiceAction
//
// @param text Message text to be shown in the chat
func NewMessageCustomServiceAction(text string) *MessageCustomServiceAction {
	messageCustomServiceActionTemp := MessageCustomServiceAction{
		tdCommon: tdCommon{Type: "messageCustomServiceAction"},
		Text:     text,
	}

	return &messageCustomServiceActionTemp
}

// GetMessageContentEnum return the enum type of this object
func (messageCustomServiceAction *MessageCustomServiceAction) GetMessageContentEnum() MessageContentEnum {
	return MessageCustomServiceActionType
}

// MessageGameScore A new high score was achieved in a game
type MessageGameScore struct {
	tdCommon
	GameMessageId int64     `json:"game_message_id"` // Identifier of the message with the game, can be an identifier of a deleted message
	GameId        JSONInt64 `json:"game_id"`         // Identifier of the game; may be different from the games presented in the message with the game
	Score         int32     `json:"score"`           // New score
}

// MessageType return the string telegram-type of MessageGameScore
func (messageGameScore *MessageGameScore) MessageType() string {
	return "messageGameScore"
}

// NewMessageGameScore creates a new MessageGameScore
//
// @param gameMessageId Identifier of the message with the game, can be an identifier of a deleted message
// @param gameId Identifier of the game; may be different from the games presented in the message with the game
// @param score New score
func NewMessageGameScore(gameMessageId int64, gameId JSONInt64, score int32) *MessageGameScore {
	messageGameScoreTemp := MessageGameScore{
		tdCommon:      tdCommon{Type: "messageGameScore"},
		GameMessageId: gameMessageId,
		GameId:        gameId,
		Score:         score,
	}

	return &messageGameScoreTemp
}

// GetMessageContentEnum return the enum type of this object
func (messageGameScore *MessageGameScore) GetMessageContentEnum() MessageContentEnum {
	return MessageGameScoreType
}

// MessagePaymentSuccessful A payment has been completed
type MessagePaymentSuccessful struct {
	tdCommon
	InvoiceChatId    int64  `json:"invoice_chat_id"`    // Identifier of the chat, containing the corresponding invoice message; 0 if unknown
	InvoiceMessageId int64  `json:"invoice_message_id"` // Identifier of the message with the corresponding invoice; can be an identifier of a deleted message
	Currency         string `json:"currency"`           // Currency for the price of the product
	TotalAmount      int64  `json:"total_amount"`       // Total price for the product, in the smallest units of the currency
}

// MessageType return the string telegram-type of MessagePaymentSuccessful
func (messagePaymentSuccessful *MessagePaymentSuccessful) MessageType() string {
	return "messagePaymentSuccessful"
}

// NewMessagePaymentSuccessful creates a new MessagePaymentSuccessful
//
// @param invoiceChatId Identifier of the chat, containing the corresponding invoice message; 0 if unknown
// @param invoiceMessageId Identifier of the message with the corresponding invoice; can be an identifier of a deleted message
// @param currency Currency for the price of the product
// @param totalAmount Total price for the product, in the smallest units of the currency
func NewMessagePaymentSuccessful(invoiceChatId int64, invoiceMessageId int64, currency string, totalAmount int64) *MessagePaymentSuccessful {
	messagePaymentSuccessfulTemp := MessagePaymentSuccessful{
		tdCommon:         tdCommon{Type: "messagePaymentSuccessful"},
		InvoiceChatId:    invoiceChatId,
		InvoiceMessageId: invoiceMessageId,
		Currency:         currency,
		TotalAmount:      totalAmount,
	}

	return &messagePaymentSuccessfulTemp
}

// GetMessageContentEnum return the enum type of this object
func (messagePaymentSuccessful *MessagePaymentSuccessful) GetMessageContentEnum() MessageContentEnum {
	return MessagePaymentSuccessfulType
}

// MessagePaymentSuccessfulBot A payment has been completed; for bots only
type MessagePaymentSuccessfulBot struct {
	tdCommon
	Currency                string     `json:"currency"`                   // Currency for price of the product
	TotalAmount             int64      `json:"total_amount"`               // Total price for the product, in the smallest units of the currency
	InvoicePayload          []byte     `json:"invoice_payload"`            // Invoice payload
	ShippingOptionId        string     `json:"shipping_option_id"`         // Identifier of the shipping option chosen by the user; may be empty if not applicable
	OrderInfo               *OrderInfo `json:"order_info"`                 // Information about the order; may be null
	TelegramPaymentChargeId string     `json:"telegram_payment_charge_id"` // Telegram payment identifier
	ProviderPaymentChargeId string     `json:"provider_payment_charge_id"` // Provider payment identifier
}

// MessageType return the string telegram-type of MessagePaymentSuccessfulBot
func (messagePaymentSuccessfulBot *MessagePaymentSuccessfulBot) MessageType() string {
	return "messagePaymentSuccessfulBot"
}

// NewMessagePaymentSuccessfulBot creates a new MessagePaymentSuccessfulBot
//
// @param currency Currency for price of the product
// @param totalAmount Total price for the product, in the smallest units of the currency
// @param invoicePayload Invoice payload
// @param shippingOptionId Identifier of the shipping option chosen by the user; may be empty if not applicable
// @param orderInfo Information about the order; may be null
// @param telegramPaymentChargeId Telegram payment identifier
// @param providerPaymentChargeId Provider payment identifier
func NewMessagePaymentSuccessfulBot(currency string, totalAmount int64, invoicePayload []byte, shippingOptionId string, orderInfo *OrderInfo, telegramPaymentChargeId string, providerPaymentChargeId string) *MessagePaymentSuccessfulBot {
	messagePaymentSuccessfulBotTemp := MessagePaymentSuccessfulBot{
		tdCommon:                tdCommon{Type: "messagePaymentSuccessfulBot"},
		Currency:                currency,
		TotalAmount:             totalAmount,
		InvoicePayload:          invoicePayload,
		ShippingOptionId:        shippingOptionId,
		OrderInfo:               orderInfo,
		TelegramPaymentChargeId: telegramPaymentChargeId,
		ProviderPaymentChargeId: providerPaymentChargeId,
	}

	return &messagePaymentSuccessfulBotTemp
}

// GetMessageContentEnum return the enum type of this object
func (messagePaymentSuccessfulBot *MessagePaymentSuccessfulBot) GetMessageContentEnum() MessageContentEnum {
	return MessagePaymentSuccessfulBotType
}

// MessageContactRegistered A contact has registered with Telegram
type MessageContactRegistered struct {
	tdCommon
}

// MessageType return the string telegram-type of MessageContactRegistered
func (messageContactRegistered *MessageContactRegistered) MessageType() string {
	return "messageContactRegistered"
}

// NewMessageContactRegistered creates a new MessageContactRegistered
//
func NewMessageContactRegistered() *MessageContactRegistered {
	messageContactRegisteredTemp := MessageContactRegistered{
		tdCommon: tdCommon{Type: "messageContactRegistered"},
	}

	return &messageContactRegisteredTemp
}

// GetMessageContentEnum return the enum type of this object
func (messageContactRegistered *MessageContactRegistered) GetMessageContentEnum() MessageContentEnum {
	return MessageContactRegisteredType
}

// MessageWebsiteConnected The current user has connected a website by logging in using Telegram Login Widget on it
type MessageWebsiteConnected struct {
	tdCommon
	DomainName string `json:"domain_name"` // Domain name of the connected website
}

// MessageType return the string telegram-type of MessageWebsiteConnected
func (messageWebsiteConnected *MessageWebsiteConnected) MessageType() string {
	return "messageWebsiteConnected"
}

// NewMessageWebsiteConnected creates a new MessageWebsiteConnected
//
// @param domainName Domain name of the connected website
func NewMessageWebsiteConnected(domainName string) *MessageWebsiteConnected {
	messageWebsiteConnectedTemp := MessageWebsiteConnected{
		tdCommon:   tdCommon{Type: "messageWebsiteConnected"},
		DomainName: domainName,
	}

	return &messageWebsiteConnectedTemp
}

// GetMessageContentEnum return the enum type of this object
func (messageWebsiteConnected *MessageWebsiteConnected) GetMessageContentEnum() MessageContentEnum {
	return MessageWebsiteConnectedType
}

// MessagePassportDataSent Telegram Passport data has been sent
type MessagePassportDataSent struct {
	tdCommon
	Types []PassportElementType `json:"types"` // List of Telegram Passport element types sent
}

// MessageType return the string telegram-type of MessagePassportDataSent
func (messagePassportDataSent *MessagePassportDataSent) MessageType() string {
	return "messagePassportDataSent"
}

// NewMessagePassportDataSent creates a new MessagePassportDataSent
//
// @param typeParams List of Telegram Passport element types sent
func NewMessagePassportDataSent(typeParams []PassportElementType) *MessagePassportDataSent {
	messagePassportDataSentTemp := MessagePassportDataSent{
		tdCommon: tdCommon{Type: "messagePassportDataSent"},
		Types:    typeParams,
	}

	return &messagePassportDataSentTemp
}

// GetMessageContentEnum return the enum type of this object
func (messagePassportDataSent *MessagePassportDataSent) GetMessageContentEnum() MessageContentEnum {
	return MessagePassportDataSentType
}

// MessagePassportDataReceived Telegram Passport data has been received; for bots only
type MessagePassportDataReceived struct {
	tdCommon
	Elements    []EncryptedPassportElement `json:"elements"`    // List of received Telegram Passport elements
	Credentials *EncryptedCredentials      `json:"credentials"` // Encrypted data credentials
}

// MessageType return the string telegram-type of MessagePassportDataReceived
func (messagePassportDataReceived *MessagePassportDataReceived) MessageType() string {
	return "messagePassportDataReceived"
}

// NewMessagePassportDataReceived creates a new MessagePassportDataReceived
//
// @param elements List of received Telegram Passport elements
// @param credentials Encrypted data credentials
func NewMessagePassportDataReceived(elements []EncryptedPassportElement, credentials *EncryptedCredentials) *MessagePassportDataReceived {
	messagePassportDataReceivedTemp := MessagePassportDataReceived{
		tdCommon:    tdCommon{Type: "messagePassportDataReceived"},
		Elements:    elements,
		Credentials: credentials,
	}

	return &messagePassportDataReceivedTemp
}

// GetMessageContentEnum return the enum type of this object
func (messagePassportDataReceived *MessagePassportDataReceived) GetMessageContentEnum() MessageContentEnum {
	return MessagePassportDataReceivedType
}

// MessageProximityAlertTriggered A user in the chat came within proximity alert range
type MessageProximityAlertTriggered struct {
	tdCommon
	TravelerId MessageSender `json:"traveler_id"` // The identifier of a user or chat that triggered the proximity alert
	WatcherId  MessageSender `json:"watcher_id"`  // The identifier of a user or chat that subscribed for the proximity alert
	Distance   int32         `json:"distance"`    // The distance between the users
}

// MessageType return the string telegram-type of MessageProximityAlertTriggered
func (messageProximityAlertTriggered *MessageProximityAlertTriggered) MessageType() string {
	return "messageProximityAlertTriggered"
}

// NewMessageProximityAlertTriggered creates a new MessageProximityAlertTriggered
//
// @param travelerId The identifier of a user or chat that triggered the proximity alert
// @param watcherId The identifier of a user or chat that subscribed for the proximity alert
// @param distance The distance between the users
func NewMessageProximityAlertTriggered(travelerId MessageSender, watcherId MessageSender, distance int32) *MessageProximityAlertTriggered {
	messageProximityAlertTriggeredTemp := MessageProximityAlertTriggered{
		tdCommon:   tdCommon{Type: "messageProximityAlertTriggered"},
		TravelerId: travelerId,
		WatcherId:  watcherId,
		Distance:   distance,
	}

	return &messageProximityAlertTriggeredTemp
}

// UnmarshalJSON unmarshal to json
func (messageProximityAlertTriggered *MessageProximityAlertTriggered) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	tempObj := struct {
		tdCommon
		Distance int32 `json:"distance"` // The distance between the users
	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	messageProximityAlertTriggered.tdCommon = tempObj.tdCommon
	messageProximityAlertTriggered.Distance = tempObj.Distance

	fieldTravelerId, _ := unmarshalMessageSender(objMap["traveler_id"])
	messageProximityAlertTriggered.TravelerId = fieldTravelerId

	fieldWatcherId, _ := unmarshalMessageSender(objMap["watcher_id"])
	messageProximityAlertTriggered.WatcherId = fieldWatcherId

	return nil
}

// GetMessageContentEnum return the enum type of this object
func (messageProximityAlertTriggered *MessageProximityAlertTriggered) GetMessageContentEnum() MessageContentEnum {
	return MessageProximityAlertTriggeredType
}

// MessageUnsupported Message content that is not supported in the current TDLib version
type MessageUnsupported struct {
	tdCommon
}

// MessageType return the string telegram-type of MessageUnsupported
func (messageUnsupported *MessageUnsupported) MessageType() string {
	return "messageUnsupported"
}

// NewMessageUnsupported creates a new MessageUnsupported
//
func NewMessageUnsupported() *MessageUnsupported {
	messageUnsupportedTemp := MessageUnsupported{
		tdCommon: tdCommon{Type: "messageUnsupported"},
	}

	return &messageUnsupportedTemp
}

// GetMessageContentEnum return the enum type of this object
func (messageUnsupported *MessageUnsupported) GetMessageContentEnum() MessageContentEnum {
	return MessageUnsupportedType
}
