package tdlib

import (
	"encoding/json"
	"fmt"
)

// BotCommands Contains a list of bot commands
type BotCommands struct {
	tdCommon
	BotUserId int64        `json:"bot_user_id"` // Bot's user identifier
	Commands  []BotCommand `json:"commands"`    // List of bot commands
}

// MessageType return the string telegram-type of BotCommands
func (botCommands *BotCommands) MessageType() string {
	return "botCommands"
}

// NewBotCommands creates a new BotCommands
//
// @param botUserId Bot's user identifier
// @param commands List of bot commands
func NewBotCommands(botUserId int64, commands []BotCommand) *BotCommands {
	botCommandsTemp := BotCommands{
		tdCommon:  tdCommon{Type: "botCommands"},
		BotUserId: botUserId,
		Commands:  commands,
	}

	return &botCommandsTemp
}

// GetCommands Returns the list of commands supported by the bot for the given user scope and language; for bots only
// @param scope The scope to which the commands are relevant; pass null to get commands in the default bot command scope
// @param languageCode A two-letter ISO 639-1 country code or an empty string
func (client *Client) GetCommands(scope BotCommandScope, languageCode string) (*BotCommands, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":         "getCommands",
		"scope":         scope,
		"language_code": languageCode,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var botCommands BotCommands
	err = json.Unmarshal(result.Raw, &botCommands)
	return &botCommands, err
}
