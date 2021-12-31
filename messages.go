package tdlib

import (
	"encoding/json"
	"fmt"
)

// Messages Contains a list of messages
type Messages struct {
	tdCommon
	TotalCount int32     `json:"total_count"` // Approximate total count of messages found
	Messages   []Message `json:"messages"`    // List of messages; messages may be null
}

// MessageType return the string telegram-type of Messages
func (messages *Messages) MessageType() string {
	return "messages"
}

// NewMessages creates a new Messages
//
// @param totalCount Approximate total count of messages found
// @param messages List of messages; messages may be null
func NewMessages(totalCount int32, messages []Message) *Messages {
	messagesTemp := Messages{
		tdCommon:   tdCommon{Type: "messages"},
		TotalCount: totalCount,
		Messages:   messages,
	}

	return &messagesTemp
}

// GetMessages Returns information about messages. If a message is not found, returns null on the corresponding position of the result
// @param chatId Identifier of the chat the messages belong to
// @param messageIds Identifiers of the messages to get
func (client *Client) GetMessages(chatId int64, messageIds []int64) (*Messages, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":       "getMessages",
		"chat_id":     chatId,
		"message_ids": messageIds,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var messages Messages
	err = json.Unmarshal(result.Raw, &messages)
	return &messages, err
}

// GetChatHistory Returns messages in a chat. The messages are returned in a reverse chronological order (i.e., in order of decreasing message_id). For optimal performance, the number of returned messages is chosen by TDLib. This is an offline request if only_local is true
// @param chatId Chat identifier
// @param fromMessageId Identifier of the message starting from which history must be fetched; use 0 to get results from the last message
// @param offset Specify 0 to get results from exactly the from_message_id or a negative offset up to 99 to get additionally some newer messages
// @param limit The maximum number of messages to be returned; must be positive and can't be greater than 100. If the offset is negative, the limit must be greater than or equal to -offset. For optimal performance, the number of returned messages is chosen by TDLib and can be smaller than the specified limit
// @param onlyLocal If true, returns only messages that are available locally without sending network requests
func (client *Client) GetChatHistory(chatId int64, fromMessageId int64, offset int32, limit int32, onlyLocal bool) (*Messages, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":           "getChatHistory",
		"chat_id":         chatId,
		"from_message_id": fromMessageId,
		"offset":          offset,
		"limit":           limit,
		"only_local":      onlyLocal,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var messages Messages
	err = json.Unmarshal(result.Raw, &messages)
	return &messages, err
}

// GetMessageThreadHistory Returns messages in a message thread of a message. Can be used only if message.can_get_message_thread == true. Message thread of a channel message is in the channel's linked supergroup. The messages are returned in a reverse chronological order (i.e., in order of decreasing message_id). For optimal performance, the number of returned messages is chosen by TDLib
// @param chatId Chat identifier
// @param messageId Message identifier, which thread history needs to be returned
// @param fromMessageId Identifier of the message starting from which history must be fetched; use 0 to get results from the last message
// @param offset Specify 0 to get results from exactly the from_message_id or a negative offset up to 99 to get additionally some newer messages
// @param limit The maximum number of messages to be returned; must be positive and can't be greater than 100. If the offset is negative, the limit must be greater than or equal to -offset. For optimal performance, the number of returned messages is chosen by TDLib and can be smaller than the specified limit
func (client *Client) GetMessageThreadHistory(chatId int64, messageId int64, fromMessageId int64, offset int32, limit int32) (*Messages, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":           "getMessageThreadHistory",
		"chat_id":         chatId,
		"message_id":      messageId,
		"from_message_id": fromMessageId,
		"offset":          offset,
		"limit":           limit,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var messages Messages
	err = json.Unmarshal(result.Raw, &messages)
	return &messages, err
}

// SearchChatMessages Searches for messages with given words in the chat. Returns the results in reverse chronological order, i.e. in order of decreasing message_id. Cannot be used in secret chats with a non-empty query (searchSecretMessages must be used instead), or without an enabled message database. For optimal performance, the number of returned messages is chosen by TDLib and can be smaller than the specified limit
// @param chatId Identifier of the chat in which to search messages
// @param query Query to search for
// @param senderId Identifier of the sender of messages to search for; pass null to search for messages from any sender. Not supported in secret chats
// @param fromMessageId Identifier of the message starting from which history must be fetched; use 0 to get results from the last message
// @param offset Specify 0 to get results from exactly the from_message_id or a negative offset to get the specified message and some newer messages
// @param limit The maximum number of messages to be returned; must be positive and can't be greater than 100. If the offset is negative, the limit must be greater than -offset. For optimal performance, the number of returned messages is chosen by TDLib and can be smaller than the specified limit
// @param filter Additional filter for messages to search; pass null to search for all messages
// @param messageThreadId If not 0, only messages in the specified thread will be returned; supergroups only
func (client *Client) SearchChatMessages(chatId int64, query string, senderId MessageSender, fromMessageId int64, offset int32, limit int32, filter SearchMessagesFilter, messageThreadId int64) (*Messages, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":             "searchChatMessages",
		"chat_id":           chatId,
		"query":             query,
		"sender_id":         senderId,
		"from_message_id":   fromMessageId,
		"offset":            offset,
		"limit":             limit,
		"filter":            filter,
		"message_thread_id": messageThreadId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var messages Messages
	err = json.Unmarshal(result.Raw, &messages)
	return &messages, err
}

// SearchMessages Searches for messages in all chats except secret chats. Returns the results in reverse chronological order (i.e., in order of decreasing (date, chat_id, message_id)). For optimal performance, the number of returned messages is chosen by TDLib and can be smaller than the specified limit
// @param chatList Chat list in which to search messages; pass null to search in all chats regardless of their chat list. Only Main and Archive chat lists are supported
// @param query Query to search for
// @param offsetDate The date of the message starting from which the results need to be fetched. Use 0 or any date in the future to get results from the last message
// @param offsetChatId The chat identifier of the last found message, or 0 for the first request
// @param offsetMessageId The message identifier of the last found message, or 0 for the first request
// @param limit The maximum number of messages to be returned; up to 100. For optimal performance, the number of returned messages is chosen by TDLib and can be smaller than the specified limit
// @param filter Additional filter for messages to search; pass null to search for all messages. Filters searchMessagesFilterMention, searchMessagesFilterUnreadMention, searchMessagesFilterFailedToSend and searchMessagesFilterPinned are unsupported in this function
// @param minDate If not 0, the minimum date of the messages to return
// @param maxDate If not 0, the maximum date of the messages to return
func (client *Client) SearchMessages(chatList ChatList, query string, offsetDate int32, offsetChatId int64, offsetMessageId int64, limit int32, filter SearchMessagesFilter, minDate int32, maxDate int32) (*Messages, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":             "searchMessages",
		"chat_list":         chatList,
		"query":             query,
		"offset_date":       offsetDate,
		"offset_chat_id":    offsetChatId,
		"offset_message_id": offsetMessageId,
		"limit":             limit,
		"filter":            filter,
		"min_date":          minDate,
		"max_date":          maxDate,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var messages Messages
	err = json.Unmarshal(result.Raw, &messages)
	return &messages, err
}

// SearchCallMessages Searches for call messages. Returns the results in reverse chronological order (i. e., in order of decreasing message_id). For optimal performance, the number of returned messages is chosen by TDLib
// @param fromMessageId Identifier of the message from which to search; use 0 to get results from the last message
// @param limit The maximum number of messages to be returned; up to 100. For optimal performance, the number of returned messages is chosen by TDLib and can be smaller than the specified limit
// @param onlyMissed If true, returns only messages with missed/declined calls
func (client *Client) SearchCallMessages(fromMessageId int64, limit int32, onlyMissed bool) (*Messages, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":           "searchCallMessages",
		"from_message_id": fromMessageId,
		"limit":           limit,
		"only_missed":     onlyMissed,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var messages Messages
	err = json.Unmarshal(result.Raw, &messages)
	return &messages, err
}

// SearchChatRecentLocationMessages Returns information about the recent locations of chat members that were sent to the chat. Returns up to 1 location message per user
// @param chatId Chat identifier
// @param limit The maximum number of messages to be returned
func (client *Client) SearchChatRecentLocationMessages(chatId int64, limit int32) (*Messages, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "searchChatRecentLocationMessages",
		"chat_id": chatId,
		"limit":   limit,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var messages Messages
	err = json.Unmarshal(result.Raw, &messages)
	return &messages, err
}

// GetActiveLiveLocationMessages Returns all active live locations that need to be updated by the application. The list is persistent across application restarts only if the message database is used
func (client *Client) GetActiveLiveLocationMessages() (*Messages, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "getActiveLiveLocationMessages",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var messages Messages
	err = json.Unmarshal(result.Raw, &messages)
	return &messages, err
}

// GetChatScheduledMessages Returns all scheduled messages in a chat. The messages are returned in a reverse chronological order (i.e., in order of decreasing message_id)
// @param chatId Chat identifier
func (client *Client) GetChatScheduledMessages(chatId int64) (*Messages, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "getChatScheduledMessages",
		"chat_id": chatId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var messages Messages
	err = json.Unmarshal(result.Raw, &messages)
	return &messages, err
}

// SendMessageAlbum Sends 2-10 messages grouped together into an album. Currently, only audio, document, photo and video messages can be grouped into an album. Documents and audio files can be only grouped in an album with messages of the same type. Returns sent messages
// @param chatId Target chat
// @param messageThreadId If not 0, a message thread identifier in which the messages will be sent
// @param replyToMessageId Identifier of a message to reply to or 0
// @param options Options to be used to send the messages; pass null to use default options
// @param inputMessageContents Contents of messages to be sent. At most 10 messages can be added to an album
func (client *Client) SendMessageAlbum(chatId int64, messageThreadId int64, replyToMessageId int64, options *MessageSendOptions, inputMessageContents []InputMessageContent) (*Messages, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":                  "sendMessageAlbum",
		"chat_id":                chatId,
		"message_thread_id":      messageThreadId,
		"reply_to_message_id":    replyToMessageId,
		"options":                options,
		"input_message_contents": inputMessageContents,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var messages Messages
	err = json.Unmarshal(result.Raw, &messages)
	return &messages, err
}

// ForwardMessages Forwards previously sent messages. Returns the forwarded messages in the same order as the message identifiers passed in message_ids. If a message can't be forwarded, null will be returned instead of the message
// @param chatId Identifier of the chat to which to forward messages
// @param fromChatId Identifier of the chat from which to forward messages
// @param messageIds Identifiers of the messages to forward. Message identifiers must be in a strictly increasing order. At most 100 messages can be forwarded simultaneously
// @param options Options to be used to send the messages; pass null to use default options
// @param sendCopy If true, content of the messages will be copied without reference to the original sender. Always true if the messages are forwarded to a secret chat or are local
// @param removeCaption If true, media caption of message copies will be removed. Ignored if send_copy is false
// @param onlyPreview If true, messages will not be forwarded and instead fake messages will be returned
func (client *Client) ForwardMessages(chatId int64, fromChatId int64, messageIds []int64, options *MessageSendOptions, sendCopy bool, removeCaption bool, onlyPreview bool) (*Messages, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":          "forwardMessages",
		"chat_id":        chatId,
		"from_chat_id":   fromChatId,
		"message_ids":    messageIds,
		"options":        options,
		"send_copy":      sendCopy,
		"remove_caption": removeCaption,
		"only_preview":   onlyPreview,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var messages Messages
	err = json.Unmarshal(result.Raw, &messages)
	return &messages, err
}

// ResendMessages Resends messages which failed to send. Can be called only for messages for which messageSendingStateFailed.can_retry is true and after specified in messageSendingStateFailed.retry_after time passed. If a message is re-sent, the corresponding failed to send message is deleted. Returns the sent messages in the same order as the message identifiers passed in message_ids. If a message can't be re-sent, null will be returned instead of the message
// @param chatId Identifier of the chat to send messages
// @param messageIds Identifiers of the messages to resend. Message identifiers must be in a strictly increasing order
func (client *Client) ResendMessages(chatId int64, messageIds []int64) (*Messages, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":       "resendMessages",
		"chat_id":     chatId,
		"message_ids": messageIds,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var messages Messages
	err = json.Unmarshal(result.Raw, &messages)
	return &messages, err
}
