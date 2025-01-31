package tdlib

import (
	"encoding/json"
	"fmt"
)

// UserPrivacySetting Describes available user privacy settings
type UserPrivacySetting interface {
	GetUserPrivacySettingEnum() UserPrivacySettingEnum
}

// UserPrivacySettingEnum Alias for abstract UserPrivacySetting 'Sub-Classes', used as constant-enum here
type UserPrivacySettingEnum string

// UserPrivacySetting enums
const (
	UserPrivacySettingShowStatusType                  UserPrivacySettingEnum = "userPrivacySettingShowStatus"
	UserPrivacySettingShowProfilePhotoType            UserPrivacySettingEnum = "userPrivacySettingShowProfilePhoto"
	UserPrivacySettingShowLinkInForwardedMessagesType UserPrivacySettingEnum = "userPrivacySettingShowLinkInForwardedMessages"
	UserPrivacySettingShowPhoneNumberType             UserPrivacySettingEnum = "userPrivacySettingShowPhoneNumber"
	UserPrivacySettingAllowChatInvitesType            UserPrivacySettingEnum = "userPrivacySettingAllowChatInvites"
	UserPrivacySettingAllowCallsType                  UserPrivacySettingEnum = "userPrivacySettingAllowCalls"
	UserPrivacySettingAllowPeerToPeerCallsType        UserPrivacySettingEnum = "userPrivacySettingAllowPeerToPeerCalls"
	UserPrivacySettingAllowFindingByPhoneNumberType   UserPrivacySettingEnum = "userPrivacySettingAllowFindingByPhoneNumber"
)

func unmarshalUserPrivacySetting(rawMsg *json.RawMessage) (UserPrivacySetting, error) {

	if rawMsg == nil {
		return nil, nil
	}
	var objMap map[string]interface{}
	err := json.Unmarshal(*rawMsg, &objMap)
	if err != nil {
		return nil, err
	}

	switch UserPrivacySettingEnum(objMap["@type"].(string)) {
	case UserPrivacySettingShowStatusType:
		var userPrivacySettingShowStatus UserPrivacySettingShowStatus
		err := json.Unmarshal(*rawMsg, &userPrivacySettingShowStatus)
		return &userPrivacySettingShowStatus, err

	case UserPrivacySettingShowProfilePhotoType:
		var userPrivacySettingShowProfilePhoto UserPrivacySettingShowProfilePhoto
		err := json.Unmarshal(*rawMsg, &userPrivacySettingShowProfilePhoto)
		return &userPrivacySettingShowProfilePhoto, err

	case UserPrivacySettingShowLinkInForwardedMessagesType:
		var userPrivacySettingShowLinkInForwardedMessages UserPrivacySettingShowLinkInForwardedMessages
		err := json.Unmarshal(*rawMsg, &userPrivacySettingShowLinkInForwardedMessages)
		return &userPrivacySettingShowLinkInForwardedMessages, err

	case UserPrivacySettingShowPhoneNumberType:
		var userPrivacySettingShowPhoneNumber UserPrivacySettingShowPhoneNumber
		err := json.Unmarshal(*rawMsg, &userPrivacySettingShowPhoneNumber)
		return &userPrivacySettingShowPhoneNumber, err

	case UserPrivacySettingAllowChatInvitesType:
		var userPrivacySettingAllowChatInvites UserPrivacySettingAllowChatInvites
		err := json.Unmarshal(*rawMsg, &userPrivacySettingAllowChatInvites)
		return &userPrivacySettingAllowChatInvites, err

	case UserPrivacySettingAllowCallsType:
		var userPrivacySettingAllowCalls UserPrivacySettingAllowCalls
		err := json.Unmarshal(*rawMsg, &userPrivacySettingAllowCalls)
		return &userPrivacySettingAllowCalls, err

	case UserPrivacySettingAllowPeerToPeerCallsType:
		var userPrivacySettingAllowPeerToPeerCalls UserPrivacySettingAllowPeerToPeerCalls
		err := json.Unmarshal(*rawMsg, &userPrivacySettingAllowPeerToPeerCalls)
		return &userPrivacySettingAllowPeerToPeerCalls, err

	case UserPrivacySettingAllowFindingByPhoneNumberType:
		var userPrivacySettingAllowFindingByPhoneNumber UserPrivacySettingAllowFindingByPhoneNumber
		err := json.Unmarshal(*rawMsg, &userPrivacySettingAllowFindingByPhoneNumber)
		return &userPrivacySettingAllowFindingByPhoneNumber, err

	default:
		return nil, fmt.Errorf("Error UnMarshaling, unknown type:" + objMap["@type"].(string))
	}
}

// UserPrivacySettingShowStatus A privacy setting for managing whether the user's online status is visible
type UserPrivacySettingShowStatus struct {
	tdCommon
}

// MessageType return the string telegram-type of UserPrivacySettingShowStatus
func (userPrivacySettingShowStatus *UserPrivacySettingShowStatus) MessageType() string {
	return "userPrivacySettingShowStatus"
}

// NewUserPrivacySettingShowStatus creates a new UserPrivacySettingShowStatus
//
func NewUserPrivacySettingShowStatus() *UserPrivacySettingShowStatus {
	userPrivacySettingShowStatusTemp := UserPrivacySettingShowStatus{
		tdCommon: tdCommon{Type: "userPrivacySettingShowStatus"},
	}

	return &userPrivacySettingShowStatusTemp
}

// GetUserPrivacySettingEnum return the enum type of this object
func (userPrivacySettingShowStatus *UserPrivacySettingShowStatus) GetUserPrivacySettingEnum() UserPrivacySettingEnum {
	return UserPrivacySettingShowStatusType
}

// UserPrivacySettingShowProfilePhoto A privacy setting for managing whether the user's profile photo is visible
type UserPrivacySettingShowProfilePhoto struct {
	tdCommon
}

// MessageType return the string telegram-type of UserPrivacySettingShowProfilePhoto
func (userPrivacySettingShowProfilePhoto *UserPrivacySettingShowProfilePhoto) MessageType() string {
	return "userPrivacySettingShowProfilePhoto"
}

// NewUserPrivacySettingShowProfilePhoto creates a new UserPrivacySettingShowProfilePhoto
//
func NewUserPrivacySettingShowProfilePhoto() *UserPrivacySettingShowProfilePhoto {
	userPrivacySettingShowProfilePhotoTemp := UserPrivacySettingShowProfilePhoto{
		tdCommon: tdCommon{Type: "userPrivacySettingShowProfilePhoto"},
	}

	return &userPrivacySettingShowProfilePhotoTemp
}

// GetUserPrivacySettingEnum return the enum type of this object
func (userPrivacySettingShowProfilePhoto *UserPrivacySettingShowProfilePhoto) GetUserPrivacySettingEnum() UserPrivacySettingEnum {
	return UserPrivacySettingShowProfilePhotoType
}

// UserPrivacySettingShowLinkInForwardedMessages A privacy setting for managing whether a link to the user's account is included in forwarded messages
type UserPrivacySettingShowLinkInForwardedMessages struct {
	tdCommon
}

// MessageType return the string telegram-type of UserPrivacySettingShowLinkInForwardedMessages
func (userPrivacySettingShowLinkInForwardedMessages *UserPrivacySettingShowLinkInForwardedMessages) MessageType() string {
	return "userPrivacySettingShowLinkInForwardedMessages"
}

// NewUserPrivacySettingShowLinkInForwardedMessages creates a new UserPrivacySettingShowLinkInForwardedMessages
//
func NewUserPrivacySettingShowLinkInForwardedMessages() *UserPrivacySettingShowLinkInForwardedMessages {
	userPrivacySettingShowLinkInForwardedMessagesTemp := UserPrivacySettingShowLinkInForwardedMessages{
		tdCommon: tdCommon{Type: "userPrivacySettingShowLinkInForwardedMessages"},
	}

	return &userPrivacySettingShowLinkInForwardedMessagesTemp
}

// GetUserPrivacySettingEnum return the enum type of this object
func (userPrivacySettingShowLinkInForwardedMessages *UserPrivacySettingShowLinkInForwardedMessages) GetUserPrivacySettingEnum() UserPrivacySettingEnum {
	return UserPrivacySettingShowLinkInForwardedMessagesType
}

// UserPrivacySettingShowPhoneNumber A privacy setting for managing whether the user's phone number is visible
type UserPrivacySettingShowPhoneNumber struct {
	tdCommon
}

// MessageType return the string telegram-type of UserPrivacySettingShowPhoneNumber
func (userPrivacySettingShowPhoneNumber *UserPrivacySettingShowPhoneNumber) MessageType() string {
	return "userPrivacySettingShowPhoneNumber"
}

// NewUserPrivacySettingShowPhoneNumber creates a new UserPrivacySettingShowPhoneNumber
//
func NewUserPrivacySettingShowPhoneNumber() *UserPrivacySettingShowPhoneNumber {
	userPrivacySettingShowPhoneNumberTemp := UserPrivacySettingShowPhoneNumber{
		tdCommon: tdCommon{Type: "userPrivacySettingShowPhoneNumber"},
	}

	return &userPrivacySettingShowPhoneNumberTemp
}

// GetUserPrivacySettingEnum return the enum type of this object
func (userPrivacySettingShowPhoneNumber *UserPrivacySettingShowPhoneNumber) GetUserPrivacySettingEnum() UserPrivacySettingEnum {
	return UserPrivacySettingShowPhoneNumberType
}

// UserPrivacySettingAllowChatInvites A privacy setting for managing whether the user can be invited to chats
type UserPrivacySettingAllowChatInvites struct {
	tdCommon
}

// MessageType return the string telegram-type of UserPrivacySettingAllowChatInvites
func (userPrivacySettingAllowChatInvites *UserPrivacySettingAllowChatInvites) MessageType() string {
	return "userPrivacySettingAllowChatInvites"
}

// NewUserPrivacySettingAllowChatInvites creates a new UserPrivacySettingAllowChatInvites
//
func NewUserPrivacySettingAllowChatInvites() *UserPrivacySettingAllowChatInvites {
	userPrivacySettingAllowChatInvitesTemp := UserPrivacySettingAllowChatInvites{
		tdCommon: tdCommon{Type: "userPrivacySettingAllowChatInvites"},
	}

	return &userPrivacySettingAllowChatInvitesTemp
}

// GetUserPrivacySettingEnum return the enum type of this object
func (userPrivacySettingAllowChatInvites *UserPrivacySettingAllowChatInvites) GetUserPrivacySettingEnum() UserPrivacySettingEnum {
	return UserPrivacySettingAllowChatInvitesType
}

// UserPrivacySettingAllowCalls A privacy setting for managing whether the user can be called
type UserPrivacySettingAllowCalls struct {
	tdCommon
}

// MessageType return the string telegram-type of UserPrivacySettingAllowCalls
func (userPrivacySettingAllowCalls *UserPrivacySettingAllowCalls) MessageType() string {
	return "userPrivacySettingAllowCalls"
}

// NewUserPrivacySettingAllowCalls creates a new UserPrivacySettingAllowCalls
//
func NewUserPrivacySettingAllowCalls() *UserPrivacySettingAllowCalls {
	userPrivacySettingAllowCallsTemp := UserPrivacySettingAllowCalls{
		tdCommon: tdCommon{Type: "userPrivacySettingAllowCalls"},
	}

	return &userPrivacySettingAllowCallsTemp
}

// GetUserPrivacySettingEnum return the enum type of this object
func (userPrivacySettingAllowCalls *UserPrivacySettingAllowCalls) GetUserPrivacySettingEnum() UserPrivacySettingEnum {
	return UserPrivacySettingAllowCallsType
}

// UserPrivacySettingAllowPeerToPeerCalls A privacy setting for managing whether peer-to-peer connections can be used for calls
type UserPrivacySettingAllowPeerToPeerCalls struct {
	tdCommon
}

// MessageType return the string telegram-type of UserPrivacySettingAllowPeerToPeerCalls
func (userPrivacySettingAllowPeerToPeerCalls *UserPrivacySettingAllowPeerToPeerCalls) MessageType() string {
	return "userPrivacySettingAllowPeerToPeerCalls"
}

// NewUserPrivacySettingAllowPeerToPeerCalls creates a new UserPrivacySettingAllowPeerToPeerCalls
//
func NewUserPrivacySettingAllowPeerToPeerCalls() *UserPrivacySettingAllowPeerToPeerCalls {
	userPrivacySettingAllowPeerToPeerCallsTemp := UserPrivacySettingAllowPeerToPeerCalls{
		tdCommon: tdCommon{Type: "userPrivacySettingAllowPeerToPeerCalls"},
	}

	return &userPrivacySettingAllowPeerToPeerCallsTemp
}

// GetUserPrivacySettingEnum return the enum type of this object
func (userPrivacySettingAllowPeerToPeerCalls *UserPrivacySettingAllowPeerToPeerCalls) GetUserPrivacySettingEnum() UserPrivacySettingEnum {
	return UserPrivacySettingAllowPeerToPeerCallsType
}

// UserPrivacySettingAllowFindingByPhoneNumber A privacy setting for managing whether the user can be found by their phone number. Checked only if the phone number is not known to the other user. Can be set only to "Allow contacts" or "Allow all"
type UserPrivacySettingAllowFindingByPhoneNumber struct {
	tdCommon
}

// MessageType return the string telegram-type of UserPrivacySettingAllowFindingByPhoneNumber
func (userPrivacySettingAllowFindingByPhoneNumber *UserPrivacySettingAllowFindingByPhoneNumber) MessageType() string {
	return "userPrivacySettingAllowFindingByPhoneNumber"
}

// NewUserPrivacySettingAllowFindingByPhoneNumber creates a new UserPrivacySettingAllowFindingByPhoneNumber
//
func NewUserPrivacySettingAllowFindingByPhoneNumber() *UserPrivacySettingAllowFindingByPhoneNumber {
	userPrivacySettingAllowFindingByPhoneNumberTemp := UserPrivacySettingAllowFindingByPhoneNumber{
		tdCommon: tdCommon{Type: "userPrivacySettingAllowFindingByPhoneNumber"},
	}

	return &userPrivacySettingAllowFindingByPhoneNumberTemp
}

// GetUserPrivacySettingEnum return the enum type of this object
func (userPrivacySettingAllowFindingByPhoneNumber *UserPrivacySettingAllowFindingByPhoneNumber) GetUserPrivacySettingEnum() UserPrivacySettingEnum {
	return UserPrivacySettingAllowFindingByPhoneNumberType
}
