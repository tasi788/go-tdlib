package tdlib

// Animation Describes an animation file. The animation must be encoded in GIF or MPEG4 format
type Animation struct {
	tdCommon
	Duration      int32          `json:"duration"`      // Duration of the animation, in seconds; as defined by the sender
	Width         int32          `json:"width"`         // Width of the animation
	Height        int32          `json:"height"`        // Height of the animation
	FileName      string         `json:"file_name"`     // Original name of the file; as defined by the sender
	MimeType      string         `json:"mime_type"`     // MIME type of the file, usually "image/gif" or "video/mp4"
	HasStickers   bool           `json:"has_stickers"`  // True, if stickers were added to the animation. The list of corresponding sticker set can be received using getAttachedStickerSets
	Minithumbnail *Minithumbnail `json:"minithumbnail"` // Animation minithumbnail; may be null
	Thumbnail     *Thumbnail     `json:"thumbnail"`     // Animation thumbnail in JPEG or MPEG4 format; may be null
	Animation     *File          `json:"animation"`     // File containing the animation
}

// MessageType return the string telegram-type of Animation
func (animation *Animation) MessageType() string {
	return "animation"
}

// NewAnimation creates a new Animation
//
// @param duration Duration of the animation, in seconds; as defined by the sender
// @param width Width of the animation
// @param height Height of the animation
// @param fileName Original name of the file; as defined by the sender
// @param mimeType MIME type of the file, usually "image/gif" or "video/mp4"
// @param hasStickers True, if stickers were added to the animation. The list of corresponding sticker set can be received using getAttachedStickerSets
// @param minithumbnail Animation minithumbnail; may be null
// @param thumbnail Animation thumbnail in JPEG or MPEG4 format; may be null
// @param animation File containing the animation
func NewAnimation(duration int32, width int32, height int32, fileName string, mimeType string, hasStickers bool, minithumbnail *Minithumbnail, thumbnail *Thumbnail, animation *File) *Animation {
	animationTemp := Animation{
		tdCommon:      tdCommon{Type: "animation"},
		Duration:      duration,
		Width:         width,
		Height:        height,
		FileName:      fileName,
		MimeType:      mimeType,
		HasStickers:   hasStickers,
		Minithumbnail: minithumbnail,
		Thumbnail:     thumbnail,
		Animation:     animation,
	}

	return &animationTemp
}
