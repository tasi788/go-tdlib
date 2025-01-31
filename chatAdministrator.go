package tdlib

// ChatAdministrator Contains information about a chat administrator
type ChatAdministrator struct {
	tdCommon
	UserId      int64  `json:"user_id"`      // User identifier of the administrator
	CustomTitle string `json:"custom_title"` // Custom title of the administrator
	IsOwner     bool   `json:"is_owner"`     // True, if the user is the owner of the chat
}

// MessageType return the string telegram-type of ChatAdministrator
func (chatAdministrator *ChatAdministrator) MessageType() string {
	return "chatAdministrator"
}

// NewChatAdministrator creates a new ChatAdministrator
//
// @param userId User identifier of the administrator
// @param customTitle Custom title of the administrator
// @param isOwner True, if the user is the owner of the chat
func NewChatAdministrator(userId int64, customTitle string, isOwner bool) *ChatAdministrator {
	chatAdministratorTemp := ChatAdministrator{
		tdCommon:    tdCommon{Type: "chatAdministrator"},
		UserId:      userId,
		CustomTitle: customTitle,
		IsOwner:     isOwner,
	}

	return &chatAdministratorTemp
}
