package tdlib

import (
	"encoding/json"
	"fmt"
)

// Background Describes a chat background
type Background struct {
	tdCommon
	Id        JSONInt64      `json:"id"`         // Unique background identifier
	IsDefault bool           `json:"is_default"` // True, if this is one of default backgrounds
	IsDark    bool           `json:"is_dark"`    // True, if the background is dark and is recommended to be used with dark theme
	Name      string         `json:"name"`       // Unique background name
	Document  *Document      `json:"document"`   // Document with the background; may be null. Null only for filled backgrounds
	Type      BackgroundType `json:"type"`       // Type of the background
}

// MessageType return the string telegram-type of Background
func (background *Background) MessageType() string {
	return "background"
}

// NewBackground creates a new Background
//
// @param id Unique background identifier
// @param isDefault True, if this is one of default backgrounds
// @param isDark True, if the background is dark and is recommended to be used with dark theme
// @param name Unique background name
// @param document Document with the background; may be null. Null only for filled backgrounds
// @param typeParam Type of the background
func NewBackground(id JSONInt64, isDefault bool, isDark bool, name string, document *Document, typeParam BackgroundType) *Background {
	backgroundTemp := Background{
		tdCommon:  tdCommon{Type: "background"},
		Id:        id,
		IsDefault: isDefault,
		IsDark:    isDark,
		Name:      name,
		Document:  document,
		Type:      typeParam,
	}

	return &backgroundTemp
}

// UnmarshalJSON unmarshal to json
func (background *Background) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	tempObj := struct {
		tdCommon
		Id        JSONInt64 `json:"id"`         // Unique background identifier
		IsDefault bool      `json:"is_default"` // True, if this is one of default backgrounds
		IsDark    bool      `json:"is_dark"`    // True, if the background is dark and is recommended to be used with dark theme
		Name      string    `json:"name"`       // Unique background name
		Document  *Document `json:"document"`   // Document with the background; may be null. Null only for filled backgrounds

	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	background.tdCommon = tempObj.tdCommon
	background.Id = tempObj.Id
	background.IsDefault = tempObj.IsDefault
	background.IsDark = tempObj.IsDark
	background.Name = tempObj.Name
	background.Document = tempObj.Document

	fieldType, _ := unmarshalBackgroundType(objMap["type"])
	background.Type = fieldType

	return nil
}

// SearchBackground Searches for a background by its name
// @param name The name of the background
func (client *Client) SearchBackground(name string) (*Background, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "searchBackground",
		"name":  name,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var background Background
	err = json.Unmarshal(result.Raw, &background)
	return &background, err
}

// SetBackground Changes the background selected by the user; adds background to the list of installed backgrounds
// @param background The input background to use; pass null to create a new filled backgrounds or to remove the current background
// @param typeParam Background type; pass null to use the default type of the remote background or to remove the current background
// @param forDarkTheme True, if the background is chosen for dark theme
func (client *Client) SetBackground(background InputBackground, typeParam BackgroundType, forDarkTheme bool) (*Background, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":          "setBackground",
		"background":     background,
		"type":           typeParam,
		"for_dark_theme": forDarkTheme,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var backgroundDummy Background
	err = json.Unmarshal(result.Raw, &backgroundDummy)
	return &backgroundDummy, err
}
