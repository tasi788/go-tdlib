package tdlib

// ChatStatisticsAdministratorActionsInfo Contains statistics about administrator actions done by a user
type ChatStatisticsAdministratorActionsInfo struct {
	tdCommon
	UserId              int64 `json:"user_id"`               // Administrator user identifier
	DeletedMessageCount int32 `json:"deleted_message_count"` // Number of messages deleted by the administrator
	BannedUserCount     int32 `json:"banned_user_count"`     // Number of users banned by the administrator
	RestrictedUserCount int32 `json:"restricted_user_count"` // Number of users restricted by the administrator
}

// MessageType return the string telegram-type of ChatStatisticsAdministratorActionsInfo
func (chatStatisticsAdministratorActionsInfo *ChatStatisticsAdministratorActionsInfo) MessageType() string {
	return "chatStatisticsAdministratorActionsInfo"
}

// NewChatStatisticsAdministratorActionsInfo creates a new ChatStatisticsAdministratorActionsInfo
//
// @param userId Administrator user identifier
// @param deletedMessageCount Number of messages deleted by the administrator
// @param bannedUserCount Number of users banned by the administrator
// @param restrictedUserCount Number of users restricted by the administrator
func NewChatStatisticsAdministratorActionsInfo(userId int64, deletedMessageCount int32, bannedUserCount int32, restrictedUserCount int32) *ChatStatisticsAdministratorActionsInfo {
	chatStatisticsAdministratorActionsInfoTemp := ChatStatisticsAdministratorActionsInfo{
		tdCommon:            tdCommon{Type: "chatStatisticsAdministratorActionsInfo"},
		UserId:              userId,
		DeletedMessageCount: deletedMessageCount,
		BannedUserCount:     bannedUserCount,
		RestrictedUserCount: restrictedUserCount,
	}

	return &chatStatisticsAdministratorActionsInfoTemp
}
