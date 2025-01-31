package tdlib

import (
	"encoding/json"
)

// Thumbnail Represents a thumbnail
type Thumbnail struct {
	tdCommon
	Format ThumbnailFormat `json:"format"` // Thumbnail format
	Width  int32           `json:"width"`  // Thumbnail width
	Height int32           `json:"height"` // Thumbnail height
	File   *File           `json:"file"`   // The thumbnail
}

// MessageType return the string telegram-type of Thumbnail
func (thumbnail *Thumbnail) MessageType() string {
	return "thumbnail"
}

// NewThumbnail creates a new Thumbnail
//
// @param format Thumbnail format
// @param width Thumbnail width
// @param height Thumbnail height
// @param file The thumbnail
func NewThumbnail(format ThumbnailFormat, width int32, height int32, file *File) *Thumbnail {
	thumbnailTemp := Thumbnail{
		tdCommon: tdCommon{Type: "thumbnail"},
		Format:   format,
		Width:    width,
		Height:   height,
		File:     file,
	}

	return &thumbnailTemp
}

// UnmarshalJSON unmarshal to json
func (thumbnail *Thumbnail) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	tempObj := struct {
		tdCommon
		Width  int32 `json:"width"`  // Thumbnail width
		Height int32 `json:"height"` // Thumbnail height
		File   *File `json:"file"`   // The thumbnail
	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	thumbnail.tdCommon = tempObj.tdCommon
	thumbnail.Width = tempObj.Width
	thumbnail.Height = tempObj.Height
	thumbnail.File = tempObj.File

	fieldFormat, _ := unmarshalThumbnailFormat(objMap["format"])
	thumbnail.Format = fieldFormat

	return nil
}
