package tdlib

import (
	"encoding/json"
	"fmt"
)

// PageBlockHorizontalAlignment Describes a horizontal alignment of a table cell content
type PageBlockHorizontalAlignment interface {
	GetPageBlockHorizontalAlignmentEnum() PageBlockHorizontalAlignmentEnum
}

// PageBlockHorizontalAlignmentEnum Alias for abstract PageBlockHorizontalAlignment 'Sub-Classes', used as constant-enum here
type PageBlockHorizontalAlignmentEnum string

// PageBlockHorizontalAlignment enums
const (
	PageBlockHorizontalAlignmentLeftType   PageBlockHorizontalAlignmentEnum = "pageBlockHorizontalAlignmentLeft"
	PageBlockHorizontalAlignmentCenterType PageBlockHorizontalAlignmentEnum = "pageBlockHorizontalAlignmentCenter"
	PageBlockHorizontalAlignmentRightType  PageBlockHorizontalAlignmentEnum = "pageBlockHorizontalAlignmentRight"
)

func unmarshalPageBlockHorizontalAlignment(rawMsg *json.RawMessage) (PageBlockHorizontalAlignment, error) {

	if rawMsg == nil {
		return nil, nil
	}
	var objMap map[string]interface{}
	err := json.Unmarshal(*rawMsg, &objMap)
	if err != nil {
		return nil, err
	}

	switch PageBlockHorizontalAlignmentEnum(objMap["@type"].(string)) {
	case PageBlockHorizontalAlignmentLeftType:
		var pageBlockHorizontalAlignmentLeft PageBlockHorizontalAlignmentLeft
		err := json.Unmarshal(*rawMsg, &pageBlockHorizontalAlignmentLeft)
		return &pageBlockHorizontalAlignmentLeft, err

	case PageBlockHorizontalAlignmentCenterType:
		var pageBlockHorizontalAlignmentCenter PageBlockHorizontalAlignmentCenter
		err := json.Unmarshal(*rawMsg, &pageBlockHorizontalAlignmentCenter)
		return &pageBlockHorizontalAlignmentCenter, err

	case PageBlockHorizontalAlignmentRightType:
		var pageBlockHorizontalAlignmentRight PageBlockHorizontalAlignmentRight
		err := json.Unmarshal(*rawMsg, &pageBlockHorizontalAlignmentRight)
		return &pageBlockHorizontalAlignmentRight, err

	default:
		return nil, fmt.Errorf("Error UnMarshaling, unknown type:" + objMap["@type"].(string))
	}
}

// PageBlockHorizontalAlignmentLeft The content must be left-aligned
type PageBlockHorizontalAlignmentLeft struct {
	tdCommon
}

// MessageType return the string telegram-type of PageBlockHorizontalAlignmentLeft
func (pageBlockHorizontalAlignmentLeft *PageBlockHorizontalAlignmentLeft) MessageType() string {
	return "pageBlockHorizontalAlignmentLeft"
}

// NewPageBlockHorizontalAlignmentLeft creates a new PageBlockHorizontalAlignmentLeft
//
func NewPageBlockHorizontalAlignmentLeft() *PageBlockHorizontalAlignmentLeft {
	pageBlockHorizontalAlignmentLeftTemp := PageBlockHorizontalAlignmentLeft{
		tdCommon: tdCommon{Type: "pageBlockHorizontalAlignmentLeft"},
	}

	return &pageBlockHorizontalAlignmentLeftTemp
}

// GetPageBlockHorizontalAlignmentEnum return the enum type of this object
func (pageBlockHorizontalAlignmentLeft *PageBlockHorizontalAlignmentLeft) GetPageBlockHorizontalAlignmentEnum() PageBlockHorizontalAlignmentEnum {
	return PageBlockHorizontalAlignmentLeftType
}

// PageBlockHorizontalAlignmentCenter The content must be center-aligned
type PageBlockHorizontalAlignmentCenter struct {
	tdCommon
}

// MessageType return the string telegram-type of PageBlockHorizontalAlignmentCenter
func (pageBlockHorizontalAlignmentCenter *PageBlockHorizontalAlignmentCenter) MessageType() string {
	return "pageBlockHorizontalAlignmentCenter"
}

// NewPageBlockHorizontalAlignmentCenter creates a new PageBlockHorizontalAlignmentCenter
//
func NewPageBlockHorizontalAlignmentCenter() *PageBlockHorizontalAlignmentCenter {
	pageBlockHorizontalAlignmentCenterTemp := PageBlockHorizontalAlignmentCenter{
		tdCommon: tdCommon{Type: "pageBlockHorizontalAlignmentCenter"},
	}

	return &pageBlockHorizontalAlignmentCenterTemp
}

// GetPageBlockHorizontalAlignmentEnum return the enum type of this object
func (pageBlockHorizontalAlignmentCenter *PageBlockHorizontalAlignmentCenter) GetPageBlockHorizontalAlignmentEnum() PageBlockHorizontalAlignmentEnum {
	return PageBlockHorizontalAlignmentCenterType
}

// PageBlockHorizontalAlignmentRight The content must be right-aligned
type PageBlockHorizontalAlignmentRight struct {
	tdCommon
}

// MessageType return the string telegram-type of PageBlockHorizontalAlignmentRight
func (pageBlockHorizontalAlignmentRight *PageBlockHorizontalAlignmentRight) MessageType() string {
	return "pageBlockHorizontalAlignmentRight"
}

// NewPageBlockHorizontalAlignmentRight creates a new PageBlockHorizontalAlignmentRight
//
func NewPageBlockHorizontalAlignmentRight() *PageBlockHorizontalAlignmentRight {
	pageBlockHorizontalAlignmentRightTemp := PageBlockHorizontalAlignmentRight{
		tdCommon: tdCommon{Type: "pageBlockHorizontalAlignmentRight"},
	}

	return &pageBlockHorizontalAlignmentRightTemp
}

// GetPageBlockHorizontalAlignmentEnum return the enum type of this object
func (pageBlockHorizontalAlignmentRight *PageBlockHorizontalAlignmentRight) GetPageBlockHorizontalAlignmentEnum() PageBlockHorizontalAlignmentEnum {
	return PageBlockHorizontalAlignmentRightType
}
