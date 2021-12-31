package tdlib

// ChatInviteLinkMember Describes a chat member joined a chat via an invite link
type ChatInviteLinkMember struct {
	tdCommon
	UserId         int64 `json:"user_id"`          // User identifier
	JoinedChatDate int32 `json:"joined_chat_date"` // Point in time (Unix timestamp) when the user joined the chat
	ApproverUserId int64 `json:"approver_user_id"` // User identifier of the chat administrator, approved user join request
}

// MessageType return the string telegram-type of ChatInviteLinkMember
func (chatInviteLinkMember *ChatInviteLinkMember) MessageType() string {
	return "chatInviteLinkMember"
}

// NewChatInviteLinkMember creates a new ChatInviteLinkMember
//
// @param userId User identifier
// @param joinedChatDate Point in time (Unix timestamp) when the user joined the chat
// @param approverUserId User identifier of the chat administrator, approved user join request
func NewChatInviteLinkMember(userId int64, joinedChatDate int32, approverUserId int64) *ChatInviteLinkMember {
	chatInviteLinkMemberTemp := ChatInviteLinkMember{
		tdCommon:       tdCommon{Type: "chatInviteLinkMember"},
		UserId:         userId,
		JoinedChatDate: joinedChatDate,
		ApproverUserId: approverUserId,
	}

	return &chatInviteLinkMemberTemp
}
