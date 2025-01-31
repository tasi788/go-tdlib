package tdlib

// ChatStatisticsInviterInfo Contains statistics about number of new members invited by a user
type ChatStatisticsInviterInfo struct {
	tdCommon
	UserId           int64 `json:"user_id"`            // User identifier
	AddedMemberCount int32 `json:"added_member_count"` // Number of new members invited by the user
}

// MessageType return the string telegram-type of ChatStatisticsInviterInfo
func (chatStatisticsInviterInfo *ChatStatisticsInviterInfo) MessageType() string {
	return "chatStatisticsInviterInfo"
}

// NewChatStatisticsInviterInfo creates a new ChatStatisticsInviterInfo
//
// @param userId User identifier
// @param addedMemberCount Number of new members invited by the user
func NewChatStatisticsInviterInfo(userId int64, addedMemberCount int32) *ChatStatisticsInviterInfo {
	chatStatisticsInviterInfoTemp := ChatStatisticsInviterInfo{
		tdCommon:         tdCommon{Type: "chatStatisticsInviterInfo"},
		UserId:           userId,
		AddedMemberCount: addedMemberCount,
	}

	return &chatStatisticsInviterInfoTemp
}
