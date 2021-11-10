package tdlib

import (
	"encoding/json"
)

// Call Describes a call
type Call struct {
	tdCommon
	Id         int32     `json:"id"`          // Call identifier, not persistent
	UserId     int64     `json:"user_id"`     // Peer user identifier
	IsOutgoing bool      `json:"is_outgoing"` // True, if the call is outgoing
	IsVideo    bool      `json:"is_video"`    // True, if the call is a video call
	State      CallState `json:"state"`       // Call state
}

// MessageType return the string telegram-type of Call
func (call *Call) MessageType() string {
	return "call"
}

// NewCall creates a new Call
//
// @param id Call identifier, not persistent
// @param userId Peer user identifier
// @param isOutgoing True, if the call is outgoing
// @param isVideo True, if the call is a video call
// @param state Call state
func NewCall(id int32, userId int64, isOutgoing bool, isVideo bool, state CallState) *Call {
	callTemp := Call{
		tdCommon:   tdCommon{Type: "call"},
		Id:         id,
		UserId:     userId,
		IsOutgoing: isOutgoing,
		IsVideo:    isVideo,
		State:      state,
	}

	return &callTemp
}

// UnmarshalJSON unmarshal to json
func (call *Call) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	tempObj := struct {
		tdCommon
		Id         int32 `json:"id"`          // Call identifier, not persistent
		UserId     int64 `json:"user_id"`     // Peer user identifier
		IsOutgoing bool  `json:"is_outgoing"` // True, if the call is outgoing
		IsVideo    bool  `json:"is_video"`    // True, if the call is a video call

	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	call.tdCommon = tempObj.tdCommon
	call.Id = tempObj.Id
	call.UserId = tempObj.UserId
	call.IsOutgoing = tempObj.IsOutgoing
	call.IsVideo = tempObj.IsVideo

	fieldState, _ := unmarshalCallState(objMap["state"])
	call.State = fieldState

	return nil
}
