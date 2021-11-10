package tdlib

import (
	"encoding/json"
	"fmt"
)

// CallId Contains the call identifier
type CallId struct {
	tdCommon
	Id int32 `json:"id"` // Call identifier
}

// MessageType return the string telegram-type of CallId
func (callId *CallId) MessageType() string {
	return "callId"
}

// NewCallId creates a new CallId
//
// @param id Call identifier
func NewCallId(id int32) *CallId {
	callIdTemp := CallId{
		tdCommon: tdCommon{Type: "callId"},
		Id:       id,
	}

	return &callIdTemp
}

// CreateCall Creates a new call
// @param userId Identifier of the user to be called
// @param protocol The call protocols supported by the application
// @param isVideo True, if a video call needs to be created
func (client *Client) CreateCall(userId int64, protocol *CallProtocol, isVideo bool) (*CallId, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":    "createCall",
		"user_id":  userId,
		"protocol": protocol,
		"is_video": isVideo,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var callId CallId
	err = json.Unmarshal(result.Raw, &callId)
	return &callId, err
}
