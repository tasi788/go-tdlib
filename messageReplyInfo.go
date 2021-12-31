package tdlib

// MessageReplyInfo Contains information about replies to a message
type MessageReplyInfo struct {
	tdCommon
	ReplyCount              int32           `json:"reply_count"`                 // Number of times the message was directly or indirectly replied
	RecentReplierIds        []MessageSender `json:"recent_replier_ids"`          // Identifiers of at most 3 recent repliers to the message; available in channels with a discussion supergroup. The users and chats are expected to be inaccessible: only their photo and name will be available
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
// @param recentReplierIds Identifiers of at most 3 recent repliers to the message; available in channels with a discussion supergroup. The users and chats are expected to be inaccessible: only their photo and name will be available
// @param lastReadInboxMessageId Identifier of the last read incoming reply to the message
// @param lastReadOutboxMessageId Identifier of the last read outgoing reply to the message
// @param lastMessageId Identifier of the last reply to the message
func NewMessageReplyInfo(replyCount int32, recentReplierIds []MessageSender, lastReadInboxMessageId int64, lastReadOutboxMessageId int64, lastMessageId int64) *MessageReplyInfo {
	messageReplyInfoTemp := MessageReplyInfo{
		tdCommon:                tdCommon{Type: "messageReplyInfo"},
		ReplyCount:              replyCount,
		RecentReplierIds:        recentReplierIds,
		LastReadInboxMessageId:  lastReadInboxMessageId,
		LastReadOutboxMessageId: lastReadOutboxMessageId,
		LastMessageId:           lastMessageId,
	}

	return &messageReplyInfoTemp
}
