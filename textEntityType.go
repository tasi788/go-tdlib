package tdlib

import (
	"encoding/json"
	"fmt"
)

// TextEntityType Represents a part of the text which must be formatted differently
type TextEntityType interface {
	GetTextEntityTypeEnum() TextEntityTypeEnum
}

// TextEntityTypeEnum Alias for abstract TextEntityType 'Sub-Classes', used as constant-enum here
type TextEntityTypeEnum string

// TextEntityType enums
const (
	TextEntityTypeMentionType        TextEntityTypeEnum = "textEntityTypeMention"
	TextEntityTypeHashtagType        TextEntityTypeEnum = "textEntityTypeHashtag"
	TextEntityTypeCashtagType        TextEntityTypeEnum = "textEntityTypeCashtag"
	TextEntityTypeBotCommandType     TextEntityTypeEnum = "textEntityTypeBotCommand"
	TextEntityTypeUrlType            TextEntityTypeEnum = "textEntityTypeUrl"
	TextEntityTypeEmailAddressType   TextEntityTypeEnum = "textEntityTypeEmailAddress"
	TextEntityTypePhoneNumberType    TextEntityTypeEnum = "textEntityTypePhoneNumber"
	TextEntityTypeBankCardNumberType TextEntityTypeEnum = "textEntityTypeBankCardNumber"
	TextEntityTypeBoldType           TextEntityTypeEnum = "textEntityTypeBold"
	TextEntityTypeItalicType         TextEntityTypeEnum = "textEntityTypeItalic"
	TextEntityTypeUnderlineType      TextEntityTypeEnum = "textEntityTypeUnderline"
	TextEntityTypeStrikethroughType  TextEntityTypeEnum = "textEntityTypeStrikethrough"
	TextEntityTypeSpoilerType        TextEntityTypeEnum = "textEntityTypeSpoiler"
	TextEntityTypeCodeType           TextEntityTypeEnum = "textEntityTypeCode"
	TextEntityTypePreType            TextEntityTypeEnum = "textEntityTypePre"
	TextEntityTypePreCodeType        TextEntityTypeEnum = "textEntityTypePreCode"
	TextEntityTypeTextUrlType        TextEntityTypeEnum = "textEntityTypeTextUrl"
	TextEntityTypeMentionNameType    TextEntityTypeEnum = "textEntityTypeMentionName"
	TextEntityTypeMediaTimestampType TextEntityTypeEnum = "textEntityTypeMediaTimestamp"
)

func unmarshalTextEntityType(rawMsg *json.RawMessage) (TextEntityType, error) {

	if rawMsg == nil {
		return nil, nil
	}
	var objMap map[string]interface{}
	err := json.Unmarshal(*rawMsg, &objMap)
	if err != nil {
		return nil, err
	}

	switch TextEntityTypeEnum(objMap["@type"].(string)) {
	case TextEntityTypeMentionType:
		var textEntityTypeMention TextEntityTypeMention
		err := json.Unmarshal(*rawMsg, &textEntityTypeMention)
		return &textEntityTypeMention, err

	case TextEntityTypeHashtagType:
		var textEntityTypeHashtag TextEntityTypeHashtag
		err := json.Unmarshal(*rawMsg, &textEntityTypeHashtag)
		return &textEntityTypeHashtag, err

	case TextEntityTypeCashtagType:
		var textEntityTypeCashtag TextEntityTypeCashtag
		err := json.Unmarshal(*rawMsg, &textEntityTypeCashtag)
		return &textEntityTypeCashtag, err

	case TextEntityTypeBotCommandType:
		var textEntityTypeBotCommand TextEntityTypeBotCommand
		err := json.Unmarshal(*rawMsg, &textEntityTypeBotCommand)
		return &textEntityTypeBotCommand, err

	case TextEntityTypeUrlType:
		var textEntityTypeUrl TextEntityTypeUrl
		err := json.Unmarshal(*rawMsg, &textEntityTypeUrl)
		return &textEntityTypeUrl, err

	case TextEntityTypeEmailAddressType:
		var textEntityTypeEmailAddress TextEntityTypeEmailAddress
		err := json.Unmarshal(*rawMsg, &textEntityTypeEmailAddress)
		return &textEntityTypeEmailAddress, err

	case TextEntityTypePhoneNumberType:
		var textEntityTypePhoneNumber TextEntityTypePhoneNumber
		err := json.Unmarshal(*rawMsg, &textEntityTypePhoneNumber)
		return &textEntityTypePhoneNumber, err

	case TextEntityTypeBankCardNumberType:
		var textEntityTypeBankCardNumber TextEntityTypeBankCardNumber
		err := json.Unmarshal(*rawMsg, &textEntityTypeBankCardNumber)
		return &textEntityTypeBankCardNumber, err

	case TextEntityTypeBoldType:
		var textEntityTypeBold TextEntityTypeBold
		err := json.Unmarshal(*rawMsg, &textEntityTypeBold)
		return &textEntityTypeBold, err

	case TextEntityTypeItalicType:
		var textEntityTypeItalic TextEntityTypeItalic
		err := json.Unmarshal(*rawMsg, &textEntityTypeItalic)
		return &textEntityTypeItalic, err

	case TextEntityTypeUnderlineType:
		var textEntityTypeUnderline TextEntityTypeUnderline
		err := json.Unmarshal(*rawMsg, &textEntityTypeUnderline)
		return &textEntityTypeUnderline, err

	case TextEntityTypeStrikethroughType:
		var textEntityTypeStrikethrough TextEntityTypeStrikethrough
		err := json.Unmarshal(*rawMsg, &textEntityTypeStrikethrough)
		return &textEntityTypeStrikethrough, err

	case TextEntityTypeSpoilerType:
		var textEntityTypeSpoiler TextEntityTypeSpoiler
		err := json.Unmarshal(*rawMsg, &textEntityTypeSpoiler)
		return &textEntityTypeSpoiler, err

	case TextEntityTypeCodeType:
		var textEntityTypeCode TextEntityTypeCode
		err := json.Unmarshal(*rawMsg, &textEntityTypeCode)
		return &textEntityTypeCode, err

	case TextEntityTypePreType:
		var textEntityTypePre TextEntityTypePre
		err := json.Unmarshal(*rawMsg, &textEntityTypePre)
		return &textEntityTypePre, err

	case TextEntityTypePreCodeType:
		var textEntityTypePreCode TextEntityTypePreCode
		err := json.Unmarshal(*rawMsg, &textEntityTypePreCode)
		return &textEntityTypePreCode, err

	case TextEntityTypeTextUrlType:
		var textEntityTypeTextUrl TextEntityTypeTextUrl
		err := json.Unmarshal(*rawMsg, &textEntityTypeTextUrl)
		return &textEntityTypeTextUrl, err

	case TextEntityTypeMentionNameType:
		var textEntityTypeMentionName TextEntityTypeMentionName
		err := json.Unmarshal(*rawMsg, &textEntityTypeMentionName)
		return &textEntityTypeMentionName, err

	case TextEntityTypeMediaTimestampType:
		var textEntityTypeMediaTimestamp TextEntityTypeMediaTimestamp
		err := json.Unmarshal(*rawMsg, &textEntityTypeMediaTimestamp)
		return &textEntityTypeMediaTimestamp, err

	default:
		return nil, fmt.Errorf("Error UnMarshaling, unknown type:" + objMap["@type"].(string))
	}
}

// TextEntityTypeMention A mention of a user by their username
type TextEntityTypeMention struct {
	tdCommon
}

// MessageType return the string telegram-type of TextEntityTypeMention
func (textEntityTypeMention *TextEntityTypeMention) MessageType() string {
	return "textEntityTypeMention"
}

// NewTextEntityTypeMention creates a new TextEntityTypeMention
//
func NewTextEntityTypeMention() *TextEntityTypeMention {
	textEntityTypeMentionTemp := TextEntityTypeMention{
		tdCommon: tdCommon{Type: "textEntityTypeMention"},
	}

	return &textEntityTypeMentionTemp
}

// GetTextEntityTypeEnum return the enum type of this object
func (textEntityTypeMention *TextEntityTypeMention) GetTextEntityTypeEnum() TextEntityTypeEnum {
	return TextEntityTypeMentionType
}

// TextEntityTypeHashtag A hashtag text, beginning with "#"
type TextEntityTypeHashtag struct {
	tdCommon
}

// MessageType return the string telegram-type of TextEntityTypeHashtag
func (textEntityTypeHashtag *TextEntityTypeHashtag) MessageType() string {
	return "textEntityTypeHashtag"
}

// NewTextEntityTypeHashtag creates a new TextEntityTypeHashtag
//
func NewTextEntityTypeHashtag() *TextEntityTypeHashtag {
	textEntityTypeHashtagTemp := TextEntityTypeHashtag{
		tdCommon: tdCommon{Type: "textEntityTypeHashtag"},
	}

	return &textEntityTypeHashtagTemp
}

// GetTextEntityTypeEnum return the enum type of this object
func (textEntityTypeHashtag *TextEntityTypeHashtag) GetTextEntityTypeEnum() TextEntityTypeEnum {
	return TextEntityTypeHashtagType
}

// TextEntityTypeCashtag A cashtag text, beginning with "$" and consisting of capital English letters (e.g., "$USD")
type TextEntityTypeCashtag struct {
	tdCommon
}

// MessageType return the string telegram-type of TextEntityTypeCashtag
func (textEntityTypeCashtag *TextEntityTypeCashtag) MessageType() string {
	return "textEntityTypeCashtag"
}

// NewTextEntityTypeCashtag creates a new TextEntityTypeCashtag
//
func NewTextEntityTypeCashtag() *TextEntityTypeCashtag {
	textEntityTypeCashtagTemp := TextEntityTypeCashtag{
		tdCommon: tdCommon{Type: "textEntityTypeCashtag"},
	}

	return &textEntityTypeCashtagTemp
}

// GetTextEntityTypeEnum return the enum type of this object
func (textEntityTypeCashtag *TextEntityTypeCashtag) GetTextEntityTypeEnum() TextEntityTypeEnum {
	return TextEntityTypeCashtagType
}

// TextEntityTypeBotCommand A bot command, beginning with "/"
type TextEntityTypeBotCommand struct {
	tdCommon
}

// MessageType return the string telegram-type of TextEntityTypeBotCommand
func (textEntityTypeBotCommand *TextEntityTypeBotCommand) MessageType() string {
	return "textEntityTypeBotCommand"
}

// NewTextEntityTypeBotCommand creates a new TextEntityTypeBotCommand
//
func NewTextEntityTypeBotCommand() *TextEntityTypeBotCommand {
	textEntityTypeBotCommandTemp := TextEntityTypeBotCommand{
		tdCommon: tdCommon{Type: "textEntityTypeBotCommand"},
	}

	return &textEntityTypeBotCommandTemp
}

// GetTextEntityTypeEnum return the enum type of this object
func (textEntityTypeBotCommand *TextEntityTypeBotCommand) GetTextEntityTypeEnum() TextEntityTypeEnum {
	return TextEntityTypeBotCommandType
}

// TextEntityTypeUrl An HTTP URL
type TextEntityTypeUrl struct {
	tdCommon
}

// MessageType return the string telegram-type of TextEntityTypeUrl
func (textEntityTypeUrl *TextEntityTypeUrl) MessageType() string {
	return "textEntityTypeUrl"
}

// NewTextEntityTypeUrl creates a new TextEntityTypeUrl
//
func NewTextEntityTypeUrl() *TextEntityTypeUrl {
	textEntityTypeUrlTemp := TextEntityTypeUrl{
		tdCommon: tdCommon{Type: "textEntityTypeUrl"},
	}

	return &textEntityTypeUrlTemp
}

// GetTextEntityTypeEnum return the enum type of this object
func (textEntityTypeUrl *TextEntityTypeUrl) GetTextEntityTypeEnum() TextEntityTypeEnum {
	return TextEntityTypeUrlType
}

// TextEntityTypeEmailAddress An email address
type TextEntityTypeEmailAddress struct {
	tdCommon
}

// MessageType return the string telegram-type of TextEntityTypeEmailAddress
func (textEntityTypeEmailAddress *TextEntityTypeEmailAddress) MessageType() string {
	return "textEntityTypeEmailAddress"
}

// NewTextEntityTypeEmailAddress creates a new TextEntityTypeEmailAddress
//
func NewTextEntityTypeEmailAddress() *TextEntityTypeEmailAddress {
	textEntityTypeEmailAddressTemp := TextEntityTypeEmailAddress{
		tdCommon: tdCommon{Type: "textEntityTypeEmailAddress"},
	}

	return &textEntityTypeEmailAddressTemp
}

// GetTextEntityTypeEnum return the enum type of this object
func (textEntityTypeEmailAddress *TextEntityTypeEmailAddress) GetTextEntityTypeEnum() TextEntityTypeEnum {
	return TextEntityTypeEmailAddressType
}

// TextEntityTypePhoneNumber A phone number
type TextEntityTypePhoneNumber struct {
	tdCommon
}

// MessageType return the string telegram-type of TextEntityTypePhoneNumber
func (textEntityTypePhoneNumber *TextEntityTypePhoneNumber) MessageType() string {
	return "textEntityTypePhoneNumber"
}

// NewTextEntityTypePhoneNumber creates a new TextEntityTypePhoneNumber
//
func NewTextEntityTypePhoneNumber() *TextEntityTypePhoneNumber {
	textEntityTypePhoneNumberTemp := TextEntityTypePhoneNumber{
		tdCommon: tdCommon{Type: "textEntityTypePhoneNumber"},
	}

	return &textEntityTypePhoneNumberTemp
}

// GetTextEntityTypeEnum return the enum type of this object
func (textEntityTypePhoneNumber *TextEntityTypePhoneNumber) GetTextEntityTypeEnum() TextEntityTypeEnum {
	return TextEntityTypePhoneNumberType
}

// TextEntityTypeBankCardNumber A bank card number. The getBankCardInfo method can be used to get information about the bank card
type TextEntityTypeBankCardNumber struct {
	tdCommon
}

// MessageType return the string telegram-type of TextEntityTypeBankCardNumber
func (textEntityTypeBankCardNumber *TextEntityTypeBankCardNumber) MessageType() string {
	return "textEntityTypeBankCardNumber"
}

// NewTextEntityTypeBankCardNumber creates a new TextEntityTypeBankCardNumber
//
func NewTextEntityTypeBankCardNumber() *TextEntityTypeBankCardNumber {
	textEntityTypeBankCardNumberTemp := TextEntityTypeBankCardNumber{
		tdCommon: tdCommon{Type: "textEntityTypeBankCardNumber"},
	}

	return &textEntityTypeBankCardNumberTemp
}

// GetTextEntityTypeEnum return the enum type of this object
func (textEntityTypeBankCardNumber *TextEntityTypeBankCardNumber) GetTextEntityTypeEnum() TextEntityTypeEnum {
	return TextEntityTypeBankCardNumberType
}

// TextEntityTypeBold A bold text
type TextEntityTypeBold struct {
	tdCommon
}

// MessageType return the string telegram-type of TextEntityTypeBold
func (textEntityTypeBold *TextEntityTypeBold) MessageType() string {
	return "textEntityTypeBold"
}

// NewTextEntityTypeBold creates a new TextEntityTypeBold
//
func NewTextEntityTypeBold() *TextEntityTypeBold {
	textEntityTypeBoldTemp := TextEntityTypeBold{
		tdCommon: tdCommon{Type: "textEntityTypeBold"},
	}

	return &textEntityTypeBoldTemp
}

// GetTextEntityTypeEnum return the enum type of this object
func (textEntityTypeBold *TextEntityTypeBold) GetTextEntityTypeEnum() TextEntityTypeEnum {
	return TextEntityTypeBoldType
}

// TextEntityTypeItalic An italic text
type TextEntityTypeItalic struct {
	tdCommon
}

// MessageType return the string telegram-type of TextEntityTypeItalic
func (textEntityTypeItalic *TextEntityTypeItalic) MessageType() string {
	return "textEntityTypeItalic"
}

// NewTextEntityTypeItalic creates a new TextEntityTypeItalic
//
func NewTextEntityTypeItalic() *TextEntityTypeItalic {
	textEntityTypeItalicTemp := TextEntityTypeItalic{
		tdCommon: tdCommon{Type: "textEntityTypeItalic"},
	}

	return &textEntityTypeItalicTemp
}

// GetTextEntityTypeEnum return the enum type of this object
func (textEntityTypeItalic *TextEntityTypeItalic) GetTextEntityTypeEnum() TextEntityTypeEnum {
	return TextEntityTypeItalicType
}

// TextEntityTypeUnderline An underlined text
type TextEntityTypeUnderline struct {
	tdCommon
}

// MessageType return the string telegram-type of TextEntityTypeUnderline
func (textEntityTypeUnderline *TextEntityTypeUnderline) MessageType() string {
	return "textEntityTypeUnderline"
}

// NewTextEntityTypeUnderline creates a new TextEntityTypeUnderline
//
func NewTextEntityTypeUnderline() *TextEntityTypeUnderline {
	textEntityTypeUnderlineTemp := TextEntityTypeUnderline{
		tdCommon: tdCommon{Type: "textEntityTypeUnderline"},
	}

	return &textEntityTypeUnderlineTemp
}

// GetTextEntityTypeEnum return the enum type of this object
func (textEntityTypeUnderline *TextEntityTypeUnderline) GetTextEntityTypeEnum() TextEntityTypeEnum {
	return TextEntityTypeUnderlineType
}

// TextEntityTypeStrikethrough A strikethrough text
type TextEntityTypeStrikethrough struct {
	tdCommon
}

// MessageType return the string telegram-type of TextEntityTypeStrikethrough
func (textEntityTypeStrikethrough *TextEntityTypeStrikethrough) MessageType() string {
	return "textEntityTypeStrikethrough"
}

// NewTextEntityTypeStrikethrough creates a new TextEntityTypeStrikethrough
//
func NewTextEntityTypeStrikethrough() *TextEntityTypeStrikethrough {
	textEntityTypeStrikethroughTemp := TextEntityTypeStrikethrough{
		tdCommon: tdCommon{Type: "textEntityTypeStrikethrough"},
	}

	return &textEntityTypeStrikethroughTemp
}

// GetTextEntityTypeEnum return the enum type of this object
func (textEntityTypeStrikethrough *TextEntityTypeStrikethrough) GetTextEntityTypeEnum() TextEntityTypeEnum {
	return TextEntityTypeStrikethroughType
}

// TextEntityTypeSpoiler A spoiler text. Not supported in secret chats
type TextEntityTypeSpoiler struct {
	tdCommon
}

// MessageType return the string telegram-type of TextEntityTypeSpoiler
func (textEntityTypeSpoiler *TextEntityTypeSpoiler) MessageType() string {
	return "textEntityTypeSpoiler"
}

// NewTextEntityTypeSpoiler creates a new TextEntityTypeSpoiler
//
func NewTextEntityTypeSpoiler() *TextEntityTypeSpoiler {
	textEntityTypeSpoilerTemp := TextEntityTypeSpoiler{
		tdCommon: tdCommon{Type: "textEntityTypeSpoiler"},
	}

	return &textEntityTypeSpoilerTemp
}

// GetTextEntityTypeEnum return the enum type of this object
func (textEntityTypeSpoiler *TextEntityTypeSpoiler) GetTextEntityTypeEnum() TextEntityTypeEnum {
	return TextEntityTypeSpoilerType
}

// TextEntityTypeCode Text that must be formatted as if inside a code HTML tag
type TextEntityTypeCode struct {
	tdCommon
}

// MessageType return the string telegram-type of TextEntityTypeCode
func (textEntityTypeCode *TextEntityTypeCode) MessageType() string {
	return "textEntityTypeCode"
}

// NewTextEntityTypeCode creates a new TextEntityTypeCode
//
func NewTextEntityTypeCode() *TextEntityTypeCode {
	textEntityTypeCodeTemp := TextEntityTypeCode{
		tdCommon: tdCommon{Type: "textEntityTypeCode"},
	}

	return &textEntityTypeCodeTemp
}

// GetTextEntityTypeEnum return the enum type of this object
func (textEntityTypeCode *TextEntityTypeCode) GetTextEntityTypeEnum() TextEntityTypeEnum {
	return TextEntityTypeCodeType
}

// TextEntityTypePre Text that must be formatted as if inside a pre HTML tag
type TextEntityTypePre struct {
	tdCommon
}

// MessageType return the string telegram-type of TextEntityTypePre
func (textEntityTypePre *TextEntityTypePre) MessageType() string {
	return "textEntityTypePre"
}

// NewTextEntityTypePre creates a new TextEntityTypePre
//
func NewTextEntityTypePre() *TextEntityTypePre {
	textEntityTypePreTemp := TextEntityTypePre{
		tdCommon: tdCommon{Type: "textEntityTypePre"},
	}

	return &textEntityTypePreTemp
}

// GetTextEntityTypeEnum return the enum type of this object
func (textEntityTypePre *TextEntityTypePre) GetTextEntityTypeEnum() TextEntityTypeEnum {
	return TextEntityTypePreType
}

// TextEntityTypePreCode Text that must be formatted as if inside pre, and code HTML tags
type TextEntityTypePreCode struct {
	tdCommon
	Language string `json:"language"` // Programming language of the code; as defined by the sender
}

// MessageType return the string telegram-type of TextEntityTypePreCode
func (textEntityTypePreCode *TextEntityTypePreCode) MessageType() string {
	return "textEntityTypePreCode"
}

// NewTextEntityTypePreCode creates a new TextEntityTypePreCode
//
// @param language Programming language of the code; as defined by the sender
func NewTextEntityTypePreCode(language string) *TextEntityTypePreCode {
	textEntityTypePreCodeTemp := TextEntityTypePreCode{
		tdCommon: tdCommon{Type: "textEntityTypePreCode"},
		Language: language,
	}

	return &textEntityTypePreCodeTemp
}

// GetTextEntityTypeEnum return the enum type of this object
func (textEntityTypePreCode *TextEntityTypePreCode) GetTextEntityTypeEnum() TextEntityTypeEnum {
	return TextEntityTypePreCodeType
}

// TextEntityTypeTextUrl A text description shown instead of a raw URL
type TextEntityTypeTextUrl struct {
	tdCommon
	Url string `json:"url"` // HTTP or tg:// URL to be opened when the link is clicked
}

// MessageType return the string telegram-type of TextEntityTypeTextUrl
func (textEntityTypeTextUrl *TextEntityTypeTextUrl) MessageType() string {
	return "textEntityTypeTextUrl"
}

// NewTextEntityTypeTextUrl creates a new TextEntityTypeTextUrl
//
// @param url HTTP or tg:// URL to be opened when the link is clicked
func NewTextEntityTypeTextUrl(url string) *TextEntityTypeTextUrl {
	textEntityTypeTextUrlTemp := TextEntityTypeTextUrl{
		tdCommon: tdCommon{Type: "textEntityTypeTextUrl"},
		Url:      url,
	}

	return &textEntityTypeTextUrlTemp
}

// GetTextEntityTypeEnum return the enum type of this object
func (textEntityTypeTextUrl *TextEntityTypeTextUrl) GetTextEntityTypeEnum() TextEntityTypeEnum {
	return TextEntityTypeTextUrlType
}

// TextEntityTypeMentionName A text shows instead of a raw mention of the user (e.g., when the user has no username)
type TextEntityTypeMentionName struct {
	tdCommon
	UserId int64 `json:"user_id"` // Identifier of the mentioned user
}

// MessageType return the string telegram-type of TextEntityTypeMentionName
func (textEntityTypeMentionName *TextEntityTypeMentionName) MessageType() string {
	return "textEntityTypeMentionName"
}

// NewTextEntityTypeMentionName creates a new TextEntityTypeMentionName
//
// @param userId Identifier of the mentioned user
func NewTextEntityTypeMentionName(userId int64) *TextEntityTypeMentionName {
	textEntityTypeMentionNameTemp := TextEntityTypeMentionName{
		tdCommon: tdCommon{Type: "textEntityTypeMentionName"},
		UserId:   userId,
	}

	return &textEntityTypeMentionNameTemp
}

// GetTextEntityTypeEnum return the enum type of this object
func (textEntityTypeMentionName *TextEntityTypeMentionName) GetTextEntityTypeEnum() TextEntityTypeEnum {
	return TextEntityTypeMentionNameType
}

// TextEntityTypeMediaTimestamp A media timestamp
type TextEntityTypeMediaTimestamp struct {
	tdCommon
	MediaTimestamp int32 `json:"media_timestamp"` // Timestamp from which a video/audio/video note/voice note playing must start, in seconds. The media can be in the content or the web page preview of the current message, or in the same places in the replied message
}

// MessageType return the string telegram-type of TextEntityTypeMediaTimestamp
func (textEntityTypeMediaTimestamp *TextEntityTypeMediaTimestamp) MessageType() string {
	return "textEntityTypeMediaTimestamp"
}

// NewTextEntityTypeMediaTimestamp creates a new TextEntityTypeMediaTimestamp
//
// @param mediaTimestamp Timestamp from which a video/audio/video note/voice note playing must start, in seconds. The media can be in the content or the web page preview of the current message, or in the same places in the replied message
func NewTextEntityTypeMediaTimestamp(mediaTimestamp int32) *TextEntityTypeMediaTimestamp {
	textEntityTypeMediaTimestampTemp := TextEntityTypeMediaTimestamp{
		tdCommon:       tdCommon{Type: "textEntityTypeMediaTimestamp"},
		MediaTimestamp: mediaTimestamp,
	}

	return &textEntityTypeMediaTimestampTemp
}

// GetTextEntityTypeEnum return the enum type of this object
func (textEntityTypeMediaTimestamp *TextEntityTypeMediaTimestamp) GetTextEntityTypeEnum() TextEntityTypeEnum {
	return TextEntityTypeMediaTimestampType
}
