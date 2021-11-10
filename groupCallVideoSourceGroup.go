package tdlib

// GroupCallVideoSourceGroup Describes a group of video synchronization source identifiers
type GroupCallVideoSourceGroup struct {
	tdCommon
	Semantics string  `json:"semantics"`  // The semantics of sources, one of "SIM" or "FID"
	SourceIds []int32 `json:"source_ids"` // The list of synchronization source identifiers
}

// MessageType return the string telegram-type of GroupCallVideoSourceGroup
func (groupCallVideoSourceGroup *GroupCallVideoSourceGroup) MessageType() string {
	return "groupCallVideoSourceGroup"
}

// NewGroupCallVideoSourceGroup creates a new GroupCallVideoSourceGroup
//
// @param semantics The semantics of sources, one of "SIM" or "FID"
// @param sourceIds The list of synchronization source identifiers
func NewGroupCallVideoSourceGroup(semantics string, sourceIds []int32) *GroupCallVideoSourceGroup {
	groupCallVideoSourceGroupTemp := GroupCallVideoSourceGroup{
		tdCommon:  tdCommon{Type: "groupCallVideoSourceGroup"},
		Semantics: semantics,
		SourceIds: sourceIds,
	}

	return &groupCallVideoSourceGroupTemp
}
