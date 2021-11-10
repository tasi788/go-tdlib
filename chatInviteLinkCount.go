package tdlib

// ChatInviteLinkCount Describes a chat administrator with a number of active and revoked chat invite links
type ChatInviteLinkCount struct {
	tdCommon
	UserId                 int64 `json:"user_id"`                   // Administrator's user identifier
	InviteLinkCount        int32 `json:"invite_link_count"`         // Number of active invite links
	RevokedInviteLinkCount int32 `json:"revoked_invite_link_count"` // Number of revoked invite links
}

// MessageType return the string telegram-type of ChatInviteLinkCount
func (chatInviteLinkCount *ChatInviteLinkCount) MessageType() string {
	return "chatInviteLinkCount"
}

// NewChatInviteLinkCount creates a new ChatInviteLinkCount
//
// @param userId Administrator's user identifier
// @param inviteLinkCount Number of active invite links
// @param revokedInviteLinkCount Number of revoked invite links
func NewChatInviteLinkCount(userId int64, inviteLinkCount int32, revokedInviteLinkCount int32) *ChatInviteLinkCount {
	chatInviteLinkCountTemp := ChatInviteLinkCount{
		tdCommon:               tdCommon{Type: "chatInviteLinkCount"},
		UserId:                 userId,
		InviteLinkCount:        inviteLinkCount,
		RevokedInviteLinkCount: revokedInviteLinkCount,
	}

	return &chatInviteLinkCountTemp
}
