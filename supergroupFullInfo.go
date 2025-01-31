package tdlib

import (
	"encoding/json"
	"fmt"
)

// SupergroupFullInfo Contains full information about a supergroup or channel
type SupergroupFullInfo struct {
	tdCommon
	Photo                    *ChatPhoto      `json:"photo"`                        // Chat photo; may be null
	Description              string          `json:"description"`                  // Supergroup or channel description
	MemberCount              int32           `json:"member_count"`                 // Number of members in the supergroup or channel; 0 if unknown
	AdministratorCount       int32           `json:"administrator_count"`          // Number of privileged users in the supergroup or channel; 0 if unknown
	RestrictedCount          int32           `json:"restricted_count"`             // Number of restricted users in the supergroup; 0 if unknown
	BannedCount              int32           `json:"banned_count"`                 // Number of users banned from chat; 0 if unknown
	LinkedChatId             int64           `json:"linked_chat_id"`               // Chat identifier of a discussion group for the channel, or a channel, for which the supergroup is the designated discussion group; 0 if none or unknown
	SlowModeDelay            int32           `json:"slow_mode_delay"`              // Delay between consecutive sent messages for non-administrator supergroup members, in seconds
	SlowModeDelayExpiresIn   float64         `json:"slow_mode_delay_expires_in"`   // Time left before next message can be sent in the supergroup, in seconds. An updateSupergroupFullInfo update is not triggered when value of this field changes, but both new and old values are non-zero
	CanGetMembers            bool            `json:"can_get_members"`              // True, if members of the chat can be retrieved
	CanSetUsername           bool            `json:"can_set_username"`             // True, if the chat username can be changed
	CanSetStickerSet         bool            `json:"can_set_sticker_set"`          // True, if the supergroup sticker set can be changed
	CanSetLocation           bool            `json:"can_set_location"`             // True, if the supergroup location can be changed
	CanGetStatistics         bool            `json:"can_get_statistics"`           // True, if the supergroup or channel statistics are available
	IsAllHistoryAvailable    bool            `json:"is_all_history_available"`     // True, if new chat members will have access to old messages. In public or discussion groups and both public and private channels, old messages are always available, so this option affects only private supergroups without a linked chat. The value of this field is only available for chat administrators
	StickerSetId             JSONInt64       `json:"sticker_set_id"`               // Identifier of the supergroup sticker set; 0 if none
	Location                 *ChatLocation   `json:"location"`                     // Location to which the supergroup is connected; may be null
	InviteLink               *ChatInviteLink `json:"invite_link"`                  // Primary invite link for this chat; may be null. For chat administrators with can_invite_users right only
	BotCommands              []BotCommands   `json:"bot_commands"`                 // List of commands of bots in the group
	UpgradedFromBasicGroupId int64           `json:"upgraded_from_basic_group_id"` // Identifier of the basic group from which supergroup was upgraded; 0 if none
	UpgradedFromMaxMessageId int64           `json:"upgraded_from_max_message_id"` // Identifier of the last message in the basic group from which supergroup was upgraded; 0 if none
}

// MessageType return the string telegram-type of SupergroupFullInfo
func (supergroupFullInfo *SupergroupFullInfo) MessageType() string {
	return "supergroupFullInfo"
}

// NewSupergroupFullInfo creates a new SupergroupFullInfo
//
// @param photo Chat photo; may be null
// @param description Supergroup or channel description
// @param memberCount Number of members in the supergroup or channel; 0 if unknown
// @param administratorCount Number of privileged users in the supergroup or channel; 0 if unknown
// @param restrictedCount Number of restricted users in the supergroup; 0 if unknown
// @param bannedCount Number of users banned from chat; 0 if unknown
// @param linkedChatId Chat identifier of a discussion group for the channel, or a channel, for which the supergroup is the designated discussion group; 0 if none or unknown
// @param slowModeDelay Delay between consecutive sent messages for non-administrator supergroup members, in seconds
// @param slowModeDelayExpiresIn Time left before next message can be sent in the supergroup, in seconds. An updateSupergroupFullInfo update is not triggered when value of this field changes, but both new and old values are non-zero
// @param canGetMembers True, if members of the chat can be retrieved
// @param canSetUsername True, if the chat username can be changed
// @param canSetStickerSet True, if the supergroup sticker set can be changed
// @param canSetLocation True, if the supergroup location can be changed
// @param canGetStatistics True, if the supergroup or channel statistics are available
// @param isAllHistoryAvailable True, if new chat members will have access to old messages. In public or discussion groups and both public and private channels, old messages are always available, so this option affects only private supergroups without a linked chat. The value of this field is only available for chat administrators
// @param stickerSetId Identifier of the supergroup sticker set; 0 if none
// @param location Location to which the supergroup is connected; may be null
// @param inviteLink Primary invite link for this chat; may be null. For chat administrators with can_invite_users right only
// @param botCommands List of commands of bots in the group
// @param upgradedFromBasicGroupId Identifier of the basic group from which supergroup was upgraded; 0 if none
// @param upgradedFromMaxMessageId Identifier of the last message in the basic group from which supergroup was upgraded; 0 if none
func NewSupergroupFullInfo(photo *ChatPhoto, description string, memberCount int32, administratorCount int32, restrictedCount int32, bannedCount int32, linkedChatId int64, slowModeDelay int32, slowModeDelayExpiresIn float64, canGetMembers bool, canSetUsername bool, canSetStickerSet bool, canSetLocation bool, canGetStatistics bool, isAllHistoryAvailable bool, stickerSetId JSONInt64, location *ChatLocation, inviteLink *ChatInviteLink, botCommands []BotCommands, upgradedFromBasicGroupId int64, upgradedFromMaxMessageId int64) *SupergroupFullInfo {
	supergroupFullInfoTemp := SupergroupFullInfo{
		tdCommon:                 tdCommon{Type: "supergroupFullInfo"},
		Photo:                    photo,
		Description:              description,
		MemberCount:              memberCount,
		AdministratorCount:       administratorCount,
		RestrictedCount:          restrictedCount,
		BannedCount:              bannedCount,
		LinkedChatId:             linkedChatId,
		SlowModeDelay:            slowModeDelay,
		SlowModeDelayExpiresIn:   slowModeDelayExpiresIn,
		CanGetMembers:            canGetMembers,
		CanSetUsername:           canSetUsername,
		CanSetStickerSet:         canSetStickerSet,
		CanSetLocation:           canSetLocation,
		CanGetStatistics:         canGetStatistics,
		IsAllHistoryAvailable:    isAllHistoryAvailable,
		StickerSetId:             stickerSetId,
		Location:                 location,
		InviteLink:               inviteLink,
		BotCommands:              botCommands,
		UpgradedFromBasicGroupId: upgradedFromBasicGroupId,
		UpgradedFromMaxMessageId: upgradedFromMaxMessageId,
	}

	return &supergroupFullInfoTemp
}

// GetSupergroupFullInfo Returns full information about a supergroup or a channel by its identifier, cached for up to 1 minute
// @param supergroupId Supergroup or channel identifier
func (client *Client) GetSupergroupFullInfo(supergroupId int64) (*SupergroupFullInfo, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":         "getSupergroupFullInfo",
		"supergroup_id": supergroupId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var supergroupFullInfo SupergroupFullInfo
	err = json.Unmarshal(result.Raw, &supergroupFullInfo)
	return &supergroupFullInfo, err
}
