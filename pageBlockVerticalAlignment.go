package tdlib

import (
	"encoding/json"
	"fmt"
)

// PageBlockVerticalAlignment Describes a Vertical alignment of a table cell content
type PageBlockVerticalAlignment interface {
	GetPageBlockVerticalAlignmentEnum() PageBlockVerticalAlignmentEnum
}

// PageBlockVerticalAlignmentEnum Alias for abstract PageBlockVerticalAlignment 'Sub-Classes', used as constant-enum here
type PageBlockVerticalAlignmentEnum string

// PageBlockVerticalAlignment enums
const (
	PageBlockVerticalAlignmentTopType    PageBlockVerticalAlignmentEnum = "pageBlockVerticalAlignmentTop"
	PageBlockVerticalAlignmentMiddleType PageBlockVerticalAlignmentEnum = "pageBlockVerticalAlignmentMiddle"
	PageBlockVerticalAlignmentBottomType PageBlockVerticalAlignmentEnum = "pageBlockVerticalAlignmentBottom"
)

func unmarshalPageBlockVerticalAlignment(rawMsg *json.RawMessage) (PageBlockVerticalAlignment, error) {

	if rawMsg == nil {
		return nil, nil
	}
	var objMap map[string]interface{}
	err := json.Unmarshal(*rawMsg, &objMap)
	if err != nil {
		return nil, err
	}

	switch PageBlockVerticalAlignmentEnum(objMap["@type"].(string)) {
	case PageBlockVerticalAlignmentTopType:
		var pageBlockVerticalAlignmentTop PageBlockVerticalAlignmentTop
		err := json.Unmarshal(*rawMsg, &pageBlockVerticalAlignmentTop)
		return &pageBlockVerticalAlignmentTop, err

	case PageBlockVerticalAlignmentMiddleType:
		var pageBlockVerticalAlignmentMiddle PageBlockVerticalAlignmentMiddle
		err := json.Unmarshal(*rawMsg, &pageBlockVerticalAlignmentMiddle)
		return &pageBlockVerticalAlignmentMiddle, err

	case PageBlockVerticalAlignmentBottomType:
		var pageBlockVerticalAlignmentBottom PageBlockVerticalAlignmentBottom
		err := json.Unmarshal(*rawMsg, &pageBlockVerticalAlignmentBottom)
		return &pageBlockVerticalAlignmentBottom, err

	default:
		return nil, fmt.Errorf("Error UnMarshaling, unknown type:" + objMap["@type"].(string))
	}
}

// PageBlockVerticalAlignmentTop The content must be top-aligned
type PageBlockVerticalAlignmentTop struct {
	tdCommon
}

// MessageType return the string telegram-type of PageBlockVerticalAlignmentTop
func (pageBlockVerticalAlignmentTop *PageBlockVerticalAlignmentTop) MessageType() string {
	return "pageBlockVerticalAlignmentTop"
}

// NewPageBlockVerticalAlignmentTop creates a new PageBlockVerticalAlignmentTop
//
func NewPageBlockVerticalAlignmentTop() *PageBlockVerticalAlignmentTop {
	pageBlockVerticalAlignmentTopTemp := PageBlockVerticalAlignmentTop{
		tdCommon: tdCommon{Type: "pageBlockVerticalAlignmentTop"},
	}

	return &pageBlockVerticalAlignmentTopTemp
}

// GetPageBlockVerticalAlignmentEnum return the enum type of this object
func (pageBlockVerticalAlignmentTop *PageBlockVerticalAlignmentTop) GetPageBlockVerticalAlignmentEnum() PageBlockVerticalAlignmentEnum {
	return PageBlockVerticalAlignmentTopType
}

// PageBlockVerticalAlignmentMiddle The content must be middle-aligned
type PageBlockVerticalAlignmentMiddle struct {
	tdCommon
}

// MessageType return the string telegram-type of PageBlockVerticalAlignmentMiddle
func (pageBlockVerticalAlignmentMiddle *PageBlockVerticalAlignmentMiddle) MessageType() string {
	return "pageBlockVerticalAlignmentMiddle"
}

// NewPageBlockVerticalAlignmentMiddle creates a new PageBlockVerticalAlignmentMiddle
//
func NewPageBlockVerticalAlignmentMiddle() *PageBlockVerticalAlignmentMiddle {
	pageBlockVerticalAlignmentMiddleTemp := PageBlockVerticalAlignmentMiddle{
		tdCommon: tdCommon{Type: "pageBlockVerticalAlignmentMiddle"},
	}

	return &pageBlockVerticalAlignmentMiddleTemp
}

// GetPageBlockVerticalAlignmentEnum return the enum type of this object
func (pageBlockVerticalAlignmentMiddle *PageBlockVerticalAlignmentMiddle) GetPageBlockVerticalAlignmentEnum() PageBlockVerticalAlignmentEnum {
	return PageBlockVerticalAlignmentMiddleType
}

// PageBlockVerticalAlignmentBottom The content must be bottom-aligned
type PageBlockVerticalAlignmentBottom struct {
	tdCommon
}

// MessageType return the string telegram-type of PageBlockVerticalAlignmentBottom
func (pageBlockVerticalAlignmentBottom *PageBlockVerticalAlignmentBottom) MessageType() string {
	return "pageBlockVerticalAlignmentBottom"
}

// NewPageBlockVerticalAlignmentBottom creates a new PageBlockVerticalAlignmentBottom
//
func NewPageBlockVerticalAlignmentBottom() *PageBlockVerticalAlignmentBottom {
	pageBlockVerticalAlignmentBottomTemp := PageBlockVerticalAlignmentBottom{
		tdCommon: tdCommon{Type: "pageBlockVerticalAlignmentBottom"},
	}

	return &pageBlockVerticalAlignmentBottomTemp
}

// GetPageBlockVerticalAlignmentEnum return the enum type of this object
func (pageBlockVerticalAlignmentBottom *PageBlockVerticalAlignmentBottom) GetPageBlockVerticalAlignmentEnum() PageBlockVerticalAlignmentEnum {
	return PageBlockVerticalAlignmentBottomType
}
