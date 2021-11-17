package tdlib

import (
	"encoding/json"
	"fmt"
)

// BasicGroupFullInfo Contains full information about a basic group
type BasicGroupFullInfo struct {
	tdCommon
	Photo         *ChatPhoto      `json:"photo"`           // Chat photo; may be null
	Description   string          `json:"description"`     // Group description. Updated only after the basic group is opened
	CreatorUserId int64           `json:"creator_user_id"` // User identifier of the creator of the group; 0 if unknown
	Members       []ChatMember    `json:"members"`         // Group members
	InviteLink    *ChatInviteLink `json:"invite_link"`     // Primary invite link for this group; may be null. For chat administrators with can_invite_users right only. Updated only after the basic group is opened
	BotCommands   []BotCommands   `json:"bot_commands"`    // List of commands of bots in the group
}

// MessageType return the string telegram-type of BasicGroupFullInfo
func (basicGroupFullInfo *BasicGroupFullInfo) MessageType() string {
	return "basicGroupFullInfo"
}

// NewBasicGroupFullInfo creates a new BasicGroupFullInfo
//
// @param photo Chat photo; may be null
// @param description Group description. Updated only after the basic group is opened
// @param creatorUserId User identifier of the creator of the group; 0 if unknown
// @param members Group members
// @param inviteLink Primary invite link for this group; may be null. For chat administrators with can_invite_users right only. Updated only after the basic group is opened
// @param botCommands List of commands of bots in the group
func NewBasicGroupFullInfo(photo *ChatPhoto, description string, creatorUserId int64, members []ChatMember, inviteLink *ChatInviteLink, botCommands []BotCommands) *BasicGroupFullInfo {
	basicGroupFullInfoTemp := BasicGroupFullInfo{
		tdCommon:      tdCommon{Type: "basicGroupFullInfo"},
		Photo:         photo,
		Description:   description,
		CreatorUserId: creatorUserId,
		Members:       members,
		InviteLink:    inviteLink,
		BotCommands:   botCommands,
	}

	return &basicGroupFullInfoTemp
}

// GetBasicGroupFullInfo Returns full information about a basic group by its identifier
// @param basicGroupId Basic group identifier
func (client *Client) GetBasicGroupFullInfo(basicGroupId int64) (*BasicGroupFullInfo, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":          "getBasicGroupFullInfo",
		"basic_group_id": basicGroupId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var basicGroupFullInfo BasicGroupFullInfo
	err = json.Unmarshal(result.Raw, &basicGroupFullInfo)
	return &basicGroupFullInfo, err
}
