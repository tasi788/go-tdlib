package tdlib

import (
	"encoding/json"
	"fmt"
)

// ChatEvents Contains a list of chat events
type ChatEvents struct {
	tdCommon
	Events []ChatEvent `json:"events"` // List of events
}

// MessageType return the string telegram-type of ChatEvents
func (chatEvents *ChatEvents) MessageType() string {
	return "chatEvents"
}

// NewChatEvents creates a new ChatEvents
//
// @param events List of events
func NewChatEvents(events []ChatEvent) *ChatEvents {
	chatEventsTemp := ChatEvents{
		tdCommon: tdCommon{Type: "chatEvents"},
		Events:   events,
	}

	return &chatEventsTemp
}

// GetChatEventLog Returns a list of service actions taken by chat members and administrators in the last 48 hours. Available only for supergroups and channels. Requires administrator rights. Returns results in reverse chronological order (i. e., in order of decreasing event_id)
// @param chatId Chat identifier
// @param query Search query by which to filter events
// @param fromEventId Identifier of an event from which to return results. Use 0 to get results from the latest events
// @param limit The maximum number of events to return; up to 100
// @param filters The types of events to return; pass null to get chat events of all types
// @param userIds User identifiers by which to filter events. By default, events relating to all users will be returned
func (client *Client) GetChatEventLog(chatId int64, query string, fromEventId JSONInt64, limit int32, filters *ChatEventLogFilters, userIds []int64) (*ChatEvents, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":         "getChatEventLog",
		"chat_id":       chatId,
		"query":         query,
		"from_event_id": fromEventId,
		"limit":         limit,
		"filters":       filters,
		"user_ids":      userIds,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var chatEvents ChatEvents
	err = json.Unmarshal(result.Raw, &chatEvents)
	return &chatEvents, err
}
