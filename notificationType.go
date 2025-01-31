package tdlib

import (
	"encoding/json"
	"fmt"
)

// NotificationType Contains detailed information about a notification
type NotificationType interface {
	GetNotificationTypeEnum() NotificationTypeEnum
}

// NotificationTypeEnum Alias for abstract NotificationType 'Sub-Classes', used as constant-enum here
type NotificationTypeEnum string

// NotificationType enums
const (
	NotificationTypeNewMessageType     NotificationTypeEnum = "notificationTypeNewMessage"
	NotificationTypeNewSecretChatType  NotificationTypeEnum = "notificationTypeNewSecretChat"
	NotificationTypeNewCallType        NotificationTypeEnum = "notificationTypeNewCall"
	NotificationTypeNewPushMessageType NotificationTypeEnum = "notificationTypeNewPushMessage"
)

func unmarshalNotificationType(rawMsg *json.RawMessage) (NotificationType, error) {

	if rawMsg == nil {
		return nil, nil
	}
	var objMap map[string]interface{}
	err := json.Unmarshal(*rawMsg, &objMap)
	if err != nil {
		return nil, err
	}

	switch NotificationTypeEnum(objMap["@type"].(string)) {
	case NotificationTypeNewMessageType:
		var notificationTypeNewMessage NotificationTypeNewMessage
		err := json.Unmarshal(*rawMsg, &notificationTypeNewMessage)
		return &notificationTypeNewMessage, err

	case NotificationTypeNewSecretChatType:
		var notificationTypeNewSecretChat NotificationTypeNewSecretChat
		err := json.Unmarshal(*rawMsg, &notificationTypeNewSecretChat)
		return &notificationTypeNewSecretChat, err

	case NotificationTypeNewCallType:
		var notificationTypeNewCall NotificationTypeNewCall
		err := json.Unmarshal(*rawMsg, &notificationTypeNewCall)
		return &notificationTypeNewCall, err

	case NotificationTypeNewPushMessageType:
		var notificationTypeNewPushMessage NotificationTypeNewPushMessage
		err := json.Unmarshal(*rawMsg, &notificationTypeNewPushMessage)
		return &notificationTypeNewPushMessage, err

	default:
		return nil, fmt.Errorf("Error UnMarshaling, unknown type:" + objMap["@type"].(string))
	}
}

// NotificationTypeNewMessage New message was received
type NotificationTypeNewMessage struct {
	tdCommon
	Message *Message `json:"message"` // The message
}

// MessageType return the string telegram-type of NotificationTypeNewMessage
func (notificationTypeNewMessage *NotificationTypeNewMessage) MessageType() string {
	return "notificationTypeNewMessage"
}

// NewNotificationTypeNewMessage creates a new NotificationTypeNewMessage
//
// @param message The message
func NewNotificationTypeNewMessage(message *Message) *NotificationTypeNewMessage {
	notificationTypeNewMessageTemp := NotificationTypeNewMessage{
		tdCommon: tdCommon{Type: "notificationTypeNewMessage"},
		Message:  message,
	}

	return &notificationTypeNewMessageTemp
}

// GetNotificationTypeEnum return the enum type of this object
func (notificationTypeNewMessage *NotificationTypeNewMessage) GetNotificationTypeEnum() NotificationTypeEnum {
	return NotificationTypeNewMessageType
}

// NotificationTypeNewSecretChat New secret chat was created
type NotificationTypeNewSecretChat struct {
	tdCommon
}

// MessageType return the string telegram-type of NotificationTypeNewSecretChat
func (notificationTypeNewSecretChat *NotificationTypeNewSecretChat) MessageType() string {
	return "notificationTypeNewSecretChat"
}

// NewNotificationTypeNewSecretChat creates a new NotificationTypeNewSecretChat
//
func NewNotificationTypeNewSecretChat() *NotificationTypeNewSecretChat {
	notificationTypeNewSecretChatTemp := NotificationTypeNewSecretChat{
		tdCommon: tdCommon{Type: "notificationTypeNewSecretChat"},
	}

	return &notificationTypeNewSecretChatTemp
}

// GetNotificationTypeEnum return the enum type of this object
func (notificationTypeNewSecretChat *NotificationTypeNewSecretChat) GetNotificationTypeEnum() NotificationTypeEnum {
	return NotificationTypeNewSecretChatType
}

// NotificationTypeNewCall New call was received
type NotificationTypeNewCall struct {
	tdCommon
	CallId int32 `json:"call_id"` // Call identifier
}

// MessageType return the string telegram-type of NotificationTypeNewCall
func (notificationTypeNewCall *NotificationTypeNewCall) MessageType() string {
	return "notificationTypeNewCall"
}

// NewNotificationTypeNewCall creates a new NotificationTypeNewCall
//
// @param callId Call identifier
func NewNotificationTypeNewCall(callId int32) *NotificationTypeNewCall {
	notificationTypeNewCallTemp := NotificationTypeNewCall{
		tdCommon: tdCommon{Type: "notificationTypeNewCall"},
		CallId:   callId,
	}

	return &notificationTypeNewCallTemp
}

// GetNotificationTypeEnum return the enum type of this object
func (notificationTypeNewCall *NotificationTypeNewCall) GetNotificationTypeEnum() NotificationTypeEnum {
	return NotificationTypeNewCallType
}

// NotificationTypeNewPushMessage New message was received through a push notification
type NotificationTypeNewPushMessage struct {
	tdCommon
	MessageId  int64              `json:"message_id"`  // The message identifier. The message will not be available in the chat history, but the ID can be used in viewMessages, or as reply_to_message_id
	SenderId   MessageSender      `json:"sender_id"`   // Identifier of the sender of the message. Corresponding user or chat may be inaccessible
	SenderName string             `json:"sender_name"` // Name of the sender
	IsOutgoing bool               `json:"is_outgoing"` // True, if the message is outgoing
	Content    PushMessageContent `json:"content"`     // Push message content
}

// MessageType return the string telegram-type of NotificationTypeNewPushMessage
func (notificationTypeNewPushMessage *NotificationTypeNewPushMessage) MessageType() string {
	return "notificationTypeNewPushMessage"
}

// NewNotificationTypeNewPushMessage creates a new NotificationTypeNewPushMessage
//
// @param messageId The message identifier. The message will not be available in the chat history, but the ID can be used in viewMessages, or as reply_to_message_id
// @param senderId Identifier of the sender of the message. Corresponding user or chat may be inaccessible
// @param senderName Name of the sender
// @param isOutgoing True, if the message is outgoing
// @param content Push message content
func NewNotificationTypeNewPushMessage(messageId int64, senderId MessageSender, senderName string, isOutgoing bool, content PushMessageContent) *NotificationTypeNewPushMessage {
	notificationTypeNewPushMessageTemp := NotificationTypeNewPushMessage{
		tdCommon:   tdCommon{Type: "notificationTypeNewPushMessage"},
		MessageId:  messageId,
		SenderId:   senderId,
		SenderName: senderName,
		IsOutgoing: isOutgoing,
		Content:    content,
	}

	return &notificationTypeNewPushMessageTemp
}

// UnmarshalJSON unmarshal to json
func (notificationTypeNewPushMessage *NotificationTypeNewPushMessage) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	tempObj := struct {
		tdCommon
		MessageId  int64  `json:"message_id"`  // The message identifier. The message will not be available in the chat history, but the ID can be used in viewMessages, or as reply_to_message_id
		SenderName string `json:"sender_name"` // Name of the sender
		IsOutgoing bool   `json:"is_outgoing"` // True, if the message is outgoing

	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	notificationTypeNewPushMessage.tdCommon = tempObj.tdCommon
	notificationTypeNewPushMessage.MessageId = tempObj.MessageId
	notificationTypeNewPushMessage.SenderName = tempObj.SenderName
	notificationTypeNewPushMessage.IsOutgoing = tempObj.IsOutgoing

	fieldSenderId, _ := unmarshalMessageSender(objMap["sender_id"])
	notificationTypeNewPushMessage.SenderId = fieldSenderId

	fieldContent, _ := unmarshalPushMessageContent(objMap["content"])
	notificationTypeNewPushMessage.Content = fieldContent

	return nil
}

// GetNotificationTypeEnum return the enum type of this object
func (notificationTypeNewPushMessage *NotificationTypeNewPushMessage) GetNotificationTypeEnum() NotificationTypeEnum {
	return NotificationTypeNewPushMessageType
}
