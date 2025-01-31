package tdlib

import (
	"encoding/json"
	"fmt"
)

// ResetPasswordResult Represents result of 2-step verification password reset
type ResetPasswordResult interface {
	GetResetPasswordResultEnum() ResetPasswordResultEnum
}

// ResetPasswordResultEnum Alias for abstract ResetPasswordResult 'Sub-Classes', used as constant-enum here
type ResetPasswordResultEnum string

// ResetPasswordResult enums
const (
	ResetPasswordResultOkType       ResetPasswordResultEnum = "resetPasswordResultOk"
	ResetPasswordResultPendingType  ResetPasswordResultEnum = "resetPasswordResultPending"
	ResetPasswordResultDeclinedType ResetPasswordResultEnum = "resetPasswordResultDeclined"
)

func unmarshalResetPasswordResult(rawMsg *json.RawMessage) (ResetPasswordResult, error) {

	if rawMsg == nil {
		return nil, nil
	}
	var objMap map[string]interface{}
	err := json.Unmarshal(*rawMsg, &objMap)
	if err != nil {
		return nil, err
	}

	switch ResetPasswordResultEnum(objMap["@type"].(string)) {
	case ResetPasswordResultOkType:
		var resetPasswordResultOk ResetPasswordResultOk
		err := json.Unmarshal(*rawMsg, &resetPasswordResultOk)
		return &resetPasswordResultOk, err

	case ResetPasswordResultPendingType:
		var resetPasswordResultPending ResetPasswordResultPending
		err := json.Unmarshal(*rawMsg, &resetPasswordResultPending)
		return &resetPasswordResultPending, err

	case ResetPasswordResultDeclinedType:
		var resetPasswordResultDeclined ResetPasswordResultDeclined
		err := json.Unmarshal(*rawMsg, &resetPasswordResultDeclined)
		return &resetPasswordResultDeclined, err

	default:
		return nil, fmt.Errorf("Error UnMarshaling, unknown type:" + objMap["@type"].(string))
	}
}

// ResetPasswordResultOk The password was reset
type ResetPasswordResultOk struct {
	tdCommon
}

// MessageType return the string telegram-type of ResetPasswordResultOk
func (resetPasswordResultOk *ResetPasswordResultOk) MessageType() string {
	return "resetPasswordResultOk"
}

// NewResetPasswordResultOk creates a new ResetPasswordResultOk
//
func NewResetPasswordResultOk() *ResetPasswordResultOk {
	resetPasswordResultOkTemp := ResetPasswordResultOk{
		tdCommon: tdCommon{Type: "resetPasswordResultOk"},
	}

	return &resetPasswordResultOkTemp
}

// GetResetPasswordResultEnum return the enum type of this object
func (resetPasswordResultOk *ResetPasswordResultOk) GetResetPasswordResultEnum() ResetPasswordResultEnum {
	return ResetPasswordResultOkType
}

// ResetPasswordResultPending The password reset request is pending
type ResetPasswordResultPending struct {
	tdCommon
	PendingResetDate int32 `json:"pending_reset_date"` // Point in time (Unix timestamp) after which the password can be reset immediately using resetPassword
}

// MessageType return the string telegram-type of ResetPasswordResultPending
func (resetPasswordResultPending *ResetPasswordResultPending) MessageType() string {
	return "resetPasswordResultPending"
}

// NewResetPasswordResultPending creates a new ResetPasswordResultPending
//
// @param pendingResetDate Point in time (Unix timestamp) after which the password can be reset immediately using resetPassword
func NewResetPasswordResultPending(pendingResetDate int32) *ResetPasswordResultPending {
	resetPasswordResultPendingTemp := ResetPasswordResultPending{
		tdCommon:         tdCommon{Type: "resetPasswordResultPending"},
		PendingResetDate: pendingResetDate,
	}

	return &resetPasswordResultPendingTemp
}

// GetResetPasswordResultEnum return the enum type of this object
func (resetPasswordResultPending *ResetPasswordResultPending) GetResetPasswordResultEnum() ResetPasswordResultEnum {
	return ResetPasswordResultPendingType
}

// ResetPasswordResultDeclined The password reset request was declined
type ResetPasswordResultDeclined struct {
	tdCommon
	RetryDate int32 `json:"retry_date"` // Point in time (Unix timestamp) when the password reset can be retried
}

// MessageType return the string telegram-type of ResetPasswordResultDeclined
func (resetPasswordResultDeclined *ResetPasswordResultDeclined) MessageType() string {
	return "resetPasswordResultDeclined"
}

// NewResetPasswordResultDeclined creates a new ResetPasswordResultDeclined
//
// @param retryDate Point in time (Unix timestamp) when the password reset can be retried
func NewResetPasswordResultDeclined(retryDate int32) *ResetPasswordResultDeclined {
	resetPasswordResultDeclinedTemp := ResetPasswordResultDeclined{
		tdCommon:  tdCommon{Type: "resetPasswordResultDeclined"},
		RetryDate: retryDate,
	}

	return &resetPasswordResultDeclinedTemp
}

// GetResetPasswordResultEnum return the enum type of this object
func (resetPasswordResultDeclined *ResetPasswordResultDeclined) GetResetPasswordResultEnum() ResetPasswordResultEnum {
	return ResetPasswordResultDeclinedType
}

// ResetPassword Removes 2-step verification password without previous password and access to recovery email address. The password can't be reset immediately and the request needs to be repeated after the specified time
func (client *Client) ResetPassword() (ResetPasswordResult, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "resetPassword",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	switch ResetPasswordResultEnum(result.Data["@type"].(string)) {

	case ResetPasswordResultOkType:
		var resetPasswordResult ResetPasswordResultOk
		err = json.Unmarshal(result.Raw, &resetPasswordResult)
		return &resetPasswordResult, err

	case ResetPasswordResultPendingType:
		var resetPasswordResult ResetPasswordResultPending
		err = json.Unmarshal(result.Raw, &resetPasswordResult)
		return &resetPasswordResult, err

	case ResetPasswordResultDeclinedType:
		var resetPasswordResult ResetPasswordResultDeclined
		err = json.Unmarshal(result.Raw, &resetPasswordResult)
		return &resetPasswordResult, err

	default:
		return nil, fmt.Errorf("Invalid type")
	}
}
