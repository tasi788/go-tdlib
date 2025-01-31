package tdlib

import (
	"encoding/json"
	"fmt"
)

// CanTransferOwnershipResult Represents result of checking whether the current session can be used to transfer a chat ownership to another user
type CanTransferOwnershipResult interface {
	GetCanTransferOwnershipResultEnum() CanTransferOwnershipResultEnum
}

// CanTransferOwnershipResultEnum Alias for abstract CanTransferOwnershipResult 'Sub-Classes', used as constant-enum here
type CanTransferOwnershipResultEnum string

// CanTransferOwnershipResult enums
const (
	CanTransferOwnershipResultOkType               CanTransferOwnershipResultEnum = "canTransferOwnershipResultOk"
	CanTransferOwnershipResultPasswordNeededType   CanTransferOwnershipResultEnum = "canTransferOwnershipResultPasswordNeeded"
	CanTransferOwnershipResultPasswordTooFreshType CanTransferOwnershipResultEnum = "canTransferOwnershipResultPasswordTooFresh"
	CanTransferOwnershipResultSessionTooFreshType  CanTransferOwnershipResultEnum = "canTransferOwnershipResultSessionTooFresh"
)

func unmarshalCanTransferOwnershipResult(rawMsg *json.RawMessage) (CanTransferOwnershipResult, error) {

	if rawMsg == nil {
		return nil, nil
	}
	var objMap map[string]interface{}
	err := json.Unmarshal(*rawMsg, &objMap)
	if err != nil {
		return nil, err
	}

	switch CanTransferOwnershipResultEnum(objMap["@type"].(string)) {
	case CanTransferOwnershipResultOkType:
		var canTransferOwnershipResultOk CanTransferOwnershipResultOk
		err := json.Unmarshal(*rawMsg, &canTransferOwnershipResultOk)
		return &canTransferOwnershipResultOk, err

	case CanTransferOwnershipResultPasswordNeededType:
		var canTransferOwnershipResultPasswordNeeded CanTransferOwnershipResultPasswordNeeded
		err := json.Unmarshal(*rawMsg, &canTransferOwnershipResultPasswordNeeded)
		return &canTransferOwnershipResultPasswordNeeded, err

	case CanTransferOwnershipResultPasswordTooFreshType:
		var canTransferOwnershipResultPasswordTooFresh CanTransferOwnershipResultPasswordTooFresh
		err := json.Unmarshal(*rawMsg, &canTransferOwnershipResultPasswordTooFresh)
		return &canTransferOwnershipResultPasswordTooFresh, err

	case CanTransferOwnershipResultSessionTooFreshType:
		var canTransferOwnershipResultSessionTooFresh CanTransferOwnershipResultSessionTooFresh
		err := json.Unmarshal(*rawMsg, &canTransferOwnershipResultSessionTooFresh)
		return &canTransferOwnershipResultSessionTooFresh, err

	default:
		return nil, fmt.Errorf("Error UnMarshaling, unknown type:" + objMap["@type"].(string))
	}
}

// CanTransferOwnershipResultOk The session can be used
type CanTransferOwnershipResultOk struct {
	tdCommon
}

// MessageType return the string telegram-type of CanTransferOwnershipResultOk
func (canTransferOwnershipResultOk *CanTransferOwnershipResultOk) MessageType() string {
	return "canTransferOwnershipResultOk"
}

// NewCanTransferOwnershipResultOk creates a new CanTransferOwnershipResultOk
//
func NewCanTransferOwnershipResultOk() *CanTransferOwnershipResultOk {
	canTransferOwnershipResultOkTemp := CanTransferOwnershipResultOk{
		tdCommon: tdCommon{Type: "canTransferOwnershipResultOk"},
	}

	return &canTransferOwnershipResultOkTemp
}

// GetCanTransferOwnershipResultEnum return the enum type of this object
func (canTransferOwnershipResultOk *CanTransferOwnershipResultOk) GetCanTransferOwnershipResultEnum() CanTransferOwnershipResultEnum {
	return CanTransferOwnershipResultOkType
}

// CanTransferOwnershipResultPasswordNeeded The 2-step verification needs to be enabled first
type CanTransferOwnershipResultPasswordNeeded struct {
	tdCommon
}

// MessageType return the string telegram-type of CanTransferOwnershipResultPasswordNeeded
func (canTransferOwnershipResultPasswordNeeded *CanTransferOwnershipResultPasswordNeeded) MessageType() string {
	return "canTransferOwnershipResultPasswordNeeded"
}

// NewCanTransferOwnershipResultPasswordNeeded creates a new CanTransferOwnershipResultPasswordNeeded
//
func NewCanTransferOwnershipResultPasswordNeeded() *CanTransferOwnershipResultPasswordNeeded {
	canTransferOwnershipResultPasswordNeededTemp := CanTransferOwnershipResultPasswordNeeded{
		tdCommon: tdCommon{Type: "canTransferOwnershipResultPasswordNeeded"},
	}

	return &canTransferOwnershipResultPasswordNeededTemp
}

// GetCanTransferOwnershipResultEnum return the enum type of this object
func (canTransferOwnershipResultPasswordNeeded *CanTransferOwnershipResultPasswordNeeded) GetCanTransferOwnershipResultEnum() CanTransferOwnershipResultEnum {
	return CanTransferOwnershipResultPasswordNeededType
}

// CanTransferOwnershipResultPasswordTooFresh The 2-step verification was enabled recently, user needs to wait
type CanTransferOwnershipResultPasswordTooFresh struct {
	tdCommon
	RetryAfter int32 `json:"retry_after"` // Time left before the session can be used to transfer ownership of a chat, in seconds
}

// MessageType return the string telegram-type of CanTransferOwnershipResultPasswordTooFresh
func (canTransferOwnershipResultPasswordTooFresh *CanTransferOwnershipResultPasswordTooFresh) MessageType() string {
	return "canTransferOwnershipResultPasswordTooFresh"
}

// NewCanTransferOwnershipResultPasswordTooFresh creates a new CanTransferOwnershipResultPasswordTooFresh
//
// @param retryAfter Time left before the session can be used to transfer ownership of a chat, in seconds
func NewCanTransferOwnershipResultPasswordTooFresh(retryAfter int32) *CanTransferOwnershipResultPasswordTooFresh {
	canTransferOwnershipResultPasswordTooFreshTemp := CanTransferOwnershipResultPasswordTooFresh{
		tdCommon:   tdCommon{Type: "canTransferOwnershipResultPasswordTooFresh"},
		RetryAfter: retryAfter,
	}

	return &canTransferOwnershipResultPasswordTooFreshTemp
}

// GetCanTransferOwnershipResultEnum return the enum type of this object
func (canTransferOwnershipResultPasswordTooFresh *CanTransferOwnershipResultPasswordTooFresh) GetCanTransferOwnershipResultEnum() CanTransferOwnershipResultEnum {
	return CanTransferOwnershipResultPasswordTooFreshType
}

// CanTransferOwnershipResultSessionTooFresh The session was created recently, user needs to wait
type CanTransferOwnershipResultSessionTooFresh struct {
	tdCommon
	RetryAfter int32 `json:"retry_after"` // Time left before the session can be used to transfer ownership of a chat, in seconds
}

// MessageType return the string telegram-type of CanTransferOwnershipResultSessionTooFresh
func (canTransferOwnershipResultSessionTooFresh *CanTransferOwnershipResultSessionTooFresh) MessageType() string {
	return "canTransferOwnershipResultSessionTooFresh"
}

// NewCanTransferOwnershipResultSessionTooFresh creates a new CanTransferOwnershipResultSessionTooFresh
//
// @param retryAfter Time left before the session can be used to transfer ownership of a chat, in seconds
func NewCanTransferOwnershipResultSessionTooFresh(retryAfter int32) *CanTransferOwnershipResultSessionTooFresh {
	canTransferOwnershipResultSessionTooFreshTemp := CanTransferOwnershipResultSessionTooFresh{
		tdCommon:   tdCommon{Type: "canTransferOwnershipResultSessionTooFresh"},
		RetryAfter: retryAfter,
	}

	return &canTransferOwnershipResultSessionTooFreshTemp
}

// GetCanTransferOwnershipResultEnum return the enum type of this object
func (canTransferOwnershipResultSessionTooFresh *CanTransferOwnershipResultSessionTooFresh) GetCanTransferOwnershipResultEnum() CanTransferOwnershipResultEnum {
	return CanTransferOwnershipResultSessionTooFreshType
}

// CanTransferOwnership Checks whether the current session can be used to transfer a chat ownership to another user
func (client *Client) CanTransferOwnership() (CanTransferOwnershipResult, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "canTransferOwnership",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	switch CanTransferOwnershipResultEnum(result.Data["@type"].(string)) {

	case CanTransferOwnershipResultOkType:
		var canTransferOwnershipResult CanTransferOwnershipResultOk
		err = json.Unmarshal(result.Raw, &canTransferOwnershipResult)
		return &canTransferOwnershipResult, err

	case CanTransferOwnershipResultPasswordNeededType:
		var canTransferOwnershipResult CanTransferOwnershipResultPasswordNeeded
		err = json.Unmarshal(result.Raw, &canTransferOwnershipResult)
		return &canTransferOwnershipResult, err

	case CanTransferOwnershipResultPasswordTooFreshType:
		var canTransferOwnershipResult CanTransferOwnershipResultPasswordTooFresh
		err = json.Unmarshal(result.Raw, &canTransferOwnershipResult)
		return &canTransferOwnershipResult, err

	case CanTransferOwnershipResultSessionTooFreshType:
		var canTransferOwnershipResult CanTransferOwnershipResultSessionTooFresh
		err = json.Unmarshal(result.Raw, &canTransferOwnershipResult)
		return &canTransferOwnershipResult, err

	default:
		return nil, fmt.Errorf("Invalid type")
	}
}
