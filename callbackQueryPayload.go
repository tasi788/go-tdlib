package tdlib

import (
	"encoding/json"
	"fmt"
)

// CallbackQueryPayload Represents a payload of a callback query
type CallbackQueryPayload interface {
	GetCallbackQueryPayloadEnum() CallbackQueryPayloadEnum
}

// CallbackQueryPayloadEnum Alias for abstract CallbackQueryPayload 'Sub-Classes', used as constant-enum here
type CallbackQueryPayloadEnum string

// CallbackQueryPayload enums
const (
	CallbackQueryPayloadDataType             CallbackQueryPayloadEnum = "callbackQueryPayloadData"
	CallbackQueryPayloadDataWithPasswordType CallbackQueryPayloadEnum = "callbackQueryPayloadDataWithPassword"
	CallbackQueryPayloadGameType             CallbackQueryPayloadEnum = "callbackQueryPayloadGame"
)

func unmarshalCallbackQueryPayload(rawMsg *json.RawMessage) (CallbackQueryPayload, error) {

	if rawMsg == nil {
		return nil, nil
	}
	var objMap map[string]interface{}
	err := json.Unmarshal(*rawMsg, &objMap)
	if err != nil {
		return nil, err
	}

	switch CallbackQueryPayloadEnum(objMap["@type"].(string)) {
	case CallbackQueryPayloadDataType:
		var callbackQueryPayloadData CallbackQueryPayloadData
		err := json.Unmarshal(*rawMsg, &callbackQueryPayloadData)
		return &callbackQueryPayloadData, err

	case CallbackQueryPayloadDataWithPasswordType:
		var callbackQueryPayloadDataWithPassword CallbackQueryPayloadDataWithPassword
		err := json.Unmarshal(*rawMsg, &callbackQueryPayloadDataWithPassword)
		return &callbackQueryPayloadDataWithPassword, err

	case CallbackQueryPayloadGameType:
		var callbackQueryPayloadGame CallbackQueryPayloadGame
		err := json.Unmarshal(*rawMsg, &callbackQueryPayloadGame)
		return &callbackQueryPayloadGame, err

	default:
		return nil, fmt.Errorf("Error UnMarshaling, unknown type:" + objMap["@type"].(string))
	}
}

// CallbackQueryPayloadData The payload for a general callback button
type CallbackQueryPayloadData struct {
	tdCommon
	Data []byte `json:"data"` // Data that was attached to the callback button
}

// MessageType return the string telegram-type of CallbackQueryPayloadData
func (callbackQueryPayloadData *CallbackQueryPayloadData) MessageType() string {
	return "callbackQueryPayloadData"
}

// NewCallbackQueryPayloadData creates a new CallbackQueryPayloadData
//
// @param data Data that was attached to the callback button
func NewCallbackQueryPayloadData(data []byte) *CallbackQueryPayloadData {
	callbackQueryPayloadDataTemp := CallbackQueryPayloadData{
		tdCommon: tdCommon{Type: "callbackQueryPayloadData"},
		Data:     data,
	}

	return &callbackQueryPayloadDataTemp
}

// GetCallbackQueryPayloadEnum return the enum type of this object
func (callbackQueryPayloadData *CallbackQueryPayloadData) GetCallbackQueryPayloadEnum() CallbackQueryPayloadEnum {
	return CallbackQueryPayloadDataType
}

// CallbackQueryPayloadDataWithPassword The payload for a callback button requiring password
type CallbackQueryPayloadDataWithPassword struct {
	tdCommon
	Password string `json:"password"` // The password for the current user
	Data     []byte `json:"data"`     // Data that was attached to the callback button
}

// MessageType return the string telegram-type of CallbackQueryPayloadDataWithPassword
func (callbackQueryPayloadDataWithPassword *CallbackQueryPayloadDataWithPassword) MessageType() string {
	return "callbackQueryPayloadDataWithPassword"
}

// NewCallbackQueryPayloadDataWithPassword creates a new CallbackQueryPayloadDataWithPassword
//
// @param password The password for the current user
// @param data Data that was attached to the callback button
func NewCallbackQueryPayloadDataWithPassword(password string, data []byte) *CallbackQueryPayloadDataWithPassword {
	callbackQueryPayloadDataWithPasswordTemp := CallbackQueryPayloadDataWithPassword{
		tdCommon: tdCommon{Type: "callbackQueryPayloadDataWithPassword"},
		Password: password,
		Data:     data,
	}

	return &callbackQueryPayloadDataWithPasswordTemp
}

// GetCallbackQueryPayloadEnum return the enum type of this object
func (callbackQueryPayloadDataWithPassword *CallbackQueryPayloadDataWithPassword) GetCallbackQueryPayloadEnum() CallbackQueryPayloadEnum {
	return CallbackQueryPayloadDataWithPasswordType
}

// CallbackQueryPayloadGame The payload for a game callback button
type CallbackQueryPayloadGame struct {
	tdCommon
	GameShortName string `json:"game_short_name"` // A short name of the game that was attached to the callback button
}

// MessageType return the string telegram-type of CallbackQueryPayloadGame
func (callbackQueryPayloadGame *CallbackQueryPayloadGame) MessageType() string {
	return "callbackQueryPayloadGame"
}

// NewCallbackQueryPayloadGame creates a new CallbackQueryPayloadGame
//
// @param gameShortName A short name of the game that was attached to the callback button
func NewCallbackQueryPayloadGame(gameShortName string) *CallbackQueryPayloadGame {
	callbackQueryPayloadGameTemp := CallbackQueryPayloadGame{
		tdCommon:      tdCommon{Type: "callbackQueryPayloadGame"},
		GameShortName: gameShortName,
	}

	return &callbackQueryPayloadGameTemp
}

// GetCallbackQueryPayloadEnum return the enum type of this object
func (callbackQueryPayloadGame *CallbackQueryPayloadGame) GetCallbackQueryPayloadEnum() CallbackQueryPayloadEnum {
	return CallbackQueryPayloadGameType
}
