package tdlib

import (
	"encoding/json"
	"fmt"
)

// UserStatus Describes the last time the user was online
type UserStatus interface {
	GetUserStatusEnum() UserStatusEnum
}

// UserStatusEnum Alias for abstract UserStatus 'Sub-Classes', used as constant-enum here
type UserStatusEnum string

// UserStatus enums
const (
	UserStatusEmptyType     UserStatusEnum = "userStatusEmpty"
	UserStatusOnlineType    UserStatusEnum = "userStatusOnline"
	UserStatusOfflineType   UserStatusEnum = "userStatusOffline"
	UserStatusRecentlyType  UserStatusEnum = "userStatusRecently"
	UserStatusLastWeekType  UserStatusEnum = "userStatusLastWeek"
	UserStatusLastMonthType UserStatusEnum = "userStatusLastMonth"
)

func unmarshalUserStatus(rawMsg *json.RawMessage) (UserStatus, error) {

	if rawMsg == nil {
		return nil, nil
	}
	var objMap map[string]interface{}
	err := json.Unmarshal(*rawMsg, &objMap)
	if err != nil {
		return nil, err
	}

	switch UserStatusEnum(objMap["@type"].(string)) {
	case UserStatusEmptyType:
		var userStatusEmpty UserStatusEmpty
		err := json.Unmarshal(*rawMsg, &userStatusEmpty)
		return &userStatusEmpty, err

	case UserStatusOnlineType:
		var userStatusOnline UserStatusOnline
		err := json.Unmarshal(*rawMsg, &userStatusOnline)
		return &userStatusOnline, err

	case UserStatusOfflineType:
		var userStatusOffline UserStatusOffline
		err := json.Unmarshal(*rawMsg, &userStatusOffline)
		return &userStatusOffline, err

	case UserStatusRecentlyType:
		var userStatusRecently UserStatusRecently
		err := json.Unmarshal(*rawMsg, &userStatusRecently)
		return &userStatusRecently, err

	case UserStatusLastWeekType:
		var userStatusLastWeek UserStatusLastWeek
		err := json.Unmarshal(*rawMsg, &userStatusLastWeek)
		return &userStatusLastWeek, err

	case UserStatusLastMonthType:
		var userStatusLastMonth UserStatusLastMonth
		err := json.Unmarshal(*rawMsg, &userStatusLastMonth)
		return &userStatusLastMonth, err

	default:
		return nil, fmt.Errorf("Error UnMarshaling, unknown type:" + objMap["@type"].(string))
	}
}

// UserStatusEmpty The user status was never changed
type UserStatusEmpty struct {
	tdCommon
}

// MessageType return the string telegram-type of UserStatusEmpty
func (userStatusEmpty *UserStatusEmpty) MessageType() string {
	return "userStatusEmpty"
}

// NewUserStatusEmpty creates a new UserStatusEmpty
//
func NewUserStatusEmpty() *UserStatusEmpty {
	userStatusEmptyTemp := UserStatusEmpty{
		tdCommon: tdCommon{Type: "userStatusEmpty"},
	}

	return &userStatusEmptyTemp
}

// GetUserStatusEnum return the enum type of this object
func (userStatusEmpty *UserStatusEmpty) GetUserStatusEnum() UserStatusEnum {
	return UserStatusEmptyType
}

// UserStatusOnline The user is online
type UserStatusOnline struct {
	tdCommon
	Expires int32 `json:"expires"` // Point in time (Unix timestamp) when the user's online status will expire
}

// MessageType return the string telegram-type of UserStatusOnline
func (userStatusOnline *UserStatusOnline) MessageType() string {
	return "userStatusOnline"
}

// NewUserStatusOnline creates a new UserStatusOnline
//
// @param expires Point in time (Unix timestamp) when the user's online status will expire
func NewUserStatusOnline(expires int32) *UserStatusOnline {
	userStatusOnlineTemp := UserStatusOnline{
		tdCommon: tdCommon{Type: "userStatusOnline"},
		Expires:  expires,
	}

	return &userStatusOnlineTemp
}

// GetUserStatusEnum return the enum type of this object
func (userStatusOnline *UserStatusOnline) GetUserStatusEnum() UserStatusEnum {
	return UserStatusOnlineType
}

// UserStatusOffline The user is offline
type UserStatusOffline struct {
	tdCommon
	WasOnline int32 `json:"was_online"` // Point in time (Unix timestamp) when the user was last online
}

// MessageType return the string telegram-type of UserStatusOffline
func (userStatusOffline *UserStatusOffline) MessageType() string {
	return "userStatusOffline"
}

// NewUserStatusOffline creates a new UserStatusOffline
//
// @param wasOnline Point in time (Unix timestamp) when the user was last online
func NewUserStatusOffline(wasOnline int32) *UserStatusOffline {
	userStatusOfflineTemp := UserStatusOffline{
		tdCommon:  tdCommon{Type: "userStatusOffline"},
		WasOnline: wasOnline,
	}

	return &userStatusOfflineTemp
}

// GetUserStatusEnum return the enum type of this object
func (userStatusOffline *UserStatusOffline) GetUserStatusEnum() UserStatusEnum {
	return UserStatusOfflineType
}

// UserStatusRecently The user was online recently
type UserStatusRecently struct {
	tdCommon
}

// MessageType return the string telegram-type of UserStatusRecently
func (userStatusRecently *UserStatusRecently) MessageType() string {
	return "userStatusRecently"
}

// NewUserStatusRecently creates a new UserStatusRecently
//
func NewUserStatusRecently() *UserStatusRecently {
	userStatusRecentlyTemp := UserStatusRecently{
		tdCommon: tdCommon{Type: "userStatusRecently"},
	}

	return &userStatusRecentlyTemp
}

// GetUserStatusEnum return the enum type of this object
func (userStatusRecently *UserStatusRecently) GetUserStatusEnum() UserStatusEnum {
	return UserStatusRecentlyType
}

// UserStatusLastWeek The user is offline, but was online last week
type UserStatusLastWeek struct {
	tdCommon
}

// MessageType return the string telegram-type of UserStatusLastWeek
func (userStatusLastWeek *UserStatusLastWeek) MessageType() string {
	return "userStatusLastWeek"
}

// NewUserStatusLastWeek creates a new UserStatusLastWeek
//
func NewUserStatusLastWeek() *UserStatusLastWeek {
	userStatusLastWeekTemp := UserStatusLastWeek{
		tdCommon: tdCommon{Type: "userStatusLastWeek"},
	}

	return &userStatusLastWeekTemp
}

// GetUserStatusEnum return the enum type of this object
func (userStatusLastWeek *UserStatusLastWeek) GetUserStatusEnum() UserStatusEnum {
	return UserStatusLastWeekType
}

// UserStatusLastMonth The user is offline, but was online last month
type UserStatusLastMonth struct {
	tdCommon
}

// MessageType return the string telegram-type of UserStatusLastMonth
func (userStatusLastMonth *UserStatusLastMonth) MessageType() string {
	return "userStatusLastMonth"
}

// NewUserStatusLastMonth creates a new UserStatusLastMonth
//
func NewUserStatusLastMonth() *UserStatusLastMonth {
	userStatusLastMonthTemp := UserStatusLastMonth{
		tdCommon: tdCommon{Type: "userStatusLastMonth"},
	}

	return &userStatusLastMonthTemp
}

// GetUserStatusEnum return the enum type of this object
func (userStatusLastMonth *UserStatusLastMonth) GetUserStatusEnum() UserStatusEnum {
	return UserStatusLastMonthType
}
