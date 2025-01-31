package tdlib

// VideoNote Describes a video note. The video must be equal in width and height, cropped to a circle, and stored in MPEG4 format
type VideoNote struct {
	tdCommon
	Duration      int32          `json:"duration"`      // Duration of the video, in seconds; as defined by the sender
	Length        int32          `json:"length"`        // Video width and height; as defined by the sender
	Minithumbnail *Minithumbnail `json:"minithumbnail"` // Video minithumbnail; may be null
	Thumbnail     *Thumbnail     `json:"thumbnail"`     // Video thumbnail in JPEG format; as defined by the sender; may be null
	Video         *File          `json:"video"`         // File containing the video
}

// MessageType return the string telegram-type of VideoNote
func (videoNote *VideoNote) MessageType() string {
	return "videoNote"
}

// NewVideoNote creates a new VideoNote
//
// @param duration Duration of the video, in seconds; as defined by the sender
// @param length Video width and height; as defined by the sender
// @param minithumbnail Video minithumbnail; may be null
// @param thumbnail Video thumbnail in JPEG format; as defined by the sender; may be null
// @param video File containing the video
func NewVideoNote(duration int32, length int32, minithumbnail *Minithumbnail, thumbnail *Thumbnail, video *File) *VideoNote {
	videoNoteTemp := VideoNote{
		tdCommon:      tdCommon{Type: "videoNote"},
		Duration:      duration,
		Length:        length,
		Minithumbnail: minithumbnail,
		Thumbnail:     thumbnail,
		Video:         video,
	}

	return &videoNoteTemp
}
