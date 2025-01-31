package tdlib

import (
	"encoding/json"
)

// PageBlockTableCell Represents a cell of a table
type PageBlockTableCell struct {
	tdCommon
	Text     RichText                     `json:"text"`      // Cell text; may be null. If the text is null, then the cell must be invisible
	IsHeader bool                         `json:"is_header"` // True, if it is a header cell
	Colspan  int32                        `json:"colspan"`   // The number of columns the cell spans
	Rowspan  int32                        `json:"rowspan"`   // The number of rows the cell spans
	Align    PageBlockHorizontalAlignment `json:"align"`     // Horizontal cell content alignment
	Valign   PageBlockVerticalAlignment   `json:"valign"`    // Vertical cell content alignment
}

// MessageType return the string telegram-type of PageBlockTableCell
func (pageBlockTableCell *PageBlockTableCell) MessageType() string {
	return "pageBlockTableCell"
}

// NewPageBlockTableCell creates a new PageBlockTableCell
//
// @param text Cell text; may be null. If the text is null, then the cell must be invisible
// @param isHeader True, if it is a header cell
// @param colspan The number of columns the cell spans
// @param rowspan The number of rows the cell spans
// @param align Horizontal cell content alignment
// @param valign Vertical cell content alignment
func NewPageBlockTableCell(text RichText, isHeader bool, colspan int32, rowspan int32, align PageBlockHorizontalAlignment, valign PageBlockVerticalAlignment) *PageBlockTableCell {
	pageBlockTableCellTemp := PageBlockTableCell{
		tdCommon: tdCommon{Type: "pageBlockTableCell"},
		Text:     text,
		IsHeader: isHeader,
		Colspan:  colspan,
		Rowspan:  rowspan,
		Align:    align,
		Valign:   valign,
	}

	return &pageBlockTableCellTemp
}

// UnmarshalJSON unmarshal to json
func (pageBlockTableCell *PageBlockTableCell) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	tempObj := struct {
		tdCommon
		IsHeader bool  `json:"is_header"` // True, if it is a header cell
		Colspan  int32 `json:"colspan"`   // The number of columns the cell spans
		Rowspan  int32 `json:"rowspan"`   // The number of rows the cell spans

	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	pageBlockTableCell.tdCommon = tempObj.tdCommon
	pageBlockTableCell.IsHeader = tempObj.IsHeader
	pageBlockTableCell.Colspan = tempObj.Colspan
	pageBlockTableCell.Rowspan = tempObj.Rowspan

	fieldText, _ := unmarshalRichText(objMap["text"])
	pageBlockTableCell.Text = fieldText

	fieldAlign, _ := unmarshalPageBlockHorizontalAlignment(objMap["align"])
	pageBlockTableCell.Align = fieldAlign

	fieldValign, _ := unmarshalPageBlockVerticalAlignment(objMap["valign"])
	pageBlockTableCell.Valign = fieldValign

	return nil
}
