package tdlib

import (
	"encoding/json"
	"fmt"
)

// ChatFilterInfo Contains basic information about a chat filter
type ChatFilterInfo struct {
	tdCommon
	Id       int32  `json:"id"`        // Unique chat filter identifier
	Title    string `json:"title"`     // The title of the filter; 1-12 characters without line feeds
	IconName string `json:"icon_name"` // The chosen or default icon name for short filter representation. One of "All", "Unread", "Unmuted", "Bots", "Channels", "Groups", "Private", "Custom", "Setup", "Cat", "Crown", "Favorite", "Flower", "Game", "Home", "Love", "Mask", "Party", "Sport", "Study", "Trade", "Travel", "Work"
}

// MessageType return the string telegram-type of ChatFilterInfo
func (chatFilterInfo *ChatFilterInfo) MessageType() string {
	return "chatFilterInfo"
}

// NewChatFilterInfo creates a new ChatFilterInfo
//
// @param id Unique chat filter identifier
// @param title The title of the filter; 1-12 characters without line feeds
// @param iconName The chosen or default icon name for short filter representation. One of "All", "Unread", "Unmuted", "Bots", "Channels", "Groups", "Private", "Custom", "Setup", "Cat", "Crown", "Favorite", "Flower", "Game", "Home", "Love", "Mask", "Party", "Sport", "Study", "Trade", "Travel", "Work"
func NewChatFilterInfo(id int32, title string, iconName string) *ChatFilterInfo {
	chatFilterInfoTemp := ChatFilterInfo{
		tdCommon: tdCommon{Type: "chatFilterInfo"},
		Id:       id,
		Title:    title,
		IconName: iconName,
	}

	return &chatFilterInfoTemp
}

// CreateChatFilter Creates new chat filter. Returns information about the created chat filter
// @param filter Chat filter
func (client *Client) CreateChatFilter(filter *ChatFilter) (*ChatFilterInfo, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":  "createChatFilter",
		"filter": filter,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var chatFilterInfo ChatFilterInfo
	err = json.Unmarshal(result.Raw, &chatFilterInfo)
	return &chatFilterInfo, err
}

// EditChatFilter Edits existing chat filter. Returns information about the edited chat filter
// @param chatFilterId Chat filter identifier
// @param filter The edited chat filter
func (client *Client) EditChatFilter(chatFilterId int32, filter *ChatFilter) (*ChatFilterInfo, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":          "editChatFilter",
		"chat_filter_id": chatFilterId,
		"filter":         filter,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var chatFilterInfo ChatFilterInfo
	err = json.Unmarshal(result.Raw, &chatFilterInfo)
	return &chatFilterInfo, err
}
