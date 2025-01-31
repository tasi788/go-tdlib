package tdlib

// BotCommand Represents a command supported by a bot
type BotCommand struct {
	tdCommon
	Command     string `json:"command"`     // Text of the bot command
	Description string `json:"description"` // Description of the bot command
}

// MessageType return the string telegram-type of BotCommand
func (botCommand *BotCommand) MessageType() string {
	return "botCommand"
}

// NewBotCommand creates a new BotCommand
//
// @param command Text of the bot command
// @param description Description of the bot command
func NewBotCommand(command string, description string) *BotCommand {
	botCommandTemp := BotCommand{
		tdCommon:    tdCommon{Type: "botCommand"},
		Command:     command,
		Description: description,
	}

	return &botCommandTemp
}
