package tdlib

import (
	"encoding/json"
	"fmt"
)

// UserPrivacySettingRule Represents a single rule for managing privacy settings
type UserPrivacySettingRule interface {
	GetUserPrivacySettingRuleEnum() UserPrivacySettingRuleEnum
}

// UserPrivacySettingRuleEnum Alias for abstract UserPrivacySettingRule 'Sub-Classes', used as constant-enum here
type UserPrivacySettingRuleEnum string

// UserPrivacySettingRule enums
const (
	UserPrivacySettingRuleAllowAllType            UserPrivacySettingRuleEnum = "userPrivacySettingRuleAllowAll"
	UserPrivacySettingRuleAllowContactsType       UserPrivacySettingRuleEnum = "userPrivacySettingRuleAllowContacts"
	UserPrivacySettingRuleAllowUsersType          UserPrivacySettingRuleEnum = "userPrivacySettingRuleAllowUsers"
	UserPrivacySettingRuleAllowChatMembersType    UserPrivacySettingRuleEnum = "userPrivacySettingRuleAllowChatMembers"
	UserPrivacySettingRuleRestrictAllType         UserPrivacySettingRuleEnum = "userPrivacySettingRuleRestrictAll"
	UserPrivacySettingRuleRestrictContactsType    UserPrivacySettingRuleEnum = "userPrivacySettingRuleRestrictContacts"
	UserPrivacySettingRuleRestrictUsersType       UserPrivacySettingRuleEnum = "userPrivacySettingRuleRestrictUsers"
	UserPrivacySettingRuleRestrictChatMembersType UserPrivacySettingRuleEnum = "userPrivacySettingRuleRestrictChatMembers"
)

func unmarshalUserPrivacySettingRule(rawMsg *json.RawMessage) (UserPrivacySettingRule, error) {

	if rawMsg == nil {
		return nil, nil
	}
	var objMap map[string]interface{}
	err := json.Unmarshal(*rawMsg, &objMap)
	if err != nil {
		return nil, err
	}

	switch UserPrivacySettingRuleEnum(objMap["@type"].(string)) {
	case UserPrivacySettingRuleAllowAllType:
		var userPrivacySettingRuleAllowAll UserPrivacySettingRuleAllowAll
		err := json.Unmarshal(*rawMsg, &userPrivacySettingRuleAllowAll)
		return &userPrivacySettingRuleAllowAll, err

	case UserPrivacySettingRuleAllowContactsType:
		var userPrivacySettingRuleAllowContacts UserPrivacySettingRuleAllowContacts
		err := json.Unmarshal(*rawMsg, &userPrivacySettingRuleAllowContacts)
		return &userPrivacySettingRuleAllowContacts, err

	case UserPrivacySettingRuleAllowUsersType:
		var userPrivacySettingRuleAllowUsers UserPrivacySettingRuleAllowUsers
		err := json.Unmarshal(*rawMsg, &userPrivacySettingRuleAllowUsers)
		return &userPrivacySettingRuleAllowUsers, err

	case UserPrivacySettingRuleAllowChatMembersType:
		var userPrivacySettingRuleAllowChatMembers UserPrivacySettingRuleAllowChatMembers
		err := json.Unmarshal(*rawMsg, &userPrivacySettingRuleAllowChatMembers)
		return &userPrivacySettingRuleAllowChatMembers, err

	case UserPrivacySettingRuleRestrictAllType:
		var userPrivacySettingRuleRestrictAll UserPrivacySettingRuleRestrictAll
		err := json.Unmarshal(*rawMsg, &userPrivacySettingRuleRestrictAll)
		return &userPrivacySettingRuleRestrictAll, err

	case UserPrivacySettingRuleRestrictContactsType:
		var userPrivacySettingRuleRestrictContacts UserPrivacySettingRuleRestrictContacts
		err := json.Unmarshal(*rawMsg, &userPrivacySettingRuleRestrictContacts)
		return &userPrivacySettingRuleRestrictContacts, err

	case UserPrivacySettingRuleRestrictUsersType:
		var userPrivacySettingRuleRestrictUsers UserPrivacySettingRuleRestrictUsers
		err := json.Unmarshal(*rawMsg, &userPrivacySettingRuleRestrictUsers)
		return &userPrivacySettingRuleRestrictUsers, err

	case UserPrivacySettingRuleRestrictChatMembersType:
		var userPrivacySettingRuleRestrictChatMembers UserPrivacySettingRuleRestrictChatMembers
		err := json.Unmarshal(*rawMsg, &userPrivacySettingRuleRestrictChatMembers)
		return &userPrivacySettingRuleRestrictChatMembers, err

	default:
		return nil, fmt.Errorf("Error UnMarshaling, unknown type:" + objMap["@type"].(string))
	}
}

// UserPrivacySettingRuleAllowAll A rule to allow all users to do something
type UserPrivacySettingRuleAllowAll struct {
	tdCommon
}

// MessageType return the string telegram-type of UserPrivacySettingRuleAllowAll
func (userPrivacySettingRuleAllowAll *UserPrivacySettingRuleAllowAll) MessageType() string {
	return "userPrivacySettingRuleAllowAll"
}

// NewUserPrivacySettingRuleAllowAll creates a new UserPrivacySettingRuleAllowAll
//
func NewUserPrivacySettingRuleAllowAll() *UserPrivacySettingRuleAllowAll {
	userPrivacySettingRuleAllowAllTemp := UserPrivacySettingRuleAllowAll{
		tdCommon: tdCommon{Type: "userPrivacySettingRuleAllowAll"},
	}

	return &userPrivacySettingRuleAllowAllTemp
}

// GetUserPrivacySettingRuleEnum return the enum type of this object
func (userPrivacySettingRuleAllowAll *UserPrivacySettingRuleAllowAll) GetUserPrivacySettingRuleEnum() UserPrivacySettingRuleEnum {
	return UserPrivacySettingRuleAllowAllType
}

// UserPrivacySettingRuleAllowContacts A rule to allow all of a user's contacts to do something
type UserPrivacySettingRuleAllowContacts struct {
	tdCommon
}

// MessageType return the string telegram-type of UserPrivacySettingRuleAllowContacts
func (userPrivacySettingRuleAllowContacts *UserPrivacySettingRuleAllowContacts) MessageType() string {
	return "userPrivacySettingRuleAllowContacts"
}

// NewUserPrivacySettingRuleAllowContacts creates a new UserPrivacySettingRuleAllowContacts
//
func NewUserPrivacySettingRuleAllowContacts() *UserPrivacySettingRuleAllowContacts {
	userPrivacySettingRuleAllowContactsTemp := UserPrivacySettingRuleAllowContacts{
		tdCommon: tdCommon{Type: "userPrivacySettingRuleAllowContacts"},
	}

	return &userPrivacySettingRuleAllowContactsTemp
}

// GetUserPrivacySettingRuleEnum return the enum type of this object
func (userPrivacySettingRuleAllowContacts *UserPrivacySettingRuleAllowContacts) GetUserPrivacySettingRuleEnum() UserPrivacySettingRuleEnum {
	return UserPrivacySettingRuleAllowContactsType
}

// UserPrivacySettingRuleAllowUsers A rule to allow certain specified users to do something
type UserPrivacySettingRuleAllowUsers struct {
	tdCommon
	UserIds []int64 `json:"user_ids"` // The user identifiers, total number of users in all rules must not exceed 1000
}

// MessageType return the string telegram-type of UserPrivacySettingRuleAllowUsers
func (userPrivacySettingRuleAllowUsers *UserPrivacySettingRuleAllowUsers) MessageType() string {
	return "userPrivacySettingRuleAllowUsers"
}

// NewUserPrivacySettingRuleAllowUsers creates a new UserPrivacySettingRuleAllowUsers
//
// @param userIds The user identifiers, total number of users in all rules must not exceed 1000
func NewUserPrivacySettingRuleAllowUsers(userIds []int64) *UserPrivacySettingRuleAllowUsers {
	userPrivacySettingRuleAllowUsersTemp := UserPrivacySettingRuleAllowUsers{
		tdCommon: tdCommon{Type: "userPrivacySettingRuleAllowUsers"},
		UserIds:  userIds,
	}

	return &userPrivacySettingRuleAllowUsersTemp
}

// GetUserPrivacySettingRuleEnum return the enum type of this object
func (userPrivacySettingRuleAllowUsers *UserPrivacySettingRuleAllowUsers) GetUserPrivacySettingRuleEnum() UserPrivacySettingRuleEnum {
	return UserPrivacySettingRuleAllowUsersType
}

// UserPrivacySettingRuleAllowChatMembers A rule to allow all members of certain specified basic groups and supergroups to doing something
type UserPrivacySettingRuleAllowChatMembers struct {
	tdCommon
	ChatIds []int64 `json:"chat_ids"` // The chat identifiers, total number of chats in all rules must not exceed 20
}

// MessageType return the string telegram-type of UserPrivacySettingRuleAllowChatMembers
func (userPrivacySettingRuleAllowChatMembers *UserPrivacySettingRuleAllowChatMembers) MessageType() string {
	return "userPrivacySettingRuleAllowChatMembers"
}

// NewUserPrivacySettingRuleAllowChatMembers creates a new UserPrivacySettingRuleAllowChatMembers
//
// @param chatIds The chat identifiers, total number of chats in all rules must not exceed 20
func NewUserPrivacySettingRuleAllowChatMembers(chatIds []int64) *UserPrivacySettingRuleAllowChatMembers {
	userPrivacySettingRuleAllowChatMembersTemp := UserPrivacySettingRuleAllowChatMembers{
		tdCommon: tdCommon{Type: "userPrivacySettingRuleAllowChatMembers"},
		ChatIds:  chatIds,
	}

	return &userPrivacySettingRuleAllowChatMembersTemp
}

// GetUserPrivacySettingRuleEnum return the enum type of this object
func (userPrivacySettingRuleAllowChatMembers *UserPrivacySettingRuleAllowChatMembers) GetUserPrivacySettingRuleEnum() UserPrivacySettingRuleEnum {
	return UserPrivacySettingRuleAllowChatMembersType
}

// UserPrivacySettingRuleRestrictAll A rule to restrict all users from doing something
type UserPrivacySettingRuleRestrictAll struct {
	tdCommon
}

// MessageType return the string telegram-type of UserPrivacySettingRuleRestrictAll
func (userPrivacySettingRuleRestrictAll *UserPrivacySettingRuleRestrictAll) MessageType() string {
	return "userPrivacySettingRuleRestrictAll"
}

// NewUserPrivacySettingRuleRestrictAll creates a new UserPrivacySettingRuleRestrictAll
//
func NewUserPrivacySettingRuleRestrictAll() *UserPrivacySettingRuleRestrictAll {
	userPrivacySettingRuleRestrictAllTemp := UserPrivacySettingRuleRestrictAll{
		tdCommon: tdCommon{Type: "userPrivacySettingRuleRestrictAll"},
	}

	return &userPrivacySettingRuleRestrictAllTemp
}

// GetUserPrivacySettingRuleEnum return the enum type of this object
func (userPrivacySettingRuleRestrictAll *UserPrivacySettingRuleRestrictAll) GetUserPrivacySettingRuleEnum() UserPrivacySettingRuleEnum {
	return UserPrivacySettingRuleRestrictAllType
}

// UserPrivacySettingRuleRestrictContacts A rule to restrict all contacts of a user from doing something
type UserPrivacySettingRuleRestrictContacts struct {
	tdCommon
}

// MessageType return the string telegram-type of UserPrivacySettingRuleRestrictContacts
func (userPrivacySettingRuleRestrictContacts *UserPrivacySettingRuleRestrictContacts) MessageType() string {
	return "userPrivacySettingRuleRestrictContacts"
}

// NewUserPrivacySettingRuleRestrictContacts creates a new UserPrivacySettingRuleRestrictContacts
//
func NewUserPrivacySettingRuleRestrictContacts() *UserPrivacySettingRuleRestrictContacts {
	userPrivacySettingRuleRestrictContactsTemp := UserPrivacySettingRuleRestrictContacts{
		tdCommon: tdCommon{Type: "userPrivacySettingRuleRestrictContacts"},
	}

	return &userPrivacySettingRuleRestrictContactsTemp
}

// GetUserPrivacySettingRuleEnum return the enum type of this object
func (userPrivacySettingRuleRestrictContacts *UserPrivacySettingRuleRestrictContacts) GetUserPrivacySettingRuleEnum() UserPrivacySettingRuleEnum {
	return UserPrivacySettingRuleRestrictContactsType
}

// UserPrivacySettingRuleRestrictUsers A rule to restrict all specified users from doing something
type UserPrivacySettingRuleRestrictUsers struct {
	tdCommon
	UserIds []int64 `json:"user_ids"` // The user identifiers, total number of users in all rules must not exceed 1000
}

// MessageType return the string telegram-type of UserPrivacySettingRuleRestrictUsers
func (userPrivacySettingRuleRestrictUsers *UserPrivacySettingRuleRestrictUsers) MessageType() string {
	return "userPrivacySettingRuleRestrictUsers"
}

// NewUserPrivacySettingRuleRestrictUsers creates a new UserPrivacySettingRuleRestrictUsers
//
// @param userIds The user identifiers, total number of users in all rules must not exceed 1000
func NewUserPrivacySettingRuleRestrictUsers(userIds []int64) *UserPrivacySettingRuleRestrictUsers {
	userPrivacySettingRuleRestrictUsersTemp := UserPrivacySettingRuleRestrictUsers{
		tdCommon: tdCommon{Type: "userPrivacySettingRuleRestrictUsers"},
		UserIds:  userIds,
	}

	return &userPrivacySettingRuleRestrictUsersTemp
}

// GetUserPrivacySettingRuleEnum return the enum type of this object
func (userPrivacySettingRuleRestrictUsers *UserPrivacySettingRuleRestrictUsers) GetUserPrivacySettingRuleEnum() UserPrivacySettingRuleEnum {
	return UserPrivacySettingRuleRestrictUsersType
}

// UserPrivacySettingRuleRestrictChatMembers A rule to restrict all members of specified basic groups and supergroups from doing something
type UserPrivacySettingRuleRestrictChatMembers struct {
	tdCommon
	ChatIds []int64 `json:"chat_ids"` // The chat identifiers, total number of chats in all rules must not exceed 20
}

// MessageType return the string telegram-type of UserPrivacySettingRuleRestrictChatMembers
func (userPrivacySettingRuleRestrictChatMembers *UserPrivacySettingRuleRestrictChatMembers) MessageType() string {
	return "userPrivacySettingRuleRestrictChatMembers"
}

// NewUserPrivacySettingRuleRestrictChatMembers creates a new UserPrivacySettingRuleRestrictChatMembers
//
// @param chatIds The chat identifiers, total number of chats in all rules must not exceed 20
func NewUserPrivacySettingRuleRestrictChatMembers(chatIds []int64) *UserPrivacySettingRuleRestrictChatMembers {
	userPrivacySettingRuleRestrictChatMembersTemp := UserPrivacySettingRuleRestrictChatMembers{
		tdCommon: tdCommon{Type: "userPrivacySettingRuleRestrictChatMembers"},
		ChatIds:  chatIds,
	}

	return &userPrivacySettingRuleRestrictChatMembersTemp
}

// GetUserPrivacySettingRuleEnum return the enum type of this object
func (userPrivacySettingRuleRestrictChatMembers *UserPrivacySettingRuleRestrictChatMembers) GetUserPrivacySettingRuleEnum() UserPrivacySettingRuleEnum {
	return UserPrivacySettingRuleRestrictChatMembersType
}
