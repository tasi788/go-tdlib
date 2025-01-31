package tdlib

// Audio Describes an audio file. Audio is usually in MP3 or M4A format
type Audio struct {
	tdCommon
	Duration                int32          `json:"duration"`                  // Duration of the audio, in seconds; as defined by the sender
	Title                   string         `json:"title"`                     // Title of the audio; as defined by the sender
	Performer               string         `json:"performer"`                 // Performer of the audio; as defined by the sender
	FileName                string         `json:"file_name"`                 // Original name of the file; as defined by the sender
	MimeType                string         `json:"mime_type"`                 // The MIME type of the file; as defined by the sender
	AlbumCoverMinithumbnail *Minithumbnail `json:"album_cover_minithumbnail"` // The minithumbnail of the album cover; may be null
	AlbumCoverThumbnail     *Thumbnail     `json:"album_cover_thumbnail"`     // The thumbnail of the album cover in JPEG format; as defined by the sender. The full size thumbnail is supposed to be extracted from the downloaded file; may be null
	Audio                   *File          `json:"audio"`                     // File containing the audio
}

// MessageType return the string telegram-type of Audio
func (audio *Audio) MessageType() string {
	return "audio"
}

// NewAudio creates a new Audio
//
// @param duration Duration of the audio, in seconds; as defined by the sender
// @param title Title of the audio; as defined by the sender
// @param performer Performer of the audio; as defined by the sender
// @param fileName Original name of the file; as defined by the sender
// @param mimeType The MIME type of the file; as defined by the sender
// @param albumCoverMinithumbnail The minithumbnail of the album cover; may be null
// @param albumCoverThumbnail The thumbnail of the album cover in JPEG format; as defined by the sender. The full size thumbnail is supposed to be extracted from the downloaded file; may be null
// @param audio File containing the audio
func NewAudio(duration int32, title string, performer string, fileName string, mimeType string, albumCoverMinithumbnail *Minithumbnail, albumCoverThumbnail *Thumbnail, audio *File) *Audio {
	audioTemp := Audio{
		tdCommon:                tdCommon{Type: "audio"},
		Duration:                duration,
		Title:                   title,
		Performer:               performer,
		FileName:                fileName,
		MimeType:                mimeType,
		AlbumCoverMinithumbnail: albumCoverMinithumbnail,
		AlbumCoverThumbnail:     albumCoverThumbnail,
		Audio:                   audio,
	}

	return &audioTemp
}
