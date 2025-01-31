package tdlib

import (
	"encoding/json"
	"fmt"
)

// UserFullInfo Contains full information about a user
type UserFullInfo struct {
	tdCommon
	Photo                           *ChatPhoto   `json:"photo"`                               // User profile photo; may be null
	IsBlocked                       bool         `json:"is_blocked"`                          // True, if the user is blocked by the current user
	CanBeCalled                     bool         `json:"can_be_called"`                       // True, if the user can be called
	SupportsVideoCalls              bool         `json:"supports_video_calls"`                // True, if a video call can be created with the user
	HasPrivateCalls                 bool         `json:"has_private_calls"`                   // True, if the user can't be called due to their privacy settings
	HasPrivateForwards              bool         `json:"has_private_forwards"`                // True, if the user can't be linked in forwarded messages due to their privacy settings
	NeedPhoneNumberPrivacyException bool         `json:"need_phone_number_privacy_exception"` // True, if the current user needs to explicitly allow to share their phone number with the user when the method addContact is used
	Bio                             string       `json:"bio"`                                 // A short user bio
	ShareText                       string       `json:"share_text"`                          // For bots, the text that is shown on the bot's profile page and is sent together with the link when users share the bot
	Description                     string       `json:"description"`                         // For bots, the text shown in the chat with the bot if the chat is empty
	GroupInCommonCount              int32        `json:"group_in_common_count"`               // Number of group chats where both the other user and the current user are a member; 0 for the current user
	Commands                        []BotCommand `json:"commands"`                            // For bots, list of the bot commands
}

// MessageType return the string telegram-type of UserFullInfo
func (userFullInfo *UserFullInfo) MessageType() string {
	return "userFullInfo"
}

// NewUserFullInfo creates a new UserFullInfo
//
// @param photo User profile photo; may be null
// @param isBlocked True, if the user is blocked by the current user
// @param canBeCalled True, if the user can be called
// @param supportsVideoCalls True, if a video call can be created with the user
// @param hasPrivateCalls True, if the user can't be called due to their privacy settings
// @param hasPrivateForwards True, if the user can't be linked in forwarded messages due to their privacy settings
// @param needPhoneNumberPrivacyException True, if the current user needs to explicitly allow to share their phone number with the user when the method addContact is used
// @param bio A short user bio
// @param shareText For bots, the text that is shown on the bot's profile page and is sent together with the link when users share the bot
// @param description For bots, the text shown in the chat with the bot if the chat is empty
// @param groupInCommonCount Number of group chats where both the other user and the current user are a member; 0 for the current user
// @param commands For bots, list of the bot commands
func NewUserFullInfo(photo *ChatPhoto, isBlocked bool, canBeCalled bool, supportsVideoCalls bool, hasPrivateCalls bool, hasPrivateForwards bool, needPhoneNumberPrivacyException bool, bio string, shareText string, description string, groupInCommonCount int32, commands []BotCommand) *UserFullInfo {
	userFullInfoTemp := UserFullInfo{
		tdCommon:                        tdCommon{Type: "userFullInfo"},
		Photo:                           photo,
		IsBlocked:                       isBlocked,
		CanBeCalled:                     canBeCalled,
		SupportsVideoCalls:              supportsVideoCalls,
		HasPrivateCalls:                 hasPrivateCalls,
		HasPrivateForwards:              hasPrivateForwards,
		NeedPhoneNumberPrivacyException: needPhoneNumberPrivacyException,
		Bio:                             bio,
		ShareText:                       shareText,
		Description:                     description,
		GroupInCommonCount:              groupInCommonCount,
		Commands:                        commands,
	}

	return &userFullInfoTemp
}

// GetUserFullInfo Returns full information about a user by their identifier
// @param userId User identifier
func (client *Client) GetUserFullInfo(userId int64) (*UserFullInfo, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "getUserFullInfo",
		"user_id": userId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var userFullInfo UserFullInfo
	err = json.Unmarshal(result.Raw, &userFullInfo)
	return &userFullInfo, err
}
