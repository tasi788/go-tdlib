package tdlib

import (
	"encoding/json"
)

// GroupCallParticipant Represents a group call participant
type GroupCallParticipant struct {
	tdCommon
	ParticipantId              MessageSender                  `json:"participant_id"`                  // Identifier of the group call participant
	AudioSourceId              int32                          `json:"audio_source_id"`                 // User's audio channel synchronization source identifier
	ScreenSharingAudioSourceId int32                          `json:"screen_sharing_audio_source_id"`  // User's screen sharing audio channel synchronization source identifier
	VideoInfo                  *GroupCallParticipantVideoInfo `json:"video_info"`                      // Information about user's video channel; may be null if there is no active video
	ScreenSharingVideoInfo     *GroupCallParticipantVideoInfo `json:"screen_sharing_video_info"`       // Information about user's screen sharing video channel; may be null if there is no active screen sharing video
	Bio                        string                         `json:"bio"`                             // The participant user's bio or the participant chat's description
	IsCurrentUser              bool                           `json:"is_current_user"`                 // True, if the participant is the current user
	IsSpeaking                 bool                           `json:"is_speaking"`                     // True, if the participant is speaking as set by setGroupCallParticipantIsSpeaking
	IsHandRaised               bool                           `json:"is_hand_raised"`                  // True, if the participant hand is raised
	CanBeMutedForAllUsers      bool                           `json:"can_be_muted_for_all_users"`      // True, if the current user can mute the participant for all other group call participants
	CanBeUnmutedForAllUsers    bool                           `json:"can_be_unmuted_for_all_users"`    // True, if the current user can allow the participant to unmute themselves or unmute the participant (if the participant is the current user)
	CanBeMutedForCurrentUser   bool                           `json:"can_be_muted_for_current_user"`   // True, if the current user can mute the participant only for self
	CanBeUnmutedForCurrentUser bool                           `json:"can_be_unmuted_for_current_user"` // True, if the current user can unmute the participant for self
	IsMutedForAllUsers         bool                           `json:"is_muted_for_all_users"`          // True, if the participant is muted for all users
	IsMutedForCurrentUser      bool                           `json:"is_muted_for_current_user"`       // True, if the participant is muted for the current user
	CanUnmuteSelf              bool                           `json:"can_unmute_self"`                 // True, if the participant is muted for all users, but can unmute themselves
	VolumeLevel                int32                          `json:"volume_level"`                    // Participant's volume level; 1-20000 in hundreds of percents
	Order                      string                         `json:"order"`                           // User's order in the group call participant list. Orders must be compared lexicographically. The bigger is order, the higher is user in the list. If order is empty, the user must be removed from the participant list
}

// MessageType return the string telegram-type of GroupCallParticipant
func (groupCallParticipant *GroupCallParticipant) MessageType() string {
	return "groupCallParticipant"
}

// NewGroupCallParticipant creates a new GroupCallParticipant
//
// @param participantId Identifier of the group call participant
// @param audioSourceId User's audio channel synchronization source identifier
// @param screenSharingAudioSourceId User's screen sharing audio channel synchronization source identifier
// @param videoInfo Information about user's video channel; may be null if there is no active video
// @param screenSharingVideoInfo Information about user's screen sharing video channel; may be null if there is no active screen sharing video
// @param bio The participant user's bio or the participant chat's description
// @param isCurrentUser True, if the participant is the current user
// @param isSpeaking True, if the participant is speaking as set by setGroupCallParticipantIsSpeaking
// @param isHandRaised True, if the participant hand is raised
// @param canBeMutedForAllUsers True, if the current user can mute the participant for all other group call participants
// @param canBeUnmutedForAllUsers True, if the current user can allow the participant to unmute themselves or unmute the participant (if the participant is the current user)
// @param canBeMutedForCurrentUser True, if the current user can mute the participant only for self
// @param canBeUnmutedForCurrentUser True, if the current user can unmute the participant for self
// @param isMutedForAllUsers True, if the participant is muted for all users
// @param isMutedForCurrentUser True, if the participant is muted for the current user
// @param canUnmuteSelf True, if the participant is muted for all users, but can unmute themselves
// @param volumeLevel Participant's volume level; 1-20000 in hundreds of percents
// @param order User's order in the group call participant list. Orders must be compared lexicographically. The bigger is order, the higher is user in the list. If order is empty, the user must be removed from the participant list
func NewGroupCallParticipant(participantId MessageSender, audioSourceId int32, screenSharingAudioSourceId int32, videoInfo *GroupCallParticipantVideoInfo, screenSharingVideoInfo *GroupCallParticipantVideoInfo, bio string, isCurrentUser bool, isSpeaking bool, isHandRaised bool, canBeMutedForAllUsers bool, canBeUnmutedForAllUsers bool, canBeMutedForCurrentUser bool, canBeUnmutedForCurrentUser bool, isMutedForAllUsers bool, isMutedForCurrentUser bool, canUnmuteSelf bool, volumeLevel int32, order string) *GroupCallParticipant {
	groupCallParticipantTemp := GroupCallParticipant{
		tdCommon:                   tdCommon{Type: "groupCallParticipant"},
		ParticipantId:              participantId,
		AudioSourceId:              audioSourceId,
		ScreenSharingAudioSourceId: screenSharingAudioSourceId,
		VideoInfo:                  videoInfo,
		ScreenSharingVideoInfo:     screenSharingVideoInfo,
		Bio:                        bio,
		IsCurrentUser:              isCurrentUser,
		IsSpeaking:                 isSpeaking,
		IsHandRaised:               isHandRaised,
		CanBeMutedForAllUsers:      canBeMutedForAllUsers,
		CanBeUnmutedForAllUsers:    canBeUnmutedForAllUsers,
		CanBeMutedForCurrentUser:   canBeMutedForCurrentUser,
		CanBeUnmutedForCurrentUser: canBeUnmutedForCurrentUser,
		IsMutedForAllUsers:         isMutedForAllUsers,
		IsMutedForCurrentUser:      isMutedForCurrentUser,
		CanUnmuteSelf:              canUnmuteSelf,
		VolumeLevel:                volumeLevel,
		Order:                      order,
	}

	return &groupCallParticipantTemp
}

// UnmarshalJSON unmarshal to json
func (groupCallParticipant *GroupCallParticipant) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	tempObj := struct {
		tdCommon
		AudioSourceId              int32                          `json:"audio_source_id"`                 // User's audio channel synchronization source identifier
		ScreenSharingAudioSourceId int32                          `json:"screen_sharing_audio_source_id"`  // User's screen sharing audio channel synchronization source identifier
		VideoInfo                  *GroupCallParticipantVideoInfo `json:"video_info"`                      // Information about user's video channel; may be null if there is no active video
		ScreenSharingVideoInfo     *GroupCallParticipantVideoInfo `json:"screen_sharing_video_info"`       // Information about user's screen sharing video channel; may be null if there is no active screen sharing video
		Bio                        string                         `json:"bio"`                             // The participant user's bio or the participant chat's description
		IsCurrentUser              bool                           `json:"is_current_user"`                 // True, if the participant is the current user
		IsSpeaking                 bool                           `json:"is_speaking"`                     // True, if the participant is speaking as set by setGroupCallParticipantIsSpeaking
		IsHandRaised               bool                           `json:"is_hand_raised"`                  // True, if the participant hand is raised
		CanBeMutedForAllUsers      bool                           `json:"can_be_muted_for_all_users"`      // True, if the current user can mute the participant for all other group call participants
		CanBeUnmutedForAllUsers    bool                           `json:"can_be_unmuted_for_all_users"`    // True, if the current user can allow the participant to unmute themselves or unmute the participant (if the participant is the current user)
		CanBeMutedForCurrentUser   bool                           `json:"can_be_muted_for_current_user"`   // True, if the current user can mute the participant only for self
		CanBeUnmutedForCurrentUser bool                           `json:"can_be_unmuted_for_current_user"` // True, if the current user can unmute the participant for self
		IsMutedForAllUsers         bool                           `json:"is_muted_for_all_users"`          // True, if the participant is muted for all users
		IsMutedForCurrentUser      bool                           `json:"is_muted_for_current_user"`       // True, if the participant is muted for the current user
		CanUnmuteSelf              bool                           `json:"can_unmute_self"`                 // True, if the participant is muted for all users, but can unmute themselves
		VolumeLevel                int32                          `json:"volume_level"`                    // Participant's volume level; 1-20000 in hundreds of percents
		Order                      string                         `json:"order"`                           // User's order in the group call participant list. Orders must be compared lexicographically. The bigger is order, the higher is user in the list. If order is empty, the user must be removed from the participant list
	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	groupCallParticipant.tdCommon = tempObj.tdCommon
	groupCallParticipant.AudioSourceId = tempObj.AudioSourceId
	groupCallParticipant.ScreenSharingAudioSourceId = tempObj.ScreenSharingAudioSourceId
	groupCallParticipant.VideoInfo = tempObj.VideoInfo
	groupCallParticipant.ScreenSharingVideoInfo = tempObj.ScreenSharingVideoInfo
	groupCallParticipant.Bio = tempObj.Bio
	groupCallParticipant.IsCurrentUser = tempObj.IsCurrentUser
	groupCallParticipant.IsSpeaking = tempObj.IsSpeaking
	groupCallParticipant.IsHandRaised = tempObj.IsHandRaised
	groupCallParticipant.CanBeMutedForAllUsers = tempObj.CanBeMutedForAllUsers
	groupCallParticipant.CanBeUnmutedForAllUsers = tempObj.CanBeUnmutedForAllUsers
	groupCallParticipant.CanBeMutedForCurrentUser = tempObj.CanBeMutedForCurrentUser
	groupCallParticipant.CanBeUnmutedForCurrentUser = tempObj.CanBeUnmutedForCurrentUser
	groupCallParticipant.IsMutedForAllUsers = tempObj.IsMutedForAllUsers
	groupCallParticipant.IsMutedForCurrentUser = tempObj.IsMutedForCurrentUser
	groupCallParticipant.CanUnmuteSelf = tempObj.CanUnmuteSelf
	groupCallParticipant.VolumeLevel = tempObj.VolumeLevel
	groupCallParticipant.Order = tempObj.Order

	fieldParticipantId, _ := unmarshalMessageSender(objMap["participant_id"])
	groupCallParticipant.ParticipantId = fieldParticipantId

	return nil
}
