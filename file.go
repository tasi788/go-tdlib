package tdlib

import (
	"encoding/json"
	"fmt"
)

// File Represents a file
type File struct {
	tdCommon
	Id           int32       `json:"id"`            // Unique file identifier
	DcId         int32       `json:"dc_id"`         // File data center
	Size         int32       `json:"size"`          // File size, in bytes; 0 if unknown
	ExpectedSize int32       `json:"expected_size"` // Approximate file size in bytes in case the exact file size is unknown. Can be used to show download/upload progress
	Local        *LocalFile  `json:"local"`         // Information about the local copy of the file
	Remote       *RemoteFile `json:"remote"`        // Information about the remote copy of the file
}

// MessageType return the string telegram-type of File
func (file *File) MessageType() string {
	return "file"
}

// NewFile creates a new File
//
// @param id Unique file identifier
// @param dcId File data center
// @param size File size, in bytes; 0 if unknown
// @param expectedSize Approximate file size in bytes in case the exact file size is unknown. Can be used to show download/upload progress
// @param local Information about the local copy of the file
// @param remote Information about the remote copy of the file
func NewFile(id int32, dcId int32, size int32, expectedSize int32, local *LocalFile, remote *RemoteFile) *File {
	fileTemp := File{
		tdCommon:     tdCommon{Type: "file"},
		Id:           id,
		DcId:         dcId,
		Size:         size,
		ExpectedSize: expectedSize,
		Local:        local,
		Remote:       remote,
	}

	return &fileTemp
}

// GetFile Returns information about a file; this is an offline request
// @param fileId Identifier of the file to get
func (client *Client) GetFile(fileId int32) (*File, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "getFile",
		"file_id": fileId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var fileDummy File
	err = json.Unmarshal(result.Raw, &fileDummy)
	return &fileDummy, err
}

// GetRemoteFile Returns information about a file by its remote ID; this is an offline request. Can be used to register a URL as a file for further uploading, or sending as a message. Even the request succeeds, the file can be used only if it is still accessible to the user. For example, if the file is from a message, then the message must be not deleted and accessible to the user. If the file database is disabled, then the corresponding object with the file must be preloaded by the application
// @param remoteFileId Remote identifier of the file to get
// @param fileType File type; pass null if unknown
func (client *Client) GetRemoteFile(remoteFileId string, fileType FileType) (*File, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":          "getRemoteFile",
		"remote_file_id": remoteFileId,
		"file_type":      fileType,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var fileDummy File
	err = json.Unmarshal(result.Raw, &fileDummy)
	return &fileDummy, err
}

// DownloadFile Downloads a file from the cloud. Download progress and completion of the download will be notified through updateFile updates
// @param fileId Identifier of the file to download
// @param priority Priority of the download (1-32). The higher the priority, the earlier the file will be downloaded. If the priorities of two files are equal, then the last one for which downloadFile was called will be downloaded first
// @param offset The starting position from which the file needs to be downloaded
// @param limit Number of bytes which need to be downloaded starting from the "offset" position before the download will automatically be canceled; use 0 to download without a limit
// @param synchronous If false, this request returns file state just after the download has been started. If true, this request returns file state only after the download has succeeded, has failed, has been canceled or a new downloadFile request with different offset/limit parameters was sent
func (client *Client) DownloadFile(fileId int32, priority int32, offset int32, limit int32, synchronous bool) (*File, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":       "downloadFile",
		"file_id":     fileId,
		"priority":    priority,
		"offset":      offset,
		"limit":       limit,
		"synchronous": synchronous,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var fileDummy File
	err = json.Unmarshal(result.Raw, &fileDummy)
	return &fileDummy, err
}

// UploadFile Asynchronously uploads a file to the cloud without sending it in a message. updateFile will be used to notify about upload progress and successful completion of the upload. The file will not have a persistent remote identifier until it will be sent in a message
// @param file File to upload
// @param fileType File type; pass null if unknown
// @param priority Priority of the upload (1-32). The higher the priority, the earlier the file will be uploaded. If the priorities of two files are equal, then the first one for which uploadFile was called will be uploaded first
func (client *Client) UploadFile(file InputFile, fileType FileType, priority int32) (*File, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":     "uploadFile",
		"file":      file,
		"file_type": fileType,
		"priority":  priority,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var fileDummy File
	err = json.Unmarshal(result.Raw, &fileDummy)
	return &fileDummy, err
}

// UploadStickerFile Uploads a PNG image with a sticker; returns the uploaded file
// @param userId Sticker file owner; ignored for regular users
// @param sticker Sticker file to upload
func (client *Client) UploadStickerFile(userId int64, sticker InputSticker) (*File, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "uploadStickerFile",
		"user_id": userId,
		"sticker": sticker,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var file File
	err = json.Unmarshal(result.Raw, &file)
	return &file, err
}

// GetMapThumbnailFile Returns information about a file with a map thumbnail in PNG format. Only map thumbnail files with size less than 1MB can be downloaded
// @param location Location of the map center
// @param zoom Map zoom level; 13-20
// @param width Map width in pixels before applying scale; 16-1024
// @param height Map height in pixels before applying scale; 16-1024
// @param scale Map scale; 1-3
// @param chatId Identifier of a chat, in which the thumbnail will be shown. Use 0 if unknown
func (client *Client) GetMapThumbnailFile(location *Location, zoom int32, width int32, height int32, scale int32, chatId int64) (*File, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":    "getMapThumbnailFile",
		"location": location,
		"zoom":     zoom,
		"width":    width,
		"height":   height,
		"scale":    scale,
		"chat_id":  chatId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var file File
	err = json.Unmarshal(result.Raw, &file)
	return &file, err
}
