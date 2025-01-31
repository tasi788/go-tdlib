package tdlib

import (
	"encoding/json"
	"fmt"
)

// Chats Represents a list of chats
type Chats struct {
	tdCommon
	TotalCount int32   `json:"total_count"` // Approximate total count of chats found
	ChatIds    []int64 `json:"chat_ids"`    // List of chat identifiers
}

// MessageType return the string telegram-type of Chats
func (chats *Chats) MessageType() string {
	return "chats"
}

// NewChats creates a new Chats
//
// @param totalCount Approximate total count of chats found
// @param chatIds List of chat identifiers
func NewChats(totalCount int32, chatIds []int64) *Chats {
	chatsTemp := Chats{
		tdCommon:   tdCommon{Type: "chats"},
		TotalCount: totalCount,
		ChatIds:    chatIds,
	}

	return &chatsTemp
}

// GetChats Returns an ordered list of chats from the beginning of a chat list. For informational purposes only. Use loadChats and updates processing instead to maintain chat lists in a consistent state
// @param chatList The chat list in which to return chats; pass null to get chats from the main chat list
// @param limit The maximum number of chats to be returned
func (client *Client) GetChats(chatList ChatList, limit int32) (*Chats, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":     "getChats",
		"chat_list": chatList,
		"limit":     limit,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var chats Chats
	err = json.Unmarshal(result.Raw, &chats)
	return &chats, err
}

// SearchPublicChats Searches public chats by looking for specified query in their username and title. Currently, only private chats, supergroups and channels can be public. Returns a meaningful number of results. Excludes private chats with contacts and chats from the chat list from the results
// @param query Query to search for
func (client *Client) SearchPublicChats(query string) (*Chats, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "searchPublicChats",
		"query": query,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var chats Chats
	err = json.Unmarshal(result.Raw, &chats)
	return &chats, err
}

// SearchChats Searches for the specified query in the title and username of already known chats, this is an offline request. Returns chats in the order seen in the main chat list
// @param query Query to search for. If the query is empty, returns up to 50 recently found chats
// @param limit The maximum number of chats to be returned
func (client *Client) SearchChats(query string, limit int32) (*Chats, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "searchChats",
		"query": query,
		"limit": limit,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var chats Chats
	err = json.Unmarshal(result.Raw, &chats)
	return &chats, err
}

// SearchChatsOnServer Searches for the specified query in the title and username of already known chats via request to the server. Returns chats in the order seen in the main chat list
// @param query Query to search for
// @param limit The maximum number of chats to be returned
func (client *Client) SearchChatsOnServer(query string, limit int32) (*Chats, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "searchChatsOnServer",
		"query": query,
		"limit": limit,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var chats Chats
	err = json.Unmarshal(result.Raw, &chats)
	return &chats, err
}

// GetTopChats Returns a list of frequently used chats. Supported only if the chat info database is enabled
// @param category Category of chats to be returned
// @param limit The maximum number of chats to be returned; up to 30
func (client *Client) GetTopChats(category TopChatCategory, limit int32) (*Chats, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":    "getTopChats",
		"category": category,
		"limit":    limit,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var chats Chats
	err = json.Unmarshal(result.Raw, &chats)
	return &chats, err
}

// GetRecentlyOpenedChats Returns recently opened chats, this is an offline request. Returns chats in the order of last opening
// @param limit The maximum number of chats to be returned
func (client *Client) GetRecentlyOpenedChats(limit int32) (*Chats, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "getRecentlyOpenedChats",
		"limit": limit,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var chats Chats
	err = json.Unmarshal(result.Raw, &chats)
	return &chats, err
}

// GetCreatedPublicChats Returns a list of public chats of the specified type, owned by the user
// @param typeParam Type of the public chats to return
func (client *Client) GetCreatedPublicChats(typeParam PublicChatType) (*Chats, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "getCreatedPublicChats",
		"type":  typeParam,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var chats Chats
	err = json.Unmarshal(result.Raw, &chats)
	return &chats, err
}

// GetSuitableDiscussionChats Returns a list of basic group and supergroup chats, which can be used as a discussion group for a channel. Returned basic group chats must be first upgraded to supergroups before they can be set as a discussion group. To set a returned supergroup as a discussion group, access to its old messages must be enabled using toggleSupergroupIsAllHistoryAvailable first
func (client *Client) GetSuitableDiscussionChats() (*Chats, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "getSuitableDiscussionChats",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var chats Chats
	err = json.Unmarshal(result.Raw, &chats)
	return &chats, err
}

// GetInactiveSupergroupChats Returns a list of recently inactive supergroups and channels. Can be used when user reaches limit on the number of joined supergroups and channels and receives CHANNELS_TOO_MUCH error
func (client *Client) GetInactiveSupergroupChats() (*Chats, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "getInactiveSupergroupChats",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var chats Chats
	err = json.Unmarshal(result.Raw, &chats)
	return &chats, err
}

// GetGroupsInCommon Returns a list of common group chats with a given user. Chats are sorted by their type and creation date
// @param userId User identifier
// @param offsetChatId Chat identifier starting from which to return chats; use 0 for the first request
// @param limit The maximum number of chats to be returned; up to 100
func (client *Client) GetGroupsInCommon(userId int64, offsetChatId int64, limit int32) (*Chats, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":          "getGroupsInCommon",
		"user_id":        userId,
		"offset_chat_id": offsetChatId,
		"limit":          limit,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var chats Chats
	err = json.Unmarshal(result.Raw, &chats)
	return &chats, err
}

// GetChatNotificationSettingsExceptions Returns list of chats with non-default notification settings
// @param scope If specified, only chats from the scope will be returned; pass null to return chats from all scopes
// @param compareSound If true, also chats with non-default sound will be returned
func (client *Client) GetChatNotificationSettingsExceptions(scope NotificationSettingsScope, compareSound bool) (*Chats, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":         "getChatNotificationSettingsExceptions",
		"scope":         scope,
		"compare_sound": compareSound,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var chats Chats
	err = json.Unmarshal(result.Raw, &chats)
	return &chats, err
}
