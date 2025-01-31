package tdlib

import (
	"encoding/json"
	"fmt"
)

// ChatFilter Represents a filter of user chats
type ChatFilter struct {
	tdCommon
	Title              string  `json:"title"`                // The title of the filter; 1-12 characters without line feeds
	IconName           string  `json:"icon_name"`            // The chosen icon name for short filter representation. If non-empty, must be one of "All", "Unread", "Unmuted", "Bots", "Channels", "Groups", "Private", "Custom", "Setup", "Cat", "Crown", "Favorite", "Flower", "Game", "Home", "Love", "Mask", "Party", "Sport", "Study", "Trade", "Travel", "Work". If empty, use getChatFilterDefaultIconName to get default icon name for the filter
	PinnedChatIds      []int64 `json:"pinned_chat_ids"`      // The chat identifiers of pinned chats in the filtered chat list
	IncludedChatIds    []int64 `json:"included_chat_ids"`    // The chat identifiers of always included chats in the filtered chat list
	ExcludedChatIds    []int64 `json:"excluded_chat_ids"`    // The chat identifiers of always excluded chats in the filtered chat list
	ExcludeMuted       bool    `json:"exclude_muted"`        // True, if muted chats need to be excluded
	ExcludeRead        bool    `json:"exclude_read"`         // True, if read chats need to be excluded
	ExcludeArchived    bool    `json:"exclude_archived"`     // True, if archived chats need to be excluded
	IncludeContacts    bool    `json:"include_contacts"`     // True, if contacts need to be included
	IncludeNonContacts bool    `json:"include_non_contacts"` // True, if non-contact users need to be included
	IncludeBots        bool    `json:"include_bots"`         // True, if bots need to be included
	IncludeGroups      bool    `json:"include_groups"`       // True, if basic groups and supergroups need to be included
	IncludeChannels    bool    `json:"include_channels"`     // True, if channels need to be included
}

// MessageType return the string telegram-type of ChatFilter
func (chatFilter *ChatFilter) MessageType() string {
	return "chatFilter"
}

// NewChatFilter creates a new ChatFilter
//
// @param title The title of the filter; 1-12 characters without line feeds
// @param iconName The chosen icon name for short filter representation. If non-empty, must be one of "All", "Unread", "Unmuted", "Bots", "Channels", "Groups", "Private", "Custom", "Setup", "Cat", "Crown", "Favorite", "Flower", "Game", "Home", "Love", "Mask", "Party", "Sport", "Study", "Trade", "Travel", "Work". If empty, use getChatFilterDefaultIconName to get default icon name for the filter
// @param pinnedChatIds The chat identifiers of pinned chats in the filtered chat list
// @param includedChatIds The chat identifiers of always included chats in the filtered chat list
// @param excludedChatIds The chat identifiers of always excluded chats in the filtered chat list
// @param excludeMuted True, if muted chats need to be excluded
// @param excludeRead True, if read chats need to be excluded
// @param excludeArchived True, if archived chats need to be excluded
// @param includeContacts True, if contacts need to be included
// @param includeNonContacts True, if non-contact users need to be included
// @param includeBots True, if bots need to be included
// @param includeGroups True, if basic groups and supergroups need to be included
// @param includeChannels True, if channels need to be included
func NewChatFilter(title string, iconName string, pinnedChatIds []int64, includedChatIds []int64, excludedChatIds []int64, excludeMuted bool, excludeRead bool, excludeArchived bool, includeContacts bool, includeNonContacts bool, includeBots bool, includeGroups bool, includeChannels bool) *ChatFilter {
	chatFilterTemp := ChatFilter{
		tdCommon:           tdCommon{Type: "chatFilter"},
		Title:              title,
		IconName:           iconName,
		PinnedChatIds:      pinnedChatIds,
		IncludedChatIds:    includedChatIds,
		ExcludedChatIds:    excludedChatIds,
		ExcludeMuted:       excludeMuted,
		ExcludeRead:        excludeRead,
		ExcludeArchived:    excludeArchived,
		IncludeContacts:    includeContacts,
		IncludeNonContacts: includeNonContacts,
		IncludeBots:        includeBots,
		IncludeGroups:      includeGroups,
		IncludeChannels:    includeChannels,
	}

	return &chatFilterTemp
}

// GetChatFilter Returns information about a chat filter by its identifier
// @param chatFilterId Chat filter identifier
func (client *Client) GetChatFilter(chatFilterId int32) (*ChatFilter, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":          "getChatFilter",
		"chat_filter_id": chatFilterId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var chatFilterDummy ChatFilter
	err = json.Unmarshal(result.Raw, &chatFilterDummy)
	return &chatFilterDummy, err
}
