package tdlib

// Photo Describes a photo
type Photo struct {
	tdCommon
	HasStickers   bool           `json:"has_stickers"`  // True, if stickers were added to the photo. The list of corresponding sticker sets can be received using getAttachedStickerSets
	Minithumbnail *Minithumbnail `json:"minithumbnail"` // Photo minithumbnail; may be null
	Sizes         []PhotoSize    `json:"sizes"`         // Available variants of the photo, in different sizes
}

// MessageType return the string telegram-type of Photo
func (photo *Photo) MessageType() string {
	return "photo"
}

// NewPhoto creates a new Photo
//
// @param hasStickers True, if stickers were added to the photo. The list of corresponding sticker sets can be received using getAttachedStickerSets
// @param minithumbnail Photo minithumbnail; may be null
// @param sizes Available variants of the photo, in different sizes
func NewPhoto(hasStickers bool, minithumbnail *Minithumbnail, sizes []PhotoSize) *Photo {
	photoTemp := Photo{
		tdCommon:      tdCommon{Type: "photo"},
		HasStickers:   hasStickers,
		Minithumbnail: minithumbnail,
		Sizes:         sizes,
	}

	return &photoTemp
}
