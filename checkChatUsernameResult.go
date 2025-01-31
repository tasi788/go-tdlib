package tdlib

import (
	"encoding/json"
	"fmt"
)

// CheckChatUsernameResult Represents result of checking whether a username can be set for a chat
type CheckChatUsernameResult interface {
	GetCheckChatUsernameResultEnum() CheckChatUsernameResultEnum
}

// CheckChatUsernameResultEnum Alias for abstract CheckChatUsernameResult 'Sub-Classes', used as constant-enum here
type CheckChatUsernameResultEnum string

// CheckChatUsernameResult enums
const (
	CheckChatUsernameResultOkType                      CheckChatUsernameResultEnum = "checkChatUsernameResultOk"
	CheckChatUsernameResultUsernameInvalidType         CheckChatUsernameResultEnum = "checkChatUsernameResultUsernameInvalid"
	CheckChatUsernameResultUsernameOccupiedType        CheckChatUsernameResultEnum = "checkChatUsernameResultUsernameOccupied"
	CheckChatUsernameResultPublicChatsTooMuchType      CheckChatUsernameResultEnum = "checkChatUsernameResultPublicChatsTooMuch"
	CheckChatUsernameResultPublicGroupsUnavailableType CheckChatUsernameResultEnum = "checkChatUsernameResultPublicGroupsUnavailable"
)

func unmarshalCheckChatUsernameResult(rawMsg *json.RawMessage) (CheckChatUsernameResult, error) {

	if rawMsg == nil {
		return nil, nil
	}
	var objMap map[string]interface{}
	err := json.Unmarshal(*rawMsg, &objMap)
	if err != nil {
		return nil, err
	}

	switch CheckChatUsernameResultEnum(objMap["@type"].(string)) {
	case CheckChatUsernameResultOkType:
		var checkChatUsernameResultOk CheckChatUsernameResultOk
		err := json.Unmarshal(*rawMsg, &checkChatUsernameResultOk)
		return &checkChatUsernameResultOk, err

	case CheckChatUsernameResultUsernameInvalidType:
		var checkChatUsernameResultUsernameInvalid CheckChatUsernameResultUsernameInvalid
		err := json.Unmarshal(*rawMsg, &checkChatUsernameResultUsernameInvalid)
		return &checkChatUsernameResultUsernameInvalid, err

	case CheckChatUsernameResultUsernameOccupiedType:
		var checkChatUsernameResultUsernameOccupied CheckChatUsernameResultUsernameOccupied
		err := json.Unmarshal(*rawMsg, &checkChatUsernameResultUsernameOccupied)
		return &checkChatUsernameResultUsernameOccupied, err

	case CheckChatUsernameResultPublicChatsTooMuchType:
		var checkChatUsernameResultPublicChatsTooMuch CheckChatUsernameResultPublicChatsTooMuch
		err := json.Unmarshal(*rawMsg, &checkChatUsernameResultPublicChatsTooMuch)
		return &checkChatUsernameResultPublicChatsTooMuch, err

	case CheckChatUsernameResultPublicGroupsUnavailableType:
		var checkChatUsernameResultPublicGroupsUnavailable CheckChatUsernameResultPublicGroupsUnavailable
		err := json.Unmarshal(*rawMsg, &checkChatUsernameResultPublicGroupsUnavailable)
		return &checkChatUsernameResultPublicGroupsUnavailable, err

	default:
		return nil, fmt.Errorf("Error UnMarshaling, unknown type:" + objMap["@type"].(string))
	}
}

// CheckChatUsernameResultOk The username can be set
type CheckChatUsernameResultOk struct {
	tdCommon
}

// MessageType return the string telegram-type of CheckChatUsernameResultOk
func (checkChatUsernameResultOk *CheckChatUsernameResultOk) MessageType() string {
	return "checkChatUsernameResultOk"
}

// NewCheckChatUsernameResultOk creates a new CheckChatUsernameResultOk
//
func NewCheckChatUsernameResultOk() *CheckChatUsernameResultOk {
	checkChatUsernameResultOkTemp := CheckChatUsernameResultOk{
		tdCommon: tdCommon{Type: "checkChatUsernameResultOk"},
	}

	return &checkChatUsernameResultOkTemp
}

// GetCheckChatUsernameResultEnum return the enum type of this object
func (checkChatUsernameResultOk *CheckChatUsernameResultOk) GetCheckChatUsernameResultEnum() CheckChatUsernameResultEnum {
	return CheckChatUsernameResultOkType
}

// CheckChatUsernameResultUsernameInvalid The username is invalid
type CheckChatUsernameResultUsernameInvalid struct {
	tdCommon
}

// MessageType return the string telegram-type of CheckChatUsernameResultUsernameInvalid
func (checkChatUsernameResultUsernameInvalid *CheckChatUsernameResultUsernameInvalid) MessageType() string {
	return "checkChatUsernameResultUsernameInvalid"
}

// NewCheckChatUsernameResultUsernameInvalid creates a new CheckChatUsernameResultUsernameInvalid
//
func NewCheckChatUsernameResultUsernameInvalid() *CheckChatUsernameResultUsernameInvalid {
	checkChatUsernameResultUsernameInvalidTemp := CheckChatUsernameResultUsernameInvalid{
		tdCommon: tdCommon{Type: "checkChatUsernameResultUsernameInvalid"},
	}

	return &checkChatUsernameResultUsernameInvalidTemp
}

// GetCheckChatUsernameResultEnum return the enum type of this object
func (checkChatUsernameResultUsernameInvalid *CheckChatUsernameResultUsernameInvalid) GetCheckChatUsernameResultEnum() CheckChatUsernameResultEnum {
	return CheckChatUsernameResultUsernameInvalidType
}

// CheckChatUsernameResultUsernameOccupied The username is occupied
type CheckChatUsernameResultUsernameOccupied struct {
	tdCommon
}

// MessageType return the string telegram-type of CheckChatUsernameResultUsernameOccupied
func (checkChatUsernameResultUsernameOccupied *CheckChatUsernameResultUsernameOccupied) MessageType() string {
	return "checkChatUsernameResultUsernameOccupied"
}

// NewCheckChatUsernameResultUsernameOccupied creates a new CheckChatUsernameResultUsernameOccupied
//
func NewCheckChatUsernameResultUsernameOccupied() *CheckChatUsernameResultUsernameOccupied {
	checkChatUsernameResultUsernameOccupiedTemp := CheckChatUsernameResultUsernameOccupied{
		tdCommon: tdCommon{Type: "checkChatUsernameResultUsernameOccupied"},
	}

	return &checkChatUsernameResultUsernameOccupiedTemp
}

// GetCheckChatUsernameResultEnum return the enum type of this object
func (checkChatUsernameResultUsernameOccupied *CheckChatUsernameResultUsernameOccupied) GetCheckChatUsernameResultEnum() CheckChatUsernameResultEnum {
	return CheckChatUsernameResultUsernameOccupiedType
}

// CheckChatUsernameResultPublicChatsTooMuch The user has too much chats with username, one of them must be made private first
type CheckChatUsernameResultPublicChatsTooMuch struct {
	tdCommon
}

// MessageType return the string telegram-type of CheckChatUsernameResultPublicChatsTooMuch
func (checkChatUsernameResultPublicChatsTooMuch *CheckChatUsernameResultPublicChatsTooMuch) MessageType() string {
	return "checkChatUsernameResultPublicChatsTooMuch"
}

// NewCheckChatUsernameResultPublicChatsTooMuch creates a new CheckChatUsernameResultPublicChatsTooMuch
//
func NewCheckChatUsernameResultPublicChatsTooMuch() *CheckChatUsernameResultPublicChatsTooMuch {
	checkChatUsernameResultPublicChatsTooMuchTemp := CheckChatUsernameResultPublicChatsTooMuch{
		tdCommon: tdCommon{Type: "checkChatUsernameResultPublicChatsTooMuch"},
	}

	return &checkChatUsernameResultPublicChatsTooMuchTemp
}

// GetCheckChatUsernameResultEnum return the enum type of this object
func (checkChatUsernameResultPublicChatsTooMuch *CheckChatUsernameResultPublicChatsTooMuch) GetCheckChatUsernameResultEnum() CheckChatUsernameResultEnum {
	return CheckChatUsernameResultPublicChatsTooMuchType
}

// CheckChatUsernameResultPublicGroupsUnavailable The user can't be a member of a public supergroup
type CheckChatUsernameResultPublicGroupsUnavailable struct {
	tdCommon
}

// MessageType return the string telegram-type of CheckChatUsernameResultPublicGroupsUnavailable
func (checkChatUsernameResultPublicGroupsUnavailable *CheckChatUsernameResultPublicGroupsUnavailable) MessageType() string {
	return "checkChatUsernameResultPublicGroupsUnavailable"
}

// NewCheckChatUsernameResultPublicGroupsUnavailable creates a new CheckChatUsernameResultPublicGroupsUnavailable
//
func NewCheckChatUsernameResultPublicGroupsUnavailable() *CheckChatUsernameResultPublicGroupsUnavailable {
	checkChatUsernameResultPublicGroupsUnavailableTemp := CheckChatUsernameResultPublicGroupsUnavailable{
		tdCommon: tdCommon{Type: "checkChatUsernameResultPublicGroupsUnavailable"},
	}

	return &checkChatUsernameResultPublicGroupsUnavailableTemp
}

// GetCheckChatUsernameResultEnum return the enum type of this object
func (checkChatUsernameResultPublicGroupsUnavailable *CheckChatUsernameResultPublicGroupsUnavailable) GetCheckChatUsernameResultEnum() CheckChatUsernameResultEnum {
	return CheckChatUsernameResultPublicGroupsUnavailableType
}

// CheckChatUsername Checks whether a username can be set for a chat
// @param chatId Chat identifier; must be identifier of a supergroup chat, or a channel chat, or a private chat with self, or zero if the chat is being created
// @param username Username to be checked
func (client *Client) CheckChatUsername(chatId int64, username string) (CheckChatUsernameResult, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":    "checkChatUsername",
		"chat_id":  chatId,
		"username": username,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	switch CheckChatUsernameResultEnum(result.Data["@type"].(string)) {

	case CheckChatUsernameResultOkType:
		var checkChatUsernameResult CheckChatUsernameResultOk
		err = json.Unmarshal(result.Raw, &checkChatUsernameResult)
		return &checkChatUsernameResult, err

	case CheckChatUsernameResultUsernameInvalidType:
		var checkChatUsernameResult CheckChatUsernameResultUsernameInvalid
		err = json.Unmarshal(result.Raw, &checkChatUsernameResult)
		return &checkChatUsernameResult, err

	case CheckChatUsernameResultUsernameOccupiedType:
		var checkChatUsernameResult CheckChatUsernameResultUsernameOccupied
		err = json.Unmarshal(result.Raw, &checkChatUsernameResult)
		return &checkChatUsernameResult, err

	case CheckChatUsernameResultPublicChatsTooMuchType:
		var checkChatUsernameResult CheckChatUsernameResultPublicChatsTooMuch
		err = json.Unmarshal(result.Raw, &checkChatUsernameResult)
		return &checkChatUsernameResult, err

	case CheckChatUsernameResultPublicGroupsUnavailableType:
		var checkChatUsernameResult CheckChatUsernameResultPublicGroupsUnavailable
		err = json.Unmarshal(result.Raw, &checkChatUsernameResult)
		return &checkChatUsernameResult, err

	default:
		return nil, fmt.Errorf("Invalid type")
	}
}
