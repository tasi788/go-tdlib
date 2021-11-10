package tdlib

import (
	"encoding/json"
)

// DraftMessage Contains information about a message draft
type DraftMessage struct {
	tdCommon
	ReplyToMessageId int64               `json:"reply_to_message_id"` // Identifier of the message to reply to; 0 if none
	Date             int32               `json:"date"`                // Point in time (Unix timestamp) when the draft was created
	InputMessageText InputMessageContent `json:"input_message_text"`  // Content of the message draft; must be of the type inputMessageText
}

// MessageType return the string telegram-type of DraftMessage
func (draftMessage *DraftMessage) MessageType() string {
	return "draftMessage"
}

// NewDraftMessage creates a new DraftMessage
//
// @param replyToMessageId Identifier of the message to reply to; 0 if none
// @param date Point in time (Unix timestamp) when the draft was created
// @param inputMessageText Content of the message draft; must be of the type inputMessageText
func NewDraftMessage(replyToMessageId int64, date int32, inputMessageText InputMessageContent) *DraftMessage {
	draftMessageTemp := DraftMessage{
		tdCommon:         tdCommon{Type: "draftMessage"},
		ReplyToMessageId: replyToMessageId,
		Date:             date,
		InputMessageText: inputMessageText,
	}

	return &draftMessageTemp
}

// UnmarshalJSON unmarshal to json
func (draftMessage *DraftMessage) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	tempObj := struct {
		tdCommon
		ReplyToMessageId int64 `json:"reply_to_message_id"` // Identifier of the message to reply to; 0 if none
		Date             int32 `json:"date"`                // Point in time (Unix timestamp) when the draft was created

	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	draftMessage.tdCommon = tempObj.tdCommon
	draftMessage.ReplyToMessageId = tempObj.ReplyToMessageId
	draftMessage.Date = tempObj.Date

	fieldInputMessageText, _ := unmarshalInputMessageContent(objMap["input_message_text"])
	draftMessage.InputMessageText = fieldInputMessageText

	return nil
}
