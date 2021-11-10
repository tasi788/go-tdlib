package tdlib

// MessagePosition Contains information about a message in a specific position
type MessagePosition struct {
	tdCommon
	Position  int32 `json:"position"`   // 0-based message position in the full list of suitable messages
	MessageId int64 `json:"message_id"` // Message identifier
	Date      int32 `json:"date"`       // Point in time (Unix timestamp) when the message was sent
}

// MessageType return the string telegram-type of MessagePosition
func (messagePosition *MessagePosition) MessageType() string {
	return "messagePosition"
}

// NewMessagePosition creates a new MessagePosition
//
// @param position 0-based message position in the full list of suitable messages
// @param messageId Message identifier
// @param date Point in time (Unix timestamp) when the message was sent
func NewMessagePosition(position int32, messageId int64, date int32) *MessagePosition {
	messagePositionTemp := MessagePosition{
		tdCommon:  tdCommon{Type: "messagePosition"},
		Position:  position,
		MessageId: messageId,
		Date:      date,
	}

	return &messagePositionTemp
}
