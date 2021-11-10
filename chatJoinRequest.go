package tdlib

// ChatJoinRequest Describes a user that sent a join request and waits for administrator approval
type ChatJoinRequest struct {
	tdCommon
	UserId int64  `json:"user_id"` // User identifier
	Date   int32  `json:"date"`    // Point in time (Unix timestamp) when the user sent the join request
	Bio    string `json:"bio"`     // A short bio of the user
}

// MessageType return the string telegram-type of ChatJoinRequest
func (chatJoinRequest *ChatJoinRequest) MessageType() string {
	return "chatJoinRequest"
}

// NewChatJoinRequest creates a new ChatJoinRequest
//
// @param userId User identifier
// @param date Point in time (Unix timestamp) when the user sent the join request
// @param bio A short bio of the user
func NewChatJoinRequest(userId int64, date int32, bio string) *ChatJoinRequest {
	chatJoinRequestTemp := ChatJoinRequest{
		tdCommon: tdCommon{Type: "chatJoinRequest"},
		UserId:   userId,
		Date:     date,
		Bio:      bio,
	}

	return &chatJoinRequestTemp
}
