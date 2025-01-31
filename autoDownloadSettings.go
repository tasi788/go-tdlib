package tdlib

// AutoDownloadSettings Contains auto-download settings
type AutoDownloadSettings struct {
	tdCommon
	IsAutoDownloadEnabled bool  `json:"is_auto_download_enabled"` // True, if the auto-download is enabled
	MaxPhotoFileSize      int32 `json:"max_photo_file_size"`      // The maximum size of a photo file to be auto-downloaded, in bytes
	MaxVideoFileSize      int32 `json:"max_video_file_size"`      // The maximum size of a video file to be auto-downloaded, in bytes
	MaxOtherFileSize      int32 `json:"max_other_file_size"`      // The maximum size of other file types to be auto-downloaded, in bytes
	VideoUploadBitrate    int32 `json:"video_upload_bitrate"`     // The maximum suggested bitrate for uploaded videos, in kbit/s
	PreloadLargeVideos    bool  `json:"preload_large_videos"`     // True, if the beginning of video files needs to be preloaded for instant playback
	PreloadNextAudio      bool  `json:"preload_next_audio"`       // True, if the next audio track needs to be preloaded while the user is listening to an audio file
	UseLessDataForCalls   bool  `json:"use_less_data_for_calls"`  // True, if "use less data for calls" option needs to be enabled
}

// MessageType return the string telegram-type of AutoDownloadSettings
func (autoDownloadSettings *AutoDownloadSettings) MessageType() string {
	return "autoDownloadSettings"
}

// NewAutoDownloadSettings creates a new AutoDownloadSettings
//
// @param isAutoDownloadEnabled True, if the auto-download is enabled
// @param maxPhotoFileSize The maximum size of a photo file to be auto-downloaded, in bytes
// @param maxVideoFileSize The maximum size of a video file to be auto-downloaded, in bytes
// @param maxOtherFileSize The maximum size of other file types to be auto-downloaded, in bytes
// @param videoUploadBitrate The maximum suggested bitrate for uploaded videos, in kbit/s
// @param preloadLargeVideos True, if the beginning of video files needs to be preloaded for instant playback
// @param preloadNextAudio True, if the next audio track needs to be preloaded while the user is listening to an audio file
// @param useLessDataForCalls True, if "use less data for calls" option needs to be enabled
func NewAutoDownloadSettings(isAutoDownloadEnabled bool, maxPhotoFileSize int32, maxVideoFileSize int32, maxOtherFileSize int32, videoUploadBitrate int32, preloadLargeVideos bool, preloadNextAudio bool, useLessDataForCalls bool) *AutoDownloadSettings {
	autoDownloadSettingsTemp := AutoDownloadSettings{
		tdCommon:              tdCommon{Type: "autoDownloadSettings"},
		IsAutoDownloadEnabled: isAutoDownloadEnabled,
		MaxPhotoFileSize:      maxPhotoFileSize,
		MaxVideoFileSize:      maxVideoFileSize,
		MaxOtherFileSize:      maxOtherFileSize,
		VideoUploadBitrate:    videoUploadBitrate,
		PreloadLargeVideos:    preloadLargeVideos,
		PreloadNextAudio:      preloadNextAudio,
		UseLessDataForCalls:   useLessDataForCalls,
	}

	return &autoDownloadSettingsTemp
}
