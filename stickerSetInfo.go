package tdlib

// StickerSetInfo Represents short information about a sticker set
type StickerSetInfo struct {
	tdCommon
	Id               JSONInt64          `json:"id"`                // Identifier of the sticker set
	Title            string             `json:"title"`             // Title of the sticker set
	Name             string             `json:"name"`              // Name of the sticker set
	Thumbnail        *Thumbnail         `json:"thumbnail"`         // Sticker set thumbnail in WEBP or TGS format with width and height 100; may be null
	ThumbnailOutline []ClosedVectorPath `json:"thumbnail_outline"` // Sticker set thumbnail's outline represented as a list of closed vector paths; may be empty. The coordinate system origin is in the upper-left corner
	IsInstalled      bool               `json:"is_installed"`      // True, if the sticker set has been installed by the current user
	IsArchived       bool               `json:"is_archived"`       // True, if the sticker set has been archived. A sticker set can't be installed and archived simultaneously
	IsOfficial       bool               `json:"is_official"`       // True, if the sticker set is official
	IsAnimated       bool               `json:"is_animated"`       // True, is the stickers in the set are animated
	IsMasks          bool               `json:"is_masks"`          // True, if the stickers in the set are masks
	IsViewed         bool               `json:"is_viewed"`         // True for already viewed trending sticker sets
	Size             int32              `json:"size"`              // Total number of stickers in the set
	Covers           []Sticker          `json:"covers"`            // Up to the first 5 stickers from the set, depending on the context. If the application needs more stickers the full sticker set needs to be requested
}

// MessageType return the string telegram-type of StickerSetInfo
func (stickerSetInfo *StickerSetInfo) MessageType() string {
	return "stickerSetInfo"
}

// NewStickerSetInfo creates a new StickerSetInfo
//
// @param id Identifier of the sticker set
// @param title Title of the sticker set
// @param name Name of the sticker set
// @param thumbnail Sticker set thumbnail in WEBP or TGS format with width and height 100; may be null
// @param thumbnailOutline Sticker set thumbnail's outline represented as a list of closed vector paths; may be empty. The coordinate system origin is in the upper-left corner
// @param isInstalled True, if the sticker set has been installed by the current user
// @param isArchived True, if the sticker set has been archived. A sticker set can't be installed and archived simultaneously
// @param isOfficial True, if the sticker set is official
// @param isAnimated True, is the stickers in the set are animated
// @param isMasks True, if the stickers in the set are masks
// @param isViewed True for already viewed trending sticker sets
// @param size Total number of stickers in the set
// @param covers Up to the first 5 stickers from the set, depending on the context. If the application needs more stickers the full sticker set needs to be requested
func NewStickerSetInfo(id JSONInt64, title string, name string, thumbnail *Thumbnail, thumbnailOutline []ClosedVectorPath, isInstalled bool, isArchived bool, isOfficial bool, isAnimated bool, isMasks bool, isViewed bool, size int32, covers []Sticker) *StickerSetInfo {
	stickerSetInfoTemp := StickerSetInfo{
		tdCommon:         tdCommon{Type: "stickerSetInfo"},
		Id:               id,
		Title:            title,
		Name:             name,
		Thumbnail:        thumbnail,
		ThumbnailOutline: thumbnailOutline,
		IsInstalled:      isInstalled,
		IsArchived:       isArchived,
		IsOfficial:       isOfficial,
		IsAnimated:       isAnimated,
		IsMasks:          isMasks,
		IsViewed:         isViewed,
		Size:             size,
		Covers:           covers,
	}

	return &stickerSetInfoTemp
}
