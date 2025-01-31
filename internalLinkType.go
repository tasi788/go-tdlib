package tdlib

import (
	"encoding/json"
	"fmt"
)

// InternalLinkType Describes an internal https://t.me or tg: link, which must be processed by the app in a special way
type InternalLinkType interface {
	GetInternalLinkTypeEnum() InternalLinkTypeEnum
}

// InternalLinkTypeEnum Alias for abstract InternalLinkType 'Sub-Classes', used as constant-enum here
type InternalLinkTypeEnum string

// InternalLinkType enums
const (
	InternalLinkTypeActiveSessionsType          InternalLinkTypeEnum = "internalLinkTypeActiveSessions"
	InternalLinkTypeAuthenticationCodeType      InternalLinkTypeEnum = "internalLinkTypeAuthenticationCode"
	InternalLinkTypeBackgroundType              InternalLinkTypeEnum = "internalLinkTypeBackground"
	InternalLinkTypeBotStartType                InternalLinkTypeEnum = "internalLinkTypeBotStart"
	InternalLinkTypeBotStartInGroupType         InternalLinkTypeEnum = "internalLinkTypeBotStartInGroup"
	InternalLinkTypeChangePhoneNumberType       InternalLinkTypeEnum = "internalLinkTypeChangePhoneNumber"
	InternalLinkTypeChatInviteType              InternalLinkTypeEnum = "internalLinkTypeChatInvite"
	InternalLinkTypeFilterSettingsType          InternalLinkTypeEnum = "internalLinkTypeFilterSettings"
	InternalLinkTypeGameType                    InternalLinkTypeEnum = "internalLinkTypeGame"
	InternalLinkTypeLanguagePackType            InternalLinkTypeEnum = "internalLinkTypeLanguagePack"
	InternalLinkTypeMessageType                 InternalLinkTypeEnum = "internalLinkTypeMessage"
	InternalLinkTypeMessageDraftType            InternalLinkTypeEnum = "internalLinkTypeMessageDraft"
	InternalLinkTypePassportDataRequestType     InternalLinkTypeEnum = "internalLinkTypePassportDataRequest"
	InternalLinkTypePhoneNumberConfirmationType InternalLinkTypeEnum = "internalLinkTypePhoneNumberConfirmation"
	InternalLinkTypeProxyType                   InternalLinkTypeEnum = "internalLinkTypeProxy"
	InternalLinkTypePublicChatType              InternalLinkTypeEnum = "internalLinkTypePublicChat"
	InternalLinkTypeQrCodeAuthenticationType    InternalLinkTypeEnum = "internalLinkTypeQrCodeAuthentication"
	InternalLinkTypeSettingsType                InternalLinkTypeEnum = "internalLinkTypeSettings"
	InternalLinkTypeStickerSetType              InternalLinkTypeEnum = "internalLinkTypeStickerSet"
	InternalLinkTypeThemeType                   InternalLinkTypeEnum = "internalLinkTypeTheme"
	InternalLinkTypeThemeSettingsType           InternalLinkTypeEnum = "internalLinkTypeThemeSettings"
	InternalLinkTypeUnknownDeepLinkType         InternalLinkTypeEnum = "internalLinkTypeUnknownDeepLink"
	InternalLinkTypeUnsupportedProxyType        InternalLinkTypeEnum = "internalLinkTypeUnsupportedProxy"
	InternalLinkTypeVideoChatType               InternalLinkTypeEnum = "internalLinkTypeVideoChat"
)

func unmarshalInternalLinkType(rawMsg *json.RawMessage) (InternalLinkType, error) {

	if rawMsg == nil {
		return nil, nil
	}
	var objMap map[string]interface{}
	err := json.Unmarshal(*rawMsg, &objMap)
	if err != nil {
		return nil, err
	}

	switch InternalLinkTypeEnum(objMap["@type"].(string)) {
	case InternalLinkTypeActiveSessionsType:
		var internalLinkTypeActiveSessions InternalLinkTypeActiveSessions
		err := json.Unmarshal(*rawMsg, &internalLinkTypeActiveSessions)
		return &internalLinkTypeActiveSessions, err

	case InternalLinkTypeAuthenticationCodeType:
		var internalLinkTypeAuthenticationCode InternalLinkTypeAuthenticationCode
		err := json.Unmarshal(*rawMsg, &internalLinkTypeAuthenticationCode)
		return &internalLinkTypeAuthenticationCode, err

	case InternalLinkTypeBackgroundType:
		var internalLinkTypeBackground InternalLinkTypeBackground
		err := json.Unmarshal(*rawMsg, &internalLinkTypeBackground)
		return &internalLinkTypeBackground, err

	case InternalLinkTypeBotStartType:
		var internalLinkTypeBotStart InternalLinkTypeBotStart
		err := json.Unmarshal(*rawMsg, &internalLinkTypeBotStart)
		return &internalLinkTypeBotStart, err

	case InternalLinkTypeBotStartInGroupType:
		var internalLinkTypeBotStartInGroup InternalLinkTypeBotStartInGroup
		err := json.Unmarshal(*rawMsg, &internalLinkTypeBotStartInGroup)
		return &internalLinkTypeBotStartInGroup, err

	case InternalLinkTypeChangePhoneNumberType:
		var internalLinkTypeChangePhoneNumber InternalLinkTypeChangePhoneNumber
		err := json.Unmarshal(*rawMsg, &internalLinkTypeChangePhoneNumber)
		return &internalLinkTypeChangePhoneNumber, err

	case InternalLinkTypeChatInviteType:
		var internalLinkTypeChatInvite InternalLinkTypeChatInvite
		err := json.Unmarshal(*rawMsg, &internalLinkTypeChatInvite)
		return &internalLinkTypeChatInvite, err

	case InternalLinkTypeFilterSettingsType:
		var internalLinkTypeFilterSettings InternalLinkTypeFilterSettings
		err := json.Unmarshal(*rawMsg, &internalLinkTypeFilterSettings)
		return &internalLinkTypeFilterSettings, err

	case InternalLinkTypeGameType:
		var internalLinkTypeGame InternalLinkTypeGame
		err := json.Unmarshal(*rawMsg, &internalLinkTypeGame)
		return &internalLinkTypeGame, err

	case InternalLinkTypeLanguagePackType:
		var internalLinkTypeLanguagePack InternalLinkTypeLanguagePack
		err := json.Unmarshal(*rawMsg, &internalLinkTypeLanguagePack)
		return &internalLinkTypeLanguagePack, err

	case InternalLinkTypeMessageType:
		var internalLinkTypeMessage InternalLinkTypeMessage
		err := json.Unmarshal(*rawMsg, &internalLinkTypeMessage)
		return &internalLinkTypeMessage, err

	case InternalLinkTypeMessageDraftType:
		var internalLinkTypeMessageDraft InternalLinkTypeMessageDraft
		err := json.Unmarshal(*rawMsg, &internalLinkTypeMessageDraft)
		return &internalLinkTypeMessageDraft, err

	case InternalLinkTypePassportDataRequestType:
		var internalLinkTypePassportDataRequest InternalLinkTypePassportDataRequest
		err := json.Unmarshal(*rawMsg, &internalLinkTypePassportDataRequest)
		return &internalLinkTypePassportDataRequest, err

	case InternalLinkTypePhoneNumberConfirmationType:
		var internalLinkTypePhoneNumberConfirmation InternalLinkTypePhoneNumberConfirmation
		err := json.Unmarshal(*rawMsg, &internalLinkTypePhoneNumberConfirmation)
		return &internalLinkTypePhoneNumberConfirmation, err

	case InternalLinkTypeProxyType:
		var internalLinkTypeProxy InternalLinkTypeProxy
		err := json.Unmarshal(*rawMsg, &internalLinkTypeProxy)
		return &internalLinkTypeProxy, err

	case InternalLinkTypePublicChatType:
		var internalLinkTypePublicChat InternalLinkTypePublicChat
		err := json.Unmarshal(*rawMsg, &internalLinkTypePublicChat)
		return &internalLinkTypePublicChat, err

	case InternalLinkTypeQrCodeAuthenticationType:
		var internalLinkTypeQrCodeAuthentication InternalLinkTypeQrCodeAuthentication
		err := json.Unmarshal(*rawMsg, &internalLinkTypeQrCodeAuthentication)
		return &internalLinkTypeQrCodeAuthentication, err

	case InternalLinkTypeSettingsType:
		var internalLinkTypeSettings InternalLinkTypeSettings
		err := json.Unmarshal(*rawMsg, &internalLinkTypeSettings)
		return &internalLinkTypeSettings, err

	case InternalLinkTypeStickerSetType:
		var internalLinkTypeStickerSet InternalLinkTypeStickerSet
		err := json.Unmarshal(*rawMsg, &internalLinkTypeStickerSet)
		return &internalLinkTypeStickerSet, err

	case InternalLinkTypeThemeType:
		var internalLinkTypeTheme InternalLinkTypeTheme
		err := json.Unmarshal(*rawMsg, &internalLinkTypeTheme)
		return &internalLinkTypeTheme, err

	case InternalLinkTypeThemeSettingsType:
		var internalLinkTypeThemeSettings InternalLinkTypeThemeSettings
		err := json.Unmarshal(*rawMsg, &internalLinkTypeThemeSettings)
		return &internalLinkTypeThemeSettings, err

	case InternalLinkTypeUnknownDeepLinkType:
		var internalLinkTypeUnknownDeepLink InternalLinkTypeUnknownDeepLink
		err := json.Unmarshal(*rawMsg, &internalLinkTypeUnknownDeepLink)
		return &internalLinkTypeUnknownDeepLink, err

	case InternalLinkTypeUnsupportedProxyType:
		var internalLinkTypeUnsupportedProxy InternalLinkTypeUnsupportedProxy
		err := json.Unmarshal(*rawMsg, &internalLinkTypeUnsupportedProxy)
		return &internalLinkTypeUnsupportedProxy, err

	case InternalLinkTypeVideoChatType:
		var internalLinkTypeVideoChat InternalLinkTypeVideoChat
		err := json.Unmarshal(*rawMsg, &internalLinkTypeVideoChat)
		return &internalLinkTypeVideoChat, err

	default:
		return nil, fmt.Errorf("Error UnMarshaling, unknown type:" + objMap["@type"].(string))
	}
}

// InternalLinkTypeActiveSessions The link is a link to the active sessions section of the app. Use getActiveSessions to handle the link
type InternalLinkTypeActiveSessions struct {
	tdCommon
}

// MessageType return the string telegram-type of InternalLinkTypeActiveSessions
func (internalLinkTypeActiveSessions *InternalLinkTypeActiveSessions) MessageType() string {
	return "internalLinkTypeActiveSessions"
}

// NewInternalLinkTypeActiveSessions creates a new InternalLinkTypeActiveSessions
//
func NewInternalLinkTypeActiveSessions() *InternalLinkTypeActiveSessions {
	internalLinkTypeActiveSessionsTemp := InternalLinkTypeActiveSessions{
		tdCommon: tdCommon{Type: "internalLinkTypeActiveSessions"},
	}

	return &internalLinkTypeActiveSessionsTemp
}

// GetInternalLinkTypeEnum return the enum type of this object
func (internalLinkTypeActiveSessions *InternalLinkTypeActiveSessions) GetInternalLinkTypeEnum() InternalLinkTypeEnum {
	return InternalLinkTypeActiveSessionsType
}

// InternalLinkTypeAuthenticationCode The link contains an authentication code. Call checkAuthenticationCode with the code if the current authorization state is authorizationStateWaitCode
type InternalLinkTypeAuthenticationCode struct {
	tdCommon
	Code string `json:"code"` // The authentication code
}

// MessageType return the string telegram-type of InternalLinkTypeAuthenticationCode
func (internalLinkTypeAuthenticationCode *InternalLinkTypeAuthenticationCode) MessageType() string {
	return "internalLinkTypeAuthenticationCode"
}

// NewInternalLinkTypeAuthenticationCode creates a new InternalLinkTypeAuthenticationCode
//
// @param code The authentication code
func NewInternalLinkTypeAuthenticationCode(code string) *InternalLinkTypeAuthenticationCode {
	internalLinkTypeAuthenticationCodeTemp := InternalLinkTypeAuthenticationCode{
		tdCommon: tdCommon{Type: "internalLinkTypeAuthenticationCode"},
		Code:     code,
	}

	return &internalLinkTypeAuthenticationCodeTemp
}

// GetInternalLinkTypeEnum return the enum type of this object
func (internalLinkTypeAuthenticationCode *InternalLinkTypeAuthenticationCode) GetInternalLinkTypeEnum() InternalLinkTypeEnum {
	return InternalLinkTypeAuthenticationCodeType
}

// InternalLinkTypeBackground The link is a link to a background. Call searchBackground with the given background name to process the link
type InternalLinkTypeBackground struct {
	tdCommon
	BackgroundName string `json:"background_name"` // Name of the background
}

// MessageType return the string telegram-type of InternalLinkTypeBackground
func (internalLinkTypeBackground *InternalLinkTypeBackground) MessageType() string {
	return "internalLinkTypeBackground"
}

// NewInternalLinkTypeBackground creates a new InternalLinkTypeBackground
//
// @param backgroundName Name of the background
func NewInternalLinkTypeBackground(backgroundName string) *InternalLinkTypeBackground {
	internalLinkTypeBackgroundTemp := InternalLinkTypeBackground{
		tdCommon:       tdCommon{Type: "internalLinkTypeBackground"},
		BackgroundName: backgroundName,
	}

	return &internalLinkTypeBackgroundTemp
}

// GetInternalLinkTypeEnum return the enum type of this object
func (internalLinkTypeBackground *InternalLinkTypeBackground) GetInternalLinkTypeEnum() InternalLinkTypeEnum {
	return InternalLinkTypeBackgroundType
}

// InternalLinkTypeBotStart The link is a link to a chat with a Telegram bot. Call searchPublicChat with the given bot username, check that the user is a bot, show START button in the chat with the bot, and then call sendBotStartMessage with the given start parameter after the button is pressed
type InternalLinkTypeBotStart struct {
	tdCommon
	BotUsername    string `json:"bot_username"`    // Username of the bot
	StartParameter string `json:"start_parameter"` // The parameter to be passed to sendBotStartMessage
}

// MessageType return the string telegram-type of InternalLinkTypeBotStart
func (internalLinkTypeBotStart *InternalLinkTypeBotStart) MessageType() string {
	return "internalLinkTypeBotStart"
}

// NewInternalLinkTypeBotStart creates a new InternalLinkTypeBotStart
//
// @param botUsername Username of the bot
// @param startParameter The parameter to be passed to sendBotStartMessage
func NewInternalLinkTypeBotStart(botUsername string, startParameter string) *InternalLinkTypeBotStart {
	internalLinkTypeBotStartTemp := InternalLinkTypeBotStart{
		tdCommon:       tdCommon{Type: "internalLinkTypeBotStart"},
		BotUsername:    botUsername,
		StartParameter: startParameter,
	}

	return &internalLinkTypeBotStartTemp
}

// GetInternalLinkTypeEnum return the enum type of this object
func (internalLinkTypeBotStart *InternalLinkTypeBotStart) GetInternalLinkTypeEnum() InternalLinkTypeEnum {
	return InternalLinkTypeBotStartType
}

// InternalLinkTypeBotStartInGroup The link is a link to a Telegram bot, which is supposed to be added to a group chat. Call searchPublicChat with the given bot username, check that the user is a bot and can be added to groups, ask the current user to select a group to add the bot to, and then call sendBotStartMessage with the given start parameter and the chosen group chat. Bots can be added to a public group only by administrators of the group
type InternalLinkTypeBotStartInGroup struct {
	tdCommon
	BotUsername    string `json:"bot_username"`    // Username of the bot
	StartParameter string `json:"start_parameter"` // The parameter to be passed to sendBotStartMessage
}

// MessageType return the string telegram-type of InternalLinkTypeBotStartInGroup
func (internalLinkTypeBotStartInGroup *InternalLinkTypeBotStartInGroup) MessageType() string {
	return "internalLinkTypeBotStartInGroup"
}

// NewInternalLinkTypeBotStartInGroup creates a new InternalLinkTypeBotStartInGroup
//
// @param botUsername Username of the bot
// @param startParameter The parameter to be passed to sendBotStartMessage
func NewInternalLinkTypeBotStartInGroup(botUsername string, startParameter string) *InternalLinkTypeBotStartInGroup {
	internalLinkTypeBotStartInGroupTemp := InternalLinkTypeBotStartInGroup{
		tdCommon:       tdCommon{Type: "internalLinkTypeBotStartInGroup"},
		BotUsername:    botUsername,
		StartParameter: startParameter,
	}

	return &internalLinkTypeBotStartInGroupTemp
}

// GetInternalLinkTypeEnum return the enum type of this object
func (internalLinkTypeBotStartInGroup *InternalLinkTypeBotStartInGroup) GetInternalLinkTypeEnum() InternalLinkTypeEnum {
	return InternalLinkTypeBotStartInGroupType
}

// InternalLinkTypeChangePhoneNumber The link is a link to the change phone number section of the app
type InternalLinkTypeChangePhoneNumber struct {
	tdCommon
}

// MessageType return the string telegram-type of InternalLinkTypeChangePhoneNumber
func (internalLinkTypeChangePhoneNumber *InternalLinkTypeChangePhoneNumber) MessageType() string {
	return "internalLinkTypeChangePhoneNumber"
}

// NewInternalLinkTypeChangePhoneNumber creates a new InternalLinkTypeChangePhoneNumber
//
func NewInternalLinkTypeChangePhoneNumber() *InternalLinkTypeChangePhoneNumber {
	internalLinkTypeChangePhoneNumberTemp := InternalLinkTypeChangePhoneNumber{
		tdCommon: tdCommon{Type: "internalLinkTypeChangePhoneNumber"},
	}

	return &internalLinkTypeChangePhoneNumberTemp
}

// GetInternalLinkTypeEnum return the enum type of this object
func (internalLinkTypeChangePhoneNumber *InternalLinkTypeChangePhoneNumber) GetInternalLinkTypeEnum() InternalLinkTypeEnum {
	return InternalLinkTypeChangePhoneNumberType
}

// InternalLinkTypeChatInvite The link is a chat invite link. Call checkChatInviteLink with the given invite link to process the link
type InternalLinkTypeChatInvite struct {
	tdCommon
	InviteLink string `json:"invite_link"` // Internal representation of the invite link
}

// MessageType return the string telegram-type of InternalLinkTypeChatInvite
func (internalLinkTypeChatInvite *InternalLinkTypeChatInvite) MessageType() string {
	return "internalLinkTypeChatInvite"
}

// NewInternalLinkTypeChatInvite creates a new InternalLinkTypeChatInvite
//
// @param inviteLink Internal representation of the invite link
func NewInternalLinkTypeChatInvite(inviteLink string) *InternalLinkTypeChatInvite {
	internalLinkTypeChatInviteTemp := InternalLinkTypeChatInvite{
		tdCommon:   tdCommon{Type: "internalLinkTypeChatInvite"},
		InviteLink: inviteLink,
	}

	return &internalLinkTypeChatInviteTemp
}

// GetInternalLinkTypeEnum return the enum type of this object
func (internalLinkTypeChatInvite *InternalLinkTypeChatInvite) GetInternalLinkTypeEnum() InternalLinkTypeEnum {
	return InternalLinkTypeChatInviteType
}

// InternalLinkTypeFilterSettings The link is a link to the filter settings section of the app
type InternalLinkTypeFilterSettings struct {
	tdCommon
}

// MessageType return the string telegram-type of InternalLinkTypeFilterSettings
func (internalLinkTypeFilterSettings *InternalLinkTypeFilterSettings) MessageType() string {
	return "internalLinkTypeFilterSettings"
}

// NewInternalLinkTypeFilterSettings creates a new InternalLinkTypeFilterSettings
//
func NewInternalLinkTypeFilterSettings() *InternalLinkTypeFilterSettings {
	internalLinkTypeFilterSettingsTemp := InternalLinkTypeFilterSettings{
		tdCommon: tdCommon{Type: "internalLinkTypeFilterSettings"},
	}

	return &internalLinkTypeFilterSettingsTemp
}

// GetInternalLinkTypeEnum return the enum type of this object
func (internalLinkTypeFilterSettings *InternalLinkTypeFilterSettings) GetInternalLinkTypeEnum() InternalLinkTypeEnum {
	return InternalLinkTypeFilterSettingsType
}

// InternalLinkTypeGame The link is a link to a game. Call searchPublicChat with the given bot username, check that the user is a bot, ask the current user to select a chat to send the game, and then call sendMessage with inputMessageGame
type InternalLinkTypeGame struct {
	tdCommon
	BotUsername   string `json:"bot_username"`    // Username of the bot that owns the game
	GameShortName string `json:"game_short_name"` // Short name of the game
}

// MessageType return the string telegram-type of InternalLinkTypeGame
func (internalLinkTypeGame *InternalLinkTypeGame) MessageType() string {
	return "internalLinkTypeGame"
}

// NewInternalLinkTypeGame creates a new InternalLinkTypeGame
//
// @param botUsername Username of the bot that owns the game
// @param gameShortName Short name of the game
func NewInternalLinkTypeGame(botUsername string, gameShortName string) *InternalLinkTypeGame {
	internalLinkTypeGameTemp := InternalLinkTypeGame{
		tdCommon:      tdCommon{Type: "internalLinkTypeGame"},
		BotUsername:   botUsername,
		GameShortName: gameShortName,
	}

	return &internalLinkTypeGameTemp
}

// GetInternalLinkTypeEnum return the enum type of this object
func (internalLinkTypeGame *InternalLinkTypeGame) GetInternalLinkTypeEnum() InternalLinkTypeEnum {
	return InternalLinkTypeGameType
}

// InternalLinkTypeLanguagePack The link is a link to a language pack. Call getLanguagePackInfo with the given language pack identifier to process the link
type InternalLinkTypeLanguagePack struct {
	tdCommon
	LanguagePackId string `json:"language_pack_id"` // Language pack identifier
}

// MessageType return the string telegram-type of InternalLinkTypeLanguagePack
func (internalLinkTypeLanguagePack *InternalLinkTypeLanguagePack) MessageType() string {
	return "internalLinkTypeLanguagePack"
}

// NewInternalLinkTypeLanguagePack creates a new InternalLinkTypeLanguagePack
//
// @param languagePackId Language pack identifier
func NewInternalLinkTypeLanguagePack(languagePackId string) *InternalLinkTypeLanguagePack {
	internalLinkTypeLanguagePackTemp := InternalLinkTypeLanguagePack{
		tdCommon:       tdCommon{Type: "internalLinkTypeLanguagePack"},
		LanguagePackId: languagePackId,
	}

	return &internalLinkTypeLanguagePackTemp
}

// GetInternalLinkTypeEnum return the enum type of this object
func (internalLinkTypeLanguagePack *InternalLinkTypeLanguagePack) GetInternalLinkTypeEnum() InternalLinkTypeEnum {
	return InternalLinkTypeLanguagePackType
}

// InternalLinkTypeMessage The link is a link to a Telegram message. Call getMessageLinkInfo with the given URL to process the link
type InternalLinkTypeMessage struct {
	tdCommon
	Url string `json:"url"` // URL to be passed to getMessageLinkInfo
}

// MessageType return the string telegram-type of InternalLinkTypeMessage
func (internalLinkTypeMessage *InternalLinkTypeMessage) MessageType() string {
	return "internalLinkTypeMessage"
}

// NewInternalLinkTypeMessage creates a new InternalLinkTypeMessage
//
// @param url URL to be passed to getMessageLinkInfo
func NewInternalLinkTypeMessage(url string) *InternalLinkTypeMessage {
	internalLinkTypeMessageTemp := InternalLinkTypeMessage{
		tdCommon: tdCommon{Type: "internalLinkTypeMessage"},
		Url:      url,
	}

	return &internalLinkTypeMessageTemp
}

// GetInternalLinkTypeEnum return the enum type of this object
func (internalLinkTypeMessage *InternalLinkTypeMessage) GetInternalLinkTypeEnum() InternalLinkTypeEnum {
	return InternalLinkTypeMessageType
}

// InternalLinkTypeMessageDraft The link contains a message draft text. A share screen needs to be shown to the user, then the chosen chat must be opened and the text is added to the input field
type InternalLinkTypeMessageDraft struct {
	tdCommon
	Text         *FormattedText `json:"text"`          // Message draft text
	ContainsLink bool           `json:"contains_link"` // True, if the first line of the text contains a link. If true, the input field needs to be focused and the text after the link must be selected
}

// MessageType return the string telegram-type of InternalLinkTypeMessageDraft
func (internalLinkTypeMessageDraft *InternalLinkTypeMessageDraft) MessageType() string {
	return "internalLinkTypeMessageDraft"
}

// NewInternalLinkTypeMessageDraft creates a new InternalLinkTypeMessageDraft
//
// @param text Message draft text
// @param containsLink True, if the first line of the text contains a link. If true, the input field needs to be focused and the text after the link must be selected
func NewInternalLinkTypeMessageDraft(text *FormattedText, containsLink bool) *InternalLinkTypeMessageDraft {
	internalLinkTypeMessageDraftTemp := InternalLinkTypeMessageDraft{
		tdCommon:     tdCommon{Type: "internalLinkTypeMessageDraft"},
		Text:         text,
		ContainsLink: containsLink,
	}

	return &internalLinkTypeMessageDraftTemp
}

// GetInternalLinkTypeEnum return the enum type of this object
func (internalLinkTypeMessageDraft *InternalLinkTypeMessageDraft) GetInternalLinkTypeEnum() InternalLinkTypeEnum {
	return InternalLinkTypeMessageDraftType
}

// InternalLinkTypePassportDataRequest The link contains a request of Telegram passport data. Call getPassportAuthorizationForm with the given parameters to process the link if the link was received from outside of the app, otherwise ignore it
type InternalLinkTypePassportDataRequest struct {
	tdCommon
	BotUserId   int64  `json:"bot_user_id"`  // User identifier of the service's bot
	Scope       string `json:"scope"`        // Telegram Passport element types requested by the service
	PublicKey   string `json:"public_key"`   // Service's public key
	Nonce       string `json:"nonce"`        // Unique request identifier provided by the service
	CallbackUrl string `json:"callback_url"` // An HTTP URL to open once the request is finished or canceled with the parameter tg_passport=success or tg_passport=cancel respectively. If empty, then the link tgbot{bot_user_id}://passport/success or tgbot{bot_user_id}://passport/cancel needs to be opened instead
}

// MessageType return the string telegram-type of InternalLinkTypePassportDataRequest
func (internalLinkTypePassportDataRequest *InternalLinkTypePassportDataRequest) MessageType() string {
	return "internalLinkTypePassportDataRequest"
}

// NewInternalLinkTypePassportDataRequest creates a new InternalLinkTypePassportDataRequest
//
// @param botUserId User identifier of the service's bot
// @param scope Telegram Passport element types requested by the service
// @param publicKey Service's public key
// @param nonce Unique request identifier provided by the service
// @param callbackUrl An HTTP URL to open once the request is finished or canceled with the parameter tg_passport=success or tg_passport=cancel respectively. If empty, then the link tgbot{bot_user_id}://passport/success or tgbot{bot_user_id}://passport/cancel needs to be opened instead
func NewInternalLinkTypePassportDataRequest(botUserId int64, scope string, publicKey string, nonce string, callbackUrl string) *InternalLinkTypePassportDataRequest {
	internalLinkTypePassportDataRequestTemp := InternalLinkTypePassportDataRequest{
		tdCommon:    tdCommon{Type: "internalLinkTypePassportDataRequest"},
		BotUserId:   botUserId,
		Scope:       scope,
		PublicKey:   publicKey,
		Nonce:       nonce,
		CallbackUrl: callbackUrl,
	}

	return &internalLinkTypePassportDataRequestTemp
}

// GetInternalLinkTypeEnum return the enum type of this object
func (internalLinkTypePassportDataRequest *InternalLinkTypePassportDataRequest) GetInternalLinkTypeEnum() InternalLinkTypeEnum {
	return InternalLinkTypePassportDataRequestType
}

// InternalLinkTypePhoneNumberConfirmation The link can be used to confirm ownership of a phone number to prevent account deletion. Call sendPhoneNumberConfirmationCode with the given hash and phone number to process the link
type InternalLinkTypePhoneNumberConfirmation struct {
	tdCommon
	Hash        string `json:"hash"`         // Hash value from the link
	PhoneNumber string `json:"phone_number"` // Phone number value from the link
}

// MessageType return the string telegram-type of InternalLinkTypePhoneNumberConfirmation
func (internalLinkTypePhoneNumberConfirmation *InternalLinkTypePhoneNumberConfirmation) MessageType() string {
	return "internalLinkTypePhoneNumberConfirmation"
}

// NewInternalLinkTypePhoneNumberConfirmation creates a new InternalLinkTypePhoneNumberConfirmation
//
// @param hash Hash value from the link
// @param phoneNumber Phone number value from the link
func NewInternalLinkTypePhoneNumberConfirmation(hash string, phoneNumber string) *InternalLinkTypePhoneNumberConfirmation {
	internalLinkTypePhoneNumberConfirmationTemp := InternalLinkTypePhoneNumberConfirmation{
		tdCommon:    tdCommon{Type: "internalLinkTypePhoneNumberConfirmation"},
		Hash:        hash,
		PhoneNumber: phoneNumber,
	}

	return &internalLinkTypePhoneNumberConfirmationTemp
}

// GetInternalLinkTypeEnum return the enum type of this object
func (internalLinkTypePhoneNumberConfirmation *InternalLinkTypePhoneNumberConfirmation) GetInternalLinkTypeEnum() InternalLinkTypeEnum {
	return InternalLinkTypePhoneNumberConfirmationType
}

// InternalLinkTypeProxy The link is a link to a proxy. Call addProxy with the given parameters to process the link and add the proxy
type InternalLinkTypeProxy struct {
	tdCommon
	Server string    `json:"server"` // Proxy server IP address
	Port   int32     `json:"port"`   // Proxy server port
	Type   ProxyType `json:"type"`   // Type of the proxy
}

// MessageType return the string telegram-type of InternalLinkTypeProxy
func (internalLinkTypeProxy *InternalLinkTypeProxy) MessageType() string {
	return "internalLinkTypeProxy"
}

// NewInternalLinkTypeProxy creates a new InternalLinkTypeProxy
//
// @param server Proxy server IP address
// @param port Proxy server port
// @param typeParam Type of the proxy
func NewInternalLinkTypeProxy(server string, port int32, typeParam ProxyType) *InternalLinkTypeProxy {
	internalLinkTypeProxyTemp := InternalLinkTypeProxy{
		tdCommon: tdCommon{Type: "internalLinkTypeProxy"},
		Server:   server,
		Port:     port,
		Type:     typeParam,
	}

	return &internalLinkTypeProxyTemp
}

// UnmarshalJSON unmarshal to json
func (internalLinkTypeProxy *InternalLinkTypeProxy) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	tempObj := struct {
		tdCommon
		Server string `json:"server"` // Proxy server IP address
		Port   int32  `json:"port"`   // Proxy server port

	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	internalLinkTypeProxy.tdCommon = tempObj.tdCommon
	internalLinkTypeProxy.Server = tempObj.Server
	internalLinkTypeProxy.Port = tempObj.Port

	fieldType, _ := unmarshalProxyType(objMap["type"])
	internalLinkTypeProxy.Type = fieldType

	return nil
}

// GetInternalLinkTypeEnum return the enum type of this object
func (internalLinkTypeProxy *InternalLinkTypeProxy) GetInternalLinkTypeEnum() InternalLinkTypeEnum {
	return InternalLinkTypeProxyType
}

// InternalLinkTypePublicChat The link is a link to a chat by its username. Call searchPublicChat with the given chat username to process the link
type InternalLinkTypePublicChat struct {
	tdCommon
	ChatUsername string `json:"chat_username"` // Username of the chat
}

// MessageType return the string telegram-type of InternalLinkTypePublicChat
func (internalLinkTypePublicChat *InternalLinkTypePublicChat) MessageType() string {
	return "internalLinkTypePublicChat"
}

// NewInternalLinkTypePublicChat creates a new InternalLinkTypePublicChat
//
// @param chatUsername Username of the chat
func NewInternalLinkTypePublicChat(chatUsername string) *InternalLinkTypePublicChat {
	internalLinkTypePublicChatTemp := InternalLinkTypePublicChat{
		tdCommon:     tdCommon{Type: "internalLinkTypePublicChat"},
		ChatUsername: chatUsername,
	}

	return &internalLinkTypePublicChatTemp
}

// GetInternalLinkTypeEnum return the enum type of this object
func (internalLinkTypePublicChat *InternalLinkTypePublicChat) GetInternalLinkTypeEnum() InternalLinkTypeEnum {
	return InternalLinkTypePublicChatType
}

// InternalLinkTypeQrCodeAuthentication The link can be used to login the current user on another device, but it must be scanned from QR-code using in-app camera. An alert similar to "This code can be used to allow someone to log in to your Telegram account. To confirm Telegram login, please go to Settings > Devices > Scan QR and scan the code" needs to be shown
type InternalLinkTypeQrCodeAuthentication struct {
	tdCommon
}

// MessageType return the string telegram-type of InternalLinkTypeQrCodeAuthentication
func (internalLinkTypeQrCodeAuthentication *InternalLinkTypeQrCodeAuthentication) MessageType() string {
	return "internalLinkTypeQrCodeAuthentication"
}

// NewInternalLinkTypeQrCodeAuthentication creates a new InternalLinkTypeQrCodeAuthentication
//
func NewInternalLinkTypeQrCodeAuthentication() *InternalLinkTypeQrCodeAuthentication {
	internalLinkTypeQrCodeAuthenticationTemp := InternalLinkTypeQrCodeAuthentication{
		tdCommon: tdCommon{Type: "internalLinkTypeQrCodeAuthentication"},
	}

	return &internalLinkTypeQrCodeAuthenticationTemp
}

// GetInternalLinkTypeEnum return the enum type of this object
func (internalLinkTypeQrCodeAuthentication *InternalLinkTypeQrCodeAuthentication) GetInternalLinkTypeEnum() InternalLinkTypeEnum {
	return InternalLinkTypeQrCodeAuthenticationType
}

// InternalLinkTypeSettings The link is a link to app settings
type InternalLinkTypeSettings struct {
	tdCommon
}

// MessageType return the string telegram-type of InternalLinkTypeSettings
func (internalLinkTypeSettings *InternalLinkTypeSettings) MessageType() string {
	return "internalLinkTypeSettings"
}

// NewInternalLinkTypeSettings creates a new InternalLinkTypeSettings
//
func NewInternalLinkTypeSettings() *InternalLinkTypeSettings {
	internalLinkTypeSettingsTemp := InternalLinkTypeSettings{
		tdCommon: tdCommon{Type: "internalLinkTypeSettings"},
	}

	return &internalLinkTypeSettingsTemp
}

// GetInternalLinkTypeEnum return the enum type of this object
func (internalLinkTypeSettings *InternalLinkTypeSettings) GetInternalLinkTypeEnum() InternalLinkTypeEnum {
	return InternalLinkTypeSettingsType
}

// InternalLinkTypeStickerSet The link is a link to a sticker set. Call searchStickerSet with the given sticker set name to process the link and show the sticker set
type InternalLinkTypeStickerSet struct {
	tdCommon
	StickerSetName string `json:"sticker_set_name"` // Name of the sticker set
}

// MessageType return the string telegram-type of InternalLinkTypeStickerSet
func (internalLinkTypeStickerSet *InternalLinkTypeStickerSet) MessageType() string {
	return "internalLinkTypeStickerSet"
}

// NewInternalLinkTypeStickerSet creates a new InternalLinkTypeStickerSet
//
// @param stickerSetName Name of the sticker set
func NewInternalLinkTypeStickerSet(stickerSetName string) *InternalLinkTypeStickerSet {
	internalLinkTypeStickerSetTemp := InternalLinkTypeStickerSet{
		tdCommon:       tdCommon{Type: "internalLinkTypeStickerSet"},
		StickerSetName: stickerSetName,
	}

	return &internalLinkTypeStickerSetTemp
}

// GetInternalLinkTypeEnum return the enum type of this object
func (internalLinkTypeStickerSet *InternalLinkTypeStickerSet) GetInternalLinkTypeEnum() InternalLinkTypeEnum {
	return InternalLinkTypeStickerSetType
}

// InternalLinkTypeTheme The link is a link to a theme. TDLib has no theme support yet
type InternalLinkTypeTheme struct {
	tdCommon
	ThemeName string `json:"theme_name"` // Name of the theme
}

// MessageType return the string telegram-type of InternalLinkTypeTheme
func (internalLinkTypeTheme *InternalLinkTypeTheme) MessageType() string {
	return "internalLinkTypeTheme"
}

// NewInternalLinkTypeTheme creates a new InternalLinkTypeTheme
//
// @param themeName Name of the theme
func NewInternalLinkTypeTheme(themeName string) *InternalLinkTypeTheme {
	internalLinkTypeThemeTemp := InternalLinkTypeTheme{
		tdCommon:  tdCommon{Type: "internalLinkTypeTheme"},
		ThemeName: themeName,
	}

	return &internalLinkTypeThemeTemp
}

// GetInternalLinkTypeEnum return the enum type of this object
func (internalLinkTypeTheme *InternalLinkTypeTheme) GetInternalLinkTypeEnum() InternalLinkTypeEnum {
	return InternalLinkTypeThemeType
}

// InternalLinkTypeThemeSettings The link is a link to the theme settings section of the app
type InternalLinkTypeThemeSettings struct {
	tdCommon
}

// MessageType return the string telegram-type of InternalLinkTypeThemeSettings
func (internalLinkTypeThemeSettings *InternalLinkTypeThemeSettings) MessageType() string {
	return "internalLinkTypeThemeSettings"
}

// NewInternalLinkTypeThemeSettings creates a new InternalLinkTypeThemeSettings
//
func NewInternalLinkTypeThemeSettings() *InternalLinkTypeThemeSettings {
	internalLinkTypeThemeSettingsTemp := InternalLinkTypeThemeSettings{
		tdCommon: tdCommon{Type: "internalLinkTypeThemeSettings"},
	}

	return &internalLinkTypeThemeSettingsTemp
}

// GetInternalLinkTypeEnum return the enum type of this object
func (internalLinkTypeThemeSettings *InternalLinkTypeThemeSettings) GetInternalLinkTypeEnum() InternalLinkTypeEnum {
	return InternalLinkTypeThemeSettingsType
}

// InternalLinkTypeUnknownDeepLink The link is an unknown tg: link. Call getDeepLinkInfo to process the link
type InternalLinkTypeUnknownDeepLink struct {
	tdCommon
	Link string `json:"link"` // Link to be passed to getDeepLinkInfo
}

// MessageType return the string telegram-type of InternalLinkTypeUnknownDeepLink
func (internalLinkTypeUnknownDeepLink *InternalLinkTypeUnknownDeepLink) MessageType() string {
	return "internalLinkTypeUnknownDeepLink"
}

// NewInternalLinkTypeUnknownDeepLink creates a new InternalLinkTypeUnknownDeepLink
//
// @param link Link to be passed to getDeepLinkInfo
func NewInternalLinkTypeUnknownDeepLink(link string) *InternalLinkTypeUnknownDeepLink {
	internalLinkTypeUnknownDeepLinkTemp := InternalLinkTypeUnknownDeepLink{
		tdCommon: tdCommon{Type: "internalLinkTypeUnknownDeepLink"},
		Link:     link,
	}

	return &internalLinkTypeUnknownDeepLinkTemp
}

// GetInternalLinkTypeEnum return the enum type of this object
func (internalLinkTypeUnknownDeepLink *InternalLinkTypeUnknownDeepLink) GetInternalLinkTypeEnum() InternalLinkTypeEnum {
	return InternalLinkTypeUnknownDeepLinkType
}

// InternalLinkTypeUnsupportedProxy The link is a link to an unsupported proxy. An alert can be shown to the user
type InternalLinkTypeUnsupportedProxy struct {
	tdCommon
}

// MessageType return the string telegram-type of InternalLinkTypeUnsupportedProxy
func (internalLinkTypeUnsupportedProxy *InternalLinkTypeUnsupportedProxy) MessageType() string {
	return "internalLinkTypeUnsupportedProxy"
}

// NewInternalLinkTypeUnsupportedProxy creates a new InternalLinkTypeUnsupportedProxy
//
func NewInternalLinkTypeUnsupportedProxy() *InternalLinkTypeUnsupportedProxy {
	internalLinkTypeUnsupportedProxyTemp := InternalLinkTypeUnsupportedProxy{
		tdCommon: tdCommon{Type: "internalLinkTypeUnsupportedProxy"},
	}

	return &internalLinkTypeUnsupportedProxyTemp
}

// GetInternalLinkTypeEnum return the enum type of this object
func (internalLinkTypeUnsupportedProxy *InternalLinkTypeUnsupportedProxy) GetInternalLinkTypeEnum() InternalLinkTypeEnum {
	return InternalLinkTypeUnsupportedProxyType
}

// InternalLinkTypeVideoChat The link is a link to a video chat. Call searchPublicChat with the given chat username, and then joinGoupCall with the given invite hash to process the link
type InternalLinkTypeVideoChat struct {
	tdCommon
	ChatUsername string `json:"chat_username"`  // Username of the chat with the video chat
	InviteHash   string `json:"invite_hash"`    // If non-empty, invite hash to be used to join the video chat without being muted by administrators
	IsLiveStream bool   `json:"is_live_stream"` // True, if the video chat is expected to be a live stream in a channel or a broadcast group
}

// MessageType return the string telegram-type of InternalLinkTypeVideoChat
func (internalLinkTypeVideoChat *InternalLinkTypeVideoChat) MessageType() string {
	return "internalLinkTypeVideoChat"
}

// NewInternalLinkTypeVideoChat creates a new InternalLinkTypeVideoChat
//
// @param chatUsername Username of the chat with the video chat
// @param inviteHash If non-empty, invite hash to be used to join the video chat without being muted by administrators
// @param isLiveStream True, if the video chat is expected to be a live stream in a channel or a broadcast group
func NewInternalLinkTypeVideoChat(chatUsername string, inviteHash string, isLiveStream bool) *InternalLinkTypeVideoChat {
	internalLinkTypeVideoChatTemp := InternalLinkTypeVideoChat{
		tdCommon:     tdCommon{Type: "internalLinkTypeVideoChat"},
		ChatUsername: chatUsername,
		InviteHash:   inviteHash,
		IsLiveStream: isLiveStream,
	}

	return &internalLinkTypeVideoChatTemp
}

// GetInternalLinkTypeEnum return the enum type of this object
func (internalLinkTypeVideoChat *InternalLinkTypeVideoChat) GetInternalLinkTypeEnum() InternalLinkTypeEnum {
	return InternalLinkTypeVideoChatType
}

// GetInternalLinkType Returns information about the type of an internal link. Returns a 404 error if the link is not internal. Can be called before authorization
// @param link The link
func (client *Client) GetInternalLinkType(link string) (InternalLinkType, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "getInternalLinkType",
		"link":  link,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	switch InternalLinkTypeEnum(result.Data["@type"].(string)) {

	case InternalLinkTypeActiveSessionsType:
		var internalLinkType InternalLinkTypeActiveSessions
		err = json.Unmarshal(result.Raw, &internalLinkType)
		return &internalLinkType, err

	case InternalLinkTypeAuthenticationCodeType:
		var internalLinkType InternalLinkTypeAuthenticationCode
		err = json.Unmarshal(result.Raw, &internalLinkType)
		return &internalLinkType, err

	case InternalLinkTypeBackgroundType:
		var internalLinkType InternalLinkTypeBackground
		err = json.Unmarshal(result.Raw, &internalLinkType)
		return &internalLinkType, err

	case InternalLinkTypeBotStartType:
		var internalLinkType InternalLinkTypeBotStart
		err = json.Unmarshal(result.Raw, &internalLinkType)
		return &internalLinkType, err

	case InternalLinkTypeBotStartInGroupType:
		var internalLinkType InternalLinkTypeBotStartInGroup
		err = json.Unmarshal(result.Raw, &internalLinkType)
		return &internalLinkType, err

	case InternalLinkTypeChangePhoneNumberType:
		var internalLinkType InternalLinkTypeChangePhoneNumber
		err = json.Unmarshal(result.Raw, &internalLinkType)
		return &internalLinkType, err

	case InternalLinkTypeChatInviteType:
		var internalLinkType InternalLinkTypeChatInvite
		err = json.Unmarshal(result.Raw, &internalLinkType)
		return &internalLinkType, err

	case InternalLinkTypeFilterSettingsType:
		var internalLinkType InternalLinkTypeFilterSettings
		err = json.Unmarshal(result.Raw, &internalLinkType)
		return &internalLinkType, err

	case InternalLinkTypeGameType:
		var internalLinkType InternalLinkTypeGame
		err = json.Unmarshal(result.Raw, &internalLinkType)
		return &internalLinkType, err

	case InternalLinkTypeLanguagePackType:
		var internalLinkType InternalLinkTypeLanguagePack
		err = json.Unmarshal(result.Raw, &internalLinkType)
		return &internalLinkType, err

	case InternalLinkTypeMessageType:
		var internalLinkType InternalLinkTypeMessage
		err = json.Unmarshal(result.Raw, &internalLinkType)
		return &internalLinkType, err

	case InternalLinkTypeMessageDraftType:
		var internalLinkType InternalLinkTypeMessageDraft
		err = json.Unmarshal(result.Raw, &internalLinkType)
		return &internalLinkType, err

	case InternalLinkTypePassportDataRequestType:
		var internalLinkType InternalLinkTypePassportDataRequest
		err = json.Unmarshal(result.Raw, &internalLinkType)
		return &internalLinkType, err

	case InternalLinkTypePhoneNumberConfirmationType:
		var internalLinkType InternalLinkTypePhoneNumberConfirmation
		err = json.Unmarshal(result.Raw, &internalLinkType)
		return &internalLinkType, err

	case InternalLinkTypeProxyType:
		var internalLinkType InternalLinkTypeProxy
		err = json.Unmarshal(result.Raw, &internalLinkType)
		return &internalLinkType, err

	case InternalLinkTypePublicChatType:
		var internalLinkType InternalLinkTypePublicChat
		err = json.Unmarshal(result.Raw, &internalLinkType)
		return &internalLinkType, err

	case InternalLinkTypeQrCodeAuthenticationType:
		var internalLinkType InternalLinkTypeQrCodeAuthentication
		err = json.Unmarshal(result.Raw, &internalLinkType)
		return &internalLinkType, err

	case InternalLinkTypeSettingsType:
		var internalLinkType InternalLinkTypeSettings
		err = json.Unmarshal(result.Raw, &internalLinkType)
		return &internalLinkType, err

	case InternalLinkTypeStickerSetType:
		var internalLinkType InternalLinkTypeStickerSet
		err = json.Unmarshal(result.Raw, &internalLinkType)
		return &internalLinkType, err

	case InternalLinkTypeThemeType:
		var internalLinkType InternalLinkTypeTheme
		err = json.Unmarshal(result.Raw, &internalLinkType)
		return &internalLinkType, err

	case InternalLinkTypeThemeSettingsType:
		var internalLinkType InternalLinkTypeThemeSettings
		err = json.Unmarshal(result.Raw, &internalLinkType)
		return &internalLinkType, err

	case InternalLinkTypeUnknownDeepLinkType:
		var internalLinkType InternalLinkTypeUnknownDeepLink
		err = json.Unmarshal(result.Raw, &internalLinkType)
		return &internalLinkType, err

	case InternalLinkTypeUnsupportedProxyType:
		var internalLinkType InternalLinkTypeUnsupportedProxy
		err = json.Unmarshal(result.Raw, &internalLinkType)
		return &internalLinkType, err

	case InternalLinkTypeVideoChatType:
		var internalLinkType InternalLinkTypeVideoChat
		err = json.Unmarshal(result.Raw, &internalLinkType)
		return &internalLinkType, err

	default:
		return nil, fmt.Errorf("Invalid type")
	}
}
