package tdlib

// PhotoSize Describes an image in JPEG format
type PhotoSize struct {
	tdCommon
	Type             string  `json:"type"`              // Image type (see https://core.telegram.org/constructor/photoSize)
	Photo            *File   `json:"photo"`             // Information about the image file
	Width            int32   `json:"width"`             // Image width
	Height           int32   `json:"height"`            // Image height
	ProgressiveSizes []int32 `json:"progressive_sizes"` // Sizes of progressive JPEG file prefixes, which can be used to preliminarily show the image; in bytes
}

// MessageType return the string telegram-type of PhotoSize
func (photoSize *PhotoSize) MessageType() string {
	return "photoSize"
}

// NewPhotoSize creates a new PhotoSize
//
// @param typeParam Image type (see https://core.telegram.org/constructor/photoSize)
// @param photo Information about the image file
// @param width Image width
// @param height Image height
// @param progressiveSizes Sizes of progressive JPEG file prefixes, which can be used to preliminarily show the image; in bytes
func NewPhotoSize(typeParam string, photo *File, width int32, height int32, progressiveSizes []int32) *PhotoSize {
	photoSizeTemp := PhotoSize{
		tdCommon:         tdCommon{Type: "photoSize"},
		Type:             typeParam,
		Photo:            photo,
		Width:            width,
		Height:           height,
		ProgressiveSizes: progressiveSizes,
	}

	return &photoSizeTemp
}
