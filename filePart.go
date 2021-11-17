package tdlib

import (
	"encoding/json"
	"fmt"
)

// FilePart Contains a part of a file
type FilePart struct {
	tdCommon
	Data []byte `json:"data"` // File bytes
}

// MessageType return the string telegram-type of FilePart
func (filePart *FilePart) MessageType() string {
	return "filePart"
}

// NewFilePart creates a new FilePart
//
// @param data File bytes
func NewFilePart(data []byte) *FilePart {
	filePartTemp := FilePart{
		tdCommon: tdCommon{Type: "filePart"},
		Data:     data,
	}

	return &filePartTemp
}

// ReadFilePart Reads a part of a file from the TDLib file cache and returns read bytes. This method is intended to be used only if the application has no direct access to TDLib's file system, because it is usually slower than a direct read from the file
// @param fileId Identifier of the file. The file must be located in the TDLib file cache
// @param offset The offset from which to read the file
// @param count Number of bytes to read. An error will be returned if there are not enough bytes available in the file from the specified position. Pass 0 to read all available data from the specified position
func (client *Client) ReadFilePart(fileId int32, offset int32, count int32) (*FilePart, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "readFilePart",
		"file_id": fileId,
		"offset":  offset,
		"count":   count,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var filePart FilePart
	err = json.Unmarshal(result.Raw, &filePart)
	return &filePart, err
}

// GetGroupCallStreamSegment Returns a file with a segment of a group call stream in a modified OGG format for audio or MPEG-4 format for video
// @param groupCallId Group call identifier
// @param timeOffset Point in time when the stream segment begins; Unix timestamp in milliseconds
// @param scale Segment duration scale; 0-1. Segment's duration is 1000/(2**scale) milliseconds
// @param channelId Identifier of an audio/video channel to get as received from tgcalls
// @param videoQuality Video quality as received from tgcalls; pass null to get the worst available quality
func (client *Client) GetGroupCallStreamSegment(groupCallId int32, timeOffset int64, scale int32, channelId int32, videoQuality GroupCallVideoQuality) (*FilePart, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":         "getGroupCallStreamSegment",
		"group_call_id": groupCallId,
		"time_offset":   timeOffset,
		"scale":         scale,
		"channel_id":    channelId,
		"video_quality": videoQuality,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var filePart FilePart
	err = json.Unmarshal(result.Raw, &filePart)
	return &filePart, err
}
