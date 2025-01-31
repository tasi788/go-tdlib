package tdlib

import (
	"encoding/json"
	"fmt"
)

// Supergroup Represents a supergroup or channel with zero or more members (subscribers in the case of channels). From the point of view of the system, a channel is a special kind of a supergroup: only administrators can post and see the list of members, and posts from all administrators use the name and photo of the channel instead of individual names and profile photos. Unlike supergroups, channels can have an unlimited number of subscribers
type Supergroup struct {
	tdCommon
	Id                int64            `json:"id"`                   // Supergroup or channel identifier
	AccessHash        JSONInt64        `json:"access_hash"`          // Supergroup or channel access hash
	Username          string           `json:"username"`             // Username of the supergroup or channel; empty for private supergroups or channels
	Date              int32            `json:"date"`                 // Point in time (Unix timestamp) when the current user joined, or the point in time when the supergroup or channel was created, in case the user is not a member
	Status            ChatMemberStatus `json:"status"`               // Status of the current user in the supergroup or channel; custom title will be always empty
	MemberCount       int32            `json:"member_count"`         // Number of members in the supergroup or channel; 0 if unknown. Currently, it is guaranteed to be known only if the supergroup or channel was received through searchPublicChats, searchChatsNearby, getInactiveSupergroupChats, getSuitableDiscussionChats, getGroupsInCommon, or getUserPrivacySettingRules
	HasLinkedChat     bool             `json:"has_linked_chat"`      // True, if the channel has a discussion group, or the supergroup is the designated discussion group for a channel
	HasLocation       bool             `json:"has_location"`         // True, if the supergroup is connected to a location, i.e. the supergroup is a location-based supergroup
	SignMessages      bool             `json:"sign_messages"`        // True, if messages sent to the channel need to contain information about the sender. This field is only applicable to channels
	IsSlowModeEnabled bool             `json:"is_slow_mode_enabled"` // True, if the slow mode is enabled in the supergroup
	IsChannel         bool             `json:"is_channel"`           // True, if the supergroup is a channel
	IsBroadcastGroup  bool             `json:"is_broadcast_group"`   // True, if the supergroup is a broadcast group, i.e. only administrators can send messages and there is no limit on the number of members
	IsVerified        bool             `json:"is_verified"`          // True, if the supergroup or channel is verified
	RestrictionReason string           `json:"restriction_reason"`   // If non-empty, contains a human-readable description of the reason why access to this supergroup or channel must be restricted
	IsScam            bool             `json:"is_scam"`              // True, if many users reported this supergroup or channel as a scam
	IsFake            bool             `json:"is_fake"`              // True, if many users reported this supergroup or channel as a fake account
}

// MessageType return the string telegram-type of Supergroup
func (supergroup *Supergroup) MessageType() string {
	return "supergroup"
}

// NewSupergroup creates a new Supergroup
//
// @param id Supergroup or channel identifier
// @param accessHash Supergroup or channel access hash
// @param username Username of the supergroup or channel; empty for private supergroups or channels
// @param date Point in time (Unix timestamp) when the current user joined, or the point in time when the supergroup or channel was created, in case the user is not a member
// @param status Status of the current user in the supergroup or channel; custom title will be always empty
// @param memberCount Number of members in the supergroup or channel; 0 if unknown. Currently, it is guaranteed to be known only if the supergroup or channel was received through searchPublicChats, searchChatsNearby, getInactiveSupergroupChats, getSuitableDiscussionChats, getGroupsInCommon, or getUserPrivacySettingRules
// @param hasLinkedChat True, if the channel has a discussion group, or the supergroup is the designated discussion group for a channel
// @param hasLocation True, if the supergroup is connected to a location, i.e. the supergroup is a location-based supergroup
// @param signMessages True, if messages sent to the channel need to contain information about the sender. This field is only applicable to channels
// @param isSlowModeEnabled True, if the slow mode is enabled in the supergroup
// @param isChannel True, if the supergroup is a channel
// @param isBroadcastGroup True, if the supergroup is a broadcast group, i.e. only administrators can send messages and there is no limit on the number of members
// @param isVerified True, if the supergroup or channel is verified
// @param restrictionReason If non-empty, contains a human-readable description of the reason why access to this supergroup or channel must be restricted
// @param isScam True, if many users reported this supergroup or channel as a scam
// @param isFake True, if many users reported this supergroup or channel as a fake account
func NewSupergroup(id int64, accessHash JSONInt64, username string, date int32, status ChatMemberStatus, memberCount int32, hasLinkedChat bool, hasLocation bool, signMessages bool, isSlowModeEnabled bool, isChannel bool, isBroadcastGroup bool, isVerified bool, restrictionReason string, isScam bool, isFake bool) *Supergroup {
	supergroupTemp := Supergroup{
		tdCommon:          tdCommon{Type: "supergroup"},
		Id:                id,
		AccessHash:        accessHash,
		Username:          username,
		Date:              date,
		Status:            status,
		MemberCount:       memberCount,
		HasLinkedChat:     hasLinkedChat,
		HasLocation:       hasLocation,
		SignMessages:      signMessages,
		IsSlowModeEnabled: isSlowModeEnabled,
		IsChannel:         isChannel,
		IsBroadcastGroup:  isBroadcastGroup,
		IsVerified:        isVerified,
		RestrictionReason: restrictionReason,
		IsScam:            isScam,
		IsFake:            isFake,
	}

	return &supergroupTemp
}

// UnmarshalJSON unmarshal to json
func (supergroup *Supergroup) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	tempObj := struct {
		tdCommon
		Id                int64     `json:"id"`                   // Supergroup or channel identifier
		AccessHash        JSONInt64 `json:"access_hash"`          // Supergroup or channel access hash
		Username          string    `json:"username"`             // Username of the supergroup or channel; empty for private supergroups or channels
		Date              int32     `json:"date"`                 // Point in time (Unix timestamp) when the current user joined, or the point in time when the supergroup or channel was created, in case the user is not a member
		MemberCount       int32     `json:"member_count"`         // Number of members in the supergroup or channel; 0 if unknown. Currently, it is guaranteed to be known only if the supergroup or channel was received through searchPublicChats, searchChatsNearby, getInactiveSupergroupChats, getSuitableDiscussionChats, getGroupsInCommon, or getUserPrivacySettingRules
		HasLinkedChat     bool      `json:"has_linked_chat"`      // True, if the channel has a discussion group, or the supergroup is the designated discussion group for a channel
		HasLocation       bool      `json:"has_location"`         // True, if the supergroup is connected to a location, i.e. the supergroup is a location-based supergroup
		SignMessages      bool      `json:"sign_messages"`        // True, if messages sent to the channel need to contain information about the sender. This field is only applicable to channels
		IsSlowModeEnabled bool      `json:"is_slow_mode_enabled"` // True, if the slow mode is enabled in the supergroup
		IsChannel         bool      `json:"is_channel"`           // True, if the supergroup is a channel
		IsBroadcastGroup  bool      `json:"is_broadcast_group"`   // True, if the supergroup is a broadcast group, i.e. only administrators can send messages and there is no limit on the number of members
		IsVerified        bool      `json:"is_verified"`          // True, if the supergroup or channel is verified
		RestrictionReason string    `json:"restriction_reason"`   // If non-empty, contains a human-readable description of the reason why access to this supergroup or channel must be restricted
		IsScam            bool      `json:"is_scam"`              // True, if many users reported this supergroup or channel as a scam
		IsFake            bool      `json:"is_fake"`              // True, if many users reported this supergroup or channel as a fake account
	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	supergroup.tdCommon = tempObj.tdCommon
	supergroup.Id = tempObj.Id
	supergroup.AccessHash = tempObj.AccessHash
	supergroup.Username = tempObj.Username
	supergroup.Date = tempObj.Date
	supergroup.MemberCount = tempObj.MemberCount
	supergroup.HasLinkedChat = tempObj.HasLinkedChat
	supergroup.HasLocation = tempObj.HasLocation
	supergroup.SignMessages = tempObj.SignMessages
	supergroup.IsSlowModeEnabled = tempObj.IsSlowModeEnabled
	supergroup.IsChannel = tempObj.IsChannel
	supergroup.IsBroadcastGroup = tempObj.IsBroadcastGroup
	supergroup.IsVerified = tempObj.IsVerified
	supergroup.RestrictionReason = tempObj.RestrictionReason
	supergroup.IsScam = tempObj.IsScam
	supergroup.IsFake = tempObj.IsFake

	fieldStatus, _ := unmarshalChatMemberStatus(objMap["status"])
	supergroup.Status = fieldStatus

	return nil
}

// GetSupergroup Returns information about a supergroup or a channel by its identifier. This is an offline request if the current user is not a bot
// @param supergroupId Supergroup or channel identifier
func (client *Client) GetSupergroup(supergroupId int64) (*Supergroup, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":         "getSupergroup",
		"supergroup_id": supergroupId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var supergroupDummy Supergroup
	err = json.Unmarshal(result.Raw, &supergroupDummy)
	return &supergroupDummy, err
}
