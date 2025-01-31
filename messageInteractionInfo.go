package tdlib

// MessageInteractionInfo Contains information about interactions with a message
type MessageInteractionInfo struct {
	tdCommon
	ViewCount    int32             `json:"view_count"`    // Number of times the message was viewed
	ForwardCount int32             `json:"forward_count"` // Number of times the message was forwarded
	ReplyInfo    *MessageReplyInfo `json:"reply_info"`    // Information about direct or indirect replies to the message; may be null. Currently, available only in channels with a discussion supergroup and discussion supergroups for messages, which are not replies itself
}

// MessageType return the string telegram-type of MessageInteractionInfo
func (messageInteractionInfo *MessageInteractionInfo) MessageType() string {
	return "messageInteractionInfo"
}

// NewMessageInteractionInfo creates a new MessageInteractionInfo
//
// @param viewCount Number of times the message was viewed
// @param forwardCount Number of times the message was forwarded
// @param replyInfo Information about direct or indirect replies to the message; may be null. Currently, available only in channels with a discussion supergroup and discussion supergroups for messages, which are not replies itself
func NewMessageInteractionInfo(viewCount int32, forwardCount int32, replyInfo *MessageReplyInfo) *MessageInteractionInfo {
	messageInteractionInfoTemp := MessageInteractionInfo{
		tdCommon:     tdCommon{Type: "messageInteractionInfo"},
		ViewCount:    viewCount,
		ForwardCount: forwardCount,
		ReplyInfo:    replyInfo,
	}

	return &messageInteractionInfoTemp
}
