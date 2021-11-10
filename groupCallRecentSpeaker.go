package tdlib

import (
	"encoding/json"
)

// GroupCallRecentSpeaker Describes a recently speaking participant in a group call
type GroupCallRecentSpeaker struct {
	tdCommon
	ParticipantId MessageSender `json:"participant_id"` // Group call participant identifier
	IsSpeaking    bool          `json:"is_speaking"`    // True, is the user has spoken recently
}

// MessageType return the string telegram-type of GroupCallRecentSpeaker
func (groupCallRecentSpeaker *GroupCallRecentSpeaker) MessageType() string {
	return "groupCallRecentSpeaker"
}

// NewGroupCallRecentSpeaker creates a new GroupCallRecentSpeaker
//
// @param participantId Group call participant identifier
// @param isSpeaking True, is the user has spoken recently
func NewGroupCallRecentSpeaker(participantId MessageSender, isSpeaking bool) *GroupCallRecentSpeaker {
	groupCallRecentSpeakerTemp := GroupCallRecentSpeaker{
		tdCommon:      tdCommon{Type: "groupCallRecentSpeaker"},
		ParticipantId: participantId,
		IsSpeaking:    isSpeaking,
	}

	return &groupCallRecentSpeakerTemp
}

// UnmarshalJSON unmarshal to json
func (groupCallRecentSpeaker *GroupCallRecentSpeaker) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	tempObj := struct {
		tdCommon
		IsSpeaking bool `json:"is_speaking"` // True, is the user has spoken recently
	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	groupCallRecentSpeaker.tdCommon = tempObj.tdCommon
	groupCallRecentSpeaker.IsSpeaking = tempObj.IsSpeaking

	fieldParticipantId, _ := unmarshalMessageSender(objMap["participant_id"])
	groupCallRecentSpeaker.ParticipantId = fieldParticipantId

	return nil
}
