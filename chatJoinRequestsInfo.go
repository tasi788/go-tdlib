package tdlib

// ChatJoinRequestsInfo Contains information about pending chat join requests
type ChatJoinRequestsInfo struct {
	tdCommon
	TotalCount int32   `json:"total_count"` // Total number of pending join requests
	UserIds    []int64 `json:"user_ids"`    // Identifiers of users sent the newest pending join requests
}

// MessageType return the string telegram-type of ChatJoinRequestsInfo
func (chatJoinRequestsInfo *ChatJoinRequestsInfo) MessageType() string {
	return "chatJoinRequestsInfo"
}

// NewChatJoinRequestsInfo creates a new ChatJoinRequestsInfo
//
// @param totalCount Total number of pending join requests
// @param userIds Identifiers of users sent the newest pending join requests
func NewChatJoinRequestsInfo(totalCount int32, userIds []int64) *ChatJoinRequestsInfo {
	chatJoinRequestsInfoTemp := ChatJoinRequestsInfo{
		tdCommon:   tdCommon{Type: "chatJoinRequestsInfo"},
		TotalCount: totalCount,
		UserIds:    userIds,
	}

	return &chatJoinRequestsInfoTemp
}
