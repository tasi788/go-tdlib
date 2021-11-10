package tdlib

// ChatStatisticsMessageInteractionInfo Contains statistics about interactions with a message
type ChatStatisticsMessageInteractionInfo struct {
	tdCommon
	MessageId    int64 `json:"message_id"`    // Message identifier
	ViewCount    int32 `json:"view_count"`    // Number of times the message was viewed
	ForwardCount int32 `json:"forward_count"` // Number of times the message was forwarded
}

// MessageType return the string telegram-type of ChatStatisticsMessageInteractionInfo
func (chatStatisticsMessageInteractionInfo *ChatStatisticsMessageInteractionInfo) MessageType() string {
	return "chatStatisticsMessageInteractionInfo"
}

// NewChatStatisticsMessageInteractionInfo creates a new ChatStatisticsMessageInteractionInfo
//
// @param messageId Message identifier
// @param viewCount Number of times the message was viewed
// @param forwardCount Number of times the message was forwarded
func NewChatStatisticsMessageInteractionInfo(messageId int64, viewCount int32, forwardCount int32) *ChatStatisticsMessageInteractionInfo {
	chatStatisticsMessageInteractionInfoTemp := ChatStatisticsMessageInteractionInfo{
		tdCommon:     tdCommon{Type: "chatStatisticsMessageInteractionInfo"},
		MessageId:    messageId,
		ViewCount:    viewCount,
		ForwardCount: forwardCount,
	}

	return &chatStatisticsMessageInteractionInfoTemp
}
