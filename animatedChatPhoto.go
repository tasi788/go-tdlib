package tdlib

// AnimatedChatPhoto Animated variant of a chat photo in MPEG4 format
type AnimatedChatPhoto struct {
	tdCommon
	Length             int32   `json:"length"`               // Animation width and height
	File               *File   `json:"file"`                 // Information about the animation file
	MainFrameTimestamp float64 `json:"main_frame_timestamp"` // Timestamp of the frame, used as a static chat photo
}

// MessageType return the string telegram-type of AnimatedChatPhoto
func (animatedChatPhoto *AnimatedChatPhoto) MessageType() string {
	return "animatedChatPhoto"
}

// NewAnimatedChatPhoto creates a new AnimatedChatPhoto
//
// @param length Animation width and height
// @param file Information about the animation file
// @param mainFrameTimestamp Timestamp of the frame, used as a static chat photo
func NewAnimatedChatPhoto(length int32, file *File, mainFrameTimestamp float64) *AnimatedChatPhoto {
	animatedChatPhotoTemp := AnimatedChatPhoto{
		tdCommon:           tdCommon{Type: "animatedChatPhoto"},
		Length:             length,
		File:               file,
		MainFrameTimestamp: mainFrameTimestamp,
	}

	return &animatedChatPhotoTemp
}
