package tdlib

import (
	"encoding/json"
	"fmt"
)

// ChatMembers Contains a list of chat members
type ChatMembers struct {
	tdCommon
	TotalCount int32        `json:"total_count"` // Approximate total count of chat members found
	Members    []ChatMember `json:"members"`     // A list of chat members
}

// MessageType return the string telegram-type of ChatMembers
func (chatMembers *ChatMembers) MessageType() string {
	return "chatMembers"
}

// NewChatMembers creates a new ChatMembers
//
// @param totalCount Approximate total count of chat members found
// @param members A list of chat members
func NewChatMembers(totalCount int32, members []ChatMember) *ChatMembers {
	chatMembersTemp := ChatMembers{
		tdCommon:   tdCommon{Type: "chatMembers"},
		TotalCount: totalCount,
		Members:    members,
	}

	return &chatMembersTemp
}

// SearchChatMembers Searches for a specified query in the first name, last name and username of the members of a specified chat. Requires administrator rights in channels
// @param chatId Chat identifier
// @param query Query to search for
// @param limit The maximum number of users to be returned; up to 200
// @param filter The type of users to search for; pass null to search among all chat members
func (client *Client) SearchChatMembers(chatId int64, query string, limit int32, filter ChatMembersFilter) (*ChatMembers, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "searchChatMembers",
		"chat_id": chatId,
		"query":   query,
		"limit":   limit,
		"filter":  filter,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var chatMembers ChatMembers
	err = json.Unmarshal(result.Raw, &chatMembers)
	return &chatMembers, err
}

// GetSupergroupMembers Returns information about members or banned users in a supergroup or channel. Can be used only if supergroupFullInfo.can_get_members == true; additionally, administrator privileges may be required for some filters
// @param supergroupId Identifier of the supergroup or channel
// @param filter The type of users to return; pass null to use supergroupMembersFilterRecent
// @param offset Number of users to skip
// @param limit The maximum number of users be returned; up to 200
func (client *Client) GetSupergroupMembers(supergroupId int64, filter SupergroupMembersFilter, offset int32, limit int32) (*ChatMembers, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":         "getSupergroupMembers",
		"supergroup_id": supergroupId,
		"filter":        filter,
		"offset":        offset,
		"limit":         limit,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var chatMembers ChatMembers
	err = json.Unmarshal(result.Raw, &chatMembers)
	return &chatMembers, err
}
