package tdlib

import (
	"encoding/json"
	"fmt"
)

// FormattedText A text with some entities
type FormattedText struct {
	tdCommon
	Text     string       `json:"text"`     // The text
	Entities []TextEntity `json:"entities"` // Entities contained in the text. Entities can be nested, but must not mutually intersect with each other. Pre, Code and PreCode entities can't contain other entities. Bold, Italic, Underline, Strikethrough, and Spoiler entities can contain and to be contained in all other entities. All other entities can't contain each other
}

// MessageType return the string telegram-type of FormattedText
func (formattedText *FormattedText) MessageType() string {
	return "formattedText"
}

// NewFormattedText creates a new FormattedText
//
// @param text The text
// @param entities Entities contained in the text. Entities can be nested, but must not mutually intersect with each other. Pre, Code and PreCode entities can't contain other entities. Bold, Italic, Underline, Strikethrough, and Spoiler entities can contain and to be contained in all other entities. All other entities can't contain each other
func NewFormattedText(text string, entities []TextEntity) *FormattedText {
	formattedTextTemp := FormattedText{
		tdCommon: tdCommon{Type: "formattedText"},
		Text:     text,
		Entities: entities,
	}

	return &formattedTextTemp
}

// ParseTextEntities Parses Bold, Italic, Underline, Strikethrough, Spoiler, Code, Pre, PreCode, TextUrl and MentionName entities contained in the text. Can be called synchronously
// @param text The text to parse
// @param parseMode Text parse mode
func (client *Client) ParseTextEntities(text string, parseMode TextParseMode) (*FormattedText, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":      "parseTextEntities",
		"text":       text,
		"parse_mode": parseMode,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var formattedText FormattedText
	err = json.Unmarshal(result.Raw, &formattedText)
	return &formattedText, err
}

// ParseMarkdown Parses Markdown entities in a human-friendly format, ignoring markup errors. Can be called synchronously
// @param text The text to parse. For example, "__italic__ ~~strikethrough~~ ||spoiler|| **bold** `code` ```pre``` __[italic__ text_url](telegram.org) __italic**bold italic__bold**"
func (client *Client) ParseMarkdown(text *FormattedText) (*FormattedText, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "parseMarkdown",
		"text":  text,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var formattedText FormattedText
	err = json.Unmarshal(result.Raw, &formattedText)
	return &formattedText, err
}

// GetMarkdownText Replaces text entities with Markdown formatting in a human-friendly format. Entities that can't be represented in Markdown unambiguously are kept as is. Can be called synchronously
// @param text The text
func (client *Client) GetMarkdownText(text *FormattedText) (*FormattedText, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "getMarkdownText",
		"text":  text,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var formattedText FormattedText
	err = json.Unmarshal(result.Raw, &formattedText)
	return &formattedText, err
}
