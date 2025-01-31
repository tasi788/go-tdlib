package tdlib

import (
	"encoding/json"
)

// InputThumbnail A thumbnail to be sent along with a file; must be in JPEG or WEBP format for stickers, and less than 200 KB in size
type InputThumbnail struct {
	tdCommon
	Thumbnail InputFile `json:"thumbnail"` // Thumbnail file to send. Sending thumbnails by file_id is currently not supported
	Width     int32     `json:"width"`     // Thumbnail width, usually shouldn't exceed 320. Use 0 if unknown
	Height    int32     `json:"height"`    // Thumbnail height, usually shouldn't exceed 320. Use 0 if unknown
}

// MessageType return the string telegram-type of InputThumbnail
func (inputThumbnail *InputThumbnail) MessageType() string {
	return "inputThumbnail"
}

// NewInputThumbnail creates a new InputThumbnail
//
// @param thumbnail Thumbnail file to send. Sending thumbnails by file_id is currently not supported
// @param width Thumbnail width, usually shouldn't exceed 320. Use 0 if unknown
// @param height Thumbnail height, usually shouldn't exceed 320. Use 0 if unknown
func NewInputThumbnail(thumbnail InputFile, width int32, height int32) *InputThumbnail {
	inputThumbnailTemp := InputThumbnail{
		tdCommon:  tdCommon{Type: "inputThumbnail"},
		Thumbnail: thumbnail,
		Width:     width,
		Height:    height,
	}

	return &inputThumbnailTemp
}

// UnmarshalJSON unmarshal to json
func (inputThumbnail *InputThumbnail) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	tempObj := struct {
		tdCommon
		Width  int32 `json:"width"`  // Thumbnail width, usually shouldn't exceed 320. Use 0 if unknown
		Height int32 `json:"height"` // Thumbnail height, usually shouldn't exceed 320. Use 0 if unknown
	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	inputThumbnail.tdCommon = tempObj.tdCommon
	inputThumbnail.Width = tempObj.Width
	inputThumbnail.Height = tempObj.Height

	fieldThumbnail, _ := unmarshalInputFile(objMap["thumbnail"])
	inputThumbnail.Thumbnail = fieldThumbnail

	return nil
}
