package tdlib

import (
	"encoding/json"
	"fmt"
)

// ScopeNotificationSettings Contains information about notification settings for several chats
type ScopeNotificationSettings struct {
	tdCommon
	MuteFor                           int32  `json:"mute_for"`                             // Time left before notifications will be unmuted, in seconds
	Sound                             string `json:"sound"`                                // The name of an audio file to be used for notification sounds; only applies to iOS applications
	ShowPreview                       bool   `json:"show_preview"`                         // True, if message content must be displayed in notifications
	DisablePinnedMessageNotifications bool   `json:"disable_pinned_message_notifications"` // True, if notifications for incoming pinned messages will be created as for an ordinary unread message
	DisableMentionNotifications       bool   `json:"disable_mention_notifications"`        // True, if notifications for messages with mentions will be created as for an ordinary unread message
}

// MessageType return the string telegram-type of ScopeNotificationSettings
func (scopeNotificationSettings *ScopeNotificationSettings) MessageType() string {
	return "scopeNotificationSettings"
}

// NewScopeNotificationSettings creates a new ScopeNotificationSettings
//
// @param muteFor Time left before notifications will be unmuted, in seconds
// @param sound The name of an audio file to be used for notification sounds; only applies to iOS applications
// @param showPreview True, if message content must be displayed in notifications
// @param disablePinnedMessageNotifications True, if notifications for incoming pinned messages will be created as for an ordinary unread message
// @param disableMentionNotifications True, if notifications for messages with mentions will be created as for an ordinary unread message
func NewScopeNotificationSettings(muteFor int32, sound string, showPreview bool, disablePinnedMessageNotifications bool, disableMentionNotifications bool) *ScopeNotificationSettings {
	scopeNotificationSettingsTemp := ScopeNotificationSettings{
		tdCommon:                          tdCommon{Type: "scopeNotificationSettings"},
		MuteFor:                           muteFor,
		Sound:                             sound,
		ShowPreview:                       showPreview,
		DisablePinnedMessageNotifications: disablePinnedMessageNotifications,
		DisableMentionNotifications:       disableMentionNotifications,
	}

	return &scopeNotificationSettingsTemp
}

// GetScopeNotificationSettings Returns the notification settings for chats of a given type
// @param scope Types of chats for which to return the notification settings information
func (client *Client) GetScopeNotificationSettings(scope NotificationSettingsScope) (*ScopeNotificationSettings, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "getScopeNotificationSettings",
		"scope": scope,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var scopeNotificationSettings ScopeNotificationSettings
	err = json.Unmarshal(result.Raw, &scopeNotificationSettings)
	return &scopeNotificationSettings, err
}
