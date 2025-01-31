package tdlib

import (
	"encoding/json"
	"fmt"
)

// GroupCall Describes a group call
type GroupCall struct {
	tdCommon
	Id                           int32                    `json:"id"`                               // Group call identifier
	Title                        string                   `json:"title"`                            // Group call title
	ScheduledStartDate           int32                    `json:"scheduled_start_date"`             // Point in time (Unix timestamp) when the group call is supposed to be started by an administrator; 0 if it is already active or was ended
	EnabledStartNotification     bool                     `json:"enabled_start_notification"`       // True, if the group call is scheduled and the current user will receive a notification when the group call will start
	IsActive                     bool                     `json:"is_active"`                        // True, if the call is active
	IsJoined                     bool                     `json:"is_joined"`                        // True, if the call is joined
	NeedRejoin                   bool                     `json:"need_rejoin"`                      // True, if user was kicked from the call because of network loss and the call needs to be rejoined
	CanBeManaged                 bool                     `json:"can_be_managed"`                   // True, if the current user can manage the group call
	ParticipantCount             int32                    `json:"participant_count"`                // Number of participants in the group call
	LoadedAllParticipants        bool                     `json:"loaded_all_participants"`          // True, if all group call participants are loaded
	RecentSpeakers               []GroupCallRecentSpeaker `json:"recent_speakers"`                  // At most 3 recently speaking users in the group call
	IsMyVideoEnabled             bool                     `json:"is_my_video_enabled"`              // True, if the current user's video is enabled
	IsMyVideoPaused              bool                     `json:"is_my_video_paused"`               // True, if the current user's video is paused
	CanEnableVideo               bool                     `json:"can_enable_video"`                 // True, if the current user can broadcast video or share screen
	MuteNewParticipants          bool                     `json:"mute_new_participants"`            // True, if only group call administrators can unmute new participants
	CanToggleMuteNewParticipants bool                     `json:"can_toggle_mute_new_participants"` // True, if the current user can enable or disable mute_new_participants setting
	RecordDuration               int32                    `json:"record_duration"`                  // Duration of the ongoing group call recording, in seconds; 0 if none. An updateGroupCall update is not triggered when value of this field changes, but the same recording goes on
	IsVideoRecorded              bool                     `json:"is_video_recorded"`                // True, if a video file is being recorded for the call
	Duration                     int32                    `json:"duration"`                         // Call duration, in seconds; for ended calls only
}

// MessageType return the string telegram-type of GroupCall
func (groupCall *GroupCall) MessageType() string {
	return "groupCall"
}

// NewGroupCall creates a new GroupCall
//
// @param id Group call identifier
// @param title Group call title
// @param scheduledStartDate Point in time (Unix timestamp) when the group call is supposed to be started by an administrator; 0 if it is already active or was ended
// @param enabledStartNotification True, if the group call is scheduled and the current user will receive a notification when the group call will start
// @param isActive True, if the call is active
// @param isJoined True, if the call is joined
// @param needRejoin True, if user was kicked from the call because of network loss and the call needs to be rejoined
// @param canBeManaged True, if the current user can manage the group call
// @param participantCount Number of participants in the group call
// @param loadedAllParticipants True, if all group call participants are loaded
// @param recentSpeakers At most 3 recently speaking users in the group call
// @param isMyVideoEnabled True, if the current user's video is enabled
// @param isMyVideoPaused True, if the current user's video is paused
// @param canEnableVideo True, if the current user can broadcast video or share screen
// @param muteNewParticipants True, if only group call administrators can unmute new participants
// @param canToggleMuteNewParticipants True, if the current user can enable or disable mute_new_participants setting
// @param recordDuration Duration of the ongoing group call recording, in seconds; 0 if none. An updateGroupCall update is not triggered when value of this field changes, but the same recording goes on
// @param isVideoRecorded True, if a video file is being recorded for the call
// @param duration Call duration, in seconds; for ended calls only
func NewGroupCall(id int32, title string, scheduledStartDate int32, enabledStartNotification bool, isActive bool, isJoined bool, needRejoin bool, canBeManaged bool, participantCount int32, loadedAllParticipants bool, recentSpeakers []GroupCallRecentSpeaker, isMyVideoEnabled bool, isMyVideoPaused bool, canEnableVideo bool, muteNewParticipants bool, canToggleMuteNewParticipants bool, recordDuration int32, isVideoRecorded bool, duration int32) *GroupCall {
	groupCallTemp := GroupCall{
		tdCommon:                     tdCommon{Type: "groupCall"},
		Id:                           id,
		Title:                        title,
		ScheduledStartDate:           scheduledStartDate,
		EnabledStartNotification:     enabledStartNotification,
		IsActive:                     isActive,
		IsJoined:                     isJoined,
		NeedRejoin:                   needRejoin,
		CanBeManaged:                 canBeManaged,
		ParticipantCount:             participantCount,
		LoadedAllParticipants:        loadedAllParticipants,
		RecentSpeakers:               recentSpeakers,
		IsMyVideoEnabled:             isMyVideoEnabled,
		IsMyVideoPaused:              isMyVideoPaused,
		CanEnableVideo:               canEnableVideo,
		MuteNewParticipants:          muteNewParticipants,
		CanToggleMuteNewParticipants: canToggleMuteNewParticipants,
		RecordDuration:               recordDuration,
		IsVideoRecorded:              isVideoRecorded,
		Duration:                     duration,
	}

	return &groupCallTemp
}

// GetGroupCall Returns information about a group call
// @param groupCallId Group call identifier
func (client *Client) GetGroupCall(groupCallId int32) (*GroupCall, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":         "getGroupCall",
		"group_call_id": groupCallId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var groupCallDummy GroupCall
	err = json.Unmarshal(result.Raw, &groupCallDummy)
	return &groupCallDummy, err
}
