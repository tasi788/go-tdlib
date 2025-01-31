package tdlib

// PageBlockListItem Describes an item of a list page block
type PageBlockListItem struct {
	tdCommon
	Label      string      `json:"label"`       // Item label
	PageBlocks []PageBlock `json:"page_blocks"` // Item blocks
}

// MessageType return the string telegram-type of PageBlockListItem
func (pageBlockListItem *PageBlockListItem) MessageType() string {
	return "pageBlockListItem"
}

// NewPageBlockListItem creates a new PageBlockListItem
//
// @param label Item label
// @param pageBlocks Item blocks
func NewPageBlockListItem(label string, pageBlocks []PageBlock) *PageBlockListItem {
	pageBlockListItemTemp := PageBlockListItem{
		tdCommon:   tdCommon{Type: "pageBlockListItem"},
		Label:      label,
		PageBlocks: pageBlocks,
	}

	return &pageBlockListItemTemp
}
