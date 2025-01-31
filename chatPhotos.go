package tdlib

import (
	"encoding/json"
	"fmt"
)

// ChatPhotos Contains a list of chat or user profile photos
type ChatPhotos struct {
	tdCommon
	TotalCount int32       `json:"total_count"` // Total number of photos
	Photos     []ChatPhoto `json:"photos"`      // List of photos
}

// MessageType return the string telegram-type of ChatPhotos
func (chatPhotos *ChatPhotos) MessageType() string {
	return "chatPhotos"
}

// NewChatPhotos creates a new ChatPhotos
//
// @param totalCount Total number of photos
// @param photos List of photos
func NewChatPhotos(totalCount int32, photos []ChatPhoto) *ChatPhotos {
	chatPhotosTemp := ChatPhotos{
		tdCommon:   tdCommon{Type: "chatPhotos"},
		TotalCount: totalCount,
		Photos:     photos,
	}

	return &chatPhotosTemp
}

// GetUserProfilePhotos Returns the profile photos of a user. The result of this query may be outdated: some photos might have been deleted already
// @param userId User identifier
// @param offset The number of photos to skip; must be non-negative
// @param limit The maximum number of photos to be returned; up to 100
func (client *Client) GetUserProfilePhotos(userId int64, offset int32, limit int32) (*ChatPhotos, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "getUserProfilePhotos",
		"user_id": userId,
		"offset":  offset,
		"limit":   limit,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var chatPhotos ChatPhotos
	err = json.Unmarshal(result.Raw, &chatPhotos)
	return &chatPhotos, err
}
