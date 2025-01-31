package tdlib

// ProfilePhoto Describes a user profile photo
type ProfilePhoto struct {
	tdCommon
	Id            JSONInt64      `json:"id"`            // Photo identifier; 0 for an empty photo. Can be used to find a photo in a list of user profile photos
	Small         *File          `json:"small"`         // A small (160x160) user profile photo. The file can be downloaded only before the photo is changed
	Big           *File          `json:"big"`           // A big (640x640) user profile photo. The file can be downloaded only before the photo is changed
	Minithumbnail *Minithumbnail `json:"minithumbnail"` // User profile photo minithumbnail; may be null
	HasAnimation  bool           `json:"has_animation"` // True, if the photo has animated variant
}

// MessageType return the string telegram-type of ProfilePhoto
func (profilePhoto *ProfilePhoto) MessageType() string {
	return "profilePhoto"
}

// NewProfilePhoto creates a new ProfilePhoto
//
// @param id Photo identifier; 0 for an empty photo. Can be used to find a photo in a list of user profile photos
// @param small A small (160x160) user profile photo. The file can be downloaded only before the photo is changed
// @param big A big (640x640) user profile photo. The file can be downloaded only before the photo is changed
// @param minithumbnail User profile photo minithumbnail; may be null
// @param hasAnimation True, if the photo has animated variant
func NewProfilePhoto(id JSONInt64, small *File, big *File, minithumbnail *Minithumbnail, hasAnimation bool) *ProfilePhoto {
	profilePhotoTemp := ProfilePhoto{
		tdCommon:      tdCommon{Type: "profilePhoto"},
		Id:            id,
		Small:         small,
		Big:           big,
		Minithumbnail: minithumbnail,
		HasAnimation:  hasAnimation,
	}

	return &profilePhotoTemp
}
