package tdlib

import (
	"encoding/json"
	"fmt"
)

// TMeUrlType Describes the type of a URL linking to an internal Telegram entity
type TMeUrlType interface {
	GetTMeUrlTypeEnum() TMeUrlTypeEnum
}

// TMeUrlTypeEnum Alias for abstract TMeUrlType 'Sub-Classes', used as constant-enum here
type TMeUrlTypeEnum string

// TMeUrlType enums
const (
	TMeUrlTypeUserType       TMeUrlTypeEnum = "tMeUrlTypeUser"
	TMeUrlTypeSupergroupType TMeUrlTypeEnum = "tMeUrlTypeSupergroup"
	TMeUrlTypeChatInviteType TMeUrlTypeEnum = "tMeUrlTypeChatInvite"
	TMeUrlTypeStickerSetType TMeUrlTypeEnum = "tMeUrlTypeStickerSet"
)

func unmarshalTMeUrlType(rawMsg *json.RawMessage) (TMeUrlType, error) {

	if rawMsg == nil {
		return nil, nil
	}
	var objMap map[string]interface{}
	err := json.Unmarshal(*rawMsg, &objMap)
	if err != nil {
		return nil, err
	}

	switch TMeUrlTypeEnum(objMap["@type"].(string)) {
	case TMeUrlTypeUserType:
		var tMeUrlTypeUser TMeUrlTypeUser
		err := json.Unmarshal(*rawMsg, &tMeUrlTypeUser)
		return &tMeUrlTypeUser, err

	case TMeUrlTypeSupergroupType:
		var tMeUrlTypeSupergroup TMeUrlTypeSupergroup
		err := json.Unmarshal(*rawMsg, &tMeUrlTypeSupergroup)
		return &tMeUrlTypeSupergroup, err

	case TMeUrlTypeChatInviteType:
		var tMeUrlTypeChatInvite TMeUrlTypeChatInvite
		err := json.Unmarshal(*rawMsg, &tMeUrlTypeChatInvite)
		return &tMeUrlTypeChatInvite, err

	case TMeUrlTypeStickerSetType:
		var tMeUrlTypeStickerSet TMeUrlTypeStickerSet
		err := json.Unmarshal(*rawMsg, &tMeUrlTypeStickerSet)
		return &tMeUrlTypeStickerSet, err

	default:
		return nil, fmt.Errorf("Error UnMarshaling, unknown type:" + objMap["@type"].(string))
	}
}

// TMeUrlTypeUser A URL linking to a user
type TMeUrlTypeUser struct {
	tdCommon
	UserId int64 `json:"user_id"` // Identifier of the user
}

// MessageType return the string telegram-type of TMeUrlTypeUser
func (tMeUrlTypeUser *TMeUrlTypeUser) MessageType() string {
	return "tMeUrlTypeUser"
}

// NewTMeUrlTypeUser creates a new TMeUrlTypeUser
//
// @param userId Identifier of the user
func NewTMeUrlTypeUser(userId int64) *TMeUrlTypeUser {
	tMeUrlTypeUserTemp := TMeUrlTypeUser{
		tdCommon: tdCommon{Type: "tMeUrlTypeUser"},
		UserId:   userId,
	}

	return &tMeUrlTypeUserTemp
}

// GetTMeUrlTypeEnum return the enum type of this object
func (tMeUrlTypeUser *TMeUrlTypeUser) GetTMeUrlTypeEnum() TMeUrlTypeEnum {
	return TMeUrlTypeUserType
}

// TMeUrlTypeSupergroup A URL linking to a public supergroup or channel
type TMeUrlTypeSupergroup struct {
	tdCommon
	SupergroupId int64 `json:"supergroup_id"` // Identifier of the supergroup or channel
}

// MessageType return the string telegram-type of TMeUrlTypeSupergroup
func (tMeUrlTypeSupergroup *TMeUrlTypeSupergroup) MessageType() string {
	return "tMeUrlTypeSupergroup"
}

// NewTMeUrlTypeSupergroup creates a new TMeUrlTypeSupergroup
//
// @param supergroupId Identifier of the supergroup or channel
func NewTMeUrlTypeSupergroup(supergroupId int64) *TMeUrlTypeSupergroup {
	tMeUrlTypeSupergroupTemp := TMeUrlTypeSupergroup{
		tdCommon:     tdCommon{Type: "tMeUrlTypeSupergroup"},
		SupergroupId: supergroupId,
	}

	return &tMeUrlTypeSupergroupTemp
}

// GetTMeUrlTypeEnum return the enum type of this object
func (tMeUrlTypeSupergroup *TMeUrlTypeSupergroup) GetTMeUrlTypeEnum() TMeUrlTypeEnum {
	return TMeUrlTypeSupergroupType
}

// TMeUrlTypeChatInvite A chat invite link
type TMeUrlTypeChatInvite struct {
	tdCommon
	Info *ChatInviteLinkInfo `json:"info"` // Chat invite link info
}

// MessageType return the string telegram-type of TMeUrlTypeChatInvite
func (tMeUrlTypeChatInvite *TMeUrlTypeChatInvite) MessageType() string {
	return "tMeUrlTypeChatInvite"
}

// NewTMeUrlTypeChatInvite creates a new TMeUrlTypeChatInvite
//
// @param info Chat invite link info
func NewTMeUrlTypeChatInvite(info *ChatInviteLinkInfo) *TMeUrlTypeChatInvite {
	tMeUrlTypeChatInviteTemp := TMeUrlTypeChatInvite{
		tdCommon: tdCommon{Type: "tMeUrlTypeChatInvite"},
		Info:     info,
	}

	return &tMeUrlTypeChatInviteTemp
}

// GetTMeUrlTypeEnum return the enum type of this object
func (tMeUrlTypeChatInvite *TMeUrlTypeChatInvite) GetTMeUrlTypeEnum() TMeUrlTypeEnum {
	return TMeUrlTypeChatInviteType
}

// TMeUrlTypeStickerSet A URL linking to a sticker set
type TMeUrlTypeStickerSet struct {
	tdCommon
	StickerSetId JSONInt64 `json:"sticker_set_id"` // Identifier of the sticker set
}

// MessageType return the string telegram-type of TMeUrlTypeStickerSet
func (tMeUrlTypeStickerSet *TMeUrlTypeStickerSet) MessageType() string {
	return "tMeUrlTypeStickerSet"
}

// NewTMeUrlTypeStickerSet creates a new TMeUrlTypeStickerSet
//
// @param stickerSetId Identifier of the sticker set
func NewTMeUrlTypeStickerSet(stickerSetId JSONInt64) *TMeUrlTypeStickerSet {
	tMeUrlTypeStickerSetTemp := TMeUrlTypeStickerSet{
		tdCommon:     tdCommon{Type: "tMeUrlTypeStickerSet"},
		StickerSetId: stickerSetId,
	}

	return &tMeUrlTypeStickerSetTemp
}

// GetTMeUrlTypeEnum return the enum type of this object
func (tMeUrlTypeStickerSet *TMeUrlTypeStickerSet) GetTMeUrlTypeEnum() TMeUrlTypeEnum {
	return TMeUrlTypeStickerSetType
}
