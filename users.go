package tdlib

import (
	"encoding/json"
	"fmt"
)

// Users Represents a list of users
type Users struct {
	tdCommon
	TotalCount int32   `json:"total_count"` // Approximate total count of users found
	UserIds    []int64 `json:"user_ids"`    // A list of user identifiers
}

// MessageType return the string telegram-type of Users
func (users *Users) MessageType() string {
	return "users"
}

// NewUsers creates a new Users
//
// @param totalCount Approximate total count of users found
// @param userIds A list of user identifiers
func NewUsers(totalCount int32, userIds []int64) *Users {
	usersTemp := Users{
		tdCommon:   tdCommon{Type: "users"},
		TotalCount: totalCount,
		UserIds:    userIds,
	}

	return &usersTemp
}

// GetMessageViewers Returns viewers of a recent outgoing message in a basic group or a supergroup chat. For video notes and voice notes only users, opened content of the message, are returned. The method can be called if message.can_get_viewers == true
// @param chatId Chat identifier
// @param messageId Identifier of the message
func (client *Client) GetMessageViewers(chatId int64, messageId int64) (*Users, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":      "getMessageViewers",
		"chat_id":    chatId,
		"message_id": messageId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var users Users
	err = json.Unmarshal(result.Raw, &users)
	return &users, err
}

// GetPollVoters Returns users voted for the specified option in a non-anonymous polls. For optimal performance, the number of returned users is chosen by TDLib
// @param chatId Identifier of the chat to which the poll belongs
// @param messageId Identifier of the message containing the poll
// @param optionId 0-based identifier of the answer option
// @param offset Number of users to skip in the result; must be non-negative
// @param limit The maximum number of users to be returned; must be positive and can't be greater than 50. For optimal performance, the number of returned users is chosen by TDLib and can be smaller than the specified limit, even if the end of the voter list has not been reached
func (client *Client) GetPollVoters(chatId int64, messageId int64, optionId int32, offset int32, limit int32) (*Users, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":      "getPollVoters",
		"chat_id":    chatId,
		"message_id": messageId,
		"option_id":  optionId,
		"offset":     offset,
		"limit":      limit,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var users Users
	err = json.Unmarshal(result.Raw, &users)
	return &users, err
}

// GetContacts Returns all user contacts
func (client *Client) GetContacts() (*Users, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "getContacts",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var users Users
	err = json.Unmarshal(result.Raw, &users)
	return &users, err
}

// SearchContacts Searches for the specified query in the first names, last names and usernames of the known user contacts
// @param query Query to search for; may be empty to return all contacts
// @param limit The maximum number of users to be returned
func (client *Client) SearchContacts(query string, limit int32) (*Users, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "searchContacts",
		"query": query,
		"limit": limit,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var users Users
	err = json.Unmarshal(result.Raw, &users)
	return &users, err
}

// GetRecentInlineBots Returns up to 20 recently used inline bots in the order of their last usage
func (client *Client) GetRecentInlineBots() (*Users, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "getRecentInlineBots",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var users Users
	err = json.Unmarshal(result.Raw, &users)
	return &users, err
}
