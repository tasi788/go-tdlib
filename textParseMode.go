package tdlib

import (
	"encoding/json"
	"fmt"
)

// TextParseMode Describes the way the text needs to be parsed for TextEntities
type TextParseMode interface {
	GetTextParseModeEnum() TextParseModeEnum
}

// TextParseModeEnum Alias for abstract TextParseMode 'Sub-Classes', used as constant-enum here
type TextParseModeEnum string

// TextParseMode enums
const (
	TextParseModeMarkdownType TextParseModeEnum = "textParseModeMarkdown"
	TextParseModeHTMLType     TextParseModeEnum = "textParseModeHTML"
)

func unmarshalTextParseMode(rawMsg *json.RawMessage) (TextParseMode, error) {

	if rawMsg == nil {
		return nil, nil
	}
	var objMap map[string]interface{}
	err := json.Unmarshal(*rawMsg, &objMap)
	if err != nil {
		return nil, err
	}

	switch TextParseModeEnum(objMap["@type"].(string)) {
	case TextParseModeMarkdownType:
		var textParseModeMarkdown TextParseModeMarkdown
		err := json.Unmarshal(*rawMsg, &textParseModeMarkdown)
		return &textParseModeMarkdown, err

	case TextParseModeHTMLType:
		var textParseModeHTML TextParseModeHTML
		err := json.Unmarshal(*rawMsg, &textParseModeHTML)
		return &textParseModeHTML, err

	default:
		return nil, fmt.Errorf("Error UnMarshaling, unknown type:" + objMap["@type"].(string))
	}
}

// TextParseModeMarkdown The text uses Markdown-style formatting
type TextParseModeMarkdown struct {
	tdCommon
	Version int32 `json:"version"` // Version of the parser: 0 or 1 - Telegram Bot API "Markdown" parse mode, 2 - Telegram Bot API "MarkdownV2" parse mode
}

// MessageType return the string telegram-type of TextParseModeMarkdown
func (textParseModeMarkdown *TextParseModeMarkdown) MessageType() string {
	return "textParseModeMarkdown"
}

// NewTextParseModeMarkdown creates a new TextParseModeMarkdown
//
// @param version Version of the parser: 0 or 1 - Telegram Bot API "Markdown" parse mode, 2 - Telegram Bot API "MarkdownV2" parse mode
func NewTextParseModeMarkdown(version int32) *TextParseModeMarkdown {
	textParseModeMarkdownTemp := TextParseModeMarkdown{
		tdCommon: tdCommon{Type: "textParseModeMarkdown"},
		Version:  version,
	}

	return &textParseModeMarkdownTemp
}

// GetTextParseModeEnum return the enum type of this object
func (textParseModeMarkdown *TextParseModeMarkdown) GetTextParseModeEnum() TextParseModeEnum {
	return TextParseModeMarkdownType
}

// TextParseModeHTML The text uses HTML-style formatting. The same as Telegram Bot API "HTML" parse mode
type TextParseModeHTML struct {
	tdCommon
}

// MessageType return the string telegram-type of TextParseModeHTML
func (textParseModeHTML *TextParseModeHTML) MessageType() string {
	return "textParseModeHTML"
}

// NewTextParseModeHTML creates a new TextParseModeHTML
//
func NewTextParseModeHTML() *TextParseModeHTML {
	textParseModeHTMLTemp := TextParseModeHTML{
		tdCommon: tdCommon{Type: "textParseModeHTML"},
	}

	return &textParseModeHTMLTemp
}

// GetTextParseModeEnum return the enum type of this object
func (textParseModeHTML *TextParseModeHTML) GetTextParseModeEnum() TextParseModeEnum {
	return TextParseModeHTMLType
}
