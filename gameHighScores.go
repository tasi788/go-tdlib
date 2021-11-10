package tdlib

import (
	"encoding/json"
	"fmt"
)

// GameHighScores Contains a list of game high scores
type GameHighScores struct {
	tdCommon
	Scores []GameHighScore `json:"scores"` // A list of game high scores
}

// MessageType return the string telegram-type of GameHighScores
func (gameHighScores *GameHighScores) MessageType() string {
	return "gameHighScores"
}

// NewGameHighScores creates a new GameHighScores
//
// @param scores A list of game high scores
func NewGameHighScores(scores []GameHighScore) *GameHighScores {
	gameHighScoresTemp := GameHighScores{
		tdCommon: tdCommon{Type: "gameHighScores"},
		Scores:   scores,
	}

	return &gameHighScoresTemp
}

// GetGameHighScores Returns the high scores for a game and some part of the high score table in the range of the specified user; for bots only
// @param chatId The chat that contains the message with the game
// @param messageId Identifier of the message
// @param userId User identifier
func (client *Client) GetGameHighScores(chatId int64, messageId int64, userId int64) (*GameHighScores, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":      "getGameHighScores",
		"chat_id":    chatId,
		"message_id": messageId,
		"user_id":    userId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var gameHighScores GameHighScores
	err = json.Unmarshal(result.Raw, &gameHighScores)
	return &gameHighScores, err
}

// GetInlineGameHighScores Returns game high scores and some part of the high score table in the range of the specified user; for bots only
// @param inlineMessageId Inline message identifier
// @param userId User identifier
func (client *Client) GetInlineGameHighScores(inlineMessageId string, userId int64) (*GameHighScores, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":             "getInlineGameHighScores",
		"inline_message_id": inlineMessageId,
		"user_id":           userId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var gameHighScores GameHighScores
	err = json.Unmarshal(result.Raw, &gameHighScores)
	return &gameHighScores, err
}
