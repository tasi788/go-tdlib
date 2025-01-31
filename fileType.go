package tdlib

import (
	"encoding/json"
	"fmt"
)

// FileType Represents the type of a file
type FileType interface {
	GetFileTypeEnum() FileTypeEnum
}

// FileTypeEnum Alias for abstract FileType 'Sub-Classes', used as constant-enum here
type FileTypeEnum string

// FileType enums
const (
	FileTypeNoneType            FileTypeEnum = "fileTypeNone"
	FileTypeAnimationType       FileTypeEnum = "fileTypeAnimation"
	FileTypeAudioType           FileTypeEnum = "fileTypeAudio"
	FileTypeDocumentType        FileTypeEnum = "fileTypeDocument"
	FileTypePhotoType           FileTypeEnum = "fileTypePhoto"
	FileTypeProfilePhotoType    FileTypeEnum = "fileTypeProfilePhoto"
	FileTypeSecretType          FileTypeEnum = "fileTypeSecret"
	FileTypeSecretThumbnailType FileTypeEnum = "fileTypeSecretThumbnail"
	FileTypeSecureType          FileTypeEnum = "fileTypeSecure"
	FileTypeStickerType         FileTypeEnum = "fileTypeSticker"
	FileTypeThumbnailType       FileTypeEnum = "fileTypeThumbnail"
	FileTypeUnknownType         FileTypeEnum = "fileTypeUnknown"
	FileTypeVideoType           FileTypeEnum = "fileTypeVideo"
	FileTypeVideoNoteType       FileTypeEnum = "fileTypeVideoNote"
	FileTypeVoiceNoteType       FileTypeEnum = "fileTypeVoiceNote"
	FileTypeWallpaperType       FileTypeEnum = "fileTypeWallpaper"
)

func unmarshalFileType(rawMsg *json.RawMessage) (FileType, error) {

	if rawMsg == nil {
		return nil, nil
	}
	var objMap map[string]interface{}
	err := json.Unmarshal(*rawMsg, &objMap)
	if err != nil {
		return nil, err
	}

	switch FileTypeEnum(objMap["@type"].(string)) {
	case FileTypeNoneType:
		var fileTypeNone FileTypeNone
		err := json.Unmarshal(*rawMsg, &fileTypeNone)
		return &fileTypeNone, err

	case FileTypeAnimationType:
		var fileTypeAnimation FileTypeAnimation
		err := json.Unmarshal(*rawMsg, &fileTypeAnimation)
		return &fileTypeAnimation, err

	case FileTypeAudioType:
		var fileTypeAudio FileTypeAudio
		err := json.Unmarshal(*rawMsg, &fileTypeAudio)
		return &fileTypeAudio, err

	case FileTypeDocumentType:
		var fileTypeDocument FileTypeDocument
		err := json.Unmarshal(*rawMsg, &fileTypeDocument)
		return &fileTypeDocument, err

	case FileTypePhotoType:
		var fileTypePhoto FileTypePhoto
		err := json.Unmarshal(*rawMsg, &fileTypePhoto)
		return &fileTypePhoto, err

	case FileTypeProfilePhotoType:
		var fileTypeProfilePhoto FileTypeProfilePhoto
		err := json.Unmarshal(*rawMsg, &fileTypeProfilePhoto)
		return &fileTypeProfilePhoto, err

	case FileTypeSecretType:
		var fileTypeSecret FileTypeSecret
		err := json.Unmarshal(*rawMsg, &fileTypeSecret)
		return &fileTypeSecret, err

	case FileTypeSecretThumbnailType:
		var fileTypeSecretThumbnail FileTypeSecretThumbnail
		err := json.Unmarshal(*rawMsg, &fileTypeSecretThumbnail)
		return &fileTypeSecretThumbnail, err

	case FileTypeSecureType:
		var fileTypeSecure FileTypeSecure
		err := json.Unmarshal(*rawMsg, &fileTypeSecure)
		return &fileTypeSecure, err

	case FileTypeStickerType:
		var fileTypeSticker FileTypeSticker
		err := json.Unmarshal(*rawMsg, &fileTypeSticker)
		return &fileTypeSticker, err

	case FileTypeThumbnailType:
		var fileTypeThumbnail FileTypeThumbnail
		err := json.Unmarshal(*rawMsg, &fileTypeThumbnail)
		return &fileTypeThumbnail, err

	case FileTypeUnknownType:
		var fileTypeUnknown FileTypeUnknown
		err := json.Unmarshal(*rawMsg, &fileTypeUnknown)
		return &fileTypeUnknown, err

	case FileTypeVideoType:
		var fileTypeVideo FileTypeVideo
		err := json.Unmarshal(*rawMsg, &fileTypeVideo)
		return &fileTypeVideo, err

	case FileTypeVideoNoteType:
		var fileTypeVideoNote FileTypeVideoNote
		err := json.Unmarshal(*rawMsg, &fileTypeVideoNote)
		return &fileTypeVideoNote, err

	case FileTypeVoiceNoteType:
		var fileTypeVoiceNote FileTypeVoiceNote
		err := json.Unmarshal(*rawMsg, &fileTypeVoiceNote)
		return &fileTypeVoiceNote, err

	case FileTypeWallpaperType:
		var fileTypeWallpaper FileTypeWallpaper
		err := json.Unmarshal(*rawMsg, &fileTypeWallpaper)
		return &fileTypeWallpaper, err

	default:
		return nil, fmt.Errorf("Error UnMarshaling, unknown type:" + objMap["@type"].(string))
	}
}

// FileTypeNone The data is not a file
type FileTypeNone struct {
	tdCommon
}

// MessageType return the string telegram-type of FileTypeNone
func (fileTypeNone *FileTypeNone) MessageType() string {
	return "fileTypeNone"
}

// NewFileTypeNone creates a new FileTypeNone
//
func NewFileTypeNone() *FileTypeNone {
	fileTypeNoneTemp := FileTypeNone{
		tdCommon: tdCommon{Type: "fileTypeNone"},
	}

	return &fileTypeNoneTemp
}

// GetFileTypeEnum return the enum type of this object
func (fileTypeNone *FileTypeNone) GetFileTypeEnum() FileTypeEnum {
	return FileTypeNoneType
}

// FileTypeAnimation The file is an animation
type FileTypeAnimation struct {
	tdCommon
}

// MessageType return the string telegram-type of FileTypeAnimation
func (fileTypeAnimation *FileTypeAnimation) MessageType() string {
	return "fileTypeAnimation"
}

// NewFileTypeAnimation creates a new FileTypeAnimation
//
func NewFileTypeAnimation() *FileTypeAnimation {
	fileTypeAnimationTemp := FileTypeAnimation{
		tdCommon: tdCommon{Type: "fileTypeAnimation"},
	}

	return &fileTypeAnimationTemp
}

// GetFileTypeEnum return the enum type of this object
func (fileTypeAnimation *FileTypeAnimation) GetFileTypeEnum() FileTypeEnum {
	return FileTypeAnimationType
}

// FileTypeAudio The file is an audio file
type FileTypeAudio struct {
	tdCommon
}

// MessageType return the string telegram-type of FileTypeAudio
func (fileTypeAudio *FileTypeAudio) MessageType() string {
	return "fileTypeAudio"
}

// NewFileTypeAudio creates a new FileTypeAudio
//
func NewFileTypeAudio() *FileTypeAudio {
	fileTypeAudioTemp := FileTypeAudio{
		tdCommon: tdCommon{Type: "fileTypeAudio"},
	}

	return &fileTypeAudioTemp
}

// GetFileTypeEnum return the enum type of this object
func (fileTypeAudio *FileTypeAudio) GetFileTypeEnum() FileTypeEnum {
	return FileTypeAudioType
}

// FileTypeDocument The file is a document
type FileTypeDocument struct {
	tdCommon
}

// MessageType return the string telegram-type of FileTypeDocument
func (fileTypeDocument *FileTypeDocument) MessageType() string {
	return "fileTypeDocument"
}

// NewFileTypeDocument creates a new FileTypeDocument
//
func NewFileTypeDocument() *FileTypeDocument {
	fileTypeDocumentTemp := FileTypeDocument{
		tdCommon: tdCommon{Type: "fileTypeDocument"},
	}

	return &fileTypeDocumentTemp
}

// GetFileTypeEnum return the enum type of this object
func (fileTypeDocument *FileTypeDocument) GetFileTypeEnum() FileTypeEnum {
	return FileTypeDocumentType
}

// FileTypePhoto The file is a photo
type FileTypePhoto struct {
	tdCommon
}

// MessageType return the string telegram-type of FileTypePhoto
func (fileTypePhoto *FileTypePhoto) MessageType() string {
	return "fileTypePhoto"
}

// NewFileTypePhoto creates a new FileTypePhoto
//
func NewFileTypePhoto() *FileTypePhoto {
	fileTypePhotoTemp := FileTypePhoto{
		tdCommon: tdCommon{Type: "fileTypePhoto"},
	}

	return &fileTypePhotoTemp
}

// GetFileTypeEnum return the enum type of this object
func (fileTypePhoto *FileTypePhoto) GetFileTypeEnum() FileTypeEnum {
	return FileTypePhotoType
}

// FileTypeProfilePhoto The file is a profile photo
type FileTypeProfilePhoto struct {
	tdCommon
}

// MessageType return the string telegram-type of FileTypeProfilePhoto
func (fileTypeProfilePhoto *FileTypeProfilePhoto) MessageType() string {
	return "fileTypeProfilePhoto"
}

// NewFileTypeProfilePhoto creates a new FileTypeProfilePhoto
//
func NewFileTypeProfilePhoto() *FileTypeProfilePhoto {
	fileTypeProfilePhotoTemp := FileTypeProfilePhoto{
		tdCommon: tdCommon{Type: "fileTypeProfilePhoto"},
	}

	return &fileTypeProfilePhotoTemp
}

// GetFileTypeEnum return the enum type of this object
func (fileTypeProfilePhoto *FileTypeProfilePhoto) GetFileTypeEnum() FileTypeEnum {
	return FileTypeProfilePhotoType
}

// FileTypeSecret The file was sent to a secret chat (the file type is not known to the server)
type FileTypeSecret struct {
	tdCommon
}

// MessageType return the string telegram-type of FileTypeSecret
func (fileTypeSecret *FileTypeSecret) MessageType() string {
	return "fileTypeSecret"
}

// NewFileTypeSecret creates a new FileTypeSecret
//
func NewFileTypeSecret() *FileTypeSecret {
	fileTypeSecretTemp := FileTypeSecret{
		tdCommon: tdCommon{Type: "fileTypeSecret"},
	}

	return &fileTypeSecretTemp
}

// GetFileTypeEnum return the enum type of this object
func (fileTypeSecret *FileTypeSecret) GetFileTypeEnum() FileTypeEnum {
	return FileTypeSecretType
}

// FileTypeSecretThumbnail The file is a thumbnail of a file from a secret chat
type FileTypeSecretThumbnail struct {
	tdCommon
}

// MessageType return the string telegram-type of FileTypeSecretThumbnail
func (fileTypeSecretThumbnail *FileTypeSecretThumbnail) MessageType() string {
	return "fileTypeSecretThumbnail"
}

// NewFileTypeSecretThumbnail creates a new FileTypeSecretThumbnail
//
func NewFileTypeSecretThumbnail() *FileTypeSecretThumbnail {
	fileTypeSecretThumbnailTemp := FileTypeSecretThumbnail{
		tdCommon: tdCommon{Type: "fileTypeSecretThumbnail"},
	}

	return &fileTypeSecretThumbnailTemp
}

// GetFileTypeEnum return the enum type of this object
func (fileTypeSecretThumbnail *FileTypeSecretThumbnail) GetFileTypeEnum() FileTypeEnum {
	return FileTypeSecretThumbnailType
}

// FileTypeSecure The file is a file from Secure storage used for storing Telegram Passport files
type FileTypeSecure struct {
	tdCommon
}

// MessageType return the string telegram-type of FileTypeSecure
func (fileTypeSecure *FileTypeSecure) MessageType() string {
	return "fileTypeSecure"
}

// NewFileTypeSecure creates a new FileTypeSecure
//
func NewFileTypeSecure() *FileTypeSecure {
	fileTypeSecureTemp := FileTypeSecure{
		tdCommon: tdCommon{Type: "fileTypeSecure"},
	}

	return &fileTypeSecureTemp
}

// GetFileTypeEnum return the enum type of this object
func (fileTypeSecure *FileTypeSecure) GetFileTypeEnum() FileTypeEnum {
	return FileTypeSecureType
}

// FileTypeSticker The file is a sticker
type FileTypeSticker struct {
	tdCommon
}

// MessageType return the string telegram-type of FileTypeSticker
func (fileTypeSticker *FileTypeSticker) MessageType() string {
	return "fileTypeSticker"
}

// NewFileTypeSticker creates a new FileTypeSticker
//
func NewFileTypeSticker() *FileTypeSticker {
	fileTypeStickerTemp := FileTypeSticker{
		tdCommon: tdCommon{Type: "fileTypeSticker"},
	}

	return &fileTypeStickerTemp
}

// GetFileTypeEnum return the enum type of this object
func (fileTypeSticker *FileTypeSticker) GetFileTypeEnum() FileTypeEnum {
	return FileTypeStickerType
}

// FileTypeThumbnail The file is a thumbnail of another file
type FileTypeThumbnail struct {
	tdCommon
}

// MessageType return the string telegram-type of FileTypeThumbnail
func (fileTypeThumbnail *FileTypeThumbnail) MessageType() string {
	return "fileTypeThumbnail"
}

// NewFileTypeThumbnail creates a new FileTypeThumbnail
//
func NewFileTypeThumbnail() *FileTypeThumbnail {
	fileTypeThumbnailTemp := FileTypeThumbnail{
		tdCommon: tdCommon{Type: "fileTypeThumbnail"},
	}

	return &fileTypeThumbnailTemp
}

// GetFileTypeEnum return the enum type of this object
func (fileTypeThumbnail *FileTypeThumbnail) GetFileTypeEnum() FileTypeEnum {
	return FileTypeThumbnailType
}

// FileTypeUnknown The file type is not yet known
type FileTypeUnknown struct {
	tdCommon
}

// MessageType return the string telegram-type of FileTypeUnknown
func (fileTypeUnknown *FileTypeUnknown) MessageType() string {
	return "fileTypeUnknown"
}

// NewFileTypeUnknown creates a new FileTypeUnknown
//
func NewFileTypeUnknown() *FileTypeUnknown {
	fileTypeUnknownTemp := FileTypeUnknown{
		tdCommon: tdCommon{Type: "fileTypeUnknown"},
	}

	return &fileTypeUnknownTemp
}

// GetFileTypeEnum return the enum type of this object
func (fileTypeUnknown *FileTypeUnknown) GetFileTypeEnum() FileTypeEnum {
	return FileTypeUnknownType
}

// FileTypeVideo The file is a video
type FileTypeVideo struct {
	tdCommon
}

// MessageType return the string telegram-type of FileTypeVideo
func (fileTypeVideo *FileTypeVideo) MessageType() string {
	return "fileTypeVideo"
}

// NewFileTypeVideo creates a new FileTypeVideo
//
func NewFileTypeVideo() *FileTypeVideo {
	fileTypeVideoTemp := FileTypeVideo{
		tdCommon: tdCommon{Type: "fileTypeVideo"},
	}

	return &fileTypeVideoTemp
}

// GetFileTypeEnum return the enum type of this object
func (fileTypeVideo *FileTypeVideo) GetFileTypeEnum() FileTypeEnum {
	return FileTypeVideoType
}

// FileTypeVideoNote The file is a video note
type FileTypeVideoNote struct {
	tdCommon
}

// MessageType return the string telegram-type of FileTypeVideoNote
func (fileTypeVideoNote *FileTypeVideoNote) MessageType() string {
	return "fileTypeVideoNote"
}

// NewFileTypeVideoNote creates a new FileTypeVideoNote
//
func NewFileTypeVideoNote() *FileTypeVideoNote {
	fileTypeVideoNoteTemp := FileTypeVideoNote{
		tdCommon: tdCommon{Type: "fileTypeVideoNote"},
	}

	return &fileTypeVideoNoteTemp
}

// GetFileTypeEnum return the enum type of this object
func (fileTypeVideoNote *FileTypeVideoNote) GetFileTypeEnum() FileTypeEnum {
	return FileTypeVideoNoteType
}

// FileTypeVoiceNote The file is a voice note
type FileTypeVoiceNote struct {
	tdCommon
}

// MessageType return the string telegram-type of FileTypeVoiceNote
func (fileTypeVoiceNote *FileTypeVoiceNote) MessageType() string {
	return "fileTypeVoiceNote"
}

// NewFileTypeVoiceNote creates a new FileTypeVoiceNote
//
func NewFileTypeVoiceNote() *FileTypeVoiceNote {
	fileTypeVoiceNoteTemp := FileTypeVoiceNote{
		tdCommon: tdCommon{Type: "fileTypeVoiceNote"},
	}

	return &fileTypeVoiceNoteTemp
}

// GetFileTypeEnum return the enum type of this object
func (fileTypeVoiceNote *FileTypeVoiceNote) GetFileTypeEnum() FileTypeEnum {
	return FileTypeVoiceNoteType
}

// FileTypeWallpaper The file is a wallpaper or a background pattern
type FileTypeWallpaper struct {
	tdCommon
}

// MessageType return the string telegram-type of FileTypeWallpaper
func (fileTypeWallpaper *FileTypeWallpaper) MessageType() string {
	return "fileTypeWallpaper"
}

// NewFileTypeWallpaper creates a new FileTypeWallpaper
//
func NewFileTypeWallpaper() *FileTypeWallpaper {
	fileTypeWallpaperTemp := FileTypeWallpaper{
		tdCommon: tdCommon{Type: "fileTypeWallpaper"},
	}

	return &fileTypeWallpaperTemp
}

// GetFileTypeEnum return the enum type of this object
func (fileTypeWallpaper *FileTypeWallpaper) GetFileTypeEnum() FileTypeEnum {
	return FileTypeWallpaperType
}
