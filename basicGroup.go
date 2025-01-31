package tdlib

import (
	"encoding/json"
	"fmt"
)

// BasicGroup Represents a basic group of 0-200 users (must be upgraded to a supergroup to accommodate more than 200 users)
type BasicGroup struct {
	tdCommon
	Id                     int64            `json:"id"`                        // Group identifier
	AccessHash             JSONInt64        `json:"access_hash"`               // Group access hash
	MemberCount            int32            `json:"member_count"`              // Number of members in the group
	Status                 ChatMemberStatus `json:"status"`                    // Status of the current user in the group
	IsActive               bool             `json:"is_active"`                 // True, if the group is active
	UpgradedToSupergroupId int32            `json:"upgraded_to_supergroup_id"` // Identifier of the supergroup to which this group was upgraded; 0 if none
}

// MessageType return the string telegram-type of BasicGroup
func (basicGroup *BasicGroup) MessageType() string {
	return "basicGroup"
}

// NewBasicGroup creates a new BasicGroup
//
// @param id Group identifier
// @param accessHash Group access hash
// @param memberCount Number of members in the group
// @param status Status of the current user in the group
// @param isActive True, if the group is active
// @param upgradedToSupergroupId Identifier of the supergroup to which this group was upgraded; 0 if none
func NewBasicGroup(id int64, accessHash JSONInt64, memberCount int32, status ChatMemberStatus, isActive bool, upgradedToSupergroupId int32) *BasicGroup {
	basicGroupTemp := BasicGroup{
		tdCommon:               tdCommon{Type: "basicGroup"},
		Id:                     id,
		AccessHash:             accessHash,
		MemberCount:            memberCount,
		Status:                 status,
		IsActive:               isActive,
		UpgradedToSupergroupId: upgradedToSupergroupId,
	}

	return &basicGroupTemp
}

// UnmarshalJSON unmarshal to json
func (basicGroup *BasicGroup) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	tempObj := struct {
		tdCommon
		Id                     int64     `json:"id"`                        // Group identifier
		AccessHash             JSONInt64 `json:"access_hash"`               // Group access hash
		MemberCount            int32     `json:"member_count"`              // Number of members in the group
		IsActive               bool      `json:"is_active"`                 // True, if the group is active
		UpgradedToSupergroupId int32     `json:"upgraded_to_supergroup_id"` // Identifier of the supergroup to which this group was upgraded; 0 if none
	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	basicGroup.tdCommon = tempObj.tdCommon
	basicGroup.Id = tempObj.Id
	basicGroup.AccessHash = tempObj.AccessHash
	basicGroup.MemberCount = tempObj.MemberCount
	basicGroup.IsActive = tempObj.IsActive
	basicGroup.UpgradedToSupergroupId = tempObj.UpgradedToSupergroupId

	fieldStatus, _ := unmarshalChatMemberStatus(objMap["status"])
	basicGroup.Status = fieldStatus

	return nil
}

// GetBasicGroup Returns information about a basic group by its identifier. This is an offline request if the current user is not a bot
// @param basicGroupId Basic group identifier
func (client *Client) GetBasicGroup(basicGroupId int64) (*BasicGroup, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":          "getBasicGroup",
		"basic_group_id": basicGroupId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var basicGroupDummy BasicGroup
	err = json.Unmarshal(result.Raw, &basicGroupDummy)
	return &basicGroupDummy, err
}
