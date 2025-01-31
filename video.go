package tdlib

// Video Describes a video file
type Video struct {
	tdCommon
	Duration          int32          `json:"duration"`           // Duration of the video, in seconds; as defined by the sender
	Width             int32          `json:"width"`              // Video width; as defined by the sender
	Height            int32          `json:"height"`             // Video height; as defined by the sender
	FileName          string         `json:"file_name"`          // Original name of the file; as defined by the sender
	MimeType          string         `json:"mime_type"`          // MIME type of the file; as defined by the sender
	HasStickers       bool           `json:"has_stickers"`       // True, if stickers were added to the video. The list of corresponding sticker sets can be received using getAttachedStickerSets
	SupportsStreaming bool           `json:"supports_streaming"` // True, if the video is supposed to be streamed
	Minithumbnail     *Minithumbnail `json:"minithumbnail"`      // Video minithumbnail; may be null
	Thumbnail         *Thumbnail     `json:"thumbnail"`          // Video thumbnail in JPEG or MPEG4 format; as defined by the sender; may be null
	Video             *File          `json:"video"`              // File containing the video
}

// MessageType return the string telegram-type of Video
func (video *Video) MessageType() string {
	return "video"
}

// NewVideo creates a new Video
//
// @param duration Duration of the video, in seconds; as defined by the sender
// @param width Video width; as defined by the sender
// @param height Video height; as defined by the sender
// @param fileName Original name of the file; as defined by the sender
// @param mimeType MIME type of the file; as defined by the sender
// @param hasStickers True, if stickers were added to the video. The list of corresponding sticker sets can be received using getAttachedStickerSets
// @param supportsStreaming True, if the video is supposed to be streamed
// @param minithumbnail Video minithumbnail; may be null
// @param thumbnail Video thumbnail in JPEG or MPEG4 format; as defined by the sender; may be null
// @param video File containing the video
func NewVideo(duration int32, width int32, height int32, fileName string, mimeType string, hasStickers bool, supportsStreaming bool, minithumbnail *Minithumbnail, thumbnail *Thumbnail, video *File) *Video {
	videoTemp := Video{
		tdCommon:          tdCommon{Type: "video"},
		Duration:          duration,
		Width:             width,
		Height:            height,
		FileName:          fileName,
		MimeType:          mimeType,
		HasStickers:       hasStickers,
		SupportsStreaming: supportsStreaming,
		Minithumbnail:     minithumbnail,
		Thumbnail:         thumbnail,
		Video:             video,
	}

	return &videoTemp
}
