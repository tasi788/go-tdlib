package tdlib

// GameHighScore Contains one row of the game high score table
type GameHighScore struct {
	tdCommon
	Position int32 `json:"position"` // Position in the high score table
	UserId   int64 `json:"user_id"`  // User identifier
	Score    int32 `json:"score"`    // User score
}

// MessageType return the string telegram-type of GameHighScore
func (gameHighScore *GameHighScore) MessageType() string {
	return "gameHighScore"
}

// NewGameHighScore creates a new GameHighScore
//
// @param position Position in the high score table
// @param userId User identifier
// @param score User score
func NewGameHighScore(position int32, userId int64, score int32) *GameHighScore {
	gameHighScoreTemp := GameHighScore{
		tdCommon: tdCommon{Type: "gameHighScore"},
		Position: position,
		UserId:   userId,
		Score:    score,
	}

	return &gameHighScoreTemp
}
