package tdlib

import (
	"encoding/json"
	"fmt"
)

// Emojis Represents a list of emoji
type Emojis struct {
	tdCommon
	Emojis []string `json:"emojis"` // List of emojis
}

// MessageType return the string telegram-type of Emojis
func (emojis *Emojis) MessageType() string {
	return "emojis"
}

// NewEmojis creates a new Emojis
//
// @param emojis List of emojis
func NewEmojis(emojis []string) *Emojis {
	emojisTemp := Emojis{
		tdCommon: tdCommon{Type: "emojis"},
		Emojis:   emojis,
	}

	return &emojisTemp
}

// GetStickerEmojis Returns emoji corresponding to a sticker. The list is only for informational purposes, because a sticker is always sent with a fixed emoji from the corresponding Sticker object
// @param sticker Sticker file identifier
func (client *Client) GetStickerEmojis(sticker InputFile) (*Emojis, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "getStickerEmojis",
		"sticker": sticker,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var emojis Emojis
	err = json.Unmarshal(result.Raw, &emojis)
	return &emojis, err
}

// SearchEmojis Searches for emojis by keywords. Supported only if the file database is enabled
// @param text Text to search for
// @param exactMatch True, if only emojis, which exactly match text needs to be returned
// @param inputLanguageCodes List of possible IETF language tags of the user's input language; may be empty if unknown
func (client *Client) SearchEmojis(text string, exactMatch bool, inputLanguageCodes []string) (*Emojis, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":                "searchEmojis",
		"text":                 text,
		"exact_match":          exactMatch,
		"input_language_codes": inputLanguageCodes,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var emojis Emojis
	err = json.Unmarshal(result.Raw, &emojis)
	return &emojis, err
}
