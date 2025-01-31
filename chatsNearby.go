package tdlib

import (
	"encoding/json"
	"fmt"
)

// ChatsNearby Represents a list of chats located nearby
type ChatsNearby struct {
	tdCommon
	UsersNearby       []ChatNearby `json:"users_nearby"`       // List of users nearby
	SupergroupsNearby []ChatNearby `json:"supergroups_nearby"` // List of location-based supergroups nearby
}

// MessageType return the string telegram-type of ChatsNearby
func (chatsNearby *ChatsNearby) MessageType() string {
	return "chatsNearby"
}

// NewChatsNearby creates a new ChatsNearby
//
// @param usersNearby List of users nearby
// @param supergroupsNearby List of location-based supergroups nearby
func NewChatsNearby(usersNearby []ChatNearby, supergroupsNearby []ChatNearby) *ChatsNearby {
	chatsNearbyTemp := ChatsNearby{
		tdCommon:          tdCommon{Type: "chatsNearby"},
		UsersNearby:       usersNearby,
		SupergroupsNearby: supergroupsNearby,
	}

	return &chatsNearbyTemp
}

// SearchChatsNearby Returns a list of users and location-based supergroups nearby. The list of users nearby will be updated for 60 seconds after the request by the updates updateUsersNearby. The request must be sent again every 25 seconds with adjusted location to not miss new chats
// @param location Current user location
func (client *Client) SearchChatsNearby(location *Location) (*ChatsNearby, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":    "searchChatsNearby",
		"location": location,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var chatsNearby ChatsNearby
	err = json.Unmarshal(result.Raw, &chatsNearby)
	return &chatsNearby, err
}
