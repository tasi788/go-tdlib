package tdlib

import (
	"encoding/json"
	"fmt"
)

// StickerSets Represents a list of sticker sets
type StickerSets struct {
	tdCommon
	TotalCount int32            `json:"total_count"` // Approximate total number of sticker sets found
	Sets       []StickerSetInfo `json:"sets"`        // List of sticker sets
}

// MessageType return the string telegram-type of StickerSets
func (stickerSets *StickerSets) MessageType() string {
	return "stickerSets"
}

// NewStickerSets creates a new StickerSets
//
// @param totalCount Approximate total number of sticker sets found
// @param sets List of sticker sets
func NewStickerSets(totalCount int32, sets []StickerSetInfo) *StickerSets {
	stickerSetsTemp := StickerSets{
		tdCommon:   tdCommon{Type: "stickerSets"},
		TotalCount: totalCount,
		Sets:       sets,
	}

	return &stickerSetsTemp
}

// GetInstalledStickerSets Returns a list of installed sticker sets
// @param isMasks Pass true to return mask sticker sets; pass false to return ordinary sticker sets
func (client *Client) GetInstalledStickerSets(isMasks bool) (*StickerSets, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":    "getInstalledStickerSets",
		"is_masks": isMasks,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var stickerSets StickerSets
	err = json.Unmarshal(result.Raw, &stickerSets)
	return &stickerSets, err
}

// GetArchivedStickerSets Returns a list of archived sticker sets
// @param isMasks Pass true to return mask stickers sets; pass false to return ordinary sticker sets
// @param offsetStickerSetId Identifier of the sticker set from which to return the result
// @param limit The maximum number of sticker sets to return; up to 100
func (client *Client) GetArchivedStickerSets(isMasks bool, offsetStickerSetId JSONInt64, limit int32) (*StickerSets, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":                 "getArchivedStickerSets",
		"is_masks":              isMasks,
		"offset_sticker_set_id": offsetStickerSetId,
		"limit":                 limit,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var stickerSets StickerSets
	err = json.Unmarshal(result.Raw, &stickerSets)
	return &stickerSets, err
}

// GetTrendingStickerSets Returns a list of trending sticker sets. For optimal performance, the number of returned sticker sets is chosen by TDLib
// @param offset The offset from which to return the sticker sets; must be non-negative
// @param limit The maximum number of sticker sets to be returned; up to 100. For optimal performance, the number of returned sticker sets is chosen by TDLib and can be smaller than the specified limit, even if the end of the list has not been reached
func (client *Client) GetTrendingStickerSets(offset int32, limit int32) (*StickerSets, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":  "getTrendingStickerSets",
		"offset": offset,
		"limit":  limit,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var stickerSets StickerSets
	err = json.Unmarshal(result.Raw, &stickerSets)
	return &stickerSets, err
}

// GetAttachedStickerSets Returns a list of sticker sets attached to a file. Currently, only photos and videos can have attached sticker sets
// @param fileId File identifier
func (client *Client) GetAttachedStickerSets(fileId int32) (*StickerSets, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "getAttachedStickerSets",
		"file_id": fileId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var stickerSets StickerSets
	err = json.Unmarshal(result.Raw, &stickerSets)
	return &stickerSets, err
}

// SearchInstalledStickerSets Searches for installed sticker sets by looking for specified query in their title and name
// @param isMasks Pass true to return mask sticker sets; pass false to return ordinary sticker sets
// @param query Query to search for
// @param limit The maximum number of sticker sets to return
func (client *Client) SearchInstalledStickerSets(isMasks bool, query string, limit int32) (*StickerSets, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":    "searchInstalledStickerSets",
		"is_masks": isMasks,
		"query":    query,
		"limit":    limit,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var stickerSets StickerSets
	err = json.Unmarshal(result.Raw, &stickerSets)
	return &stickerSets, err
}

// SearchStickerSets Searches for ordinary sticker sets by looking for specified query in their title and name. Excludes installed sticker sets from the results
// @param query Query to search for
func (client *Client) SearchStickerSets(query string) (*StickerSets, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "searchStickerSets",
		"query": query,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var stickerSets StickerSets
	err = json.Unmarshal(result.Raw, &stickerSets)
	return &stickerSets, err
}
