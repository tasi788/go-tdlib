package tdlib

// ChatPhotoInfo Contains basic information about the photo of a chat
type ChatPhotoInfo struct {
	tdCommon
	Small         *File          `json:"small"`         // A small (160x160) chat photo variant in JPEG format. The file can be downloaded only before the photo is changed
	Big           *File          `json:"big"`           // A big (640x640) chat photo variant in JPEG format. The file can be downloaded only before the photo is changed
	Minithumbnail *Minithumbnail `json:"minithumbnail"` // Chat photo minithumbnail; may be null
	HasAnimation  bool           `json:"has_animation"` // True, if the photo has animated variant
}

// MessageType return the string telegram-type of ChatPhotoInfo
func (chatPhotoInfo *ChatPhotoInfo) MessageType() string {
	return "chatPhotoInfo"
}

// NewChatPhotoInfo creates a new ChatPhotoInfo
//
// @param small A small (160x160) chat photo variant in JPEG format. The file can be downloaded only before the photo is changed
// @param big A big (640x640) chat photo variant in JPEG format. The file can be downloaded only before the photo is changed
// @param minithumbnail Chat photo minithumbnail; may be null
// @param hasAnimation True, if the photo has animated variant
func NewChatPhotoInfo(small *File, big *File, minithumbnail *Minithumbnail, hasAnimation bool) *ChatPhotoInfo {
	chatPhotoInfoTemp := ChatPhotoInfo{
		tdCommon:      tdCommon{Type: "chatPhotoInfo"},
		Small:         small,
		Big:           big,
		Minithumbnail: minithumbnail,
		HasAnimation:  hasAnimation,
	}

	return &chatPhotoInfoTemp
}
