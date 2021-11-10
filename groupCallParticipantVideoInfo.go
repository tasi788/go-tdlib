package tdlib

// GroupCallParticipantVideoInfo Contains information about a group call participant's video channel
type GroupCallParticipantVideoInfo struct {
	tdCommon
	SourceGroups []GroupCallVideoSourceGroup `json:"source_groups"` // List of synchronization source groups of the video
	EndpointId   string                      `json:"endpoint_id"`   // Video channel endpoint identifier
	IsPaused     bool                        `json:"is_paused"`     // True if the video is paused. This flag needs to be ignored, if new video frames are received
}

// MessageType return the string telegram-type of GroupCallParticipantVideoInfo
func (groupCallParticipantVideoInfo *GroupCallParticipantVideoInfo) MessageType() string {
	return "groupCallParticipantVideoInfo"
}

// NewGroupCallParticipantVideoInfo creates a new GroupCallParticipantVideoInfo
//
// @param sourceGroups List of synchronization source groups of the video
// @param endpointId Video channel endpoint identifier
// @param isPaused True if the video is paused. This flag needs to be ignored, if new video frames are received
func NewGroupCallParticipantVideoInfo(sourceGroups []GroupCallVideoSourceGroup, endpointId string, isPaused bool) *GroupCallParticipantVideoInfo {
	groupCallParticipantVideoInfoTemp := GroupCallParticipantVideoInfo{
		tdCommon:     tdCommon{Type: "groupCallParticipantVideoInfo"},
		SourceGroups: sourceGroups,
		EndpointId:   endpointId,
		IsPaused:     isPaused,
	}

	return &groupCallParticipantVideoInfoTemp
}
