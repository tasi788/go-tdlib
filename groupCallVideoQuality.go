package tdlib

import (
	"encoding/json"
	"fmt"
)

// GroupCallVideoQuality Describes the quality of a group call video
type GroupCallVideoQuality interface {
	GetGroupCallVideoQualityEnum() GroupCallVideoQualityEnum
}

// GroupCallVideoQualityEnum Alias for abstract GroupCallVideoQuality 'Sub-Classes', used as constant-enum here
type GroupCallVideoQualityEnum string

// GroupCallVideoQuality enums
const (
	GroupCallVideoQualityThumbnailType GroupCallVideoQualityEnum = "groupCallVideoQualityThumbnail"
	GroupCallVideoQualityMediumType    GroupCallVideoQualityEnum = "groupCallVideoQualityMedium"
	GroupCallVideoQualityFullType      GroupCallVideoQualityEnum = "groupCallVideoQualityFull"
)

func unmarshalGroupCallVideoQuality(rawMsg *json.RawMessage) (GroupCallVideoQuality, error) {

	if rawMsg == nil {
		return nil, nil
	}
	var objMap map[string]interface{}
	err := json.Unmarshal(*rawMsg, &objMap)
	if err != nil {
		return nil, err
	}

	switch GroupCallVideoQualityEnum(objMap["@type"].(string)) {
	case GroupCallVideoQualityThumbnailType:
		var groupCallVideoQualityThumbnail GroupCallVideoQualityThumbnail
		err := json.Unmarshal(*rawMsg, &groupCallVideoQualityThumbnail)
		return &groupCallVideoQualityThumbnail, err

	case GroupCallVideoQualityMediumType:
		var groupCallVideoQualityMedium GroupCallVideoQualityMedium
		err := json.Unmarshal(*rawMsg, &groupCallVideoQualityMedium)
		return &groupCallVideoQualityMedium, err

	case GroupCallVideoQualityFullType:
		var groupCallVideoQualityFull GroupCallVideoQualityFull
		err := json.Unmarshal(*rawMsg, &groupCallVideoQualityFull)
		return &groupCallVideoQualityFull, err

	default:
		return nil, fmt.Errorf("Error UnMarshaling, unknown type:" + objMap["@type"].(string))
	}
}

// GroupCallVideoQualityThumbnail The worst available video quality
type GroupCallVideoQualityThumbnail struct {
	tdCommon
}

// MessageType return the string telegram-type of GroupCallVideoQualityThumbnail
func (groupCallVideoQualityThumbnail *GroupCallVideoQualityThumbnail) MessageType() string {
	return "groupCallVideoQualityThumbnail"
}

// NewGroupCallVideoQualityThumbnail creates a new GroupCallVideoQualityThumbnail
//
func NewGroupCallVideoQualityThumbnail() *GroupCallVideoQualityThumbnail {
	groupCallVideoQualityThumbnailTemp := GroupCallVideoQualityThumbnail{
		tdCommon: tdCommon{Type: "groupCallVideoQualityThumbnail"},
	}

	return &groupCallVideoQualityThumbnailTemp
}

// GetGroupCallVideoQualityEnum return the enum type of this object
func (groupCallVideoQualityThumbnail *GroupCallVideoQualityThumbnail) GetGroupCallVideoQualityEnum() GroupCallVideoQualityEnum {
	return GroupCallVideoQualityThumbnailType
}

// GroupCallVideoQualityMedium The medium video quality
type GroupCallVideoQualityMedium struct {
	tdCommon
}

// MessageType return the string telegram-type of GroupCallVideoQualityMedium
func (groupCallVideoQualityMedium *GroupCallVideoQualityMedium) MessageType() string {
	return "groupCallVideoQualityMedium"
}

// NewGroupCallVideoQualityMedium creates a new GroupCallVideoQualityMedium
//
func NewGroupCallVideoQualityMedium() *GroupCallVideoQualityMedium {
	groupCallVideoQualityMediumTemp := GroupCallVideoQualityMedium{
		tdCommon: tdCommon{Type: "groupCallVideoQualityMedium"},
	}

	return &groupCallVideoQualityMediumTemp
}

// GetGroupCallVideoQualityEnum return the enum type of this object
func (groupCallVideoQualityMedium *GroupCallVideoQualityMedium) GetGroupCallVideoQualityEnum() GroupCallVideoQualityEnum {
	return GroupCallVideoQualityMediumType
}

// GroupCallVideoQualityFull The best available video quality
type GroupCallVideoQualityFull struct {
	tdCommon
}

// MessageType return the string telegram-type of GroupCallVideoQualityFull
func (groupCallVideoQualityFull *GroupCallVideoQualityFull) MessageType() string {
	return "groupCallVideoQualityFull"
}

// NewGroupCallVideoQualityFull creates a new GroupCallVideoQualityFull
//
func NewGroupCallVideoQualityFull() *GroupCallVideoQualityFull {
	groupCallVideoQualityFullTemp := GroupCallVideoQualityFull{
		tdCommon: tdCommon{Type: "groupCallVideoQualityFull"},
	}

	return &groupCallVideoQualityFullTemp
}

// GetGroupCallVideoQualityEnum return the enum type of this object
func (groupCallVideoQualityFull *GroupCallVideoQualityFull) GetGroupCallVideoQualityEnum() GroupCallVideoQualityEnum {
	return GroupCallVideoQualityFullType
}
