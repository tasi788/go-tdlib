package tdlib

// Minithumbnail Thumbnail image of a very poor quality and low resolution
type Minithumbnail struct {
	tdCommon
	Width  int32  `json:"width"`  // Thumbnail width, usually doesn't exceed 40
	Height int32  `json:"height"` // Thumbnail height, usually doesn't exceed 40
	Data   []byte `json:"data"`   // The thumbnail in JPEG format
}

// MessageType return the string telegram-type of Minithumbnail
func (minithumbnail *Minithumbnail) MessageType() string {
	return "minithumbnail"
}

// NewMinithumbnail creates a new Minithumbnail
//
// @param width Thumbnail width, usually doesn't exceed 40
// @param height Thumbnail height, usually doesn't exceed 40
// @param data The thumbnail in JPEG format
func NewMinithumbnail(width int32, height int32, data []byte) *Minithumbnail {
	minithumbnailTemp := Minithumbnail{
		tdCommon: tdCommon{Type: "minithumbnail"},
		Width:    width,
		Height:   height,
		Data:     data,
	}

	return &minithumbnailTemp
}
