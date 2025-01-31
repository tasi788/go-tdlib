package tdlib

// ChatTheme Describes a chat theme
type ChatTheme struct {
	tdCommon
	Name          string         `json:"name"`           // Theme name
	LightSettings *ThemeSettings `json:"light_settings"` // Theme settings for a light chat theme
	DarkSettings  *ThemeSettings `json:"dark_settings"`  // Theme settings for a dark chat theme
}

// MessageType return the string telegram-type of ChatTheme
func (chatTheme *ChatTheme) MessageType() string {
	return "chatTheme"
}

// NewChatTheme creates a new ChatTheme
//
// @param name Theme name
// @param lightSettings Theme settings for a light chat theme
// @param darkSettings Theme settings for a dark chat theme
func NewChatTheme(name string, lightSettings *ThemeSettings, darkSettings *ThemeSettings) *ChatTheme {
	chatThemeTemp := ChatTheme{
		tdCommon:      tdCommon{Type: "chatTheme"},
		Name:          name,
		LightSettings: lightSettings,
		DarkSettings:  darkSettings,
	}

	return &chatThemeTemp
}
