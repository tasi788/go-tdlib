package tdlib

import (
	"encoding/json"
	"fmt"
)

// DiceStickers Contains animated stickers which must be used for dice animation rendering
type DiceStickers interface {
	GetDiceStickersEnum() DiceStickersEnum
}

// DiceStickersEnum Alias for abstract DiceStickers 'Sub-Classes', used as constant-enum here
type DiceStickersEnum string

// DiceStickers enums
const (
	DiceStickersRegularType     DiceStickersEnum = "diceStickersRegular"
	DiceStickersSlotMachineType DiceStickersEnum = "diceStickersSlotMachine"
)

func unmarshalDiceStickers(rawMsg *json.RawMessage) (DiceStickers, error) {

	if rawMsg == nil {
		return nil, nil
	}
	var objMap map[string]interface{}
	err := json.Unmarshal(*rawMsg, &objMap)
	if err != nil {
		return nil, err
	}

	switch DiceStickersEnum(objMap["@type"].(string)) {
	case DiceStickersRegularType:
		var diceStickersRegular DiceStickersRegular
		err := json.Unmarshal(*rawMsg, &diceStickersRegular)
		return &diceStickersRegular, err

	case DiceStickersSlotMachineType:
		var diceStickersSlotMachine DiceStickersSlotMachine
		err := json.Unmarshal(*rawMsg, &diceStickersSlotMachine)
		return &diceStickersSlotMachine, err

	default:
		return nil, fmt.Errorf("Error UnMarshaling, unknown type:" + objMap["@type"].(string))
	}
}

// DiceStickersRegular A regular animated sticker
type DiceStickersRegular struct {
	tdCommon
	Sticker *Sticker `json:"sticker"` // The animated sticker with the dice animation
}

// MessageType return the string telegram-type of DiceStickersRegular
func (diceStickersRegular *DiceStickersRegular) MessageType() string {
	return "diceStickersRegular"
}

// NewDiceStickersRegular creates a new DiceStickersRegular
//
// @param sticker The animated sticker with the dice animation
func NewDiceStickersRegular(sticker *Sticker) *DiceStickersRegular {
	diceStickersRegularTemp := DiceStickersRegular{
		tdCommon: tdCommon{Type: "diceStickersRegular"},
		Sticker:  sticker,
	}

	return &diceStickersRegularTemp
}

// GetDiceStickersEnum return the enum type of this object
func (diceStickersRegular *DiceStickersRegular) GetDiceStickersEnum() DiceStickersEnum {
	return DiceStickersRegularType
}

// DiceStickersSlotMachine Animated stickers to be combined into a slot machine
type DiceStickersSlotMachine struct {
	tdCommon
	Background *Sticker `json:"background"`  // The animated sticker with the slot machine background. The background animation must start playing after all reel animations finish
	Lever      *Sticker `json:"lever"`       // The animated sticker with the lever animation. The lever animation must play once in the initial dice state
	LeftReel   *Sticker `json:"left_reel"`   // The animated sticker with the left reel
	CenterReel *Sticker `json:"center_reel"` // The animated sticker with the center reel
	RightReel  *Sticker `json:"right_reel"`  // The animated sticker with the right reel
}

// MessageType return the string telegram-type of DiceStickersSlotMachine
func (diceStickersSlotMachine *DiceStickersSlotMachine) MessageType() string {
	return "diceStickersSlotMachine"
}

// NewDiceStickersSlotMachine creates a new DiceStickersSlotMachine
//
// @param background The animated sticker with the slot machine background. The background animation must start playing after all reel animations finish
// @param lever The animated sticker with the lever animation. The lever animation must play once in the initial dice state
// @param leftReel The animated sticker with the left reel
// @param centerReel The animated sticker with the center reel
// @param rightReel The animated sticker with the right reel
func NewDiceStickersSlotMachine(background *Sticker, lever *Sticker, leftReel *Sticker, centerReel *Sticker, rightReel *Sticker) *DiceStickersSlotMachine {
	diceStickersSlotMachineTemp := DiceStickersSlotMachine{
		tdCommon:   tdCommon{Type: "diceStickersSlotMachine"},
		Background: background,
		Lever:      lever,
		LeftReel:   leftReel,
		CenterReel: centerReel,
		RightReel:  rightReel,
	}

	return &diceStickersSlotMachineTemp
}

// GetDiceStickersEnum return the enum type of this object
func (diceStickersSlotMachine *DiceStickersSlotMachine) GetDiceStickersEnum() DiceStickersEnum {
	return DiceStickersSlotMachineType
}
