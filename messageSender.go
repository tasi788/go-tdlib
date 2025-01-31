package tdlib

import (
	"encoding/json"
	"fmt"
)

// MessageSender Contains information about the sender of a message
type MessageSender interface {
	GetMessageSenderEnum() MessageSenderEnum
}

// MessageSenderEnum Alias for abstract MessageSender 'Sub-Classes', used as constant-enum here
type MessageSenderEnum string

// MessageSender enums
const (
	MessageSenderUserType MessageSenderEnum = "messageSenderUser"
	MessageSenderChatType MessageSenderEnum = "messageSenderChat"
)

func unmarshalMessageSender(rawMsg *json.RawMessage) (MessageSender, error) {

	if rawMsg == nil {
		return nil, nil
	}
	var objMap map[string]interface{}
	err := json.Unmarshal(*rawMsg, &objMap)
	if err != nil {
		return nil, err
	}

	switch MessageSenderEnum(objMap["@type"].(string)) {
	case MessageSenderUserType:
		var messageSenderUser MessageSenderUser
		err := json.Unmarshal(*rawMsg, &messageSenderUser)
		return &messageSenderUser, err

	case MessageSenderChatType:
		var messageSenderChat MessageSenderChat
		err := json.Unmarshal(*rawMsg, &messageSenderChat)
		return &messageSenderChat, err

	default:
		return nil, fmt.Errorf("Error UnMarshaling, unknown type:" + objMap["@type"].(string))
	}
}

// MessageSenderUser The message was sent by a known user
type MessageSenderUser struct {
	tdCommon
	UserId int64 `json:"user_id"` // Identifier of the user that sent the message
}

// MessageType return the string telegram-type of MessageSenderUser
func (messageSenderUser *MessageSenderUser) MessageType() string {
	return "messageSenderUser"
}

// NewMessageSenderUser creates a new MessageSenderUser
//
// @param userId Identifier of the user that sent the message
func NewMessageSenderUser(userId int64) *MessageSenderUser {
	messageSenderUserTemp := MessageSenderUser{
		tdCommon: tdCommon{Type: "messageSenderUser"},
		UserId:   userId,
	}

	return &messageSenderUserTemp
}

// GetMessageSenderEnum return the enum type of this object
func (messageSenderUser *MessageSenderUser) GetMessageSenderEnum() MessageSenderEnum {
	return MessageSenderUserType
}

// MessageSenderChat The message was sent on behalf of a chat
type MessageSenderChat struct {
	tdCommon
	ChatId int64 `json:"chat_id"` // Identifier of the chat that sent the message
}

// MessageType return the string telegram-type of MessageSenderChat
func (messageSenderChat *MessageSenderChat) MessageType() string {
	return "messageSenderChat"
}

// NewMessageSenderChat creates a new MessageSenderChat
//
// @param chatId Identifier of the chat that sent the message
func NewMessageSenderChat(chatId int64) *MessageSenderChat {
	messageSenderChatTemp := MessageSenderChat{
		tdCommon: tdCommon{Type: "messageSenderChat"},
		ChatId:   chatId,
	}

	return &messageSenderChatTemp
}

// GetMessageSenderEnum return the enum type of this object
func (messageSenderChat *MessageSenderChat) GetMessageSenderEnum() MessageSenderEnum {
	return MessageSenderChatType
}
