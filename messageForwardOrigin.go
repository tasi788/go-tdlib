package tdlib

import (
	"encoding/json"
	"fmt"
)

// MessageForwardOrigin Contains information about the origin of a forwarded message
type MessageForwardOrigin interface {
	GetMessageForwardOriginEnum() MessageForwardOriginEnum
}

// MessageForwardOriginEnum Alias for abstract MessageForwardOrigin 'Sub-Classes', used as constant-enum here
type MessageForwardOriginEnum string

// MessageForwardOrigin enums
const (
	MessageForwardOriginUserType          MessageForwardOriginEnum = "messageForwardOriginUser"
	MessageForwardOriginChatType          MessageForwardOriginEnum = "messageForwardOriginChat"
	MessageForwardOriginHiddenUserType    MessageForwardOriginEnum = "messageForwardOriginHiddenUser"
	MessageForwardOriginChannelType       MessageForwardOriginEnum = "messageForwardOriginChannel"
	MessageForwardOriginMessageImportType MessageForwardOriginEnum = "messageForwardOriginMessageImport"
)

func unmarshalMessageForwardOrigin(rawMsg *json.RawMessage) (MessageForwardOrigin, error) {

	if rawMsg == nil {
		return nil, nil
	}
	var objMap map[string]interface{}
	err := json.Unmarshal(*rawMsg, &objMap)
	if err != nil {
		return nil, err
	}

	switch MessageForwardOriginEnum(objMap["@type"].(string)) {
	case MessageForwardOriginUserType:
		var messageForwardOriginUser MessageForwardOriginUser
		err := json.Unmarshal(*rawMsg, &messageForwardOriginUser)
		return &messageForwardOriginUser, err

	case MessageForwardOriginChatType:
		var messageForwardOriginChat MessageForwardOriginChat
		err := json.Unmarshal(*rawMsg, &messageForwardOriginChat)
		return &messageForwardOriginChat, err

	case MessageForwardOriginHiddenUserType:
		var messageForwardOriginHiddenUser MessageForwardOriginHiddenUser
		err := json.Unmarshal(*rawMsg, &messageForwardOriginHiddenUser)
		return &messageForwardOriginHiddenUser, err

	case MessageForwardOriginChannelType:
		var messageForwardOriginChannel MessageForwardOriginChannel
		err := json.Unmarshal(*rawMsg, &messageForwardOriginChannel)
		return &messageForwardOriginChannel, err

	case MessageForwardOriginMessageImportType:
		var messageForwardOriginMessageImport MessageForwardOriginMessageImport
		err := json.Unmarshal(*rawMsg, &messageForwardOriginMessageImport)
		return &messageForwardOriginMessageImport, err

	default:
		return nil, fmt.Errorf("Error UnMarshaling, unknown type:" + objMap["@type"].(string))
	}
}

// MessageForwardOriginUser The message was originally sent by a known user
type MessageForwardOriginUser struct {
	tdCommon
	SenderUserId int64 `json:"sender_user_id"` // Identifier of the user that originally sent the message
}

// MessageType return the string telegram-type of MessageForwardOriginUser
func (messageForwardOriginUser *MessageForwardOriginUser) MessageType() string {
	return "messageForwardOriginUser"
}

// NewMessageForwardOriginUser creates a new MessageForwardOriginUser
//
// @param senderUserId Identifier of the user that originally sent the message
func NewMessageForwardOriginUser(senderUserId int64) *MessageForwardOriginUser {
	messageForwardOriginUserTemp := MessageForwardOriginUser{
		tdCommon:     tdCommon{Type: "messageForwardOriginUser"},
		SenderUserId: senderUserId,
	}

	return &messageForwardOriginUserTemp
}

// GetMessageForwardOriginEnum return the enum type of this object
func (messageForwardOriginUser *MessageForwardOriginUser) GetMessageForwardOriginEnum() MessageForwardOriginEnum {
	return MessageForwardOriginUserType
}

// MessageForwardOriginChat The message was originally sent on behalf of a chat
type MessageForwardOriginChat struct {
	tdCommon
	SenderChatId    int64  `json:"sender_chat_id"`   // Identifier of the chat that originally sent the message
	AuthorSignature string `json:"author_signature"` // For messages originally sent by an anonymous chat administrator, original message author signature
}

// MessageType return the string telegram-type of MessageForwardOriginChat
func (messageForwardOriginChat *MessageForwardOriginChat) MessageType() string {
	return "messageForwardOriginChat"
}

// NewMessageForwardOriginChat creates a new MessageForwardOriginChat
//
// @param senderChatId Identifier of the chat that originally sent the message
// @param authorSignature For messages originally sent by an anonymous chat administrator, original message author signature
func NewMessageForwardOriginChat(senderChatId int64, authorSignature string) *MessageForwardOriginChat {
	messageForwardOriginChatTemp := MessageForwardOriginChat{
		tdCommon:        tdCommon{Type: "messageForwardOriginChat"},
		SenderChatId:    senderChatId,
		AuthorSignature: authorSignature,
	}

	return &messageForwardOriginChatTemp
}

// GetMessageForwardOriginEnum return the enum type of this object
func (messageForwardOriginChat *MessageForwardOriginChat) GetMessageForwardOriginEnum() MessageForwardOriginEnum {
	return MessageForwardOriginChatType
}

// MessageForwardOriginHiddenUser The message was originally sent by a user, which is hidden by their privacy settings
type MessageForwardOriginHiddenUser struct {
	tdCommon
	SenderName string `json:"sender_name"` // Name of the sender
}

// MessageType return the string telegram-type of MessageForwardOriginHiddenUser
func (messageForwardOriginHiddenUser *MessageForwardOriginHiddenUser) MessageType() string {
	return "messageForwardOriginHiddenUser"
}

// NewMessageForwardOriginHiddenUser creates a new MessageForwardOriginHiddenUser
//
// @param senderName Name of the sender
func NewMessageForwardOriginHiddenUser(senderName string) *MessageForwardOriginHiddenUser {
	messageForwardOriginHiddenUserTemp := MessageForwardOriginHiddenUser{
		tdCommon:   tdCommon{Type: "messageForwardOriginHiddenUser"},
		SenderName: senderName,
	}

	return &messageForwardOriginHiddenUserTemp
}

// GetMessageForwardOriginEnum return the enum type of this object
func (messageForwardOriginHiddenUser *MessageForwardOriginHiddenUser) GetMessageForwardOriginEnum() MessageForwardOriginEnum {
	return MessageForwardOriginHiddenUserType
}

// MessageForwardOriginChannel The message was originally a post in a channel
type MessageForwardOriginChannel struct {
	tdCommon
	ChatId          int64  `json:"chat_id"`          // Identifier of the chat from which the message was originally forwarded
	MessageId       int64  `json:"message_id"`       // Message identifier of the original message
	AuthorSignature string `json:"author_signature"` // Original post author signature
}

// MessageType return the string telegram-type of MessageForwardOriginChannel
func (messageForwardOriginChannel *MessageForwardOriginChannel) MessageType() string {
	return "messageForwardOriginChannel"
}

// NewMessageForwardOriginChannel creates a new MessageForwardOriginChannel
//
// @param chatId Identifier of the chat from which the message was originally forwarded
// @param messageId Message identifier of the original message
// @param authorSignature Original post author signature
func NewMessageForwardOriginChannel(chatId int64, messageId int64, authorSignature string) *MessageForwardOriginChannel {
	messageForwardOriginChannelTemp := MessageForwardOriginChannel{
		tdCommon:        tdCommon{Type: "messageForwardOriginChannel"},
		ChatId:          chatId,
		MessageId:       messageId,
		AuthorSignature: authorSignature,
	}

	return &messageForwardOriginChannelTemp
}

// GetMessageForwardOriginEnum return the enum type of this object
func (messageForwardOriginChannel *MessageForwardOriginChannel) GetMessageForwardOriginEnum() MessageForwardOriginEnum {
	return MessageForwardOriginChannelType
}

// MessageForwardOriginMessageImport The message was imported from an exported message history
type MessageForwardOriginMessageImport struct {
	tdCommon
	SenderName string `json:"sender_name"` // Name of the sender
}

// MessageType return the string telegram-type of MessageForwardOriginMessageImport
func (messageForwardOriginMessageImport *MessageForwardOriginMessageImport) MessageType() string {
	return "messageForwardOriginMessageImport"
}

// NewMessageForwardOriginMessageImport creates a new MessageForwardOriginMessageImport
//
// @param senderName Name of the sender
func NewMessageForwardOriginMessageImport(senderName string) *MessageForwardOriginMessageImport {
	messageForwardOriginMessageImportTemp := MessageForwardOriginMessageImport{
		tdCommon:   tdCommon{Type: "messageForwardOriginMessageImport"},
		SenderName: senderName,
	}

	return &messageForwardOriginMessageImportTemp
}

// GetMessageForwardOriginEnum return the enum type of this object
func (messageForwardOriginMessageImport *MessageForwardOriginMessageImport) GetMessageForwardOriginEnum() MessageForwardOriginEnum {
	return MessageForwardOriginMessageImportType
}
