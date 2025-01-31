package tdlib

// ChatNotificationSettings Contains information about notification settings for a chat
type ChatNotificationSettings struct {
	tdCommon
	UseDefaultMuteFor                           bool   `json:"use_default_mute_for"`                             // If true, mute_for is ignored and the value for the relevant type of chat is used instead
	MuteFor                                     int32  `json:"mute_for"`                                         // Time left before notifications will be unmuted, in seconds
	UseDefaultSound                             bool   `json:"use_default_sound"`                                // If true, sound is ignored and the value for the relevant type of chat is used instead
	Sound                                       string `json:"sound"`                                            // The name of an audio file to be used for notification sounds; only applies to iOS applications
	UseDefaultShowPreview                       bool   `json:"use_default_show_preview"`                         // If true, show_preview is ignored and the value for the relevant type of chat is used instead
	ShowPreview                                 bool   `json:"show_preview"`                                     // True, if message content must be displayed in notifications
	UseDefaultDisablePinnedMessageNotifications bool   `json:"use_default_disable_pinned_message_notifications"` // If true, disable_pinned_message_notifications is ignored and the value for the relevant type of chat is used instead
	DisablePinnedMessageNotifications           bool   `json:"disable_pinned_message_notifications"`             // If true, notifications for incoming pinned messages will be created as for an ordinary unread message
	UseDefaultDisableMentionNotifications       bool   `json:"use_default_disable_mention_notifications"`        // If true, disable_mention_notifications is ignored and the value for the relevant type of chat is used instead
	DisableMentionNotifications                 bool   `json:"disable_mention_notifications"`                    // If true, notifications for messages with mentions will be created as for an ordinary unread message
}

// MessageType return the string telegram-type of ChatNotificationSettings
func (chatNotificationSettings *ChatNotificationSettings) MessageType() string {
	return "chatNotificationSettings"
}

// NewChatNotificationSettings creates a new ChatNotificationSettings
//
// @param useDefaultMuteFor If true, mute_for is ignored and the value for the relevant type of chat is used instead
// @param muteFor Time left before notifications will be unmuted, in seconds
// @param useDefaultSound If true, sound is ignored and the value for the relevant type of chat is used instead
// @param sound The name of an audio file to be used for notification sounds; only applies to iOS applications
// @param useDefaultShowPreview If true, show_preview is ignored and the value for the relevant type of chat is used instead
// @param showPreview True, if message content must be displayed in notifications
// @param useDefaultDisablePinnedMessageNotifications If true, disable_pinned_message_notifications is ignored and the value for the relevant type of chat is used instead
// @param disablePinnedMessageNotifications If true, notifications for incoming pinned messages will be created as for an ordinary unread message
// @param useDefaultDisableMentionNotifications If true, disable_mention_notifications is ignored and the value for the relevant type of chat is used instead
// @param disableMentionNotifications If true, notifications for messages with mentions will be created as for an ordinary unread message
func NewChatNotificationSettings(useDefaultMuteFor bool, muteFor int32, useDefaultSound bool, sound string, useDefaultShowPreview bool, showPreview bool, useDefaultDisablePinnedMessageNotifications bool, disablePinnedMessageNotifications bool, useDefaultDisableMentionNotifications bool, disableMentionNotifications bool) *ChatNotificationSettings {
	chatNotificationSettingsTemp := ChatNotificationSettings{
		tdCommon:              tdCommon{Type: "chatNotificationSettings"},
		UseDefaultMuteFor:     useDefaultMuteFor,
		MuteFor:               muteFor,
		UseDefaultSound:       useDefaultSound,
		Sound:                 sound,
		UseDefaultShowPreview: useDefaultShowPreview,
		ShowPreview:           showPreview,
		UseDefaultDisablePinnedMessageNotifications: useDefaultDisablePinnedMessageNotifications,
		DisablePinnedMessageNotifications:           disablePinnedMessageNotifications,
		UseDefaultDisableMentionNotifications:       useDefaultDisableMentionNotifications,
		DisableMentionNotifications:                 disableMentionNotifications,
	}

	return &chatNotificationSettingsTemp
}
