package tdlib

import (
	"encoding/json"
	"fmt"
)

// PublicChatType Describes a type of public chats
type PublicChatType interface {
	GetPublicChatTypeEnum() PublicChatTypeEnum
}

// PublicChatTypeEnum Alias for abstract PublicChatType 'Sub-Classes', used as constant-enum here
type PublicChatTypeEnum string

// PublicChatType enums
const (
	PublicChatTypeHasUsernameType     PublicChatTypeEnum = "publicChatTypeHasUsername"
	PublicChatTypeIsLocationBasedType PublicChatTypeEnum = "publicChatTypeIsLocationBased"
)

func unmarshalPublicChatType(rawMsg *json.RawMessage) (PublicChatType, error) {

	if rawMsg == nil {
		return nil, nil
	}
	var objMap map[string]interface{}
	err := json.Unmarshal(*rawMsg, &objMap)
	if err != nil {
		return nil, err
	}

	switch PublicChatTypeEnum(objMap["@type"].(string)) {
	case PublicChatTypeHasUsernameType:
		var publicChatTypeHasUsername PublicChatTypeHasUsername
		err := json.Unmarshal(*rawMsg, &publicChatTypeHasUsername)
		return &publicChatTypeHasUsername, err

	case PublicChatTypeIsLocationBasedType:
		var publicChatTypeIsLocationBased PublicChatTypeIsLocationBased
		err := json.Unmarshal(*rawMsg, &publicChatTypeIsLocationBased)
		return &publicChatTypeIsLocationBased, err

	default:
		return nil, fmt.Errorf("Error UnMarshaling, unknown type:" + objMap["@type"].(string))
	}
}

// PublicChatTypeHasUsername The chat is public, because it has username
type PublicChatTypeHasUsername struct {
	tdCommon
}

// MessageType return the string telegram-type of PublicChatTypeHasUsername
func (publicChatTypeHasUsername *PublicChatTypeHasUsername) MessageType() string {
	return "publicChatTypeHasUsername"
}

// NewPublicChatTypeHasUsername creates a new PublicChatTypeHasUsername
//
func NewPublicChatTypeHasUsername() *PublicChatTypeHasUsername {
	publicChatTypeHasUsernameTemp := PublicChatTypeHasUsername{
		tdCommon: tdCommon{Type: "publicChatTypeHasUsername"},
	}

	return &publicChatTypeHasUsernameTemp
}

// GetPublicChatTypeEnum return the enum type of this object
func (publicChatTypeHasUsername *PublicChatTypeHasUsername) GetPublicChatTypeEnum() PublicChatTypeEnum {
	return PublicChatTypeHasUsernameType
}

// PublicChatTypeIsLocationBased The chat is public, because it is a location-based supergroup
type PublicChatTypeIsLocationBased struct {
	tdCommon
}

// MessageType return the string telegram-type of PublicChatTypeIsLocationBased
func (publicChatTypeIsLocationBased *PublicChatTypeIsLocationBased) MessageType() string {
	return "publicChatTypeIsLocationBased"
}

// NewPublicChatTypeIsLocationBased creates a new PublicChatTypeIsLocationBased
//
func NewPublicChatTypeIsLocationBased() *PublicChatTypeIsLocationBased {
	publicChatTypeIsLocationBasedTemp := PublicChatTypeIsLocationBased{
		tdCommon: tdCommon{Type: "publicChatTypeIsLocationBased"},
	}

	return &publicChatTypeIsLocationBasedTemp
}

// GetPublicChatTypeEnum return the enum type of this object
func (publicChatTypeIsLocationBased *PublicChatTypeIsLocationBased) GetPublicChatTypeEnum() PublicChatTypeEnum {
	return PublicChatTypeIsLocationBasedType
}
