package tdlib

import (
	"encoding/json"
	"fmt"
)

// StickerSet Represents a sticker set
type StickerSet struct {
	tdCommon
	Id               JSONInt64          `json:"id"`                // Identifier of the sticker set
	Title            string             `json:"title"`             // Title of the sticker set
	Name             string             `json:"name"`              // Name of the sticker set
	Thumbnail        *Thumbnail         `json:"thumbnail"`         // Sticker set thumbnail in WEBP or TGS format with width and height 100; may be null. The file can be downloaded only before the thumbnail is changed
	ThumbnailOutline []ClosedVectorPath `json:"thumbnail_outline"` // Sticker set thumbnail's outline represented as a list of closed vector paths; may be empty. The coordinate system origin is in the upper-left corner
	IsInstalled      bool               `json:"is_installed"`      // True, if the sticker set has been installed by the current user
	IsArchived       bool               `json:"is_archived"`       // True, if the sticker set has been archived. A sticker set can't be installed and archived simultaneously
	IsOfficial       bool               `json:"is_official"`       // True, if the sticker set is official
	IsAnimated       bool               `json:"is_animated"`       // True, is the stickers in the set are animated
	IsMasks          bool               `json:"is_masks"`          // True, if the stickers in the set are masks
	IsViewed         bool               `json:"is_viewed"`         // True for already viewed trending sticker sets
	Stickers         []Sticker          `json:"stickers"`          // List of stickers in this set
	Emojis           []Emojis           `json:"emojis"`            // A list of emoji corresponding to the stickers in the same order. The list is only for informational purposes, because a sticker is always sent with a fixed emoji from the corresponding Sticker object
}

// MessageType return the string telegram-type of StickerSet
func (stickerSet *StickerSet) MessageType() string {
	return "stickerSet"
}

// NewStickerSet creates a new StickerSet
//
// @param id Identifier of the sticker set
// @param title Title of the sticker set
// @param name Name of the sticker set
// @param thumbnail Sticker set thumbnail in WEBP or TGS format with width and height 100; may be null. The file can be downloaded only before the thumbnail is changed
// @param thumbnailOutline Sticker set thumbnail's outline represented as a list of closed vector paths; may be empty. The coordinate system origin is in the upper-left corner
// @param isInstalled True, if the sticker set has been installed by the current user
// @param isArchived True, if the sticker set has been archived. A sticker set can't be installed and archived simultaneously
// @param isOfficial True, if the sticker set is official
// @param isAnimated True, is the stickers in the set are animated
// @param isMasks True, if the stickers in the set are masks
// @param isViewed True for already viewed trending sticker sets
// @param stickers List of stickers in this set
// @param emojis A list of emoji corresponding to the stickers in the same order. The list is only for informational purposes, because a sticker is always sent with a fixed emoji from the corresponding Sticker object
func NewStickerSet(id JSONInt64, title string, name string, thumbnail *Thumbnail, thumbnailOutline []ClosedVectorPath, isInstalled bool, isArchived bool, isOfficial bool, isAnimated bool, isMasks bool, isViewed bool, stickers []Sticker, emojis []Emojis) *StickerSet {
	stickerSetTemp := StickerSet{
		tdCommon:         tdCommon{Type: "stickerSet"},
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
		Stickers:         stickers,
		Emojis:           emojis,
	}

	return &stickerSetTemp
}

// GetStickerSet Returns information about a sticker set by its identifier
// @param setId Identifier of the sticker set
func (client *Client) GetStickerSet(setId JSONInt64) (*StickerSet, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":  "getStickerSet",
		"set_id": setId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var stickerSet StickerSet
	err = json.Unmarshal(result.Raw, &stickerSet)
	return &stickerSet, err
}

// SearchStickerSet Searches for a sticker set by its name
// @param name Name of the sticker set
func (client *Client) SearchStickerSet(name string) (*StickerSet, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "searchStickerSet",
		"name":  name,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var stickerSet StickerSet
	err = json.Unmarshal(result.Raw, &stickerSet)
	return &stickerSet, err
}

// CreateNewStickerSet Creates a new sticker set. Returns the newly created sticker set
// @param userId Sticker set owner; ignored for regular users
// @param title Sticker set title; 1-64 characters
// @param name Sticker set name. Can contain only English letters, digits and underscores. Must end with *"_by_<bot username>"* (*<bot_username>* is case insensitive) for bots; 1-64 characters
// @param isMasks True, if stickers are masks. Animated stickers can't be masks
// @param stickers List of stickers to be added to the set; must be non-empty. All stickers must be of the same type. For animated stickers, uploadStickerFile must be used before the sticker is shown
// @param source Source of the sticker set; may be empty if unknown
func (client *Client) CreateNewStickerSet(userId int64, title string, name string, isMasks bool, stickers []InputSticker, source string) (*StickerSet, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":    "createNewStickerSet",
		"user_id":  userId,
		"title":    title,
		"name":     name,
		"is_masks": isMasks,
		"stickers": stickers,
		"source":   source,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var stickerSet StickerSet
	err = json.Unmarshal(result.Raw, &stickerSet)
	return &stickerSet, err
}

// AddStickerToSet Adds a new sticker to a set; for bots only. Returns the sticker set
// @param userId Sticker set owner
// @param name Sticker set name
// @param sticker Sticker to add to the set
func (client *Client) AddStickerToSet(userId int64, name string, sticker InputSticker) (*StickerSet, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "addStickerToSet",
		"user_id": userId,
		"name":    name,
		"sticker": sticker,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var stickerSet StickerSet
	err = json.Unmarshal(result.Raw, &stickerSet)
	return &stickerSet, err
}

// SetStickerSetThumbnail Sets a sticker set thumbnail; for bots only. Returns the sticker set
// @param userId Sticker set owner
// @param name Sticker set name
// @param thumbnail Thumbnail to set in PNG or TGS format; pass null to remove the sticker set thumbnail. Animated thumbnail must be set for animated sticker sets and only for them
func (client *Client) SetStickerSetThumbnail(userId int64, name string, thumbnail InputFile) (*StickerSet, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":     "setStickerSetThumbnail",
		"user_id":   userId,
		"name":      name,
		"thumbnail": thumbnail,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var stickerSet StickerSet
	err = json.Unmarshal(result.Raw, &stickerSet)
	return &stickerSet, err
}
