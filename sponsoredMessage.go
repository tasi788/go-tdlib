package tdlib

import (
	"encoding/json"
	"fmt"
)

// SponsoredMessage Describes a sponsored message
type SponsoredMessage struct {
	tdCommon
	MessageId       int64               `json:"message_id"`        // Message identifier; unique for the chat to which the sponsored message belongs among both ordinary and sponsored messages
	SponsorChatId   int64               `json:"sponsor_chat_id"`   // Sponsor chat identifier; 0 if the sponsor chat is accessible through an invite link
	SponsorChatInfo *ChatInviteLinkInfo `json:"sponsor_chat_info"` // Information about the sponsor chat; may be null unless sponsor_chat_id == 0
	Link            InternalLinkType    `json:"link"`              // An internal link to be opened when the sponsored message is clicked; may be null. If null, the sponsor chat needs to be opened instead
	Content         MessageContent      `json:"content"`           // Content of the message. Currently, can be only of the type messageText
}

// MessageType return the string telegram-type of SponsoredMessage
func (sponsoredMessage *SponsoredMessage) MessageType() string {
	return "sponsoredMessage"
}

// NewSponsoredMessage creates a new SponsoredMessage
//
// @param messageId Message identifier; unique for the chat to which the sponsored message belongs among both ordinary and sponsored messages
// @param sponsorChatId Sponsor chat identifier; 0 if the sponsor chat is accessible through an invite link
// @param sponsorChatInfo Information about the sponsor chat; may be null unless sponsor_chat_id == 0
// @param link An internal link to be opened when the sponsored message is clicked; may be null. If null, the sponsor chat needs to be opened instead
// @param content Content of the message. Currently, can be only of the type messageText
func NewSponsoredMessage(messageId int64, sponsorChatId int64, sponsorChatInfo *ChatInviteLinkInfo, link InternalLinkType, content MessageContent) *SponsoredMessage {
	sponsoredMessageTemp := SponsoredMessage{
		tdCommon:        tdCommon{Type: "sponsoredMessage"},
		MessageId:       messageId,
		SponsorChatId:   sponsorChatId,
		SponsorChatInfo: sponsorChatInfo,
		Link:            link,
		Content:         content,
	}

	return &sponsoredMessageTemp
}

// UnmarshalJSON unmarshal to json
func (sponsoredMessage *SponsoredMessage) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	tempObj := struct {
		tdCommon
		MessageId       int64               `json:"message_id"`        // Message identifier; unique for the chat to which the sponsored message belongs among both ordinary and sponsored messages
		SponsorChatId   int64               `json:"sponsor_chat_id"`   // Sponsor chat identifier; 0 if the sponsor chat is accessible through an invite link
		SponsorChatInfo *ChatInviteLinkInfo `json:"sponsor_chat_info"` // Information about the sponsor chat; may be null unless sponsor_chat_id == 0

	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	sponsoredMessage.tdCommon = tempObj.tdCommon
	sponsoredMessage.MessageId = tempObj.MessageId
	sponsoredMessage.SponsorChatId = tempObj.SponsorChatId
	sponsoredMessage.SponsorChatInfo = tempObj.SponsorChatInfo

	fieldLink, _ := unmarshalInternalLinkType(objMap["link"])
	sponsoredMessage.Link = fieldLink

	fieldContent, _ := unmarshalMessageContent(objMap["content"])
	sponsoredMessage.Content = fieldContent

	return nil
}

// GetChatSponsoredMessage Returns sponsored message to be shown in a chat; for channel chats only. Returns a 404 error if there is no sponsored message in the chat
// @param chatId Identifier of the chat
func (client *Client) GetChatSponsoredMessage(chatId int64) (*SponsoredMessage, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "getChatSponsoredMessage",
		"chat_id": chatId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var sponsoredMessage SponsoredMessage
	err = json.Unmarshal(result.Raw, &sponsoredMessage)
	return &sponsoredMessage, err
}
