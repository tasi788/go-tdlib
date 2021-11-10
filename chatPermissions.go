package tdlib

// ChatPermissions Describes actions that a user is allowed to take in a chat
type ChatPermissions struct {
	tdCommon
	CanSendMessages       bool `json:"can_send_messages"`         // True, if the user can send text messages, contacts, locations, and venues
	CanSendMediaMessages  bool `json:"can_send_media_messages"`   // True, if the user can send audio files, documents, photos, videos, video notes, and voice notes. Implies can_send_messages permissions
	CanSendPolls          bool `json:"can_send_polls"`            // True, if the user can send polls. Implies can_send_messages permissions
	CanSendStickers       bool `json:"can_send_stickers"`         // True, if the user can send stickers. Implies can_send_messages permissions
	CanSendAnimations     bool `json:"can_send_animations"`       // True, if the user can send animations. Implies can_send_messages permissions
	CanSendGames          bool `json:"can_send_games"`            // True, if the user can send games. Implies can_send_messages permissions
	CanUseInlineBots      bool `json:"can_use_inline_bots"`       // True, if the user can use inline bots. Implies can_send_messages permissions
	CanAddWebPagePreviews bool `json:"can_add_web_page_previews"` // True, if the user may add a web page preview to their messages. Implies can_send_messages permissions
	CanChangeInfo         bool `json:"can_change_info"`           // True, if the user can change the chat title, photo, and other settings
	CanInviteUsers        bool `json:"can_invite_users"`          // True, if the user can invite new users to the chat
	CanPinMessages        bool `json:"can_pin_messages"`          // True, if the user can pin messages
}

// MessageType return the string telegram-type of ChatPermissions
func (chatPermissions *ChatPermissions) MessageType() string {
	return "chatPermissions"
}

// NewChatPermissions creates a new ChatPermissions
//
// @param canSendMessages True, if the user can send text messages, contacts, locations, and venues
// @param canSendMediaMessages True, if the user can send audio files, documents, photos, videos, video notes, and voice notes. Implies can_send_messages permissions
// @param canSendPolls True, if the user can send polls. Implies can_send_messages permissions
// @param canSendStickers True, if the user can send stickers. Implies can_send_messages permissions
// @param canSendAnimations True, if the user can send animations. Implies can_send_messages permissions
// @param canSendGames True, if the user can send games. Implies can_send_messages permissions
// @param canUseInlineBots True, if the user can use inline bots. Implies can_send_messages permissions
// @param canAddWebPagePreviews True, if the user may add a web page preview to their messages. Implies can_send_messages permissions
// @param canChangeInfo True, if the user can change the chat title, photo, and other settings
// @param canInviteUsers True, if the user can invite new users to the chat
// @param canPinMessages True, if the user can pin messages
func NewChatPermissions(canSendMessages bool, canSendMediaMessages bool, canSendPolls bool, canSendStickers bool, canSendAnimations bool, canSendGames bool, canUseInlineBots bool, canAddWebPagePreviews bool, canChangeInfo bool, canInviteUsers bool, canPinMessages bool) *ChatPermissions {
	chatPermissionsTemp := ChatPermissions{
		tdCommon:              tdCommon{Type: "chatPermissions"},
		CanSendMessages:       canSendMessages,
		CanSendMediaMessages:  canSendMediaMessages,
		CanSendPolls:          canSendPolls,
		CanSendStickers:       canSendStickers,
		CanSendAnimations:     canSendAnimations,
		CanSendGames:          canSendGames,
		CanUseInlineBots:      canUseInlineBots,
		CanAddWebPagePreviews: canAddWebPagePreviews,
		CanChangeInfo:         canChangeInfo,
		CanInviteUsers:        canInviteUsers,
		CanPinMessages:        canPinMessages,
	}

	return &chatPermissionsTemp
}
