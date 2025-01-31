package tdlib

// Game Describes a game
type Game struct {
	tdCommon
	Id          JSONInt64      `json:"id"`          // Game ID
	ShortName   string         `json:"short_name"`  // Game short name. To share a game use the URL https://t.me/{bot_username}?game={game_short_name}
	Title       string         `json:"title"`       // Game title
	Text        *FormattedText `json:"text"`        // Game text, usually containing scoreboards for a game
	Description string         `json:"description"` // Game description
	Photo       *Photo         `json:"photo"`       // Game photo
	Animation   *Animation     `json:"animation"`   // Game animation; may be null
}

// MessageType return the string telegram-type of Game
func (game *Game) MessageType() string {
	return "game"
}

// NewGame creates a new Game
//
// @param id Game ID
// @param shortName Game short name. To share a game use the URL https://t.me/{bot_username}?game={game_short_name}
// @param title Game title
// @param text Game text, usually containing scoreboards for a game
// @param description Game description
// @param photo Game photo
// @param animation Game animation; may be null
func NewGame(id JSONInt64, shortName string, title string, text *FormattedText, description string, photo *Photo, animation *Animation) *Game {
	gameTemp := Game{
		tdCommon:    tdCommon{Type: "game"},
		Id:          id,
		ShortName:   shortName,
		Title:       title,
		Text:        text,
		Description: description,
		Photo:       photo,
		Animation:   animation,
	}

	return &gameTemp
}
