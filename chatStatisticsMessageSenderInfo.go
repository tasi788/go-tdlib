package tdlib

// ChatStatisticsMessageSenderInfo Contains statistics about messages sent by a user
type ChatStatisticsMessageSenderInfo struct {
	tdCommon
	UserId                int64 `json:"user_id"`                 // User identifier
	SentMessageCount      int32 `json:"sent_message_count"`      // Number of sent messages
	AverageCharacterCount int32 `json:"average_character_count"` // Average number of characters in sent messages; 0 if unknown
}

// MessageType return the string telegram-type of ChatStatisticsMessageSenderInfo
func (chatStatisticsMessageSenderInfo *ChatStatisticsMessageSenderInfo) MessageType() string {
	return "chatStatisticsMessageSenderInfo"
}

// NewChatStatisticsMessageSenderInfo creates a new ChatStatisticsMessageSenderInfo
//
// @param userId User identifier
// @param sentMessageCount Number of sent messages
// @param averageCharacterCount Average number of characters in sent messages; 0 if unknown
func NewChatStatisticsMessageSenderInfo(userId int64, sentMessageCount int32, averageCharacterCount int32) *ChatStatisticsMessageSenderInfo {
	chatStatisticsMessageSenderInfoTemp := ChatStatisticsMessageSenderInfo{
		tdCommon:              tdCommon{Type: "chatStatisticsMessageSenderInfo"},
		UserId:                userId,
		SentMessageCount:      sentMessageCount,
		AverageCharacterCount: averageCharacterCount,
	}

	return &chatStatisticsMessageSenderInfoTemp
}
