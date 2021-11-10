package tdlib

import (
	"encoding/json"
)

// MessageForwardInfo Contains information about a forwarded message
type MessageForwardInfo struct {
	tdCommon
	Origin                        MessageForwardOrigin `json:"origin"`                           // Origin of a forwarded message
	Date                          int32                `json:"date"`                             // Point in time (Unix timestamp) when the message was originally sent
	PublicServiceAnnouncementType string               `json:"public_service_announcement_type"` // The type of a public service announcement for the forwarded message
	FromChatId                    int64                `json:"from_chat_id"`                     // For messages forwarded to the chat with the current user (Saved Messages), to the Replies bot chat, or to the channel's discussion group, the identifier of the chat from which the message was forwarded last time; 0 if unknown
	FromMessageId                 int64                `json:"from_message_id"`                  // For messages forwarded to the chat with the current user (Saved Messages), to the Replies bot chat, or to the channel's discussion group, the identifier of the original message from which the new message was forwarded last time; 0 if unknown
}

// MessageType return the string telegram-type of MessageForwardInfo
func (messageForwardInfo *MessageForwardInfo) MessageType() string {
	return "messageForwardInfo"
}

// NewMessageForwardInfo creates a new MessageForwardInfo
//
// @param origin Origin of a forwarded message
// @param date Point in time (Unix timestamp) when the message was originally sent
// @param publicServiceAnnouncementType The type of a public service announcement for the forwarded message
// @param fromChatId For messages forwarded to the chat with the current user (Saved Messages), to the Replies bot chat, or to the channel's discussion group, the identifier of the chat from which the message was forwarded last time; 0 if unknown
// @param fromMessageId For messages forwarded to the chat with the current user (Saved Messages), to the Replies bot chat, or to the channel's discussion group, the identifier of the original message from which the new message was forwarded last time; 0 if unknown
func NewMessageForwardInfo(origin MessageForwardOrigin, date int32, publicServiceAnnouncementType string, fromChatId int64, fromMessageId int64) *MessageForwardInfo {
	messageForwardInfoTemp := MessageForwardInfo{
		tdCommon:                      tdCommon{Type: "messageForwardInfo"},
		Origin:                        origin,
		Date:                          date,
		PublicServiceAnnouncementType: publicServiceAnnouncementType,
		FromChatId:                    fromChatId,
		FromMessageId:                 fromMessageId,
	}

	return &messageForwardInfoTemp
}

// UnmarshalJSON unmarshal to json
func (messageForwardInfo *MessageForwardInfo) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	tempObj := struct {
		tdCommon
		Date                          int32  `json:"date"`                             // Point in time (Unix timestamp) when the message was originally sent
		PublicServiceAnnouncementType string `json:"public_service_announcement_type"` // The type of a public service announcement for the forwarded message
		FromChatId                    int64  `json:"from_chat_id"`                     // For messages forwarded to the chat with the current user (Saved Messages), to the Replies bot chat, or to the channel's discussion group, the identifier of the chat from which the message was forwarded last time; 0 if unknown
		FromMessageId                 int64  `json:"from_message_id"`                  // For messages forwarded to the chat with the current user (Saved Messages), to the Replies bot chat, or to the channel's discussion group, the identifier of the original message from which the new message was forwarded last time; 0 if unknown
	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	messageForwardInfo.tdCommon = tempObj.tdCommon
	messageForwardInfo.Date = tempObj.Date
	messageForwardInfo.PublicServiceAnnouncementType = tempObj.PublicServiceAnnouncementType
	messageForwardInfo.FromChatId = tempObj.FromChatId
	messageForwardInfo.FromMessageId = tempObj.FromMessageId

	fieldOrigin, _ := unmarshalMessageForwardOrigin(objMap["origin"])
	messageForwardInfo.Origin = fieldOrigin

	return nil
}
