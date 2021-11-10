package tdlib

// MessageCalendarDay Contains information about found messages sent in a specific day
type MessageCalendarDay struct {
	tdCommon
	TotalCount int32    `json:"total_count"` // Total number of found messages sent in the day
	Message    *Message `json:"message"`     // First message sent in the day
}

// MessageType return the string telegram-type of MessageCalendarDay
func (messageCalendarDay *MessageCalendarDay) MessageType() string {
	return "messageCalendarDay"
}

// NewMessageCalendarDay creates a new MessageCalendarDay
//
// @param totalCount Total number of found messages sent in the day
// @param message First message sent in the day
func NewMessageCalendarDay(totalCount int32, message *Message) *MessageCalendarDay {
	messageCalendarDayTemp := MessageCalendarDay{
		tdCommon:   tdCommon{Type: "messageCalendarDay"},
		TotalCount: totalCount,
		Message:    message,
	}

	return &messageCalendarDayTemp
}
