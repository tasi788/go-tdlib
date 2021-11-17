package tdlib

import (
	"encoding/json"
	"fmt"
)

// MessageFileType Contains information about a file with messages exported from another app
type MessageFileType interface {
	GetMessageFileTypeEnum() MessageFileTypeEnum
}

// MessageFileTypeEnum Alias for abstract MessageFileType 'Sub-Classes', used as constant-enum here
type MessageFileTypeEnum string

// MessageFileType enums
const (
	MessageFileTypePrivateType MessageFileTypeEnum = "messageFileTypePrivate"
	MessageFileTypeGroupType   MessageFileTypeEnum = "messageFileTypeGroup"
	MessageFileTypeUnknownType MessageFileTypeEnum = "messageFileTypeUnknown"
)

func unmarshalMessageFileType(rawMsg *json.RawMessage) (MessageFileType, error) {

	if rawMsg == nil {
		return nil, nil
	}
	var objMap map[string]interface{}
	err := json.Unmarshal(*rawMsg, &objMap)
	if err != nil {
		return nil, err
	}

	switch MessageFileTypeEnum(objMap["@type"].(string)) {
	case MessageFileTypePrivateType:
		var messageFileTypePrivate MessageFileTypePrivate
		err := json.Unmarshal(*rawMsg, &messageFileTypePrivate)
		return &messageFileTypePrivate, err

	case MessageFileTypeGroupType:
		var messageFileTypeGroup MessageFileTypeGroup
		err := json.Unmarshal(*rawMsg, &messageFileTypeGroup)
		return &messageFileTypeGroup, err

	case MessageFileTypeUnknownType:
		var messageFileTypeUnknown MessageFileTypeUnknown
		err := json.Unmarshal(*rawMsg, &messageFileTypeUnknown)
		return &messageFileTypeUnknown, err

	default:
		return nil, fmt.Errorf("Error UnMarshaling, unknown type:" + objMap["@type"].(string))
	}
}

// MessageFileTypePrivate The messages was exported from a private chat
type MessageFileTypePrivate struct {
	tdCommon
	Name string `json:"name"` // Name of the other party; may be empty if unrecognized
}

// MessageType return the string telegram-type of MessageFileTypePrivate
func (messageFileTypePrivate *MessageFileTypePrivate) MessageType() string {
	return "messageFileTypePrivate"
}

// NewMessageFileTypePrivate creates a new MessageFileTypePrivate
//
// @param name Name of the other party; may be empty if unrecognized
func NewMessageFileTypePrivate(name string) *MessageFileTypePrivate {
	messageFileTypePrivateTemp := MessageFileTypePrivate{
		tdCommon: tdCommon{Type: "messageFileTypePrivate"},
		Name:     name,
	}

	return &messageFileTypePrivateTemp
}

// GetMessageFileTypeEnum return the enum type of this object
func (messageFileTypePrivate *MessageFileTypePrivate) GetMessageFileTypeEnum() MessageFileTypeEnum {
	return MessageFileTypePrivateType
}

// MessageFileTypeGroup The messages was exported from a group chat
type MessageFileTypeGroup struct {
	tdCommon
	Title string `json:"title"` // Title of the group chat; may be empty if unrecognized
}

// MessageType return the string telegram-type of MessageFileTypeGroup
func (messageFileTypeGroup *MessageFileTypeGroup) MessageType() string {
	return "messageFileTypeGroup"
}

// NewMessageFileTypeGroup creates a new MessageFileTypeGroup
//
// @param title Title of the group chat; may be empty if unrecognized
func NewMessageFileTypeGroup(title string) *MessageFileTypeGroup {
	messageFileTypeGroupTemp := MessageFileTypeGroup{
		tdCommon: tdCommon{Type: "messageFileTypeGroup"},
		Title:    title,
	}

	return &messageFileTypeGroupTemp
}

// GetMessageFileTypeEnum return the enum type of this object
func (messageFileTypeGroup *MessageFileTypeGroup) GetMessageFileTypeEnum() MessageFileTypeEnum {
	return MessageFileTypeGroupType
}

// MessageFileTypeUnknown The messages was exported from a chat of unknown type
type MessageFileTypeUnknown struct {
	tdCommon
}

// MessageType return the string telegram-type of MessageFileTypeUnknown
func (messageFileTypeUnknown *MessageFileTypeUnknown) MessageType() string {
	return "messageFileTypeUnknown"
}

// NewMessageFileTypeUnknown creates a new MessageFileTypeUnknown
//
func NewMessageFileTypeUnknown() *MessageFileTypeUnknown {
	messageFileTypeUnknownTemp := MessageFileTypeUnknown{
		tdCommon: tdCommon{Type: "messageFileTypeUnknown"},
	}

	return &messageFileTypeUnknownTemp
}

// GetMessageFileTypeEnum return the enum type of this object
func (messageFileTypeUnknown *MessageFileTypeUnknown) GetMessageFileTypeEnum() MessageFileTypeEnum {
	return MessageFileTypeUnknownType
}

// GetMessageFileType Returns information about a file with messages exported from another app
// @param messageFileHead Beginning of the message file; up to 100 first lines
func (client *Client) GetMessageFileType(messageFileHead string) (MessageFileType, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":             "getMessageFileType",
		"message_file_head": messageFileHead,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	switch MessageFileTypeEnum(result.Data["@type"].(string)) {

	case MessageFileTypePrivateType:
		var messageFileType MessageFileTypePrivate
		err = json.Unmarshal(result.Raw, &messageFileType)
		return &messageFileType, err

	case MessageFileTypeGroupType:
		var messageFileType MessageFileTypeGroup
		err = json.Unmarshal(result.Raw, &messageFileType)
		return &messageFileType, err

	case MessageFileTypeUnknownType:
		var messageFileType MessageFileTypeUnknown
		err = json.Unmarshal(result.Raw, &messageFileType)
		return &messageFileType, err

	default:
		return nil, fmt.Errorf("Invalid type")
	}
}
