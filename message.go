package tdlib

import (
	"encoding/json"
	"fmt"
)

// Message Describes a message
type Message struct {
	tdCommon
	Id                        int64                   `json:"id"`                            // Message identifier; unique for the chat to which the message belongs
	SenderId                  MessageSender           `json:"sender_id"`                     // Identifier of the sender of the message
	ChatId                    int64                   `json:"chat_id"`                       // Chat identifier
	SendingState              MessageSendingState     `json:"sending_state"`                 // The sending state of the message; may be null
	SchedulingState           MessageSchedulingState  `json:"scheduling_state"`              // The scheduling state of the message; may be null
	IsOutgoing                bool                    `json:"is_outgoing"`                   // True, if the message is outgoing
	IsPinned                  bool                    `json:"is_pinned"`                     // True, if the message is pinned
	CanBeEdited               bool                    `json:"can_be_edited"`                 // True, if the message can be edited. For live location and poll messages this fields shows whether editMessageLiveLocation or stopPoll can be used with this message by the application
	CanBeForwarded            bool                    `json:"can_be_forwarded"`              // True, if the message can be forwarded
	CanBeSaved                bool                    `json:"can_be_saved"`                  // True, if content of the message can be saved locally or copied
	CanBeDeletedOnlyForSelf   bool                    `json:"can_be_deleted_only_for_self"`  // True, if the message can be deleted only for the current user while other users will continue to see it
	CanBeDeletedForAllUsers   bool                    `json:"can_be_deleted_for_all_users"`  // True, if the message can be deleted for all users
	CanGetStatistics          bool                    `json:"can_get_statistics"`            // True, if the message statistics are available
	CanGetMessageThread       bool                    `json:"can_get_message_thread"`        // True, if the message thread info is available
	CanGetViewers             bool                    `json:"can_get_viewers"`               // True, if chat members already viewed the message can be received through getMessageViewers
	CanGetMediaTimestampLinks bool                    `json:"can_get_media_timestamp_links"` // True, if media timestamp links can be generated for media timestamp entities in the message text, caption or web page description
	HasTimestampedMedia       bool                    `json:"has_timestamped_media"`         // True, if media timestamp entities refers to a media in this message as opposed to a media in the replied message
	IsChannelPost             bool                    `json:"is_channel_post"`               // True, if the message is a channel post. All messages to channels are channel posts, all other messages are not channel posts
	ContainsUnreadMention     bool                    `json:"contains_unread_mention"`       // True, if the message contains an unread mention for the current user
	Date                      int32                   `json:"date"`                          // Point in time (Unix timestamp) when the message was sent
	EditDate                  int32                   `json:"edit_date"`                     // Point in time (Unix timestamp) when the message was last edited
	ForwardInfo               *MessageForwardInfo     `json:"forward_info"`                  // Information about the initial message sender; may be null
	InteractionInfo           *MessageInteractionInfo `json:"interaction_info"`              // Information about interactions with the message; may be null
	ReplyInChatId             int64                   `json:"reply_in_chat_id"`              // If non-zero, the identifier of the chat to which the replied message belongs; Currently, only messages in the Replies chat can have different reply_in_chat_id and chat_id
	ReplyToMessageId          int64                   `json:"reply_to_message_id"`           // If non-zero, the identifier of the message this message is replying to; can be the identifier of a deleted message
	MessageThreadId           int64                   `json:"message_thread_id"`             // If non-zero, the identifier of the message thread the message belongs to; unique within the chat to which the message belongs
	Ttl                       int32                   `json:"ttl"`                           // For self-destructing messages, the message's TTL (Time To Live), in seconds; 0 if none. TDLib will send updateDeleteMessages or updateMessageContent once the TTL expires
	TtlExpiresIn              float64                 `json:"ttl_expires_in"`                // Time left before the message expires, in seconds. If the TTL timer isn't started yet, equals to the value of the ttl field
	ViaBotUserId              int64                   `json:"via_bot_user_id"`               // If non-zero, the user identifier of the bot through which this message was sent
	AuthorSignature           string                  `json:"author_signature"`              // For channel posts and anonymous group messages, optional author signature
	MediaAlbumId              JSONInt64               `json:"media_album_id"`                // Unique identifier of an album this message belongs to. Only audios, documents, photos and videos can be grouped together in albums
	RestrictionReason         string                  `json:"restriction_reason"`            // If non-empty, contains a human-readable description of the reason why access to this message must be restricted
	Content                   MessageContent          `json:"content"`                       // Content of the message
	ReplyMarkup               ReplyMarkup             `json:"reply_markup"`                  // Reply markup for the message; may be null
}

// MessageType return the string telegram-type of Message
func (message *Message) MessageType() string {
	return "message"
}

// NewMessage creates a new Message
//
// @param id Message identifier; unique for the chat to which the message belongs
// @param senderId Identifier of the sender of the message
// @param chatId Chat identifier
// @param sendingState The sending state of the message; may be null
// @param schedulingState The scheduling state of the message; may be null
// @param isOutgoing True, if the message is outgoing
// @param isPinned True, if the message is pinned
// @param canBeEdited True, if the message can be edited. For live location and poll messages this fields shows whether editMessageLiveLocation or stopPoll can be used with this message by the application
// @param canBeForwarded True, if the message can be forwarded
// @param canBeSaved True, if content of the message can be saved locally or copied
// @param canBeDeletedOnlyForSelf True, if the message can be deleted only for the current user while other users will continue to see it
// @param canBeDeletedForAllUsers True, if the message can be deleted for all users
// @param canGetStatistics True, if the message statistics are available
// @param canGetMessageThread True, if the message thread info is available
// @param canGetViewers True, if chat members already viewed the message can be received through getMessageViewers
// @param canGetMediaTimestampLinks True, if media timestamp links can be generated for media timestamp entities in the message text, caption or web page description
// @param hasTimestampedMedia True, if media timestamp entities refers to a media in this message as opposed to a media in the replied message
// @param isChannelPost True, if the message is a channel post. All messages to channels are channel posts, all other messages are not channel posts
// @param containsUnreadMention True, if the message contains an unread mention for the current user
// @param date Point in time (Unix timestamp) when the message was sent
// @param editDate Point in time (Unix timestamp) when the message was last edited
// @param forwardInfo Information about the initial message sender; may be null
// @param interactionInfo Information about interactions with the message; may be null
// @param replyInChatId If non-zero, the identifier of the chat to which the replied message belongs; Currently, only messages in the Replies chat can have different reply_in_chat_id and chat_id
// @param replyToMessageId If non-zero, the identifier of the message this message is replying to; can be the identifier of a deleted message
// @param messageThreadId If non-zero, the identifier of the message thread the message belongs to; unique within the chat to which the message belongs
// @param ttl For self-destructing messages, the message's TTL (Time To Live), in seconds; 0 if none. TDLib will send updateDeleteMessages or updateMessageContent once the TTL expires
// @param ttlExpiresIn Time left before the message expires, in seconds. If the TTL timer isn't started yet, equals to the value of the ttl field
// @param viaBotUserId If non-zero, the user identifier of the bot through which this message was sent
// @param authorSignature For channel posts and anonymous group messages, optional author signature
// @param mediaAlbumId Unique identifier of an album this message belongs to. Only audios, documents, photos and videos can be grouped together in albums
// @param restrictionReason If non-empty, contains a human-readable description of the reason why access to this message must be restricted
// @param content Content of the message
// @param replyMarkup Reply markup for the message; may be null
func NewMessage(id int64, senderId MessageSender, chatId int64, sendingState MessageSendingState, schedulingState MessageSchedulingState, isOutgoing bool, isPinned bool, canBeEdited bool, canBeForwarded bool, canBeSaved bool, canBeDeletedOnlyForSelf bool, canBeDeletedForAllUsers bool, canGetStatistics bool, canGetMessageThread bool, canGetViewers bool, canGetMediaTimestampLinks bool, hasTimestampedMedia bool, isChannelPost bool, containsUnreadMention bool, date int32, editDate int32, forwardInfo *MessageForwardInfo, interactionInfo *MessageInteractionInfo, replyInChatId int64, replyToMessageId int64, messageThreadId int64, ttl int32, ttlExpiresIn float64, viaBotUserId int64, authorSignature string, mediaAlbumId JSONInt64, restrictionReason string, content MessageContent, replyMarkup ReplyMarkup) *Message {
	messageTemp := Message{
		tdCommon:                  tdCommon{Type: "message"},
		Id:                        id,
		SenderId:                  senderId,
		ChatId:                    chatId,
		SendingState:              sendingState,
		SchedulingState:           schedulingState,
		IsOutgoing:                isOutgoing,
		IsPinned:                  isPinned,
		CanBeEdited:               canBeEdited,
		CanBeForwarded:            canBeForwarded,
		CanBeSaved:                canBeSaved,
		CanBeDeletedOnlyForSelf:   canBeDeletedOnlyForSelf,
		CanBeDeletedForAllUsers:   canBeDeletedForAllUsers,
		CanGetStatistics:          canGetStatistics,
		CanGetMessageThread:       canGetMessageThread,
		CanGetViewers:             canGetViewers,
		CanGetMediaTimestampLinks: canGetMediaTimestampLinks,
		HasTimestampedMedia:       hasTimestampedMedia,
		IsChannelPost:             isChannelPost,
		ContainsUnreadMention:     containsUnreadMention,
		Date:                      date,
		EditDate:                  editDate,
		ForwardInfo:               forwardInfo,
		InteractionInfo:           interactionInfo,
		ReplyInChatId:             replyInChatId,
		ReplyToMessageId:          replyToMessageId,
		MessageThreadId:           messageThreadId,
		Ttl:                       ttl,
		TtlExpiresIn:              ttlExpiresIn,
		ViaBotUserId:              viaBotUserId,
		AuthorSignature:           authorSignature,
		MediaAlbumId:              mediaAlbumId,
		RestrictionReason:         restrictionReason,
		Content:                   content,
		ReplyMarkup:               replyMarkup,
	}

	return &messageTemp
}

// UnmarshalJSON unmarshal to json
func (message *Message) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	tempObj := struct {
		tdCommon
		Id                        int64               `json:"id"`                            // Message identifier; unique for the chat to which the message belongs
		ChatId                    int64               `json:"chat_id"`                       // Chat identifier
		IsOutgoing                bool                `json:"is_outgoing"`                   // True, if the message is outgoing
		IsPinned                  bool                `json:"is_pinned"`                     // True, if the message is pinned
		CanBeEdited               bool                `json:"can_be_edited"`                 // True, if the message can be edited. For live location and poll messages this fields shows whether editMessageLiveLocation or stopPoll can be used with this message by the application
		CanBeForwarded            bool                `json:"can_be_forwarded"`              // True, if the message can be forwarded
		CanBeSaved                bool                `json:"can_be_saved"`                  // True, if content of the message can be saved locally or copied
		CanBeDeletedOnlyForSelf   bool                `json:"can_be_deleted_only_for_self"`  // True, if the message can be deleted only for the current user while other users will continue to see it
		CanBeDeletedForAllUsers   bool                `json:"can_be_deleted_for_all_users"`  // True, if the message can be deleted for all users
		CanGetStatistics          bool                `json:"can_get_statistics"`            // True, if the message statistics are available
		CanGetMessageThread       bool                `json:"can_get_message_thread"`        // True, if the message thread info is available
		CanGetViewers             bool                `json:"can_get_viewers"`               // True, if chat members already viewed the message can be received through getMessageViewers
		CanGetMediaTimestampLinks bool                `json:"can_get_media_timestamp_links"` // True, if media timestamp links can be generated for media timestamp entities in the message text, caption or web page description
		HasTimestampedMedia       bool                `json:"has_timestamped_media"`         // True, if media timestamp entities refers to a media in this message as opposed to a media in the replied message
		IsChannelPost             bool                `json:"is_channel_post"`               // True, if the message is a channel post. All messages to channels are channel posts, all other messages are not channel posts
		ContainsUnreadMention     bool                `json:"contains_unread_mention"`       // True, if the message contains an unread mention for the current user
		Date                      int32               `json:"date"`                          // Point in time (Unix timestamp) when the message was sent
		EditDate                  int32               `json:"edit_date"`                     // Point in time (Unix timestamp) when the message was last edited
		ForwardInfo               *MessageForwardInfo `json:"forward_info"`                  // Information about the initial message sender; may be null
		ReplyInChatId             int64               `json:"reply_in_chat_id"`              // If non-zero, the identifier of the chat to which the replied message belongs; Currently, only messages in the Replies chat can have different reply_in_chat_id and chat_id
		ReplyToMessageId          int64               `json:"reply_to_message_id"`           // If non-zero, the identifier of the message this message is replying to; can be the identifier of a deleted message
		MessageThreadId           int64               `json:"message_thread_id"`             // If non-zero, the identifier of the message thread the message belongs to; unique within the chat to which the message belongs
		Ttl                       int32               `json:"ttl"`                           // For self-destructing messages, the message's TTL (Time To Live), in seconds; 0 if none. TDLib will send updateDeleteMessages or updateMessageContent once the TTL expires
		TtlExpiresIn              float64             `json:"ttl_expires_in"`                // Time left before the message expires, in seconds. If the TTL timer isn't started yet, equals to the value of the ttl field
		ViaBotUserId              int64               `json:"via_bot_user_id"`               // If non-zero, the user identifier of the bot through which this message was sent
		AuthorSignature           string              `json:"author_signature"`              // For channel posts and anonymous group messages, optional author signature
		MediaAlbumId              JSONInt64           `json:"media_album_id"`                // Unique identifier of an album this message belongs to. Only audios, documents, photos and videos can be grouped together in albums
		RestrictionReason         string              `json:"restriction_reason"`            // If non-empty, contains a human-readable description of the reason why access to this message must be restricted

	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	message.tdCommon = tempObj.tdCommon
	message.Id = tempObj.Id
	message.ChatId = tempObj.ChatId
	message.IsOutgoing = tempObj.IsOutgoing
	message.IsPinned = tempObj.IsPinned
	message.CanBeEdited = tempObj.CanBeEdited
	message.CanBeForwarded = tempObj.CanBeForwarded
	message.CanBeSaved = tempObj.CanBeSaved
	message.CanBeDeletedOnlyForSelf = tempObj.CanBeDeletedOnlyForSelf
	message.CanBeDeletedForAllUsers = tempObj.CanBeDeletedForAllUsers
	message.CanGetStatistics = tempObj.CanGetStatistics
	message.CanGetMessageThread = tempObj.CanGetMessageThread
	message.CanGetViewers = tempObj.CanGetViewers
	message.CanGetMediaTimestampLinks = tempObj.CanGetMediaTimestampLinks
	message.HasTimestampedMedia = tempObj.HasTimestampedMedia
	message.IsChannelPost = tempObj.IsChannelPost
	message.ContainsUnreadMention = tempObj.ContainsUnreadMention
	message.Date = tempObj.Date
	message.EditDate = tempObj.EditDate
	message.ForwardInfo = tempObj.ForwardInfo
	message.ReplyInChatId = tempObj.ReplyInChatId
	message.ReplyToMessageId = tempObj.ReplyToMessageId
	message.MessageThreadId = tempObj.MessageThreadId
	message.Ttl = tempObj.Ttl
	message.TtlExpiresIn = tempObj.TtlExpiresIn
	message.ViaBotUserId = tempObj.ViaBotUserId
	message.AuthorSignature = tempObj.AuthorSignature
	message.MediaAlbumId = tempObj.MediaAlbumId
	message.RestrictionReason = tempObj.RestrictionReason

	fieldSenderId, _ := unmarshalMessageSender(objMap["sender_id"])
	message.SenderId = fieldSenderId

	// FYI: Arman92/go-tl-parser can't parse TDLib 1.7 new type `interaction_info`
	// and that go-tl-parser is HARD TO MAINTENANCE!!!
	// So we need convert `interaction_info` to Go struct without codegen.
	if objMap["interaction_info"] != nil {
		info, _ := objMap["interaction_info"].MarshalJSON()

		infoObj := struct {
			Type         string `json:"@type"`
			ViewCount    int32  `json:"view_count"`
			ForwardCount int32  `json:"forward_count"`
			ReplyInfo    struct {
				Type             string `json:"@type"`
				ReplyCount       int32  `json:"reply_count"`
				RecentReplierIds []struct {
					Type   string `json:"@type"`
					ChatID int64  `json:"chat_id"`
					UserID int64  `json:"user_id"`
				} `json:"recent_repliers_ids"`
				LastReadInboxMessageID  int64 `json:"last_read_inbox_message_id"`
				LastReadOutboxMessageID int64 `json:"last_read_outbox_message_id"`
				LastMessageID           int64 `json:"last_message_id"`
			} `json:"reply_info"`
		}{}

		err = json.Unmarshal(info, &infoObj)
		if err != nil {
			return err
		}

		message.InteractionInfo = NewMessageInteractionInfo(infoObj.ViewCount, infoObj.ForwardCount, nil)
		message.InteractionInfo.ReplyInfo = NewMessageReplyInfo(infoObj.ReplyInfo.ReplyCount, nil, infoObj.ReplyInfo.LastReadInboxMessageID, infoObj.ReplyInfo.LastReadOutboxMessageID, infoObj.ReplyInfo.LastMessageID)
		for _, reply := range infoObj.ReplyInfo.RecentReplierIds {
			switch reply.Type {
			case "messageSenderUser":
				{
					message.InteractionInfo.ReplyInfo.RecentReplierIds = append(message.InteractionInfo.ReplyInfo.RecentReplierIds, NewMessageSenderUser(reply.UserID))
				}
			case "messageSenderChat":
				{
					message.InteractionInfo.ReplyInfo.RecentReplierIds = append(message.InteractionInfo.ReplyInfo.RecentReplierIds, NewMessageSenderChat(reply.ChatID))
				}
			default:
				return fmt.Errorf("Error unmarshaling, unknown messageSender type:" + reply.Type)
			}
		}
	}

	fieldSendingState, _ := unmarshalMessageSendingState(objMap["sending_state"])
	message.SendingState = fieldSendingState

	fieldSchedulingState, _ := unmarshalMessageSchedulingState(objMap["scheduling_state"])
	message.SchedulingState = fieldSchedulingState

	fieldContent, _ := unmarshalMessageContent(objMap["content"])
	message.Content = fieldContent

	fieldReplyMarkup, _ := unmarshalReplyMarkup(objMap["reply_markup"])
	message.ReplyMarkup = fieldReplyMarkup

	return nil
}

// GetMessage Returns information about a message
// @param chatId Identifier of the chat the message belongs to
// @param messageId Identifier of the message to get
func (client *Client) GetMessage(chatId int64, messageId int64) (*Message, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":      "getMessage",
		"chat_id":    chatId,
		"message_id": messageId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var messageDummy Message
	err = json.Unmarshal(result.Raw, &messageDummy)
	return &messageDummy, err
}

// GetMessageLocally Returns information about a message, if it is available locally without sending network request. This is an offline request
// @param chatId Identifier of the chat the message belongs to
// @param messageId Identifier of the message to get
func (client *Client) GetMessageLocally(chatId int64, messageId int64) (*Message, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":      "getMessageLocally",
		"chat_id":    chatId,
		"message_id": messageId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var messageDummy Message
	err = json.Unmarshal(result.Raw, &messageDummy)
	return &messageDummy, err
}

// GetRepliedMessage Returns information about a message that is replied by a given message. Also returns the pinned message, the game message, and the invoice message for messages of the types messagePinMessage, messageGameScore, and messagePaymentSuccessful respectively
// @param chatId Identifier of the chat the message belongs to
// @param messageId Identifier of the reply message
func (client *Client) GetRepliedMessage(chatId int64, messageId int64) (*Message, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":      "getRepliedMessage",
		"chat_id":    chatId,
		"message_id": messageId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var messageDummy Message
	err = json.Unmarshal(result.Raw, &messageDummy)
	return &messageDummy, err
}

// GetChatPinnedMessage Returns information about a newest pinned message in the chat
// @param chatId Identifier of the chat the message belongs to
func (client *Client) GetChatPinnedMessage(chatId int64) (*Message, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "getChatPinnedMessage",
		"chat_id": chatId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var message Message
	err = json.Unmarshal(result.Raw, &message)
	return &message, err
}

// GetCallbackQueryMessage Returns information about a message with the callback button that originated a callback query; for bots only
// @param chatId Identifier of the chat the message belongs to
// @param messageId Message identifier
// @param callbackQueryId Identifier of the callback query
func (client *Client) GetCallbackQueryMessage(chatId int64, messageId int64, callbackQueryId JSONInt64) (*Message, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":             "getCallbackQueryMessage",
		"chat_id":           chatId,
		"message_id":        messageId,
		"callback_query_id": callbackQueryId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var messageDummy Message
	err = json.Unmarshal(result.Raw, &messageDummy)
	return &messageDummy, err
}

// GetChatMessageByDate Returns the last message sent in a chat no later than the specified date
// @param chatId Chat identifier
// @param date Point in time (Unix timestamp) relative to which to search for messages
func (client *Client) GetChatMessageByDate(chatId int64, date int32) (*Message, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "getChatMessageByDate",
		"chat_id": chatId,
		"date":    date,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var message Message
	err = json.Unmarshal(result.Raw, &message)
	return &message, err
}

// SendMessage Sends a message. Returns the sent message
// @param chatId Target chat
// @param messageThreadId If not 0, a message thread identifier in which the message will be sent
// @param replyToMessageId Identifier of the message to reply to or 0
// @param options Options to be used to send the message; pass null to use default options
// @param replyMarkup Markup for replying to the message; pass null if none; for bots only
// @param inputMessageContent The content of the message to be sent
func (client *Client) SendMessage(chatId int64, messageThreadId int64, replyToMessageId int64, options *MessageSendOptions, replyMarkup ReplyMarkup, inputMessageContent InputMessageContent) (*Message, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":                 "sendMessage",
		"chat_id":               chatId,
		"message_thread_id":     messageThreadId,
		"reply_to_message_id":   replyToMessageId,
		"options":               options,
		"reply_markup":          replyMarkup,
		"input_message_content": inputMessageContent,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var messageDummy Message
	err = json.Unmarshal(result.Raw, &messageDummy)
	return &messageDummy, err
}

// SendBotStartMessage Invites a bot to a chat (if it is not yet a member) and sends it the /start command. Bots can't be invited to a private chat other than the chat with the bot. Bots can't be invited to channels (although they can be added as admins) and secret chats. Returns the sent message
// @param botUserId Identifier of the bot
// @param chatId Identifier of the target chat
// @param parameter A hidden parameter sent to the bot for deep linking purposes (https://core.telegram.org/bots#deep-linking)
func (client *Client) SendBotStartMessage(botUserId int64, chatId int64, parameter string) (*Message, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":       "sendBotStartMessage",
		"bot_user_id": botUserId,
		"chat_id":     chatId,
		"parameter":   parameter,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var message Message
	err = json.Unmarshal(result.Raw, &message)
	return &message, err
}

// SendInlineQueryResultMessage Sends the result of an inline query as a message. Returns the sent message. Always clears a chat draft message
// @param chatId Target chat
// @param messageThreadId If not 0, a message thread identifier in which the message will be sent
// @param replyToMessageId Identifier of a message to reply to or 0
// @param options Options to be used to send the message; pass null to use default options
// @param queryId Identifier of the inline query
// @param resultId Identifier of the inline result
// @param hideViaBot If true, there will be no mention of a bot, via which the message is sent. Can be used only for bots GetOption("animation_search_bot_username"), GetOption("photo_search_bot_username") and GetOption("venue_search_bot_username")
func (client *Client) SendInlineQueryResultMessage(chatId int64, messageThreadId int64, replyToMessageId int64, options *MessageSendOptions, queryId JSONInt64, resultId string, hideViaBot bool) (*Message, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":               "sendInlineQueryResultMessage",
		"chat_id":             chatId,
		"message_thread_id":   messageThreadId,
		"reply_to_message_id": replyToMessageId,
		"options":             options,
		"query_id":            queryId,
		"result_id":           resultId,
		"hide_via_bot":        hideViaBot,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var messageDummy Message
	err = json.Unmarshal(result.Raw, &messageDummy)
	return &messageDummy, err
}

// AddLocalMessage Adds a local message to a chat. The message is persistent across application restarts only if the message database is used. Returns the added message
// @param chatId Target chat
// @param senderId Identifier of the sender of the message
// @param replyToMessageId Identifier of the message to reply to or 0
// @param disableNotification Pass true to disable notification for the message
// @param inputMessageContent The content of the message to be added
func (client *Client) AddLocalMessage(chatId int64, senderId MessageSender, replyToMessageId int64, disableNotification bool, inputMessageContent InputMessageContent) (*Message, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":                 "addLocalMessage",
		"chat_id":               chatId,
		"sender_id":             senderId,
		"reply_to_message_id":   replyToMessageId,
		"disable_notification":  disableNotification,
		"input_message_content": inputMessageContent,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var messageDummy Message
	err = json.Unmarshal(result.Raw, &messageDummy)
	return &messageDummy, err
}

// EditMessageText Edits the text of a message (or a text of a game message). Returns the edited message after the edit is completed on the server side
// @param chatId The chat the message belongs to
// @param messageId Identifier of the message
// @param replyMarkup The new message reply markup; pass null if none; for bots only
// @param inputMessageContent New text content of the message. Must be of type inputMessageText
func (client *Client) EditMessageText(chatId int64, messageId int64, replyMarkup ReplyMarkup, inputMessageContent InputMessageContent) (*Message, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":                 "editMessageText",
		"chat_id":               chatId,
		"message_id":            messageId,
		"reply_markup":          replyMarkup,
		"input_message_content": inputMessageContent,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var messageDummy Message
	err = json.Unmarshal(result.Raw, &messageDummy)
	return &messageDummy, err
}

// EditMessageLiveLocation Edits the message content of a live location. Messages can be edited for a limited period of time specified in the live location. Returns the edited message after the edit is completed on the server side
// @param chatId The chat the message belongs to
// @param messageId Identifier of the message
// @param replyMarkup The new message reply markup; pass null if none; for bots only
// @param location New location content of the message; pass null to stop sharing the live location
// @param heading The new direction in which the location moves, in degrees; 1-360. Pass 0 if unknown
// @param proximityAlertRadius The new maximum distance for proximity alerts, in meters (0-100000). Pass 0 if the notification is disabled
func (client *Client) EditMessageLiveLocation(chatId int64, messageId int64, replyMarkup ReplyMarkup, location *Location, heading int32, proximityAlertRadius int32) (*Message, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":                  "editMessageLiveLocation",
		"chat_id":                chatId,
		"message_id":             messageId,
		"reply_markup":           replyMarkup,
		"location":               location,
		"heading":                heading,
		"proximity_alert_radius": proximityAlertRadius,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var messageDummy Message
	err = json.Unmarshal(result.Raw, &messageDummy)
	return &messageDummy, err
}

// EditMessageMedia Edits the content of a message with an animation, an audio, a document, a photo or a video, including message caption. If only the caption needs to be edited, use editMessageCaption instead. The media can't be edited if the message was set to self-destruct or to a self-destructing media. The type of message content in an album can't be changed with exception of replacing a photo with a video or vice versa. Returns the edited message after the edit is completed on the server side
// @param chatId The chat the message belongs to
// @param messageId Identifier of the message
// @param replyMarkup The new message reply markup; pass null if none; for bots only
// @param inputMessageContent New content of the message. Must be one of the following types: inputMessageAnimation, inputMessageAudio, inputMessageDocument, inputMessagePhoto or inputMessageVideo
func (client *Client) EditMessageMedia(chatId int64, messageId int64, replyMarkup ReplyMarkup, inputMessageContent InputMessageContent) (*Message, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":                 "editMessageMedia",
		"chat_id":               chatId,
		"message_id":            messageId,
		"reply_markup":          replyMarkup,
		"input_message_content": inputMessageContent,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var messageDummy Message
	err = json.Unmarshal(result.Raw, &messageDummy)
	return &messageDummy, err
}

// EditMessageCaption Edits the message content caption. Returns the edited message after the edit is completed on the server side
// @param chatId The chat the message belongs to
// @param messageId Identifier of the message
// @param replyMarkup The new message reply markup; pass null if none; for bots only
// @param caption New message content caption; 0-GetOption("message_caption_length_max") characters; pass null to remove caption
func (client *Client) EditMessageCaption(chatId int64, messageId int64, replyMarkup ReplyMarkup, caption *FormattedText) (*Message, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":        "editMessageCaption",
		"chat_id":      chatId,
		"message_id":   messageId,
		"reply_markup": replyMarkup,
		"caption":      caption,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var messageDummy Message
	err = json.Unmarshal(result.Raw, &messageDummy)
	return &messageDummy, err
}

// EditMessageReplyMarkup Edits the message reply markup; for bots only. Returns the edited message after the edit is completed on the server side
// @param chatId The chat the message belongs to
// @param messageId Identifier of the message
// @param replyMarkup The new message reply markup; pass null if none
func (client *Client) EditMessageReplyMarkup(chatId int64, messageId int64, replyMarkup ReplyMarkup) (*Message, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":        "editMessageReplyMarkup",
		"chat_id":      chatId,
		"message_id":   messageId,
		"reply_markup": replyMarkup,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var messageDummy Message
	err = json.Unmarshal(result.Raw, &messageDummy)
	return &messageDummy, err
}

// SetGameScore Updates the game score of the specified user in the game; for bots only
// @param chatId The chat to which the message with the game belongs
// @param messageId Identifier of the message
// @param editMessage True, if the message needs to be edited
// @param userId User identifier
// @param score The new score
// @param force Pass true to update the score even if it decreases. If the score is 0, the user will be deleted from the high score table
func (client *Client) SetGameScore(chatId int64, messageId int64, editMessage bool, userId int64, score int32, force bool) (*Message, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":        "setGameScore",
		"chat_id":      chatId,
		"message_id":   messageId,
		"edit_message": editMessage,
		"user_id":      userId,
		"score":        score,
		"force":        force,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var messageDummy Message
	err = json.Unmarshal(result.Raw, &messageDummy)
	return &messageDummy, err
}
