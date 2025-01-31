package tdlib

// ChatLocation Represents a location to which a chat is connected
type ChatLocation struct {
	tdCommon
	Location *Location `json:"location"` // The location
	Address  string    `json:"address"`  // Location address; 1-64 characters, as defined by the chat owner
}

// MessageType return the string telegram-type of ChatLocation
func (chatLocation *ChatLocation) MessageType() string {
	return "chatLocation"
}

// NewChatLocation creates a new ChatLocation
//
// @param location The location
// @param address Location address; 1-64 characters, as defined by the chat owner
func NewChatLocation(location *Location, address string) *ChatLocation {
	chatLocationTemp := ChatLocation{
		tdCommon: tdCommon{Type: "chatLocation"},
		Location: location,
		Address:  address,
	}

	return &chatLocationTemp
}
