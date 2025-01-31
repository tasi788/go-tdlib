package tdlib

import (
	"encoding/json"
	"fmt"
)

// ChatReportReason Describes the reason why a chat is reported
type ChatReportReason interface {
	GetChatReportReasonEnum() ChatReportReasonEnum
}

// ChatReportReasonEnum Alias for abstract ChatReportReason 'Sub-Classes', used as constant-enum here
type ChatReportReasonEnum string

// ChatReportReason enums
const (
	ChatReportReasonSpamType              ChatReportReasonEnum = "chatReportReasonSpam"
	ChatReportReasonViolenceType          ChatReportReasonEnum = "chatReportReasonViolence"
	ChatReportReasonPornographyType       ChatReportReasonEnum = "chatReportReasonPornography"
	ChatReportReasonChildAbuseType        ChatReportReasonEnum = "chatReportReasonChildAbuse"
	ChatReportReasonCopyrightType         ChatReportReasonEnum = "chatReportReasonCopyright"
	ChatReportReasonUnrelatedLocationType ChatReportReasonEnum = "chatReportReasonUnrelatedLocation"
	ChatReportReasonFakeType              ChatReportReasonEnum = "chatReportReasonFake"
	ChatReportReasonCustomType            ChatReportReasonEnum = "chatReportReasonCustom"
)

func unmarshalChatReportReason(rawMsg *json.RawMessage) (ChatReportReason, error) {

	if rawMsg == nil {
		return nil, nil
	}
	var objMap map[string]interface{}
	err := json.Unmarshal(*rawMsg, &objMap)
	if err != nil {
		return nil, err
	}

	switch ChatReportReasonEnum(objMap["@type"].(string)) {
	case ChatReportReasonSpamType:
		var chatReportReasonSpam ChatReportReasonSpam
		err := json.Unmarshal(*rawMsg, &chatReportReasonSpam)
		return &chatReportReasonSpam, err

	case ChatReportReasonViolenceType:
		var chatReportReasonViolence ChatReportReasonViolence
		err := json.Unmarshal(*rawMsg, &chatReportReasonViolence)
		return &chatReportReasonViolence, err

	case ChatReportReasonPornographyType:
		var chatReportReasonPornography ChatReportReasonPornography
		err := json.Unmarshal(*rawMsg, &chatReportReasonPornography)
		return &chatReportReasonPornography, err

	case ChatReportReasonChildAbuseType:
		var chatReportReasonChildAbuse ChatReportReasonChildAbuse
		err := json.Unmarshal(*rawMsg, &chatReportReasonChildAbuse)
		return &chatReportReasonChildAbuse, err

	case ChatReportReasonCopyrightType:
		var chatReportReasonCopyright ChatReportReasonCopyright
		err := json.Unmarshal(*rawMsg, &chatReportReasonCopyright)
		return &chatReportReasonCopyright, err

	case ChatReportReasonUnrelatedLocationType:
		var chatReportReasonUnrelatedLocation ChatReportReasonUnrelatedLocation
		err := json.Unmarshal(*rawMsg, &chatReportReasonUnrelatedLocation)
		return &chatReportReasonUnrelatedLocation, err

	case ChatReportReasonFakeType:
		var chatReportReasonFake ChatReportReasonFake
		err := json.Unmarshal(*rawMsg, &chatReportReasonFake)
		return &chatReportReasonFake, err

	case ChatReportReasonCustomType:
		var chatReportReasonCustom ChatReportReasonCustom
		err := json.Unmarshal(*rawMsg, &chatReportReasonCustom)
		return &chatReportReasonCustom, err

	default:
		return nil, fmt.Errorf("Error UnMarshaling, unknown type:" + objMap["@type"].(string))
	}
}

// ChatReportReasonSpam The chat contains spam messages
type ChatReportReasonSpam struct {
	tdCommon
}

// MessageType return the string telegram-type of ChatReportReasonSpam
func (chatReportReasonSpam *ChatReportReasonSpam) MessageType() string {
	return "chatReportReasonSpam"
}

// NewChatReportReasonSpam creates a new ChatReportReasonSpam
//
func NewChatReportReasonSpam() *ChatReportReasonSpam {
	chatReportReasonSpamTemp := ChatReportReasonSpam{
		tdCommon: tdCommon{Type: "chatReportReasonSpam"},
	}

	return &chatReportReasonSpamTemp
}

// GetChatReportReasonEnum return the enum type of this object
func (chatReportReasonSpam *ChatReportReasonSpam) GetChatReportReasonEnum() ChatReportReasonEnum {
	return ChatReportReasonSpamType
}

// ChatReportReasonViolence The chat promotes violence
type ChatReportReasonViolence struct {
	tdCommon
}

// MessageType return the string telegram-type of ChatReportReasonViolence
func (chatReportReasonViolence *ChatReportReasonViolence) MessageType() string {
	return "chatReportReasonViolence"
}

// NewChatReportReasonViolence creates a new ChatReportReasonViolence
//
func NewChatReportReasonViolence() *ChatReportReasonViolence {
	chatReportReasonViolenceTemp := ChatReportReasonViolence{
		tdCommon: tdCommon{Type: "chatReportReasonViolence"},
	}

	return &chatReportReasonViolenceTemp
}

// GetChatReportReasonEnum return the enum type of this object
func (chatReportReasonViolence *ChatReportReasonViolence) GetChatReportReasonEnum() ChatReportReasonEnum {
	return ChatReportReasonViolenceType
}

// ChatReportReasonPornography The chat contains pornographic messages
type ChatReportReasonPornography struct {
	tdCommon
}

// MessageType return the string telegram-type of ChatReportReasonPornography
func (chatReportReasonPornography *ChatReportReasonPornography) MessageType() string {
	return "chatReportReasonPornography"
}

// NewChatReportReasonPornography creates a new ChatReportReasonPornography
//
func NewChatReportReasonPornography() *ChatReportReasonPornography {
	chatReportReasonPornographyTemp := ChatReportReasonPornography{
		tdCommon: tdCommon{Type: "chatReportReasonPornography"},
	}

	return &chatReportReasonPornographyTemp
}

// GetChatReportReasonEnum return the enum type of this object
func (chatReportReasonPornography *ChatReportReasonPornography) GetChatReportReasonEnum() ChatReportReasonEnum {
	return ChatReportReasonPornographyType
}

// ChatReportReasonChildAbuse The chat has child abuse related content
type ChatReportReasonChildAbuse struct {
	tdCommon
}

// MessageType return the string telegram-type of ChatReportReasonChildAbuse
func (chatReportReasonChildAbuse *ChatReportReasonChildAbuse) MessageType() string {
	return "chatReportReasonChildAbuse"
}

// NewChatReportReasonChildAbuse creates a new ChatReportReasonChildAbuse
//
func NewChatReportReasonChildAbuse() *ChatReportReasonChildAbuse {
	chatReportReasonChildAbuseTemp := ChatReportReasonChildAbuse{
		tdCommon: tdCommon{Type: "chatReportReasonChildAbuse"},
	}

	return &chatReportReasonChildAbuseTemp
}

// GetChatReportReasonEnum return the enum type of this object
func (chatReportReasonChildAbuse *ChatReportReasonChildAbuse) GetChatReportReasonEnum() ChatReportReasonEnum {
	return ChatReportReasonChildAbuseType
}

// ChatReportReasonCopyright The chat contains copyrighted content
type ChatReportReasonCopyright struct {
	tdCommon
}

// MessageType return the string telegram-type of ChatReportReasonCopyright
func (chatReportReasonCopyright *ChatReportReasonCopyright) MessageType() string {
	return "chatReportReasonCopyright"
}

// NewChatReportReasonCopyright creates a new ChatReportReasonCopyright
//
func NewChatReportReasonCopyright() *ChatReportReasonCopyright {
	chatReportReasonCopyrightTemp := ChatReportReasonCopyright{
		tdCommon: tdCommon{Type: "chatReportReasonCopyright"},
	}

	return &chatReportReasonCopyrightTemp
}

// GetChatReportReasonEnum return the enum type of this object
func (chatReportReasonCopyright *ChatReportReasonCopyright) GetChatReportReasonEnum() ChatReportReasonEnum {
	return ChatReportReasonCopyrightType
}

// ChatReportReasonUnrelatedLocation The location-based chat is unrelated to its stated location
type ChatReportReasonUnrelatedLocation struct {
	tdCommon
}

// MessageType return the string telegram-type of ChatReportReasonUnrelatedLocation
func (chatReportReasonUnrelatedLocation *ChatReportReasonUnrelatedLocation) MessageType() string {
	return "chatReportReasonUnrelatedLocation"
}

// NewChatReportReasonUnrelatedLocation creates a new ChatReportReasonUnrelatedLocation
//
func NewChatReportReasonUnrelatedLocation() *ChatReportReasonUnrelatedLocation {
	chatReportReasonUnrelatedLocationTemp := ChatReportReasonUnrelatedLocation{
		tdCommon: tdCommon{Type: "chatReportReasonUnrelatedLocation"},
	}

	return &chatReportReasonUnrelatedLocationTemp
}

// GetChatReportReasonEnum return the enum type of this object
func (chatReportReasonUnrelatedLocation *ChatReportReasonUnrelatedLocation) GetChatReportReasonEnum() ChatReportReasonEnum {
	return ChatReportReasonUnrelatedLocationType
}

// ChatReportReasonFake The chat represents a fake account
type ChatReportReasonFake struct {
	tdCommon
}

// MessageType return the string telegram-type of ChatReportReasonFake
func (chatReportReasonFake *ChatReportReasonFake) MessageType() string {
	return "chatReportReasonFake"
}

// NewChatReportReasonFake creates a new ChatReportReasonFake
//
func NewChatReportReasonFake() *ChatReportReasonFake {
	chatReportReasonFakeTemp := ChatReportReasonFake{
		tdCommon: tdCommon{Type: "chatReportReasonFake"},
	}

	return &chatReportReasonFakeTemp
}

// GetChatReportReasonEnum return the enum type of this object
func (chatReportReasonFake *ChatReportReasonFake) GetChatReportReasonEnum() ChatReportReasonEnum {
	return ChatReportReasonFakeType
}

// ChatReportReasonCustom A custom reason provided by the user
type ChatReportReasonCustom struct {
	tdCommon
}

// MessageType return the string telegram-type of ChatReportReasonCustom
func (chatReportReasonCustom *ChatReportReasonCustom) MessageType() string {
	return "chatReportReasonCustom"
}

// NewChatReportReasonCustom creates a new ChatReportReasonCustom
//
func NewChatReportReasonCustom() *ChatReportReasonCustom {
	chatReportReasonCustomTemp := ChatReportReasonCustom{
		tdCommon: tdCommon{Type: "chatReportReasonCustom"},
	}

	return &chatReportReasonCustomTemp
}

// GetChatReportReasonEnum return the enum type of this object
func (chatReportReasonCustom *ChatReportReasonCustom) GetChatReportReasonEnum() ChatReportReasonEnum {
	return ChatReportReasonCustomType
}
