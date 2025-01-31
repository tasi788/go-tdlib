package tdlib

import (
	"encoding/json"
	"fmt"
)

// ChatSource Describes a reason why an external chat is shown in a chat list
type ChatSource interface {
	GetChatSourceEnum() ChatSourceEnum
}

// ChatSourceEnum Alias for abstract ChatSource 'Sub-Classes', used as constant-enum here
type ChatSourceEnum string

// ChatSource enums
const (
	ChatSourceMtprotoProxyType              ChatSourceEnum = "chatSourceMtprotoProxy"
	ChatSourcePublicServiceAnnouncementType ChatSourceEnum = "chatSourcePublicServiceAnnouncement"
)

func unmarshalChatSource(rawMsg *json.RawMessage) (ChatSource, error) {

	if rawMsg == nil {
		return nil, nil
	}
	var objMap map[string]interface{}
	err := json.Unmarshal(*rawMsg, &objMap)
	if err != nil {
		return nil, err
	}

	switch ChatSourceEnum(objMap["@type"].(string)) {
	case ChatSourceMtprotoProxyType:
		var chatSourceMtprotoProxy ChatSourceMtprotoProxy
		err := json.Unmarshal(*rawMsg, &chatSourceMtprotoProxy)
		return &chatSourceMtprotoProxy, err

	case ChatSourcePublicServiceAnnouncementType:
		var chatSourcePublicServiceAnnouncement ChatSourcePublicServiceAnnouncement
		err := json.Unmarshal(*rawMsg, &chatSourcePublicServiceAnnouncement)
		return &chatSourcePublicServiceAnnouncement, err

	default:
		return nil, fmt.Errorf("Error UnMarshaling, unknown type:" + objMap["@type"].(string))
	}
}

// ChatSourceMtprotoProxy The chat is sponsored by the user's MTProxy server
type ChatSourceMtprotoProxy struct {
	tdCommon
}

// MessageType return the string telegram-type of ChatSourceMtprotoProxy
func (chatSourceMtprotoProxy *ChatSourceMtprotoProxy) MessageType() string {
	return "chatSourceMtprotoProxy"
}

// NewChatSourceMtprotoProxy creates a new ChatSourceMtprotoProxy
//
func NewChatSourceMtprotoProxy() *ChatSourceMtprotoProxy {
	chatSourceMtprotoProxyTemp := ChatSourceMtprotoProxy{
		tdCommon: tdCommon{Type: "chatSourceMtprotoProxy"},
	}

	return &chatSourceMtprotoProxyTemp
}

// GetChatSourceEnum return the enum type of this object
func (chatSourceMtprotoProxy *ChatSourceMtprotoProxy) GetChatSourceEnum() ChatSourceEnum {
	return ChatSourceMtprotoProxyType
}

// ChatSourcePublicServiceAnnouncement The chat contains a public service announcement
type ChatSourcePublicServiceAnnouncement struct {
	tdCommon
	Type string `json:"type"` // The type of the announcement
	Text string `json:"text"` // The text of the announcement
}

// MessageType return the string telegram-type of ChatSourcePublicServiceAnnouncement
func (chatSourcePublicServiceAnnouncement *ChatSourcePublicServiceAnnouncement) MessageType() string {
	return "chatSourcePublicServiceAnnouncement"
}

// NewChatSourcePublicServiceAnnouncement creates a new ChatSourcePublicServiceAnnouncement
//
// @param typeParam The type of the announcement
// @param text The text of the announcement
func NewChatSourcePublicServiceAnnouncement(typeParam string, text string) *ChatSourcePublicServiceAnnouncement {
	chatSourcePublicServiceAnnouncementTemp := ChatSourcePublicServiceAnnouncement{
		tdCommon: tdCommon{Type: "chatSourcePublicServiceAnnouncement"},
		Type:     typeParam,
		Text:     text,
	}

	return &chatSourcePublicServiceAnnouncementTemp
}

// GetChatSourceEnum return the enum type of this object
func (chatSourcePublicServiceAnnouncement *ChatSourcePublicServiceAnnouncement) GetChatSourceEnum() ChatSourceEnum {
	return ChatSourcePublicServiceAnnouncementType
}
