package tdlib

// ChatNearby Describes a chat located nearby
type ChatNearby struct {
	tdCommon
	ChatId   int64 `json:"chat_id"`  // Chat identifier
	Distance int32 `json:"distance"` // Distance to the chat location, in meters
}

// MessageType return the string telegram-type of ChatNearby
func (chatNearby *ChatNearby) MessageType() string {
	return "chatNearby"
}

// NewChatNearby creates a new ChatNearby
//
// @param chatId Chat identifier
// @param distance Distance to the chat location, in meters
func NewChatNearby(chatId int64, distance int32) *ChatNearby {
	chatNearbyTemp := ChatNearby{
		tdCommon: tdCommon{Type: "chatNearby"},
		ChatId:   chatId,
		Distance: distance,
	}

	return &chatNearbyTemp
}
