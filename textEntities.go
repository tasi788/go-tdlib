package tdlib

import (
	"encoding/json"
	"fmt"
)

// TextEntities Contains a list of text entities
type TextEntities struct {
	tdCommon
	Entities []TextEntity `json:"entities"` // List of text entities
}

// MessageType return the string telegram-type of TextEntities
func (textEntities *TextEntities) MessageType() string {
	return "textEntities"
}

// NewTextEntities creates a new TextEntities
//
// @param entities List of text entities
func NewTextEntities(entities []TextEntity) *TextEntities {
	textEntitiesTemp := TextEntities{
		tdCommon: tdCommon{Type: "textEntities"},
		Entities: entities,
	}

	return &textEntitiesTemp
}

// GetTextEntities Returns all entities (mentions, hashtags, cashtags, bot commands, bank card numbers, URLs, and email addresses) contained in the text. Can be called synchronously
// @param text The text in which to look for entites
func (client *Client) GetTextEntities(text string) (*TextEntities, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "getTextEntities",
		"text":  text,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var textEntities TextEntities
	err = json.Unmarshal(result.Raw, &textEntities)
	return &textEntities, err
}
