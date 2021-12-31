package tdlib

import (
	"encoding/json"
	"fmt"
)

// MessageCalendar Contains information about found messages, split by days according to the option "utc_time_offset"
type MessageCalendar struct {
	tdCommon
	TotalCount int32                `json:"total_count"` // Total number of found messages
	Days       []MessageCalendarDay `json:"days"`        // Information about messages sent
}

// MessageType return the string telegram-type of MessageCalendar
func (messageCalendar *MessageCalendar) MessageType() string {
	return "messageCalendar"
}

// NewMessageCalendar creates a new MessageCalendar
//
// @param totalCount Total number of found messages
// @param days Information about messages sent
func NewMessageCalendar(totalCount int32, days []MessageCalendarDay) *MessageCalendar {
	messageCalendarTemp := MessageCalendar{
		tdCommon:   tdCommon{Type: "messageCalendar"},
		TotalCount: totalCount,
		Days:       days,
	}

	return &messageCalendarTemp
}

// GetChatMessageCalendar Returns information about the next messages of the specified type in the chat split by days. Returns the results in reverse chronological order. Can return partial result for the last returned day. Behavior of this method depends on the value of the option "utc_time_offset"
// @param chatId Identifier of the chat in which to return information about messages
// @param filter Filter for message content. Filters searchMessagesFilterEmpty, searchMessagesFilterMention and searchMessagesFilterUnreadMention are unsupported in this function
// @param fromMessageId The message identifier from which to return information about messages; use 0 to get results from the last message
func (client *Client) GetChatMessageCalendar(chatId int64, filter SearchMessagesFilter, fromMessageId int64) (*MessageCalendar, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":           "getChatMessageCalendar",
		"chat_id":         chatId,
		"filter":          filter,
		"from_message_id": fromMessageId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var messageCalendar MessageCalendar
	err = json.Unmarshal(result.Raw, &messageCalendar)
	return &messageCalendar, err
}
