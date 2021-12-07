package tdlib

import (
	"encoding/json"
	"fmt"
)

// AnimatedEmoji Describes an animated representation of an emoji
type AnimatedEmoji struct {
	tdCommon
	Sticker         *Sticker `json:"sticker"`          // Animated sticker for the emoji
	FitzpatrickType int32    `json:"fitzpatrick_type"` // Emoji modifier fitzpatrick type; 0-6; 0 if none
	Sound           *File    `json:"sound"`            // File containing the sound to be played when the animated emoji is clicked if any; may be null. The sound is encoded with the Opus codec, and stored inside an OGG container
}

// MessageType return the string telegram-type of AnimatedEmoji
func (animatedEmoji *AnimatedEmoji) MessageType() string {
	return "animatedEmoji"
}

// NewAnimatedEmoji creates a new AnimatedEmoji
//
// @param sticker Animated sticker for the emoji
// @param fitzpatrickType Emoji modifier fitzpatrick type; 0-6; 0 if none
// @param sound File containing the sound to be played when the animated emoji is clicked if any; may be null. The sound is encoded with the Opus codec, and stored inside an OGG container
func NewAnimatedEmoji(sticker *Sticker, fitzpatrickType int32, sound *File) *AnimatedEmoji {
	animatedEmojiTemp := AnimatedEmoji{
		tdCommon:        tdCommon{Type: "animatedEmoji"},
		Sticker:         sticker,
		FitzpatrickType: fitzpatrickType,
		Sound:           sound,
	}

	return &animatedEmojiTemp
}

// GetAnimatedEmoji Returns an animated emoji corresponding to a given emoji. Returns a 404 error if the emoji has no animated emoji
// @param emoji The emoji
func (client *Client) GetAnimatedEmoji(emoji string) (*AnimatedEmoji, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "getAnimatedEmoji",
		"emoji": emoji,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var animatedEmoji AnimatedEmoji
	err = json.Unmarshal(result.Raw, &animatedEmoji)
	return &animatedEmoji, err
}
