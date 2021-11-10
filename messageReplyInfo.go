package tdlib

// MessageReplyInfo Contains information about replies to a message
type MessageReplyInfo struct {
	tdCommon
	ReplyCount              int32           `json:"reply_count"`                 // Number of times the message was directly or indirectly replied
	RecentRepliers          []MessageSender `json:"recent_repliers"`             // Recent repliers to the message; available in channels with a discussion supergroup
	LastReadInboxMessageId  int64           `json:"last_read_inbox_message_id"`  // Identifier of the last read incoming reply to the message
	LastReadOutboxMessageId int64           `json:"last_read_outbox_message_id"` // Identifier of the last read outgoing reply to the message
	LastMessageId           int64           `json:"last_message_id"`             // Identifier of the last reply to the message
}

// MessageType return the string telegram-type of MessageReplyInfo
func (messageReplyInfo *MessageReplyInfo) MessageType() string {
	return "messageReplyInfo"
}

// NewMessageReplyInfo creates a new MessageReplyInfo
//
// @param replyCount Number of times the message was directly or indirectly replied
// @param recentRepliers Recent repliers to the message; available in channels with a discussion supergroup
// @param lastReadInboxMessageId Identifier of the last read incoming reply to the message
// @param lastReadOutboxMessageId Identifier of the last read outgoing reply to the message
// @param lastMessageId Identifier of the last reply to the message
func NewMessageReplyInfo(replyCount int32, recentRepliers []MessageSender, lastReadInboxMessageId int64, lastReadOutboxMessageId int64, lastMessageId int64) *MessageReplyInfo {
	messageReplyInfoTemp := MessageReplyInfo{
		tdCommon:                tdCommon{Type: "messageReplyInfo"},
		ReplyCount:              replyCount,
		RecentRepliers:          recentRepliers,
		LastReadInboxMessageId:  lastReadInboxMessageId,
		LastReadOutboxMessageId: lastReadOutboxMessageId,
		LastMessageId:           lastMessageId,
	}

	return &messageReplyInfoTemp
}
