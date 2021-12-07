package tdlib

import (
	"encoding/json"
	"fmt"
)

// SuggestedAction Describes an action suggested to the current user
type SuggestedAction interface {
	GetSuggestedActionEnum() SuggestedActionEnum
}

// SuggestedActionEnum Alias for abstract SuggestedAction 'Sub-Classes', used as constant-enum here
type SuggestedActionEnum string

// SuggestedAction enums
const (
	SuggestedActionEnableArchiveAndMuteNewChatsType SuggestedActionEnum = "suggestedActionEnableArchiveAndMuteNewChats"
	SuggestedActionCheckPasswordType                SuggestedActionEnum = "suggestedActionCheckPassword"
	SuggestedActionCheckPhoneNumberType             SuggestedActionEnum = "suggestedActionCheckPhoneNumber"
	SuggestedActionSeeTicksHintType                 SuggestedActionEnum = "suggestedActionSeeTicksHint"
	SuggestedActionConvertToBroadcastGroupType      SuggestedActionEnum = "suggestedActionConvertToBroadcastGroup"
	SuggestedActionSetPasswordType                  SuggestedActionEnum = "suggestedActionSetPassword"
)

func unmarshalSuggestedAction(rawMsg *json.RawMessage) (SuggestedAction, error) {

	if rawMsg == nil {
		return nil, nil
	}
	var objMap map[string]interface{}
	err := json.Unmarshal(*rawMsg, &objMap)
	if err != nil {
		return nil, err
	}

	switch SuggestedActionEnum(objMap["@type"].(string)) {
	case SuggestedActionEnableArchiveAndMuteNewChatsType:
		var suggestedActionEnableArchiveAndMuteNewChats SuggestedActionEnableArchiveAndMuteNewChats
		err := json.Unmarshal(*rawMsg, &suggestedActionEnableArchiveAndMuteNewChats)
		return &suggestedActionEnableArchiveAndMuteNewChats, err

	case SuggestedActionCheckPasswordType:
		var suggestedActionCheckPassword SuggestedActionCheckPassword
		err := json.Unmarshal(*rawMsg, &suggestedActionCheckPassword)
		return &suggestedActionCheckPassword, err

	case SuggestedActionCheckPhoneNumberType:
		var suggestedActionCheckPhoneNumber SuggestedActionCheckPhoneNumber
		err := json.Unmarshal(*rawMsg, &suggestedActionCheckPhoneNumber)
		return &suggestedActionCheckPhoneNumber, err

	case SuggestedActionSeeTicksHintType:
		var suggestedActionSeeTicksHint SuggestedActionSeeTicksHint
		err := json.Unmarshal(*rawMsg, &suggestedActionSeeTicksHint)
		return &suggestedActionSeeTicksHint, err

	case SuggestedActionConvertToBroadcastGroupType:
		var suggestedActionConvertToBroadcastGroup SuggestedActionConvertToBroadcastGroup
		err := json.Unmarshal(*rawMsg, &suggestedActionConvertToBroadcastGroup)
		return &suggestedActionConvertToBroadcastGroup, err

	case SuggestedActionSetPasswordType:
		var suggestedActionSetPassword SuggestedActionSetPassword
		err := json.Unmarshal(*rawMsg, &suggestedActionSetPassword)
		return &suggestedActionSetPassword, err

	default:
		return nil, fmt.Errorf("Error UnMarshaling, unknown type:" + objMap["@type"].(string))
	}
}

// SuggestedActionEnableArchiveAndMuteNewChats Suggests the user to enable "archive_and_mute_new_chats_from_unknown_users" option
type SuggestedActionEnableArchiveAndMuteNewChats struct {
	tdCommon
}

// MessageType return the string telegram-type of SuggestedActionEnableArchiveAndMuteNewChats
func (suggestedActionEnableArchiveAndMuteNewChats *SuggestedActionEnableArchiveAndMuteNewChats) MessageType() string {
	return "suggestedActionEnableArchiveAndMuteNewChats"
}

// NewSuggestedActionEnableArchiveAndMuteNewChats creates a new SuggestedActionEnableArchiveAndMuteNewChats
//
func NewSuggestedActionEnableArchiveAndMuteNewChats() *SuggestedActionEnableArchiveAndMuteNewChats {
	suggestedActionEnableArchiveAndMuteNewChatsTemp := SuggestedActionEnableArchiveAndMuteNewChats{
		tdCommon: tdCommon{Type: "suggestedActionEnableArchiveAndMuteNewChats"},
	}

	return &suggestedActionEnableArchiveAndMuteNewChatsTemp
}

// GetSuggestedActionEnum return the enum type of this object
func (suggestedActionEnableArchiveAndMuteNewChats *SuggestedActionEnableArchiveAndMuteNewChats) GetSuggestedActionEnum() SuggestedActionEnum {
	return SuggestedActionEnableArchiveAndMuteNewChatsType
}

// SuggestedActionCheckPassword Suggests the user to check whether 2-step verification password is still remembered
type SuggestedActionCheckPassword struct {
	tdCommon
}

// MessageType return the string telegram-type of SuggestedActionCheckPassword
func (suggestedActionCheckPassword *SuggestedActionCheckPassword) MessageType() string {
	return "suggestedActionCheckPassword"
}

// NewSuggestedActionCheckPassword creates a new SuggestedActionCheckPassword
//
func NewSuggestedActionCheckPassword() *SuggestedActionCheckPassword {
	suggestedActionCheckPasswordTemp := SuggestedActionCheckPassword{
		tdCommon: tdCommon{Type: "suggestedActionCheckPassword"},
	}

	return &suggestedActionCheckPasswordTemp
}

// GetSuggestedActionEnum return the enum type of this object
func (suggestedActionCheckPassword *SuggestedActionCheckPassword) GetSuggestedActionEnum() SuggestedActionEnum {
	return SuggestedActionCheckPasswordType
}

// SuggestedActionCheckPhoneNumber Suggests the user to check whether authorization phone number is correct and change the phone number if it is inaccessible
type SuggestedActionCheckPhoneNumber struct {
	tdCommon
}

// MessageType return the string telegram-type of SuggestedActionCheckPhoneNumber
func (suggestedActionCheckPhoneNumber *SuggestedActionCheckPhoneNumber) MessageType() string {
	return "suggestedActionCheckPhoneNumber"
}

// NewSuggestedActionCheckPhoneNumber creates a new SuggestedActionCheckPhoneNumber
//
func NewSuggestedActionCheckPhoneNumber() *SuggestedActionCheckPhoneNumber {
	suggestedActionCheckPhoneNumberTemp := SuggestedActionCheckPhoneNumber{
		tdCommon: tdCommon{Type: "suggestedActionCheckPhoneNumber"},
	}

	return &suggestedActionCheckPhoneNumberTemp
}

// GetSuggestedActionEnum return the enum type of this object
func (suggestedActionCheckPhoneNumber *SuggestedActionCheckPhoneNumber) GetSuggestedActionEnum() SuggestedActionEnum {
	return SuggestedActionCheckPhoneNumberType
}

// SuggestedActionSeeTicksHint Suggests the user to see a hint about meaning of one and two ticks on sent message
type SuggestedActionSeeTicksHint struct {
	tdCommon
}

// MessageType return the string telegram-type of SuggestedActionSeeTicksHint
func (suggestedActionSeeTicksHint *SuggestedActionSeeTicksHint) MessageType() string {
	return "suggestedActionSeeTicksHint"
}

// NewSuggestedActionSeeTicksHint creates a new SuggestedActionSeeTicksHint
//
func NewSuggestedActionSeeTicksHint() *SuggestedActionSeeTicksHint {
	suggestedActionSeeTicksHintTemp := SuggestedActionSeeTicksHint{
		tdCommon: tdCommon{Type: "suggestedActionSeeTicksHint"},
	}

	return &suggestedActionSeeTicksHintTemp
}

// GetSuggestedActionEnum return the enum type of this object
func (suggestedActionSeeTicksHint *SuggestedActionSeeTicksHint) GetSuggestedActionEnum() SuggestedActionEnum {
	return SuggestedActionSeeTicksHintType
}

// SuggestedActionConvertToBroadcastGroup Suggests the user to convert specified supergroup to a broadcast group
type SuggestedActionConvertToBroadcastGroup struct {
	tdCommon
	SupergroupId int64 `json:"supergroup_id"` // Supergroup identifier
}

// MessageType return the string telegram-type of SuggestedActionConvertToBroadcastGroup
func (suggestedActionConvertToBroadcastGroup *SuggestedActionConvertToBroadcastGroup) MessageType() string {
	return "suggestedActionConvertToBroadcastGroup"
}

// NewSuggestedActionConvertToBroadcastGroup creates a new SuggestedActionConvertToBroadcastGroup
//
// @param supergroupId Supergroup identifier
func NewSuggestedActionConvertToBroadcastGroup(supergroupId int64) *SuggestedActionConvertToBroadcastGroup {
	suggestedActionConvertToBroadcastGroupTemp := SuggestedActionConvertToBroadcastGroup{
		tdCommon:     tdCommon{Type: "suggestedActionConvertToBroadcastGroup"},
		SupergroupId: supergroupId,
	}

	return &suggestedActionConvertToBroadcastGroupTemp
}

// GetSuggestedActionEnum return the enum type of this object
func (suggestedActionConvertToBroadcastGroup *SuggestedActionConvertToBroadcastGroup) GetSuggestedActionEnum() SuggestedActionEnum {
	return SuggestedActionConvertToBroadcastGroupType
}

// SuggestedActionSetPassword Suggests the user to set a 2-step verification password to be able to log in again
type SuggestedActionSetPassword struct {
	tdCommon
	AuthorizationDelay int32 `json:"authorization_delay"` // The number of days to pass between consecutive authorizations if the user declines to set password
}

// MessageType return the string telegram-type of SuggestedActionSetPassword
func (suggestedActionSetPassword *SuggestedActionSetPassword) MessageType() string {
	return "suggestedActionSetPassword"
}

// NewSuggestedActionSetPassword creates a new SuggestedActionSetPassword
//
// @param authorizationDelay The number of days to pass between consecutive authorizations if the user declines to set password
func NewSuggestedActionSetPassword(authorizationDelay int32) *SuggestedActionSetPassword {
	suggestedActionSetPasswordTemp := SuggestedActionSetPassword{
		tdCommon:           tdCommon{Type: "suggestedActionSetPassword"},
		AuthorizationDelay: authorizationDelay,
	}

	return &suggestedActionSetPasswordTemp
}

// GetSuggestedActionEnum return the enum type of this object
func (suggestedActionSetPassword *SuggestedActionSetPassword) GetSuggestedActionEnum() SuggestedActionEnum {
	return SuggestedActionSetPasswordType
}
