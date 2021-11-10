package tdlib

import (
	"encoding/json"
	"fmt"
)

// ChatJoinRequests Contains a list of chat join requests
type ChatJoinRequests struct {
	tdCommon
	TotalCount int32             `json:"total_count"` // Approximate total count of requests found
	Requests   []ChatJoinRequest `json:"requests"`    // List of the requests
}

// MessageType return the string telegram-type of ChatJoinRequests
func (chatJoinRequests *ChatJoinRequests) MessageType() string {
	return "chatJoinRequests"
}

// NewChatJoinRequests creates a new ChatJoinRequests
//
// @param totalCount Approximate total count of requests found
// @param requests List of the requests
func NewChatJoinRequests(totalCount int32, requests []ChatJoinRequest) *ChatJoinRequests {
	chatJoinRequestsTemp := ChatJoinRequests{
		tdCommon:   tdCommon{Type: "chatJoinRequests"},
		TotalCount: totalCount,
		Requests:   requests,
	}

	return &chatJoinRequestsTemp
}

// GetChatJoinRequests Returns pending join requests in a chat
// @param chatId Chat identifier
// @param inviteLink Invite link for which to return join requests. If empty, all join requests will be returned. Requires administrator privileges and can_invite_users right in the chat for own links and owner privileges for other links
// @param query A query to search for in the first names, last names and usernames of the users to return
// @param offsetRequest A chat join request from which to return next requests; pass null to get results from the beginning
// @param limit The maximum number of chat join requests to return
func (client *Client) GetChatJoinRequests(chatId int64, inviteLink string, query string, offsetRequest *ChatJoinRequest, limit int32) (*ChatJoinRequests, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":          "getChatJoinRequests",
		"chat_id":        chatId,
		"invite_link":    inviteLink,
		"query":          query,
		"offset_request": offsetRequest,
		"limit":          limit,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var chatJoinRequests ChatJoinRequests
	err = json.Unmarshal(result.Raw, &chatJoinRequests)
	return &chatJoinRequests, err
}
