package tdlib

import (
	"encoding/json"
)

// VideoChat Describes a video chat
type VideoChat struct {
	tdCommon
	GroupCallId          int32         `json:"group_call_id"`          // Group call identifier of an active video chat; 0 if none. Full information about the video chat can be received through the method getGroupCall
	HasParticipants      bool          `json:"has_participants"`       // True, if the video chat has participants
	DefaultParticipantId MessageSender `json:"default_participant_id"` // Default group call participant identifier to join the video chat; may be null
}

// MessageType return the string telegram-type of VideoChat
func (videoChat *VideoChat) MessageType() string {
	return "videoChat"
}

// NewVideoChat creates a new VideoChat
//
// @param groupCallId Group call identifier of an active video chat; 0 if none. Full information about the video chat can be received through the method getGroupCall
// @param hasParticipants True, if the video chat has participants
// @param defaultParticipantId Default group call participant identifier to join the video chat; may be null
func NewVideoChat(groupCallId int32, hasParticipants bool, defaultParticipantId MessageSender) *VideoChat {
	videoChatTemp := VideoChat{
		tdCommon:             tdCommon{Type: "videoChat"},
		GroupCallId:          groupCallId,
		HasParticipants:      hasParticipants,
		DefaultParticipantId: defaultParticipantId,
	}

	return &videoChatTemp
}

// UnmarshalJSON unmarshal to json
func (videoChat *VideoChat) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	tempObj := struct {
		tdCommon
		GroupCallId     int32 `json:"group_call_id"`    // Group call identifier of an active video chat; 0 if none. Full information about the video chat can be received through the method getGroupCall
		HasParticipants bool  `json:"has_participants"` // True, if the video chat has participants

	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	videoChat.tdCommon = tempObj.tdCommon
	videoChat.GroupCallId = tempObj.GroupCallId
	videoChat.HasParticipants = tempObj.HasParticipants

	fieldDefaultParticipantId, _ := unmarshalMessageSender(objMap["default_participant_id"])
	videoChat.DefaultParticipantId = fieldDefaultParticipantId

	return nil
}
