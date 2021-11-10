package tdlib

import (
	"encoding/json"
	"fmt"
)

// ChatType Describes the type of a chat
type ChatType interface {
	GetChatTypeEnum() ChatTypeEnum
}

// ChatTypeEnum Alias for abstract ChatType 'Sub-Classes', used as constant-enum here
type ChatTypeEnum string

// ChatType enums
const (
	ChatTypePrivateType    ChatTypeEnum = "chatTypePrivate"
	ChatTypeBasicGroupType ChatTypeEnum = "chatTypeBasicGroup"
	ChatTypeSupergroupType ChatTypeEnum = "chatTypeSupergroup"
	ChatTypeSecretType     ChatTypeEnum = "chatTypeSecret"
)

func unmarshalChatType(rawMsg *json.RawMessage) (ChatType, error) {

	if rawMsg == nil {
		return nil, nil
	}
	var objMap map[string]interface{}
	err := json.Unmarshal(*rawMsg, &objMap)
	if err != nil {
		return nil, err
	}

	switch ChatTypeEnum(objMap["@type"].(string)) {
	case ChatTypePrivateType:
		var chatTypePrivate ChatTypePrivate
		err := json.Unmarshal(*rawMsg, &chatTypePrivate)
		return &chatTypePrivate, err

	case ChatTypeBasicGroupType:
		var chatTypeBasicGroup ChatTypeBasicGroup
		err := json.Unmarshal(*rawMsg, &chatTypeBasicGroup)
		return &chatTypeBasicGroup, err

	case ChatTypeSupergroupType:
		var chatTypeSupergroup ChatTypeSupergroup
		err := json.Unmarshal(*rawMsg, &chatTypeSupergroup)
		return &chatTypeSupergroup, err

	case ChatTypeSecretType:
		var chatTypeSecret ChatTypeSecret
		err := json.Unmarshal(*rawMsg, &chatTypeSecret)
		return &chatTypeSecret, err

	default:
		return nil, fmt.Errorf("Error UnMarshaling, unknown type:" + objMap["@type"].(string))
	}
}

// ChatTypePrivate An ordinary chat with a user
type ChatTypePrivate struct {
	tdCommon
	UserId int64 `json:"user_id"` // User identifier
}

// MessageType return the string telegram-type of ChatTypePrivate
func (chatTypePrivate *ChatTypePrivate) MessageType() string {
	return "chatTypePrivate"
}

// NewChatTypePrivate creates a new ChatTypePrivate
//
// @param userId User identifier
func NewChatTypePrivate(userId int64) *ChatTypePrivate {
	chatTypePrivateTemp := ChatTypePrivate{
		tdCommon: tdCommon{Type: "chatTypePrivate"},
		UserId:   userId,
	}

	return &chatTypePrivateTemp
}

// GetChatTypeEnum return the enum type of this object
func (chatTypePrivate *ChatTypePrivate) GetChatTypeEnum() ChatTypeEnum {
	return ChatTypePrivateType
}

// ChatTypeBasicGroup A basic group (a chat with 0-200 other users)
type ChatTypeBasicGroup struct {
	tdCommon
	BasicGroupId int64 `json:"basic_group_id"` // Basic group identifier
}

// MessageType return the string telegram-type of ChatTypeBasicGroup
func (chatTypeBasicGroup *ChatTypeBasicGroup) MessageType() string {
	return "chatTypeBasicGroup"
}

// NewChatTypeBasicGroup creates a new ChatTypeBasicGroup
//
// @param basicGroupId Basic group identifier
func NewChatTypeBasicGroup(basicGroupId int64) *ChatTypeBasicGroup {
	chatTypeBasicGroupTemp := ChatTypeBasicGroup{
		tdCommon:     tdCommon{Type: "chatTypeBasicGroup"},
		BasicGroupId: basicGroupId,
	}

	return &chatTypeBasicGroupTemp
}

// GetChatTypeEnum return the enum type of this object
func (chatTypeBasicGroup *ChatTypeBasicGroup) GetChatTypeEnum() ChatTypeEnum {
	return ChatTypeBasicGroupType
}

// ChatTypeSupergroup A supergroup or channel (with unlimited members)
type ChatTypeSupergroup struct {
	tdCommon
	SupergroupId int64 `json:"supergroup_id"` // Supergroup or channel identifier
	IsChannel    bool  `json:"is_channel"`    // True, if the supergroup is a channel
}

// MessageType return the string telegram-type of ChatTypeSupergroup
func (chatTypeSupergroup *ChatTypeSupergroup) MessageType() string {
	return "chatTypeSupergroup"
}

// NewChatTypeSupergroup creates a new ChatTypeSupergroup
//
// @param supergroupId Supergroup or channel identifier
// @param isChannel True, if the supergroup is a channel
func NewChatTypeSupergroup(supergroupId int64, isChannel bool) *ChatTypeSupergroup {
	chatTypeSupergroupTemp := ChatTypeSupergroup{
		tdCommon:     tdCommon{Type: "chatTypeSupergroup"},
		SupergroupId: supergroupId,
		IsChannel:    isChannel,
	}

	return &chatTypeSupergroupTemp
}

// GetChatTypeEnum return the enum type of this object
func (chatTypeSupergroup *ChatTypeSupergroup) GetChatTypeEnum() ChatTypeEnum {
	return ChatTypeSupergroupType
}

// ChatTypeSecret A secret chat with a user
type ChatTypeSecret struct {
	tdCommon
	SecretChatId int32 `json:"secret_chat_id"` // Secret chat identifier
	UserId       int64 `json:"user_id"`        // User identifier of the secret chat peer
}

// MessageType return the string telegram-type of ChatTypeSecret
func (chatTypeSecret *ChatTypeSecret) MessageType() string {
	return "chatTypeSecret"
}

// NewChatTypeSecret creates a new ChatTypeSecret
//
// @param secretChatId Secret chat identifier
// @param userId User identifier of the secret chat peer
func NewChatTypeSecret(secretChatId int32, userId int64) *ChatTypeSecret {
	chatTypeSecretTemp := ChatTypeSecret{
		tdCommon:     tdCommon{Type: "chatTypeSecret"},
		SecretChatId: secretChatId,
		UserId:       userId,
	}

	return &chatTypeSecretTemp
}

// GetChatTypeEnum return the enum type of this object
func (chatTypeSecret *ChatTypeSecret) GetChatTypeEnum() ChatTypeEnum {
	return ChatTypeSecretType
}
