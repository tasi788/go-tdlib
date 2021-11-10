package tdlib

import (
	"encoding/json"
)

// SponsoredMessage Describes a sponsored message
type SponsoredMessage struct {
	tdCommon
	Id            int32            `json:"id"`              // Unique sponsored message identifier
	SponsorChatId int64            `json:"sponsor_chat_id"` // Chat identifier
	Link          InternalLinkType `json:"link"`            // An internal link to be opened when the sponsored message is clicked; may be null. If null, the sponsor chat needs to be opened instead
	Content       MessageContent   `json:"content"`         // Content of the message
}

// MessageType return the string telegram-type of SponsoredMessage
func (sponsoredMessage *SponsoredMessage) MessageType() string {
	return "sponsoredMessage"
}

// NewSponsoredMessage creates a new SponsoredMessage
//
// @param id Unique sponsored message identifier
// @param sponsorChatId Chat identifier
// @param link An internal link to be opened when the sponsored message is clicked; may be null. If null, the sponsor chat needs to be opened instead
// @param content Content of the message
func NewSponsoredMessage(id int32, sponsorChatId int64, link InternalLinkType, content MessageContent) *SponsoredMessage {
	sponsoredMessageTemp := SponsoredMessage{
		tdCommon:      tdCommon{Type: "sponsoredMessage"},
		Id:            id,
		SponsorChatId: sponsorChatId,
		Link:          link,
		Content:       content,
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
		Id            int32 `json:"id"`              // Unique sponsored message identifier
		SponsorChatId int64 `json:"sponsor_chat_id"` // Chat identifier

	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	sponsoredMessage.tdCommon = tempObj.tdCommon
	sponsoredMessage.Id = tempObj.Id
	sponsoredMessage.SponsorChatId = tempObj.SponsorChatId

	fieldLink, _ := unmarshalInternalLinkType(objMap["link"])
	sponsoredMessage.Link = fieldLink

	fieldContent, _ := unmarshalMessageContent(objMap["content"])
	sponsoredMessage.Content = fieldContent

	return nil
}
