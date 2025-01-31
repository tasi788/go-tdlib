package tdlib

import (
	"encoding/json"
)

// TextEntity Represents a part of the text that needs to be formatted in some unusual way
type TextEntity struct {
	tdCommon
	Offset int32          `json:"offset"` // Offset of the entity, in UTF-16 code units
	Length int32          `json:"length"` // Length of the entity, in UTF-16 code units
	Type   TextEntityType `json:"type"`   // Type of the entity
}

// MessageType return the string telegram-type of TextEntity
func (textEntity *TextEntity) MessageType() string {
	return "textEntity"
}

// NewTextEntity creates a new TextEntity
//
// @param offset Offset of the entity, in UTF-16 code units
// @param length Length of the entity, in UTF-16 code units
// @param typeParam Type of the entity
func NewTextEntity(offset int32, length int32, typeParam TextEntityType) *TextEntity {
	textEntityTemp := TextEntity{
		tdCommon: tdCommon{Type: "textEntity"},
		Offset:   offset,
		Length:   length,
		Type:     typeParam,
	}

	return &textEntityTemp
}

// UnmarshalJSON unmarshal to json
func (textEntity *TextEntity) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	tempObj := struct {
		tdCommon
		Offset int32 `json:"offset"` // Offset of the entity, in UTF-16 code units
		Length int32 `json:"length"` // Length of the entity, in UTF-16 code units

	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	textEntity.tdCommon = tempObj.tdCommon
	textEntity.Offset = tempObj.Offset
	textEntity.Length = tempObj.Length

	fieldType, _ := unmarshalTextEntityType(objMap["type"])
	textEntity.Type = fieldType

	return nil
}
