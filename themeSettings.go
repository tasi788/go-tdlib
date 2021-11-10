package tdlib

import (
	"encoding/json"
)

// ThemeSettings Describes theme settings
type ThemeSettings struct {
	tdCommon
	AccentColor                int32          `json:"accent_color"`                  // Theme accent color in ARGB format
	Background                 *Background    `json:"background"`                    // The background to be used in chats; may be null
	OutgoingMessageFill        BackgroundFill `json:"outgoing_message_fill"`         // The fill to be used as a background for outgoing messages
	AnimateOutgoingMessageFill bool           `json:"animate_outgoing_message_fill"` // If true, the freeform gradient fill needs to be animated on every sent message
	OutgoingMessageAccentColor int32          `json:"outgoing_message_accent_color"` // Accent color of outgoing messages in ARGB format
}

// MessageType return the string telegram-type of ThemeSettings
func (themeSettings *ThemeSettings) MessageType() string {
	return "themeSettings"
}

// NewThemeSettings creates a new ThemeSettings
//
// @param accentColor Theme accent color in ARGB format
// @param background The background to be used in chats; may be null
// @param outgoingMessageFill The fill to be used as a background for outgoing messages
// @param animateOutgoingMessageFill If true, the freeform gradient fill needs to be animated on every sent message
// @param outgoingMessageAccentColor Accent color of outgoing messages in ARGB format
func NewThemeSettings(accentColor int32, background *Background, outgoingMessageFill BackgroundFill, animateOutgoingMessageFill bool, outgoingMessageAccentColor int32) *ThemeSettings {
	themeSettingsTemp := ThemeSettings{
		tdCommon:                   tdCommon{Type: "themeSettings"},
		AccentColor:                accentColor,
		Background:                 background,
		OutgoingMessageFill:        outgoingMessageFill,
		AnimateOutgoingMessageFill: animateOutgoingMessageFill,
		OutgoingMessageAccentColor: outgoingMessageAccentColor,
	}

	return &themeSettingsTemp
}

// UnmarshalJSON unmarshal to json
func (themeSettings *ThemeSettings) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	tempObj := struct {
		tdCommon
		AccentColor                int32       `json:"accent_color"`                  // Theme accent color in ARGB format
		Background                 *Background `json:"background"`                    // The background to be used in chats; may be null
		AnimateOutgoingMessageFill bool        `json:"animate_outgoing_message_fill"` // If true, the freeform gradient fill needs to be animated on every sent message
		OutgoingMessageAccentColor int32       `json:"outgoing_message_accent_color"` // Accent color of outgoing messages in ARGB format
	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	themeSettings.tdCommon = tempObj.tdCommon
	themeSettings.AccentColor = tempObj.AccentColor
	themeSettings.Background = tempObj.Background
	themeSettings.AnimateOutgoingMessageFill = tempObj.AnimateOutgoingMessageFill
	themeSettings.OutgoingMessageAccentColor = tempObj.OutgoingMessageAccentColor

	fieldOutgoingMessageFill, _ := unmarshalBackgroundFill(objMap["outgoing_message_fill"])
	themeSettings.OutgoingMessageFill = fieldOutgoingMessageFill

	return nil
}
